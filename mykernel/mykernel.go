package mykernel

import (
	"github.com/YamiOdymel/microkernel/kernel"
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
