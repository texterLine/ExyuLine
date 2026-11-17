// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package item

import "github.com/texterLine/exyuline/server/world"

// Beef is a food item obtained from cows. It can be cooked in a furnace, smoker, or campfire.
type Beef struct {
	defaultFood

	// Cooked is whether the beef is cooked.
	Cooked bool
}

// Consume ...
func (b Beef) Consume(_ *world.Tx, c Consumer) Stack {
	if b.Cooked {
		c.Saturate(8, 12.8)
	} else {
		c.Saturate(3, 1.8)
	}
	return Stack{}
}

// SmeltInfo ...
func (b Beef) SmeltInfo() SmeltInfo {
	if b.Cooked {
		return SmeltInfo{}
	}
	return newFoodSmeltInfo(NewStack(Beef{Cooked: true}, 1), 0.35)
}

// EncodeItem ...
func (b Beef) EncodeItem() (name string, meta int16) {
	if b.Cooked {
		return "minecraft:cooked_beef", 0
	}
	return "minecraft:beef", 0
}
