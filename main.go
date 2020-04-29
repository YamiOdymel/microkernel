package main

import (
	"github.com/YamiOdymel/microkernel/mykernel"
	"github.com/YamiOdymel/microkernel/services/bird"
	"github.com/YamiOdymel/microkernel/services/snake"
)

func main() {
	// k 會初始化一個自製的核心，而這個核心基礎來自 `kernel`。
	k := mykernel.New()

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
