package misc

import (
	"fmt"
	"image"
	"image/draw"
	"math/rand"
	"os"
	"sort"
	"time"

	"gitlab.algotoolchain.com/keelakay/utils/errors"
)

type Head string

const (
	HeadNone             Head = "None"
	HeadAdventureHat     Head = "Adventure Hat"
	HeadAngelHalo        Head = "Angel Halo"
	HeadBeret            Head = "Beret"
	HeadBikerHelmet      Head = "Biker Helmet"
	HeadBucketHat        Head = "Bucket Hat"
	HeadCowboyHat        Head = "Cowboy Hat"
	HeadCrown            Head = "Crown"
	HeadDemonHorns       Head = "Demon Horns"
	HeadDuck             Head = "Duck"
	HeadEarTag           Head = "Ear Tag"
	HeadEgg              Head = "Egg"
	HeadFishermanHat     Head = "Fisherman Hat"
	HeadFlower           Head = "Flower"
	HeadDemonHalo        Head = "Demon Halo"
	HeadLeaf             Head = "Leaf"
	HeadMageHat          Head = "Mage Hat"
	HeadMushroomHead     Head = "Mushroom Head"
	HeadPartyHat         Head = "Party Hat"
	HeadPopeHat          Head = "Pope Hat"
	HeadRubyDemonHorns   Head = "Ruby Demon Horns"
	HeadSaggyBeanie      Head = "Saggy Beanie"
	HeadSailorHat        Head = "Sailor Hat"
	HeadSapphireMageHat  Head = "Sapphire Mage Hat"
	HeadBeanie           Head = "Beanie"
	HeadSkull            Head = "Skull"
	HeadThirdEye         Head = "Third Eye"
	HeadTinyDog          Head = "Tiny Dog"
	HeadUnicornHorn      Head = "Unicorn Horn"
	HeadUshanka          Head = "Ushanka"
	HeadVikingHelmet     Head = "Viking Helmet"
	HeadWatermelonHelmet Head = "Watermelon Helmet"

	// Combos
	HeadDemonHornsWithThirdEye     Head = "Demon Horns, Third Eye"
	HeadRubyDemonHornsWithThirdEye Head = "Ruby Demon Horns, Third Eye"
	HaloEgg                        Head = "Halo, Egg"
	HaloEarTag                     Head = "Halo, Ear Tag"
	HaloWithThirdEye               Head = "Halo, Third Eye"

	// COLLABS / SPECIALS
	BlackCap Head = "Black Cap"
)

var HeadTraits = []Head{
	HeadNone,
	HeadAdventureHat,
	HeadAngelHalo,
	HeadBeret,
	HeadBikerHelmet,
	HeadBucketHat,
	HeadCowboyHat,
	HeadCrown,
	HeadDemonHorns,
	HeadDuck,
	HeadEarTag,
	HeadEgg,
	HeadFishermanHat,
	HeadFlower,
	HeadDemonHalo,
	HeadLeaf,
	HeadMageHat,
	HeadMushroomHead,
	HeadPartyHat,
	HeadPopeHat,
	HeadRubyDemonHorns,
	HeadSaggyBeanie,
	HeadSailorHat,
	HeadSapphireMageHat,
	HeadBeanie,
	HeadSkull,
	HeadThirdEye,
	HeadTinyDog,
	HeadUnicornHorn,
	HeadUshanka,
	HeadVikingHelmet,
	HeadWatermelonHelmet,
	HeadDemonHornsWithThirdEye,
	HeadRubyDemonHornsWithThirdEye,
	HaloEgg,
	HaloEarTag,
	HaloWithThirdEye,
	BlackCap,
}

func (h Head) GetTraits() []Head {
	return HeadTraits
}

type Eyes string

const (
	EyesBase          Eyes = "Base"
	EyesAngry         Eyes = "Angry"
	EyesAviators      Eyes = "Aviators"
	EyesBlackholes    Eyes = "BlackHoles"
	EyesBlindfold     Eyes = "Blindfold"
	EyesBushyBrows    Eyes = "Bushy Brows"
	EyesGlowing       Eyes = "Glowing"
	EyesGoogly        Eyes = "Googly"
	EyesHearts        Eyes = "Hearts"
	EyesJeweler       Eyes = "Jeweler"
	EyesEyePatch      Eyes = "Eye Patch"
	EyesSad           Eyes = "Sad"
	EyesSafetyGoggles Eyes = "Safety Goggles"
	EyesScar          Eyes = "Scar"
	EyesSnorkel       Eyes = "Snorkel"
	EyesSnowGoggles   Eyes = "Snow Goggles"
	EyesSpirals       Eyes = "Spirals"
	EyesSquareGlasses Eyes = "Square Glasses"
	EyesSunglasses    Eyes = "Sunglasses"
	EyesTechGlasses   Eyes = "Tech Glasses"
)

var EyeTraits = []Eyes{
	EyesBase,
	EyesAngry,
	EyesAviators,
	EyesBlackholes,
	EyesBlindfold,
	EyesBushyBrows,
	EyesGlowing,
	EyesGoogly,
	EyesHearts,
	EyesJeweler,
	EyesEyePatch,
	EyesSad,
	EyesSafetyGoggles,
	EyesScar,
	EyesSnorkel,
	EyesSnowGoggles,
	EyesSpirals,
	EyesSquareGlasses,
	EyesSunglasses,
	EyesTechGlasses,
}

func (e Eyes) GetTraits() []Eyes {
	return EyeTraits
}

type Snout string

const (
	SnoutNone               Snout = "None"
	SnoutBone               Snout = "Bone"
	SnoutBubblegum          Snout = "Bubblegum"
	SnoutDagger             Snout = "Dagger"
	SnoutDogToy             Snout = "Dog Toy"
	SnoutDetectivePipe      Snout = "Detective Pipe"
	SnoutDonut              Snout = "Donut"
	SnoutDriftwood          Snout = "Driftwood"
	SnoutDrool              Snout = "Drool"
	SnoutFakeBeard          Snout = "Fake Beard"
	SnoutFootball           Snout = "Football"
	SnoutGasMask            Snout = "Gas Mask"
	SnoutHotDog             Snout = "Hot Dog"
	SnoutKazoo              Snout = "Kazoo"
	SnoutLollyPop           Snout = "LollyPop"
	SnoutMouthGuard         Snout = "MouthGuard"
	SnoutMuzzle             Snout = "Muzzle"
	SnoutNoseRing           Snout = "NoseRing"
	SnoutOldShoe            Snout = "OldShoe"
	SnoutRabbitsFoot        Snout = "RabbitsFoot"
	SnoutRoyalFlush         Snout = "Royal Flush"
	SnoutScar               Snout = "Scar"
	SnoutSideTongue         Snout = "Side Tongue"
	SnoutSnarl              Snout = "Snarl"
	SnoutSpeechBubble       Snout = "Speech Bubble"
	SnoutTopazDetectivePipe Snout = "Topaz Detective Pipe"
	SnoutYoinkBall          Snout = "Yoink Ball"
	SnoutYoYo               Snout = "Yo-yo"
)

var SnoutTraits = []Snout{
	SnoutNone,
	SnoutBone,
	SnoutBubblegum,
	SnoutDagger,
	SnoutDogToy,
	SnoutDetectivePipe,
	SnoutDonut,
	SnoutDriftwood,
	SnoutDrool,
	SnoutFakeBeard,
	SnoutFootball,
	SnoutGasMask,
	SnoutHotDog,
	SnoutKazoo,
	SnoutLollyPop,
	SnoutMouthGuard,
	SnoutMuzzle,
	SnoutNoseRing,
	SnoutOldShoe,
	SnoutRabbitsFoot,
	SnoutRoyalFlush,
	SnoutScar,
	SnoutSideTongue,
	SnoutSnarl,
	SnoutSpeechBubble,
	SnoutTopazDetectivePipe,
	SnoutYoinkBall,
	SnoutYoYo,
}

func (s Snout) GetTraits() []Snout {
	return SnoutTraits
}

type Body string

const (
	BodyNone               Body = "None"
	BodyButcher            Body = "Butcher Apron"
	BodyCamo               Body = "Camo Uniform"
	BodyFlannel            Body = "Flannel"
	BodyCaveman            Body = "Caveman"
	BodyCloak              Body = "Cloak"
	BodyCyberTux           Body = "CyberTux"
	BodyDetective          Body = "Detective Suit"
	BodyDinoSuit           Body = "Dino Suit"
	BodyDiver              Body = "Diving Suit"
	BodyDoctorsCoat        Body = "Doctor Coat"
	BodyEmeraldCloak       Body = "Emerald Cloak"
	BodyAdventuringOutfit  Body = "Adventuring Outfit"
	BodyFishSuit           Body = "Fish Suit"
	BodyFootballUniform    Body = "Football Uniform"
	BodyGoldMedal          Body = "Gold Medal"
	BodyHazmatSuit         Body = "Hazmat Suit"
	BodyTrackSuit          Body = "Tracksuit"
	BodyLeatherArmor       Body = "Leather Armor"
	BodyMageCloak          Body = "Mage Cloak"
	BodyMonkRobes          Body = "Monk Robes"
	BodyPopeRobes          Body = "Pope Robes"
	BodyPuffyJacket        Body = "Puffy Jacket"
	BodyRoyalCape          Body = "Royal Cape"
	BodyRubyCyberTux       Body = "Ruby CyberTux"
	BodySailor             Body = "Sailor Suit"
	BodySapphireMageCloak  Body = "Sapphire Mage Cloak"
	BodyScarf              Body = "Scarf"
	BodyLabCoat            Body = "Lab Coat"
	BodySoccerJersey       Body = "Soccer Jersey"
	BodySweater            Body = "Sweater"
	BodyTieDyeHoodie       Body = "Tie Dye Hoodie"
	BodyTieDyeShirt        Body = "Tie Dye Shirt"
	BodyToga               Body = "Toga"
	BodyTopazDetectiveSuit Body = "Topaz Detective Suit"
	BodyWhiteButtonUp      Body = "White Button Up"

	// COLLABS / SPECIALS
	BodyGreenScarf Body = "Green Scarf"
)

