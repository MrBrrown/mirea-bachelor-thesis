package detector

import (
	t "example.com/coomper/assistantCore/commands"
)

type Detector interface {
	Process(string) (t.Command, error)
}

var detector GigaDetector

func InitDetectot() error {
	detector = GigaDetector{Token{Token: "", ExpiresAt: 0}}
	return nil
}

func GetDetector() Detector {
	return &detector
}
