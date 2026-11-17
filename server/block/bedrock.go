// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

// Bedrock is a block that is indestructible in survival.
type Bedrock struct {
	solid
	bassDrum

	// InfiniteBurning specifies if the bedrock block is set aflame and will burn forever. This is the case
	// for bedrock found under end crystals on top of the end pillars.
	InfiniteBurning bool
}

// EncodeItem ...
func (Bedrock) EncodeItem() (name string, meta int16) {
	return "minecraft:bedrock", 0
}

// EncodeBlock ...
func (b Bedrock) EncodeBlock() (name string, properties map[string]any) {
	//noinspection SpellCheckingInspection
	return "minecraft:bedrock", map[string]any{"infiniburn_bit": b.InfiniteBurning}
}

// SupportsEndCrystal always returns true.
func (Bedrock) SupportsEndCrystal() bool {
	return true
}
