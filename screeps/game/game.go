package game

import "syscall/js"

var Game = js.Global().Get("Game")

func GetTime() int {
	return Game.Get("time").Int()
}
