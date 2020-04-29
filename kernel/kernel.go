package kernel

// Kernel 定義了一個最底層的核心應該有什麼樣的功能。
type Kernel interface {
	// CPUInfo 要能夠回傳處理器資訊供服務取得。
	CPUInfo() string
	// DiskInfo 要能夠回傳硬碟資訊供服務取得。
	DiskInfo() string
	// Registry 是這個核心至少要有一個服務中心。
	Registry
}
