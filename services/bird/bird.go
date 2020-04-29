package bird

import (
	"log"

	"github.com/YamiOdymel/microkernel/kernel"
)

// bird 定義了一個服務。
type bird struct {
}

// OnCall 會在被核心呼叫的時候執行。
func (b bird) OnCall(k kernel.Kernel) {
	log.Println("來自 Bird 服務的資訊：" + k.CPUInfo())
}

// New 會初始化一個服務。
func New() kernel.Service {
	return new(bird)
}
