// +build windows

package com

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/moutend/CoreServer/pkg/types"
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

func csGetIUIAutomationElement(v *ICoreServer, direction types.TreeWalkerDirection, pRootElement uintptr) (pElement *IUIAutomationElement, err error) {
	hr, _, _ := syscall.Syscall6(
		v.VTable().GetIUIAutomationElement,
		4,
		uintptr(unsafe.Pointer(v)),
		uintptr(direction),
		pRootElement,
		uintptr(unsafe.Pointer(&pElement)),
		0,
		0)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return pElement, nil
}

func csUpdateTreeWalker(v *ICoreServer) error {
	hr, _, _ := syscall.Syscall(
		v.VTable().UpdateTreeWalker,
		1,
		uintptr(unsafe.Pointer(v)),
		0,
		0)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}
