// +build !windows

package com

import (
	"github.com/go-ole/go-ole"
)

func csStart(v *ICoreServer) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func csStop(v *ICoreServer) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func csSetUIEventHandler(v *ICoreServer, handleFunc UIEventHandler) error {
	return ole.NewError(ole.E_NOTIMPL)
}
