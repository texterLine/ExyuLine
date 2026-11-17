// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package packbuilder

import (
	"os"
	"path/filepath"
	"strings"
)

// buildLanguageFile creates a lang file and writes all of the language entries to the pack.
func buildLanguageFile(dir string, lang []string) {
	if err := os.Mkdir(filepath.Join(dir, "texts"), os.ModePerm); err != nil {
		panic(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "texts/en_US.lang"), []byte(strings.Join(lang, "\n")), 0666); err != nil {
		panic(err)
	}
}
