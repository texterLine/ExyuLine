// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package creative

// Category represents a category of items in the creative inventory which are shown as different tabs.
type Category struct {
	category
}

type category uint8

// ConstructionCategory is the construction category which contains only blocks that do not fall under
// any other category.
func ConstructionCategory() Category {
	return Category{1}
}

// NatureCategory is the nature category which contains blocks and items that can be naturally found in the
// world.
func NatureCategory() Category {
	return Category{2}
}

// EquipmentCategory is the equipment category which contains tools, armour, food and any other form of
// equipment.
func EquipmentCategory() Category {
	return Category{3}
}

// ItemsCategory is the items category for all the miscellaneous items that do not fall under any other
// category.
func ItemsCategory() Category {
	return Category{4}
}

// Uint8 returns the category type as a uint8.
func (s category) Uint8() uint8 {
	return uint8(s)
}
