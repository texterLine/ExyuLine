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

// BeetrootSoup is an unstackable food item.
type BeetrootSoup struct {
	defaultFood
}

// MaxCount ...
func (BeetrootSoup) MaxCount() int {
	return 1
}

// Consume ...
func (BeetrootSoup) Consume(_ *world.Tx, c Consumer) Stack {
	c.Saturate(6, 7.2)
	return NewStack(Bowl{}, 1)
}

// EncodeItem ...
func (BeetrootSoup) EncodeItem() (name string, meta int16) {
	return "minecraft:beetroot_soup", 0
}
