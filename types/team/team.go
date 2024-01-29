package team

import (
	"time"
)

type Team struct {
	Tabs                []string         `json:"tabs,omitempty"`
	AllAvailableSeasons []interface{}    `json:"allAvailableSeasons,omitempty"`
	Details             *Details         `json:"details,omitempty"`
	Seostr              *string          `json:"seostr,omitempty"`
	QAData              []*QADatum       `json:"QAData,omitempty"`
	Table               []*OverviewTable `json:"table,omitempty"`
	Transfers           *Transfers       `json:"transfers,omitempty"`
	Overview            *Overview        `json:"overview,omitempty"`
	Stats               *TeamStats       `json:"stats,omitempty"`
	Fixtures            *TeamFixtures    `json:"fixtures,omitempty"`
	Squad               [][]*SquadClass  `json:"squad,omitempty"`
	History             *History         `json:"history,omitempty"`
}

type QADatum struct {
	Question *string `json:"question,omitempty"`
	Answer   *string `json:"answer,omitempty"`
}

type Details struct {
	ID               *int              `json:"id,omitempty"`
	Type             *string           `json:"type,omitempty"`
	Name             *string           `json:"name,omitempty"`
	LatestSeason     *string           `json:"latestSeason,omitempty"`
	ShortName        *string           `json:"shortName,omitempty"`
	Country          *Country          `json:"country,omitempty"`
	FaqJSONLD        *FAQJSONLD        `json:"faqJSONLD,omitempty"`
	SportsTeamJSONLD *SportsTeamJSONLD `json:"sportsTeamJSONLD,omitempty"`
	BreadcrumbJSONLD *BreadcrumbJSONLD `json:"breadcrumbJSONLD,omitempty"`
	CanSyncCalendar  *bool             `json:"canSyncCalendar,omitempty"`
}

type BreadcrumbJSONLD struct {
	Context         *string            `json:"@context,omitempty"`
	Type            *string            `json:"@type,omitempty"`
	ItemListElement []*ItemListElement `json:"itemListElement,omitempty"`
}

type ItemListElement struct {
	Type     *string `json:"@type,omitempty"`
	Position *int    `json:"position,omitempty"`
	Name     *string `json:"name,omitempty"`
	Item     *string `json:"item,omitempty"`
}

type Country string

const (
	Eng Country = "ENG"
	Int Country = "INT"
)

