package mp

import (
    "fmt"
    "time"
)
//MP3播放器
type MP3Player struct {
    stat     int
    progress int
}

//实现了Player接口
func (p *MP3Player) Play(source string) {
    fmt.Println("Playing MP3 music", source) //source 音乐的位置

    p.progress = 0

    for p.progress < 100 {
        time.Sleep(100 * time.Millisecond)    //假装正在播放，播放100秒
        fmt.Print(".")
        p.progress +=10
    }
    fmt.Println("\nFinished playing",source)
}
