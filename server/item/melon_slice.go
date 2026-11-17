// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package item

import (
	"github.com/texterLine/exyuline/server/world"
	"time"
)

// MelonSlice is a food item dropped by melon blocks.
type MelonSlice struct{}

// AlwaysConsumable ...
func (m MelonSlice) AlwaysConsumable() bool {
	return false
}

// ConsumeDuration ...
func (m MelonSlice) ConsumeDuration() time.Duration {
	return DefaultConsumeDuration
}

// Consume ...
func (m MelonSlice) Consume(_ *world.Tx, c Consumer) Stack {
	c.Saturate(2, 1.2)
	return Stack{}
}

// CompostChance ...
func (MelonSlice) CompostChance() float64 {
	return 0.5
}

// EncodeItem ...
func (m MelonSlice) EncodeItem() (name string, meta int16) {
	return "minecraft:melon_slice", 0
}
