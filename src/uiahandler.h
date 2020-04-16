#pragma once

#include <UIAutomationClient.h>
#include <oleauto.h>
#include <windows.h>

#include "context.h"

class FocusChangeEventHandler : public IUIAutomationFocusChangedEventHandler {
public:
  // IUnknown methods
  STDMETHODIMP QueryInterface(REFIID riid, void **ppvObject);
  STDMETHODIMP_(ULONG) AddRef();
  STDMETHODIMP_(ULONG) Release();

  // IUIAutomationFocusChangedEventHandler methods
  STDMETHODIMP HandleFocusChangedEvent(IUIAutomationElement *pSender);

  FocusChangeEventHandler(UIALoopContext *ctx);

private:
  LONG mRefCount;
  UIALoopContext *mUIALoopContext;
};

class PropertyChangeEventHandler
    : public IUIAutomationPropertyChangedEventHandler {
public:
  // IUnknown methods
  STDMETHODIMP QueryInterface(REFIID riid, void **ppvObject);
  STDMETHODIMP_(ULONG) AddRef();
  STDMETHODIMP_(ULONG) Release();

  // IUIAutomationPropertyChangedEventHandler methods
  STDMETHODIMP HandlePropertyChangedEvent(IUIAutomationElement *sender,
                                          PROPERTYID propertyId,
                                          VARIANT newValue);

  PropertyChangeEventHandler(UIALoopContext *ctx);

private:
  LONG mRefCount;
  UIALoopContext *mUIALoopContext;
};

class AutomationEventHandler : public IUIAutomationEventHandler {
public:
  // IUnknown methods
  STDMETHODIMP QueryInterface(REFIID riid, void **ppvObject);
  STDMETHODIMP_(ULONG) AddRef();
  STDMETHODIMP_(ULONG) Release();

  // IUIAutomationEventHandler methods
  STDMETHODIMP HandleAutomationEvent(IUIAutomationElement *sender,
                                     EVENTID eventId);

  AutomationEventHandler(UIALoopContext *ctx);

private:
  LONG mRefCount;
  UIALoopContext *mUIALoopContext;
};

class StructureChangeEventHandler
    : public IUIAutomationStructureChangedEventHandler {
public:
  // IUnknown methods
  STDMETHODIMP QueryInterface(REFIID riid, void **ppvObject);
  STDMETHODIMP_(ULONG) AddRef();
  STDMETHODIMP_(ULONG) Release();

  // IUIAutomationStructureChangedEventHandler methods
  STDMETHODIMP HandleStructureChangedEvent(IUIAutomationElement *pSender,
                                           StructureChangeType changeType,
                                           SAFEARRAY(int) runtimeId);

  StructureChangeEventHandler(UIALoopContext *ctx);

private:
  LONG mRefCount;
  UIALoopContext *mUIALoopContext;
};
