// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

// GrindstoneAttachment represents a type of attachment for a Grindstone.
type GrindstoneAttachment struct {
	grindstoneAttachment
}

// StandingGrindstoneAttachment is a type of attachment for a standing Grindstone.
func StandingGrindstoneAttachment() GrindstoneAttachment {
	return GrindstoneAttachment{0}
}

// HangingGrindstoneAttachment is a type of attachment for a hanging Grindstone.
func HangingGrindstoneAttachment() GrindstoneAttachment {
	return GrindstoneAttachment{1}
}

// WallGrindstoneAttachment is a type of attachment for a wall Grindstone.
func WallGrindstoneAttachment() GrindstoneAttachment {
	return GrindstoneAttachment{2}
}

// GrindstoneAttachments returns all possible GrindstoneAttachments.
func GrindstoneAttachments() []GrindstoneAttachment {
	return []GrindstoneAttachment{StandingGrindstoneAttachment(), HangingGrindstoneAttachment(), WallGrindstoneAttachment()}
}

type grindstoneAttachment uint8

// Uint8 returns the GrindstoneAttachment as a uint8.
func (g grindstoneAttachment) Uint8() uint8 {
	return uint8(g)
}

// String returns the GrindstoneAttachment as a string.
func (g grindstoneAttachment) String() string {
	switch g {
	case 0:
		return "standing"
	case 1:
		return "hanging"
	case 2:
		return "side"
	}
	panic("should never happen")
}
