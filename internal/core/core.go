package core

import (
	"log"

	"github.com/go-ole/go-ole"
	"github.com/moutend/CoreServer/pkg/com"
)

var (
	isRunning bool
	server    *com.ICoreServer
)

func Setup() error {
	if isRunning {
		return nil
	}
	if err := ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED); err != nil {
		return err
	}
	if err := com.CoCreateInstance(com.CLSID_CoreServer, 0, com.CLSCTX_ALL, com.IID_ICoreServer, &server); err != nil {
		return err
	}
	if err := server.Start(); err != nil {
		return err
	}

	isRunning = true

	log.Println("core: Setup() is done")

	return nil
}

func Teardown() error {
	server.Stop()
	server.Release()
	ole.CoUninitialize()

	isRunning = false

	log.Println("core: Teardown() is done")

	return nil
}

func SetUIAEventHandler(fn com.UIAEventHandler) {
	server.SetUIAEventHandler(fn)
}
