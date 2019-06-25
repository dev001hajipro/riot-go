package common

// LOLQueue is league type.
type LOLQueue string

const (
	// RankedSolo5x5 is a Ranked match.
	RankedSolo5x5 LOLQueue = "RANKED_SOLO_5x5"
	// RankedFlexSR is for team match.
	RankedFlexSR LOLQueue = "RANKED_FLEX_SR"
	// RankedFlexTT is for team match.
	RankedFlexTT LOLQueue = "RANKED_FLEX_TT"
)

// Tier is a tier.
type Tier string

const (
	// Diamond tier.
	Diamond Tier = "DIAMOND"
	// Platinum tier.
	Platinum Tier = "PLATINUM"
	// Gold tier.
	Gold Tier = "GOLD"
	// Silver tier.
	Silver Tier = "SILVER"
	// Bronze tier.
	Bronze Tier = "BRONZE"
	// Iron tier.
	Iron Tier = "IRON"
)

// Division is a divison.
type Division string

const (
	// DivisionI is I.
	DivisionI Division = "I"
	// DivisionII is II.
	DivisionII Division = "II"
	// DivisionIII is III.
	DivisionIII Division = "III"
	// DivisionIV is IV.
	DivisionIV Division = "IV"
)

type Region string

const (
	RU   Region = "RU"
	KR   Region = "KR"
	PBE1 Region = "PBE1"
	BR1  Region = "BR1"
	OC1  Region = "OC1"
	JP1  Region = "JP1"
	NA1  Region = "NA1"
	EUN1 Region = "EUN1"
	EUW1 Region = "EUW1"
	TR1  Region = "TR1"
	LA1  Region = "LA1"
	LA2  Region = "LA2"
)
