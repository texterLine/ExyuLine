// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package leveldat

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"strconv"
	"strings"
)

// Version is the current version stored in level.dat files.
const Version = 10

// minimumCompatibleClientVersion is the minimum compatible client version,
// required by the latest Minecraft data provider.
var minimumCompatibleClientVersion []int32

// init initialises the minimum compatible client version.
func init() {
	fullVersion := append(strings.Split(protocol.CurrentVersion, "."), "0", "0")
	for _, v := range fullVersion {
		i, _ := strconv.Atoi(v)
		minimumCompatibleClientVersion = append(minimumCompatibleClientVersion, int32(i))
	}
}
