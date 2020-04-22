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
  IAccessible *pAcc{};
  VARIANT vChild{};

  hr = AccessibleObjectFromEvent(hWindow, objectId, childId, &pAcc, &vChild);

  if (FAILED(hr)) {
    Log->Info(L"Failed to get IAccessible", GetCurrentThreadId(), __LONGFILE__);
    goto CLEANUP;
  }
  if (ctx->IAEventHandleFunc != nullptr) {
    ctx->IAEventHandleFunc(eventId, pAcc);
  }

CLEANUP:

  SafeRelease(&pAcc);
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
