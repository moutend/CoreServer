package com

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type UIEventHandler func(eventId int64, pInterface uintptr) int64

type ICoreServer struct {
	ole.IDispatch
}

type ICoreServerVtbl struct {
	ole.IDispatchVtbl
	Start              uintptr
	Stop               uintptr
	SetIAEventHandler  uintptr
	SetIUIEventHandler uintptr
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

func (v *ICoreServer) SetIUIEventHandler(handleFunc UIEventHandler) error {
	return csSetIUIEventHandler(v, handleFunc)
}

func (v *ICoreServer) SetIAEventHandler(handleFunc UIEventHandler) error {
	return csSetIAEventHandler(v, handleFunc)
}
