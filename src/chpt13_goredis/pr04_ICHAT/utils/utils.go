package utils

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

// 清屏
func ClearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// 打点进度条
func ProgressBar(n int) {
	for i := 0; i < n; i++ {
		fmt.Print(".")
		time.Sleep(time.Second)
	}
}
