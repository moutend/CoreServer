#pragma once

#include <ICoreServer/ICoreServer.h>
#include <mutex>
#include <windows.h>

class CCoreServer : public ICoreServer {
public:
  // IUnknown methods
  STDMETHODIMP QueryInterface(REFIID riid, void **ppvObject);
  STDMETHODIMP_(ULONG) AddRef();
  STDMETHODIMP_(ULONG) Release();

  // IDispatch methods
  STDMETHODIMP GetTypeInfoCount(UINT *pctinfo);
  STDMETHODIMP GetTypeInfo(UINT iTInfo, LCID lcid, ITypeInfo **ppTInfo);
  STDMETHODIMP GetIDsOfNames(REFIID riid, LPOLESTR *rgszNames, UINT cNames,
                             LCID lcid, DISPID *rgDispId);
  STDMETHODIMP Invoke(DISPID dispIdMember, REFIID riid, LCID lcid, WORD wFlags,
                      DISPPARAMS *pDispParams, VARIANT *pVarResult,
                      EXCEPINFO *pExcepInfo, UINT *puArgErr);

  // ICoreServer methods
  STDMETHODIMP Start();
  STDMETHODIMP Stop();
  STDMETHODIMP SetUIEventHandler(UIEventHandler handler);

  CCoreServer();
  ~CCoreServer();

private:
  LONG mReferenceCount;
  ITypeInfo *mTypeInfo;

  std::mutex mCoreServerMutex;
};

class CCoreServerFactory : public IClassFactory {
public:
  // IUnknown methods
  STDMETHODIMP QueryInterface(REFIID riid, void **ppvObject);
  STDMETHODIMP_(ULONG) AddRef();
  STDMETHODIMP_(ULONG) Release();

  // IClassFactory methods
  STDMETHODIMP CreateInstance(IUnknown *pUnkOuter, REFIID riid,
                              void **ppvObject);
  STDMETHODIMP LockServer(BOOL fLock);
};

extern void LockModule(BOOL bLock);
extern BOOL CreateRegistryKey(HKEY hKeyRoot, LPTSTR lpszKey, LPTSTR lpszValue,
                              LPTSTR lpszData);
extern void GetGuidString(REFGUID rguid, LPTSTR lpszGuid);
