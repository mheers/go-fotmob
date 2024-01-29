package player

import "time"

type Player struct {
	ID                  *int                 `json:"id,omitempty"`
	Name                *string              `json:"name,omitempty"`
	BirthDate           *BirthDate           `json:"birthDate,omitempty"`
	IsCoach             *bool                `json:"isCoach,omitempty"`
	IsCaptain           *bool                `json:"isCaptain,omitempty"`
	PrimaryTeam         *PrimaryTeam         `json:"primaryTeam,omitempty"`
	PositionDescription *PositionDescription `json:"positionDescription,omitempty"`
	InjuryInformation   interface{}          `json:"injuryInformation,omitempty"`
	PlayerInformation   []PlayerInformation  `json:"playerInformation,omitempty"`
	MainLeague          *MainLeague          `json:"mainLeague,omitempty"`
	Trophies            *Trophies            `json:"trophies,omitempty"`
	RecentMatches       []RecentMatch        `json:"recentMatches,omitempty"`
	CareerHistory       *CareerHistory       `json:"careerHistory,omitempty"`
	Traits              *Traits              `json:"traits,omitempty"`
	Meta                *Meta                `json:"meta,omitempty"`
	CoachStats          interface{}          `json:"coachStats,omitempty"`
	StatSeasons         []StatSeason         `json:"statSeasons,omitempty"`
}

type BirthDate struct {
	UTCTime  *time.Time `json:"utcTime,omitempty"`
	Timezone *string    `json:"timezone,omitempty"`
}

type CareerHistory struct {
	ShowFootnote *bool        `json:"showFootnote,omitempty"`
	CareerItems  *CareerItems `json:"careerItems,omitempty"`
	FullCareer   *bool        `json:"fullCareer,omitempty"`
}

type CareerItems struct {
	Senior       *NationalTeam `json:"senior,omitempty"`
	Youth        *Youth        `json:"youth,omitempty"`
	NationalTeam *NationalTeam `json:"national team,omitempty"`
}

type NationalTeam struct {
	TeamEntries   []TeamEntry               `json:"teamEntries,omitempty"`
	SeasonEntries []NationalTeamSeasonEntry `json:"seasonEntries,omitempty"`
}

type NationalTeamSeasonEntry struct {
	SeasonName      *string                `json:"seasonName,omitempty"`
	Appearances     *string                `json:"appearances,omitempty"`
	Goals           *string                `json:"goals,omitempty"`
	Assists         *string                `json:"assists,omitempty"`
	Rating          *Rating                `json:"rating,omitempty"`
	TournamentStats []PurpleTournamentStat `json:"tournamentStats,omitempty"`
	TeamID          *int                   `json:"teamId,omitempty"`
	Team            *string                `json:"team,omitempty"`
	TeamGender      *string                `json:"teamGender,omitempty"`
	TransferType    interface{}            `json:"transferType,omitempty"`
}

type Rating struct {
	Num     *string `json:"num,omitempty"`
	Bgcolor *string `json:"bgcolor,omitempty"`
}

type PurpleTournamentStat struct {
	LeagueID     *int    `json:"leagueId,omitempty"`
	TournamentID *int    `json:"tournamentId,omitempty"`
	LeagueName   *string `json:"leagueName,omitempty"`
	SeasonRating *int    `json:"seasonRating,omitempty"`
	IsFriendly   *bool   `json:"isFriendly,omitempty"`
	SeasonName   *string `json:"seasonName,omitempty"`
	Goals        *string `json:"goals,omitempty"`
	Assists      *string `json:"assists,omitempty"`
	Appearances  *string `json:"appearances,omitempty"`
	Rating       *Rating `json:"rating,omitempty"`
}

type TeamEntry struct {
	ParticipantID    *int        `json:"participantId,omitempty"`
	TeamID           *int        `json:"teamId,omitempty"`
	Team             *string     `json:"team,omitempty"`
	TeamGender       *string     `json:"teamGender,omitempty"`
	TransferType     interface{} `json:"transferType,omitempty"`
	StartDate        *time.Time  `json:"startDate,omitempty"`
	EndDate          *time.Time  `json:"endDate,omitempty"`
	Active           *bool       `json:"active,omitempty"`
	Role             interface{} `json:"role,omitempty"`
	Appearances      *string     `json:"appearances,omitempty"`
	Goals            *string     `json:"goals,omitempty"`
	Assists          *string     `json:"assists,omitempty"`
	HasUncertainData *bool       `json:"hasUncertainData,omitempty"`
}

type Youth struct {
	TeamEntries   []TeamEntry        `json:"teamEntries,omitempty"`
	SeasonEntries []YouthSeasonEntry `json:"seasonEntries,omitempty"`
}

