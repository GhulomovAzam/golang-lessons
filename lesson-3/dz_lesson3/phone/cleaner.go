package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	flag.Parse()
phoneNumber := flag.Arg(0) 
right := strings.ReplaceAll(phoneNumber, "(", "")
right = strings.ReplaceAll(right, ")", "")
right = strings.ReplaceAll(right, "-", "")
right = strings.ReplaceAll(right, " ", "")
fmt.Print(right)

}