package team

import (
	"encoding/json"
	"testing"

	"github.com/aws/smithy-go/ptr"
	"github.com/stretchr/testify/require"
)

func TestUnmarshalParticipant(t *testing.T) {
	data := []byte(`{
		"id": 530859,
		"name": "Leroy Sané",
		"rank": 1,
		"goals": null,
		"assists": null,
		"rating": null,
		"positionId": null,
		"ccode": "GER",
		"cname": null,
		"teamId": 9823,
		"teamName": "Bayern München",
		"showRole": null,
		"showCountryFlag": null,
		"showTeamFlag": true,
		"excludeFromRanking": false,
		"value": "8.18"
	}`)
	p := &Participant{}
	err := json.Unmarshal(data, p)
	require.NoError(t, err)
	require.Equal(t, ptr.Float32(8.18), p.Value)
}
