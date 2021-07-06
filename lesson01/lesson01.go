package main

import (
	"github.com/kyokomi/emoji"
)

func main() {
	message := emoji.Sprint("Hello, world :relaxed:")
	emoji.Println(message)
}
