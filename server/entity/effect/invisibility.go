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

// Invisibility is a lasting effect that causes the affected entity to turn
// invisible. While invisible, the entity's armour is still visible and effect
// particles will still be displayed.
var Invisibility invisibility

type invisibility struct {
	nopLasting
}

// Start ...
func (invisibility) Start(e world.Entity, _ int) {
	if i, ok := e.(interface {
		SetInvisible()
		SetVisible()
	}); ok {
		i.SetInvisible()
	}
}

// End ...
func (invisibility) End(e world.Entity, _ int) {
	if i, ok := e.(interface {
		SetInvisible()
		SetVisible()
	}); ok {
		i.SetVisible()
	}
}

// RGBA ...
func (invisibility) RGBA() color.RGBA {
	return color.RGBA{R: 0xf6, G: 0xf6, B: 0xf6, A: 0xff}
}
