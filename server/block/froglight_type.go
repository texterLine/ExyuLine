// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

// FroglightType represents a type of froglight.
type FroglightType struct {
	froglight
}

type froglight uint8

// Pearlescent is the purple variant of froglight.
func Pearlescent() FroglightType {
	return FroglightType{0}
}

// Verdant is the green variant of froglight.
func Verdant() FroglightType {
	return FroglightType{1}
}

// Ochre is the yellow variant of froglight.
func Ochre() FroglightType {
	return FroglightType{2}
}

// Uint8 ...
func (f froglight) Uint8() uint8 {
	return uint8(f)
}

// Name ...
func (f froglight) Name() string {
	switch f {
	case 0:
		return "Pearlescent Froglight"
	case 1:
		return "Verdant Froglight"
	case 2:
		return "Ochre Froglight"
	}
	panic("unknown froglight type")
}

// String ...
func (f froglight) String() string {
	switch f {
	case 0:
		return "pearlescent"
	case 1:
		return "verdant"
	case 2:
		return "ochre"
	}
	panic("unknown froglight type")
}

// FroglightTypes ...
func FroglightTypes() []FroglightType {
	return []FroglightType{Pearlescent(), Verdant(), Ochre()}
}
