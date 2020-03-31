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
		fmt.Printf("@@@received %x\n", pInterface)
		uiae := &com.IUIAutomationElement{}
		vtblPtr := *(*uintptr)(unsafe.Pointer(pInterface))
		uiae.RawVTable = (*interface{})(unsafe.Pointer(vtblPtr))

		err := uiae.GetRuntimeId()
		fmt.Println("@@@bounding", err)

		res := uiae.Release()
		fmt.Println("@@@Release", res)

		return 0
	})

	time.Sleep(30 * time.Second)

	err = foo.Stop()
	fmt.Println("Called Foo::Stop()")

	return nil
}
