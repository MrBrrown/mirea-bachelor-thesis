package assambly

import (
	"os/exec"

	htgotts "github.com/hegedustibor/htgo-tts"
	"github.com/hegedustibor/htgo-tts/voices"
)

type Text2Speech struct {
	speech *htgotts.Speech
}

func NewText2Speech() *Text2Speech {
	return &Text2Speech{
		speech: &htgotts.Speech{Folder: "audio", Language: voices.Russian},
	}
}

func (p *Text2Speech) Configure() error {
	return nil
}

func (p *Text2Speech) Process(sentance string) error {
	player := exec.Command("say", sentance)
	return player.Run()
}
