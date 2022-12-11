package main

import (
	"fmt"

	"github.com/mattn/go-runewidth"
)

var emojis = []string{
	// 1 byte
	"⚗",

	"➕",
	"➖",
	"✨",
	"✅",
	"⏪",

	// 2 bytes
	"🎨",
	"🔥",
	"🐛",
	"🚑",
	"📝",
	"🚀",
	"💄",
	"🎉",
	"🔒",
	"🔖",
	"🚨",
	"🚧",
	"💚",
	"📌",
	"👷",
	"📈",
	"🔧",
	"🔨",
	"🌐",
	"💩",
	"🔀",
	"📦",
	"👽",
	"🚚",
	"📄",
	"💥",
	"🍱",
	"💡",
	"🍻",
	"💬",
	"🔊",
	"🔇",
	"👥",
	"🚸",
	"📱",
	"🤡",
	"🥚",
	"🙈",
	"📸",
	"🔍",
	"🌱",
	"🚩",
	"🥅",
	"💫",

	"🗃",
	"🏗",
	"🏷",
	"🗑",

	"⬇️",
	"⬆️",
	"♻️",
	"✏️",

	"♿️",

	// 3 bytes
	"⚡️ ",
}

const measurement = "                -123-"

const (
	runeWidth = iota
	stringWidth
)

func main() {
	fmt.Println("Rune Width")
	alignByWidth(runeWidth)
	fmt.Println("")

	fmt.Println("String Width")
	alignByWidth(stringWidth)
}

func alignByWidth(w int) {
	var width int

	fmt.Println(measurement)
	for _, emoji := range emojis {
		width = 0
		switch w {
		case runeWidth:
			for _, r := range emoji {
				width += runewidth.RuneWidth(r)
			}
		case stringWidth:
			width = runewidth.StringWidth(emoji)
		}

		padding := 3 - width

		fmt.Printf("Width: %v Emoji: |%v%*v|\n", width, emoji, padding, "")
	}
	fmt.Println(measurement)
}
