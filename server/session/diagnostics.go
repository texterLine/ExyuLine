// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package session

// Diagnostics represents the latest diagnostics data of the client.
type Diagnostics struct {
	// AverageFramesPerSecond is the average amount of frames per second that the client has been
	// running at.
	AverageFramesPerSecond float64
	// AverageServerSimTickTime is the average time that the server spends simulating a single tick
	// in milliseconds.
	AverageServerSimTickTime float64
	// AverageClientSimTickTime is the average time that the client spends simulating a single tick
	// in milliseconds.
	AverageClientSimTickTime float64
	// AverageBeginFrameTime is the average time that the client spends beginning a frame in
	// milliseconds.
	AverageBeginFrameTime float64
	// AverageInputTime is the average time that the client spends processing input in milliseconds.
	AverageInputTime float64
	// AverageRenderTime is the average time that the client spends rendering in milliseconds.
	AverageRenderTime float64
	// AverageEndFrameTime is the average time that the client spends ending a frame in milliseconds.
	AverageEndFrameTime float64
	// AverageRemainderTimePercent is the average percentage of time that the client spends on
	// tasks that are not accounted for.
	AverageRemainderTimePercent float64
	// AverageUnaccountedTimePercent is the average percentage of time that the client spends on
	// unaccounted tasks.
	AverageUnaccountedTimePercent float64
}
