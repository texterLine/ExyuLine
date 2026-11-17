// ExyuLine Software - Copyright (c) 2026 texterLine - https://github.com/texterLine
// All modifications, additions and original work by texterLine are the exclusive intellectual property of
// texterLine. Unauthorized copying, redistribution or removal of this notice is strictly prohibited and
// will be pursued to the full extent of the law.
// Portions of this software derive from Dragonfly (Copyright (c) 2019 Dragonfly Tech), used under the MIT
// License, whose copyright and license notices are preserved in the LICENSE file.

package block

var (
	bannerPatternsMap = map[string]BannerPatternType{}
	bannerPatternIDs  = map[BannerPatternType]string{}
)

// init initialises all default banner patterns to the registry.
func init() {
	registerBannerPattern("bo", BorderBannerPattern())
	registerBannerPattern("bri", BricksBannerPattern())
	registerBannerPattern("mc", CircleBannerPattern())
	registerBannerPattern("cre", CreeperBannerPattern())
	registerBannerPattern("cr", CrossBannerPattern())
	registerBannerPattern("cbo", CurlyBorderBannerPattern())
	registerBannerPattern("lud", DiagonalLeftBannerPattern())
	registerBannerPattern("rd", DiagonalRightBannerPattern())
	registerBannerPattern("ld", DiagonalUpLeftBannerPattern())
	registerBannerPattern("rud", DiagonalUpRightBannerPattern())
	registerBannerPattern("flo", FlowerBannerPattern())
	registerBannerPattern("gra", GradientBannerPattern())
	registerBannerPattern("gru", GradientUpBannerPattern())
	registerBannerPattern("hh", HalfHorizontalBannerPattern())
	registerBannerPattern("hhb", HalfHorizontalBottomBannerPattern())
	registerBannerPattern("vh", HalfVerticalBannerPattern())
	registerBannerPattern("vhr", HalfVerticalRightBannerPattern())
	registerBannerPattern("moj", MojangBannerPattern())
	registerBannerPattern("mr", RhombusBannerPattern())
	registerBannerPattern("sku", SkullBannerPattern())
	registerBannerPattern("ss", SmallStripesBannerPattern())
	registerBannerPattern("bl", SquareBottomLeftBannerPattern())
	registerBannerPattern("br", SquareBottomRightBannerPattern())
	registerBannerPattern("tl", SquareTopLeftBannerPattern())
	registerBannerPattern("tr", SquareTopRightBannerPattern())
	registerBannerPattern("sc", StraightCrossBannerPattern())
	registerBannerPattern("bs", StripeBottomBannerPattern())
	registerBannerPattern("cs", StripeCentreBannerPattern())
	registerBannerPattern("dls", StripeDownLeftBannerPattern())
	registerBannerPattern("drs", StripeDownRightBannerPattern())
	registerBannerPattern("ls", StripeLeftBannerPattern())
	registerBannerPattern("ms", StripeMiddleBannerPattern())
	registerBannerPattern("rs", StripeRightBannerPattern())
	registerBannerPattern("ts", StripeTopBannerPattern())
	registerBannerPattern("bt", TriangleBottomBannerPattern())
	registerBannerPattern("tt", TriangleTopBannerPattern())
	registerBannerPattern("bts", TrianglesBottomBannerPattern())
	registerBannerPattern("tts", TrianglesTopBannerPattern())
	registerBannerPattern("glb", GlobeBannerPattern())
	registerBannerPattern("pig", PiglinBannerPattern())
	registerBannerPattern("flw", FlowBannerPattern())
	registerBannerPattern("gus", GusterBannerPattern())
}

// registerBannerPattern registers a banner pattern with the ID passed.
func registerBannerPattern(id string, pattern BannerPatternType) {
	bannerPatternsMap[id] = pattern
	bannerPatternIDs[pattern] = id
}

// BannerPatternByID returns a banner pattern by the ID it was registered with. Second return value describes whether
// a banner pattern with the ID was found.
func BannerPatternByID(id string) (BannerPatternType, bool) {
	b, ok := bannerPatternsMap[id]
	return b, ok
}

// bannerPatternID returns the ID a banner pattern was registered with.
func bannerPatternID(pattern BannerPatternType) string {
	id, ok := bannerPatternIDs[pattern]
	if !ok {
		panic("should never happen")
	}
	return id
}
