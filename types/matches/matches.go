package matches

import "time"

type Matches struct {
	Leagues []League `json:"leagues,omitempty"`
	Date    *string  `json:"date,omitempty"`
}

type League struct {
	IsGroup          *bool   `json:"isGroup,omitempty"`
	GroupName        *string `json:"groupName,omitempty"`
	Ccode            *string `json:"ccode,omitempty"`
	ID               *int    `json:"id,omitempty"`
	PrimaryID        *int    `json:"primaryId,omitempty"`
	Name             *string `json:"name,omitempty"`
	Matches          []Match `json:"matches,omitempty"`
	ParentLeagueID   *int    `json:"parentLeagueId,omitempty"`
	ParentLeagueName *string `json:"parentLeagueName,omitempty"`
	InternalRank     *int    `json:"internalRank,omitempty"`
	LiveRank         *int    `json:"liveRank,omitempty"`
	SimpleLeague     *bool   `json:"simpleLeague,omitempty"`
}

type Match struct {
	ID               *int    `json:"id,omitempty"`
	LeagueID         *int    `json:"leagueId,omitempty"`
	Time             *string `json:"time,omitempty"`
	Home             *Away   `json:"home,omitempty"`
	Away             *Away   `json:"away,omitempty"`
	EliminatedTeamID *int    `json:"eliminatedTeamId,omitempty"`
	StatusID         *int    `json:"statusId,omitempty"`
	TournamentStage  *string `json:"tournamentStage,omitempty"`
	Status           *Status `json:"status,omitempty"`
	TimeTS           *int    `json:"timeTS,omitempty"`
}

type Away struct {
	ID       *int    `json:"id,omitempty"`
	Score    *int    `json:"score,omitempty"`
	Name     *string `json:"name,omitempty"`
	LongName *string `json:"longName,omitempty"`
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

type Long string

const (
	LongCancelled Long = "Cancelled"
	LongFullTime  Long = "Full-Time"
	LongPostponed Long = "Postponed"
)

type LongKey string

const (
	LongKeyCancelled LongKey = "cancelled"
	LongKeyFinished  LongKey = "finished"
	LongKeyPostponed LongKey = "postponed"
)

type Short string

const (
	ShortCan Short = "Can"
	ShortFt  Short = "FT"
	ShortPp  Short = "PP"
)

type ShortKey string

const (
	ShortKeyCancelledShort ShortKey = "cancelled_short"
	ShortKeyFulltimeShort  ShortKey = "fulltime_short"
	ShortKeyPostponedShort ShortKey = "postponed_short"
)
