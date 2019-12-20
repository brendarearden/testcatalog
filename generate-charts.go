package main

import (
	"fmt"
	"os/exec"
)

func main() {
	for i := 1; i < 4001; i++ {
		chartName := fmt.Sprintf("foo%v", i)
		cmd := exec.Command("helm", "create", chartName)
		cmd.Run()
	}
}
