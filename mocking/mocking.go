package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWorld = "Go!"
const countdownStart = 3

func Countdown(out io.Writer) {
	for i := countdownStart; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintln(out, i)
	}
	fmt.Fprint(out, finalWorld)
}

func main() {
	Countdown(os.Stdout)
}
