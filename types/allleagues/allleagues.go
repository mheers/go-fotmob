package allleagues

type AllLeagues struct {
	Favourite     interface{}     `json:"favourite"`
	Popular       []League        `json:"popular"`
	International []International `json:"international"`
	Countries     []Country       `json:"countries"`
	UserSettings  interface{}     `json:"userSettings"`
}

type Country struct {
	Ccode         string   `json:"ccode"`
	Name          string   `json:"name"`
	Leagues       []League `json:"leagues"`
	LocalizedName string   `json:"localizedName"`
}

type League struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	LocalizedName string `json:"localizedName"`
	PageURL       string `json:"pageUrl"`
}

type International struct {
	Ccode         string   `json:"ccode"`
	Name          string   `json:"name"`
	Leagues       []League `json:"leagues"`
	LocalizedName string   `json:"localizedName"`
}

type SourceStr string

const (
	FotMob SourceStr = "FotMob"
)
