package main
import . "./player"

func main() {
	var player Player = MusicPlayer {Src:"music.mp3"}
	play(player)

	gameSound := GameSoundPlayer {Src:"game.mid"}
	gameAdapter := GameSoundAdapter {SoundPlayer:gameSound}
	play(gameAdapter)
}

func play(player Player) {
	player.PlayMusic()
}