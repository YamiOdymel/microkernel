package snake

import (
	"log"

	"github.com/YamiOdymel/microkernel/kernel"
)

// snake 定義了一個服務。
type snake struct {
}

// OnCall 會在被核心呼叫的時候執行。
func (s snake) OnCall(k kernel.Kernel) {
	log.Println("來自 Snake 服務的資訊：" + k.DiskInfo())
}

// New 會初始化一個服務。
func New() kernel.Service {
	return new(snake)
}
