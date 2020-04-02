// +build windows

package com

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func csStart(v *ICoreServer) error {
	hr, _, _ := syscall.Syscall(
		v.VTable().Start,
		0,
		uintptr(unsafe.Pointer(v)),
		0,
		0)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func csStop(v *ICoreServer) error {
	hr, _, _ := syscall.Syscall(
		v.VTable().Stop,
		0,
		uintptr(unsafe.Pointer(v)),
		0,
		0)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func csSetMSAAEventHandler(v *ICoreServer, handleFunc MSAAEventHandler) error {
	hr, _, _ := syscall.Syscall(
		v.VTable().SetMSAAEventHandler,
		2,
		uintptr(unsafe.Pointer(v)),
		syscall.NewCallback(handleFunc),
		0)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func csSetUIAEventHandler(v *ICoreServer, handleFunc UIAEventHandler) error {
	hr, _, _ := syscall.Syscall(
		v.VTable().SetUIAEventHandler,
		2,
		uintptr(unsafe.Pointer(v)),
		syscall.NewCallback(handleFunc),
		0)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}
