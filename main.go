package main

import (
	"github.com/YamiOdymel/microkernel/kernel"
	"github.com/YamiOdymel/microkernel/services/bird"
	"github.com/YamiOdymel/microkernel/services/snake"
)

// mykernel 實作了一個底層 kernel 核心。
type mykernel struct {
	kernel.Registry
}

// CPUInfo 會實作取得處理器資訊的功能。
func (k *mykernel) CPUInfo() string {
	return "CPU 使用率：10%"
}

// DiskInfo 會實作取得硬碟資訊的功能。
func (k *mykernel) DiskInfo() string {
	return "硬碟使用率：32%"
}

// New 會初始化一個新的自製核心，並以 `kernel` 的定義為基礎。
func New() kernel.Kernel {
	// 先初始化這個核心。
	k := new(mykernel)
	// 替這個核心初始化一個服務中心，並將自己傳入服務中心。
	r := kernel.NewRegistry(k)
	// 然後再把服務中心放回核心。
	k.Registry = r
	return k
}

func main() {
	// k 會初始化一個自製的核心，而這個核心基礎來自 `kernel`。
	k := New()

	// 在自己核心的服務中心裡註冊兩個服務。
	k.Register("bird", bird.New())
	k.Register("snake", snake.New())
	// 呼叫所有的服務。
	k.Call()

	// 註銷其中一個服務。
	k.Destory("bird")
	// 再次呼叫所有的服務，這個時候應該只剩一個回應。
	k.Call()
}
