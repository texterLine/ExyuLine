// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package debug

// Renderer represents an interface for a renderer that can manage debug shapes.
type Renderer interface {
	// AddDebugShape adds a debug shape to the renderer, which should be rendered to the player. If the shape
	// already exists, it will be updated with the new information.
	AddDebugShape(shape Shape)
	// RemoveDebugShape removes a debug shape from the renderer by its unique identifier.
	RemoveDebugShape(shape Shape)
	// VisibleDebugShapes returns a slice of all debug shapes that are currently being shown by the renderer.
	VisibleDebugShapes() []Shape
	// RemoveAllDebugShapes clears all debug shapes from the renderer, removing them from the view of the player.
	RemoveAllDebugShapes()
}
