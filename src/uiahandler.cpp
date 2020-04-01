#include <UIAutomationClient.h>
#include <UIAutomationCore.h>
#include <cpplogger/cpplogger.h>

#include <strsafe.h>

#include "types.h"
#include "uiahandler.h"
#include "util.h"

extern Logger::Logger *Log;

FocusChangeEventHandler::FocusChangeEventHandler(UIALoopContext *ctx)
    : mUIALoopContext(ctx) {}

ULONG FocusChangeEventHandler::AddRef() {
  ULONG ret = InterlockedIncrement(&mRefCount);
  return ret;
}

ULONG FocusChangeEventHandler::Release() {
  ULONG ret = InterlockedDecrement(&mRefCount);
  if (ret == 0) {
    delete this;
    return 0;
  }
  return ret;
}

HRESULT FocusChangeEventHandler::QueryInterface(REFIID riid,
                                                void **ppInterface) {
  if (riid == __uuidof(IUnknown))
    *ppInterface = static_cast<IUIAutomationFocusChangedEventHandler *>(this);
  else if (riid == __uuidof(IUIAutomationFocusChangedEventHandler))
    *ppInterface = static_cast<IUIAutomationFocusChangedEventHandler *>(this);
  else {
    *ppInterface = nullptr;
    return E_NOINTERFACE;
  }
  this->AddRef();
  return S_OK;
}

HRESULT
FocusChangeEventHandler::HandleFocusChangedEvent(
    IUIAutomationElement *pSender) {
  Log->Info(L"IUIAutomation Focus change event received", GetCurrentThreadId(),
            __LONGFILE__);
  {
    wchar_t *buffer = new wchar_t[256]{};
    StringCbPrintfW(buffer, 512, L"UIA focus event (0x%p)", pSender);
    Log->Info(buffer, GetCurrentThreadId(), __LONGFILE__);
    delete[] buffer;
    buffer = nullptr;
  }
  if (mUIALoopContext != nullptr && mUIALoopContext->HandleFunc != nullptr) {
    mUIALoopContext->HandleFunc(0, pSender);
  }

  return S_OK;
}

PropertyChangeEventHandler::PropertyChangeEventHandler(UIALoopContext *ctx)
    : mUIALoopContext(ctx) {}

ULONG PropertyChangeEventHandler::AddRef() {
  ULONG ret = InterlockedIncrement(&mRefCount);
  return ret;
}

ULONG PropertyChangeEventHandler::Release() {
  ULONG ret = InterlockedDecrement(&mRefCount);
  if (ret == 0) {
    delete this;
    return 0;
  }
  return ret;
}

HRESULT PropertyChangeEventHandler::QueryInterface(REFIID riid,
                                                   void **ppInterface) {
  if (riid == __uuidof(IUnknown))
    *ppInterface =
        static_cast<IUIAutomationPropertyChangedEventHandler *>(this);
  else if (riid == __uuidof(IUIAutomationPropertyChangedEventHandler))
    *ppInterface =
        static_cast<IUIAutomationPropertyChangedEventHandler *>(this);
  else {
    *ppInterface = nullptr;
    return E_NOINTERFACE;
  }
  this->AddRef();
  return S_OK;
}

HRESULT
PropertyChangeEventHandler::HandlePropertyChangedEvent(
    IUIAutomationElement *pSender, PROPERTYID propertyId, VARIANT newValue) {
  Log->Info(L"IUIAutomation Property change event received",
            GetCurrentThreadId(), __LONGFILE__);
  {
    wchar_t *buffer = new wchar_t[256]{};
    StringCbPrintfW(buffer, 512, L"UIA focus event (0x%p)", pSender);
    Log->Info(buffer, GetCurrentThreadId(), __LONGFILE__);
    delete[] buffer;
    buffer = nullptr;
  }
  if (mUIALoopContext != nullptr && mUIALoopContext->HandleFunc != nullptr) {
    mUIALoopContext->HandleFunc(0, pSender);
  }

  return S_OK;
}

AutomationEventHandler::AutomationEventHandler(UIALoopContext *ctx)
    : mUIALoopContext(ctx) {}

ULONG AutomationEventHandler::AddRef() {
  ULONG ret = InterlockedIncrement(&mRefCount);
  return ret;
}

ULONG AutomationEventHandler::Release() {
  ULONG ret = InterlockedDecrement(&mRefCount);
  if (ret == 0) {
    delete this;
    return 0;
  }
  return ret;
}

HRESULT AutomationEventHandler::QueryInterface(REFIID riid,
                                               void **ppInterface) {
  if (riid == __uuidof(IUnknown))
    *ppInterface = static_cast<IUIAutomationEventHandler *>(this);
  else if (riid == __uuidof(IUIAutomationEventHandler))
    *ppInterface = static_cast<IUIAutomationEventHandler *>(this);
  else {
    *ppInterface = nullptr;
    return E_NOINTERFACE;
  }
  this->AddRef();
  return S_OK;
}

HRESULT
AutomationEventHandler::HandleAutomationEvent(IUIAutomationElement *pSender,
                                              EVENTID eventId) {
  Log->Info(L"IUIAutomation Automation event received", GetCurrentThreadId(),
            __LONGFILE__);
  {
    wchar_t *buffer = new wchar_t[256]{};
    StringCbPrintfW(buffer, 512, L"UIA focus event (0x%p)", pSender);
    Log->Info(buffer, GetCurrentThreadId(), __LONGFILE__);
    delete[] buffer;
    buffer = nullptr;
  }
  if (mUIALoopContext != nullptr && mUIALoopContext->HandleFunc != nullptr) {
    mUIALoopContext->HandleFunc(0, pSender);
  }

  return S_OK;
}
