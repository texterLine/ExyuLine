// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package item

import (
	"github.com/texterLine/exyuline/server/world/sound"
)

// MusicDisc is an item that can be played in jukeboxes.
type MusicDisc struct {
	// DiscType is the disc type of the music disc.
	DiscType sound.DiscType
}

// MaxCount always returns 1.
func (MusicDisc) MaxCount() int {
	return 1
}

// EncodeItem ...
func (m MusicDisc) EncodeItem() (name string, meta int16) {
	return "minecraft:music_disc_" + m.DiscType.String(), 0
}