var BodyTraits = []Body{
	BodyNone,
	BodyButcher,
	BodyCamo,
	BodyFlannel,
	BodyCaveman,
	BodyCloak,
	BodyCyberTux,
	BodyDetective,
	BodyDinoSuit,
	BodyDiver,
	BodyDoctorsCoat,
	BodyEmeraldCloak,
	BodyAdventuringOutfit,
	BodyFishSuit,
	BodyFootballUniform,
	BodyGoldMedal,
	BodyHazmatSuit,
	BodyTrackSuit,
	BodyLeatherArmor,
	BodyMageCloak,
	BodyMonkRobes,
	BodyPopeRobes,
	BodyPuffyJacket,
	BodyRoyalCape,
	BodyRubyCyberTux,
	BodySailor,
	BodySapphireMageCloak,
	BodyScarf,
	BodyLabCoat,
	BodySoccerJersey,
	BodySweater,
	BodyTieDyeHoodie,
	BodyTieDyeShirt,
	BodyToga,
	BodyTopazDetectiveSuit,
	BodyWhiteButtonUp,
	BodyGreenScarf,
}

func (b Body) GetTraits() []Body {
	return BodyTraits
}

type Back string

const (
	BackNone              Back = "None"
	BackAnchor            Back = "Anchor"
	BackAngelWings        Back = "Angel Wings"
	BackArrows            Back = "Arrows"
	BackAxe               Back = "Axe"
	BackBalloons          Back = "Balloons"
	BackBaseballBat       Back = "Baseball Bat"
	BackBattery           Back = "Battery"
	BackClub              Back = "Club"
	BackDemonWings        Back = "Demon Wings"
	BackEmeraldArrows     Back = "Emerald Arrows"
	BackFishingPole       Back = "FishingPole"
	BackGuitar            Back = "Guitar"
	BackHikingBackpack    Back = "Hiking Backpack"
	BackMachete           Back = "Machete"
	BackIcePick           Back = "IcePick"
	BackDemonPitchfork    Back = "Demon Pitchfork"
	BackRobotArm          Back = "Robot Arm"
	BackSandwhich         Back = "Sandwhich"
	BackSkateboard        Back = "Skateboard"
	BackSkis              Back = "Skis"
	BackSurfboard         Back = "Surfboard"
	BackWoodenSword       Back = "Wooden Sword"
	BackRadioactiveLiquid Back = "Radioactive Liquid"
	BackRoyalSceptre      Back = "Royal Sceptre"
)

var BackTraits = []Back{
	BackNone,
	BackAnchor,
	BackAngelWings,
	BackArrows,
	BackAxe,
	BackBalloons,
	BackBaseballBat,
	BackBattery,
	BackClub,
	BackDemonWings,
	BackEmeraldArrows,
	BackFishingPole,
	BackGuitar,
	BackHikingBackpack,
	BackMachete,
	BackIcePick,
	BackDemonPitchfork,
	BackRobotArm,
	BackSandwhich,
	BackSkateboard,
	BackSkis,
	BackSurfboard,
	BackWoodenSword,
	BackRadioactiveLiquid,
	BackRoyalSceptre,
}

func (b Back) GetTraits() []Back {
	return BackTraits
}

type Fur string

const (
	FurAngel        Fur = "Angel"
	FurBlack        Fur = "Black"
	FurBlackAndTan  Fur = "Black and Tan"
	FurBlonde       Fur = "Blonde"
	FurBlueBrindle  Fur = "Blue Brindle"
	FurBrownBrindle Fur = "Brown Brindle"
	FurDemon        Fur = "Demon"
	FurDiamond      Fur = "Diamond"
	FurEmerald      Fur = "Emerald"
	FurGoo          Fur = "Goo"
	FurGold         Fur = "Gold"
	FurMarble       Fur = "Marble"
	FurPink         Fur = "Pink"
	FurPlatinum     Fur = "Platinum"
	FurRainbow      Fur = "Rainbow"
	FurRed          Fur = "Red"
	FurRobot        Fur = "Robot"
	FurRuby         Fur = "Ruby"
	FurSapphire     Fur = "Sapphire"
	FurSpotted      Fur = "Spotted"
	FurTopaz        Fur = "Topaz"
)

var FurTraits = []Fur{
	FurAngel,
	FurBlack,
	FurBlackAndTan,
	FurBlonde,
	FurBlueBrindle,
	FurBrownBrindle,
	FurDemon,
	FurDiamond,
	FurEmerald,
	FurGoo,
	FurGold,
	FurMarble,
	FurPink,
	FurPlatinum,
	FurRainbow,
	FurRed,
	FurRobot,
	FurRuby,
	FurSapphire,
	FurSpotted,
	FurTopaz,
}

func (f Fur) GetTraits() []Fur {
	return FurTraits
}

type FurPlus string

const (
	FurPlusNone      FurPlus = "None"
	FurPlusBandanges FurPlus = "Bandages"
	FurPlusScars     FurPlus = "Scars"
	FurPlusStickers  FurPlus = "Stickers"
	FurPlusTattoos   FurPlus = "Tattoos"
)

var FurPlusTraits = []FurPlus{
	FurPlusNone,
	FurPlusBandanges,
	FurPlusScars,
	FurPlusStickers,
	FurPlusTattoos,
}

func (f FurPlus) GetTraits() []FurPlus {
	return FurPlusTraits
}

type Background string

const (
	BackgroundBeach          Background = "Beach"
	BackgroundBlackMarble    Background = "Black Marble"
	BackgroundBlue           Background = "Blue"
	BackgroundDiamond        Background = "Diamond"
	BackgroundEmerald        Background = "Emerald"
	BackgroundGoo            Background = "Goo"
	BackgroundGray           Background = "Gray"
	BackgroundGreen          Background = "Green"
	BackgroundOrange         Background = "Orange"
	BackgroundPurple         Background = "Purple"
	BackgroundRainbow        Background = "Rainbow"
	BackgroundRed            Background = "Red"
	BackgroundRuby           Background = "Ruby"
	BackgroundSapphire       Background = "Sapphire"
	BackgroundSnowyMountains Background = "Snowy Mountains"
	BackgroundSpace          Background = "Space"
	BackgroundTopaz          Background = "Topaz"
	BackgroundYellow         Background = "Yellow"
	BackgroundGraffiti       Background = "Graffiti"
)

var BackgroundTraits = []Background{
	BackgroundBeach,
	BackgroundBlackMarble,
	BackgroundBlue,
	BackgroundDiamond,
	BackgroundEmerald,
	BackgroundGoo,
	BackgroundGray,
	BackgroundGreen,
	BackgroundOrange,
	BackgroundPurple,
	BackgroundRainbow,
	BackgroundRed,
	BackgroundRuby,
	BackgroundSapphire,
	BackgroundSnowyMountains,
	BackgroundSpace,
	BackgroundTopaz,
	BackgroundYellow,
	BackgroundGraffiti,
}

func (b Background) GetTraits() []Background {
	return BackgroundTraits
}

type BackgroundPlus string

const (
	BackgroundPlusNone     BackgroundPlus = "None"
	BackgroundPlusConfetti BackgroundPlus = "Confetti"
	BackgroundPlusMoney    BackgroundPlus = "Money"
)

var BackgroundPlusTraits = []BackgroundPlus{
	BackgroundPlusNone,
	BackgroundPlusConfetti,
	BackgroundPlusMoney,
}

func (b BackgroundPlus) GetTraits() []BackgroundPlus {
	return BackgroundPlusTraits
}

type Rarity string

const (
	RarityCommon    Rarity = "Common"
	RarityUncommon  Rarity = "Uncommon"
	RarityRare      Rarity = "Rare"
	RarityLegendary Rarity = "Legendary"
)

var RarityTraits = []Rarity{
	RarityCommon,
	RarityUncommon,
	RarityRare,
	RarityLegendary,
}

type GenTwoTraits interface {
	Head | Eyes | Snout | Body | Back | Fur | FurPlus | Background | BackgroundPlus | Rarity
}

type TraitsMap struct {
	Head           []Head
	Eyes           []Eyes
	Snout          []Snout
	Body           []Body
	Back           []Back
	Fur            []Fur
	FurPlus        []FurPlus
	Background     []Background
	BackgroundPlus []BackgroundPlus
	Rarity         []Rarity
}

