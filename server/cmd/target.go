// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package cmd

import (
	"github.com/go-gl/mathgl/mgl64"
	"github.com/texterLine/exyuline/server/internal/sliceutil"
	"github.com/texterLine/exyuline/server/world"
	"slices"
)

// Target represents the target of a command. A []Target may be used as command parameter
// types to allow passing targets to the command.
type Target interface {
	// Position returns the position of the Target as an mgl64.Vec3.
	Position() mgl64.Vec3
}

// NamedTarget is a Target that has a name.
type NamedTarget interface {
	Target
	// Name returns a name of the Target. Note that this name needs not to be and is not unique for a Target.
	Name() string
}

// targets returns all Targets selectable by the Source passed.
func targets(tx *world.Tx) (entities []Target, players []NamedTarget) {
	if tx == nil {
		return nil, nil
	}
	ent := sliceutil.Convert[Target](slices.Collect(tx.Entities()))
	pl := sliceutil.Convert[NamedTarget](slices.Collect(tx.Players()))
	return ent, pl
}
