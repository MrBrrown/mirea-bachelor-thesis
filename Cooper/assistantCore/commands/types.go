package commands

type Command int

const (
	CNCInfo Command = iota
	CNCProcess
	CNCProcessInfo
	CNCProcessParams
	CNCParamInfo
	CNCParamValue

	Unknown
)
