package assambly

import (
	"context"
	"fmt"
	"os"

	"github.com/AssemblyAI/assemblyai-go-sdk"
)

type AssamblyS2T struct {
	apiKey string
}

func NewAssamblyS2T() (*AssamblyS2T, error) {
	k, exist := os.LookupEnv("ASSAMBLY_API_KEY")
	if !exist {
		return nil, fmt.Errorf("api key not found")
	}

	return &AssamblyS2T{apiKey: k}, nil
}

func (p *AssamblyS2T) Process(fileName string) (string, error) {
	ctx := context.Background()
	client := assemblyai.NewClient(p.apiKey)
	config := assemblyai.TranscriptOptionalParams{
		LanguageCode: "ru",
	}

	f, err := os.Open(fileName)
	if err != nil {
		f.Close()
		return "", err
	}

	transcript, err := client.Transcripts.TranscribeFromReader(ctx, f, &config)
	if err != nil {
		return "", err
	}

	return *transcript.Text, nil
}
