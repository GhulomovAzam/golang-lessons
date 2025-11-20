package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	//welcome
flag.Parse()
username := flag.Arg(0)
password := flag.Arg(1)

greeting := `Добро пожаловать, {username}! Ваш код : {password}!`
massage := strings.ReplaceAll(greeting, "{username}", username)
massage = strings.ReplaceAll(massage, "{password}", password)

fmt.Print(massage)

	
}