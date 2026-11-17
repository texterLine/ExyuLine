// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package player

import (
	"github.com/texterLine/exyuline/server/world"
)

// Context is the context passed to player event callbacks. It embeds the
// world Context, so world operations and Cancel are available directly, and
// adds the Player the event concerns. It is valid only during the callback.
type Context struct {
	*world.Context
	p *Player
}

// NewEventContext returns a fresh event context for p.
// tx and p must come from the same active owner callback.
func NewEventContext(tx *world.Tx, p *Player) *Context {
	if tx == nil || p == nil || p.tx != tx {
		panic("player: transaction and player do not belong to the same callback")
	}
	_ = tx.World() // Fail immediately if tx has already finished.
	return &Context{Context: tx.Event(), p: p}
}

// Player returns the player the event concerns, valid only during the
// callback.
func (ctx *Context) Player() *Player { return ctx.p }

// Defer schedules f to run on the owner after the current callback completes,
// with the player re-resolved for that moment. The task fails with
// world.ErrEntityClosed if the player's handle closed, or with
// world.ErrEntityNotInWorld if the player left this transaction's world.
func (ctx *Context) Defer(f func(ctx *Context)) *world.Task {
	return ctx.DeferErr(func(ctx *Context) error {
		f(ctx)
		return nil
	})
}

// DeferErr schedules f like Defer and records its returned error on the Task.
func (ctx *Context) DeferErr(f func(ctx *Context) error) *world.Task {
	h := ctx.p.H()
	return ctx.Context.DeferErr(func(tx *world.Tx) error {
		if e, ok := h.Entity(tx); ok {
			return f(NewEventContext(tx, e.(*Player)))
		}
		if h.Closed() {
			return world.ErrEntityClosed
		}
		return world.ErrEntityNotInWorld
	})
}
