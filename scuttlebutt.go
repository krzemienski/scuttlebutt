package scuttlebutt

import (
	"strings"
	"time"
)

// Config represents the configuration used by Scuttlebutt.
type Config struct {
	AppKey    string     `toml:"app_key"`
	AppSecret string     `toml:"app_secret"`
	Interval  Duration   `toml:"interval"`
	Accounts  []*Account `toml:"account"`
}

// Account represents a Twitter account that tweets occassional trending repos.
type Account struct {
	Username string `toml:"username"`
	Language string `toml:"language"`
	Key      string `toml:"key"`
	Secret   string `toml:"secret"`
}

// Notify sends a tweet for an account about a given repository.
func (a *Account) Notify(r *Repository) error {
	// TODO(benbjohnson): Create client.
	// TODO(benbjohnson): Construct tweet text.
	// TODO(benbjohnson): Send tweet.
	panic("NOT IMPLEMENTED: Account.Notify()")
	return nil
}

// AccountStatus represents status information about a given account.
type AccountStatus struct {
	NotifyTime time.Time `json:"notifyTime"`
}

// Repository represents a code repository.
type Repository struct {
	ID       string     `json:"id"`
	Language string     `json:"language"`
	Messages []*Message `json:"messages"`
}

// Message represents a message associated with a project and language.
type Message struct {
	ID string `json:"id"`
}

// Duration is a helper type for unmarshaling durations in TOML.
type Duration time.Duration

func (d *Duration) UnmarshalText(text []byte) error {
	duration, err := time.ParseDuration(string(text))
	if err != nil {
		return err
	}
	*d = Duration(duration)
	return nil
}

func splitRepositoryID(id string) (string, string, string) {
	s := strings.Split(id, "/")
	if len(s) != 3 {
		return "", "", ""
	}
	return s[0], s[1], s[2]
}
