package pokeAPI

const apiEndpoint = "https://pokeapi.co/api/v2/"

const AreaEndpoint = apiEndpoint + "location-area/"

const PokemonEndpoint = apiEndpoint + "pokemon/"

type EnumeratedResp struct {
	Count    int          `json:"count"`
	Next     *string      `json:"next"`
	Previous *string      `json:"previous"`
	Results  []NameAndUrl `json:"results"`
}

type NameAndUrl struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationData struct {
	EncounterMethodRates []EncounterMethodRates `json:"encounter_method_rates"`
	GameIndex            int                    `json:"game_index"`
	ID                   int                    `json:"id"`
	LocationNUrl         NameAndUrl             `json:"location"`
	Name                 string                 `json:"name"`
	NamesNUrl            []NameAndUrl           `json:"names"`
	PokemonEncounters    []PokemonEncounters    `json:"pokemon_encounters"`
}
type EncounterMethodRates struct {
	EncounterMethodNUrl NameAndUrl       `json:"encounter_method"`
	VersionDetails      []VersionDetails `json:"version_details"`
}

type EncounterDetails struct {
	Chance          int        `json:"chance"`
	ConditionValues []any      `json:"condition_values"`
	MaxLevel        int        `json:"max_level"`
	MethodNUrl      NameAndUrl `json:"method"`
	MinLevel        int        `json:"min_level"`
}
type VersionDetails struct {
	EncounterDetails []EncounterDetails `json:"encounter_details"`
	MaxChance        int                `json:"max_chance"`
	Rarity           int                `json:"rarity"`
	VersionNUrl      NameAndUrl         `json:"version"`
}
type PokemonEncounters struct {
	PokemonNUrl    NameAndUrl       `json:"pokemon"`
	VersionDetails []VersionDetails `json:"version_details"`
}

type Pokemon struct {
	Abilities              []Abilities   `json:"abilities"`
	BaseExperience         int           `json:"base_experience"`
	Cries                  Cries         `json:"cries"`
	FormsNUrl              []NameAndUrl  `json:"forms"`
	GameIndices            []GameIndices `json:"game_indices"`
	Height                 int           `json:"height"`
	HeldItems              []HeldItems   `json:"held_items"`
	ID                     int           `json:"id"`
	IsDefault              bool          `json:"is_default"`
	LocationAreaEncounters string        `json:"location_area_encounters"`
	Moves                  []Moves       `json:"moves"`
	Name                   string        `json:"name"`
	Order                  int           `json:"order"`
	PastAbilities          []any         `json:"past_abilities"`
	PastTypes              []any         `json:"past_types"`
	SpeciesNUrl            NameAndUrl    `json:"species"`
	Sprites                Sprites       `json:"sprites"`
	Stats                  []Stats       `json:"stats"`
	Types                  []Types       `json:"types"`
	Weight                 int           `json:"weight"`
}

