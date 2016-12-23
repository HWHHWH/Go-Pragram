package mp

import (
	"fmt"
	"time"
)

//WAVPlayer播放器
type WAVPlayer struct {
	stat     int
	progress int
}

//实现了Player接口
func (p *WAVPlayer) Play(source string) {
	fmt.Println("Playing WAV ", source) //source 音乐的位置

	p.progress = 0

	for p.progress < 100 {
		time.Sleep(100 * time.Millisecond) //假装正在播放，播放100秒
		fmt.Print(".")
		p.progress += 10
	}
	fmt.Println("\nFinished playing", source)
}
