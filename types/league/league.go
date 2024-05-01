package league

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/smithy-go/ptr"
)

type League struct {
	Tabs                []string       `json:"tabs,omitempty"`
	AllAvailableSeasons []string       `json:"allAvailableSeasons,omitempty"`
	Details             *Details       `json:"details,omitempty"`
	Seostr              *string        `json:"seostr,omitempty"`
	QAData              []interface{}  `json:"QAData,omitempty"`
	Table               []TableElement `json:"table,omitempty"`
	Transfers           *Transfers     `json:"transfers,omitempty"`
	Overview            *Overview      `json:"overview,omitempty"`
	Stats               *Stats         `json:"stats,omitempty"`
	Matches             *Matches       `json:"matches,omitempty"`
	Playoff             interface{}    `json:"playoff,omitempty"`
}

type Details struct {
	ID               *int              `json:"id,omitempty"`
	Type             *string           `json:"type,omitempty"`
	Name             *string           `json:"name,omitempty"`
	SelectedSeason   *string           `json:"selectedSeason,omitempty"`
	LatestSeason     *string           `json:"latestSeason,omitempty"`
	ShortName        *string           `json:"shortName,omitempty"`
	Country          *string           `json:"country,omitempty"`
	FaqJSONLD        interface{}       `json:"faqJSONLD,omitempty"`
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

type Matches struct {
	FirstUnplayedMatch *FirstUnplayedMatch `json:"firstUnplayedMatch,omitempty"`
	AllMatches         []AllMatch          `json:"allMatches,omitempty"`
	HasOngoingMatch    *bool               `json:"hasOngoingMatch,omitempty"`
}

type AllMatch struct {
	Round     *string            `json:"round,omitempty"`
	RoundName *string            `json:"roundName,omitempty"`
	PageUrl   *string            `json:"pageUrl,omitempty"`
	ID        *string            `json:"id,omitempty"`
	Home      *NextOpponentClass `json:"home,omitempty"`
	Away      *NextOpponentClass `json:"away,omitempty"`
	Status    *Status            `json:"status,omitempty"`
}

func (a *AllMatch) UnmarshalJSON(b []byte) error {
	var s map[string]interface{}
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	roundName := s["roundName"]
	s["roundName"] = nil

	d, err := json.Marshal(s)
	if err != nil {
		return err
	}

	type Alias AllMatch
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(a),
	}
	if err := json.Unmarshal(d, &aux); err != nil {
		return err
	}

	*a = AllMatch(*aux.Alias)

	switch roundName := roundName.(type) {
	case float64:
		a.RoundName = ptr.String(fmt.Sprintf("%f", roundName))
	case string:
		a.RoundName = ptr.String(roundName)
	}

	return nil
}

type NextOpponentClass struct {
	Name      *string `json:"name,omitempty"`
	ShortName *string `json:"shortName,omitempty"`
	ID        *string `json:"id,omitempty"`
}

type Status struct {
	UTCTime   *time.Time `json:"utcTime,omitempty"`
	Finished  *bool      `json:"finished,omitempty"`
	Started   *bool      `json:"started,omitempty"`
	Cancelled *bool      `json:"cancelled,omitempty"`
	ScoreStr  *string    `json:"scoreStr,omitempty"`
	Reason    *Reason    `json:"reason,omitempty"`
}

type Reason struct {
	Short    *Short    `json:"short,omitempty"`
	ShortKey *ShortKey `json:"shortKey,omitempty"`
	Long     *Long     `json:"long,omitempty"`
	LongKey  *LongKey  `json:"longKey,omitempty"`
}

type Short string

const (
	ShortAb Short = "Ab"
	ShortFt Short = "FT"
	ShortPp Short = "PP"
)

type ShortKey string

const (
	ShortKeyAbortedShort   ShortKey = "aborted_short"
	ShortKeyFulltimeShort  ShortKey = "fulltime_short"
	ShortKeyPostponedShort ShortKey = "postponed_short"
)

type Long string

const (
	LongAbandoned Long = "Abandoned"
	LongFullTime  Long = "Full-Time"
	LongPostponed Long = "Postponed"
)

type LongKey string

const (
	LongKeyAborted   LongKey = "aborted"
	LongKeyFinished  LongKey = "finished"
	LongKeyPostponed LongKey = "postponed"
)

type FirstUnplayedMatch struct {
	FirstRoundWithUnplayedMatch *string `json:"firstRoundWithUnplayedMatch,omitempty"`
	FirstUnplayedMatchIndex     *int    `json:"firstUnplayedMatchIndex,omitempty"`
	FirstUnplayedMatchID        *string `json:"firstUnplayedMatchId,omitempty"`
}

