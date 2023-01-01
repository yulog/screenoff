package main

import (
	"os"

	"github.com/lxn/win"
)

func main() {
	win.SendMessage(win.HWND_TOPMOST, 0x0112, 0xF170, 2)
	os.Exit(0)
}
