package player

type GameSoundAdapter struct {
	SoundPlayer GameSoundPlayer
}

func (p GameSoundAdapter) PlayMusic() {
	p.SoundPlayer.PlaySound()
}