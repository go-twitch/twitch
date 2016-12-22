package twitch

import "time"

type Channel struct {
	ID                  int64     `json:"_id"`
	BroadcasterLaungage string    `json:"broadcaster_laungage"`
	CreatedAt           time.Time `json:"created_at"`
	Name                string    `json:"name"`
}