var Traits = TraitsMap{
	Head:           HeadTraits,
	Eyes:           EyeTraits,
	Snout:          SnoutTraits,
	Body:           BodyTraits,
	Back:           BackTraits,
	Fur:            FurTraits,
	FurPlus:        FurPlusTraits,
	Background:     BackgroundTraits,
	BackgroundPlus: BackgroundPlusTraits,
	Rarity:         RarityTraits,
}

type FileMap struct {
	Head           FileMapDetails[Head]
	Eyes           FileMapDetails[Eyes]
	EyesVariant    FileMapDetails[Eyes]
	Snout          FileMapDetails[Snout]
	SnoutVariant   FileMapDetails[Snout]
	Body           FileMapDetails[Body]
	BodyVariant    FileMapDetails[Body]
	Back           FileMapDetails[Back]
	Fur            FileMapDetails[Fur]
	FurNoEars      FileMapDetails[Fur]
	FurPlus        FileMapDetails[FurPlus]
	FurPlusNoEars  FileMapDetails[FurPlus]
	Background     FileMapDetails[Background]
	BackgroundPlus FileMapDetails[BackgroundPlus]
}

type FileMapDetails[V GenTwoTraits] struct {
	Prefix string
	Map    map[V][]string
}

var Files = FileMap{
	Head: FileMapDetails[Head]{
		Prefix: "8_head/",
		Map: map[Head][]string{
			HeadNone:         {""},
			HeadAdventureHat: {"adventure_head.png"},
			HeadAngelHalo:    {"angelhalo_head.png"},
			HeadBeret:        {"beret_head.png"},
			HeadBikerHelmet:  {"biker_head.png"},
			HeadBucketHat:    {"buckethat_head.png"},
			HeadCowboyHat:    {"cowboy_head.png"},
			HeadCrown:        {"crown_head.png"},
			HeadDemonHorns:   {"devilhorn_head.png"},
			// HeadScrubCap:         {"doctor_head.png"},
			HeadDuck:             {"duck_head.png"},
			HeadEarTag:           {"eartag_head.png"},
			HeadEgg:              {"egg_head.png"},
			HeadFishermanHat:     {"fishing_head.png"},
			HeadFlower:           {"flower_head.png"},
			HeadDemonHalo:        {"ironhalo_head.png"},
			HeadLeaf:             {"leaf_head.png"},
			HeadMageHat:          {"mage_head.png"},
			HeadMushroomHead:     {"mushroom_head.png"},
			HeadPartyHat:         {"party_head.png"},
			HeadPopeHat:          {"pope_head.png"},
			HeadRubyDemonHorns:   {"ruby_devilhorn_head.png"},
			HeadSaggyBeanie:      {"saggybeanie_head.png"},
			HeadSailorHat:        {"sailor_head.png"},
			HeadSapphireMageHat:  {"sapphire_mage_head.png"},
			HeadBeanie:           {"seabeanie_head.png"},
			HeadSkull:            {"skull_head.png"},
			HeadThirdEye:         {"thirdeye_head.png"},
			HeadTinyDog:          {"tinydog_head.png"},
			HeadUnicornHorn:      {"unicorn_head.png"},
			HeadUshanka:          {"ushanka_head.png"},
			HeadVikingHelmet:     {"viking_head.png"},
			HeadWatermelonHelmet: {"watermellon_head.png"},
			// combos
			HeadDemonHornsWithThirdEye: {
				"devilhorn_head.png",
				"thirdeye_head.png",
			},
			HeadRubyDemonHornsWithThirdEye: {
				"ruby_devilhorn_head.png",
				"thirdeye_head.png",
			},
			HaloEgg: {
				"angelhalo_head.png",
				"egg_head.png",
			},
			HaloEarTag: {
				"angelhalo_head.png",
				"eartag_head.png",
			},
			HaloWithThirdEye: {
				"angelhalo_head.png",
				"thirdeye_head.png",
			},
			BlackCap: {"blackcap_head.png"},
		},
	},
	Eyes: FileMapDetails[Eyes]{
		Prefix: "4_eyes/",
		Map: map[Eyes][]string{
			EyesBase:          {"base_eye.png"},
			EyesAngry:         {"angry_eye.png"},
			EyesAviators:      {"aviators_eye.png"},
			EyesBlackholes:    {"blackhole_eye.png"},
			EyesBlindfold:     {"blindfold_eye.png"},
			EyesBushyBrows:    {"eyebrow_eye.png"},
			EyesGlowing:       {"Glowing_eye.png"},
			EyesGoogly:        {"google_eye.png"},
			EyesHearts:        {"heart_eye.png"},
			EyesJeweler:       {"Jeweler_eye.png"},
			EyesEyePatch:      {"patch_eye.png"},
			EyesSad:           {"sad_eye.png"},
			EyesSafetyGoggles: {"safety_eye.png"},
			EyesScar:          {"scar_eye.png"},
			// EyesSnorkel:       {"snorkel_eye.png"},
			EyesSnowGoggles:   {"snowgoggle_eye.png"},
			EyesSpirals:       {"spiral_eye.png"},
			EyesSquareGlasses: {"squareglasses_eye.png"},
			EyesSunglasses:    {"sunglasses_eye.png"},
			EyesTechGlasses:   {"tech_eye.png"},
		},
	},
	EyesVariant: FileMapDetails[Eyes]{
		Prefix: "4_eyes/",
		Map: map[Eyes][]string{
			EyesBase:          {"base_eye.png"},
			EyesAngry:         {"angry_eye.png"},
			EyesAviators:      {"aviators_eye.png"},
			EyesBlackholes:    {"blackhole_eye.png"},
			EyesBlindfold:     {"blindfold_eye.png"},
			EyesBushyBrows:    {"eyebrow_eye.png"},
			EyesGlowing:       {"Glowing_eye.png"},
			EyesGoogly:        {"google_eye.png"},
			EyesHearts:        {"heart_eye.png"},
			EyesJeweler:       {"Jeweler_eye.png"},
			EyesEyePatch:      {"patch_eye.png"},
			EyesSad:           {"sad_eye.png"},
			EyesSafetyGoggles: {"safety_eye.png"},
			EyesScar:          {"scar_eye.png"},
			EyesSnorkel: {
				"snorkelback_eye.png",
				"snorkelfront_eye.png",
			},
			EyesSnowGoggles:   {"snowgoggle_eye.png"},
			EyesSpirals:       {"spiral_eye.png"},
			EyesSquareGlasses: {"squareglasses_eye.png"},
			EyesSunglasses:    {"sunglasses_eye.png"},
			EyesTechGlasses:   {"tech_eye.png"},
		},
	},
	Snout: FileMapDetails[Snout]{
		Prefix: "7_snout/",
		Map: map[Snout][]string{
			SnoutNone:               {"base_snout.png"},
			SnoutBone:               {"bone_snout.png"},
			SnoutBubblegum:          {"bubble_snout.png"},
			SnoutDagger:             {"dagger_snout.png"},
			SnoutDogToy:             {"chewtoy_snout.png"},
			SnoutDetectivePipe:      {"detective_snout.png"},
			SnoutDonut:              {"donut_snout.png"},
			SnoutDriftwood:          {"driftwood_snout.png"},
			SnoutDrool:              {"drool_snout.png"},
			SnoutFakeBeard:          {"Fakebeard_snout.png"},
			SnoutFootball:           {"football_snout.png"},
			SnoutGasMask:            {"hazmatback_snout.png"},
			SnoutHotDog:             {"hotdog_snout.png"},
			SnoutKazoo:              {"kazoo_snout.png"},
			SnoutLollyPop:           {"lollypop_snout.png"},
			SnoutMouthGuard:         {"mouthguard_snout.png"},
			SnoutMuzzle:             {"muzzleback_snout.png"},
			SnoutNoseRing:           {"nosering_snout.png"},
			SnoutOldShoe:            {"oldshoe_snout.png"},
			SnoutRabbitsFoot:        {"rabbitsfoot_snout.png"},
			SnoutRoyalFlush:         {"royalflush_snout.png"},
			SnoutScar:               {"scar_snout.png"},
			SnoutSideTongue:         {"sidetongue_snout.png"},
			SnoutSnarl:              {"snarl_snout.png"},
			SnoutSpeechBubble:       {"speechbubble_snout.png"},
			SnoutTopazDetectivePipe: {"topaz_pipe_snout.png"},
			SnoutYoinkBall:          {"Yoink_snout.png"},
			SnoutYoYo:               {"yoyo_snout.png"},
		},
	},
	SnoutVariant: FileMapDetails[Snout]{
		Prefix: "7_snout/",
		Map: map[Snout][]string{
			SnoutNone:          {"base_snout.png"},
			SnoutBone:          {"bone_snout.png"},
			SnoutBubblegum:     {"bubble_snout.png"},
			SnoutDagger:        {"dagger_snout.png"},
			SnoutDogToy:        {"chewtoy_snout.png"},
			SnoutDetectivePipe: {"detective_snout.png"},
			SnoutDonut:         {"donut_snout.png"},
			SnoutDriftwood:     {"driftwood_snout.png"},
			SnoutDrool:         {"drool_snout.png"},
			SnoutFakeBeard:     {"Fakebeard_snout.png"},
			SnoutFootball:      {"football_snout.png"},
			SnoutGasMask: {
				"hazmatback_snout.png",
				"hazmatfront_snout.png",
			},
			SnoutHotDog:     {"hotdog_snout.png"},
			SnoutKazoo:      {"kazoo_snout.png"},
			SnoutLollyPop:   {"lollypop_snout.png"},
			SnoutMouthGuard: {"mouthguard_snout.png"},
			SnoutMuzzle: {
				"muzzleback_snout.png",
				"muzzlefront_snout.png",
			},
			SnoutNoseRing:           {"nosering_snout.png"},
			SnoutOldShoe:            {"oldshoe_snout.png"},
			SnoutRabbitsFoot:        {"rabbitsfoot_snout.png"},
			SnoutRoyalFlush:         {"royalflush_snout.png"},
			SnoutScar:               {"scar_snout.png"},
			SnoutSideTongue:         {"sidetongue_snout.png"},
			SnoutSnarl:              {"snarl_snout.png"},
			SnoutSpeechBubble:       {"speechbubble_snout.png"},
			SnoutTopazDetectivePipe: {"topaz_pipe_snout.png"},
			SnoutYoinkBall:          {"Yoink_snout.png"},
			SnoutYoYo:               {"yoyo_snout.png"},
		},
	},
	Body: FileMapDetails[Body]{
		Prefix: "6_body/",
		Map: map[Body][]string{
			BodyNone:               {""},
			BodyButcher:            {"butcher_body.png"},
			BodyCamo:               {"camo_body.png"},
			BodyFlannel:            {"campshirt_body.png"},
			BodyCaveman:            {"caveman_body.png"},
			BodyCloak:              {"cloak_body.png"},
			BodyCyberTux:           {"Cybertux_body.png"},
			BodyDetective:          {"detective_body.png"},
			BodyDinoSuit:           {"dinosuit_body.png"},
			BodyDiver:              {"diver_body.png"},
			BodyDoctorsCoat:        {"doctor_body.png"},
			BodyEmeraldCloak:       {"emerald_cloak_body.png"},
			BodyAdventuringOutfit:  {"explorer_body.png"},
			BodyFishSuit:           {"fishsuit_body.png"},
			BodyFootballUniform:    {"football_body.png"},
			BodyGoldMedal:          {"goldmedal_body.png"},
			BodyHazmatSuit:         {"hazmat_body.png"},
			BodyTrackSuit:          {"jogger_body.png"},
			BodyLeatherArmor:       {"leatherarmor_body.png"},
			BodyMageCloak:          {"mage_body.png"},
			BodyMonkRobes:          {"monk_body.png"},
			BodyPopeRobes:          {"popeshirt_body.png"},
			BodyPuffyJacket:        {"puffjacket_body.png"},
			BodyRoyalCape:          {"royalcape_body.png"},
			BodyRubyCyberTux:       {"ruby_cybertux_body.png"},
			BodySailor:             {"sailor_body.png"},
			BodySapphireMageCloak:  {"sapphire_mage_body.png"},
			BodyScarf:              {"scarf_body.png"},
			BodyLabCoat:            {"scientist_body.png"},
			BodySoccerJersey:       {"soccer_body.png"},
			BodySweater:            {"sweater_body.png"},
			BodyTieDyeHoodie:       {"Tiedyehood_body.png"},
			BodyTieDyeShirt:        {"Tiedyeshirt_body.png"},
			BodyToga:               {"toga_body.png"},
			BodyTopazDetectiveSuit: {"topaz_detective_body.png"},
			BodyWhiteButtonUp:      {"Whitebutton_body.png"},
			BodyGreenScarf:         {"greenscarf_body.png"},
		},
	},
	BodyVariant: FileMapDetails[Body]{
		Prefix: "6_body/",
		Map: map[Body][]string{
			BodyNone:               {""},
			BodyButcher:            {"butcher_body.png"},
			BodyCamo:               {"camo_body.png"},
			BodyFlannel:            {"campshirt_body.png"},
			BodyCaveman:            {"caveman_body.png"},
			BodyCloak:              {"cloak_body.png"},
			BodyCyberTux:           {"Cybertux_body.png"},
			BodyDetective:          {"detective_body.png"},
			BodyDinoSuit:           {"dinosuit_down_body.png"},
			BodyDiver:              {"diver_body.png"},
			BodyDoctorsCoat:        {"doctor_body.png"},
			BodyEmeraldCloak:       {"emerald_cloak_body.png"},
			BodyAdventuringOutfit:  {"explorer_body.png"},
			BodyFishSuit:           {"fishsuit_down_body.png"},
			BodyFootballUniform:    {"football_body.png"},
			BodyGoldMedal:          {"goldmedal_body.png"},
			BodyHazmatSuit:         {"hazmat_body.png"},
			BodyTrackSuit:          {"jogger_body.png"},
			BodyLeatherArmor:       {"leatherarmor_body.png"},
			BodyMageCloak:          {"mage_body.png"},
			BodyMonkRobes:          {"monk_body.png"},
			BodyPopeRobes:          {"popeshirt_body.png"},
			BodyPuffyJacket:        {"puffjacket_body.png"},
			BodyRoyalCape:          {"royalcape_body.png"},
			BodyRubyCyberTux:       {"ruby_cybertux_body.png"},
			BodySailor:             {"sailor_body.png"},
			BodySapphireMageCloak:  {"sapphire_mage_body.png"},
			BodyScarf:              {"scarf_body.png"},
			BodyLabCoat:            {"scientist_body.png"},
			BodySoccerJersey:       {"soccer_body.png"},
			BodySweater:            {"sweater_body.png"},
			BodyTieDyeHoodie:       {"Tiedyehood_body.png"},
			BodyTieDyeShirt:        {"Tiedyeshirt_body.png"},
			BodyToga:               {"toga_body.png"},
			BodyTopazDetectiveSuit: {"topaz_detective_body.png"},
			BodyWhiteButtonUp:      {"Whitebutton_body.png"},
		},
	},
	Back: FileMapDetails[Back]{
		Prefix: "2_back/",
		Map: map[Back][]string{
			BackNone:              {""},
			BackAnchor:            {"anchor_back.png"},
			BackAngelWings:        {"angel_back.png"},
			BackArrows:            {"arrows_back.png"},
			BackAxe:               {"axe_back.png"},
			BackBalloons:          {"Balloons_back.png"},
			BackBaseballBat:       {"bat_back.png"},
			BackBattery:           {"battery_back.png"},
			BackClub:              {"club_back.png"},
			BackDemonWings:        {"demon_back.png"},
			BackEmeraldArrows:     {"emerald_arrows_back.png"},
			BackFishingPole:       {"fishing_back.png"},
			BackGuitar:            {"guitar_back.png"},
			BackHikingBackpack:    {"hike_back.png"},
			BackMachete:           {"machete_back.png"},
			BackIcePick:           {"pick_back.png"},
			BackDemonPitchfork:    {"pitchfork_back.png"},
			BackRadioactiveLiquid: {"radioactive_back.png"},
			BackRobotArm:          {"roboarm_back.png"},
			BackRoyalSceptre:      {"scepter_back.png"},
			BackSandwhich:         {"sandwich_back.png"},
			BackSkateboard:        {"skate_back.png"},
			BackSkis:              {"ski_back.png"},
			BackSurfboard:         {"surfing_back.png"},
			BackWoodenSword:       {"woodsword_back.png"},
		},
	},
	Fur: FileMapDetails[Fur]{
		Prefix: "3_fur/",
		Map: map[Fur][]string{
			FurAngel:        {"angelic_fur.png"},
			FurBlack:        {"black_fur.png"},
			FurBlackAndTan:  {"blacktan_fur.png"},
			FurBlonde:       {"blonde_fur.png"},
			FurBlueBrindle:  {"blue_fur.png"},
			FurBrownBrindle: {"brown_fur.png"},
			FurDemon:        {"devilish_fur.png"},
			FurDiamond:      {"Diamond_fur.png"},
			FurEmerald:      {"emerald_fur.png"},
			FurGoo:          {"goo_fur.png"},
			FurGold:         {"Gold_fur.png"},
			FurMarble:       {"marble_fur.png"},
			FurPink:         {"pink_fur.png"},
			FurPlatinum:     {"platinum_fur.png"},
			FurRainbow:      {"rainbow_fur.png"},
			FurRed:          {"red_fur.png"},
			FurRobot:        {"robot_fur.png"},
			FurRuby:         {"ruby_fur.png"},
			FurSapphire:     {"sapphire_fur.png"},
			FurSpotted:      {"spotted_fur.png"},
			FurTopaz:        {"topaz_fur.png"},
		},
	},
	FurNoEars: FileMapDetails[Fur]{
		Prefix: "10_noear2/",
		Map: map[Fur][]string{
			FurAngel:        {"angelic_furnoear2.png"},
			FurBlack:        {"black_furnoear2.png"},
			FurBlackAndTan:  {"blacktan_furnoear2.png"},
			FurBlonde:       {"Blonde_furnoear2.png"},
			FurBlueBrindle:  {"blue_furnoear2.png"},
			FurBrownBrindle: {"brown_furnoear2.png"},
			FurDemon:        {"demonic_furnoear2.png"},
			FurDiamond:      {"diamond_furnoear2.png"},
			FurEmerald:      {"emerald_furnoear2.png"},
			FurGoo:          {"goo_furnoear2.png"},
			FurGold:         {"gold_furnoear2.png"},
			FurMarble:       {"marble_furnoear2.png"},
			FurPink:         {"pink_furnoear2.png"},
			FurPlatinum:     {"platinum_furnoear2.png"},
			FurRainbow:      {"rainbow_furnoear2.png"},
			FurRed:          {"red_furnoear2.png"},
			FurRobot:        {"robot_furnoear2.png"},
			FurRuby:         {"ruby_furnoear2.png"},
			FurSapphire:     {"sapphire_furnoear2.png"},
			FurSpotted:      {"spotted_furnoear2.png"},
			FurTopaz:        {"topaz_furnoear2.png"},
		},
	},
	FurPlus: FileMapDetails[FurPlus]{
		Prefix: "5_fur_plus/",
		Map: map[FurPlus][]string{
			FurPlusBandanges: {"bandage_furplus.png"},
			FurPlusScars:     {"scar_furplus.png"},
			FurPlusStickers:  {"stickers_furplus.png"},
			FurPlusTattoos:   {"tattoo_furplus.png"},
		},
	},
	FurPlusNoEars: FileMapDetails[FurPlus]{
		Prefix: "5_fur_plus/",
		Map: map[FurPlus][]string{
			FurPlusBandanges: {"bandage_furplusnoear.png"},
			FurPlusScars:     {"scar_furplus.png"},
			FurPlusStickers:  {"stickers_furplus.png"},
			FurPlusTattoos:   {"tattoo_furplus.png"},
		},
	},
	Background: FileMapDetails[Background]{
		Prefix: "0_backgrounds/",
		Map: map[Background][]string{
			BackgroundBeach:          {"beach_bg.png"},
			BackgroundBlackMarble:    {"black_marble_bg.png"},
			BackgroundBlue:           {"blue_bg.png"},
			BackgroundDiamond:        {"Diamond_bg.png"},
			BackgroundEmerald:        {"emerald_bg.png"},
			BackgroundGoo:            {"goo_bg.png"},
			BackgroundGray:           {"gray_bg.png"},
			BackgroundGreen:          {"green_bg.png"},
			BackgroundOrange:         {"orange_bg.png"},
			BackgroundPurple:         {"purple_bg.png"},
			BackgroundRainbow:        {"rainbow_bg.png"},
			BackgroundRed:            {"red_bg.png"},
			BackgroundRuby:           {"ruby_bg.png"},
			BackgroundSapphire:       {"sapphire_bg.png"},
			BackgroundSnowyMountains: {"snow_bg.png"},
			BackgroundSpace:          {"space_bg.png"},
			BackgroundTopaz:          {"topaz_bg.png"},
			BackgroundYellow:         {"yellow_bg.png"},
			BackgroundGraffiti:       {"graffiti_bg.jpeg"},
		},
	},
	BackgroundPlus: FileMapDetails[BackgroundPlus]{
		Prefix: "1_background_plus/",
		Map: map[BackgroundPlus][]string{
			BackgroundPlusConfetti: {"Confetti_bgplus.png"},
			BackgroundPlusMoney:    {"cash_bgplus.png"},
		},
	},
}

