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

// HealthBoost causes the affected entity to have its maximum health changed
// for a specific duration.
var HealthBoost healthBoost

type healthBoost struct {
	nopLasting
}

// Start ...
func (healthBoost) Start(e world.Entity, lvl int) {
	if l, ok := e.(living); ok {
		l.SetMaxHealth(l.MaxHealth() + 4*float64(lvl))
	}
}

// End ...
func (healthBoost) End(e world.Entity, lvl int) {
	if l, ok := e.(living); ok {
		l.SetMaxHealth(l.MaxHealth() - 4*float64(lvl))
	}
}

// RGBA ...
func (healthBoost) RGBA() color.RGBA {
	return color.RGBA{R: 0xf8, G: 0x7d, B: 0x23, A: 0xff}
}