type Overview struct {
	Season                *string                `json:"season,omitempty"`
	SelectedSeason        *string                `json:"selectedSeason,omitempty"`
	Table                 []*TableElement        `json:"table,omitempty"`
	TopPlayers            *TopPlayers            `json:"topPlayers,omitempty"`
	HasTotw               *bool                  `json:"hasTotw,omitempty"`
	LeagueOverviewMatches []*LeagueOverviewMatch `json:"leagueOverviewMatches,omitempty"`
	HasOngoingMatch       *bool                  `json:"hasOngoingMatch,omitempty"`
}

type LeagueOverviewMatch struct {
	ID                *string        `json:"id,omitempty"`
	PageURL           *string        `json:"pageUrl,omitempty"`
	Opponent          *OpponentClass `json:"opponent,omitempty"`
	Home              *OpponentClass `json:"home,omitempty"`
	Away              *OpponentClass `json:"away,omitempty"`
	DisplayTournament *bool          `json:"displayTournament,omitempty"`
	NotStarted        *bool          `json:"notStarted,omitempty"`
	Tournament        *Tournament    `json:"tournament,omitempty"`
	Status            *Status        `json:"status,omitempty"`
}

type OpponentClass struct {
	ID    *string `json:"id,omitempty"`
	Name  *string `json:"name,omitempty"`
	Score *int    `json:"score,omitempty"`
}

type Tournament struct {
}

type TableElement struct {
	Data         *Data                    `json:"data,omitempty"`
	TeamForm     map[string][]*TeamForm   `json:"teamForm,omitempty"`
	NextOpponent map[string][]interface{} `json:"nextOpponent,omitempty"`
	TableHeader  *TableHeader             `json:"tableHeader,omitempty"`
}

type Data struct {
	Composite        *bool         `json:"composite,omitempty"`
	Ccode            *string       `json:"ccode,omitempty"`
	LeagueID         *int          `json:"leagueId,omitempty"`
	PageURL          *string       `json:"pageUrl,omitempty"`
	LeagueName       *string       `json:"leagueName,omitempty"`
	Legend           []*Legend     `json:"legend,omitempty"`
	TableFilterTypes []string      `json:"tableFilterTypes,omitempty"`
	Tables           []*Tables     `json:"tables,omitempty"`
	Ongoing          []interface{} `json:"ongoing,omitempty"`
	SelectedSeason   *string       `json:"selectedSeason,omitempty"`
	IsCurrentSeason  *bool         `json:"isCurrentSeason,omitempty"`
}

type Legend struct {
	Title   *string `json:"title,omitempty"`
	TKey    *string `json:"tKey,omitempty"`
	Color   *string `json:"color,omitempty"`
	Indices []int   `json:"indices,omitempty"`
}

type Tables struct {
	Ccode      *string    `json:"ccode,omitempty"`
	LeagueID   *int       `json:"leagueId,omitempty"`
	PageURL    *string    `json:"pageUrl,omitempty"`
	LeagueName *string    `json:"leagueName,omitempty"`
	Legend     []*Legend  `json:"legend,omitempty"`
	Table      *DataTable `json:"table,omitempty"`
}

type DataTable struct {
	All  []*All `json:"all,omitempty"`
	Home []*All `json:"home,omitempty"`
	Away []*All `json:"away,omitempty"`
	Form []*All `json:"form,omitempty"`
}

type All struct {
	Name            *string     `json:"name,omitempty"`
	ShortName       *string     `json:"shortName,omitempty"`
	ID              *int        `json:"id,omitempty"`
	PageURL         *string     `json:"pageUrl,omitempty"`
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
	FeaturedInMatch *bool       `json:"featuredInMatch,omitempty"`
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
	Date         interface{}   `json:"date,omitempty"`
	TeamPageURL  *string       `json:"teamPageUrl,omitempty"`
	TooltipText  *TooltipText  `json:"tooltipText,omitempty"`
	Score        *string       `json:"score,omitempty"`
	Home         *TeamFormAway `json:"home,omitempty"`
	Away         *TeamFormAway `json:"away,omitempty"`
}

type TeamFormAway struct {
	IsOurTeam *bool `json:"isOurTeam,omitempty"`
}

type ResultString string