type NormalAkc struct {
	Name           *string        `json:"Name,omitempty"`
	Fur            Fur            `json:"Fur"`
	FurPlus        FurPlus        `json:"Fur+"`
	Head           Head           `json:"Head"`
	Eyes           Eyes           `json:"Eyes"`
	Snout          Snout          `json:"Snout"`
	Body           Body           `json:"Body"`
	Back           Back           `json:"Back"`
	Background     Background     `json:"Background"`
	BackgroundPlus BackgroundPlus `json:"Background+"`
	Rarity         *Rarity        `json:"Rarity,omitempty"`
}

type SpecialAkc struct {
	Name    *string `json:"Name,omitempty"`
	Rarity  *Rarity `json:"Rarity,omitempty"`
	Special *string `json:"Special,omitempty"`
}

type AKC interface {
	NormalAkc | SpecialAkc
}

type LayerType string

const (
	LayerTypeHead           LayerType = "Head"
	LayerTypeEyes           LayerType = "Eyes"
	LayerTypeSnout          LayerType = "Snout"
	LayerTypeBody           LayerType = "Body"
	LayerTypeFur            LayerType = "Fur"
	LayerTypeFurPlus        LayerType = "FurPlus"
	LayerTypeBack           LayerType = "Back"
	LayerTypeBackgroundPlus LayerType = "BackgroundPlus"
	LayerTypeBackground     LayerType = "Background"
)

