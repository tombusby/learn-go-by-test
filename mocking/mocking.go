package main

import (
	"fmt"
	"io"
	"os"
)

const finalWorld = "Go!"
const countdownStart = 3

func Countdown(out io.Writer) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
	}
	fmt.Fprint(out, finalWorld)
}

func main() {
	Countdown(os.Stdout)
}
