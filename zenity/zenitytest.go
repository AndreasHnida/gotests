package main

import (
	"fmt"

	"github.com/ncruces/zenity"
)

func main() {
	ask()
}

func ask() {
	response, _ := zenity.Entry("Enter new text:",
		zenity.Title("Add a new entry"))
	fmt.Println(response)
}
