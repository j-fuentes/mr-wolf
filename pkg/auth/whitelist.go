package auth

import (
	"fmt"

	tb "gopkg.in/tucnak/telebot.v2"
)

// Auth provides authentication.
type Auth struct {
	// whitelist is a list of allowed user IDs.
	whitelist []int
	// admin is the ID of the admin user.
	admin int
}

// NewAuth creates a new Auth.
func NewAuth(whitelist []int, admin int) (*Auth, error) {
	if len(whitelist) > 0 {
		return &Auth{
			whitelist: whitelist,
			admin:     admin,
		}, nil
	}

	return nil, fmt.Errorf("cannot create Auth with empty whitelist")
}

// UserAllowed returns true if the user is allowed.
func (a *Auth) UserAllowed(user *tb.User) bool {
	for _, id := range a.whitelist {
		if user.ID == id {
			return true
		}
	}

	return false
}

// AdminUser returns the admin user.
func (a *Auth) AdminUser() *tb.User {
	return &tb.User{
		ID: a.admin,
	}
}
