package main

import (
	"fmt"
	"os/exec"
)

func main() {
	for i := 1; i < 2000; i++ {
		chartName := fmt.Sprintf("bar%v", i)
		cmd := exec.Command("helm", "create", chartName)
		cmd.Run()
	}
}
