// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package entity

// BaseBehaviour provides shared runtime state for Ent behaviours. Embed it
// to inherit common functionality, or forward methods to another instance.
type BaseBehaviour struct {
	portalTravel *PortalTravelComputer
}

// NewBaseBehaviour returns a BaseBehaviour initialised with the default Ent runtime behaviour.
func NewBaseBehaviour() BaseBehaviour {
	return BaseBehaviour{portalTravel: NewPortalTravelComputer()}
}

// PortalTravelComputer returns the portal travel state for a behaviour.
func (b *BaseBehaviour) PortalTravelComputer() *PortalTravelComputer {
	if b.portalTravel == nil {
		b.portalTravel = NewPortalTravelComputer()
	}
	return b.portalTravel
}
