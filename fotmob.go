package fotmob

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"

	"github.com/mheers/go-fotmob/types/allleagues"
	"github.com/mheers/go-fotmob/types/league"
	"github.com/mheers/go-fotmob/types/matchdetails"
	"github.com/mheers/go-fotmob/types/matches"
	"github.com/mheers/go-fotmob/types/player"
	"github.com/mheers/go-fotmob/types/team"
	"github.com/mheers/go-fotmob/types/teamseasonstats"
	"github.com/mheers/go-fotmob/types/transfers"
	"github.com/mheers/go-fotmob/types/worldnews"
)

const baseUrl = "https://www.fotmob.com/api/"

type Fotmob struct {
	MatchesUrl          string
	LeaguesUrl          string
	AllLeaguesUrl       string
	TeamsUrl            string
	PlayerUrl           string
	TeamsSeasonStatsUrl string
	MatchDetailsUrl     string
	SearchUrl           string
	TransfersUrl        string
	WorldNewsUrl        string
	DefaultTimeZone     string
	DefaultLang         string
	Map                 map[interface{}]interface{}
}

func NewFotmob() *Fotmob {
	return &Fotmob{
		MatchesUrl:          baseUrl + "matches",
		LeaguesUrl:          baseUrl + "leagues",
		AllLeaguesUrl:       baseUrl + "allLeagues",
		TeamsUrl:            baseUrl + "teams",
		TeamsSeasonStatsUrl: baseUrl + "teamseasonstats",
		PlayerUrl:           baseUrl + "playerData",
		MatchDetailsUrl:     baseUrl + "matchDetails",
		SearchUrl:           baseUrl + "searchapi",
		TransfersUrl:        baseUrl + "transfers",
		WorldNewsUrl:        baseUrl + "worldnews",
		DefaultTimeZone:     "Europe/Berlin",
		DefaultLang:         "en",
	}
}

func (f *Fotmob) CheckDate(date string) ([]string, error) {
	re := regexp.MustCompile(`(20\d{2})(\d{2})(\d{2})`)
	return re.FindStringSubmatch(date), nil
}

func (f *Fotmob) SafeTypeCastFetch(url string, dst any) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// // write the response body to a file
	// data, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	return err
	// }
	// if err := os.WriteFile("response.json", data, 0644); err != nil {
	// 	return err
	// }

	// if err := json.NewDecoder(res.Body).Decode(&dst); err != nil {
	// 	return err
	// }

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, &dst); err != nil {
		return err
	}

	return nil
}

func (f *Fotmob) GetMatchesByDate(date string) (*matches.Matches, error) {
	if _, err := f.CheckDate(date); err == nil {
		url := fmt.Sprintf("%s?date=%s", f.MatchesUrl, date)
		m := &matches.Matches{}
		if err := f.SafeTypeCastFetch(url, m); err != nil {
			return nil, err
		}

		return m, nil
	}
	return nil, nil
}

// GetLeague returns a league by id; tab, leagueType and timeZone are optional
func (f *Fotmob) GetLeague(id int, tab string, leagueType string, timeZone string) (*league.League, error) {
	if tab == "" {
		tab = "overview"
	}

	if leagueType == "" {
		leagueType = "league"
	}

	if timeZone == "" {
		timeZone = f.DefaultTimeZone
	}

	url := fmt.Sprintf("%s?id=%d&tab=%s&type=%s&timeZone=%s", f.LeaguesUrl, id, tab, leagueType, timeZone)
	l := &league.League{}
	if err := f.SafeTypeCastFetch(url, l); err != nil {
		return nil, err
	}

	return l, nil
}

func (f *Fotmob) GetAllLeagues() (*allleagues.AllLeagues, error) {
	url := f.AllLeaguesUrl
	l := &allleagues.AllLeagues{}
	if err := f.SafeTypeCastFetch(url, l); err != nil {
		return nil, err
	}

	return l, nil
}

// GetTeam returns a team by id; tab, teamType and timeZone are optional
func (f *Fotmob) GetTeam(id int, tab string, teamType string, timeZone string) (*team.Team, error) {
	if tab == "" {
		tab = "overview"
	}

	if teamType == "" {
		teamType = "league"
	}

	if timeZone == "" {
		timeZone = f.DefaultTimeZone
	}

	url := fmt.Sprintf("%s?id=%d&tab=%s&type=%s&timeZone=%s", f.TeamsUrl, id, tab, teamType, timeZone)
	l := &team.Team{}

	// TODO: fix: json: cannot unmarshal string into Go struct field Player.stats.players.participant of type float32
	if err := f.SafeTypeCastFetch(url, l); err != nil {
		return nil, err
	}

	return l, nil
}

func (f *Fotmob) GetTeamSeasonStats(teamId int, seasonId int) (*teamseasonstats.TeamSeasonStats, error) {
	url := fmt.Sprintf("%s?teamId=%d&tournamentId=%d", f.TeamsSeasonStatsUrl, teamId, seasonId)
	l := &teamseasonstats.TeamSeasonStats{}
	if err := f.SafeTypeCastFetch(url, l); err != nil {
		return nil, err
	}

	return l, nil
}

func (f *Fotmob) GetPlayer(id int) (*player.Player, error) {
	url := fmt.Sprintf("%s?id=%d", f.PlayerUrl, id)
	l := &player.Player{}

	// TODO: fix: parsing time "2023-08-12T00:00:00" as "2006-01-02T15:04:05Z07:00": cannot parse "" as "Z07:00"
	if err := f.SafeTypeCastFetch(url, l); err != nil {
		return nil, err
	}

	return l, nil
}

func (f *Fotmob) GetMatchDetails(matchId int) (*matchdetails.MatchDetails, error) {
	url := fmt.Sprintf("%s?matchId=%d", f.MatchDetailsUrl, matchId)
	l := &matchdetails.MatchDetails{}

	// TODO: fix: Time.UnmarshalJSON: input is not a JSON string
	if err := f.SafeTypeCastFetch(url, l); err != nil {
		return nil, err
	}

	return l, nil
}

func (f *Fotmob) GetWorldNews(page int, lang string) ([]*worldnews.WorldNews, error) {
	if page == 0 {
		page = 1
	}

	if lang == "" {
		lang = f.DefaultLang
	}

	url := fmt.Sprintf("%s?page=%d&lang=%s", f.WorldNewsUrl, page, lang)
	l := []*worldnews.WorldNews{}
	if err := f.SafeTypeCastFetch(url, l); err != nil {
		return nil, err
	}

	return l, nil
}

func (f *Fotmob) GetTransfers(page int, lang string) (*transfers.Transfers, error) {
	if page == 0 {
		page = 1
	}

	if lang == "" {
		lang = f.DefaultLang
	}

	url := fmt.Sprintf("%s?page=%d&lang=%s", f.TransfersUrl, page, lang)
	l := &transfers.Transfers{}
	if err := f.SafeTypeCastFetch(url, l); err != nil {
		return nil, err
	}

	return l, nil
}
