// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package item

import (
	"github.com/texterLine/exyuline/server/item/potion"
	"github.com/texterLine/exyuline/server/world"
	"time"
)

// Potion is an item that grants effects on consumption.
type Potion struct {
	// Type is the type of potion.
	Type potion.Potion
}

// MaxCount ...
func (p Potion) MaxCount() int {
	return 1
}

// AlwaysConsumable ...
func (p Potion) AlwaysConsumable() bool {
	return true
}

// ConsumeDuration ...
func (p Potion) ConsumeDuration() time.Duration {
	return DefaultConsumeDuration
}

// Consume ...
func (p Potion) Consume(_ *world.Tx, c Consumer) Stack {
	for _, effect := range p.Type.Effects() {
		c.AddEffect(effect)
	}
	return NewStack(GlassBottle{}, 1)
}

// EncodeItem ...
func (p Potion) EncodeItem() (name string, meta int16) {
	return "minecraft:potion", int16(p.Type.Uint8())
}
