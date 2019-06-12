package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls", "-la")
	stdout, err := cmd.StdoutPipe()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// コマンドの結果を1行ずつ実行
	cmd.Start()
	s := bufio.NewScanner(stdout)
	for s.Scan() {
		fmt.Println(s.Text())
	}

	cmd.Wait()
}
