// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package inventory

import (
	"github.com/texterLine/exyuline/server/event"
	"github.com/texterLine/exyuline/server/item"
)

type Holder interface{}

type Context = event.Context[Holder]

// Handler is a type that may be used to handle actions performed on an inventory by a player.
type Handler interface {
	// HandleTake handles an item.Stack being taken from a slot in the inventory. This item might be the whole stack or
	// part of the stack currently present in that slot.
	HandleTake(ctx *Context, slot int, it item.Stack)
	// HandlePlace handles an item.Stack being placed in a slot of the inventory. It might either be added to an empty
	// slot or a slot that contains an item of the same type.
	HandlePlace(ctx *Context, slot int, it item.Stack)
	// HandleDrop handles the dropping of an item.Stack in a slot out of the inventory.
	HandleDrop(ctx *Context, slot int, it item.Stack)
}

// Check to make sure NopHandler implements Handler.
var _ Handler = NopHandler{}

// NopHandler is an implementation of Handler that does not run any code in any of its methods. It is the default
// Handler of an Inventory.
type NopHandler struct{}

func (NopHandler) HandleTake(*Context, int, item.Stack)  {}
func (NopHandler) HandlePlace(*Context, int, item.Stack) {}
func (NopHandler) HandleDrop(*Context, int, item.Stack)  {}
