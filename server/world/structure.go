// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package world

// Structure represents a structure which may be placed in the world. It has fixed dimensions.
type Structure interface {
	// Dimensions returns the dimensions of the structure. It returns an int array with the width, height and
	// length respectively.
	Dimensions() [3]int
	// At returns the block at a specific location in the structure. When the structure is placed in the
	// world, this method is called for every location within the dimensions of the structure. Additionally,
	// At can return a Liquid to be placed in the same place as the block.
	// At can return nil to not place any block at the position. Returning Air will set any block at that
	// position to air, but returning nil will not do anything.
	// In addition to the coordinates, At will have a function passed that may be used to get a block at a
	// specific position. In scope of At(), structures should use this over World.Block(), due to the way
	// chunks are locked.
	At(x, y, z int, blockAt func(x, y, z int) Block) (Block, Liquid)
}
