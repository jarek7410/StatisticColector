package dbStats

import "errors"

var (
	errNoUser = errors.New("user not found in database")
)

type Repository interface {
	Migrate() error

	CreateUser(User) (*User, error)
	GetUsers() ([]User, error)
	GetUser(uint) (*User, error)
	KillUser(User) error

	CrateStat(user User, stat Stat) (*Stat, error)
	GetStats(User) ([]Stat, error)
	DeleteStat(Stat) error
}
