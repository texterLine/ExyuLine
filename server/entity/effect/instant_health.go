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

// InstantHealth is an instant effect that causes the player that it is applied
// to immediately regain some health. The amount of health regained depends on
// the effect level and potency.
var InstantHealth instantHealth

type instantHealth struct{}

// Apply instantly heals the world.Entity passed for a bit of health, depending on the effect level and
// potency.
func (i instantHealth) Apply(e world.Entity, eff Effect) {
	base := 2 << eff.Level()
	if l, ok := e.(living); ok {
		l.Heal(float64(base)*eff.potency, InstantHealingSource{})
	}
}

// RGBA ...
func (instantHealth) RGBA() color.RGBA {
	return color.RGBA{R: 0xf8, G: 0x24, B: 0x23, A: 0xff}
}

// InstantHealingSource is a healing source used when an entity regains
// health from an effect.instantHealth.
type InstantHealingSource struct{}

func (InstantHealingSource) HealingSource() {}
