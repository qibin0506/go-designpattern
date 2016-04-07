package player
import "fmt"

type GameSoundPlayer struct {
	Src string
}

func (p GameSoundPlayer) PlaySound() {
	fmt.Println("play sound: " + p.Src)
}