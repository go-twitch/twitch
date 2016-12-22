package twitch

import "time"

type User struct {
	DisplayName string    `json:"display_name"`
	ID          int64     `json:"_id,string"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Bio         string    `json:"bio"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Logo        *string   `json:"logo"`
}

type UserSelf struct {
	DisplayName      string    `json:"display_name"`
	ID               int64     `json:"_id"`
	Name             string    `json:"name"`
	Type             string    `json:"type"`
	Bio              string    `json:"bio"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	Logo             *string   `json:"logo"`
	Email            *string   `json:"email"`
	EmailVerified    *bool     `json:"email_verified"`
	Partnered        bool      `json:"partnered"`
	TwitterConnected bool      `json:"twitter_connected"`
	Notifications    struct {
		Push  bool `json:"push"`
		Email bool `json:"email"`
	} `json:"notifications"`
}

type UserSubscription struct {
	ID        string    `json:"_id"`
	Channel   Channel   `json:"channel"`
	CreatedAt time.Time `json:"created_at"`
}

type UserFollows struct {
	Total   int          `json:"_total"`
	Follows []UserFollow `json:"follows"`
}

type UserFollow struct {
	CreatedAt     time.Time `json:"created_at"`
	Notifications bool      `json:"notifications"`
	Channel       Channel   `json:"channel"`
}

type UserBlocks struct {
	Total  int         `json:"_total"`
	Blocks []UserBlock `json:"blocks"`
}

type UserBlock struct {
	ID        int64     `json:"_id"`
	CreatedAt time.Time `json:"created_at"`
	User      User      `json:"user"`
}
