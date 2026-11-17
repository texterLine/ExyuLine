// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package item

import "time"

// Charcoal is an item obtained by smelting logs or wood.
type Charcoal struct{}

// FuelInfo ...
func (Charcoal) FuelInfo() FuelInfo {
	return newFuelInfo(time.Second * 80)
}

// EncodeItem ...
func (Charcoal) EncodeItem() (name string, meta int16) {
	return "minecraft:charcoal", 0
}
