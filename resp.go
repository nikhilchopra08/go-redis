package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
)


func main(){
input := "$6\r\nNikhil\r\n"

reader := bufio.NewReader(strings.NewReader(input))

b, _ := reader.ReadByte()

if b != '$' {
	fmt.Println("Invalid type, expecting only string")
	os.Exit(1)
}

size, _ := reader.ReadByte()

strsize, _ :=  strconv.ParseInt(string(size), 10, 64)

// consure /r/n
reader.ReadByte()
reader.ReadByte()

name := make([]byte, strsize)
reader.Read(name)

fmt.Println(string(name))

}