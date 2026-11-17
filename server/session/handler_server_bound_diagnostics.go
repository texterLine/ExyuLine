// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package session

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
	"github.com/texterLine/exyuline/server/world"
)

// ServerBoundDiagnosticsHandler handles diagnostic updates from the client.
type ServerBoundDiagnosticsHandler struct{}

// Handle ...
func (h *ServerBoundDiagnosticsHandler) Handle(p packet.Packet, _ *Session, _ *world.Tx, c Controllable) error {
	pk := p.(*packet.ServerBoundDiagnostics)
	c.UpdateDiagnostics(Diagnostics{
		AverageFramesPerSecond:        float64(pk.AverageFramesPerSecond),
		AverageServerSimTickTime:      float64(pk.AverageServerSimTickTime),
		AverageClientSimTickTime:      float64(pk.AverageClientSimTickTime),
		AverageBeginFrameTime:         float64(pk.AverageBeginFrameTime),
		AverageInputTime:              float64(pk.AverageInputTime),
		AverageRenderTime:             float64(pk.AverageRenderTime),
		AverageEndFrameTime:           float64(pk.AverageEndFrameTime),
		AverageRemainderTimePercent:   float64(pk.AverageRemainderTimePercent),
		AverageUnaccountedTimePercent: float64(pk.AverageUnaccountedTimePercent),
	})
	return nil
}
