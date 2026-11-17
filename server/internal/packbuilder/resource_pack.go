// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package packbuilder

import (
	_ "embed"
	"os"

	"github.com/sandertv/gophertunnel/minecraft/resource"
	"github.com/texterLine/exyuline/server/world"
	"golang.org/x/mod/sumdb/dirhash"
)

//go:embed pack_icon.png
var packIcon []byte

// BuildResourcePack builds a resource pack based on custom features that have been registered to the server.
// It creates a UUID based on the hash of the directory so the client will only be prompted to download it
// once it is changed.
func BuildResourcePack(reg world.BlockRegistry) (*resource.Pack, bool) {
	dir, err := os.MkdirTemp("", "exyuline_resource_pack-")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(dir)

	var assets int
	var lang []string

	itemCount, itemLang := buildItems(dir)
	assets += itemCount
	lang = append(lang, itemLang...)

	blockCount, blockLang := buildBlocks(reg, dir)
	assets += blockCount
	lang = append(lang, blockLang...)

	if assets > 0 {
		buildLanguageFile(dir, lang)
		if err := os.WriteFile(dir+"/pack_icon.png", packIcon, 0666); err != nil {
			panic(err)
		}
		hash, err := dirhash.HashDir(dir, "", dirhash.Hash1)
		if err != nil {
			panic(err)
		}
		var header, module [16]byte
		copy(header[:], hash)
		copy(module[:], hash[16:])
		buildManifest(dir, header, module)
		return resource.MustReadPath(dir), true
	}
	return nil, false
}
