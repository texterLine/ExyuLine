// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

import (
	"fmt"

	"github.com/texterLine/exyuline/server/internal/nbtconv"
	"github.com/texterLine/exyuline/server/item"
)

// BannerPatternLayer is a wrapper over BannerPatternType with a colour property.
type BannerPatternLayer struct {
	// Type represents the type of banner pattern.
	Type BannerPatternType
	// Colour is the colour the pattern should be rendered in.
	Colour item.Colour
}

// EncodeNBT encodes the given BannerPatternLayer into an NBT map.
func (b BannerPatternLayer) EncodeNBT() map[string]any {
	return map[string]any{
		"Pattern": bannerPatternID(b.Type),
		"Color":   int32(invertColour(b.Colour)),
	}
}

// DecodeNBT decodes the given NBT map into a BannerPatternLayer and returns it.
func (b BannerPatternLayer) DecodeNBT(data map[string]any) any {
	id := nbtconv.String(data, "Pattern")
	pattern, exists := BannerPatternByID(id)
	if !exists {
		panic(fmt.Errorf("unknown banner pattern id %q", id))
	}
	b.Type = pattern
	b.Colour = invertColourID(int16(nbtconv.Int32(data, "Color")))
	return b
}
