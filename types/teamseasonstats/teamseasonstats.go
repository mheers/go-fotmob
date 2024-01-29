package teamseasonstats

type TeamSeasonStats struct {
	Players           []Player           `json:"players,omitempty"`
	Teams             []Team             `json:"teams,omitempty"`
	TournamentSeasons []TournamentSeason `json:"tournamentSeasons,omitempty"`
	CurrentTournament *string            `json:"currentTournament,omitempty"`
	TournamentStats   *TournamentStats   `json:"tournamentStats,omitempty"`
	TournamentID      *string            `json:"tournamentId,omitempty"`
}

type Player struct {
	Header           *string             `json:"header,omitempty"`
	Participant      *PlayerParticipant  `json:"participant,omitempty"`
	FetchAllURL      *string             `json:"fetchAllUrl,omitempty"`
	TopThree         []PlayerParticipant `json:"topThree,omitempty"`
	Order            *int                `json:"order,omitempty"`
	Name             *string             `json:"name,omitempty"`
	LocalizedTitleID *string             `json:"localizedTitleId,omitempty"`
}

type PlayerParticipant struct {
	ID                 *int        `json:"id,omitempty"`
	Name               *string     `json:"name,omitempty"`
	Rank               *int        `json:"rank,omitempty"`
	Goals              interface{} `json:"goals,omitempty"`
	Assists            interface{} `json:"assists,omitempty"`
	Rating             interface{} `json:"rating,omitempty"`
	PositionID         interface{} `json:"positionId,omitempty"`
	Ccode              interface{} `json:"ccode,omitempty"`
	Cname              interface{} `json:"cname,omitempty"`
	TeamID             *int        `json:"teamId,omitempty"`
	TeamName           *string     `json:"teamName,omitempty"`
	ShowRole           interface{} `json:"showRole,omitempty"`
	ShowCountryFlag    interface{} `json:"showCountryFlag,omitempty"`
	ShowTeamFlag       *bool       `json:"showTeamFlag,omitempty"`
	ExcludeFromRanking *bool       `json:"excludeFromRanking,omitempty"`
	Value              interface{} `json:"value,omitempty"`
	TeamColors         *TeamColors `json:"teamColors,omitempty"`
}

type TeamColors struct {
	DarkMode      *string `json:"darkMode,omitempty"`
	LightMode     *string `json:"lightMode,omitempty"`
	FontDarkMode  *string `json:"fontDarkMode,omitempty"`
	FontLightMode *string `json:"fontLightMode,omitempty"`
}

type Team struct {
	Header           *string           `json:"header,omitempty"`
	LocalizedTitleID *string           `json:"localizedTitleId,omitempty"`
	Participant      *TeamParticipant  `json:"participant,omitempty"`
	FetchAllURL      *string           `json:"fetchAllUrl,omitempty"`
	TopThree         []TeamParticipant `json:"topThree,omitempty"`
	Order            *int              `json:"order,omitempty"`
	Stat             *string           `json:"stat,omitempty"`
}

type TeamParticipant struct {
	ID                 *int        `json:"id,omitempty"`
	Name               *string     `json:"name,omitempty"`
	Rank               *int        `json:"rank,omitempty"`
	Goals              interface{} `json:"goals,omitempty"`
	Assists            interface{} `json:"assists,omitempty"`
	Rating             interface{} `json:"rating,omitempty"`
	PositionID         interface{} `json:"positionId,omitempty"`
	Ccode              interface{} `json:"ccode,omitempty"`
	Cname              interface{} `json:"cname,omitempty"`
	TeamID             *int        `json:"teamId,omitempty"`
	TeamName           interface{} `json:"teamName,omitempty"`
	ShowRole           interface{} `json:"showRole,omitempty"`
	ShowCountryFlag    interface{} `json:"showCountryFlag,omitempty"`
	ShowTeamFlag       *bool       `json:"showTeamFlag,omitempty"`
	ExcludeFromRanking *bool       `json:"excludeFromRanking,omitempty"`
	Value              interface{} `json:"value,omitempty"`
	TeamColors         *TeamColors `json:"teamColors,omitempty"`
}

type TournamentSeason struct {
	Name           *string     `json:"name,omitempty"`
	Season         *string     `json:"season,omitempty"`
	LeagueName     *LeagueName `json:"leagueName,omitempty"`
	TournamentID   *string     `json:"tournamentId,omitempty"`
	ParentLeagueID *string     `json:"parentLeagueId,omitempty"`
}

type LeagueName string

