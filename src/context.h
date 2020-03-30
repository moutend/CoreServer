#pragma once

#include <ICoreServer/ICoreServer.h>
#include <windows.h>

#include "types.h"

struct LogLoopContext {
  HANDLE QuitEvent = nullptr;
};

struct UIALoopContext {
  HANDLE QuitEvent = nullptr;
  UIEventHandler HandleFunc = nullptr;
};

struct WinEventLoopContext {
  HANDLE QuitEvent = nullptr;
  UIEventHandler HandleFunc = nullptr;
  bool IsActive = true;
};