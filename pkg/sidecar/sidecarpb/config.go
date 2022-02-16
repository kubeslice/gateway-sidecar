package sidecar


// classType - Type of the Class
type classType string

// classType - Qdisc Class type  -  HTB
const (
	HTB classType = "HTB"
)

// tcInfo - the TC information
type TcInfo struct {
	// ClassType
	class classType
	// Bandwidth Ceiling in Kbps
	bwCeiling uint32
	// Bandwidth Guaranteed
	bwGuaranteed uint32
	// Priority
	priority uint32
}