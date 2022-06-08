package oodle

import (
	"errors"
	"sync"
	"syscall"
	"unsafe"
)

var oodleDLLOnce struct {
	sync.Once
	dll *syscall.DLL

	procOodleNetwork1UDP_State_Size    *syscall.Proc
	procOodleNetwork1_Shared_Size      *syscall.Proc
	procOodleNetwork1_Shared_SetWindow *syscall.Proc
	procOodleNetwork1UDP_Train         *syscall.Proc
	procOodleNetwork1UDP_Decode        *syscall.Proc

	err error
}

type DecompressorState struct {
	state      []byte
	sharedDict []byte
}

func Init(libPath string) (*DecompressorState, error) {
	oodleDLLOnce.Do(func() {
		dll, err := syscall.LoadDLL(libPath)
		if err != nil {
			oodleDLLOnce.err = err
			return
		}
		oodleDLLOnce.dll = dll
		oodleDLLOnce.procOodleNetwork1UDP_State_Size, err = dll.FindProc("OodleNetwork1UDP_State_Size")
		if err != nil {
			oodleDLLOnce.err = err
			return
		}
		oodleDLLOnce.procOodleNetwork1_Shared_Size, err = dll.FindProc("OodleNetwork1_Shared_Size")
		if err != nil {
			oodleDLLOnce.err = err
			return
		}
		oodleDLLOnce.procOodleNetwork1_Shared_SetWindow, err = dll.FindProc("OodleNetwork1_Shared_SetWindow")
		if err != nil {
			oodleDLLOnce.err = err
			return
		}
		oodleDLLOnce.procOodleNetwork1UDP_Train, err = dll.FindProc("OodleNetwork1UDP_Train")
		if err != nil {
			oodleDLLOnce.err = err
			return
		}
		oodleDLLOnce.procOodleNetwork1UDP_Decode, err = dll.FindProc("OodleNetwork1UDP_Decode")
		if err != nil {
			oodleDLLOnce.err = err
			return
		}
	})
	if oodleDLLOnce.err != nil {
		return nil, oodleDLLOnce.err
	}
	stateSize, err := procCallWrapper(oodleDLLOnce.procOodleNetwork1UDP_State_Size)
	if err != nil {
		return nil, err
	}
	sharedDictSize, err := procCallWrapper(oodleDLLOnce.procOodleNetwork1_Shared_Size, 0x13)
	if err != nil {
		return nil, err
	}

	decompressorState := DecompressorState{
		state:      make([]byte, stateSize),
		sharedDict: make([]byte, sharedDictSize),
	}
	initDict := make([]byte, 0x8000)

	_, err = procCallWrapper(
		oodleDLLOnce.procOodleNetwork1_Shared_SetWindow,
		uintptr(unsafe.Pointer(&decompressorState.sharedDict[0])),
		0x13,
		uintptr(unsafe.Pointer(&initDict[0])),
		0x8000,
	)
	if err != nil {
		return nil, err
	}

	_, err = procCallWrapper(
		oodleDLLOnce.procOodleNetwork1UDP_Train,
		uintptr(unsafe.Pointer(&decompressorState.state[0])),
		uintptr(unsafe.Pointer(&decompressorState.sharedDict[0])),
		0,
		0,
		0,
	)
	if err != nil {
		return nil, err
	}

	return &decompressorState, nil
}

func (d *DecompressorState) Decompress(input []byte, outputSize int64) ([]byte, error) {
	output := make([]byte, outputSize)

	res, _, err := syscall.SyscallN(
		oodleDLLOnce.procOodleNetwork1UDP_Decode.Addr(),
		uintptr(unsafe.Pointer(&d.state[0])),
		uintptr(unsafe.Pointer(&d.sharedDict[0])),
		uintptr(unsafe.Pointer(&input[0])),
		uintptr(len(input)),
		uintptr(unsafe.Pointer(&output[0])),
		uintptr(outputSize),
	)
	if err != 0 {
		return nil, err
	}
	if res == 0 {
		return nil, errors.New("unable to decompress")
	}
	return output, nil
}

func procCallWrapper(proc *syscall.Proc, args ...uintptr) (uintptr, error) {
	res, _, err := syscall.SyscallN(proc.Addr(), args...)
	if err != 0 {
		return 0, err
	}
	return res, nil
}
