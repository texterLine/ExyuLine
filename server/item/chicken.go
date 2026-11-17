// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package item

import (
	"github.com/texterLine/exyuline/server/entity/effect"
	"github.com/texterLine/exyuline/server/world"
	"math/rand/v2"
	"time"
)

// Chicken is a food item obtained from chickens. It can be cooked in a furnace, smoker, or campfire.
type Chicken struct {
	defaultFood

	// Cooked is whether the chicken is cooked.
	Cooked bool
}

// Consume ...
func (c Chicken) Consume(_ *world.Tx, co Consumer) Stack {
	if c.Cooked {
		co.Saturate(6, 7.2)
	} else {
		co.Saturate(2, 1.2)
		if rand.Float64() < 0.3 {
			co.AddEffect(effect.New(effect.Hunger, 1, 30*time.Second))
		}
	}
	return Stack{}
}

// SmeltInfo ...
func (c Chicken) SmeltInfo() SmeltInfo {
	if c.Cooked {
		return SmeltInfo{}
	}
	return newFoodSmeltInfo(NewStack(Chicken{Cooked: true}, 1), 0.35)
}

// EncodeItem ...
func (c Chicken) EncodeItem() (name string, meta int16) {
	if c.Cooked {
		return "minecraft:cooked_chicken", 0
	}
	return "minecraft:chicken", 0
}
