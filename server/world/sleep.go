// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package world

import (
	"github.com/google/uuid"
	"github.com/texterLine/exyuline/server/block/cube"
	"github.com/texterLine/exyuline/server/player/chat"
)

// Sleeper represents an entity that can sleep.
type Sleeper interface {
	Entity

	Name() string
	UUID() uuid.UUID

	Messaget(t chat.Translation, a ...any)
	SendSleepingIndicator(sleeping, max int)

	Sleep(pos cube.Pos)
	Sleeping() (cube.Pos, bool)
	Wake()
}

// Time constants for sleep usage.
const (
	TimeSleep         = 12542
	TimeWake          = 23459
	TimeSleepWithRain = 12010
	TimeWakeWithRain  = 23991
	TimeFull          = 24000
)

// tryAdvanceDay attempts to advance the day of the world, by first ensuring that all sleepers are sleeping, and then
// updating the time of day.
func (ticker) tryAdvanceDay(tx *Tx, timeCycle bool) {
	sleepers := tx.Sleepers()
	time := tx.w.Time() % TimeFull

	for s := range sleepers {
		if !tx.Thundering() {
			if !tx.Raining() && (time <= TimeSleep || time >= TimeWake) {
				return
			}
			if time <= TimeSleepWithRain || time >= TimeWakeWithRain {
				return
			}
		}

		if _, ok := s.Sleeping(); !ok {
			// We can't advance the time - not everyone is sleeping.
			return
		}
	}

	for s := range sleepers {
		s.Wake()
	}

	totalTime := tx.w.Time()
	if timeCycle {
		tx.w.SetTime(totalTime + TimeFull - time)
	}
	tx.w.StopRaining()
}
