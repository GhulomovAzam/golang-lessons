package main

import (
	"fmt"
	"strings"
)

func main() {
	//ДЗ-1
	template := `Добро пожаловать, {username}!`
	massage := strings.ReplaceAll(template, "{username}","Алишер")//Пакет strings у которого есть функция ReplaceAll то есть заменить
	fmt.Println(massage)
}