type WIPAKC struct {
	Head           Head
	Eyes           Eyes
	Snout          Snout
	Body           Body
	Back           Back
	Fur            Fur
	FurPlus        FurPlus
	Background     Background
	BackgroundPlus BackgroundPlus
}

type CompatTypes string

const (
	EarClipping      CompatTypes = "Ear Clipping"
	HoodieDown       CompatTypes = "Hoodie Down"
	BringForward     CompatTypes = "Bring Forward"
	MoveBack         CompatTypes = "Move Back"
	MultiLayerWrap   CompatTypes = "MultiLayerWrap"
	IncompatibleWith CompatTypes = "Incompatible With"
)

type CompatMap struct {
	Head           CompatDetails[Head]
	Eyes           CompatDetails[Eyes]
	Snout          CompatDetails[Snout]
	Body           CompatDetails[Body]
	Back           CompatDetails[Back]
	Fur            CompatDetails[Fur]
	FurPlus        CompatDetails[FurPlus]
	Background     CompatDetails[Background]
	BackgroundPlus CompatDetails[BackgroundPlus]
}

type CompatDetails[V GenTwoTraits] map[V][]CompatTypes

var Compats = CompatMap{
	Head: map[Head][]CompatTypes{
		HeadNone:         {},
		HeadAdventureHat: {EarClipping, HoodieDown, BringForward},
		HeadAngelHalo:    {},
		HeadBeret:        {EarClipping, HoodieDown},
		HeadBikerHelmet:  {EarClipping, HoodieDown},
		HeadBucketHat:    {EarClipping, HoodieDown, BringForward},
		HeadCowboyHat:    {EarClipping, HoodieDown, BringForward},
		HeadCrown:        {HoodieDown},
		HeadDemonHorns:   {HoodieDown},
		// HeadScrubCap:                   {EarClipping, HoodieDown},
		HeadDuck:                       {HoodieDown},
		HeadEarTag:                     {HoodieDown},
		HeadEgg:                        {HoodieDown},
		HeadFishermanHat:               {EarClipping, HoodieDown, BringForward},
		HeadFlower:                     {HoodieDown},
		HeadDemonHalo:                  {},
		HeadLeaf:                       {HoodieDown},
		HeadMageHat:                    {EarClipping, HoodieDown},
		HeadMushroomHead:               {EarClipping, HoodieDown},
		HeadPartyHat:                   {HoodieDown},
		HeadPopeHat:                    {HoodieDown},
		HeadRubyDemonHorns:             {HoodieDown},
		HeadSaggyBeanie:                {EarClipping, HoodieDown},
		HeadSailorHat:                  {HoodieDown},
		HeadSapphireMageHat:            {EarClipping, HoodieDown},
		HeadBeanie:                     {EarClipping, HoodieDown},
		HeadSkull:                      {HoodieDown},
		HeadThirdEye:                   {HoodieDown, MoveBack},
		HeadTinyDog:                    {HoodieDown},
		HeadUnicornHorn:                {HoodieDown},
		HeadUshanka:                    {EarClipping, HoodieDown},
		HeadVikingHelmet:               {EarClipping, HoodieDown},
		HeadWatermelonHelmet:           {EarClipping, HoodieDown},
		HeadDemonHornsWithThirdEye:     {HoodieDown, MoveBack},
		HeadRubyDemonHornsWithThirdEye: {HoodieDown, MoveBack},
		HaloEgg:                        {HoodieDown},
		HaloEarTag:                     {HoodieDown},
		HaloWithThirdEye:               {HoodieDown, MoveBack},
		BlackCap:                       {EarClipping},
	},
	Eyes: map[Eyes][]CompatTypes{
		EyesBase:          {},
		EyesAngry:         {},
		EyesAviators:      {},
		EyesBlackholes:    {},
		EyesBlindfold:     {MoveBack, HoodieDown},
		EyesBushyBrows:    {},
		EyesGlowing:       {},
		EyesGoogly:        {},
		EyesHearts:        {},
		EyesJeweler:       {MoveBack},
		EyesEyePatch:      {MoveBack},
		EyesSad:           {MoveBack},
		EyesSafetyGoggles: {},
		EyesScar:          {MoveBack},
		EyesSnorkel:       {MultiLayerWrap},
		EyesSnowGoggles:   {},
		EyesSpirals:       {MoveBack},
		EyesSquareGlasses: {},
		EyesSunglasses:    {},
		EyesTechGlasses:   {MoveBack},
	},
	Snout: map[Snout][]CompatTypes{
		SnoutNone:               {},
		SnoutBone:               {},
		SnoutBubblegum:          {},
		SnoutDagger:             {},
		SnoutDogToy:             {},
		SnoutDetectivePipe:      {},
		SnoutDonut:              {},
		SnoutDriftwood:          {},
		SnoutDrool:              {},
		SnoutFakeBeard:          {MoveBack},
		SnoutFootball:           {},
		SnoutGasMask:            {MoveBack},
		SnoutHotDog:             {},
		SnoutKazoo:              {},
		SnoutLollyPop:           {},
		SnoutMouthGuard:         {},
		SnoutMuzzle:             {MoveBack},
		SnoutNoseRing:           {},
		SnoutOldShoe:            {},
		SnoutRabbitsFoot:        {},
		SnoutRoyalFlush:         {},
		SnoutScar:               {MoveBack},
		SnoutSideTongue:         {},
		SnoutSnarl:              {MoveBack},
		SnoutSpeechBubble:       {},
		SnoutTopazDetectivePipe: {},
		SnoutYoYo:               {},
	},
	Body: map[Body][]CompatTypes{
		BodyNone:               {},
		BodyButcher:            {},
		BodyCamo:               {},
		BodyFlannel:            {},
		BodyCaveman:            {},
		BodyCloak:              {},
		BodyCyberTux:           {},
		BodyDetective:          {},
		BodyDinoSuit:           {EarClipping},
		BodyDiver:              {},
		BodyDoctorsCoat:        {},
		BodyEmeraldCloak:       {},
		BodyAdventuringOutfit:  {},
		BodyFishSuit:           {EarClipping},
		BodyFootballUniform:    {},
		BodyGoldMedal:          {},
		BodyHazmatSuit:         {},
		BodyTrackSuit:          {},
		BodyLeatherArmor:       {},
		BodyMageCloak:          {},
		BodyMonkRobes:          {},
		BodyPopeRobes:          {},
		BodyPuffyJacket:        {},
		BodyRoyalCape:          {},
		BodyRubyCyberTux:       {},
		BodySailor:             {},
		BodySapphireMageCloak:  {},
		BodyScarf:              {},
		BodyLabCoat:            {},
		BodySoccerJersey:       {},
		BodySweater:            {},
		BodyTieDyeHoodie:       {},
		BodyTieDyeShirt:        {},
		BodyToga:               {},
		BodyTopazDetectiveSuit: {},
		BodyWhiteButtonUp:      {},
	},
	Back: map[Back][]CompatTypes{
		BackNone:              {},
		BackAnchor:            {},
		BackAngelWings:        {},
		BackArrows:            {},
		BackAxe:               {},
		BackBalloons:          {},
		BackBaseballBat:       {},
		BackBattery:           {},
		BackClub:              {},
		BackDemonWings:        {},
		BackEmeraldArrows:     {},
		BackFishingPole:       {},
		BackGuitar:            {},
		BackHikingBackpack:    {},
		BackMachete:           {},
		BackIcePick:           {},
		BackDemonPitchfork:    {},
		BackRadioactiveLiquid: {},
		BackRobotArm:          {},
		BackRoyalSceptre:      {},
		BackSandwhich:         {},
		BackSkateboard:        {},
		BackSkis:              {},
		BackSurfboard:         {},
		BackWoodenSword:       {},
	},
	Fur: map[Fur][]CompatTypes{
		FurAngel:        {},
		FurBlack:        {},
		FurBlackAndTan:  {},
		FurBlonde:       {},
		FurBlueBrindle:  {},
		FurBrownBrindle: {},
		FurDemon:        {},
		FurDiamond:      {},
		FurEmerald:      {},
		FurGold:         {},
		FurMarble:       {},
		FurPink:         {},
		FurPlatinum:     {},
		FurRainbow:      {},
		FurRed:          {},
		FurRobot:        {},
		FurRuby:         {},
		FurSapphire:     {},
		FurSpotted:      {},
		FurTopaz:        {},
	},
	FurPlus: map[FurPlus][]CompatTypes{
		FurPlusBandanges: {},
		FurPlusScars:     {},
		FurPlusStickers:  {},
		FurPlusTattoos:   {},
	},
	Background: map[Background][]CompatTypes{
		BackgroundBeach:          {},
		BackgroundBlackMarble:    {},
		BackgroundBlue:           {},
		BackgroundDiamond:        {},
		BackgroundEmerald:        {},
		BackgroundGoo:            {},
		BackgroundGray:           {},
		BackgroundGreen:          {},
		BackgroundOrange:         {},
		BackgroundPurple:         {},
		BackgroundRainbow:        {},
		BackgroundRed:            {},
		BackgroundRuby:           {},
		BackgroundSapphire:       {},
		BackgroundSnowyMountains: {},
		BackgroundSpace:          {},
		BackgroundTopaz:          {},
		BackgroundYellow:         {},
	},
	BackgroundPlus: map[BackgroundPlus][]CompatTypes{
		BackgroundPlusConfetti: {},
		BackgroundPlusMoney:    {},
	},
}

