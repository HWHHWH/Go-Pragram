package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"Go-Pragram/mplayer/library"
	"Go-Pragram/mplayer/mp"
)

var (
	lib          *library.MusicManager //音乐播放器
	id           int
	ctrl, signal chan int
)

//音乐库控制器，参数为控制命令：list，add，remove；使用格式为"lib command xxx"
func handleLibCommands(tokens []string) {
	switch tokens[1] {
	case "list": //lib list
		for i := 0; i < lib.Len(); i++ {
			e, _ := lib.Get(i)
			fmt.Println(i+1, ":", e.Name, e.Artist, e.Source, e.Type)
		}

	case "add": //lib add <name> <artist> <source> <type>
		if len(tokens) == 6 {
			id++
			lib.Add(&library.MusicEntry{strconv.Itoa(id), tokens[2], tokens[3], tokens[4], tokens[5]})
		} else {
			fmt.Println("USAGE:lib add <name><artist><source><type>")
		}

	case "remove": //ib remove <name>
		if len(tokens) == 3 {
			lib.RemoveByName(tokens[2])
		} else {
			fmt.Println("USAGE: lib remove <name>")
		}
	default:
		fmt.Println("Unrecognized lib command:", tokens[1])
	}
}

//播放操作
func handlePlayCommand(tokens []string) {
	if len(tokens) != 2 {
		fmt.Println("USADE: play <name>")
	}

	e := lib.Find(tokens[1])
	if e == nil {
		fmt.Println("The music", tokens[1], "does not exist")
		return
	}
	mp.Play(e.Source, e.Type)
}

//主函数
func main() {
	//打印说明
	fmt.Println(`    
                Enter following commands to control the player:
                lib list --view the existing music lib
                lib add <name><artist><source><type>  --add a music to the music lib
                lib remove <name> --remove the specified music from the lib
                play <name> --play the specified music
    `)
	//创建音乐播放器
	lib = library.NewMusicManager()
	//读取输入信息
	r := bufio.NewReader(os.Stdin)
	//解析输入参数
	for {
		fmt.Print("Enter command-> ")

		rawLine, _, _ := r.ReadLine()

		line := string(rawLine)

		if line == "q" || line == "e" {
			break
		}

		tokens := strings.Split(line, " ")

		if tokens[0] == "lib" {
			handleLibCommands(tokens)
		} else if tokens[0] == "play" {
			handlePlayCommand(tokens)
		} else {
			fmt.Println("Unrecongnized command:", tokens[0])
		}
	}
}
