// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package item

import (
	"github.com/texterLine/exyuline/server/world"
)

// RabbitStew is a food item that can be eaten by the player.
type RabbitStew struct {
	defaultFood
}

// MaxCount ...
func (RabbitStew) MaxCount() int {
	return 1
}

// Consume ...
func (RabbitStew) Consume(_ *world.Tx, c Consumer) Stack {
	c.Saturate(10, 12)
	return NewStack(Bowl{}, 1)
}

// EncodeItem ...
func (RabbitStew) EncodeItem() (name string, meta int16) {
	return "minecraft:rabbit_stew", 0
}
