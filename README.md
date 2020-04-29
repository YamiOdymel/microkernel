* `mykernel` 實作了 `kernel.Kernel` 的定義。
* `kernel.Kernel` 並不知道 `services`，但他內部有專門控管服務的 `kernel.Registry`。
* 透過 `kernel.Registry` 將服務 `bird` 與 `snake` 註冊到 `mykernel`。

```bash
yamiodymel@OMEN:/mnt/c/Users/YamiOdymel/go/src/github.com/YamiOdymel/microkernel$ ./microkernel
2020/04/29 16:09:03 成功註冊 bird 服務！
2020/04/29 16:09:04 成功註冊 snake 服務！
2020/04/29 16:09:04 來自 Bird 服務的資訊：CPU 使用率：10%
2020/04/29 16:09:04 來自 Snake 服務的資訊：硬碟使用率：32%
2020/04/29 16:09:04 已經註銷 bird 服務！
2020/04/29 16:09:04 來自 Snake 服務的資訊：硬碟使用率：32%
```