// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package nbtconv

import (
	"github.com/texterLine/exyuline/server/item"
	"github.com/texterLine/exyuline/server/item/inventory"
)

// InvFromNBT decodes the data of an NBT slice into the inventory passed.
func InvFromNBT(inv *inventory.Inventory, items []any) {
	for _, itemData := range items {
		data, _ := itemData.(map[string]any)
		it := item.ReadNBT(data, nil)
		if it.Empty() {
			continue
		}
		_ = inv.SetItem(int(Uint8(data, "Slot")), it)
	}
}

// InvToNBT encodes an inventory to a data slice which may be encoded as NBT.
func InvToNBT(inv *inventory.Inventory) []any {
	var items []any
	for index, i := range inv.Slots() {
		if i.Empty() {
			continue
		}
		data := item.WriteNBT(i, true)
		data["Slot"] = byte(index)
		items = append(items, data)
	}
	return items
}
