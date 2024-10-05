package opc

import crypto "example.com/coomper/crypto"

type Provider interface {
	Init(connectionString string, node string) error

	GetCNCInfo(infoType string) (string, error)

	GetCurProcess() (string, error)
	GetProcessInfo(processId string) (string, error)
	GetProcessParams(processId string) ([]string, error)

	GetParamInfo(paramId string) (string, error)
	GetCurValue(paramId string) (float64, error)
}

func providerFactory() Provider {
	return &MocServer{}
}

func InitServer(data []byte) (Provider, error) {
	cs, node, err := crypto.GetConnectionAtrribs(data)
	if err != nil {
		return nil, err
	}

	p := providerFactory()

	return p, p.Init(cs, node)
}
