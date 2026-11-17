// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package playerdb

import "github.com/texterLine/exyuline/server/entity/effect"

func effectsToData(effects []effect.Effect) []jsonEffect {
	data := make([]jsonEffect, len(effects))
	for key, eff := range effects {
		id, ok := effect.ID(eff.Type())
		if !ok {
			continue
		}
		data[key] = jsonEffect{
			ID:              id,
			Duration:        eff.Duration(),
			Level:           eff.Level(),
			Ambient:         eff.Ambient(),
			ParticlesHidden: eff.ParticlesHidden(),
			Infinite:        eff.Infinite(),
		}
	}
	return data
}

func dataToEffects(data []jsonEffect) []effect.Effect {
	effects := make([]effect.Effect, len(data))
	for i, d := range data {
		e, ok := effect.ByID(d.ID)
		if !ok {
			continue
		}
		switch eff := e.(type) {
		case effect.LastingType:
			switch {
			case d.Ambient:
				effects[i] = effect.NewAmbient(eff, d.Level, d.Duration)
			case d.Infinite:
				effects[i] = effect.NewInfinite(eff, d.Level)
			default:
				effects[i] = effect.New(eff, d.Level, d.Duration)
			}

			if d.ParticlesHidden {
				effects[i] = effects[i].WithoutParticles()
			}
		default:
			effects[i] = effect.NewInstant(eff, d.Level)
		}
	}
	return effects
}
