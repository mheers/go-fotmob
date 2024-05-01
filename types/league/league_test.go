package league

import (
	"encoding/json"
	"testing"
)

const topPlayersJSON = `{"byRating":{"players":[{"id":909907,"name":"Cameron Puertas","teamId":7978,"teamName":"Union Saint-Gilloise","rating":7.85,"teamColors":{"darkMode":"#F6E000","lightMode":"#daa000","fontDarkMode":"rgba(29, 29, 29, 1.0)","fontLightMode":"rgba(255, 255, 255, 1.0)"}},{"id":482397,"name":"Julien De Sart","teamId":9991,"teamName":"KAA Gent","rating":7.76,"teamColors":{"darkMode":"#0062ce","lightMode":"#004898","fontDarkMode":"rgba(255, 255, 255, 1.0)","fontLightMode":"rgba(255, 255, 255, 1.0)"}},{"id":641344,"name":"Anders Dreyer","teamId":8635,"teamName":"RSC Anderlecht","rating":7.69,"teamColors":{"darkMode":"#524d9d","lightMode":"#3F2B72","fontDarkMode":"rgba(255, 255, 255, 1.0)","fontLightMode":"rgba(255, 255, 255, 1.0)"}}]},"byGoals":{"players":[{"id":820477,"name":"Kevin Denkey","teamId":9984,"teamName":"Cercle Brugge","goals":25,"teamColors":{"darkMode":"#006a33","lightMode":"#00521f","fontDarkMode":"rgba(255, 255, 255, 1.0)","fontLightMode":"rgba(255, 255, 255, 1.0)"}},{"id":1134206,"name":"Mohamed Amoura","teamId":7978,"teamName":"Union Saint-Gilloise","goals":18,"teamColors":{"darkMode":"#F6E000","lightMode":"#daa000","fontDarkMode":"rgba(29, 29, 29, 1.0)","fontLightMode":"rgba(255, 255, 255, 1.0)"}},{"id":641344,"name":"Anders Dreyer","teamId":8635,"teamName":"RSC Anderlecht","goals":18,"teamColors":{"darkMode":"#524d9d","lightMode":"#3F2B72","fontDarkMode":"rgba(255, 255, 255, 1.0)","fontLightMode":"rgba(255, 255, 255, 1.0)"}}]},"byAssists":{"players":[{"id":909907,"name":"Cameron Puertas","teamId":7978,"teamName":"Union Saint-Gilloise","assists":15,"teamColors":{"darkMode":"#F6E000","lightMode":"#daa000","fontDarkMode":"rgba(29, 29, 29, 1.0)","fontLightMode":"rgba(255, 255, 255, 1.0)"}},{"id":641344,"name":"Anders Dreyer","teamId":8635,"teamName":"RSC Anderlecht","assists":9,"teamColors":{"darkMode":"#524d9d","lightMode":"#3F2B72","fontDarkMode":"rgba(255, 255, 255, 1.0)","fontLightMode":"rgba(255, 255, 255, 1.0)"}},{"id":1134871,"name":"Maxim De Cuyper","teamId":8342,"teamName":"Club Brugge","assists":8,"teamColors":{"darkMode":"#006ad5","lightMode":"#006ad5","fontDarkMode":"rgba(255, 255, 255, 1.0)","fontLightMode":"rgba(255, 255, 255, 1.0)"}}]},"seeAllUrl":"stats"}`

func TestUnmarshalTopPlayers(t *testing.T) {
	// Test the unmarshal function for the top players
	t.Run("TestUnmarshalTopPlayers", func(t *testing.T) {
		// Create a new top players object
		topPlayers := &TopPlayers{}
		// Unmarshal the json data into the top players object
		err := json.Unmarshal([]byte(topPlayersJSON), &topPlayers)
		// Check if there was an error unmarshalling the json data
		if err != nil {
			// Log the error
			t.Errorf("Error unmarshalling json data: %v", err)
		}
		// Check if the top players object is nil
		if topPlayers == nil {
			// Log an error
			t.Errorf("Top players object is nil")
		}
		// Check if the top players object is empty
		if len(topPlayers.ByRating.ByPlayers) == 0 {
			// Log an error
			t.Errorf("Top players object is empty")
		}
	})
}
