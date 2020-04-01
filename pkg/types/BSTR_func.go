// +build !windows

package types

import "github.com/go-ole/go-ole"

func bstrString(p BSTR) (string, error) {
	return "", ole.NewError(ole.E_NOTIMPL)
}
