package app

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"time"
	"unsafe"

	"github.com/moutend/CoreServer/internal/core"
	"github.com/moutend/CoreServer/internal/util"
	"github.com/moutend/CoreServer/pkg/com"
	"github.com/moutend/CoreServer/pkg/types"

	"github.com/go-chi/chi"
	"github.com/moutend/CoreServer/internal/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var RootCommand = &cobra.Command{
	Use:  "todoapp",
	RunE: rootRunE,
}

func rootRunE(cmd *cobra.Command, args []string) error {
	if path, _ := cmd.Flags().GetString("config"); path != "" {
		viper.SetConfigFile(path)
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	if err := core.Setup(); err != nil {
		return err
	}

	defer core.Teardown()

	core.SetUIAEventHandler(func(eventId types.UIAEvent, pInterface uintptr) int64 {
		go http.Post("http://192.168.1.102:7902/v1/audio", "application/json", bytes.NewBufferString(`{"isForcePush":true,"commands": [{"type": 1, "value":10}]}`))
		if res, err := core.GetIUIAutomationElement(types.TW_NEXT, pInterface); err != nil {
			log.Println("@@@err", err)
		} else if res != nil {
			name, _ := res.CurrentName()
			log.Println("@@@next", name)
		}

		e := (*com.IUIAutomationElement)(unsafe.Pointer(pInterface))

		name, _ := e.CurrentName()

		log.Println("@@@current name", name)

		rect, err := e.CurrentBoundingRectangle()

		if err != nil {
			fmt.Println("@@@err", err)
			return 0
		}
		if rect.IsZero() {
			fmt.Println("@@@skipped")
			return 0
		}

		// name, _ := e.CurrentName()
		className, _ := e.CurrentClassName()
		framework, _ := e.CurrentFrameworkId()
		itemType, _ := e.CurrentItemType()
		ariaRole, _ := e.CurrentAriaRole()
		ariaProperties, _ := e.CurrentAriaProperties()
		return 0
		fmt.Printf("@@@Event:%q,Name:%q,ClassName:%q,Framework:%q,ItemType:%q,AriaRole:%q,AriaProperties:%q\n", eventId, name, className, framework, itemType, ariaRole, ariaProperties)

		return 0
	})

	router := chi.NewRouter()
	api.Setup(router)

	address := "localhost:7903"

	if a := viper.GetString("server.address"); a != "" {
		address = a
	}

	u, err := user.Current()

	if err != nil {
		return err
	}

	rand.Seed(time.Now().Unix())
	p := make([]byte, 16)

	if _, err := rand.Read(p); err != nil {
		return err
	}

	fileName := fmt.Sprintf("CoreServer-%s.txt", hex.EncodeToString(p))
	outputPath := filepath.Join(u.HomeDir, "AppData", "Roaming", "ScreenReaderX", "SystemLog", fileName)
	os.MkdirAll(filepath.Dir(outputPath), 0755)

	output := util.NewBackgroundWriter(outputPath)
	defer output.Close()

	log.SetFlags(log.Ldate | log.Ltime | log.LUTC | log.Llongfile)
	log.SetOutput(output)

	log.Printf("Listening on %s\n", address)

	return http.ListenAndServe(address, router)
}

func init() {
	RootCommand.PersistentFlags().StringP("config", "c", "", "Path to configuration file")
}