type FAQJSONLD struct {
	Context    *string       `json:"@context,omitempty"`
	Type       *string       `json:"@type,omitempty"`
	MainEntity []*MainEntity `json:"mainEntity,omitempty"`
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

type SportsTeamJSONLD struct {
	Context  *string   `json:"@context,omitempty"`
	Type     *string   `json:"@type,omitempty"`
	Name     *string   `json:"name,omitempty"`
	Sport    *string   `json:"sport,omitempty"`
	Gender   *string   `json:"gender,omitempty"`
	Logo     *string   `json:"logo,omitempty"`
	Url      *string   `json:"url,omitempty"`
	Coach    *Coach    `json:"coach,omitempty"`
	Athlete  []*Coach  `json:"athlete,omitempty"`
	Location *Location `json:"location,omitempty"`
	MemberOf *MemberOf `json:"memberOf,omitempty"`
}

type Coach struct {
	Context     *string      `json:"@context,omitempty"`
	Type        *CoachType   `json:"@type,omitempty"`
	Name        *string      `json:"name,omitempty"`
	URL         *string      `json:"url,omitempty"`
	Nationality *Nationality `json:"nationality,omitempty"`
	Affiliation *interface{} `json:"affiliation,omitempty"`
	Height      *interface{} `json:"height,omitempty"`
	Weight      *interface{} `json:"weight,omitempty"`
}

type CoachType string

const (
	Person CoachType = "Person"
)

type Nationality struct {
	Type *string `json:"@type,omitempty"`
	Name *string `json:"name,omitempty"`
}

type Location struct {
	Type    *string  `json:"@type,omitempty"`
	Name    *string  `json:"name,omitempty"`
	Address *Address `json:"address,omitempty"`
	Geo     *Geo     `json:"geo,omitempty"`
}

type Address struct {
	Type            *string `json:"@type,omitempty"`
	AddressCountry  *string `json:"addressCountry,omitempty"`
	AddressLocality *string `json:"addressLocality,omitempty"`
}

type Geo struct {
	Type      *string `json:"@type,omitempty"`
	Latitude  *string `json:"latitude,omitempty"`
	Longitude *string `json:"longitude,omitempty"`
}

type MemberOf struct {
	Type *string `json:"@type,omitempty"`
	Name *string `json:"name,omitempty"`
	URL  *string `json:"url,omitempty"`
}

type TeamFixtures struct {
	AllFixtures         *AllFixtures      `json:"allFixtures,omitempty"`
	PreviousFixturesURL *string           `json:"previousFixturesUrl,omitempty"`
	HasOngoingMatch     *bool             `json:"hasOngoingMatch,omitempty"`
	Fixtures            *FixturesFixtures `json:"fixtures,omitempty"`
}

type AllFixtures struct {
	Fixtures  []*OverviewFixture `json:"fixtures,omitempty"`
	NextMatch *NextMatch         `json:"nextMatch,omitempty"`
	LastMatch *LastMatch         `json:"lastMatch,omitempty"`
}

type OverviewFixture struct {
	ID                *int                   `json:"id,omitempty"`
	PageURL           *string                `json:"pageUrl,omitempty"`
	Opponent          *OpponentClass         `json:"opponent,omitempty"`
	Home              *OpponentClass         `json:"home,omitempty"`
	Away              *OpponentClass         `json:"away,omitempty"`
	DisplayTournament *bool                  `json:"displayTournament,omitempty"`
	Result            *int                   `json:"result,omitempty"`
	NotStarted        *bool                  `json:"notStarted,omitempty"`
	Tournament        *Tournament            `json:"tournament,omitempty"`
	Status            *OverviewFixtureStatus `json:"status,omitempty"`
	MonthYearKey      *string                `json:"monthYearKey,omitempty"`
	IsPastMatch       *bool                  `json:"isPastMatch,omitempty"`
	LastPlayedMatch   *bool                  `json:"lastPlayedMatch,omitempty"`
}

type OpponentClass struct {
	ID    *int    `json:"id,omitempty"`
	Name  *string `json:"name,omitempty"`
	Score *int    `json:"score,omitempty"`
}

type OverviewFixtureStatus struct {
	UTCTime   *time.Time  `json:"utcTime,omitempty"`
	Finished  *bool       `json:"finished,omitempty"`
	Started   *bool       `json:"started,omitempty"`
	Cancelled *bool       `json:"cancelled,omitempty"`
	ScoreStr  *string     `json:"scoreStr,omitempty"`
	Reason    *Reason     `json:"reason,omitempty"`
	Timezone  interface{} `json:"timezone,omitempty"`
}

type Reason struct {
	Short    *Short    `json:"short,omitempty"`
	ShortKey *ShortKey `json:"shortKey,omitempty"`
	Long     *Long     `json:"long,omitempty"`
	LongKey  *LongKey  `json:"longKey,omitempty"`
}

type Long string

const (
	FullTime Long = "Full-Time"
)

type LongKey string

const (
	Finished LongKey = "finished"
)

type Short string

const (
	Ft Short = "FT"
)

type ShortKey string

const (
	FulltimeShort ShortKey = "fulltime_short"
)

type Tournament struct {
	Name     *string `json:"name,omitempty"`
	LeagueID *int    `json:"leagueId,omitempty"`
}

type LastMatch struct {
	ID                *int                   `json:"id,omitempty"`
	PageURL           *string                `json:"pageUrl,omitempty"`
	Opponent          *OpponentClass         `json:"opponent,omitempty"`
	Home              *OpponentClass         `json:"home,omitempty"`
	Away              *OpponentClass         `json:"away,omitempty"`
	DisplayTournament *bool                  `json:"displayTournament,omitempty"`
	Result            *int                   `json:"result,omitempty"`
	NotStarted        *bool                  `json:"notStarted,omitempty"`
	Tournament        *Tournament            `json:"tournament,omitempty"`
	Status            *OverviewFixtureStatus `json:"status,omitempty"`
}

type NextMatch struct {
	ID                *int             `json:"id,omitempty"`
	PageURL           *string          `json:"pageUrl,omitempty"`
	Opponent          *OpponentClass   `json:"opponent,omitempty"`
	Home              *OpponentClass   `json:"home,omitempty"`
	Away              *OpponentClass   `json:"away,omitempty"`
	DisplayTournament *bool            `json:"displayTournament,omitempty"`
	NotStarted        *bool            `json:"notStarted,omitempty"`
	Tournament        *Tournament      `json:"tournament,omitempty"`
	Status            *NextMatchStatus `json:"status,omitempty"`
}

type NextMatchStatus struct {
	UTCTime   *time.Time  `json:"utcTime,omitempty"`
	Timezone  interface{} `json:"timezone,omitempty"`
	Started   *bool       `json:"started,omitempty"`
	Cancelled *bool       `json:"cancelled,omitempty"`
	Finished  *bool       `json:"finished,omitempty"`
}

type FixturesFixtures struct {
	July2023      []*OverviewFixture `json:"July 2023,omitempty"`
	August2023    []*OverviewFixture `json:"August 2023,omitempty"`
	September2023 []*OverviewFixture `json:"September 2023,omitempty"`
	October2023   []*OverviewFixture `json:"October 2023,omitempty"`
	November2023  []*OverviewFixture `json:"November 2023,omitempty"`
	December2023  []*OverviewFixture `json:"December 2023,omitempty"`
	January2024   []*The2024         `json:"January 2024,omitempty"`
	February2024  []*The2024         `json:"February 2024,omitempty"`
	March2024     []*The2024         `json:"March 2024,omitempty"`
	April2024     []*The2024         `json:"April 2024,omitempty"`
	May2024       []*The2024         `json:"May 2024,omitempty"`
}

type The2024 struct {
	ID                *int             `json:"id,omitempty"`
	PageURL           *string          `json:"pageUrl,omitempty"`
	MonthYearKey      *string          `json:"monthYearKey,omitempty"`
	Opponent          *OpponentClass   `json:"opponent,omitempty"`
	Home              *OpponentClass   `json:"home,omitempty"`
	Away              *OpponentClass   `json:"away,omitempty"`
	DisplayTournament *bool            `json:"displayTournament,omitempty"`
	IsPastMatch       *bool            `json:"isPastMatch,omitempty"`
	LastPlayedMatch   *bool            `json:"lastPlayedMatch,omitempty"`
	NotStarted        *bool            `json:"notStarted,omitempty"`
	Tournament        *Tournament      `json:"tournament,omitempty"`
	Status            *NextMatchStatus `json:"status,omitempty"`
}

type History struct {
	TrophyList          []*TrophyList        `json:"trophyList,omitempty"`
	HistoricalTableData *HistoricalTableData `json:"historicalTableData,omitempty"`
	TeamColor           *string              `json:"teamColor,omitempty"`
	TeamColorMap        *TeamColorMap        `json:"teamColorMap,omitempty"`
	Tables              *Tables              `json:"tables,omitempty"`
}

type HistoricalTableData struct {
	Divisions []*Division `json:"divisions,omitempty"`
	Ranks     []*Rank     `json:"ranks,omitempty"`
	IsValid   *bool       `json:"isValid,omitempty"`
}

type Division struct {
	DivisionRank *int    `json:"divisionRank,omitempty"`
	Name         *string `json:"name,omitempty"`
	TemplateID   *int    `json:"templateId,omitempty"`
}

type Rank struct {
	StageID        *int       `json:"stageId,omitempty"`
	TournamentName *string    `json:"tournamentName,omitempty"`
	TournamentID   *int       `json:"tournamentId,omitempty"`
	TemplateID     *int       `json:"templateId,omitempty"`
	SeasonName     *string    `json:"seasonName,omitempty"`
	Position       *int       `json:"position,omitempty"`
	NumberOfTeams  *int       `json:"numberOfTeams,omitempty"`
	Stats          *RankStats `json:"stats,omitempty"`
	TableLink      *string    `json:"tableLink,omitempty"`
	IsConsecutive  *bool      `json:"isConsecutive,omitempty"`
}

type RankStats struct {
	Points *int `json:"points,omitempty"`
	Wins   *int `json:"wins,omitempty"`
	Draws  *int `json:"draws,omitempty"`
	Loss   *int `json:"loss,omitempty"`
}

type Tables struct {
	Current  []*Current `json:"current,omitempty"`
	Historic []*Current `json:"historic,omitempty"`
}

type Current struct {
	Link []*Link `json:"link,omitempty"`
}

type Link struct {
	// _            *string  `json:"_,omitempty"`
	Name         []string `json:"name,omitempty"`
	Ccode        []string `json:"ccode,omitempty"`
	Season       []string `json:"season,omitempty"`
	StageID      []string `json:"stage_id,omitempty"`
	TournamentID []string `json:"tournament_id,omitempty"`
	TemplateID   []string `json:"template_id,omitempty"`
	LeagueID     []string `json:"league_id,omitempty"`
}

type TeamColorMap struct {
	Color              *string `json:"color,omitempty"`
	ColorAlternate     *string `json:"colorAlternate,omitempty"`
	ColorAway          *string `json:"colorAway,omitempty"`
	ColorAwayAlternate *string `json:"colorAwayAlternate,omitempty"`
}

type TrophyList struct {
	Name                 []string `json:"name,omitempty"`
	TournamentTemplateId []string `json:"tournamentTemplateId,omitempty"`
	Area                 []string `json:"area,omitempty"`
	Ccode                []string `json:"ccode,omitempty"`
	Won                  []string `json:"won,omitempty"`
	Runnerup             []string `json:"runnerup,omitempty"`
	SeasonWon            []string `json:"season_won,omitempty"`
	SeasonRunnerup       []string `json:"season_runnerup,omitempty"`
}

type Overview struct {
	Season              *string            `json:"season,omitempty"`
	SelectedSeason      *string            `json:"selectedSeason,omitempty"`
	Table               []*OverviewTable   `json:"table,omitempty"`
	TopPlayers          *TopPlayers        `json:"topPlayers,omitempty"`
	Venue               *Venue             `json:"venue,omitempty"`
	OverviewFixtures    []*OverviewFixture `json:"overviewFixtures,omitempty"`
	NextMatch           *NextMatch         `json:"nextMatch,omitempty"`
	LastMatch           *LastMatch         `json:"lastMatch,omitempty"`
	TeamForm            []*TeamForm        `json:"teamForm,omitempty"`
	HasOngoingMatch     *bool              `json:"hasOngoingMatch,omitempty"`
	PreviousFixturesUrl *string            `json:"previousFixturesUrl,omitempty"`
	TeamColors          *TeamColors        `json:"teamColors,omitempty"`
}

type OverviewTable struct {
	Data         *TableData               `json:"data,omitempty"`
	TeamForm     map[string][]*TeamForm   `json:"teamForm,omitempty"`
	NextOpponent map[string][]interface{} `json:"nextOpponent,omitempty"`
	TableHeader  *TableHeader             `json:"tableHeader,omitempty"`
}

type TableData struct {
	Ccode            *string        `json:"ccode,omitempty"`
	LeagueId         *int           `json:"leagueId,omitempty"`
	PageUrl          *string        `json:"pageUrl,omitempty"`
	LeagueName       *string        `json:"leagueName,omitempty"`
	Legend           []*Legend      `json:"legend,omitempty"`
	Ongoing          []interface{}  `json:"ongoing,omitempty"`
	Table            *PurpleTable   `json:"table,omitempty"`
	TableFilterTypes []string       `json:"tableFilterTypes,omitempty"`
	Composite        *bool          `json:"composite,omitempty"`
	Tables           []*FluffyTable `json:"tables,omitempty"`
}

type Legend struct {
	Title   *string `json:"title,omitempty"`
	TKey    *string `json:"tKey,omitempty"`
	Color   *string `json:"color,omitempty"`
	Indices []int   `json:"indices,omitempty"`
}

type PurpleTable struct {
	All  []*All `json:"all,omitempty"`
	Home []*All `json:"home,omitempty"`
	Away []*All `json:"away,omitempty"`
	Form []*All `json:"form,omitempty"`
}

type All struct {
	Name            *string     `json:"name,omitempty"`
	ShortName       *string     `json:"shortName,omitempty"`
	Id              *int        `json:"id,omitempty"`
	PageUrl         *string     `json:"pageUrl,omitempty"`
	FeaturedInMatch *bool       `json:"featuredInMatch,omitempty"`
	Deduction       *int        `json:"deduction,omitempty"`
	Ongoing         interface{} `json:"ongoing,omitempty"`
	Played          *int        `json:"played,omitempty"`
	Wins            *int        `json:"wins,omitempty"`
	Draws           *int        `json:"draws,omitempty"`
	Losses          *int        `json:"losses,omitempty"`
	ScoresStr       *string     `json:"scoresStr,omitempty"`
	GoalConDiff     *int        `json:"goalConDiff,omitempty"`
	Pts             *int        `json:"pts,omitempty"`
	Idx             *int        `json:"idx,omitempty"`
	QualColor       *string     `json:"qualColor,omitempty"`
}

type FluffyTable struct {
	Ccode      *string       `json:"ccode,omitempty"`
	LeagueID   *int          `json:"leagueId,omitempty"`
	PageURL    *string       `json:"pageUrl,omitempty"`
	LeagueName *string       `json:"leagueName,omitempty"`
	Legend     []*Legend     `json:"legend,omitempty"`
	Table      *TableTable   `json:"table,omitempty"`
	Ongoing    []interface{} `json:"ongoing,omitempty"`
}

type TableTable struct {
	All  []*All `json:"all,omitempty"`
	Home []*All `json:"home,omitempty"`
	Away []*All `json:"away,omitempty"`
}

type NextOpponentClass struct {
	ID        *string `json:"id,omitempty"`
	Name      *string `json:"name,omitempty"`
	ShortName *string `json:"shortName,omitempty"`
}

type TableHeader struct {
	StaticTableHeaders  []string `json:"staticTableHeaders,omitempty"`
	DynamicTableHeaders []string `json:"dynamicTableHeaders,omitempty"`
}

type TeamForm struct {
	Result       *int          `json:"result,omitempty"`
	ResultString *ResultString `json:"resultString,omitempty"`
	ImageURL     *string       `json:"imageUrl,omitempty"`
	LinkToMatch  *string       `json:"linkToMatch,omitempty"`
	Date         *DateClass    `json:"date,omitempty"`
	TeamPageURL  *string       `json:"teamPageUrl,omitempty"`
	TooltipText  *TooltipText  `json:"tooltipText,omitempty"`
	Score        *string       `json:"score,omitempty"`
	Home         *TeamFormAway `json:"home,omitempty"`
	Away         *TeamFormAway `json:"away,omitempty"`
}

type TeamFormAway struct {
	IsOurTeam *bool `json:"isOurTeam,omitempty"`
}

type DateClass struct {
	UTCTime *time.Time `json:"utcTime,omitempty"`
}

type ResultString string

const (
	D ResultString = "D"
	L ResultString = "L"
	W ResultString = "W"
)

type TooltipText struct {
	UTCTime    *time.Time `json:"utcTime,omitempty"`
	HomeTeam   *string    `json:"homeTeam,omitempty"`
	HomeTeamID *int       `json:"homeTeamId,omitempty"`
	HomeScore  *int       `json:"homeScore,omitempty"`
	AwayTeam   *string    `json:"awayTeam,omitempty"`
	AwayTeamID *int       `json:"awayTeamId,omitempty"`
	AwayScore  *int       `json:"awayScore,omitempty"`
}

type TeamColors struct {
	DarkMode      *string `json:"darkMode,omitempty"`
	LightMode     *string `json:"lightMode,omitempty"`
	FontDarkMode  *string `json:"fontDarkMode,omitempty"`
	FontLightMode *string `json:"fontLightMode,omitempty"`
}

type TopPlayers struct {
	ByRating  *By     `json:"byRating,omitempty"`
	ByGoals   *By     `json:"byGoals,omitempty"`
	ByAssists *By     `json:"byAssists,omitempty"`
	SeeAllURL *string `json:"seeAllUrl,omitempty"`
}

type By struct {
	Players    []*Participant `json:"players,omitempty"`
	SeeAllLink *string        `json:"seeAllLink,omitempty"`
}

type Participant struct {
	ID                 *int        `json:"id,omitempty"`
	Name               *string     `json:"name,omitempty"`
	Goals              *string     `json:"goals,omitempty"`
	Assists            *string     `json:"assists,omitempty"`
	Rating             *string     `json:"rating,omitempty"`
	PositionID         *int        `json:"positionId,omitempty"`
	Ccode              *string     `json:"ccode,omitempty"`
	Cname              *string     `json:"cname,omitempty"`
	TeamID             *int        `json:"teamId,omitempty"`
	TeamName           *string     `json:"teamName,omitempty"`
	ShowRole           *bool       `json:"showRole,omitempty"`
	ShowCountryFlag    *bool       `json:"showCountryFlag,omitempty"`
	ShowTeamFlag       bool        `json:"showTeamFlag"`
	ExcludeFromRanking bool        `json:"excludeFromRanking"`
	Rcards             string      `json:"rcards"`
	Ycards             string      `json:"ycards"`
	TeamColors         *TeamColors `json:"teamColors,omitempty"`
	Injured            bool        `json:"injured"`
	Rank               *int        `json:"rank,omitempty"`
	Value              *float32    `json:"value,omitempty"`
}

type Venue struct {
	Widget    *Widget         `json:"widget,omitempty"`
	StatPairs [][]interface{} `json:"statPairs,omitempty"`
}

type Widget struct {
	Name     *string  `json:"name,omitempty"`
	Location []string `json:"location,omitempty"`
	City     *string  `json:"city,omitempty"`
}

type SquadClass struct {
	ID        *int    `json:"id,omitempty"`
	Name      *string `json:"name,omitempty"`
	Ccode     *string `json:"ccode,omitempty"`
	Cname     *string `json:"cname,omitempty"`
	Role      Role    `json:"role,omitempty"`
	IsInjured bool    `json:"isInjured"`
}

type Role string

const (
	Attackers   Role = "attackers"
	Defenders   Role = "defenders"
	Goalkeepers Role = "goalkeepers"
	Midfielders Role = "midfielders"
)

type TeamStats struct {
	TeamID            *string             `json:"teamId,omitempty"`
	PrimaryLeagueID   *int                `json:"primaryLeagueId,omitempty"`
	PrimarySeasonID   *string             `json:"primarySeasonId,omitempty"`
	Players           []*Player           `json:"players,omitempty"`
	Teams             []*TeamElement      `json:"teams,omitempty"`
	TournamentID      *string             `json:"tournamentId,omitempty"`
	TournamentSeasons []*TournamentSeason `json:"tournamentSeasons,omitempty"`
}

type Player struct {
	Header           *string        `json:"header,omitempty"`
	Participant      *Participant   `json:"participant,omitempty"`
	FetchAllURL      *string        `json:"fetchAllUrl,omitempty"`
	TopThree         []*Participant `json:"topThree,omitempty"`
	Order            *int           `json:"order,omitempty"`
	Name             *string        `json:"name,omitempty"`
	LocalizedTitleID *string        `json:"localizedTitleId,omitempty"`
}

type TeamElement struct {
	Header           *string        `json:"header,omitempty"`
	LocalizedTitleID *string        `json:"localizedTitleId,omitempty"`
	Participant      *Participant   `json:"participant,omitempty"`
	FetchAllURL      *string        `json:"fetchAllUrl,omitempty"`
	TopThree         []*Participant `json:"topThree,omitempty"`
	Order            *int           `json:"order,omitempty"`
	Stat             *string        `json:"stat,omitempty"`
}

type TournamentSeason struct {
	Name           *string `json:"name,omitempty"`
	Season         *string `json:"season,omitempty"`
	LeagueName     *string `json:"leagueName,omitempty"`
	TournamentID   *string `json:"tournamentId,omitempty"`
	ParentLeagueID *string `json:"parentLeagueId,omitempty"`
}

type Transfers struct {
	Type *string        `json:"type,omitempty"`
	Data *TransfersData `json:"data,omitempty"`
}

type TransfersData struct {
	PlayersIn         []*Player            `json:"Players in,omitempty"`
	PlayersOut        []*Player            `json:"Players out,omitempty"`
	ContractExtension []*ContractExtension `json:"Contract extension,omitempty"`
}

type ContractExtension struct {
	Name              *string       `json:"name,omitempty"`
	PlayerID          *int          `json:"playerId,omitempty"`
	Position          *Position     `json:"position,omitempty"`
	TransferDate      *time.Time    `json:"transferDate,omitempty"`
	TransferText      []interface{} `json:"transferText,omitempty"`
	FromClub          *string       `json:"fromClub,omitempty"`
	FromClubID        *int          `json:"fromClubId,omitempty"`
	ToClub            *string       `json:"toClub,omitempty"`
	ToClubID          *int          `json:"toClubId,omitempty"`
	Fee               interface{}   `json:"fee,omitempty"`
	TransferType      *TransferType `json:"transferType,omitempty"`
	ContractExtension bool          `json:"contractExtension"`
	OnLoan            bool          `json:"onLoan"`
	FromDate          *time.Time    `json:"fromDate,omitempty"`
	ToDate            *time.Time    `json:"toDate,omitempty"`
	MarketValue       *string       `json:"marketValue,omitempty"`
}

type Position struct {
	Label *string `json:"label,omitempty"`
	Key   *string `json:"key,omitempty"`
}

type TransferType struct {
	Text            *Text            `json:"text,omitempty"`
	LocalizationKey *LocalizationKey `json:"localizationKey,omitempty"`
}

type LocalizationKey string

const (
	ContractLocalizationKey LocalizationKey = "contract"
	OnLoanLocalizationKey   LocalizationKey = "on_loan"
)

type Text string

const (
	ContractText Text = "contract"
	OnLoanText   Text = "on loan"
)

type Players struct {
	Name              *string       `json:"name,omitempty"`
	PlayerID          *int          `json:"playerId,omitempty"`
	Position          *Position     `json:"position,omitempty"`
	TransferDate      *time.Time    `json:"transferDate,omitempty"`
	TransferText      []interface{} `json:"transferText,omitempty"`
	FromClub          *string       `json:"fromClub,omitempty"`
	FromClubID        *int          `json:"fromClubId,omitempty"`
	ToClub            *string       `json:"toClub,omitempty"`
	ToClubID          *int          `json:"toClubId,omitempty"`
	Fee               interface{}   `json:"fee,omitempty"`
	TransferType      *TransferType `json:"transferType,omitempty"`
	ContractExtension bool          `json:"contractExtension"`
	OnLoan            bool          `json:"onLoan"`
	FromDate          *time.Time    `json:"fromDate,omitempty"`
	ToDate            *time.Time    `json:"toDate,omitempty"`
	MarketValue       *string       `json:"marketValue,omitempty"`
}

type Fee struct {
	FeeText          *FeeText          `json:"feeText,omitempty"`
	LocalizedFeeText *LocalizedFeeText `json:"localizedFeeText,omitempty"`
	Value            *string           `json:"value,omitempty"`
}

type FeeText string

const (
	FeeFeeText          FeeText = "fee"
	FreeTransferFeeText FeeText = "free transfer"
	OnLoanFeeText       FeeText = "on loan"
)

type LocalizedFeeText string

const (
	OnLoanLocalizedFeeText                   LocalizedFeeText = "on_loan"
	TransferFeeLocalizedFeeText              LocalizedFeeText = "transfer_fee"
	TransferTypeFreeTransferLocalizedFeeText LocalizedFeeText = "transfer_type_free_transfer"
)
