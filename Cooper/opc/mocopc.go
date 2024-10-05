package opc

type MocServer struct {
}

func (s *MocServer) Init(cs string, node string) error {
	return nil
}

func (s *MocServer) GetCNCInfo(infoType string) (string, error) {
	return "", nil
}

func (s *MocServer) GetCurProcess() (string, error) {
	return "", nil
}

func (s *MocServer) GetProcessInfo(processId string) (string, error) {
	return "", nil
}

func (s *MocServer) GetProcessParams(processId string) ([]string, error) {
	return nil, nil
}

func (s *MocServer) GetParamInfo(paramId string) (string, error) {
	return "", nil
}

func (s *MocServer) GetCurValue(paramId string) (float64, error) {
	return 0, nil
}
