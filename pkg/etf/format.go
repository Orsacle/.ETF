package etf

const MagicSize = 4

var Magic = [MagicSize]byte{'E', 'T', 'F', '1'}

const FormatVersion uint8 = 1

const (
	ColorTypeGray uint8 = iota
	ColorTypeRGB
	ColorTypeRGBA
)

const (
	CompressionNone uint8 = iota
	CompressionLZ4
)

//--> Magic(4) + Version(1) + ColorType(1) + Compression(1) + Width(4) + Height(4) + UncompressedSize(4) + CompressedSize(4) + Checksum(4)
const HeaderSize = MagicSize + 1 + 1 + 1 + 4 + 4 + 4 + 4 + 4

type Header struct {
	Magic [MagicSize]byte
	Version uint8
	ColorType uint8
	Compression uint8
	Width uint32
	Height uint32
	UncompressedSize uint32
	CompressedSize uint32
	Checksum uint32
}
