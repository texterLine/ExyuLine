// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package item

import "github.com/sandertv/gophertunnel/minecraft/text"

// Emerald is a rare mineral obtained from emerald ore or from villagers.
type Emerald struct{}

// EncodeItem ...
func (Emerald) EncodeItem() (name string, meta int16) {
	return "minecraft:emerald", 0
}

// TrimMaterial ...
func (Emerald) TrimMaterial() string {
	return "emerald"
}

// MaterialColour ...
func (Emerald) MaterialColour() string {
	return text.Emerald
}

// PayableForBeacon ...
func (Emerald) PayableForBeacon() bool {
	return true
}
