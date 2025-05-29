package main

import (
	"app/internal/telegram"
	f "fmt"
)

const tg_token = "8109905612:AAHgYQW085wSPqJwhxSzWOHpoScdB8rEU4c"

func main() {
	if err := telegram.Run(tg_token); err != nil {
		f.Println(err)
		return
	}
}
