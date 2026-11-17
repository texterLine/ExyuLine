// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package cmd

// Source represents a source of a command execution. Commands may limit the sources that can run them by
// implementing the Allower interface.
// Source implements Target. A Source must always be able to target itself.
type Source interface {
	Target
	// SendCommandOutput sends a command output to the source. The way the output is applied, depends on what
	// kind of source it is.
	// SendCommandOutput is called by a Command automatically after being run.
	SendCommandOutput(o *Output)
}
