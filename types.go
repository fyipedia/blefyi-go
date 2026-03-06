package blefyi

// SearchResult represents the API search response.
type SearchResult struct {
	Query   string       `json:"query"`
	Results []SearchItem `json:"results"`
	Total   int          `json:"total"`
}

// SearchItem represents a single search result item.
type SearchItem struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
	Type string `json:"type"`
	URL  string `json:"url"`
}

// ChipDetail represents a BLE chip.
type ChipDetail struct {
	Name         string `json:"name"`
	Slug         string `json:"slug"`
	Description  string `json:"description"`
	Manufacturer string `json:"manufacturer"`
	URL          string `json:"url"`
}

// ProfileDetail represents a BLE profile.
type ProfileDetail struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

// VersionDetail represents a BLE version.
type VersionDetail struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

// BeaconDetail represents a BLE beacon type.
type BeaconDetail struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

// UsecaseDetail represents a BLE use case.
type UsecaseDetail struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

// ManufacturerDetail represents a BLE manufacturer.
type ManufacturerDetail struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

// GlossaryTerm represents a glossary term.
type GlossaryTerm struct {
	Term       string `json:"term"`
	Slug       string `json:"slug"`
	Definition string `json:"definition"`
	URL        string `json:"url"`
}

// CompareResult represents a comparison between two BLE chips.
type CompareResult struct {
	ChipA interface{} `json:"chip_a"`
	ChipB interface{} `json:"chip_b"`
	URL   string      `json:"url"`
}
