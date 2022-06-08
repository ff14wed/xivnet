//go:build darwin || dragonfly || freebsd || linux || netbsd || openbsd
// +build darwin dragonfly freebsd linux netbsd openbsd

package oodle

import (
	"errors"
)

type DecompressorState struct{}

func Init(libPath string) (*DecompressorState, error) {
	return nil, nil
}

func (d *DecompressorState) Decompress(input []byte, outputSize int64) ([]byte, error) {
	return nil, errors.New("not implemented")
}
