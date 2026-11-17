// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

// InvisibleBedrock is an indestructible, solid block, similar to bedrock and has the appearance of air.
// It shares many of its properties with barriers.
type InvisibleBedrock struct {
	transparent
	solid
}

// EncodeItem ...
func (InvisibleBedrock) EncodeItem() (name string, meta int16) {
	return "minecraft:invisible_bedrock", 0
}

// EncodeBlock ...
func (InvisibleBedrock) EncodeBlock() (string, map[string]any) {
	return "minecraft:invisible_bedrock", nil
}
