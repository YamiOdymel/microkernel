package kernel

// Service 定義了一個服務該有的功能。
type Service interface {
	// OnCall 表示服務至少也要能被核心呼叫。
	OnCall(k Kernel)
}
