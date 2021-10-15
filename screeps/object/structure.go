package object

import (
	"github.com/whs/whscreeps/screeps/types"
	"github.com/whs/whscreeps/screeps/utils"
)

type Structure struct {
	RoomObject
}

// Destroy this structure immediately.
func (s Structure) Destroy() error {
	return utils.R(s.Raw.Call("destroy"))
}

// IsActive check whether this structure can be used.
// If room controller level is insufficient, then this method will return false
// and the structure will be highlighted with red in the game.
func (s Structure) IsActive() bool {
	return s.Raw.Call("isActive").Bool()
}

// NotifyWhenAttacked toggle auto notification when the structure is under attack.
// The notification will be sent to your account email. Turned on by default.
func (s Structure) NotifyWhenAttacked(enabled bool) error {
	return utils.R(s.Raw.Call("notifyWhenAttacked", enabled))
}

func (s Structure) Type() types.Structure {
	return s.Raw.Get("structureType").String()
}

func (s Structure) Store() Store {
	return Store{Raw: s.Raw.Get("store")}
}
