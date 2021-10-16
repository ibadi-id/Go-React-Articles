package helper

// Package ini hanya opsional untuk bantuan menampilkan pesan server

import (
	"fmt"

	"github.com/fatih/color"
)

func LineSeparator() {
	fmt.Println("=========================================")
}

func StartMessage(msg1 string) {
	LineSeparator()
	color.Yellow(msg1)
	LineSeparator()
}