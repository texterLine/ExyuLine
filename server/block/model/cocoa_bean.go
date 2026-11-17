// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package model

import (
	"github.com/texterLine/exyuline/server/block/cube"
	"github.com/texterLine/exyuline/server/world"
)

// CocoaBean is a model used by cocoa bean blocks.
type CocoaBean struct {
	// Facing is the face that the cocoa bean faces. It is the opposite of the face that the CocoaBean is attached to.
	Facing cube.Direction
	// Age is the age of the CocoaBean. The age influences the size of the CocoaBean. The maximum age value of a cocoa
	// bean is 3.
	Age int
}

// BBox returns a single physics.BBox whose size depends on the age of the CocoaBean.
func (c CocoaBean) BBox(cube.Pos, world.BlockSource) []cube.BBox {
	return []cube.BBox{full.
		Stretch(c.Facing.RotateRight().Face().Axis(), -(6-float64(c.Age))/16).
		ExtendTowards(cube.FaceUp, -0.25).
		ExtendTowards(cube.FaceDown, -((7-float64(c.Age)*2)/16)).
		ExtendTowards(c.Facing.Face(), -0.0625).
		ExtendTowards(c.Facing.Opposite().Face(), -((11 - float64(c.Age)*2) / 16))}
}

// FaceSolid always returns false.
func (c CocoaBean) FaceSolid(cube.Pos, cube.Face, world.BlockSource) bool {
	return false
}