const (
	PremierLeague LeagueName = "Premier League"
	EuropaLeague  LeagueName = "Europa League"
	FACup         LeagueName = "FA Cup"
	EFLCup        LeagueName = "EFL Cup"
)

type TournamentStats struct {
	Table              *Table             `json:"table,omitempty"`
	Goals              *Goals             `json:"goals,omitempty"`
	Rating             *float64           `json:"rating,omitempty"`
	AverageStageRating *float64           `json:"averageStageRating,omitempty"`
	TeamID             *int               `json:"teamId,omitempty"`
	StageID            *int               `json:"stageId,omitempty"`
	RelegationInfo     *RelegationInfo    `json:"relegationInfo,omitempty"`
	CleanSheets        *int               `json:"cleanSheets,omitempty"`
	YellowCards        *int               `json:"yellowCards,omitempty"`
	RedCards           *int               `json:"redCards,omitempty"`
	TeamName           *string            `json:"teamName,omitempty"`
	GoalsAgainst       *Goals             `json:"goalsAgainst,omitempty"`
	TopStatRanking     []*TopStatRanking  `json:"topStatRanking,omitempty"`
	TopStats           []*TopStat         `json:"topStats,omitempty"`
	StatCardSections   []*StatCardSection `json:"statCardSections,omitempty"`
}

type Goals struct {
	Total                 *int     `json:"total,omitempty"`
	OpenPlayGoals         *int     `json:"openPlayGoals,omitempty"`
	PenaltyGoals          *int     `json:"penaltyGoals,omitempty"`
	FreeKickGoals         *int     `json:"freeKickGoals,omitempty"`
	AccruedOwnGoals       *int     `json:"accruedOwnGoals,omitempty"`
	FastBreakGoals        *int     `json:"fastBreakGoals,omitempty"`
	SetPiece              *int     `json:"setPiece,omitempty"`
	ExpectedGoals         *float64 `json:"expectedGoals,omitempty"`
	ExpectedGoalsConceded *float64 `json:"expectedGoalsConceded,omitempty"`
}

type RelegationInfo struct {
	Promotion       *int    `json:"promotion,omitempty"`
	Relegation      *int    `json:"relegation,omitempty"`
	ChampionsLeague *string `json:"championsLeague,omitempty"`
	UefaCup         *string `json:"uefaCup,omitempty"`
	StageID         *int    `json:"stageId,omitempty"`
	Rules           []*Rule `json:"rules,omitempty"`
}

type Rule struct {
	ID             *int    `json:"id,omitempty"`
	Description    *string `json:"description,omitempty"`
	Color          *string `json:"color,omitempty"`
	LeagacyID      *string `json:"leagacyId,omitempty"`
	Value          *string `json:"value,omitempty"`
	Name           *string `json:"name,omitempty"`
	TranslationKey *string `json:"translationKey,omitempty"`
	SortOrder      *int    `json:"sortOrder,omitempty"`
}

type StatCardSection struct {
	Title               *string              `json:"title,omitempty"`
	LocalizedTitleID    *string              `json:"localizedTitleId,omitempty"`
	LocalizedCategoryID *LocalizedCategoryID `json:"localizedCategoryId,omitempty"`
	Stats               []string             `json:"stats,omitempty"`
}

type LocalizedCategoryID string

const (
	TopStats    LocalizedCategoryID = "top_stats"
	Attack      LocalizedCategoryID = "attack"
	Defense     LocalizedCategoryID = "defense"
	Discipline  LocalizedCategoryID = "discipline"
	Goalkeeping LocalizedCategoryID = "goalkeeping"
)

type Table struct {
	Lines               []*Line       `json:"lines,omitempty"`
	SubTables           []interface{} `json:"subTables,omitempty"`
	WinningPositions    *int          `json:"winningPositions,omitempty"`
	TopPositions        *int          `json:"topPositions,omitempty"`
	QualifyingPositions *int          `json:"qualifyingPositions,omitempty"`
	RelegationPositions *int          `json:"relegationPositions,omitempty"`
	LeagueName          *string       `json:"leagueName,omitempty"`
	SeasonID            *int          `json:"seasonId,omitempty"`
	LeagueID            *int          `json:"leagueId,omitempty"`
	IsComposite         *bool         `json:"isComposite,omitempty"`
	CountryCode         *string       `json:"countryCode,omitempty"`
	HasXgTable          *bool         `json:"hasXgTable,omitempty"`
}

type CountryCode string

