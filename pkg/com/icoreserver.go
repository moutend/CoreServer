package com

import (
	"unsafe"

	"github.com/go-ole/go-ole"
)

type UIEventHandler func(eventId, eAPI int64, pInterface uintptr) int64

type ICoreServer struct {
	ole.IDispatch
}

type ICoreServerVtbl struct {
	ole.IDispatchVtbl
	Start             uintptr
	Stop              uintptr
	SetUIEventHandler uintptr
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

func (v *ICoreServer) SetUIEventHandler(handleFunc UIEventHandler) error {
	return csSetUIEventHandler(v, handleFunc)
}
