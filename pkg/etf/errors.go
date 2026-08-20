package etf

import "errors"

var (
	ErrInvalidMagic = errors.New("etf: invalid magic bytes")
	ErrCorrupted    = errors.New("etf: corrupted data")
)
