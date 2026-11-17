// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package item

import "github.com/texterLine/exyuline/server/world"

// Mutton is a food item obtained from sheep. It can be cooked in a furnace, smoker, or campfire.
type Mutton struct {
	defaultFood

	// Cooked is whether the mutton is cooked.
	Cooked bool
}

// Consume ...
func (m Mutton) Consume(_ *world.Tx, c Consumer) Stack {
	if m.Cooked {
		c.Saturate(6, 9.6)
	} else {
		c.Saturate(2, 1.2)
	}
	return Stack{}
}

// SmeltInfo ...
func (m Mutton) SmeltInfo() SmeltInfo {
	if m.Cooked {
		return SmeltInfo{}
	}
	return newFoodSmeltInfo(NewStack(Mutton{Cooked: true}, 1), 0.35)
}

// EncodeItem ...
func (m Mutton) EncodeItem() (name string, meta int16) {
	if m.Cooked {
		return "minecraft:cooked_mutton", 0
	}
	return "minecraft:mutton", 0
}