type Abilities struct {
	AbilityNUrl NameAndUrl `json:"ability"`
	IsHidden    bool       `json:"is_hidden"`
	Slot        int        `json:"slot"`
}
type Cries struct {
	Latest string `json:"latest"`
	Legacy string `json:"legacy"`
}
type GameIndices struct {
	GameIndex   int        `json:"game_index"`
	VersionNUrl NameAndUrl `json:"version"`
}
type HeldItems struct {
	ItemNUrl       NameAndUrl       `json:"item"`
	VersionDetails []VersionDetails `json:"version_details"`
}
type VersionGroupDetails struct {
	LevelLearnedAt      int        `json:"level_learned_at"`
	MoveLearnMethodNUrl NameAndUrl `json:"move_learn_method"`
	Order               any        `json:"order"`
	VersionGroupNUrl    NameAndUrl `json:"version_group"`
}
type Moves struct {
	MoveNUrl            NameAndUrl            `json:"move"`
	VersionGroupDetails []VersionGroupDetails `json:"version_group_details"`
}
type DreamWorld struct {
	FrontDefault string `json:"front_default"`
	FrontFemale  any    `json:"front_female"`
}
type Home struct {
	FrontDefault     string `json:"front_default"`
	FrontFemale      any    `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale any    `json:"front_shiny_female"`
}
type OfficialArtwork struct {
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}
type Showdown struct {
	BackDefault      string `json:"back_default"`
	BackFemale       any    `json:"back_female"`
	BackShiny        string `json:"back_shiny"`
	BackShinyFemale  any    `json:"back_shiny_female"`
	FrontDefault     string `json:"front_default"`
	FrontFemale      any    `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale any    `json:"front_shiny_female"`
}
type Other struct {
	DreamWorld      DreamWorld      `json:"dream_world"`
	Home            Home            `json:"home"`
	OfficialArtwork OfficialArtwork `json:"official-artwork"`
	Showdown        Showdown        `json:"showdown"`
}
type RedBlue struct {
	BackDefault      string `json:"back_default"`
	BackGray         string `json:"back_gray"`
	BackTransparent  string `json:"back_transparent"`
	FrontDefault     string `json:"front_default"`
	FrontGray        string `json:"front_gray"`
	FrontTransparent string `json:"front_transparent"`
}
type Yellow struct {
	BackDefault      string `json:"back_default"`
	BackGray         string `json:"back_gray"`
	BackTransparent  string `json:"back_transparent"`
	FrontDefault     string `json:"front_default"`
	FrontGray        string `json:"front_gray"`
	FrontTransparent string `json:"front_transparent"`
}
type GenerationI struct {
	RedBlue RedBlue `json:"red-blue"`
	Yellow  Yellow  `json:"yellow"`
}
type Crystal struct {
	BackDefault           string `json:"back_default"`
	BackShiny             string `json:"back_shiny"`
	BackShinyTransparent  string `json:"back_shiny_transparent"`
	BackTransparent       string `json:"back_transparent"`
	FrontDefault          string `json:"front_default"`
	FrontShiny            string `json:"front_shiny"`
	FrontShinyTransparent string `json:"front_shiny_transparent"`
	FrontTransparent      string `json:"front_transparent"`
}
type Gold struct {
	BackDefault      string `json:"back_default"`
	BackShiny        string `json:"back_shiny"`
	FrontDefault     string `json:"front_default"`
	FrontShiny       string `json:"front_shiny"`
	FrontTransparent string `json:"front_transparent"`
}
type Silver struct {
	BackDefault      string `json:"back_default"`
	BackShiny        string `json:"back_shiny"`
	FrontDefault     string `json:"front_default"`
	FrontShiny       string `json:"front_shiny"`
	FrontTransparent string `json:"front_transparent"`
}
type GenerationIi struct {
	Crystal Crystal `json:"crystal"`
	Gold    Gold    `json:"gold"`
	Silver  Silver  `json:"silver"`
}
type Emerald struct {
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}
type FireredLeafgreen struct {
	BackDefault  string `json:"back_default"`
	BackShiny    string `json:"back_shiny"`
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}
type RubySapphire struct {
	BackDefault  string `json:"back_default"`
	BackShiny    string `json:"back_shiny"`
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}
type GenerationIii struct {
	Emerald          Emerald          `json:"emerald"`
	FireredLeafgreen FireredLeafgreen `json:"firered-leafgreen"`
	RubySapphire     RubySapphire     `json:"ruby-sapphire"`
}
type DiamondPearl struct {
	BackDefault      string `json:"back_default"`
	BackFemale       any    `json:"back_female"`
	BackShiny        string `json:"back_shiny"`
	BackShinyFemale  any    `json:"back_shiny_female"`
	FrontDefault     string `json:"front_default"`
	FrontFemale      any    `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale any    `json:"front_shiny_female"`
}
type HeartgoldSoulsilver struct {
	BackDefault      string `json:"back_default"`
	BackFemale       any    `json:"back_female"`
	BackShiny        string `json:"back_shiny"`
	BackShinyFemale  any    `json:"back_shiny_female"`
	FrontDefault     string `json:"front_default"`
	FrontFemale      any    `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale any    `json:"front_shiny_female"`
}
type Platinum struct {
	BackDefault      string `json:"back_default"`
	BackFemale       any    `json:"back_female"`
	BackShiny        string `json:"back_shiny"`
	BackShinyFemale  any    `json:"back_shiny_female"`
	FrontDefault     string `json:"front_default"`
	FrontFemale      any    `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale any    `json:"front_shiny_female"`
}
type GenerationIv struct {
	DiamondPearl        DiamondPearl        `json:"diamond-pearl"`
	HeartgoldSoulsilver HeartgoldSoulsilver `json:"heartgold-soulsilver"`
	Platinum            Platinum            `json:"platinum"`
}
type Animated struct {
	BackDefault      string `json:"back_default"`
	BackFemale       any    `json:"back_female"`
	BackShiny        string `json:"back_shiny"`
	BackShinyFemale  any    `json:"back_shiny_female"`
	FrontDefault     string `json:"front_default"`
	FrontFemale      any    `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale any    `json:"front_shiny_female"`
}
type BlackWhite struct {
	Animated         Animated `json:"animated"`
	BackDefault      string   `json:"back_default"`
	BackFemale       any      `json:"back_female"`
	BackShiny        string   `json:"back_shiny"`
	BackShinyFemale  any      `json:"back_shiny_female"`
	FrontDefault     string   `json:"front_default"`
	FrontFemale      any      `json:"front_female"`
	FrontShiny       string   `json:"front_shiny"`
	FrontShinyFemale any      `json:"front_shiny_female"`
}
type GenerationV struct {
	BlackWhite BlackWhite `json:"black-white"`
}
type OmegarubyAlphasapphire struct {
	FrontDefault     string `json:"front_default"`
	FrontFemale      any    `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale any    `json:"front_shiny_female"`
}
type XY struct {
	FrontDefault     string `json:"front_default"`
	FrontFemale      any    `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale any    `json:"front_shiny_female"`
}
type GenerationVi struct {
	OmegarubyAlphasapphire OmegarubyAlphasapphire `json:"omegaruby-alphasapphire"`
	XY                     XY                     `json:"x-y"`
}
type Icons struct {
	FrontDefault string `json:"front_default"`
	FrontFemale  any    `json:"front_female"`
}
type UltraSunUltraMoon struct {
	FrontDefault     string `json:"front_default"`
	FrontFemale      any    `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale any    `json:"front_shiny_female"`
}
type GenerationVii struct {
	Icons             Icons             `json:"icons"`
	UltraSunUltraMoon UltraSunUltraMoon `json:"ultra-sun-ultra-moon"`
}
type GenerationViii struct {
	Icons Icons `json:"icons"`
}
type Versions struct {
	GenerationI    GenerationI    `json:"generation-i"`
	GenerationIi   GenerationIi   `json:"generation-ii"`
	GenerationIii  GenerationIii  `json:"generation-iii"`
	GenerationIv   GenerationIv   `json:"generation-iv"`
	GenerationV    GenerationV    `json:"generation-v"`
	GenerationVi   GenerationVi   `json:"generation-vi"`
	GenerationVii  GenerationVii  `json:"generation-vii"`
	GenerationViii GenerationViii `json:"generation-viii"`
}
type Sprites struct {
	BackDefault      string   `json:"back_default"`
	BackFemale       any      `json:"back_female"`
	BackShiny        string   `json:"back_shiny"`
	BackShinyFemale  any      `json:"back_shiny_female"`
	FrontDefault     string   `json:"front_default"`
	FrontFemale      any      `json:"front_female"`
	FrontShiny       string   `json:"front_shiny"`
	FrontShinyFemale any      `json:"front_shiny_female"`
	Other            Other    `json:"other"`
	Versions         Versions `json:"versions"`
}
type Stats struct {
	BaseStat int        `json:"base_stat"`
	Effort   int        `json:"effort"`
	StatNUrl NameAndUrl `json:"stat"`
}
type Types struct {
	Slot     int        `json:"slot"`
	TypeNUrl NameAndUrl `json:"type"`
}