type IncompatibilityMap struct {
	Head  map[Head][]Eyes
	Body  map[Body][]Eyes
	Snout map[Snout][]Eyes
}

var InCompats = IncompatibilityMap{
	Head: map[Head][]Eyes{
		HeadNone:         {},
		HeadAdventureHat: {EyesSnorkel},
		HeadAngelHalo:    {},
		HeadBeret:        {EyesSnorkel},
		HeadBikerHelmet:  {EyesSafetyGoggles, EyesSnorkel, EyesSnowGoggles, EyesTechGlasses},
		HeadBucketHat:    {},
		HeadCowboyHat:    {EyesSafetyGoggles, EyesSnorkel, EyesSnowGoggles, EyesAviators},
		HeadCrown:        {},
		HeadDemonHorns:   {},
		// HeadScrubCap:                   {},
		HeadDuck:                       {},
		HeadEarTag:                     {},
		HeadEgg:                        {},
		HeadFishermanHat:               {},
		HeadFlower:                     {},
		HeadDemonHalo:                  {},
		HeadLeaf:                       {},
		HeadMageHat:                    {EyesSafetyGoggles, EyesSnorkel, EyesSnowGoggles, EyesBushyBrows},
		HeadMushroomHead:               {EyesSafetyGoggles, EyesSnorkel, EyesSnowGoggles},
		HeadPartyHat:                   {},
		HeadPopeHat:                    {},
		HeadRubyDemonHorns:             {},
		HeadSaggyBeanie:                {},
		HeadSailorHat:                  {},
		HeadSapphireMageHat:            {EyesSafetyGoggles, EyesSnorkel, EyesSnowGoggles, EyesBushyBrows},
		HeadBeanie:                     {EyesSafetyGoggles},
		HeadSkull:                      {},
		HeadThirdEye:                   {EyesSafetyGoggles, EyesSnorkel, EyesSnowGoggles},
		HeadTinyDog:                    {},
		HeadUnicornHorn:                {},
		HeadUshanka:                    {EyesSafetyGoggles, EyesSnorkel, EyesSnowGoggles},
		HeadVikingHelmet:               {EyesSafetyGoggles, EyesSnorkel, EyesSnowGoggles, EyesTechGlasses},
		HeadWatermelonHelmet:           {EyesSafetyGoggles, EyesSnorkel, EyesSnowGoggles, EyesTechGlasses},
		HeadDemonHornsWithThirdEye:     {EyesSafetyGoggles, EyesSnorkel, EyesSnowGoggles},
		HeadRubyDemonHornsWithThirdEye: {EyesSafetyGoggles, EyesSnorkel, EyesSnowGoggles},
		HaloEgg:                        {},
		HaloEarTag:                     {},
		HaloWithThirdEye:               {EyesSafetyGoggles, EyesSnorkel, EyesSnowGoggles},
	},
	Body: map[Body][]Eyes{
		BodyNone:               {},
		BodyButcher:            {},
		BodyCamo:               {},
		BodyFlannel:            {},
		BodyCaveman:            {},
		BodyCloak:              {},
		BodyCyberTux:           {},
		BodyDetective:          {},
		BodyDinoSuit:           {EyesSafetyGoggles, EyesSnorkel, EyesSnowGoggles, EyesTechGlasses},
		BodyDiver:              {},
		BodyDoctorsCoat:        {},
		BodyEmeraldCloak:       {},
		BodyAdventuringOutfit:  {},
		BodyFishSuit:           {EyesSafetyGoggles, EyesSnorkel, EyesSnowGoggles, EyesTechGlasses},
		BodyFootballUniform:    {},
		BodyGoldMedal:          {},
		BodyHazmatSuit:         {},
		BodyTrackSuit:          {},
		BodyLeatherArmor:       {},
		BodyMageCloak:          {},
		BodyMonkRobes:          {},
		BodyPopeRobes:          {},
		BodyPuffyJacket:        {},
		BodyRoyalCape:          {},
		BodyRubyCyberTux:       {},
		BodySailor:             {},
		BodySapphireMageCloak:  {},
		BodyScarf:              {},
		BodyLabCoat:            {},
		BodySoccerJersey:       {},
		BodySweater:            {},
		BodyTieDyeHoodie:       {},
		BodyTieDyeShirt:        {},
		BodyToga:               {},
		BodyTopazDetectiveSuit: {},
		BodyWhiteButtonUp:      {},
	},
}

func needsEarClipping(h Head, b Body) bool {
	if InSlice(EarClipping, Compats.Head[h]) || (InSlice(EarClipping, Compats.Body[b]) && !InSlice(HoodieDown, Compats.Head[h])) {
		return true
	}
	return false
}

