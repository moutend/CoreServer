#pragma once

#include <ICoreServer/ICoreServer.h>
#include <windows.h>

#include "types.h"

struct LogLoopContext {
  HANDLE QuitEvent = nullptr;
};

struct AutomationContext {
  bool IsActive = true;
  HANDLE QuitEvent = nullptr;
  HANDLE FocusEvent = nullptr;
  IUIAutomation5 *UIAutomation = nullptr;
  SAFEARRAY *FocusElementRuntimeId = nullptr;
  IUIEventHandler HandleFunc = nullptr;
  IUIAutomationElement *RootElement = nullptr;
  IUIAutomationTreeWalker *BaseTreeWalker = nullptr;
  IUIAutomationCacheRequest *BaseCacheRequest = nullptr;
};
