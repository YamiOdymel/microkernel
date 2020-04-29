**概念**

* `mykernel` 實作了 `kernel.Kernel` 的定義。
* `kernel.Kernel` 並不知道 `services`，但他內部有專門控管服務的 `kernel.Registry`。
* 透過 `kernel.Registry` 將服務 `bird` 與 `snake` 註冊到 `mykernel`。

**核心定義**

```go
// Kernel 定義了一個最底層的核心應該有什麼樣的功能。
type Kernel interface {
	// CPUInfo 要能夠回傳處理器資訊供服務取得。
	CPUInfo() string
	// DiskInfo 要能夠回傳硬碟資訊供服務取得。
	DiskInfo() string
	// Registry 是這個核心至少要有一個服務中心。
	Registry
}
```

**中心定義**

```go
// Registry 定義了一個良好的服務中心應該有什麼樣的機能。
type Registry interface {
	// Register 會在服務中心註冊一個新的服務。
	Register(name string, service Service)
	// Destory 會摧毀一個在服務中心內的服務。
	Destory(name string)
	// Call 會呼叫所有的服務。
	Call()
}
```

**服務定義**

```go
// Service 定義了一個服務該有的功能。
type Service interface {
	// OnCall 表示服務至少也要能被核心呼叫。
	OnCall(k Kernel)
}
```

**執行效果**

```bash
yamiodymel@OMEN:/mnt/c/Users/YamiOdymel/go/src/github.com/YamiOdymel/microkernel$ ./microkernel
2020/04/29 16:09:03 成功註冊 bird 服務！
2020/04/29 16:09:04 成功註冊 snake 服務！
2020/04/29 16:09:04 來自 Bird 服務的資訊：CPU 使用率：10%
2020/04/29 16:09:04 來自 Snake 服務的資訊：硬碟使用率：32%
2020/04/29 16:09:04 已經註銷 bird 服務！
2020/04/29 16:09:04 來自 Snake 服務的資訊：硬碟使用率：32%
```