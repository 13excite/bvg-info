package store

import "time"

// nearby stations models
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

// Stops departures models
type StopDepartures struct {
	TripID string `json:"tripId"`
	Stop   struct {
		Type     string `json:"type"`
		ID       string `json:"id"`
		Name     string `json:"name"`
		Location struct {
			Type      string  `json:"type"`
			ID        string  `json:"id"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"location"`
		Products struct {
			Suburban bool `json:"suburban"`
			Subway   bool `json:"subway"`
			Tram     bool `json:"tram"`
			Bus      bool `json:"bus"`
			Ferry    bool `json:"ferry"`
			Express  bool `json:"express"`
			Regional bool `json:"regional"`
		} `json:"products"`
		StationDHID string `json:"stationDHID"`
	} `json:"stop"`
	When            time.Time   `json:"when"`
	PlannedWhen     time.Time   `json:"plannedWhen"`
	Delay           int         `json:"delay"`
	Platform        string      `json:"platform"`
	PlannedPlatform string      `json:"plannedPlatform"`
	PrognosisType   string      `json:"prognosisType"`
	Direction       string      `json:"direction"`
	Provenance      interface{} `json:"provenance"`
	Line            struct {
		Type        string `json:"type"`
		ID          string `json:"id"`
		FahrtNr     string `json:"fahrtNr"`
		Name        string `json:"name"`
		Public      bool   `json:"public"`
		AdminCode   string `json:"adminCode"`
		ProductName string `json:"productName"`
		Mode        string `json:"mode"`
		Product     string `json:"product"`
		Operator    struct {
			Type string `json:"type"`
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"operator"`
		Symbol  string `json:"symbol"`
		Nr      int    `json:"nr"`
		Metro   bool   `json:"metro"`
		Express bool   `json:"express"`
		Night   bool   `json:"night"`
		Color   struct {
			Fg string `json:"fg"`
			Bg string `json:"bg"`
		} `json:"color"`
	} `json:"line"`
	Remarks []struct {
		Type    string `json:"type"`
		Code    string `json:"code,omitempty"`
		Text    string `json:"text"`
		ID      string `json:"id,omitempty"`
		Summary string `json:"summary,omitempty"`
		Icon    struct {
			Type  string      `json:"type"`
			Title interface{} `json:"title"`
		} `json:"icon,omitempty"`
		Priority int `json:"priority,omitempty"`
		Products struct {
			Suburban bool `json:"suburban"`
			Subway   bool `json:"subway"`
			Tram     bool `json:"tram"`
			Bus      bool `json:"bus"`
			Ferry    bool `json:"ferry"`
			Express  bool `json:"express"`
			Regional bool `json:"regional"`
		} `json:"products,omitempty"`
		Company    string    `json:"company,omitempty"`
		Categories []int     `json:"categories,omitempty"`
		ValidFrom  time.Time `json:"validFrom,omitempty"`
		ValidUntil time.Time `json:"validUntil,omitempty"`
	} `json:"remarks"`
	Origin      interface{} `json:"origin"`
	Destination struct {
		Type     string `json:"type"`
		ID       string `json:"id"`
		Name     string `json:"name"`
		Location struct {
			Type      string  `json:"type"`
			ID        string  `json:"id"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"location"`
		Products struct {
			Suburban bool `json:"suburban"`
			Subway   bool `json:"subway"`
			Tram     bool `json:"tram"`
			Bus      bool `json:"bus"`
			Ferry    bool `json:"ferry"`
			Express  bool `json:"express"`
			Regional bool `json:"regional"`
		} `json:"products"`
		StationDHID string `json:"stationDHID"`
	} `json:"destination"`
	CurrentTripPosition struct {
		Type      string  `json:"type"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"currentTripPosition"`
}

type CachedStops struct {
	Name     string
	Departes []StopDepartures
}
