package player

import (
	"encoding/json"
	"testing"

	"github.com/aws/smithy-go/ptr"
	"github.com/stretchr/testify/require"
)

func TestUnmarshalTeamEntry(t *testing.T) {
	data := []byte(`{
		"participantId": 194165,
		"teamId": 9823,
		"team": "Bayern München",
		"teamGender": "male",
		"transferType": null,
		"startDate": "2023-08-12T00:00:00",
		"endDate": null,
		"active": true,
		"role": null,
		"appearances": "26",
		"goals": "27",
		"assists": "8",
		"hasUncertainData": false
	}`)

	teamEntry := &TeamEntry{}
	err := json.Unmarshal(data, teamEntry)
	require.NoError(t, err)
	require.Equal(t, ptr.Int(194165), teamEntry.ParticipantID)
}

func TestUnmarshalFallback(t *testing.T) {
	data := []byte(`[
        {
            "value": {
                "key": null,
                "fallback": "188 cm"
            },
            "title": "Height",
            "translationKey": "height_sentencecase"
        },
        {
            "value": {
                "key": null,
                "fallback": 9
            },
            "title": "Shirt",
            "translationKey": "shirt"
        }
	]`)
	pi := []*PlayerInformation{}
	err := json.Unmarshal(data, &pi)
	require.NoError(t, err)

	fb1 := pi[0].Value.Fallback
	require.Equal(t, ptr.String("188 cm"), fb1.String)

	fb2 := pi[1].Value.Fallback
	require.Equal(t, ptr.Float64(9), fb2.Number)
}