func needsBodyVariant(h Head, b Body) bool {
	if (b == BodyDinoSuit || b == BodyFishSuit) && InSlice(HoodieDown, Compats.Head[h]) {
		return true
	}
	return false
}

func GetRandomTrait[V GenTwoTraits](lt LayerType) (V, error) {
	randSource := rand.NewSource(time.Now().UnixNano())
	r := rand.New(randSource)
	switch lt {
	case LayerTypeHead:
		l := len(HeadTraits)
		traitIndex := r.Intn(l)
		t := HeadTraits[traitIndex]
		return V(t), nil
	case LayerTypeEyes:
		l := len(EyeTraits)
		traitIndex := r.Intn(l)
		t := EyeTraits[traitIndex]
		return V(t), nil
	case LayerTypeSnout:
		l := len(SnoutTraits)
		traitIndex := r.Intn(l)
		t := SnoutTraits[traitIndex]
		return V(t), nil
	case LayerTypeBody:
		l := len(BodyTraits)
		traitIndex := r.Intn(l)
		t := BodyTraits[traitIndex]
		return V(t), nil
	case LayerTypeBack:
		l := len(BackTraits)
		traitIndex := r.Intn(l)
		t := BackTraits[traitIndex]
		return V(t), nil
	case LayerTypeFur:
		l := len(FurTraits)
		traitIndex := r.Intn(l)
		t := FurTraits[traitIndex]
		return V(t), nil
	case LayerTypeFurPlus:
		l := len(FurPlusTraits)
		traitIndex := r.Intn(l)
		t := FurPlusTraits[traitIndex]
		return V(t), nil
	case LayerTypeBackground:
		l := len(BackgroundTraits)
		traitIndex := r.Intn(l)
		t := BackgroundTraits[traitIndex]
		return V(t), nil
	case LayerTypeBackgroundPlus:
		l := len(BackgroundPlusTraits)
		traitIndex := r.Intn(l)
		t := BackgroundPlusTraits[traitIndex]
		return V(t), nil
	default:
		return V(HeadBeanie), fmt.Errorf("not sure what trait this is")
	}
}

func getRandomTraitWithOddsMap[V GenTwoTraits](lt LayerType, od OddsMap) (V, error) {
	randSource := rand.NewSource(time.Now().UnixNano())
	r := rand.New(randSource)
	switch lt {
	case LayerTypeHead:
		list := insideOutOddsMap(od.Head)
		traitIndex := r.Intn(99)
		t := list[traitIndex]
		return V(t), nil
	case LayerTypeEyes:
		list := insideOutOddsMap(od.Eyes)
		traitIndex := r.Intn(99)
		t := list[traitIndex]
		return V(t), nil
	case LayerTypeSnout:
		list := insideOutOddsMap(od.Snout)
		traitIndex := r.Intn(99)
		t := list[traitIndex]
		return V(t), nil
	case LayerTypeBody:
		list := insideOutOddsMap(od.Body)
		traitIndex := r.Intn(99)
		t := list[traitIndex]
		return V(t), nil
	case LayerTypeBack:
		list := insideOutOddsMap(od.Back)
		traitIndex := r.Intn(99)
		t := list[traitIndex]
		return V(t), nil
	case LayerTypeFur:
		list := insideOutOddsMap(od.Fur)
		traitIndex := r.Intn(99)
		t := list[traitIndex]
		return V(t), nil
	case LayerTypeFurPlus:
		list := insideOutOddsMap(od.FurPlus)
		traitIndex := r.Intn(99)
		t := list[traitIndex]
		return V(t), nil
	case LayerTypeBackground:
		list := insideOutOddsMap(od.Background)
		traitIndex := r.Intn(99)
		t := list[traitIndex]
		return V(t), nil
	case LayerTypeBackgroundPlus:
		list := insideOutOddsMap(od.BackgroundPlus)
		traitIndex := r.Intn(99)
		t := list[traitIndex]
		return V(t), nil
	default:
		return V(HeadBeanie), fmt.Errorf("not sure what trait this is")
	}
}

func getTraitImages[V GenTwoTraits](t LayerType, trait V) ([]image.Image, error) {
	data := []image.Image{}
	var (
		filePaths []string
		prefix    string
	)

	switch t {
	case LayerTypeHead:
		prefix = Files.Head.Prefix
		filePaths = Files.Head.Map[Head(trait)]

	case LayerTypeEyes:
		prefix = Files.Eyes.Prefix
		filePaths = Files.Eyes.Map[Eyes(trait)]

	case LayerTypeSnout:
		prefix = Files.Snout.Prefix
		filePaths = Files.Snout.Map[Snout(trait)]

	case LayerTypeBody:
		prefix = Files.Body.Prefix
		filePaths = Files.Body.Map[Body(trait)]

	case LayerTypeBack:
		prefix = Files.Back.Prefix
		filePaths = Files.Back.Map[Back(trait)]

	case LayerTypeFur:
		prefix = Files.Fur.Prefix
		filePaths = Files.Fur.Map[Fur(trait)]

	case LayerTypeFurPlus:
		prefix = Files.FurPlus.Prefix
		filePaths = Files.FurPlus.Map[FurPlus(trait)]

	case LayerTypeBackground:
		prefix = Files.Background.Prefix
		filePaths = Files.Background.Map[Background(trait)]

	case LayerTypeBackgroundPlus:
		prefix = Files.BackgroundPlus.Prefix
		filePaths = Files.BackgroundPlus.Map[BackgroundPlus(trait)]
	}

	for _, path := range filePaths {
		if path == "" {
			continue
		}

		fileRoute := fmt.Sprintf("./assets/%v%v", prefix, path)
		// fmt.Println(fileRoute)

		catFile, err := os.Open(fileRoute)
		if err != nil {
			fmt.Println(err)
			return data, err
		}
		defer catFile.Close()

		imgData, _, err := image.Decode(catFile)
		if err != nil {
			fmt.Println(err)
			return data, err
		}
		data = append(data, imgData)
	}

	return data, nil
}

func getClippedFurEar(f Fur) ([]image.Image, error) {
	prefix := Files.FurNoEars.Prefix
	filePaths := Files.FurNoEars.Map[f]
	data := []image.Image{}
	for _, path := range filePaths {
		if path == "" {
			continue
		}

		fileRoute := fmt.Sprintf("./assets/%v%v", prefix, path)
		// fmt.Println(fileRoute)

		catFile, err := os.Open(fileRoute)
		if err != nil {
			fmt.Println(err)
			return data, err
		}
		defer catFile.Close()

		imgData, _, err := image.Decode(catFile)
		if err != nil {
			fmt.Println(err)
			return data, err
		}
		data = append(data, imgData)
	}

	return data, nil
}

func getClippedEarFurPlus(fp FurPlus) ([]image.Image, error) {
	prefix := Files.FurPlusNoEars.Prefix
	filePaths := Files.FurPlusNoEars.Map[fp]
	data := []image.Image{}
	for _, path := range filePaths {
		if path == "" {
			continue
		}

		fileRoute := fmt.Sprintf("./assets/%v%v", prefix, path)
		// fmt.Println(fileRoute)

		catFile, err := os.Open(fileRoute)
		if err != nil {
			fmt.Println(err)
			return data, err
		}
		defer catFile.Close()

		imgData, _, err := image.Decode(catFile)
		if err != nil {
			fmt.Println(err)
			return data, err
		}
		data = append(data, imgData)
	}

	return data, nil
}

func getBodyVariant(b Body) ([]image.Image, error) {
	prefix := Files.BodyVariant.Prefix
	filePaths := Files.BodyVariant.Map[b]
	data := []image.Image{}
	for _, path := range filePaths {
		if path == "" {
			continue
		}

		fileRoute := fmt.Sprintf("./assets/%v%v", prefix, path)
		// fmt.Println(fileRoute)

		catFile, err := os.Open(fileRoute)
		if err != nil {
			fmt.Println(err)
			return data, err
		}
		defer catFile.Close()

		imgData, _, err := image.Decode(catFile)
		if err != nil {
			fmt.Println(err)
			return data, err
		}
		data = append(data, imgData)
	}

	return data, nil
}

func getSnorkelImages() ([]image.Image, error) {
	prefix := Files.EyesVariant.Prefix
	filePaths := Files.EyesVariant.Map[EyesSnorkel]
	data := []image.Image{}
	for _, path := range filePaths {
		if path == "" {
			continue
		}

		fileRoute := fmt.Sprintf("./assets/%v%v", prefix, path)
		// fmt.Println(fileRoute)

		catFile, err := os.Open(fileRoute)
		if err != nil {
			fmt.Println(err)
			return data, err
		}
		defer catFile.Close()

		imgData, _, err := image.Decode(catFile)
		if err != nil {
			fmt.Println(err)
			return data, err
		}
		data = append(data, imgData)
	}

	return data, nil
}

func getSnoutImages(trait Snout) ([]image.Image, error) {
	prefix := Files.SnoutVariant.Prefix
	filePaths := Files.SnoutVariant.Map[trait]
	data := []image.Image{}
	for _, path := range filePaths {
		if path == "" {
			continue
		}

		fileRoute := fmt.Sprintf("./assets/%v%v", prefix, path)
		// fmt.Println(fileRoute)

		catFile, err := os.Open(fileRoute)
		if err != nil {
			fmt.Println(err)
			return data, err
		}
		defer catFile.Close()

		imgData, _, err := image.Decode(catFile)
		if err != nil {
			fmt.Println(err)
			return data, err
		}
		data = append(data, imgData)
	}

	return data, nil
}

