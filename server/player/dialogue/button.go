// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package dialogue

import "encoding/json"

// Button represents a button added to a dialogue menu and consists of just
// text.
type Button struct {
	// Text holds the text displayed on the button. It may use Minecraft
	// formatting codes.
	Text string
}

// MarshalJSON ...
func (b Button) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"button_name": b.Text,
		"text":        "",
		"mode":        0, // "Click" activation
		"type":        1, // "Command" type
	})
}
