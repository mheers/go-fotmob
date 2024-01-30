package matchdetails

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/smithy-go/ptr"
)

type MatchDetails struct {
	General       *General `json:"general,omitempty"`
	Header        *Header  `json:"header,omitempty"`
	Nav           []string `json:"nav,omitempty"`
	Ongoing       *bool    `json:"ongoing,omitempty"`
	HasPendingVAR *bool    `json:"hasPendingVAR,omitempty"`
	Content       *Content `json:"content,omitempty"`
	SEO           *SEO     `json:"seo,omitempty"`
}

type Content struct {
	MatchFacts *MatchFacts    `json:"matchFacts,omitempty"`
	Liveticker *Liveticker    `json:"liveticker,omitempty"`
	Superlive  *Superlive     `json:"superlive,omitempty"`
	Buzz       interface{}    `json:"buzz,omitempty"`
	Stats      *ContentStats  `json:"stats,omitempty"`
	Shotmap    *ShotmapClass  `json:"shotmap,omitempty"`
	Lineup     *ContentLineup `json:"lineup,omitempty"`
	Playoff    *bool          `json:"playoff,omitempty"`
	Table      *Table         `json:"table,omitempty"`
	H2H        *H2H           `json:"h2h,omitempty"`
	Momentum   *Momentum      `json:"momentum,omitempty"`
}

type H2H struct {
	Summary []int   `json:"summary,omitempty"`
	Matches []Match `json:"matches,omitempty"`
}

type Match struct {
	Time     *MatchDate `json:"time,omitempty"`
	MatchURL *string    `json:"matchUrl,omitempty"`
	League   *League    `json:"league,omitempty"`
	Home     *MatchAway `json:"home,omitempty"`
	Status   *Status    `json:"status,omitempty"`
	Finished *bool      `json:"finished,omitempty"`
	Away     *MatchAway `json:"away,omitempty"`
}

type MatchAway struct {
	Name *string `json:"name,omitempty"`
	ID   *string `json:"id,omitempty"`
}

type League struct {
	Name    *string `json:"name,omitempty"`
	ID      *string `json:"id,omitempty"`
	PageURL *string `json:"pageUrl,omitempty"`
}

type Status struct {
	// UTCTime             *time.Time  `json:"utcTime,omitempty"` // TODO: fix this
	Started             *bool       `json:"started,omitempty"`
	Cancelled           *bool       `json:"cancelled,omitempty"`
	Finished            *bool       `json:"finished,omitempty"`
	ScoreStr            *string     `json:"scoreStr,omitempty"`
	Reason              *Reason     `json:"reason,omitempty"`
	WhoLostOnPenalties  interface{} `json:"whoLostOnPenalties,omitempty"`
	WhoLostOnAggregated *string     `json:"whoLostOnAggregated,omitempty"`
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
)

type ShortKey string

const (
	ShortKeyAbortedShort  ShortKey = "aborted_short"
	ShortKeyFulltimeShort ShortKey = "fulltime_short"
)

type Long string

const (
	LongAbandoned Long = "Abandoned"
	LongFullTime  Long = "Full-Time"
)

type LongKey string

const (
	LongKeyAborted  LongKey = "aborted"
	LongKeyFinished LongKey = "finished"
)

type MatchDate struct {
	// UTCTime *time.Time `json:"utcTime,omitempty"` // TODO: fix this
}

// func (m *MatchDate) UnmarshalJSON(b []byte) error {
// 	var s map[string]interface{}
// 	if err := json.Unmarshal(b, &s); err != nil {
// 		return err
// 	}

// 	t, err := time.Parse("2006-01-02T15:04:05Z", s["utcTime"].(string))
// 	if err != nil {
// 		return err
// 	}

// 	m.UTCTime = &t

// 	return nil
// }

type ContentLineup struct {
	Lineup               []LineupElement `json:"lineup,omitempty"`
	Bench                *PurpleBench    `json:"bench,omitempty"`
	NaPlayers            *NaPlayers      `json:"naPlayers,omitempty"`
	Coaches              *Coaches        `json:"coaches,omitempty"`
	TeamRatings          *TeamRatings    `json:"teamRatings,omitempty"`
	HasFantasy           *bool           `json:"hasFantasy,omitempty"`
	UsingEnetpulseLineup *bool           `json:"usingEnetpulseLineup,omitempty"`
	UsingOptaLineup      *bool           `json:"usingOptaLineup,omitempty"`
	SimpleLineup         *bool           `json:"simpleLineup,omitempty"`
}

type PurpleBench struct {
	Sides    []*Side              `json:"sides,omitempty"`
	BenchArr [][]*BenchArrElement `json:"benchArr,omitempty"`
}

type BenchArrElement struct {
	ID                *string            `json:"id,omitempty"`
	PositionID        *int               `json:"positionId,omitempty"`
	Position          *Position          `json:"position,omitempty"`
	LocalizedPosition *interface{}       `json:"localizedPosition,omitempty"`
	Name              *NameClass         `json:"name,omitempty"`
	Shirt             *int               `json:"shirt,omitempty"`
	UsualPosition     *int               `json:"usualPosition,omitempty"`
	UsingOptaID       *bool              `json:"usingOptaId,omitempty"`
	TeamID            *string            `json:"teamId,omitempty"`
	ImageURL          *string            `json:"imageUrl,omitempty"`
	PageURL           *string            `json:"pageUrl,omitempty"`
	IsHomeTeam        *bool              `json:"isHomeTeam,omitempty"`
	TimeSubbedOn      *int               `json:"timeSubbedOn,omitempty"`
	TimeSubbedOff     *interface{}       `json:"timeSubbedOff,omitempty"`
	PositionRow       *int               `json:"positionRow,omitempty"`
	Role              *Role              `json:"role,omitempty"`
	IsCaptain         *bool              `json:"isCaptain,omitempty"`
	Events            *BenchArrEvents    `json:"events,omitempty"`
	Rating            *BenchArrRating    `json:"rating,omitempty"`
	FantasyScore      *FantasyScoreClass `json:"fantasyScore,omitempty"`
	MinutesPlayed     *int               `json:"minutesPlayed,omitempty"`
	Shotmap           []*Shot            `json:"shotmap,omitempty"`
	Stats             []*BenchArrStat    `json:"stats,omitempty"`
	TeamData          *TeamData          `json:"teamData,omitempty"`
	IsGoalkeeper      *bool              `json:"isGoalkeeper,omitempty"`
	ShortName         *string            `json:"shortName,omitempty"`
}

type BenchArrEvents struct {
	Sub *PurpleSub `json:"sub,omitempty"`
	Yc  *int       `json:"yc,omitempty"`
}

type PurpleSub struct {
	SubbedIn *int `json:"subbedIn,omitempty"`
}

type FantasyScoreClass struct {
	Num     *float64 `json:"num,omitempty"`
	Bgcolor *string  `json:"bgcolor,omitempty"`
}

