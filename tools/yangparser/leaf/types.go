package leaf

type LeafType uint

const (
	REFERENCE  LeafType = iota
	UINT32
	STRING
	BOOL
	IDENTIFY
	ENUMERATION
)
