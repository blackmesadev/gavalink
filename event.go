package gavalink

// EventHandler defines events that Lavalink may send to a player
type EventHandler struct {
	OnTrackEnd       OnTrackEnd
	OnTrackException OnTrackException
	OnTrackStuck     OnTrackStuck
}

type OnTrackEnd func(*Player, string, string) error
type OnTrackException func(*Player, string, string) error
type OnTrackStuck func(*Player, string, int) error

// DummyEventHandler provides an empty event handler for users who
// wish to drop events outright. This is not recommended.
type DummyEventHandler struct{}

// OnTrackEnd is raised when a track ends
func (d DummyEventHandler) OnTrackEnd(player *Player, track string, reason string) error {
	return nil
}

// OnTrackException is raised when a track throws an exception
func (d DummyEventHandler) OnTrackException(player *Player, track string, reason string) error {
	return nil
}

// OnTrackStuck is raised when a track gets stuck
func (d DummyEventHandler) OnTrackStuck(player *Player, track string, threshold int) error {
	return nil
}
