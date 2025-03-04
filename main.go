package main

import (
	"fmt"
	"syscall"
)

func main() {
	h, e := syscall.LoadLibrary("bass.dll")
	if e != nil {
		fmt.Printf("Failed to Found BASS.dll")
	}
	defer syscall.FreeLibrary(h)
	proc, e := syscall.GetProcAddress(h, "BASS_Init") //One of the functions in the DLL
	r, _, _ := syscall.SyscallN(uintptr(proc), 1, 44100, 0, 0, 0)
	fmt.Printf("BASS_Init Function has been returned: %v\n", r)
	procm, e := syscall.GetProcAddress(h, "BASS_Start")
	var qw, _, _ = syscall.SyscallN(uintptr(procm))
	fmt.Printf("Returned Value: %d", qw)
	mn, _ := syscall.GetProcAddress(h, "BASS_StreamCreateFile")
	kl, _, _ := syscall.SyscallN(uintptr(mn), 0, uintptr(len("Example.mp3")), 0, 0, 0x4)
	qrt, _ := syscall.GetProcAddress(h, "BASS_ChannelPlay")
	syscall.SyscallN(uintptr(qrt), uintptr(kl), 0)
	for {
		fmt.Printf("HelloWorld!!!")
	}
}
