package mp

import "fmt"

//播放器接口，包含一个播放方法
type Player interface{
    Play(source string)
}

//播放方法
func Play(source,mtype string) {

    var p Player

    switch mtype{   //通过接口类型来返回具体的对象
    case "MP3" :
        p = &MP3Player{}
    case "WAV" :
        p = &WAVPlayer{}
    default:
        fmt.Println("Unsurported music type",mtype)
    return
    }

    p.Play(source)
}