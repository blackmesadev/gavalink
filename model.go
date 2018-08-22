package gavalink

const (
	// TrackLoaded is a Tracks Type for a succesful single track load
	TrackLoaded = "TRACK_LOADED"
	// PlaylistLoaded is a Tracks Type for a succseful playlist load
	PlaylistLoaded = "PLAYLIST_LOADED"
	// SearchResult is a Tracks Type for a search containing many tracks
	SearchResult = "SEARCH_RESULT"
	// NoMatches is a Tracks Type for a query yielding no matches
	NoMatches = "NO_MATCHES"
	// LoadFailed is a Tracks Type for an internal Lavalink error
	LoadFailed = "LOAD_FAILED"
)

// Tracks contains data for a Lavalink Tracks response
type Tracks struct {
	// Type contains the type of response
	//
	// This will be one of TrackLoaded, PlaylistLoaded, SearchResult,
	// NoMatches, or LoadFailed
	Type         string        `json:"loadType"`
	PlaylistInfo *PlaylistInfo `json:"playlistInfo"`
	Tracks       []Track       `json:"tracks"`
}

// PlaylistInfo contains information about a loaded playlist
type PlaylistInfo struct {
	// Name is the friendly of the playlist
	Name string `json:"name"`
	// SelectedTrack is the index of the track that loaded the playlist,
	// if one is present.
	SelectedTrack int `json:"selectedTrack"`
}

// Track contains information about a loaded track
type Track struct {
	// Data contains the base64 encoded Lavaplayer track
	Data string    `json:"track"`
	Info TrackInfo `json:"info"`
}

// TrackInfo contains more data about a loaded track
type TrackInfo struct {
	Identifier string `json:"identifier"`
	Title      string `json:"title"`
	Author     string `json:"author"`
	URI        string `json:"uri"`
	Seekable   bool   `json:"isSeekable"`
	Stream     bool   `json:"isStream"`
	Length     int    `json:"length"`
	Position   int    `json:"position"`
}
