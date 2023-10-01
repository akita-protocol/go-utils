package cp

type Tier string

const (
	Diamond  Tier = "diamond"
	Platinum Tier = "platinum"
	Gold     Tier = "gold"
	Silver   Tier = "silver"
	Bronze   Tier = "bronze"
	None     Tier = ""
)

func GetRanking(score uint64) Tier {
	if score >= 30_000_000_000_000 {
		return Diamond
	} else if score >= 20_000_000_000_000 {
		return Platinum
	} else if score >= 10_000_000_000_000 {
		return Gold
	} else if score >= 5_000_000_000_000 {
		return Silver
	} else if score >= 1_000_000_000_000 {
		return Bronze
	} else {
		return None
	}
}
