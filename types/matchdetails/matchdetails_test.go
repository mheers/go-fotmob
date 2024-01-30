package matchdetails

import (
	"encoding/json"
	"testing"

	"github.com/aws/smithy-go/ptr"
	"github.com/stretchr/testify/require"
)

func TestUnmarshalEvent(t *testing.T) {
	data := []byte(`{
		"reactKey": "4Goal530859undefinedfalse",
		"timeStr": 4,
		"type": "Goal",
		"time": 4,
		"overloadTime": null,
		"eventId": 10417041,
		"player": {
			"id": 530859,
			"name": "Leroy Sané",
			"profileUrl": "/players/530859/leroy-sane"
		},
		"profileUrl": "/players/530859/leroy-sane",
		"overloadTimeStr": false,
		"isHome": false,
		"ownGoal": null,
		"goalDescription": null,
		"goalDescriptionKey": null,
		"suffix": null,
		"suffixKey": null,
		"isPenaltyShootoutEvent": false,
		"nameStr": "Leroy Sané",
		"firstName": "Leroy",
		"lastName": "Sané",
		"fullName": "Leroy Sané",
		"playerId": 530859,
		"newScore": [
			0,
			1
		],
		"penShootoutScore": null,
		"shotmapEvent": {
			"id": 2578492825,
			"eventType": "Goal",
			"teamId": 9823,
			"playerId": 530859,
			"playerName": "Leroy Sané",
			"x": 92.4,
			"y": 30.232,
			"min": 4,
			"minAdded": null,
			"isBlocked": false,
			"isOnTarget": true,
			"blockedX": null,
			"blockedY": null,
			"goalCrossedY": 36.36375,
			"goalCrossedZ": 0.0385263156,
			"expectedGoals": 0.6170995235443115,
			"expectedGoalsOnTarget": 0.8651,
			"shotType": "RightFoot",
			"situation": "FastBreak",
			"period": "FirstHalf",
			"isOwnGoal": false,
			"onGoalShot": {
				"x": 0.37466931216931143,
				"y": 0.010192146984126984,
				"zoomRatio": 1
			},
			"isSavedOffLine": false,
			"firstName": "Leroy",
			"lastName": "Sané",
			"fullName": "Leroy Sané",
			"teamColor": "#C60428"
		},
		"assistStr": "assist by Harry Kane",
		"assistProfileUrl": "/players/194165/harry-kane",
		"assistPlayerId": 194165,
		"assistKey": "assist_by",
		"assistInput": "Harry Kane"
	}`)

	event := &Event{}
	err := json.Unmarshal(data, event)
	require.NoError(t, err)
	require.Equal(t, ptr.String("4.000000"), event.TimeStr)
	require.Equal(t, ptr.String("Goal"), event.Type)
}

func TestUnmarshalEventTypes(t *testing.T) {
	data := []byte(`[
			"Goal",
			"Assist",
			"Card",
			"AddedTime",
			"Half",
			"Substitution",
			"Yellow",
			"Injuries",
			null,
			[]
		]`)
	eventTypes := &EventTypes{}
	err := json.Unmarshal(data, eventTypes)
	require.NoError(t, err)
	require.Equal(t, 8, len(*eventTypes))
}
