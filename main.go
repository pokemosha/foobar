package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func recursion(output io.Writer, i uint, num uint) {
	fmt.Fprintf(output, "%d Foo", i)
	if i%3 == 0 {
		fmt.Fprint(output, " Bar")
	}
	fmt.Fprintln(output)
	if i < num {
		recursion(output, i+1, num)
	}
}

func foobar(writer io.Writer, reader io.Reader) error {
	var number uint
	_, err := fmt.Fscanln(reader, &number)
	if err != nil {
		return err
	}
	if number == 0 {
		return fmt.Errorf("expected integer more than 0")
	}
	recursion(writer, 1, number)
	return nil
}

func main() {
	err := foobar(os.Stdout, os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
}
