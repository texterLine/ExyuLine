// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package session

import (
	"fmt"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/texterLine/exyuline/server/item/recipe"
	"github.com/texterLine/exyuline/server/world"
)

// stonecutterInputSlot is the slot index of the input item in the stonecutter.
const stonecutterInputSlot = 0x03

// handleStonecutting handles a CraftRecipe stack request action made using a stonecutter.
func (h *ItemStackRequestHandler) handleStonecutting(a *protocol.CraftRecipeStackRequestAction, s *Session, tx *world.Tx) error {
	craft, ok := s.recipes[a.RecipeNetworkID]
	if !ok {
		return fmt.Errorf("recipe with network id %v does not exist", a.RecipeNetworkID)
	}
	if _, shapeless := craft.(recipe.Shapeless); !shapeless {
		return fmt.Errorf("recipe with network id %v is not a shapeless recipe", a.RecipeNetworkID)
	}
	if craft.Block() != "stonecutter" {
		return fmt.Errorf("recipe with network id %v is not a stonecutter recipe", a.RecipeNetworkID)
	}

	timesCrafted := int(a.NumberOfCrafts)
	if timesCrafted < 1 {
		return fmt.Errorf("times crafted must be at least 1")
	}

	expectedInputs := craft.Input()
	input, _ := h.itemInSlot(protocol.StackRequestSlotInfo{
		Container: protocol.FullContainerName{ContainerID: protocol.ContainerStonecutterInput},
		Slot:      stonecutterInputSlot,
	}, s, tx)
	if input.Count() < timesCrafted {
		return fmt.Errorf("input item count is less than number of crafts")
	}
	if !matchingStacks(input, expectedInputs[0]) {
		return fmt.Errorf("input item is not the same as expected input")
	}

	output := craft.Output()
	h.setItemInSlot(protocol.StackRequestSlotInfo{
		Container: protocol.FullContainerName{ContainerID: protocol.ContainerStonecutterInput},
		Slot:      stonecutterInputSlot,
	}, input.Grow(-timesCrafted), s, tx)
	return h.createResults(s, tx, repeatStacks(output, timesCrafted)...)
}
