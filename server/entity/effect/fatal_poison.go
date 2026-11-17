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

// FatalPoison is a lasting effect that causes the affected entity to lose
// health gradually. fatalPoison, unlike poison, can kill the entity it is
// applied to.
var FatalPoison fatalPoison

type fatalPoison struct {
	nopLasting
}

// Apply ...
func (fatalPoison) Apply(e world.Entity, eff Effect) {
	interval := max(25>>(eff.Level()-1), 1)
	if eff.Tick()%interval == 0 {
		if l, ok := e.(living); ok {
			l.Hurt(1, PoisonDamageSource{Fatal: true})
		}
	}
}

// RGBA ...
func (fatalPoison) RGBA() color.RGBA {
	return color.RGBA{R: 0x4e, G: 0x93, B: 0x31, A: 0xff}
}
