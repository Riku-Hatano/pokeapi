package structs

type ResponseStatsMoves struct {
	Accuracy      int `json:"accuracy"`
	ContestCombos struct {
		Name      string `json:"name"`
		UseBefore string `json:"use_before"`
		UseAfter  string `json:"use_after"`
	} `json:"contest_combos"`
	CotestEffect struct {
		Url string `json:"url"`
	} `json:"contest_effect"`
	ContestType struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"contest_type"`
	DamageClass struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"damage_class"`
	// EffectChange []struct {}
	EffectChance  int `json:"effect_chance"`
	EffectEntries []struct {
		Effect   string `json:"effect"`
		Language struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"language"`
		ShortEffect string `json:"short_effect"`
	} `json:"effect_entries"`
	FalvorTextEntries []struct {
		FlavorText string `json:"flavor_text"`
		Language   struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"language"`
		VersionGroup struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"version_group"`
	} `json:"flavor_text_entries"`
	Generation struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"generation"`
	Id               int `json:"id"`
	LearnedByPokemon []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"learned_by_pokemon"`
	Machines []struct {
		Machine struct {
			Url string `json:"url"`
		} `json:"machine"`
		VersionGroup struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"version_group"`
	} `json:"machines"`
	Meta struct {
		Aliment struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"aliment"`
		AlimentChance int `json:"aliment_chance"`
		Category      struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"category"`
		CritRate     int `json:"crit_rate"`
		Drain        int `json:"drain"`
		FlinchChance int `json:"flinch_chance"`
		Healing      int `json:"healing"`
		MaxHits      int `json:"max_hits"`
		MaxTurns     int `json:"max_turns"`
		MinHits      int `json:"min_hits"`
		MinTurns     int `json:"min_turns"`
		StatChance   int `json:"stat_chance"`
	} `json:"meta"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PastValues []struct {
		Accuracy     int `json:"accuracy"`
		EffectChance int `json:"effect_chance"`
		// EffectEntries
		Power int `json:"power"`
		PP    int `json:"pp"`
		Type  struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"type"`
		VesionGroup struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"version_group"`
	} `json:"past_values"`
	Power       int `json:"power"`
	PP          int `json:"pp"`
	Priority    int `json:"priority"`
	StatChanges []struct {
		Change int `json:"change"`
		Stat   struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"stat"`
	} `json:"stat_changes"`
	SuperContestEffect struct {
		Url string `json:"url"`
	} `json:"super_contest_effect"`
	Target struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"target"`
	Type struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"type"`
}