type NameClass struct {
	FirstName *string `json:"firstName,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
	FullName  *string `json:"fullName,omitempty"`
}

type Position string

const (
	PositionSubstitute Position = "Substitute"
)

type BenchArrRating struct {
	Num     *string `json:"num,omitempty"`
	Bgcolor *string `json:"bgcolor,omitempty"`
	IsTop   *IsTop  `json:"isTop,omitempty"`
}

type IsTop struct {
	IsTopRating     *bool `json:"isTopRating,omitempty"`
	IsMatchFinished *bool `json:"isMatchFinished,omitempty"`
}

type Role string

const (
	RoleAttacker   Role = "Attacker"
	RoleDefender   Role = "Defender"
	RoleKeeper     Role = "Keeper"
	RoleMidfielder Role = "Midfielder"
)

type Shot struct {
	ID                    *int        `json:"id,omitempty"`
	EventType             *EventType  `json:"eventType,omitempty"`
	TeamID                *int        `json:"teamId,omitempty"`
	PlayerID              *int        `json:"playerId,omitempty"`
	PlayerName            *string     `json:"playerName,omitempty"`
	X                     *float64    `json:"x,omitempty"`
	Y                     *float64    `json:"y,omitempty"`
	Min                   *int        `json:"min,omitempty"`
	MinAdded              *int        `json:"minAdded,omitempty"`
	IsBlocked             *bool       `json:"isBlocked,omitempty"`
	IsOnTarget            *bool       `json:"isOnTarget,omitempty"`
	BlockedX              *float64    `json:"blockedX,omitempty"`
	BlockedY              *float64    `json:"blockedY,omitempty"`
	GoalCrossedY          *float64    `json:"goalCrossedY,omitempty"`
	GoalCrossedZ          *float64    `json:"goalCrossedZ,omitempty"`
	ExpectedGoals         *float64    `json:"expectedGoals,omitempty"`
	ExpectedGoalsOnTarget *float64    `json:"expectedGoalsOnTarget,omitempty"`
	ShotType              *ShotType   `json:"shotType,omitempty"`
	Situation             *Situation  `json:"situation,omitempty"`
	Period                *Period     `json:"period,omitempty"`
	IsOwnGoal             *bool       `json:"isOwnGoal,omitempty"`
	OnGoalShot            *OnGoalShot `json:"onGoalShot,omitempty"`
	IsSavedOffLine        *bool       `json:"isSavedOffLine,omitempty"`
	TeamColor             *string     `json:"teamColor,omitempty"`
	FirstName             *string     `json:"firstName,omitempty"`
	LastName              *string     `json:"lastName,omitempty"`
	FullName              *string     `json:"fullName,omitempty"`
}

type EventType string

const (
	EventTypeAttemptSaved EventType = "AttemptSaved"
	EventTypeMiss         EventType = "Miss"
)

type OnGoalShot struct {
	X         *float64 `json:"x,omitempty"`
	Y         *float64 `json:"y,omitempty"`
	ZoomRatio *float64 `json:"zoomRatio,omitempty"`
}

type Period string

const (
	PeriodFirstHalf  Period = "FirstHalf"
	PeriodSecondHalf Period = "SecondHalf"
)

type ShotType string

const (
	ShotTypeHeader    ShotType = "Header"
	ShotTypeLeftFoot  ShotType = "LeftFoot"
	ShotTypeRightFoot ShotType = "RightFoot"
)

type Situation string

const (
	SituationFastBreak   Situation = "FastBreak"
	SituationFromCorner  Situation = "FromCorner"
	SituationRegularPlay Situation = "RegularPlay"
	SituationSetPiece    Situation = "SetPiece"
)

type BenchArrStat struct {
	Title *Title       `json:"title,omitempty"`
	Key   *StatKey     `json:"key,omitempty"`
	Stats *PurpleStats `json:"stats,omitempty"`
}

type StatKey string

const (
	StatKeyAttack   StatKey = "attack"
	StatKeyDefense  StatKey = "defense"
	StatKeyDuels    StatKey = "duels"
	StatKeyTopStats StatKey = "top_stats"
)

type PurpleStats struct {
	FotMobRating              *Assists         `json:"FotMob rating,omitempty"`
	MinutesPlayed             *Assists         `json:"Minutes played,omitempty"`
	Goals                     *Assists         `json:"Goals,omitempty"`
	Assists                   *Assists         `json:"Assists,omitempty"`
	TotalShots                *Assists         `json:"Total shots,omitempty"`
	Shotmap                   *Shotmap         `json:"Shotmap,omitempty"`
	AccuratePasses            *AccurateCrosses `json:"Accurate passes,omitempty"`
	ChancesCreated            *Assists         `json:"Chances created,omitempty"`
	ExpectedGoalsXG           *AccurateCrosses `json:"Expected goals (xG),omitempty"`
	ExpectedAssistsXA         *AccurateCrosses `json:"Expected assists (xA),omitempty"`
	XGPlusXA                  *AccurateCrosses `json:"xG + xA,omitempty"`
	FantasyPoints             *Assists         `json:"Fantasy points,omitempty"`
	ShotAccuracy              *AccurateCrosses `json:"Shot accuracy,omitempty"`
	BigChancesMissed          *Assists         `json:"Big chances missed,omitempty"`
	BlockedShots              *Assists         `json:"Blocked shots,omitempty"`
	Touches                   *Assists         `json:"Touches,omitempty"`
	TouchesInOppositionBox    *Assists         `json:"Touches in opposition box,omitempty"`
	PassesIntoFinalThird      *Assists         `json:"Passes into final third,omitempty"`
	AccurateCrosses           *AccurateCrosses `json:"Accurate crosses,omitempty"`
	Dispossessed              *Assists         `json:"Dispossessed,omitempty"`
	XGNonPenalty              *AccurateCrosses `json:"xG Non-penalty,omitempty"`
	TacklesWon                *AerialDuelsWon  `json:"Tackles won,omitempty"`
	DefensiveActions          *Assists         `json:"Defensive actions,omitempty"`
	GroundDuelsWon            *AerialDuelsWon  `json:"Ground duels won,omitempty"`
	AerialDuelsWon            *AerialDuelsWon  `json:"Aerial duels won,omitempty"`
	WasFouled                 *Assists         `json:"Was fouled,omitempty"`
	FoulsCommitted            *Assists         `json:"Fouls committed,omitempty"`
	AccurateLongBalls         *AccurateCrosses `json:"Accurate long balls,omitempty"`
	Recoveries                *Assists         `json:"Recoveries,omitempty"`
	DuelsLost                 *Assists         `json:"Duels lost,omitempty"`
	DuelsWon                  *Assists         `json:"Duels won,omitempty"`
	ExpectedGoalsOnTargetXGOT *AccurateCrosses `json:"Expected goals on target (xGOT),omitempty"`
	DribbledPast              *Assists         `json:"Dribbled past,omitempty"`
	Clearances                *Assists         `json:"Clearances,omitempty"`
	SuccessfulDribbles        *AccurateCrosses `json:"Successful dribbles,omitempty"`
	HeadedClearance           *Assists         `json:"Headed clearance,omitempty"`
}

type AccurateCrosses struct {
	Key   *AccurateCrossesKey `json:"key,omitempty"`
	Value *string             `json:"value,omitempty"`
}

func (a *AccurateCrosses) UnmarshalJSON(b []byte) error {
	var s map[string]interface{}
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	value := s["value"]
	key := s["key"].(string)

	switch value := value.(type) {
	case float64:
		a.Value = ptr.String(fmt.Sprintf("%f", value))
	case string:
		a.Value = ptr.String(value)
	}

	k := AccurateCrossesKey(key)
	a.Key = &k

	return nil
}

type AccurateCrossesKey string

const (
	AccurateCrossesN             AccurateCrossesKey = "accurate_crosses"
	AccuratePasses               AccurateCrossesKey = "accurate_passes"
	AerialsWon                   AccurateCrossesKey = "aerials_won"
	DribblesSucceeded            AccurateCrossesKey = "dribbles_succeeded"
	ExpectedAssists              AccurateCrossesKey = "expected_assists"
	ExpectedGoals                AccurateCrossesKey = "expected_goals"
	ExpectedGoalsNonPenalty      AccurateCrossesKey = "expected_goals_non_penalty"
	ExpectedGoalsOnTargetFaced   AccurateCrossesKey = "expected_goals_on_target_faced"
	ExpectedGoalsOnTargetVariant AccurateCrossesKey = "expected_goals_on_target_variant"
	FantasyPointsAccurateCrosses AccurateCrossesKey = "fantasy_points"
	GoalsPrevented               AccurateCrossesKey = "goals_prevented"
	GroundDuelsWon               AccurateCrossesKey = "ground_duels_won"
	LongBallsAccurate            AccurateCrossesKey = "long_balls_accurate"
	Saves                        AccurateCrossesKey = "saves"
	ShotAccuracy                 AccurateCrossesKey = "shot_accuracy"
	TacklesSucceeded             AccurateCrossesKey = "tackles_succeeded"
	XgAndXa                      AccurateCrossesKey = "xg_and_xa"
)

type AerialDuelsWon struct {
	Key   *AccurateCrossesKey `json:"key,omitempty"`
	Value *string             `json:"value,omitempty"`
}

func (a *AerialDuelsWon) UnmarshalJSON(b []byte) error {
	var s map[string]interface{}
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	value := s["value"]
	key := s["key"].(string)

	switch value := value.(type) {
	case float64:
		a.Value = ptr.String(fmt.Sprintf("%f", value))
	case string:
		a.Value = ptr.String(value)
	}

	k := AccurateCrossesKey(key)
	a.Key = &k

	return nil
}

type Assists struct {
	Key   *AssistsKey `json:"key,omitempty"`
	Value *float64    `json:"value,omitempty"`
}

type AssistsKey string

const (
	AssistsN             AssistsKey = "assists"
	BigChanceMissedTitle AssistsKey = "big_chance_missed_title"
	BlockedShots         AssistsKey = "blocked_shots"
	ChancesCreated       AssistsKey = "chances_created"
	Clearances           AssistsKey = "clearances"
	Corners              AssistsKey = "corners"
	DefensiveActions     AssistsKey = "defensive_actions"
	Dispossessed         AssistsKey = "dispossessed"
	DribbledPast         AssistsKey = "dribbled_past"
	DuelLost             AssistsKey = "duel_lost"
	DuelWon              AssistsKey = "duel_won"
	FantasyPoints        AssistsKey = "fantasy_points"
	Fouls                AssistsKey = "fouls"
	Goals                AssistsKey = "goals"
	GoalsConceded        AssistsKey = "goals_conceded"
	HeadedClearance      AssistsKey = "headed_clearance"
	Interceptions        AssistsKey = "interceptions"
	KeeperDivingSave     AssistsKey = "keeper_diving_save"
	KeeperHighClaim      AssistsKey = "keeper_high_claim"
	KeeperSweeper        AssistsKey = "keeper_sweeper"
	LastManTackle        AssistsKey = "last_man_tackle"
	MinutesPlayed        AssistsKey = "minutes_played"
	Offsides             AssistsKey = "Offsides"
	PassesIntoFinalThird AssistsKey = "passes_into_final_third"
	PlayerThrows         AssistsKey = "player_throws"
	Punches              AssistsKey = "punches"
	RatingTitle          AssistsKey = "rating_title"
	Recoveries           AssistsKey = "recoveries"
	SavesInsideBox       AssistsKey = "saves_inside_box"
	ShotBlocks           AssistsKey = "shot_blocks"
	TotalShots           AssistsKey = "total_shots"
	Touches              AssistsKey = "touches"
	TouchesOppBox        AssistsKey = "touches_opp_box"
	WasFouled            AssistsKey = "was_fouled"
)

type Shotmap struct {
	Key   *string `json:"key,omitempty"`
	Value *bool   `json:"value,omitempty"`
}

type Title string

const (
	Attack   Title = "Attack"
	Defense  Title = "Defense"
	Duels    Title = "Duels"
	TopStats Title = "Top stats"
)

type TeamData struct {
	Home *TeamDataAway `json:"home,omitempty"`
	Away *TeamDataAway `json:"away,omitempty"`
}

type TeamDataAway struct {
	ID    *int    `json:"id,omitempty"`
	Color *string `json:"color,omitempty"`
}

type Side string

const (
	Away  Side = "away"
	Equal Side = "equal"
	Home  Side = "home"
)

type Coaches struct {
	Sides      []Side    `json:"sides,omitempty"`
	CoachesArr [][]Coach `json:"coachesArr,omitempty"`
}

type Coach struct {
	ID            *string           `json:"id,omitempty"`
	Name          *NameClass        `json:"name,omitempty"`
	UsualPosition *int              `json:"usualPosition,omitempty"`
	ImageURL      *string           `json:"imageUrl,omitempty"`
	PageURL       *string           `json:"pageUrl,omitempty"`
	IsHomeTeam    *bool             `json:"isHomeTeam,omitempty"`
	TimeSubbedOn  *string           `json:"timeSubbedOn,omitempty"`
	TimeSubbedOff *string           `json:"timeSubbedOff,omitempty"`
	PositionRow   *string           `json:"positionRow,omitempty"`
	Role          *string           `json:"role,omitempty"`
	IsCaptain     *bool             `json:"isCaptain,omitempty"`
	Events        *CoachesArrEvents `json:"events,omitempty"`
}

type CoachesArrEvents struct {
}

type LineupElement struct {
	TeamID              *int              `json:"teamId,omitempty"`
	TeamName            *string           `json:"teamName,omitempty"`
	Bench               []BenchArrElement `json:"bench,omitempty"`
	Coach               []Coach           `json:"coach,omitempty"`
	Players             [][]PlayerElement `json:"players,omitempty"`
	Lineup              *string           `json:"lineup,omitempty"`
	NonAvailablePlayers [][]NaPlayersArr  `json:"nonAvailablePlayers,omitempty"`
	OptaLineup          *OptaLineup       `json:"optaLineup,omitempty"`
}

type NaPlayersArr struct {
	ID            *int              `json:"id,omitempty"`
	Name          *NameClass        `json:"name,omitempty"`
	ShortName     *string           `json:"shortName,omitempty"`
	ImageURL      *string           `json:"imageUrl,omitempty"`
	PageURL       *string           `json:"pageUrl,omitempty"`
	IsHomeTeam    *bool             `json:"isHomeTeam,omitempty"`
	TimeSubbedOn  *interface{}      `json:"timeSubbedOn,omitempty"`
	TimeSubbedOff *interface{}      `json:"timeSubbedOff,omitempty"`
	PositionRow   *interface{}      `json:"positionRow,omitempty"`
	IsCaptain     *bool             `json:"isCaptain,omitempty"`
	Events        *CoachesArrEvents `json:"events,omitempty"`
	NaInfo        *NaInfo           `json:"naInfo,omitempty"`
}

type NaInfo struct {
	ID             *int            `json:"id,omitempty"`
	Name           *string         `json:"name,omitempty"`
	NaReason       *string         `json:"naReason,omitempty"`
	NaReasonKey    *string         `json:"naReasonKey,omitempty"`
	Injury         *bool           `json:"injury,omitempty"`
	ExpectedReturn *ExpectedReturn `json:"expectedReturn,omitempty"`
}

type ExpectedReturn struct {
	ExpectedReturnKey       *string `json:"expectedReturnKey,omitempty"`
	ExpectedReturnDateParam *string `json:"expectedReturnDateParam,omitempty"`
	ExpectedReturnFallback  *string `json:"expectedReturnFallback,omitempty"`
}

type OptaLineup struct {
	Bench               []BenchArrElement `json:"bench,omitempty"`
	Coach               []Coach           `json:"coach,omitempty"`
	Players             [][]PlayerElement `json:"players,omitempty"`
	Lineup              *string           `json:"lineup,omitempty"`
	NonAvailablePlayers [][]NaPlayersArr  `json:"nonAvailablePlayers,omitempty"`
}

type PlayerElement struct {
	ID                  *string         `json:"id,omitempty"`
	PositionID          *int            `json:"positionId,omitempty"`
	Position            *Role           `json:"position,omitempty"`
	PositionStringShort *string         `json:"positionStringShort,omitempty"`
	LocalizedPosition   *PositionLabel  `json:"localizedPosition,omitempty"`
	Name                *NameClass      `json:"name,omitempty"`
	Shirt               *int            `json:"shirt,omitempty"`
	UsualPosition       *int            `json:"usualPosition,omitempty"`
	UsingOptaID         *bool           `json:"usingOptaId,omitempty"`
	TeamID              *string         `json:"teamId,omitempty"`
	ImageURL            *string         `json:"imageUrl,omitempty"`
	PageURL             *string         `json:"pageUrl,omitempty"`
	IsHomeTeam          *bool           `json:"isHomeTeam,omitempty"`
	TimeSubbedOn        *interface{}    `json:"timeSubbedOn,omitempty"`
	TimeSubbedOff       *int            `json:"timeSubbedOff,omitempty"`
	PositionRow         *int            `json:"positionRow,omitempty"`
	Role                *Role           `json:"role,omitempty"`
	IsCaptain           *bool           `json:"isCaptain,omitempty"`
	Events              *PlayerEvents   `json:"events,omitempty"`
	Rating              *BenchArrRating `json:"rating,omitempty"`
	FantasyScore        *FantasyScore   `json:"fantasyScore,omitempty"`
	MinutesPlayed       *int            `json:"minutesPlayed,omitempty"`
	Shotmap             []Shot          `json:"shotmap,omitempty"`
	Stats               []PlayerStat    `json:"stats,omitempty"`
	TeamData            *TeamData       `json:"teamData,omitempty"`
	IsGoalkeeper        *bool           `json:"isGoalkeeper,omitempty"`
}

type PlayerEvents struct {
	Sub  *FluffySub `json:"sub,omitempty"`
	Yc   *int       `json:"yc,omitempty"`
	Ycrc *int       `json:"ycrc,omitempty"`
}

type FluffySub struct {
	SubbedOut *int `json:"subbedOut,omitempty"`
}

type FantasyScore struct {
	Num     *string `json:"num,omitempty"`
	Bgcolor *string `json:"bgcolor,omitempty"`
}

type PositionLabel struct {
	Label *string `json:"label,omitempty"`
	Key   *string `json:"key,omitempty"`
}

type PlayerStat struct {
	Title *Title       `json:"title,omitempty"`
	Key   *StatKey     `json:"key,omitempty"`
	Stats *FluffyStats `json:"stats,omitempty"`
}

type FluffyStats struct {
	FotMobRating              *Assists         `json:"FotMob rating,omitempty"`
	MinutesPlayed             *Assists         `json:"Minutes played,omitempty"`
	Saves                     *AccurateCrosses `json:"Saves,omitempty"`
	GoalsConceded             *Assists         `json:"Goals conceded,omitempty"`
	XGOTFaced                 *AccurateCrosses `json:"xGOT faced,omitempty"`
	GoalsPrevented            *AccurateCrosses `json:"Goals prevented,omitempty"`
	AccuratePasses            *AccurateCrosses `json:"Accurate passes,omitempty"`
	AccurateLongBalls         *AccurateCrosses `json:"Accurate long balls,omitempty"`
	DivingSave                *Assists         `json:"Diving save,omitempty"`
	SavesInsideBox            *Assists         `json:"Saves inside box,omitempty"`
	ActedAsSweeper            *Assists         `json:"Acted as sweeper,omitempty"`
	Punches                   *Assists         `json:"Punches,omitempty"`
	Throws                    *Assists         `json:"Throws,omitempty"`
	HighClaim                 *Assists         `json:"High claim,omitempty"`
	Recoveries                *Assists         `json:"Recoveries,omitempty"`
	FantasyPoints             *AerialDuelsWon  `json:"Fantasy points,omitempty"`
	Touches                   *Assists         `json:"Touches,omitempty"`
	Goals                     *Assists         `json:"Goals,omitempty"`
	Assists                   *Assists         `json:"Assists,omitempty"`
	TotalShots                *Assists         `json:"Total shots,omitempty"`
	Shotmap                   *Shotmap         `json:"Shotmap,omitempty"`
	ChancesCreated            *Assists         `json:"Chances created,omitempty"`
	ExpectedGoalsXG           *AccurateCrosses `json:"Expected goals (xG),omitempty"`
	ExpectedAssistsXA         *AccurateCrosses `json:"Expected assists (xA),omitempty"`
	XGXa                      *AccurateCrosses `json:"xG + xA,omitempty"`
	ShotAccuracy              *AccurateCrosses `json:"Shot accuracy,omitempty"`
	BlockedShots              *Assists         `json:"Blocked shots,omitempty"`
	TouchesInOppositionBox    *Assists         `json:"Touches in opposition box,omitempty"`
	SuccessfulDribbles        *AccurateCrosses `json:"Successful dribbles,omitempty"`
	PassesIntoFinalThird      *Assists         `json:"Passes into final third,omitempty"`
	AccurateCrosses           *AccurateCrosses `json:"Accurate crosses,omitempty"`
	Corners                   *Assists         `json:"Corners,omitempty"`
	Dispossessed              *Assists         `json:"Dispossessed,omitempty"`
	XGNonPenalty              *AccurateCrosses `json:"xG Non-penalty,omitempty"`
	TacklesWon                *AerialDuelsWon  `json:"Tackles won,omitempty"`
	LastManTackle             *Assists         `json:"Last man tackle,omitempty"`
	Clearances                *Assists         `json:"Clearances,omitempty"`
	Interceptions             *Assists         `json:"Interceptions,omitempty"`
	DefensiveActions          *Assists         `json:"Defensive actions,omitempty"`
	DribbledPast              *Assists         `json:"Dribbled past,omitempty"`
	DuelsWon                  *Assists         `json:"Duels won,omitempty"`
	DuelsLost                 *Assists         `json:"Duels lost,omitempty"`
	GroundDuelsWon            *AerialDuelsWon  `json:"Ground duels won,omitempty"`
	AerialDuelsWon            *AerialDuelsWon  `json:"Aerial duels won,omitempty"`
	WasFouled                 *Assists         `json:"Was fouled,omitempty"`
	FoulsCommitted            *Assists         `json:"Fouls committed,omitempty"`
	ExpectedGoalsOnTargetXGOT *AccurateCrosses `json:"Expected goals on target (xGOT),omitempty"`
	HeadedClearance           *Assists         `json:"Headed clearance,omitempty"`
	Blocks                    *Assists         `json:"Blocks,omitempty"`
	Offsides                  *Assists         `json:"Offsides,omitempty"`
	BigChancesMissed          *Assists         `json:"Big chances missed,omitempty"`
}

type NaPlayers struct {
	Sides        []Side           `json:"sides,omitempty"`
	NaPlayersArr [][]NaPlayersArr `json:"naPlayersArr,omitempty"`
}

type TeamRatings struct {
	Home *FantasyScoreClass `json:"home,omitempty"`
	Away *FantasyScoreClass `json:"away,omitempty"`
}

type Liveticker struct {
	Langs *string  `json:"langs,omitempty"`
	Teams []string `json:"teams,omitempty"`
}

type MatchFacts struct {
	MatchID          *int              `json:"matchId,omitempty"`
	Highlights       *Highlights       `json:"highlights,omitempty"`
	PlayerOfTheMatch *PlayerOfTheMatch `json:"playerOfTheMatch,omitempty"`
	MatchesInRound   []*MatchesInRound `json:"matchesInRound,omitempty"`
	Events           *MatchFactsEvents `json:"events,omitempty"`
	InfoBox          *InfoBox          `json:"infoBox,omitempty"`
	TeamForm         [][]*TeamForm     `json:"teamForm,omitempty"`
	Poll             *Poll             `json:"poll,omitempty"`
	TopPlayers       *TopPlayers       `json:"topPlayers,omitempty"`
	Insights         []*Insight        `json:"insights,omitempty"`
	TopScorers       *TopScorers       `json:"topScorers,omitempty"`
	Momentum         *Momentum         `json:"momentum,omitempty"`
	CountryCode      *string           `json:"countryCode,omitempty"`
	QAData           []*QADatum        `json:"QAData,omitempty"`
}

type QADatum struct {
	Question *string `json:"question,omitempty"`
	Answer   *string `json:"answer,omitempty"`
}

type MatchFactsEvents struct {
	Ongoing               *bool        `json:"ongoing,omitempty"`
	Events                []*Event     `json:"events,omitempty"`
	EventTypes            EventTypes   `json:"eventTypes,omitempty"`
	PenaltyShootoutEvents *interface{} `json:"penaltyShootoutEvents,omitempty"`
}

type EventTypes []*string

func (mfe *EventTypes) UnmarshalJSON(data []byte) error {
	var err error
	var aux []any
	if err = json.Unmarshal(data, &aux); err != nil {
		return err
	}

	for _, v := range aux {
		switch v := v.(type) {
		case string:
			*mfe = append(*mfe, &v)
		}
	}

	return nil
}

type Event struct {
	ReactKey          *string `json:"reactKey,omitempty"`
	TimeStr           *string `json:"timeStr,omitempty"`
	Type              *string `json:"type,omitempty"`
	Time              *int    `json:"time,omitempty"`
	OverloadTime      *int    `json:"overloadTime,omitempty"`
	EventID           *int    `json:"eventId,omitempty"`
	Player            *Player `json:"player,omitempty"`
	ProfileURL        *string `json:"profileUrl,omitempty"`
	OverloadTimeStr   *string `json:"overloadTimeStr,omitempty"`
	IsHome            *bool   `json:"isHome,omitempty"`
	NameStr           *string `json:"nameStr,omitempty"`
	Card              *string `json:"card,omitempty"`
	CardDescription   *string `json:"cardDescription,omitempty"`
	MinutesAddedStr   *string `json:"minutesAddedStr,omitempty"`
	MinutesAddedKey   *string `json:"minutesAddedKey,omitempty"`
	MinutesAddedInput *int    `json:"minutesAddedInput,omitempty"`
	HalfStrShort      *string `json:"halfStrShort,omitempty"`
	HalfStrKey        *string `json:"halfStrKey,omitempty"`
	InjuredPlayerOut  *bool   `json:"injuredPlayerOut,omitempty"`
	Swap              []*Swap `json:"swap,omitempty"`
}

func (e *Event) UnmarshalJSON(data []byte) error {
	var v map[string]interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	timeStr := fmt.Sprintf("%f", v["timeStr"])

	overloadTimeResult := ""
	overloadTime := v["overloadTimeStr"]
	switch overloadTime := overloadTime.(type) {
	case bool:
		overloadTimeResult = ""
	case string:
		overloadTimeResult = overloadTime
	}

	v2 := v
	v2["timeStr"] = nil
	v2["overloadTimeStr"] = nil

	d, err := json.Marshal(v2)
	if err != nil {
		return err
	}

	type Alias Event
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(e),
	}
	if err := json.Unmarshal(d, &aux); err != nil {
		return err
	}

	*e = Event(*aux.Alias)

	e.TimeStr = ptr.String(timeStr)
	e.OverloadTimeStr = ptr.String(overloadTimeResult)

	return nil
}

type Player struct {
	ID         *int    `json:"id,omitempty"`
	Name       *string `json:"name,omitempty"`
	ProfileURL *string `json:"profileUrl,omitempty"`
}

type Swap struct {
	Name       *string `json:"name,omitempty"`
	ID         *string `json:"id,omitempty"`
	ProfileURL *string `json:"profileUrl,omitempty"`
}

type EventTypeEnum string

const (
	AddedTime    EventTypeEnum = "AddedTime"
	Card         EventTypeEnum = "Card"
	Half         EventTypeEnum = "Half"
	Substitution EventTypeEnum = "Substitution"
)

type Highlights struct {
	Image  *string `json:"image,omitempty"`
	URL    *string `json:"url,omitempty"`
	Source *string `json:"source,omitempty"`
}

type InfoBox struct {
	MatchDate  *MatchDate  `json:"Match Date,omitempty"`
	Tournament *Tournament `json:"Tournament,omitempty"`
	Stadium    *Stadium    `json:"Stadium,omitempty"`
	Referee    *Referee    `json:"Referee,omitempty"`
	Attendance *int        `json:"Attendance,omitempty"`
}

type Referee struct {
	ImgURL  *string `json:"imgUrl,omitempty"`
	Text    *string `json:"text,omitempty"`
	Country *string `json:"country,omitempty"`
}

type Stadium struct {
	Name    *string  `json:"name,omitempty"`
	City    *string  `json:"city,omitempty"`
	Country *string  `json:"country,omitempty"`
	Lat     *float64 `json:"lat,omitempty"`
	Long    *float64 `json:"long,omitempty"`
}

type Tournament struct {
	ID              *int    `json:"id,omitempty"`
	ParentLeagueID  *int    `json:"parentLeagueId,omitempty"`
	Link            *string `json:"link,omitempty"`
	LeagueName      *string `json:"leagueName,omitempty"`
	RoundName       *string `json:"roundName,omitempty"`
	Round           *string `json:"round,omitempty"`
	SelectedSeason  *string `json:"selectedSeason,omitempty"`
	IsCurrentSeason *bool   `json:"isCurrentSeason,omitempty"`
}

type Insight struct {
	Type            *string      `json:"type,omitempty"`
	PlayerID        *int         `json:"playerId,omitempty"`
	TeamID          *int         `json:"teamId,omitempty"`
	Priority        *int         `json:"priority,omitempty"`
	DefaultText     *string      `json:"defaultText,omitempty"`
	LocalizedTextID *string      `json:"localizedTextId,omitempty"`
	StatValues      []*StatValue `json:"statValues,omitempty"`
	Text            *string      `json:"text,omitempty"`
	Color           *string      `json:"color,omitempty"`
}

type StatValue struct {
	Value *float64 `json:"value,omitempty"`
	Name  *string  `json:"name,omitempty"`
	Type  *string  `json:"type,omitempty"`
}

type MatchesInRound struct {
	ID        *string             `json:"id,omitempty"`
	Home      *MatchesInRoundAway `json:"home,omitempty"`
	Away      *MatchesInRoundAway `json:"away,omitempty"`
	MatchDate *string             `json:"matchDate,omitempty"`
	Status    *Status             `json:"status,omitempty"`
	RoundID   *int                `json:"roundId,omitempty"`
	RoundName *string             `json:"roundName,omitempty"`
	HomeScore *int                `json:"homeScore,omitempty"`
	AwayScore *int                `json:"awayScore,omitempty"`
}

type MatchesInRoundAway struct {
	ID        *string `json:"id,omitempty"`
	Name      *string `json:"name,omitempty"`
	ShortName *string `json:"shortName,omitempty"`
}

type Momentum struct {
	Main            *Main          `json:"main,omitempty"`
	AlternateModels []*interface{} `json:"alternateModels,omitempty"`
}

type Main struct {
	Data       []*Datum `json:"data,omitempty"`
	DebugTitle *string  `json:"debugTitle,omitempty"`
}

type Datum struct {
	Minute *float64 `json:"minute,omitempty"`
	Value  *float64 `json:"value,omitempty"`
}

type PlayerOfTheMatch struct {
	ID            *int                    `json:"id,omitempty"`
	Name          *NameClass              `json:"name,omitempty"`
	TeamName      *string                 `json:"teamName,omitempty"`
	TeamID        *int                    `json:"teamId,omitempty"`
	Rating        *PlayerOfTheMatchRating `json:"rating,omitempty"`
	MinutesPlayed *int                    `json:"minutesPlayed,omitempty"`
	Shotmap       []*Shot                 `json:"shotmap,omitempty"`
	PageURL       *string                 `json:"pageUrl,omitempty"`
	IsHomeTeam    *bool                   `json:"isHomeTeam,omitempty"`
	Stats         []*PlayerOfTheMatchStat `json:"stats,omitempty"`
	Role          *Role                   `json:"role,omitempty"`
	TeamData      *TeamData               `json:"teamData,omitempty"`
}

type PlayerOfTheMatchRating struct {
	Num   *string `json:"num,omitempty"`
	IsTop *IsTop  `json:"isTop,omitempty"`
}

type PlayerOfTheMatchStat struct {
	Title *Title          `json:"title,omitempty"`
	Key   *StatKey        `json:"key,omitempty"`
	Stats *TentacledStats `json:"stats,omitempty"`
}

type TentacledStats struct {
	FotMobRating           *Assists         `json:"FotMob rating,omitempty"`
	MinutesPlayed          *Assists         `json:"Minutes played,omitempty"`
	Goals                  *Assists         `json:"Goals,omitempty"`
	Assists                *Assists         `json:"Assists,omitempty"`
	TotalShots             *Assists         `json:"Total shots,omitempty"`
	Shotmap                *Shotmap         `json:"Shotmap,omitempty"`
	AccuratePasses         *AccurateCrosses `json:"Accurate passes,omitempty"`
	ChancesCreated         *Assists         `json:"Chances created,omitempty"`
	ExpectedGoals          *AccurateCrosses `json:"Expected goals (xG),omitempty"`
	ExpectedAssists        *AccurateCrosses `json:"Expected assists (xA),omitempty"`
	XGPlusXA               *AccurateCrosses `json:"xG + xA,omitempty"`
	FantasyPoints          *AccurateCrosses `json:"Fantasy points,omitempty"`
	ShotAccuracy           *AccurateCrosses `json:"Shot accuracy,omitempty"`
	BlockedShots           *Assists         `json:"Blocked shots,omitempty"`
	Touches                *Assists         `json:"Touches,omitempty"`
	TouchesInOppositionBox *Assists         `json:"Touches in opposition box,omitempty"`
	SuccessfulDribbles     *AccurateCrosses `json:"Successful dribbles,omitempty"`
	PassesIntoFinalThird   *Assists         `json:"Passes into final third,omitempty"`
	AccurateCrosses        *AccurateCrosses `json:"Accurate crosses,omitempty"`
	AccurateLongBalls      *AccurateCrosses `json:"Accurate long balls,omitempty"`
	Corners                *Assists         `json:"Corners,omitempty"`
	Dispossessed           *Assists         `json:"Dispossessed,omitempty"`
	XGNonPenalty           *AccurateCrosses `json:"xG Non-penalty,omitempty"`
	TacklesWon             *AccurateCrosses `json:"Tackles won,omitempty"`
	LastManTackle          *Assists         `json:"Last man tackle,omitempty"`
	Clearances             *Assists         `json:"Clearances,omitempty"`
	Interceptions          *Assists         `json:"Interceptions,omitempty"`
	DefensiveActions       *Assists         `json:"Defensive actions,omitempty"`
	Recoveries             *Assists         `json:"Recoveries,omitempty"`
	DribbledPast           *Assists         `json:"Dribbled past,omitempty"`
	DuelsWon               *Assists         `json:"Duels won,omitempty"`
	DuelsLost              *Assists         `json:"Duels lost,omitempty"`
	GroundDuelsWon         *AccurateCrosses `json:"Ground duels won,omitempty"`
	AerialDuelsWon         *AccurateCrosses `json:"Aerial duels won,omitempty"`
	WasFouled              *Assists         `json:"Was fouled,omitempty"`
	FoulsCommitted         *Assists         `json:"Fouls committed,omitempty"`
}

type Poll struct {
	Oddspoll    *Oddspoll   `json:"oddspoll,omitempty"`
	VoteResult  *VoteResult `json:"voteResult,omitempty"`
	RenderToTop *bool       `json:"renderToTop,omitempty"`
}

type Oddspoll struct {
	PollName   *string `json:"PollName,omitempty"`
	MatchId    *int    `json:"MatchId,omitempty"`
	HomeTeamId *int    `json:"HomeTeamId,omitempty"`
	AwayTeamId *int    `json:"AwayTeamId,omitempty"`
	HomeTeam   *string `json:"HomeTeam,omitempty"`
	AwayTeam   *string `json:"AwayTeam,omitempty"`
	Facts      []*Fact `json:"Facts,omitempty"`
}

type Fact struct {
	OddsType        *string  `json:"OddsType,omitempty"`
	DefaultLabel    *string  `json:"DefaultLabel,omitempty"`
	TextLabelId     *string  `json:"TextLabelId,omitempty"`
	DefaultTemplate *string  `json:"DefaultTemplate,omitempty"`
	TextTemplateId  *string  `json:"TextTemplateId,omitempty"`
	StatValues      []string `json:"StatValues,omitempty"`
	DefaultLabels   []string `json:"DefaultLabels,omitempty"`
	LabelTemplates  []string `json:"LabelTemplates,omitempty"`
	Icon            *string  `json:"Icon,omitempty"`
	DefaultText     *string  `json:"defaultText,omitempty"`
}

type VoteResult struct {
	PollName *string `json:"PollName,omitempty"`
	Votes    []*Vote `json:"Votes,omitempty"`
}

type Vote struct {
	Name  *string `json:"Name,omitempty"`
	Votes []int   `json:"Votes,omitempty"`
}

type TeamForm struct {
	Result       *int          `json:"result,omitempty"`
	ResultString *ResultString `json:"resultString,omitempty"`
	ImageUrl     *string       `json:"imageUrl,omitempty"`
	LinkToMatch  *string       `json:"linkToMatch,omitempty"`
	// Date         *time.Time    `json:"date,omitempty"` // TODO: fix
	TeamPageUrl *string       `json:"teamPageUrl,omitempty"`
	TooltipText *TooltipText  `json:"tooltipText,omitempty"`
	Score       *string       `json:"score,omitempty"`
	Home        *TeamFormAway `json:"home,omitempty"`
	Away        *TeamFormAway `json:"away,omitempty"`
}

type TeamFormAway struct {
	ID        *string `json:"id,omitempty"`
	Name      *string `json:"name,omitempty"`
	IsOurTeam *bool   `json:"isOurTeam,omitempty"`
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
	HomeTeamId *int       `json:"homeTeamId,omitempty"`
	HomeScore  *string    `json:"homeScore,omitempty"`
	AwayTeam   *string    `json:"awayTeam,omitempty"`
	AwayTeamId *int       `json:"awayTeamId,omitempty"`
	AwayScore  *string    `json:"awayScore,omitempty"`
}

type TopPlayers struct {
	HomeTopPlayers []*TopPlayer `json:"homeTopPlayers,omitempty"`
	AwayTopPlayers []*TopPlayer `json:"awayTopPlayers,omitempty"`
}

type TopPlayer struct {
	PlayerID            *int           `json:"playerId,omitempty"`
	Name                *NameClass     `json:"name,omitempty"`
	PlayerRatingRounded *string        `json:"playerRatingRounded,omitempty"`
	Color               *string        `json:"color,omitempty"`
	ManOfTheMatch       *bool          `json:"manOfTheMatch,omitempty"`
	TeamID              *string        `json:"teamId,omitempty"`
	PositionLabel       *PositionLabel `json:"positionLabel,omitempty"`
}

type TopScorers struct {
	Section    *string      `json:"section,omitempty"`
	SectionID  *int         `json:"sectionId,omitempty"`
	HomePlayer *PlayerStats `json:"homePlayer,omitempty"`
	AwayPlayer *PlayerStats `json:"awayPlayer,omitempty"`
	StatsOrder []string     `json:"statsOrder,omitempty"`
}

type PlayerStats struct {
	PlayerID          *int             `json:"playerId,omitempty"`
	Position          *int             `json:"position,omitempty"`
	MatchesWithRating *int             `json:"matchesWithRating,omitempty"`
	LastName          *string          `json:"lastName,omitempty"`
	FullName          *string          `json:"fullName,omitempty"`
	Stats             *AwayPlayerStats `json:"stats,omitempty"`
}

type AwayPlayerStats struct {
	Goals              *int `json:"goals,omitempty"`
	GoalAssist         *int `json:"goalAssist,omitempty"`
	OntargetScoringAtt *int `json:"ontargetScoringAtt,omitempty"`
	Motm               *int `json:"motm,omitempty"`
	GamesPlayed        *int `json:"gamesPlayed,omitempty"`
	MinsPlayed         *int `json:"minsPlayed,omitempty"`
	MinsPlayedGoal     *int `json:"minsPlayedGoal,omitempty"`
	ExpectedGoals      *int `json:"expectedGoals,omitempty"`
	PlayerRating       *int `json:"playerRating,omitempty"`
}

type ShotmapClass struct {
	Shots   []*Shot         `json:"shots,omitempty"`
	Periods *ShotmapPeriods `json:"Periods,omitempty"`
}

type ShotmapPeriods struct {
	All        []*Shot `json:"All,omitempty"`
	FirstHalf  []*Shot `json:"FirstHalf,omitempty"`
	SecondHalf []*Shot `json:"SecondHalf,omitempty"`
}

type ContentStats struct {
	Periods *StatsPeriods `json:"Periods,omitempty"`
}

type StatsPeriods struct {
	All        *All `json:"All,omitempty"`
	FirstHalf  *All `json:"FirstHalf,omitempty"`
	SecondHalf *All `json:"SecondHalf,omitempty"`
}

type All struct {
	Stats      []*AllStat  `json:"stats,omitempty"`
	TeamColors *TeamColors `json:"teamColors,omitempty"`
}

type AllStat struct {
	Title       *string       `json:"title,omitempty"`
	Key         *string       `json:"key,omitempty"`
	Stats       []interface{} `json:"stats,omitempty"`
	Type        *StatType     `json:"type,omitempty"`
	Highlighted *Side         `json:"highlighted,omitempty"`
}

type StatType string

const (
	Graph         StatType = "graph"
	Text          StatType = "text"
	TitleStatType StatType = "title"
)

type TeamColors struct {
	DarkMode      *Mode `json:"darkMode,omitempty"`
	LightMode     *Mode `json:"lightMode,omitempty"`
	FontDarkMode  *Mode `json:"fontDarkMode,omitempty"`
	FontLightMode *Mode `json:"fontLightMode,omitempty"`
}

type Mode struct {
	Home *string `json:"home,omitempty"`
	Away *string `json:"away,omitempty"`
}

type Superlive struct {
	SuperLiveUrl  *string `json:"superLiveUrl,omitempty"`
	ShowSuperLive *bool   `json:"showSuperLive,omitempty"`
}

type Table struct {
	Url                  *string `json:"url,omitempty"`
	Teams                []int   `json:"teams,omitempty"`
	TournamentNameForUrl *string `json:"tournamentNameForUrl,omitempty"`
	ParentLeagueId       *int    `json:"parentLeagueId,omitempty"`
	ParentLeagueName     *string `json:"parentLeagueName,omitempty"`
	IsCurrentSeason      *bool   `json:"isCurrentSeason,omitempty"`
	ParentLeagueSeason   *string `json:"parentLeagueSeason,omitempty"`
	CountryCode          *string `json:"countryCode,omitempty"`
}

type General struct {
	MatchId                   *string          `json:"matchId,omitempty"`
	MatchName                 *string          `json:"matchName,omitempty"`
	MatchRound                *string          `json:"matchRound,omitempty"`
	TeamColors                *TeamColors      `json:"teamColors,omitempty"`
	LeagueId                  *int             `json:"leagueId,omitempty"`
	LeagueName                *string          `json:"leagueName,omitempty"`
	LeagueRoundName           *string          `json:"leagueRoundName,omitempty"`
	ParentLeagueId            *int             `json:"parentLeagueId,omitempty"`
	CountryCode               *string          `json:"countryCode,omitempty"`
	ParentLeagueName          *string          `json:"parentLeagueName,omitempty"`
	ParentLeagueSeason        *string          `json:"parentLeagueSeason,omitempty"`
	ParentLeagueTopScorerLink *string          `json:"parentLeagueTopScorerLink,omitempty"`
	ParentLeagueTournamentId  *int             `json:"parentLeagueTournamentId,omitempty"`
	HomeTeam                  *GeneralAwayTeam `json:"homeTeam,omitempty"`
	AwayTeam                  *GeneralAwayTeam `json:"awayTeam,omitempty"`
	CoverageLevel             *string          `json:"coverageLevel,omitempty"`
	MatchTimeUTC              *string          `json:"matchTimeUTC,omitempty"`
	MatchTimeUTCDate          *time.Time       `json:"matchTimeUTCDate,omitempty"`
	Started                   *bool            `json:"started,omitempty"`
	Finished                  *bool            `json:"finished,omitempty"`
}

type GeneralAwayTeam struct {
	Name *string `json:"name,omitempty"`
	ID   *int    `json:"id,omitempty"`
}

type Header struct {
	Teams  []*Team      `json:"teams,omitempty"`
	Status *Status      `json:"status,omitempty"`
	Events *interface{} `json:"events,omitempty"`
}

type Team struct {
	Name     *string      `json:"name,omitempty"`
	ID       *int         `json:"id,omitempty"`
	Score    *int         `json:"score,omitempty"`
	ImageUrl *string      `json:"imageUrl,omitempty"`
	PageUrl  *string      `json:"pageUrl,omitempty"`
	FifaRank *interface{} `json:"fifaRank,omitempty"`
}

type SEO struct {
	Path             *string             `json:"path,omitempty"`
	EventJSONLD      *EventJSONLD        `json:"eventJSONLD,omitempty"`
	BreadcrumbJSONLD []*BreadcrumbJSONLD `json:"breadcrumbJSONLD,omitempty"`
	FAQJSONLD        *FAQJSONLD          `json:"faqJSONLD,omitempty"`
}

type BreadcrumbJSONLD struct {
	Context         *string            `json:"@context,omitempty"`
	Type            *string            `json:"@type,omitempty"`
	ItemListElement []*ItemListElement `json:"itemListElement,omitempty"`
}

type ItemListElement struct {
	Type     *Type   `json:"@type,omitempty"`
	Position *int    `json:"position,omitempty"`
	Name     *string `json:"name,omitempty"`
	Item     *string `json:"item,omitempty"`
}

type Type string

const (
	ListItem Type = "ListItem"
)

type EventJSONLD struct {
	Context             *string              `json:"@context,omitempty"`
	Type                *string              `json:"@type,omitempty"`
	Sport               *string              `json:"sport,omitempty"`
	HomeTeam            *EventJSONLDAwayTeam `json:"homeTeam,omitempty"`
	AwayTeam            *EventJSONLDAwayTeam `json:"awayTeam,omitempty"`
	Name                *string              `json:"name,omitempty"`
	Description         *string              `json:"description,omitempty"`
	StartDate           *time.Time           `json:"startDate,omitempty"`
	EndDate             *time.Time           `json:"endDate,omitempty"`
	EventStatus         *string              `json:"eventStatus,omitempty"`
	EventAttendanceMode *string              `json:"eventAttendanceMode,omitempty"`
	Location            *Location            `json:"location,omitempty"`
	Image               []string             `json:"image,omitempty"`
	Organizer           *Organizer           `json:"organizer,omitempty"`
	Offers              *Offers              `json:"offers,omitempty"`
	Performer           []*Organizer         `json:"performer,omitempty"`
}

type EventJSONLDAwayTeam struct {
	Context  *string     `json:"@context,omitempty"`
	Type     *string     `json:"@type,omitempty"`
	Name     *string     `json:"name,omitempty"`
	Sport    *string     `json:"sport,omitempty"`
	Logo     *string     `json:"logo,omitempty"`
	URL      *string     `json:"url,omitempty"`
	Location interface{} `json:"location,omitempty"`
	MemberOf interface{} `json:"memberOf,omitempty"`
}

type Location struct {
	Type *string `json:"@type,omitempty"`
	URL  *string `json:"url,omitempty"`
}

type Offers struct {
	Type          *string    `json:"@type,omitempty"`
	URL           *string    `json:"url,omitempty"`
	Availability  *string    `json:"availability,omitempty"`
	Price         *string    `json:"price,omitempty"`
	PriceCurrency *string    `json:"priceCurrency,omitempty"`
	ValidFrom     *time.Time `json:"validFrom,omitempty"`
}

type Organizer struct {
	Type *string `json:"@type,omitempty"`
	Name *string `json:"name,omitempty"`
	URL  *string `json:"url,omitempty"`
	Logo *string `json:"logo,omitempty"`
}

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
