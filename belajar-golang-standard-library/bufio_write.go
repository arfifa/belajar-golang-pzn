package main

import (
	"bufio"
	"os"
)

func main() {
	write := bufio.NewWriter(os.Stdout)
	_, _ = write.WriteString("Hello world\n")
	_, _ = write.WriteString("Selamat belajar\n")
	write.Flush()	
}