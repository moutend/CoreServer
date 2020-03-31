package com

import (
	"unsafe"

	"github.com/moutend/CoreServer/pkg/types"

	"github.com/go-ole/go-ole"
)

type IUIAutomationElement struct {
	ole.IUnknown
}

type IUIAutomationElementVtbl struct {
	ole.IUnknownVtbl
	SetFocus                    uintptr
	GetRuntimeId                uintptr
	FindFirst                   uintptr
	FindAll                     uintptr
	FindFirstBuildCache         uintptr
	FindAllBuildCache           uintptr
	BuildUpdatedCache           uintptr
	GetCurrentPropertyValue     uintptr
	GetCurrentPropertyValueEx   uintptr
	GetCachedPropertyValue      uintptr
	GetCachedPropertyValueEx    uintptr
	GetCurrentPatternAs         uintptr
	GetCachedPatternAs          uintptr
	GetCurrentPattern           uintptr
	GetCachedPattern            uintptr
	GetCachedParent             uintptr
	GetCachedChildren           uintptr
	CurrentProcessId            uintptr
	CurrentControlType          uintptr
	CurrentLocalizedControlType uintptr
	CurrentName                 uintptr
	CurrentAcceleratorKey       uintptr
	CurrentAccessKey            uintptr
	CurrentHasKeyboardFocus     uintptr
	CurrentIsKeyboardFocusable  uintptr
	CurrentIsEnabled            uintptr
	CurrentAutomationId         uintptr
	CurrentClassName            uintptr
	CurrentHelpText             uintptr
	CurrentCulture              uintptr
	CurrentIsControlElement     uintptr
	CurrentIsContentElement     uintptr
	CurrentIsPassword           uintptr
	CurrentNativeWindowHandle   uintptr
	CurrentItemType             uintptr
	CurrentIsOffscreen          uintptr
	CurrentOrientation          uintptr
	CurrentFrameworkId          uintptr
	CurrentIsRequiredForForm    uintptr
	CurrentItemStatus           uintptr
	CurrentBoundingRectangle    uintptr
	CurrentLabeledBy            uintptr
	CurrentAriaRole             uintptr
	CurrentAriaProperties       uintptr
	CurrentIsDataValidForForm   uintptr
	CurrentControllerFor        uintptr
	CurrentDescribedBy          uintptr
	CurrentFlowsTo              uintptr
	CurrentProviderDescription  uintptr
	CachedProcessId             uintptr
	CachedControlType           uintptr
	CachedLocalizedControlType  uintptr
	CachedName                  uintptr
	CachedAcceleratorKey        uintptr
	CachedAccessKey             uintptr
	CachedHasKeyboardFocus      uintptr
	CachedIsKeyboardFocusable   uintptr
	CachedIsEnabled             uintptr
	CachedAutomationId          uintptr
	CachedClassName             uintptr
	CachedHelpText              uintptr
	CachedCulture               uintptr
	CachedIsControlElement      uintptr
	CachedIsContentElement      uintptr
	CachedIsPassword            uintptr
	CachedNativeWindowHandle    uintptr
	CachedItemType              uintptr
	CachedIsOffscreen           uintptr
	CachedOrientation           uintptr
	CachedFrameworkId           uintptr
	CachedIsRequiredForForm     uintptr
	CachedItemStatus            uintptr
	CachedBoundingRectangle     uintptr
	CachedLabeledBy             uintptr
	CachedAriaRole              uintptr
	CachedAriaProperties        uintptr
	CachedIsDataValidForForm    uintptr
	CachedControllerFor         uintptr
	CachedDescribedBy           uintptr
	CachedFlowsTo               uintptr
	CachedProviderDescription   uintptr
	GetClickablePoint           uintptr
}

