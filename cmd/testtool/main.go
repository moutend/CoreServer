package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"unsafe"

	"github.com/moutend/CoreServer/pkg/types"

	"github.com/moutend/CoreServer/pkg/com"

	"github.com/go-ole/go-ole"
)

func main() {
	log.SetPrefix("error: ")
	log.SetFlags(0)

	if err := run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(args []string) error {
	if err := ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED); err != nil {
		return err
	}
	defer ole.CoUninitialize()

	var foo *com.ICoreServer

	if err := com.CoCreateInstance(com.CLSID_CoreServer, 0, com.CLSCTX_ALL, com.IID_ICoreServer, &foo); err != nil {
		return err
	}

	fmt.Println("@@@instance created")

	defer foo.Release()

	err := foo.Start()
	fmt.Println("Called ICoreServer::Start", err)

	foo.SetIUIEventHandler(func(eventId types.UIAEvent, pInterface uintptr) int64 {
		e := (*com.IUIAutomationElement)(unsafe.Pointer(pInterface))

		rect, err := e.CurrentBoundingRectangle()

		if err != nil {
			fmt.Println("@@@err", err)
			return 0
		}
		if rect.IsZero() {
			fmt.Println("@@@skipped")
			return 0
		}

		name, _ := e.CurrentName()
		className, _ := e.CurrentClassName()
		framework, _ := e.CurrentFrameworkId()
		itemType, _ := e.CurrentItemType()
		ariaRole, _ := e.CurrentAriaRole()
		ariaProperties, _ := e.CurrentAriaProperties()
		fmt.Printf("@@@Event:%q,Name:%q,ClassName:%q,Framework:%q,ItemType:%q,AriaRole:%q,AriaProperties:%q\n", eventId, name, className, framework, itemType, ariaRole, ariaProperties)

		return 0
	})

	time.Sleep(30 * time.Second)

	err = foo.Stop()
	fmt.Println("Called Foo::Stop()")

	return nil
}
