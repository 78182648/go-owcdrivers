package hypercashTransaction

const (
	OP_RETURN     byte = 0x6a
	TxTreeRegular byte = 0
	TxTreeStake   byte = 1
	SigHashAll    byte = 1

	DefaultTxVersion uint16 = 1
	DefaultPkScriptVersion uint16 = 0

	NoExpiryValue uint32 = 0

)


const (
	TxSerializeFull uint16 = iota
	TxSerializeNoWitness
	TxSerializeOnlyWitness
	TxSerializeWitnessSigning
	TxSerializeWitnessValueSigning
)


const (
	PropertyID_Omni    uint32 = 1
	PropertyID_LEEK    uint32 = 11
	PropertyID_YAX     uint32 = 18
	PropertyID_ENTCASH uint32 = 19
)

var (
	PKHAddressPrefix = []byte{0x09, 0x7f}
)


var (
	CurveOrder     = []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFE, 0xBA, 0xAE, 0xDC, 0xE6, 0xAF, 0x48, 0xA0, 0x3B, 0xBF, 0xD2, 0x5E, 0x8C, 0xD0, 0x36, 0x41, 0x41}
	HalfCurveOrder = []byte{0x7F, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x5D, 0x57, 0x6E, 0x73, 0x57, 0xA4, 0x50, 0x1D, 0xDF, 0xE9, 0x2F, 0x46, 0x68, 0x1B, 0x20, 0xA0}
)