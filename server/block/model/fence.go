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

// Fence is a model used by fences of any type. It can attach to blocks with solid faces and other fences of the same
// type and has a model height just slightly over 1.
type Fence struct {
	// Wood specifies if the Fence is made from wood. This field is used to check if two fences are able to attach to
	// each other.
	Wood bool
}

const (
	fenceHeight = 1.5
	fenceInset  = 0.375
)

// BBox returns multiple physics.BBox depending on how many connections it has with the surrounding blocks.
func (f Fence) BBox(pos cube.Pos, s world.BlockSource) []cube.BBox {
	boxes := make([]cube.BBox, 0, 2)

	connectWest, connectEast := f.Connects(pos, cube.FaceWest, s), f.Connects(pos, cube.FaceEast, s)
	connectNorth, connectSouth := f.Connects(pos, cube.FaceNorth, s), f.Connects(pos, cube.FaceSouth, s)

	// Check if we have any connections on the Z axis
	if connectWest || connectEast {
		sideBox := cube.Box(0, 0, 0, 1, fenceHeight, 1).Stretch(cube.Z, -fenceInset)
		if connectWest {
			boxes = append(boxes, sideBox.ExtendTowards(cube.FaceEast, -fenceInset))
		}
		if connectEast {
			boxes = append(boxes, sideBox.ExtendTowards(cube.FaceWest, -fenceInset))
		}
	}

	// Check if we have any connections on the X axis
	if connectNorth || connectSouth {
		sideBox := cube.Box(0, 0, 0, 1, fenceHeight, 1).Stretch(cube.X, -fenceInset)
		if connectNorth {
			boxes = append(boxes, sideBox.ExtendTowards(cube.FaceSouth, -fenceInset))
		}
		if connectSouth {
			boxes = append(boxes, sideBox.ExtendTowards(cube.FaceNorth, -fenceInset))
		}
	}

	// If no connections, create a center post box
	if len(boxes) == 0 {
		boxes = append(boxes, cube.Box(fenceInset, 0, fenceInset, 1-fenceInset, fenceHeight, 1-fenceInset))
	}

	return boxes
}

// FaceSolid returns true if the face is cube.FaceDown or cube.FaceUp.
func (f Fence) FaceSolid(_ cube.Pos, face cube.Face, _ world.BlockSource) bool {
	return face == cube.FaceDown || face == cube.FaceUp
}

// Connects returns true if the Fence connects to the block at the face passed.
// A Fence connects to fences of the same kind, to fence gates and to blocks
// with a solid face towards it.
func (f Fence) Connects(pos cube.Pos, face cube.Face, src world.BlockSource) bool {
	sidePos := pos.Side(face)
	sideBlock := src.Block(sidePos)
	if fence, ok := sideBlock.Model().(Fence); ok && fence.Wood == f.Wood {
		return true
	}
	if sideBlock.Model().FaceSolid(sidePos, face.Opposite(), src) {
		return true
	}
	_, ok := sideBlock.Model().(FenceGate)
	return ok
}
