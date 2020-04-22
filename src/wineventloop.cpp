#include <UIAutomationClient.h>
#include <UIAutomationCore.h>
#include <cpplogger/cpplogger.h>
#include <windows.h>

#include <Commctrl.h>
#include <oleacc.h>

#include "context.h"
#include "types.h"
#include "util.h"
#include "wineventloop.h"

extern Logger::Logger *Log;

static AutomationContext *ctx{};
static LONG prevLeft{};
static LONG prevTop{};
static LONG prevRight{};
static LONG prevBottom{};

void eventCallback(HWINEVENTHOOK hHook, DWORD eventId, HWND hWindow,
                   LONG objectId, LONG childId, DWORD threadId,
                   DWORD eventTime) {
  if (objectId <= OBJID_ALERT) {
    return;
  }
  if (!IsWindow(hWindow)) {
    return;
  }
  if (objectId == 0 && childId == 0) {
    objectId = OBJID_CLIENT;
  }

  Log->Info(L"IAccessible event received", GetCurrentThreadId(), __LONGFILE__);

  HRESULT hr{};
  RECT rectangle{};
  IUIAutomationElement *pSender{};

  if (ctx->UIAutomation == nullptr) {
    Log->Info(L"IUIAutomation is not available", GetCurrentThreadId(),
              __LONGFILE__);
    goto CLEANUP;
  }

  hr = ctx->UIAutomation->ElementFromHandle(hWindow, &pSender);

  if (FAILED(hr)) {
    Log->Info(L"Failed to get IUIAutomationElement from IAccessible",
              GetCurrentThreadId(), __LONGFILE__);
    goto CLEANUP;
  }

  hr = pSender->get_CurrentBoundingRectangle(&rectangle);

  if (FAILED(hr)) {
    goto CLEANUP;
  }

  LONG l = prevtLeft;
  LONG t = prevTop;
  LONG r = prevRight;
  LONG b = prevBottom;

  prevLeft = rectangle.left;
  prevTop = rectangle.top;
  prevRight = rectangle.right;
  prevBottom = rectangle.bottom;

  bool isSameBoundingRectangle = l == rectangle.left && t == rectangle.top &&
                                 r == rectangle.right && b == rectangle.bottom;

  if (isSameBoundingRectangle) {
    goto CLEANUP;
  }
  if (ctx->HandleFunc != nullptr) {
    ctx->HandleFunc(UIA_AutomationFocusChangedEventId, pSender);
  }

  SAFEARRAY *runtimeId{};

  hr = pSender->GetRuntimeId(&runtimeId);

  if (FAILED(hr)) {
    Log->Warn(L"Failed to call IUIAutomationElement::GetRuntimeId()",
              GetCurrentThreadId(), __LONGFILE__);
    goto CLEANUP;
  }
  if (runtimeId == nullptr) {
    goto CLEANUP;
  }

  hr = SafeArrayCopy(runtimeId, &(ctx->FocusElementRuntimeId));

  if (FAILED(hr)) {
    Log->Warn(L"Failed to call SafeArrayCopy()", GetCurrentThreadId(),
              __LONGFILE__);
    goto CLEANUP;
  }

CLEANUP:

  SafeRelease(&pSender);
}

DWORD WINAPI windowsEventThread(LPVOID context) {
  Log->Info(L"Start Windows event thread", GetCurrentThreadId(), __LONGFILE__);

  HRESULT hr = CoInitializeEx(nullptr, COINIT_MULTITHREADED);

  if (FAILED(hr)) {
    Log->Fail(L"Failed to call CoInitializeEx", GetCurrentThreadId(),
              __LONGFILE__);
    return hr;
  }

  ctx = static_cast<AutomationContext *>(context);

  if (ctx == nullptr) {
    Log->Fail(L"Failed to obtain context", GetCurrentThreadId(), __LONGFILE__);
    return E_FAIL;
  }

  HWINEVENTHOOK hookIds[3]{};
  const DWORD events[3] = {EVENT_SYSTEM_DESKTOPSWITCH, EVENT_OBJECT_FOCUS,
                           EVENT_SYSTEM_FOREGROUND};

  for (int i = 0; i < 3; i++) {
    hookIds[i] = SetWinEventHook(events[i], events[i], nullptr, eventCallback,
                                 0, 0, WINEVENT_OUTOFCONTEXT);

    if (hookIds[i] == 0) {
      Log->Warn(L"Failed to call SetWinEventHook()", GetCurrentThreadId(),
                __LONGFILE__);
    }
  }

  Log->Info(L"Register callbacks", GetCurrentThreadId(), __LONGFILE__);

  UINT_PTR timerId = SetTimer(nullptr, 0, 3000, nullptr);
  MSG msg;

  while (GetMessage(&msg, nullptr, 0, 0)) {
    TranslateMessage(&msg);
    DispatchMessage(&msg);

    if (!ctx->IsActive) {
      Log->Info(L"Win events no longer used", GetCurrentThreadId(),
                __LONGFILE__);
      break;
    }
  }

  KillTimer(nullptr, timerId);

  for (int i = 0; i < 3; i++) {
    if (hookIds[i] == 0) {
      continue;
    }

    UnhookWinEvent(hookIds[i]);
  }

  CoUninitialize();

  Log->Info(L"End Windows event thread", GetCurrentThreadId(), __LONGFILE__);

  return S_OK;
}