func CheckAllOdds() error {
	const op errors.Op = "CheckAllOdds"
	err := mapCheck(RubyOdds)
	if err != nil {
		return errors.E(op, err)
	}

	err = mapCheck(SapphireOdds)
	if err != nil {
		return errors.E(op, err)
	}

	err = mapCheck(TopazOdds)
	if err != nil {
		return errors.E(op, err)
	}

	err = mapCheck(EmeraldOdds)
	if err != nil {
		return errors.E(op, err)
	}

	err = mapCheck(NormalOdds)
	if err != nil {
		return errors.E(op, err)
	}
	return nil
}

func GetRandomAKC(ot OddsTable) (*WIPAKC, error) {
	const op errors.Op = "GetRandomAKC"
	var (
		oddsMap OddsMap
		err     error
	)
	r := WIPAKC{}

	// ensure odds add up to 100
	switch ot {
	// case Diamond:
	case Ruby:
		oddsMap = RubyOdds
	case Sapphire:
		oddsMap = SapphireOdds
	case Topaz:
		oddsMap = TopazOdds
	case Emerald:
		oddsMap = EmeraldOdds
	case Normal:
		oddsMap = NormalOdds
	}

	r.Background, err = getRandomTraitWithOddsMap[Background](LayerTypeBackground, oddsMap)
	if err != nil {
		return nil, err
	}

	r.BackgroundPlus, err = getRandomTraitWithOddsMap[BackgroundPlus](LayerTypeBackgroundPlus, oddsMap)
	if err != nil {
		return nil, err
	}

	r.Back, err = getRandomTraitWithOddsMap[Back](LayerTypeBack, oddsMap)
	if err != nil {
		return nil, err
	}

	r.Fur, err = getRandomTraitWithOddsMap[Fur](LayerTypeFur, oddsMap)
	if err != nil {
		return nil, err
	}

	r.FurPlus, err = getRandomTraitWithOddsMap[FurPlus](LayerTypeFurPlus, oddsMap)
	if err != nil {
		return nil, err
	}

	r.Body, err = getRandomTraitWithOddsMap[Body](LayerTypeBody, oddsMap)
	if err != nil {
		return nil, err
	}

	r.Head, err = getRandomTraitWithOddsMap[Head](LayerTypeHead, oddsMap)
	if err != nil {
		return nil, err
	}

	valid := false
	for !valid {
		r.Eyes, err = getRandomTraitWithOddsMap[Eyes](LayerTypeEyes, oddsMap)
		if err != nil {
			return nil, err
		}

		if !InSlice(r.Eyes, InCompats.Head[r.Head]) && !InSlice(r.Eyes, InCompats.Body[r.Body]) {
			valid = true
		}
	}

	r.Snout, err = getRandomTraitWithOddsMap[Snout](LayerTypeSnout, oddsMap)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

type WIPAKCIMGS struct {
	Prepend        []image.Image
	Head           WIPIMG
	Eyes           WIPIMG
	Snout          WIPIMG
	Body           WIPIMG
	Back           WIPIMG
	Fur            WIPIMG
	FurPlus        WIPIMG
	Background     WIPIMG
	BackgroundPlus WIPIMG
	Append         []image.Image
}

func (w *WIPAKCIMGS) ToSlice() []image.Image {
	s := []WIPIMG{w.Head, w.Eyes, w.Snout, w.Body, w.Back, w.Fur, w.FurPlus, w.Background, w.BackgroundPlus}
	sort.Slice(s, func(i, j int) bool {
		return s[i].Order < s[j].Order
	})
	r := []image.Image{}

	for _, t := range s {
		r = append(r, t.Imgs...)
	}

	if len(w.Prepend) > 0 {
		if len(w.BackgroundPlus.Imgs) > 0 {
			r = append(r[:3], r[2:]...)
			r[2] = w.Prepend[0]
		} else {
			r = append(r[:2], r[1:]...)
			r[1] = w.Prepend[0]
		}
	}

	if len(w.Append) > 0 {
		r = append(r, w.Append...)
	}

	return r
}

type WIPIMG struct {
	Imgs  []image.Image
	Order int
}

func BuildImg(wip *WIPAKC) (*image.RGBA, error) {
	const op errors.Op = "BuildRandomAKC"
	var err error
	wipImgs := WIPAKCIMGS{
		Background:     WIPIMG{Order: 0},
		BackgroundPlus: WIPIMG{Order: 1},
		Back:           WIPIMG{Order: 2},
		Fur:            WIPIMG{Order: 3},
		FurPlus:        WIPIMG{Order: 4},
		Body:           WIPIMG{Order: 5},
		Head:           WIPIMG{Order: 6},
		Eyes:           WIPIMG{Order: 7},
		Snout:          WIPIMG{Order: 8},
	}

	wipImgs.Background.Imgs, err = getTraitImages(LayerTypeBackground, wip.Background)
	if err != nil {
		return nil, errors.E(op, err)
	}

	wipImgs.BackgroundPlus.Imgs, err = getTraitImages(LayerTypeBackgroundPlus, wip.BackgroundPlus)
	if err != nil {
		return nil, errors.E(op, err)
	}

	wipImgs.Back.Imgs, err = getTraitImages(LayerTypeBack, wip.Back)
	if err != nil {
		return nil, errors.E(op, err)
	}

	wipImgs.Fur.Imgs, err = getTraitImages(LayerTypeFur, wip.Fur)
	if err != nil {
		return nil, errors.E(op, err)
	}

	wipImgs.FurPlus.Imgs, err = getTraitImages(LayerTypeFurPlus, wip.FurPlus)
	if err != nil {
		return nil, errors.E(op, err)
	}

	wipImgs.Body.Imgs, err = getTraitImages(LayerTypeBody, wip.Body)
	if err != nil {
		return nil, errors.E(op, err)
	}

	wipImgs.Head.Imgs, err = getTraitImages(LayerTypeHead, wip.Head)
	if err != nil {
		return nil, errors.E(op, err)
	}

	wipImgs.Eyes.Imgs, err = getTraitImages(LayerTypeEyes, wip.Eyes)
	if err != nil {
		return nil, errors.E(op, err)
	}

	wipImgs.Snout.Imgs, err = getTraitImages(LayerTypeSnout, wip.Snout)
	if err != nil {
		return nil, errors.E(op, err)
	}

	clip := needsEarClipping(wip.Head, wip.Body)
	if clip {
		wipImgs.Fur.Imgs, err = getClippedFurEar(wip.Fur)
		if err != nil {
			return nil, errors.E(op, err)
		}

		if wip.FurPlus == FurPlusBandanges {
			wipImgs.FurPlus.Imgs, err = getClippedEarFurPlus(wip.FurPlus)
			if err != nil {
				return nil, errors.E(op, err)
			}
		}
	}

	if needsBodyVariant(wip.Head, wip.Body) {
		wipImgs.Body.Imgs, err = getBodyVariant(wip.Body)
		if err != nil {
			return nil, errors.E(op, err)
		}
	}

	if InSlice(BringForward, Compats.Head[wip.Head]) {
		wipImgs.Eyes.Order = 6
		wipImgs.Snout.Order = 7
		wipImgs.Head.Order = 8
	} else if InSlice(MoveBack, Compats.Eyes[wip.Eyes]) && InSlice(MoveBack, Compats.Snout[wip.Snout]) {
		wipImgs.Snout.Order = 6
		wipImgs.Eyes.Order = 7
		wipImgs.Head.Order = 8
	} else if InSlice(MoveBack, Compats.Eyes[wip.Eyes]) {
		wipImgs.Eyes.Order = 6
		wipImgs.Head.Order = 7
		wipImgs.Snout.Order = 8
	} else if InSlice(MoveBack, Compats.Snout[wip.Snout]) {
		wipImgs.Head.Order = 6
		wipImgs.Snout.Order = 7
		wipImgs.Eyes.Order = 8

		if wip.Snout == SnoutMuzzle || wip.Snout == SnoutGasMask {
			wipImgs.Snout.Order = 6
			wipImgs.Head.Order = 7
			wipImgs.Eyes.Order = 8
		}
	}

	if wip.Eyes == EyesSnorkel {

		newEyesImg, err := getSnorkelImages()
		if err != nil {
			return nil, errors.E(op, err)
		}
		wipImgs.Prepend = []image.Image{newEyesImg[0]}
		wipImgs.Eyes.Imgs = []image.Image{newEyesImg[1]}
	}

	if wip.Snout == SnoutMuzzle || wip.Snout == SnoutGasMask {

		newSnoutImg, err := getSnoutImages(wip.Snout)
		if err != nil {
			return nil, errors.E(op, err)
		}

		wipImgs.Snout.Imgs = []image.Image{newSnoutImg[0]}
		wipImgs.Append = []image.Image{newSnoutImg[1]}
	}

	images := wipImgs.ToSlice()
	base := image.NewRGBA(images[0].Bounds())
	for _, img := range images {
		draw.Draw(base, img.Bounds(), img, image.Point{}, draw.Over)
	}

	return base, nil
}
