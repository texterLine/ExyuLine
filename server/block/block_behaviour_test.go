// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block_test

import (
	"context"
	"testing"

	"github.com/texterLine/exyuline/server/block"
	"github.com/texterLine/exyuline/server/block/cube"
	"github.com/texterLine/exyuline/server/entity"
	"github.com/texterLine/exyuline/server/world"
)

// TestTorchBreaksWithoutSupport verifies that a torch is broken by a neighbour
// update on the tick after its supporting block is removed, using a
// synchronous World to make the tick deterministic.
func TestTorchBreaksWithoutSupport(t *testing.T) {
	w := world.Config{Synchronous: true, Entities: entity.DefaultRegistry}.New()
	defer w.Close()

	support, torch := cube.Pos{0, 0, 0}, cube.Pos{0, 1, 0}
	w.Do(func(tx *world.Tx) {
		tx.SetBlock(support, block.Stone{}, nil)
		tx.SetBlock(torch, block.Torch{Facing: cube.FaceDown}, nil)
		tx.SetBlock(support, block.Air{}, nil)
	})
	w.AdvanceTick()

	b, err := world.Call(context.Background(), w, func(tx *world.Tx) (world.Block, error) {
		return tx.Block(torch), nil
	})
	if err != nil {
		t.Fatalf("read torch block: %v", err)
	}
	if b != (block.Air{}) {
		t.Errorf("expected torch to break after removing its support, got %v", b)
	}
}
