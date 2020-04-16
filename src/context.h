#pragma once

#include <ICoreServer/ICoreServer.h>
#include <windows.h>

#include "types.h"

struct LogLoopContext {
  HANDLE QuitEvent = nullptr;
};

struct UIALoopContext {
  HANDLE QuitEvent = nullptr;
  IUIEventHandler HandleFunc = nullptr;
  IUIAutomationElement *RootElement = nullptr;
  IUIAutomation5 *UIAutomation = nullptr;
  IUIAutomationTreeWalker *BaseTreeWalker = nullptr;
  IUIAutomationCacheRequest *BaseCacheRequest = nullptr;
};

struct WinEventLoopContext {
  HANDLE QuitEvent = nullptr;
  IAEventHandler HandleFunc = nullptr;
  bool IsActive = true;
};