func (v *IUIAutomationElement) VTable() *IUIAutomationElementVtbl {
	return (*IUIAutomationElementVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IUIAutomationElement) SetFocus() error {
	return uiaeSetFocus(v)
}

func (v *IUIAutomationElement) GetRuntimeId() error {
	return uiaeGetRuntimeId(v)
}

func (v *IUIAutomationElement) FindFirst() error {
	return uiaeFindFirst(v)
}

func (v *IUIAutomationElement) FindAll() error {
	return uiaeFindAll(v)
}

func (v *IUIAutomationElement) FindFirstBuildCache() error {
	return uiaeFindFirstBuildCache(v)
}

func (v *IUIAutomationElement) FindAllBuildCache() error {
	return uiaeFindAllBuildCache(v)
}

func (v *IUIAutomationElement) BuildUpdatedCache() error {
	return uiaeBuildUpdatedCache(v)
}

func (v *IUIAutomationElement) GetCurrentPropertyValue() error {
	return uiaeGetCurrentPropertyValue(v)
}

func (v *IUIAutomationElement) GetCurrentPropertyValueEx() error {
	return uiaeGetCurrentPropertyValueEx(v)
}

func (v *IUIAutomationElement) GetCachedPropertyValue() error {
	return uiaeGetCachedPropertyValue(v)
}

func (v *IUIAutomationElement) GetCachedPropertyValueEx() error {
	return uiaeGetCachedPropertyValueEx(v)
}

func (v *IUIAutomationElement) GetCurrentPatternAs() error {
	return uiaeGetCurrentPatternAs(v)
}

func (v *IUIAutomationElement) GetCachedPatternAs() error {
	return uiaeGetCachedPatternAs(v)
}

func (v *IUIAutomationElement) GetCurrentPattern() error {
	return uiaeGetCurrentPattern(v)
}

func (v *IUIAutomationElement) GetCachedPattern() error {
	return uiaeGetCachedPattern(v)
}

func (v *IUIAutomationElement) GetCachedParent() error {
	return uiaeGetCachedParent(v)
}

func (v *IUIAutomationElement) GetCachedChildren() error {
	return uiaeGetCachedChildren(v)
}

func (v *IUIAutomationElement) CurrentProcessId() error {
	return uiaeCurrentProcessId(v)
}

func (v *IUIAutomationElement) CurrentControlType() error {
	return uiaeCurrentControlType(v)
}

func (v *IUIAutomationElement) CurrentLocalizedControlType() error {
	return uiaeCurrentLocalizedControlType(v)
}

func (v *IUIAutomationElement) CurrentName() error {
	return uiaeCurrentName(v)
}

func (v *IUIAutomationElement) CurrentAcceleratorKey() error {
	return uiaeCurrentAcceleratorKey(v)
}

func (v *IUIAutomationElement) CurrentAccessKey() error {
	return uiaeCurrentAccessKey(v)
}

func (v *IUIAutomationElement) CurrentHasKeyboardFocus() error {
	return uiaeCurrentHasKeyboardFocus(v)
}

func (v *IUIAutomationElement) CurrentIsKeyboardFocusable() error {
	return uiaeCurrentIsKeyboardFocusable(v)
}

func (v *IUIAutomationElement) CurrentIsEnabled() error {
	return uiaeCurrentIsEnabled(v)
}

func (v *IUIAutomationElement) CurrentAutomationId() error {
	return uiaeCurrentAutomationId(v)
}

func (v *IUIAutomationElement) CurrentClassName() error {
	return uiaeCurrentClassName(v)
}

func (v *IUIAutomationElement) CurrentHelpText() error {
	return uiaeCurrentHelpText(v)
}

func (v *IUIAutomationElement) CurrentCulture() error {
	return uiaeCurrentCulture(v)
}

func (v *IUIAutomationElement) CurrentIsControlElement() error {
	return uiaeCurrentIsControlElement(v)
}

func (v *IUIAutomationElement) CurrentIsContentElement() error {
	return uiaeCurrentIsContentElement(v)
}

func (v *IUIAutomationElement) CurrentIsPassword() error {
	return uiaeCurrentIsPassword(v)
}

func (v *IUIAutomationElement) CurrentNativeWindowHandle() error {
	return uiaeCurrentNativeWindowHandle(v)
}

func (v *IUIAutomationElement) CurrentItemType() error {
	return uiaeCurrentItemType(v)
}

func (v *IUIAutomationElement) CurrentIsOffscreen() error {
	return uiaeCurrentIsOffscreen(v)
}

func (v *IUIAutomationElement) CurrentOrientation() error {
	return uiaeCurrentOrientation(v)
}

func (v *IUIAutomationElement) CurrentFrameworkId() error {
	return uiaeCurrentFrameworkId(v)
}

func (v *IUIAutomationElement) CurrentIsRequiredForForm() error {
	return uiaeCurrentIsRequiredForForm(v)
}

func (v *IUIAutomationElement) CurrentItemStatus() error {
	return uiaeCurrentItemStatus(v)
}

func (v *IUIAutomationElement) CurrentBoundingRectangle() (*types.RECT, error) {
	return uiaeCurrentBoundingRectangle(v)
}

func (v *IUIAutomationElement) CurrentLabeledBy() error {
	return uiaeCurrentLabeledBy(v)
}

func (v *IUIAutomationElement) CurrentAriaRole() error {
	return uiaeCurrentAriaRole(v)
}

func (v *IUIAutomationElement) CurrentAriaProperties() error {
	return uiaeCurrentAriaProperties(v)
}

func (v *IUIAutomationElement) CurrentIsDataValidForForm() error {
	return uiaeCurrentIsDataValidForForm(v)
}

func (v *IUIAutomationElement) CurrentControllerFor() error {
	return uiaeCurrentControllerFor(v)
}

func (v *IUIAutomationElement) CurrentDescribedBy() error {
	return uiaeCurrentDescribedBy(v)
}

func (v *IUIAutomationElement) CurrentFlowsTo() error {
	return uiaeCurrentFlowsTo(v)
}

func (v *IUIAutomationElement) CurrentProviderDescription() error {
	return uiaeCurrentProviderDescription(v)
}

func (v *IUIAutomationElement) CachedProcessId() error {
	return uiaeCachedProcessId(v)
}

func (v *IUIAutomationElement) CachedControlType() error {
	return uiaeCachedControlType(v)
}

func (v *IUIAutomationElement) CachedLocalizedControlType() error {
	return uiaeCachedLocalizedControlType(v)
}

func (v *IUIAutomationElement) CachedName() error {
	return uiaeCachedName(v)
}

func (v *IUIAutomationElement) CachedAcceleratorKey() error {
	return uiaeCachedAcceleratorKey(v)
}

func (v *IUIAutomationElement) CachedAccessKey() error {
	return uiaeCachedAccessKey(v)
}

func (v *IUIAutomationElement) CachedHasKeyboardFocus() error {
	return uiaeCachedHasKeyboardFocus(v)
}

func (v *IUIAutomationElement) CachedIsKeyboardFocusable() error {
	return uiaeCachedIsKeyboardFocusable(v)
}

func (v *IUIAutomationElement) CachedIsEnabled() error {
	return uiaeCachedIsEnabled(v)
}

func (v *IUIAutomationElement) CachedAutomationId() error {
	return uiaeCachedAutomationId(v)
}

func (v *IUIAutomationElement) CachedClassName() error {
	return uiaeCachedClassName(v)
}

func (v *IUIAutomationElement) CachedHelpText() error {
	return uiaeCachedHelpText(v)
}

func (v *IUIAutomationElement) CachedCulture() error {
	return uiaeCachedCulture(v)
}

func (v *IUIAutomationElement) CachedIsControlElement() error {
	return uiaeCachedIsControlElement(v)
}

func (v *IUIAutomationElement) CachedIsContentElement() error {
	return uiaeCachedIsContentElement(v)
}

func (v *IUIAutomationElement) CachedIsPassword() error {
	return uiaeCachedIsPassword(v)
}

func (v *IUIAutomationElement) CachedNativeWindowHandle() error {
	return uiaeCachedNativeWindowHandle(v)
}

func (v *IUIAutomationElement) CachedItemType() error {
	return uiaeCachedItemType(v)
}

func (v *IUIAutomationElement) CachedIsOffscreen() error {
	return uiaeCachedIsOffscreen(v)
}

func (v *IUIAutomationElement) CachedOrientation() error {
	return uiaeCachedOrientation(v)
}

func (v *IUIAutomationElement) CachedFrameworkId() error {
	return uiaeCachedFrameworkId(v)
}

func (v *IUIAutomationElement) CachedIsRequiredForForm() error {
	return uiaeCachedIsRequiredForForm(v)
}

func (v *IUIAutomationElement) CachedItemStatus() error {
	return uiaeCachedItemStatus(v)
}

func (v *IUIAutomationElement) CachedBoundingRectangle() (*types.RECT, error) {
	return uiaeCachedBoundingRectangle(v)
}

func (v *IUIAutomationElement) CachedLabeledBy() error {
	return uiaeCachedLabeledBy(v)
}

func (v *IUIAutomationElement) CachedAriaRole() error {
	return uiaeCachedAriaRole(v)
}

func (v *IUIAutomationElement) CachedAriaProperties() error {
	return uiaeCachedAriaProperties(v)
}

func (v *IUIAutomationElement) CachedIsDataValidForForm() error {
	return uiaeCachedIsDataValidForForm(v)
}

func (v *IUIAutomationElement) CachedControllerFor() error {
	return uiaeCachedControllerFor(v)
}

func (v *IUIAutomationElement) CachedDescribedBy() error {
	return uiaeCachedDescribedBy(v)
}

func (v *IUIAutomationElement) CachedFlowsTo() error {
	return uiaeCachedFlowsTo(v)
}

func (v *IUIAutomationElement) CachedProviderDescription() error {
	return uiaeCachedProviderDescription(v)
}

func (v *IUIAutomationElement) GetClickablePoint() error {
	return uiaeGetClickablePoint(v)
}
