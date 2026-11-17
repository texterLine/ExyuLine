// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package generator

import (
	"github.com/texterLine/exyuline/server/block/cube"
	"github.com/texterLine/exyuline/server/world"
	"github.com/texterLine/exyuline/server/world/chunk"
)

// Flat is the flat generator of World. It generates flat worlds (like those in vanilla) with no other
// decoration. It may be constructed by calling NewFlat.
type Flat struct {
	// biome is the encoded biome that the generator should use.
	biome uint32
	// layers is a list of block runtime ID layers placed by the Flat generator. The layers are ordered in a way where
	// the last element in the slice is placed as the bottom-most block of the chunk.
	layers []uint32
}

// NewFlat creates a new Flat generator. Chunks generated are completely filled with the world.Biome passed. layers is a
// list of block layers placed by the Flat generator. The layers are ordered in a way where the last element in the
// slice is placed as the bottom-most block of the chunk.
func NewFlat(biome world.Biome, layers []world.Block) Flat {
	return NewFlatWithRegistry(biome, layers, world.DefaultBlockRegistry)
}

// NewFlatWithRegistry creates a new Flat generator using the block registry passed to resolve layers to runtime IDs.
// Use this constructor when the generator is used in a World with a non-default block registry.
func NewFlatWithRegistry(biome world.Biome, layers []world.Block, br world.BlockRegistry) Flat {
	f := Flat{
		biome:  uint32(biome.EncodeBiome()),
		layers: make([]uint32, len(layers)),
	}
	for i, b := range layers {
		f.layers[i] = br.BlockRuntimeID(b)
	}
	return f
}

// GenerateChunk ...
func (f Flat) GenerateChunk(pos world.ChunkPos, chunk *chunk.Chunk) {
	min, max := int16(chunk.Range().Min()), int16(chunk.Range().Max())
	n := int16(len(f.layers))

	for x := range uint8(16) {
		for z := range uint8(16) {
			for y := int16(0); y <= max; y++ {
				if y < n {
					chunk.SetBlock(x, min+y, z, 0, f.layers[n-y-1])
				}
				chunk.SetBiome(x, min+y, z, f.biome)
			}
		}
	}
}

// DefaultSpawn ...
func (f Flat) DefaultSpawn(dim world.Dimension) cube.Pos {
	return cube.Pos{0, dim.Range().Min() + len(f.layers) + 1, 0}
}
