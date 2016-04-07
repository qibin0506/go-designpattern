package player
import "fmt"

type MusicPlayer struct {
	Src string
}

func (p MusicPlayer) PlayMusic() {
	fmt.Println("play music: " + p.Src)
}