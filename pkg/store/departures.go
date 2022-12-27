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
