// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package effect

import (
	"github.com/texterLine/exyuline/server/world"
	"image/color"
)

// Regeneration is an effect that causes the entity that it is added to slowly
// regenerate health. The level of the effect influences the speed with which
// the entity regenerates.
var Regeneration regeneration

type regeneration struct {
	nopLasting
}

// Apply applies health to the world.Entity passed if the duration of the effect is at the right tick.
func (regeneration) Apply(e world.Entity, eff Effect) {
	interval := max(50>>(eff.Level()-1), 1)
	if eff.Tick()%interval == 0 {
		if l, ok := e.(living); ok {
			l.Heal(1, RegenerationHealingSource{})
		}
	}
}

// RGBA ...
func (regeneration) RGBA() color.RGBA {
	return color.RGBA{R: 0xcd, G: 0x5c, B: 0xab, A: 0xff}
}

// RegenerationHealingSource is a healing source used when an entity regenerates
// health from an effect.regeneration.
type RegenerationHealingSource struct{}

func (RegenerationHealingSource) HealingSource() {}
