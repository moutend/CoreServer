package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"unsafe"

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

	foo.SetUIEventHandler(func(eventId int64, eAPI int64, pInterface uintptr) int64 {
		fmt.Println("@@@received", eventId, eAPI)
		uiae := &com.IUIAutomationElement{}
		uiae.RawVTable = *(*(***interface{})(unsafe.Pointer(pInterface)))

		rect, err := uiae.CurrentBoundingRectangle()

		fmt.Println("@@@rect", rect, err)

		return 0
	})

	time.Sleep(30 * time.Second)

	err = foo.Stop()
	fmt.Println("Called Foo::Stop()")

	return nil
}