const (
	ENG CountryCode = "ENG"
	MLI CountryCode = "MLI"
	SCO CountryCode = "SCO"
	UKR CountryCode = "UKR"
	NED CountryCode = "NED"
	POR CountryCode = "POR"
	SEN CountryCode = "SEN"
	BEL CountryCode = "BEL"
)

type Line struct {
	TeamId          *int    `json:"teamId,omitempty"`
	TeamName        *string `json:"teamName,omitempty"`
	Position        *int    `json:"position,omitempty"`
	Won             *int    `json:"won,omitempty"`
	Drawn           *int    `json:"drawn,omitempty"`
	Lost            *int    `json:"lost,omitempty"`
	Scored          *int    `json:"scored,omitempty"`
	Conceeded       *int    `json:"conceeded,omitempty"`
	Points          *int    `json:"points,omitempty"`
	Conference      *string `json:"conference,omitempty"`
	PointsHome      *int    `json:"pointsHome,omitempty"`
	WonHome         *int    `json:"wonHome,omitempty"`
	DrawnHome       *int    `json:"drawnHome,omitempty"`
	LostHome        *int    `json:"lostHome,omitempty"`
	ScoredHome      *int    `json:"scoredHome,omitempty"`
	ConceededHome   *int    `json:"conceededHome,omitempty"`
	PointsDeduction *int    `json:"pointsDeduction,omitempty"`
}

type TopStatRanking struct {
	StatName            *string             `json:"statName,omitempty"`
	Title               *string             `json:"title,omitempty"`
	LocalizedTitleId    *string             `json:"localizedTitleId,omitempty"`
	StatList            []StatList          `json:"statList,omitempty"`
	StatFormat          StatFormat          `json:"statFormat,omitempty"`
	StatDecimals        *int                `json:"statDecimals,omitempty"`
	SubstatDecimals     *int                `json:"substatDecimals,omitempty"`
	Order               *int                `json:"order,omitempty"`
	Category            Category            `json:"category,omitempty"`
	LocalizedCategoryId LocalizedCategoryID `json:"localizedCategoryId,omitempty"`
	Subtitle            *string             `json:"subtitle,omitempty"`
	LocalizedSubtitleId *string             `json:"localizedSubtitleId,omitempty"`
	SubstatFormat       StatFormat          `json:"substatFormat,omitempty"`
}

type Category string

const (
	TopStatCategory     Category = "Top Stat"
	Attacking           Category = "Attacking"
	Defending           Category = "Defending"
	DisciplineCategory  Category = "Discipline"
	GoalkeepingCategory Category = "Goalkeeping"
)

type StatFormat string

const (
	Fraction StatFormat = "fraction"
	Percent  StatFormat = "percent"
	Number   StatFormat = "number"
)

type StatList struct {
	ParticipantName        *string     `json:"participantName,omitempty"`
	ParticiantId           *int        `json:"particiantId,omitempty"`
	TeamId                 *int        `json:"teamId,omitempty"`
	TeamColor              *string     `json:"teamColor,omitempty"`
	StatValue              *float64    `json:"statValue,omitempty"`
	SubStatValue           *float64    `json:"subStatValue,omitempty"`
	MinutesPlayed          *int        `json:"minutesPlayed,omitempty"`
	MatchesPlayed          *int        `json:"matchesPlayed,omitempty"`
	StatValueCount         *int        `json:"statValueCount,omitempty"`
	Rank                   *int        `json:"rank,omitempty"`
	ParticipantCountryCode CountryCode `json:"participantCountryCode,omitempty"`
	TeamName               *string     `json:"teamName,omitempty"`
}

type TopStat struct {
	StatName            *string             `json:"statName,omitempty"`
	Title               *string             `json:"title,omitempty"`
	Subtitle            *string             `json:"subtitle,omitempty"`
	LocalizedTitleId    *string             `json:"localizedTitleId,omitempty"`
	LocalizedSubtitleId *string             `json:"localizedSubtitleId,omitempty"`
	StatList            []StatList          `json:"statList,omitempty"`
	StatFormat          StatFormat          `json:"statFormat,omitempty"`
	SubstatFormat       StatFormat          `json:"substatFormat,omitempty"`
	StatDecimals        *int                `json:"statDecimals,omitempty"`
	SubstatDecimals     *int                `json:"substatDecimals,omitempty"`
	StatLocation        *string             `json:"statLocation,omitempty"`
	Order               *int                `json:"order,omitempty"`
	Category            Category            `json:"category,omitempty"`
	LocalizedCategoryId LocalizedCategoryID `json:"localizedCategoryId,omitempty"`
}
