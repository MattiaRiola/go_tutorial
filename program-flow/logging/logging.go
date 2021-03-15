package logging

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

var debug bool

func init() {
	fmt.Println("enter your name")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	Name(input)
	Greetings()
}
func Debug(b bool) {
	debug = b
}

func Log(statement string) {
	if !debug {
		return
	}

	fmt.Printf("%s %s\n", time.Now().Format(time.RFC3339), statement)
}