type YouthSeasonEntry struct {
	SeasonName      *string                `json:"seasonName,omitempty"`
	Appearances     *string                `json:"appearances,omitempty"`
	Goals           *string                `json:"goals,omitempty"`
	Assists         *string                `json:"assists,omitempty"`
	Rating          *Rating                `json:"rating,omitempty"`
	TournamentStats []FluffyTournamentStat `json:"tournamentStats,omitempty"`
	TeamID          *int                   `json:"teamId,omitempty"`
	Team            *string                `json:"team,omitempty"`
	TeamGender      *string                `json:"teamGender,omitempty"`
	TransferType    interface{}            `json:"transferType,omitempty"`
}

type FluffyTournamentStat struct {
	TournamentID *int    `json:"tournamentId,omitempty"`
	LeagueName   *string `json:"leagueName,omitempty"`
	IsFriendly   *bool   `json:"isFriendly,omitempty"`
	SeasonName   *string `json:"seasonName,omitempty"`
	Goals        *string `json:"goals,omitempty"`
	Assists      *string `json:"assists,omitempty"`
	Appearances  *string `json:"appearances,omitempty"`
	Rating       *Rating `json:"rating,omitempty"`
}

type MainLeague struct {
	TopStats   *TopStats `json:"topStats,omitempty"`
	LeagueID   *int      `json:"leagueId,omitempty"`
	LeagueName *string   `json:"leagueName,omitempty"`
}

type TopStats struct {
	ID      *string        `json:"id,omitempty"`
	Type    *string        `json:"type,omitempty"`
	Display *string        `json:"display,omitempty"`
	Items   []TopStatsItem `json:"items,omitempty"`
}

type TopStatsItem struct {
	Title               *string  `json:"title,omitempty"`
	LocalizedTitleID    *string  `json:"localizedTitleId,omitempty"`
	StatValue           *string  `json:"statValue,omitempty"`
	Per90               *float64 `json:"per90,omitempty"`
	PercentileRank      *float64 `json:"percentileRank,omitempty"`
	PercentileRankPer90 *float64 `json:"percentileRankPer90,omitempty"`
	StatFormat          *string  `json:"statFormat,omitempty"`
}

type Meta struct {
	Seopath          *string           `json:"seopath,omitempty"`
	Pageurl          *string           `json:"pageurl,omitempty"`
	FaqJSONLD        *FAQJSONLD        `json:"faqJSONLD,omitempty"`
	PersonJSONLD     *PersonJSONLD     `json:"personJSONLD,omitempty"`
	BreadcrumbJSONLD *BreadcrumbJSONLD `json:"breadcrumbJSONLD,omitempty"`
}

type BreadcrumbJSONLD struct {
	Context         *string           `json:"@context,omitempty"`
	Type            *string           `json:"@type,omitempty"`
	ItemListElement []ItemListElement `json:"itemListElement,omitempty"`
}

type ItemListElement struct {
	Type     *string `json:"@type,omitempty"`
	Position *int    `json:"position,omitempty"`
	Name     *string `json:"name,omitempty"`
	Item     *string `json:"item,omitempty"`
}

type FAQJSONLD struct {
	Context    *string      `json:"@context,omitempty"`
	Type       *string      `json:"@type,omitempty"`
	MainEntity []MainEntity `json:"mainEntity,omitempty"`
}

type MainEntity struct {
	Type           *string         `json:"@type,omitempty"`
	Name           *string         `json:"name,omitempty"`
	AcceptedAnswer *AcceptedAnswer `json:"acceptedAnswer,omitempty"`
}

type AcceptedAnswer struct {
	Type *string `json:"@type,omitempty"`
	Text *string `json:"text,omitempty"`
}

type PersonJSONLD struct {
	Context     *string      `json:"@context,omitempty"`
	Type        *string      `json:"@type,omitempty"`
	Name        *string      `json:"name,omitempty"`
	BirthDate   *time.Time   `json:"birthDate,omitempty"`
	URL         *string      `json:"url,omitempty"`
	Nationality *Affiliation `json:"nationality,omitempty"`
	Affiliation *Affiliation `json:"affiliation,omitempty"`
	Gender      *string      `json:"gender,omitempty"`
	Height      *Eight       `json:"height,omitempty"`
	Weight      *Eight       `json:"weight,omitempty"`
}

type Affiliation struct {
	Type *string `json:"@type,omitempty"`
	Name *string `json:"name,omitempty"`
}

type Eight struct {
	Type     *string `json:"@type,omitempty"`
	UnitText *string `json:"unitText,omitempty"`
	Value    *string `json:"value,omitempty"`
}

type PlayerInformation struct {
	Value          *Value  `json:"value,omitempty"`
	Title          *string `json:"title,omitempty"`
	TranslationKey *string `json:"translationKey,omitempty"`
	Icon           *Icon   `json:"icon,omitempty"`
	CountryCode    *string `json:"countryCode,omitempty"`
}

type Icon struct {
	Type *string `json:"type,omitempty"`
	ID   *string `json:"id,omitempty"`
}

type Value struct {
	Key      *string   `json:"key,omitempty"`
	Fallback *Fallback `json:"fallback,omitempty"`
}

type Fallback struct {
	Number *float64 `json:"number,omitempty"`
	String *string  `json:"string,omitempty"`
}

