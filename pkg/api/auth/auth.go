package auth

import (
	root "github.com/DarioCalovic/secretify"
)

// Authenticate tries to authenticate the user provided by username and password
func (a Auth) Authenticate(username, password string) (root.AuthToken, error) {
	u, err := a.rep.FindByUsername(a.db, username)
	if err != nil {
		return root.AuthToken{}, root.ErrUnauthorized
	}
	// TODO : check password
	return a.tokenize(u)
}

func (a Auth) tokenize(u root.User) (root.AuthToken, error) {
	if !u.Active {
		return root.AuthToken{}, root.ErrUnauthorized
	}

	token, err := a.tg.GenerateToken(u)
	if err != nil {
		return root.AuthToken{}, root.ErrUnauthorized
	}

	u.UpdateLastLogin(a.sec.Token(token))

	if err := a.rep.Update(a.db, u); err != nil {
		return root.AuthToken{}, err
	}
	return root.AuthToken{Token: token, RefreshToken: u.Token}, nil
}

// Me returns info about currently logged user
func (a Auth) Me() (root.AuthUser, error) {
	// u := a.rbac.User(c)
	u := root.AuthUser{}
	return u, nil
}
