package opc

type OpcServer struct {
	connectionString string
}

func (s *OpcServer) Init(cs string, node string) error {
	s.connectionString = cs
	return nil
}

func (s *OpcServer) GetCNCInfo(infoType string) (string, error) {
	return "", nil
}

func (s *OpcServer) GetCurProcess() (string, error) {
	return "", nil
}

func (s *OpcServer) GetProcessInfo(processId string) (string, error) {
	return "", nil
}

func (s *OpcServer) GetProcessParams(processId string) ([]string, error) {
	return nil, nil
}

func (s *OpcServer) GetParamInfo(paramId string) (string, error) {
	return "", nil
}

func (s *OpcServer) GetCurValue(paramId string) (float64, error) {
	return 0, nil
}