const (
	ResultStringD ResultString = "D"
	ResultStringL ResultString = "L"
	ResultStringW ResultString = "W"
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

type TopPlayers struct {
	ByRating  *ByRating  `json:"byRating,omitempty"`
	ByGoals   *ByGoals   `json:"byGoals,omitempty"`
	ByAssists *ByAssists `json:"byAssists,omitempty"`
	SeeAllURL *string    `json:"seeAllUrl,omitempty"`
}

type ByRating struct {
	ByPlayers []*ByRatingPlayer `json:"players,omitempty"`
}
type ByGoals struct {
	ByPlayers []*ByGoalsPlayer `json:"players,omitempty"`
}
type ByAssists struct {
	ByPlayers []*ByAssistsPlayer `json:"players,omitempty"`
}

type ByPlayer struct {
	ID         *int        `json:"id,omitempty"`
	Name       *string     `json:"name,omitempty"`
	TeamID     *int        `json:"teamId,omitempty"`
	TeamName   *string     `json:"teamName,omitempty"`
	TeamColors interface{} `json:"teamColors,omitempty"`
}

type ByRatingPlayer struct {
	ByPlayer
	Rating float64 `json:"rating,omitempty"`
}
type ByGoalsPlayer struct {
	ByPlayer
	Goals int `json:"goals,omitempty"`
}
type ByAssistsPlayer struct {
	ByPlayer
	Assists int `json:"assists,omitempty"`
}

type Stats struct {
	Players          []*Player         `json:"players,omitempty"`
	Teams            []*Player         `json:"teams,omitempty"`
	Trophies         []*Trophy         `json:"trophies,omitempty"`
	SeasonStatLinks  []*SeasonStatLink `json:"seasonStatLinks,omitempty"`
	SeasonsWithLinks []string          `json:"seasonsWithLinks,omitempty"`
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

type Participant struct {
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

type SeasonStatLink struct {
	Name           *string `json:"Name,omitempty"`
	CountryCode    *string `json:"CountryCode,omitempty"`
	RelativePath   *string `json:"RelativePath,omitempty"`
	League         *string `json:"League,omitempty"`
	StageID        *int    `json:"StageId,omitempty"`
	TemplateID     *int    `json:"TemplateId,omitempty"`
	TournamentID   *int    `json:"TournamentId,omitempty"`
	TotwRoundsLink *string `json:"TotwRoundsLink,omitempty"`
}

type Trophy struct {
	SeasonName *string `json:"seasonName,omitempty"`
	Winner     *Loser  `json:"winner,omitempty"`
	Loser      *Loser  `json:"loser,omitempty"`
}

type Loser struct {
	ID         *int    `json:"id,omitempty"`
	SeasonName *string `json:"seasonName,omitempty"`
	Name       *string `json:"name,omitempty"`
	Winner     *bool   `json:"winner,omitempty"`
}

type Transfers struct {
	Type *string  `json:"type,omitempty"`
	Data []*Datum `json:"data,omitempty"`
}

type Datum struct {
	Name              *string       `json:"name,omitempty"`
	PlayerID          *int          `json:"playerId,omitempty"`
	Position          interface{}   `json:"position,omitempty"`
	TransferDate      *time.Time    `json:"transferDate,omitempty"`
	TransferText      []interface{} `json:"transferText,omitempty"`
	FromClub          *string       `json:"fromClub,omitempty"`
	FromClubID        *int          `json:"fromClubId,omitempty"`
	ToClub            *string       `json:"toClub,omitempty"`
	ToClubID          *int          `json:"toClubId,omitempty"`
	Fee               interface{}   `json:"fee,omitempty"`
	TransferType      interface{}   `json:"transferType,omitempty"`
	ContractExtension *bool         `json:"contractExtension,omitempty"`
	OnLoan            *bool         `json:"onLoan,omitempty"`
	FromDate          *time.Time    `json:"fromDate,omitempty"`
	MarketValue       *string       `json:"marketValue,omitempty"`
	ToDate            *time.Time    `json:"toDate,omitempty"`
}

func (a *Datum) UnmarshalJSON(b []byte) error {
	var s map[string]interface{}
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	marketValue := s["marketValue"]
	s["marketValue"] = nil

	d, err := json.Marshal(s)
	if err != nil {
		return err
	}

	type Alias Datum
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(a),
	}
	if err := json.Unmarshal(d, &aux); err != nil {
		return err
	}

	*a = Datum(*aux.Alias)

	switch marketValue := marketValue.(type) {
	case float64:
		a.MarketValue = ptr.String(fmt.Sprintf("%f", marketValue))
	case string:
		a.MarketValue = ptr.String(marketValue)
	}

	return nil
}

type Fee struct {
	FeeText          interface{} `json:"feeText,omitempty"`
	LocalizedFeeText interface{} `json:"localizedFeeText,omitempty"`
	Value            *string     `json:"value,omitempty"`
}

type FeeText string

const (
	FeeTextFee          FeeText = "fee"
	FeeTextFreeTransfer FeeText = "free transfer"
	FeeTextOnLoan       FeeText = "on loan"
)

type LocalizedFeeText string

const (
	LocalizedFeeTextOnLoan                   LocalizedFeeText = "on_loan"
	LocalizedFeeTextTransferFee              LocalizedFeeText = "transfer_fee"
	LocalizedFeeTextTransferTypeFreeTransfer LocalizedFeeText = "transfer_type_free_transfer"
)

type Position struct {
	Label *string `json:"label,omitempty"`
	Key   *string `json:"key,omitempty"`
}

type Text string

const (
	TextContract Text = "contract"
	TextOnLoan   Text = "on loan"
)

type LocalizationKey string

const (
	LocalizationKeyContract LocalizationKey = "contract"
	LocalizationKeyOnLoan   LocalizationKey = "on_loan"
)

type TransferType struct {
	Text            *Text            `json:"text,omitempty"`
	LocalizationKey *LocalizationKey `json:"localizationKey,omitempty"`
}
