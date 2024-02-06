interface StorageData {
  value: unknown
  expire: number | null
}

export function setLocal(key: string, value: unknown, expire?: number) {
  if (value === undefined) {
    return
  }
  const storageData: StorageData = {
    value,
    expire: expire ? new Date().getTime() + expire * 1000 : null,
  }
  window.localStorage.setItem(key, JSON.stringify(storageData))
}

export function getLocal<T>(key: string) {
  const json = window.localStorage.getItem(key)
  if (json) {
    let storageData: StorageData | null = null
    try {
      storageData = JSON.parse(json)
    } catch {}
    if (storageData) {
      const { value, expire } = storageData
      // 在有效期内直接返回
      if (expire === null || expire >= Date.now()) {
        return value as T
      }
    }
    removeLocal(key)
    return null
  }
  return null
}

export function removeLocal(key: string) {
  window.localStorage.removeItem(key)
}

export function clearLocal() {
  window.localStorage.clear()
}
