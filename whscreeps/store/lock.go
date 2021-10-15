package store

import (
	"fmt"
	"github.com/whs/whscreeps/screeps/game"
	"sync"
)

type LockToken string

type Mutex struct {
	_     [0]sync.Mutex // no copy
	Token *LockToken
}

func (l *Mutex) Lock() LockToken {
	if l.Token != nil {
		panic("Lock() called on locked object")
	}

	token := lockToken()
	l.Token = &token
	return token
}

func (l *Mutex) IsMyLock(token LockToken) bool {
	if l.Token == nil {
		return false
	}
	return *l.Token == token
}

func (l *Mutex) Unlock(token LockToken) {
	if !l.IsMyLock(token) {
		panic("Unlock() called with wrong owner")
	}

	l.Token = nil
}

var counter = 0

func lockToken() LockToken {
	counter += 1
	return LockToken(fmt.Sprintf("%d:%d", game.GetTime(), counter))
}
