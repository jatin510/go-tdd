package mocks

import (
	"fmt"
	"io"
	"os"
)

func Countdown(out io.Writer) {
	fmt.Fprintf(out, "3")
}

func main() {
	Countdown(os.Stdout)
}
