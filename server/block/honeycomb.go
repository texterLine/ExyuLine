// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

import (
	"github.com/texterLine/exyuline/server/world/sound"
)

// Honeycomb is a decorative blocks crafted from honeycombs.
type Honeycomb struct {
	solid
}

// Instrument ...
func (h Honeycomb) Instrument() sound.Instrument {
	return sound.Flute()
}

// BreakInfo ...
func (h Honeycomb) BreakInfo() BreakInfo {
	return newBreakInfo(0.6, alwaysHarvestable, nothingEffective, oneOf(h))
}

// EncodeItem ...
func (Honeycomb) EncodeItem() (name string, meta int16) {
	return "minecraft:honeycomb_block", 0
}

// EncodeBlock ...
func (Honeycomb) EncodeBlock() (string, map[string]any) {
	return "minecraft:honeycomb_block", nil
}
