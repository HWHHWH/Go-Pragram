package main

import(
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"

    "Go-Pragram/mplayer/mp"
    "Go-Pragram/mplayer/library"
)

var (
    lib *library.MusicManager
    id int = 1
    ctrl，signal chan int
)

func handleLibCommands (tokens []string){
    switch tokens[1] {
        case "list":
            for i:= 0; i <lib.Len();i++ {
                e,_ := lib.Get(i)
                fmt.Println(i+1,":",e.Name,e.Artist,e.Source,e.Type)
            }
        }
        case "add":{
            if len(tokens) == 6 {
                id++
                lib.Add(&library.MusicEntry{
                    strconv.Itoa(id),tokens[2],tokens[3],tokens[4],tokens][5]
                })
            } else{
                fmt.Println("USAGE:lib add <name><artist><source><type>")
            }
        }
        case "remove":
            if len(tokens) == 3 {
                lib.RemoveByName(tokens[2])
            } else {
                fmt.Println("USAGE: lib remove <name>")
            }
        default:
            fmt.Println("Unrecognized lib command:",token[1])
    }
}

func handlePlayCommand(tokens []string)  {
    if len(tokens) != 2 {
        fmt.Println("USADE: play <name>")
    }

    e := lib.Find(tokens[1])
    if e == nil {
        fmt.Println("The music",tockens[1],"does not exist")
        return
    }
    m.play(e.Source,e.Type,ctrl，signal)
}

func main() {
    fmt.Println(`
                Enter following commands to control the player:
                lib list --view the existing music lib
                lib add <name><artist><source><type>  --add a music to the music lib
                lib remove <name> --remove the specified music from the lib
                play <name> --play the specified music


    `)
    lib=library.NewMusicManager()
    r:=bufio.NewReader(os.Stdin)

    for{
        fmt.Print("Enter command-> ")

        rawLine,_,_ := r.ReadLine()

        line := string(rawLine)

        if line== "q" || line =="e"{
            break
        }

        tokens := strings.Split(line," ")

        if tokens[0] == "lib"{
            handleLibCommands(tokens)
        } else if tokens[0] == "play"{
            handlePlayCommand(tokens)
        } else{
            fmt.Println("Unrecongnized command:",tokens[0])
        }
    }
}