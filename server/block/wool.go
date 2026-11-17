// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

import (
	"github.com/texterLine/exyuline/server/item"
	"github.com/texterLine/exyuline/server/world"
	"github.com/texterLine/exyuline/server/world/sound"
)

// Wool is a colourful block that can be obtained by killing/shearing sheep, or crafted using four string.
type Wool struct {
	solid

	// Colour is the colour of the wool.
	Colour item.Colour
}

// Instrument ...
func (w Wool) Instrument() sound.Instrument {
	return sound.Guitar()
}

// FlammabilityInfo ...
func (w Wool) FlammabilityInfo() FlammabilityInfo {
	return newFlammabilityInfo(30, 60, true)
}

// BreakInfo ...
func (w Wool) BreakInfo() BreakInfo {
	return newBreakInfo(0.8, alwaysHarvestable, shearsEffective, oneOf(w))
}

// EncodeItem ...
func (w Wool) EncodeItem() (name string, meta int16) {
	return "minecraft:" + w.Colour.String() + "_wool", 0
}

// EncodeBlock ...
func (w Wool) EncodeBlock() (name string, properties map[string]any) {
	return "minecraft:" + w.Colour.String() + "_wool", nil
}

// allWool returns wool blocks with all possible colours.
func allWool() []world.Block {
	b := make([]world.Block, 0, 16)
	for _, c := range item.Colours() {
		b = append(b, Wool{Colour: c})
	}
	return b
}
