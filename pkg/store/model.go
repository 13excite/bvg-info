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

type Departures struct {
	Departures []StopDepartures `json:"departures"`
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
			NationalExpress bool `json:"nationalExpress"`
			National        bool `json:"national"`
			RegionalExpress bool `json:"regionalExpress"`
			Regional        bool `json:"regional"`
			Suburban        bool `json:"suburban"`
			Bus             bool `json:"bus"`
			Ferry           bool `json:"ferry"`
			Subway          bool `json:"subway"`
			Tram            bool `json:"tram"`
			Taxi            bool `json:"taxi"`
		} `json:"products"`
		Station struct {
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
				NationalExpress bool `json:"nationalExpress"`
				National        bool `json:"national"`
				RegionalExpress bool `json:"regionalExpress"`
				Regional        bool `json:"regional"`
				Suburban        bool `json:"suburban"`
				Bus             bool `json:"bus"`
				Ferry           bool `json:"ferry"`
				Subway          bool `json:"subway"`
				Tram            bool `json:"tram"`
				Taxi            bool `json:"taxi"`
			} `json:"products"`
		} `json:"station"`
	} `json:"stop"`
	When            time.Time   `json:"when"`
	PlannedWhen     time.Time   `json:"plannedWhen"`
	Delay           interface{} `json:"delay"`
	Platform        interface{} `json:"platform"`
	PlannedPlatform interface{} `json:"plannedPlatform"`
	PrognosisType   interface{} `json:"prognosisType"`
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
	} `json:"line"`
	Remarks     []interface{} `json:"remarks"`
	Origin      interface{}   `json:"origin"`
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
			NationalExpress bool `json:"nationalExpress"`
			National        bool `json:"national"`
			RegionalExpress bool `json:"regionalExpress"`
			Regional        bool `json:"regional"`
			Suburban        bool `json:"suburban"`
			Bus             bool `json:"bus"`
			Ferry           bool `json:"ferry"`
			Subway          bool `json:"subway"`
			Tram            bool `json:"tram"`
			Taxi            bool `json:"taxi"`
		} `json:"products"`
	} `json:"destination"`
	CurrentTripPosition struct {
		Type      string  `json:"type"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"currentTripPosition"`
}

type CachedStops struct {
	Name     string
	Departes []CachedStop
}

type CachedStop struct {
	ID          string
	Name        string
	Time        time.Time
	PlannedTime time.Time
	Direction   string
	LineName    string
	ProductName string
	Remarks     []string
}
