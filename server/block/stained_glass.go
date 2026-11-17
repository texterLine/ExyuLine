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
)

// StainedGlass is a decorative, fully transparent solid block that is dyed into a different colour.
type StainedGlass struct {
	transparent
	solid
	clicksAndSticks

	// Colour specifies the colour of the block.
	Colour item.Colour
}

// BreakInfo ...
func (g StainedGlass) BreakInfo() BreakInfo {
	return newBreakInfo(0.3, alwaysHarvestable, nothingEffective, silkTouchOnlyDrop(g))
}

// EncodeItem ...
func (g StainedGlass) EncodeItem() (name string, meta int16) {
	return "minecraft:" + g.Colour.String() + "_stained_glass", 0
}

// EncodeBlock ...
func (g StainedGlass) EncodeBlock() (name string, properties map[string]any) {
	return "minecraft:" + g.Colour.String() + "_stained_glass", nil
}

// allStainedGlass returns stained-glass blocks with all possible colours.
func allStainedGlass() []world.Block {
	b := make([]world.Block, 0, 16)
	for _, c := range item.Colours() {
		b = append(b, StainedGlass{Colour: c})
	}
	return b
}
