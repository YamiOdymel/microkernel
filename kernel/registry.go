package kernel

import "log"

// Registry 定義了一個良好的服務中心應該有什麼樣的機能。
type Registry interface {
	// Register 會在服務中心註冊一個新的服務。
	Register(name string, service Service)
	// Destory 會摧毀一個在服務中心內的服務。
	Destory(name string)
	// Call 會呼叫所有的服務。
	Call()
}

// Service 定義了一個服務該有的功能。
type Service interface {
	// OnCall 表示服務至少也要能被核心呼叫。
	OnCall(k Kernel)
}

// registry 會實作一個服務中心。
type registry struct {
	// kernel 是這個服務中心所屬的核心資訊。
	kernel Kernel
	// services 是所有已經註冊的服務。
	services map[string]Service
}

// NewRegistry 會建立並實做一個服務中心。
func NewRegistry(kernel Kernel) Registry {
	return &registry{
		kernel:   kernel,
		services: make(map[string]Service),
	}
}

// Register 會在服務中心註冊一個新的服務。
func (r *registry) Register(name string, svc Service) {
	log.Println("成功註冊 " + name + " 服務！")
	r.services[name] = svc
}

// Destory 會摧毀一個在服務中心內的服務。
func (r *registry) Destory(name string) {
	log.Println("已經註銷 " + name + " 服務！")
	delete(r.services, name)
}

// Call 會呼叫所有註冊的服務。
func (r *registry) Call() {
	for _, v := range r.services {
		v.OnCall(r.kernel)
	}
}
