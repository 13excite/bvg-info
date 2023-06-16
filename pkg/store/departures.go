package store

// Hardcode nearby station (move to config????)
const (
	Sudostallee_Kongisheide string = "sudost_konigsheide"
	Schnellerstr_135               = "schnel_str_135"
	S_Schöneweide                  = "s_schoneweide"
	S_Schöneweide_sterndamm        = "s_schoneweide_sterndamm"
	S_Schöneweide_Vorplatz         = "s_schoneweide_sterndamm_vorplatz"
)

// NearbyDepartures returns map of nearby public transport stations
func NearbyDepartures() map[string]Stop {
	return map[string]Stop{
		Sudostallee_Kongisheide: {
			ID:   "733612",
			Name: "Südostallee/Königsheide",
			Location: Location{
				ID: "733612",
			},
			Products: Products{
				Bus: true,
			},
		},
		Schnellerstr_135: {
			ID:   "362510",
			Name: "Schnellerstr. 135",
			Location: Location{
				ID: "362510",
			},
			Products: Products{
				Bus: true,
			},
		},
		S_Schöneweide: {
			ID:   "733559",
			Name: "S Schöneweide",
			Location: Location{
				ID: "733559",
			},
			Products: Products{
				Bus:      true,
				Tram:     true,
				Suburban: true,
				Regional: true,
			},
		},
		S_Schöneweide_sterndamm: {
			ID:   "733587",
			Name: "Schöneweide (S)/Sterndamm",
			Location: Location{
				ID: "733587",
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
