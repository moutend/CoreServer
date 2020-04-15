package com

import (
	"unsafe"

	"github.com/moutend/CoreServer/pkg/types"

	"github.com/go-ole/go-ole"
)

type MSAAEventHandler func(eventId types.MSAAEvent, pInterface uintptr) int64

type UIAEventHandler func(eventId types.UIAEvent, pInterface uintptr) int64

type ICoreServer struct {
	ole.IDispatch
}

type ICoreServerVtbl struct {
	ole.IDispatchVtbl
	Start               uintptr
	Stop                uintptr
	SetMSAAEventHandler uintptr
	SetUIAEventHandler  uintptr
}

func (v *ICoreServer) VTable() *ICoreServerVtbl {
	return (*ICoreServerVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *ICoreServer) Start() error {
	return csStart(v)
}

func (v *ICoreServer) Stop() error {
	return csStop(v)
}

func (v *ICoreServer) SetMSAAEventHandler(handleFunc MSAAEventHandler) error {
	return csSetMSAAEventHandler(v, handleFunc)
}

func (v *ICoreServer) SetUIAEventHandler(handleFunc UIAEventHandler) error {
	return csSetUIAEventHandler(v, handleFunc)
}
