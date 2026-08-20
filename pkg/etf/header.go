package etf

import (
	"encoding/binary"
	"io"
)

func WriteHeader(w io.Writer, h *Header) error {
	buf := make([]byte, HeaderSize)

	copy(buf[0:4], h.Magic[:])
	buf[4] = h.Version
	buf[5] = h.ColorType
	buf[6] = h.Compression
	binary.BigEndian.PutUint32(buf[7:11], h.Width)
	binary.BigEndian.PutUint32(buf[11:15], h.Height)
	binary.BigEndian.PutUint32(buf[15:19], h.UncompressedSize)
	binary.BigEndian.PutUint32(buf[19:23], h.CompressedSize)
	binary.BigEndian.PutUint32(buf[23:27], h.Checksum)

	_, err := w.Write(buf)
	return err
}

func ReadHeader(r io.Reader) (*Header, error) {
	buf := make([]byte, HeaderSize)
	if _, err := io.ReadFull(r, buf); err != nil {
		return nil, err
	}

	h := &Header{}
	copy(h.Magic[:], buf[0:4])
	if h.Magic != Magic {
		return nil, ErrInvalidMagic
	}

	h.Version = buf[4]
	h.ColorType = buf[5]
	h.Compression = buf[6]
	h.Width = binary.BigEndian.Uint32(buf[7:11])
	h.Height = binary.BigEndian.Uint32(buf[11:15])
	h.UncompressedSize = binary.BigEndian.Uint32(buf[15:19])
	h.CompressedSize = binary.BigEndian.Uint32(buf[19:23])
	h.Checksum = binary.BigEndian.Uint32(buf[23:27])

	return h, nil
}
