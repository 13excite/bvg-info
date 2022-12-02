package bvv

const (
	Sudostallee_Kongisheide string = "sudost_konigsheide"
	Schnellerstr_135               = "schnel_str_135"
	S_Schöneweide                  = "s_schoneweide"
	S_Schöneweide_sterndamm        = "s_schoneweide_sterndamm"
	S_Schöneweide_Vorplatz         = "s_schoneweide_sterndamm_vorplatz"
)

type Stop struct {
	Type     string   `json:"type"`
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Location Location `json:"location"`
	Products Products `json:"products"`
}

type Location struct {
	Type      string  `json:"type"`
	ID        string  `json:"id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Products struct {
	Suburban bool `json:"suburban"`
	Subway   bool `json:"subway"`
	Tram     bool `json:"tram"`
	Bus      bool `json:"bus"`
	Ferry    bool `json:"ferry"`
	Express  bool `json:"express"`
	Regional bool `json:"regional"`
}

// NearbyDepartures returns map of nearby public transport stations
func NearbyDepartures() map[string]Stop {
	return map[string]Stop{
		Sudostallee_Kongisheide: {
			ID:   "900000194519",
			Name: "Südostallee/Königsheide",
			Location: Location{
				ID: "900194519",
			},
			Products: Products{
				Bus: true,
			},
		},
		Schnellerstr_135: {
			ID:   "900000192510",
			Name: "Schnellerstr. 135",
			Location: Location{
				ID: "900192510",
			},
			Products: Products{
				Bus: true,
			},
		},
		S_Schöneweide: {
			ID:   "900000192001",
			Name: "S Schöneweide",
			Location: Location{
				ID: "900192001",
			},
			Products: Products{
				Bus:      true,
				Tram:     true,
				Suburban: true,
				Regional: true,
			},
		},
		S_Schöneweide_sterndamm: {
			ID:   "900000194006",
			Name: "S Schöneweide/Sterndamm",
			Location: Location{
				ID: "900194006",
			},
			Products: Products{
				Bus:      true,
				Tram:     true,
				Suburban: true,
				Regional: true,
			},
		},
		S_Schöneweide_Vorplatz: {
			ID:   "900000192701",
			Name: "S Schöneweide [Vorplatz]",
			Location: Location{
				ID: "900192701",
			},
			Products: Products{
				Bus:      true,
				Tram:     true,
				Suburban: true,
				Regional: true,
			},
		},
	}
}
