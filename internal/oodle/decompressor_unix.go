//go:build darwin || dragonfly || freebsd || linux || netbsd || openbsd || wasm
// +build darwin dragonfly freebsd linux netbsd openbsd wasm

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
