package object

import (
	"github.com/whs/whscreeps/screeps/game"
	"syscall/js"
)

type RoomObject struct {
	Raw js.Value
}

func (r RoomObject) JSValue() js.Value {
	return r.Raw
}

func (r RoomObject) GetPosition() Position {
	return Position{Value: r.Raw.Get("position")}
}

func (r RoomObject) ID() string {
	return r.Raw.Get("id").String()
}

func RoomObjectById(id string) *RoomObject {
	val := game.Game.Call("getObjectById", id)
	if val.IsNull() {
		return nil
	}
	return &RoomObject{Raw: val}
}

type Position struct {
	js.Value
}

func (p Position) JSValue() js.Value {
	return p.Value
}

func (p Position) GetX() int {
	return p.Get("x").Int()
}

func (p Position) GetY() int {
	return p.Get("y").Int()
}

func (p Position) GetRoomName() string {
	return p.Get("roomName").String()
}