type PositionDescription struct {
	Positions           []Position        `json:"positions,omitempty"`
	PrimaryPosition     *PrimaryPosition  `json:"primaryPosition,omitempty"`
	NonPrimaryPositions []PrimaryPosition `json:"nonPrimaryPositions,omitempty"`
}

type PrimaryPosition struct {
	Label *string `json:"label,omitempty"`
	Key   *string `json:"key,omitempty"`
}

type Position struct {
	StrPos            *PrimaryPosition   `json:"strPos,omitempty"`
	StrPosShort       *PrimaryPosition   `json:"strPosShort,omitempty"`
	Occurrences       *int               `json:"occurrences,omitempty"`
	Position          *string            `json:"position,omitempty"`
	IsMainPosition    *bool              `json:"isMainPosition,omitempty"`
	PitchPositionData *PitchPositionData `json:"pitchPositionData,omitempty"`
}

type PitchPositionData struct {
	Right *float64 `json:"right,omitempty"`
	Top   *float64 `json:"top,omitempty"`
	Ratio *float64 `json:"ratio,omitempty"`
}

type PrimaryTeam struct {
	TeamID     *int        `json:"teamId,omitempty"`
	TeamName   *string     `json:"teamName,omitempty"`
	OnLoan     *bool       `json:"onLoan,omitempty"`
	TeamColors *TeamColors `json:"teamColors,omitempty"`
}

type TeamColors struct {
	Color              *string `json:"color,omitempty"`
	ColorAlternate     *string `json:"colorAlternate,omitempty"`
	ColorAway          *string `json:"colorAway,omitempty"`
	ColorAwayAlternate *string `json:"colorAwayAlternate,omitempty"`
}

type RecentMatch struct {
	TeamID           *int         `json:"teamId,omitempty"`
	TeamName         *string      `json:"teamName,omitempty"`
	OpponentTeamID   *int         `json:"opponentTeamId,omitempty"`
	OpponentTeamName *string      `json:"opponentTeamName,omitempty"`
	IsHomeTeam       *bool        `json:"isHomeTeam,omitempty"`
	ID               *int         `json:"id,omitempty"`
	MatchDate        *MatchDate   `json:"matchDate,omitempty"`
	MatchPageURL     *string      `json:"matchPageUrl,omitempty"`
	LeagueID         *int         `json:"leagueId,omitempty"`
	LeagueName       *string      `json:"leagueName,omitempty"`
	Stage            *string      `json:"stage,omitempty"`
	HomeScore        *int         `json:"homeScore,omitempty"`
	AwayScore        *int         `json:"awayScore,omitempty"`
	MinutesPlayed    *int         `json:"minutesPlayed,omitempty"`
	Goals            *int         `json:"goals,omitempty"`
	Assists          *int         `json:"assists,omitempty"`
	YellowCards      *int         `json:"yellowCards,omitempty"`
	RedCards         *int         `json:"redCards,omitempty"`
	RatingProps      *RatingProps `json:"ratingProps,omitempty"`
	PlayerOfTheMatch *bool        `json:"playerOfTheMatch,omitempty"`
	OnBench          *bool        `json:"onBench,omitempty"`
}

type MatchDate struct {
	UTCTime *time.Time `json:"utcTime,omitempty"`
}

type RatingProps struct {
	Num     *Fallback `json:"num,omitempty"`
	Bgcolor *string   `json:"bgcolor,omitempty"`
}

type StatSeason struct {
	SeasonName  *string                `json:"seasonName,omitempty"`
	Tournaments []StatSeasonTournament `json:"tournaments,omitempty"`
}

type StatSeasonTournament struct {
	Name         *string `json:"name,omitempty"`
	TournamentID *int    `json:"tournamentId,omitempty"`
	EntryID      *string `json:"entryId,omitempty"`
}

type Traits struct {
	Key   *string      `json:"key,omitempty"`
	Title *string      `json:"title,omitempty"`
	Items []TraitsItem `json:"items,omitempty"`
}

type TraitsItem struct {
	Key   *string `json:"key,omitempty"`
	Title *string `json:"title,omitempty"`
	Value *int    `json:"value,omitempty"`
}

type Trophies struct {
	PlayerTrophies []PlayerTrophy `json:"playerTrophies,omitempty"`
	CoachTrophies  []interface{}  `json:"coachTrophies,omitempty"`
}

type PlayerTrophy struct {
	Ccode       *string                  `json:"ccode,omitempty"`
	TeamID      *int                     `json:"teamId,omitempty"`
	TeamName    *string                  `json:"teamName,omitempty"`
	Tournaments []PlayerTrophyTournament `json:"tournaments,omitempty"`
}

type PlayerTrophyTournament struct {
	Ccode           *string  `json:"ccode,omitempty"`
	LeagueID        *int     `json:"leagueId,omitempty"`
	LeagueName      *string  `json:"leagueName,omitempty"`
	SeasonsWon      []string `json:"seasonsWon,omitempty"`
	SeasonsRunnerUp []string `json:"seasonsRunnerUp,omitempty"`
}
