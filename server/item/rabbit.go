// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package item

import "github.com/texterLine/exyuline/server/world"

// Rabbit is a food item obtained from rabbits. It can be cooked in a furnace, smoker, or campfire.
type Rabbit struct {
	defaultFood

	// Cooked is whether the rabbit is cooked.
	Cooked bool
}

// Consume ...
func (r Rabbit) Consume(_ *world.Tx, c Consumer) Stack {
	if r.Cooked {
		c.Saturate(5, 6)
	} else {
		c.Saturate(3, 1.8)
	}
	return Stack{}
}

// SmeltInfo ...
func (r Rabbit) SmeltInfo() SmeltInfo {
	if r.Cooked {
		return SmeltInfo{}
	}
	return newFoodSmeltInfo(NewStack(Rabbit{Cooked: true}, 1), 0.35)
}

// EncodeItem ...
func (r Rabbit) EncodeItem() (name string, meta int16) {
	if r.Cooked {
		return "minecraft:cooked_rabbit", 0
	}
	return "minecraft:rabbit", 0
}
