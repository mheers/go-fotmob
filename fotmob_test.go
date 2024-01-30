package fotmob

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetMatches(t *testing.T) {
	f := NewFotmob()
	matches, err := f.GetMatchesByDate("20190301")
	require.NoError(t, err)
	require.NotEmpty(t, matches)
	require.Len(t, matches.Leagues, 66)
}

func TestGetLeague(t *testing.T) {
	f := NewFotmob()
	league, err := f.GetLeague(54, "overview", "league", "Europe/Berlin") // 1. Bundesliga
	require.NoError(t, err)
	require.NotEmpty(t, league)
}

func TestGetAllLeagues(t *testing.T) {
	f := NewFotmob()
	leagues, err := f.GetAllLeagues()
	require.NoError(t, err)
	require.NotEmpty(t, leagues)
	require.GreaterOrEqual(t, len(leagues.Countries), 80)
	require.GreaterOrEqual(t, len(leagues.International), 1)
	require.GreaterOrEqual(t, len(leagues.Popular), 10)
}

func TestGetTeam(t *testing.T) {
	f := NewFotmob()
	team, err := f.GetTeam(9823, "overview", "league", "Europe/Berlin") // FC Bayern München
	require.NoError(t, err)
	require.NotEmpty(t, team)
}

func TestGetTeamSeasonStats(t *testing.T) {
	f := NewFotmob()
	teamSeasonStats, err := f.GetTeamSeasonStats(9823, 2023) // FC Bayern München
	require.NoError(t, err)
	require.NotEmpty(t, teamSeasonStats)
}

func TestGetPlayer(t *testing.T) {
	f := NewFotmob()
	player, err := f.GetPlayer(194165) // Harry Kane
	require.NoError(t, err)
	require.NotEmpty(t, player)
}

func TestGetMatchDetails(t *testing.T) {
	f := NewFotmob()
	matchDetails, err := f.GetMatchDetails(4221721) // Werder Bremen vs. FC Bayern München
	require.NoError(t, err)
	require.NotEmpty(t, matchDetails)
}

func TestGetWorldNews(t *testing.T) {
	f := NewFotmob()
	_, err := f.GetWorldNews(1, "en")
	require.NoError(t, err)
	// require.NotEmpty(t, worldNews)
}

func TestGetTransfers(t *testing.T) {
	f := NewFotmob()
	transfers, err := f.GetTransfers(1, "en")
	require.NoError(t, err)
	require.NotEmpty(t, transfers)
	require.GreaterOrEqual(t, len(transfers.Transfers), 1)
}
