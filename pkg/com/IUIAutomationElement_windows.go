// +build windows

package com

import (
	"syscall"
	"unsafe"

	"github.com/moutend/CoreServer/pkg/types"

	"github.com/go-ole/go-ole"
)

func uiaeSetFocus(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeGetRuntimeId(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeFindFirst(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeFindAll(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeFindFirstBuildCache(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeFindAllBuildCache(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeBuildUpdatedCache(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeGetCurrentPropertyValue(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeGetCurrentPropertyValueEx(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeGetCachedPropertyValue(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeGetCachedPropertyValueEx(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeGetCurrentPatternAs(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeGetCachedPatternAs(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeGetCurrentPattern(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeGetCachedPattern(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeGetCachedParent(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeGetCachedChildren(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentProcessId(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentControlType(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentLocalizedControlType(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentName(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentAcceleratorKey(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentAccessKey(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentHasKeyboardFocus(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentIsKeyboardFocusable(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentIsEnabled(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentAutomationId(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentClassName(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentHelpText(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentCulture(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentIsControlElement(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentIsContentElement(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentIsPassword(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentNativeWindowHandle(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentItemType(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentIsOffscreen(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentOrientation(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentFrameworkId(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentIsRequiredForForm(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentItemStatus(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentBoundingRectangle(v *IUIAutomationElement) (*types.RECT, error) {
	var rect types.RECT

	hr, _, _ := syscall.Syscall(
		v.VTable().CurrentBoundingRectangle,
		1,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&rect)),
		0)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return &rect, nil
}

func uiaeCurrentLabeledBy(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentAriaRole(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentAriaProperties(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentIsDataValidForForm(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentControllerFor(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentDescribedBy(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentFlowsTo(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCurrentProviderDescription(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedProcessId(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedControlType(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedLocalizedControlType(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedName(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedAcceleratorKey(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedAccessKey(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedHasKeyboardFocus(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedIsKeyboardFocusable(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedIsEnabled(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedAutomationId(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedClassName(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedHelpText(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedCulture(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedIsControlElement(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedIsContentElement(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedIsPassword(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedNativeWindowHandle(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedItemType(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedIsOffscreen(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedOrientation(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedFrameworkId(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedIsRequiredForForm(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedItemStatus(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedBoundingRectangle(v *IUIAutomationElement) (*types.RECT, error) {
	var rect types.RECT

	hr, _, _ := syscall.Syscall(
		v.VTable().CachedBoundingRectangle,
		1,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&rect)),
		0)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return &rect, nil
}

func uiaeCachedLabeledBy(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedAriaRole(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedAriaProperties(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedIsDataValidForForm(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedControllerFor(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedDescribedBy(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedFlowsTo(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeCachedProviderDescription(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}

func uiaeGetClickablePoint(v *IUIAutomationElement) error {
	return ole.NewError(ole.E_NOTIMPL)
}
