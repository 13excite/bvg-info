package store

// Hardcode nearby station (move to config????)
const (
	Sudostallee_Kongisheide string = "sudost_konigsheide"
	Schnellerstr_135        string = "schnel_str_135"
	S_Schöneweide           string = "s_schoneweide"
	S_Schöneweide_sterndamm string = "s_schoneweide_sterndamm"
	S_Schöneweide_Vorplatz  string = "s_schoneweide_sterndamm_vorplatz"
)

// NearbyDepartures returns map of nearby public transport stations
func NearbyDepartures() map[string]Departure {
	return map[string]Departure{
		Sudostallee_Kongisheide: {
			ID:   "733612",
			Name: "Südostallee/Königsheide",
		},
		Schnellerstr_135: {
			ID:   "362510",
			Name: "Schnellerstr. 135",
		},
		S_Schöneweide: {
			ID:   "733559",
			Name: "S Schöneweide",
		},
		S_Schöneweide_sterndamm: {
			ID:   "733587",
			Name: "Schöneweide (S)/Sterndamm",
		},
	}
}
