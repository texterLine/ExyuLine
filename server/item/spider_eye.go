// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package item

import (
	"github.com/texterLine/exyuline/server/entity/effect"
	"github.com/texterLine/exyuline/server/world"
	"time"
)

// SpiderEye is a poisonous food and brewing item.
type SpiderEye struct {
	defaultFood
}

// Consume ...
func (SpiderEye) Consume(_ *world.Tx, c Consumer) Stack {
	c.Saturate(2, 3.2)
	c.AddEffect(effect.New(effect.Poison, 1, time.Second*5))
	return Stack{}
}

// EncodeItem ...
func (SpiderEye) EncodeItem() (name string, meta int16) {
	return "minecraft:spider_eye", 0
}
