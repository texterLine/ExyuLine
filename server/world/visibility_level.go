// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package world

// VisibilityLevel controls whether a ViewLayer should use an entity's public visibility or force it
// visible or invisible for a viewer.
type VisibilityLevel struct {
	visibility
}

// PublicVisibility is the default visibility level where the entity is (in)visible
// depending on how it already is publicly viewed.
func PublicVisibility() VisibilityLevel {
	return VisibilityLevel{0}
}

// EnforceInvisible is the visibility level where the entity is always invisible to the viewer.
func EnforceInvisible() VisibilityLevel {
	return VisibilityLevel{1}
}

// EnforceVisible is the visibility level where the entity is always visible to the viewer.
func EnforceVisible() VisibilityLevel {
	return VisibilityLevel{2}
}

type visibility uint8

// EnforceVisibility returns whether metadata should be changed.
func (v visibility) EnforceVisibility() bool {
	return v > 0
}
