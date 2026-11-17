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

// SuspiciousStew is a food item that can give the player a status effect that depends on the flower used to craft it.
type SuspiciousStew struct {
	defaultFood

	// Type specifies the type of effect will be given to the player
	Type StewType
}

// MaxCount ...
func (SuspiciousStew) MaxCount() int {
	return 1
}

// AlwaysConsumable ...
func (SuspiciousStew) AlwaysConsumable() bool {
	return true
}

// EncodeItem ...
func (s SuspiciousStew) EncodeItem() (name string, meta int16) {
	return "minecraft:suspicious_stew", int16(s.Type.Uint8())
}

// Consume ...
func (s SuspiciousStew) Consume(_ *world.Tx, c Consumer) Stack {
	for _, effect := range s.Type.Effects() {
		c.AddEffect(effect)
	}
	c.Saturate(6, 7.2)

	return NewStack(Bowl{}, 1)
}
