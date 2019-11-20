package Keyboard

import (
	"log"
	"syscall"
	"unsafe"
)

const VkReturn = 0x0D
const VkVolumeMute = 0xAD
const VkVolumeDown = 0xAE
const VkVolumeUp = 0xAF
const VkMediaNextTrack = 0xB0
const VkMediaPrevTrack = 0xB1
const VkMediaStop = 0xB2
const VkMediaPause = 0xB3
const VkPlay = 0xFA

type keyboardInput struct {
	wVk         uint16
	wScan       uint16
	dwFlags     uint32
	time        uint32
	dwExtraInfo uint64
}

type input struct {
	inputType uint32
	ki        keyboardInput
	padding   uint64
}

var (
	user32        = syscall.NewLazyDLL("user32.dll")
	sendInputProc = user32.NewProc("SendInput")
)

func Press(vkCode uint16) {
	var i input
	i.inputType = 1
	i.ki.wVk = vkCode
	ret, _, err := sendInputProc.Call(
		uintptr(1),
		uintptr(unsafe.Pointer(&i)),
		uintptr(unsafe.Sizeof(i)),
	)
	log.Printf("ret: %v error: %v", ret, err)
}
