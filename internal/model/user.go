package model

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

type User struct {
	modelData
	Email          string `json:"email"`
	Password       string `json:"password"`
	NewPassword    string `json:"newpass,omitempty"`
	NewPassCompare string `json:"newpasscompare,omitempty"`
	Token          string `json:"token,omitempty"`
}
type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// newUser model user
func newUser(data modelData, user *User) *User {
	if user == nil {
		user = &User{}
	}
	user.modelData = data
	return user
}

// LoginUser
func (u *User) Login() (*User, error) {
	u.logger.Info().Msg("Login")
	result := &User{}
	if u.Email != "" {
		password := GetMD5Hash(u.Password)

		userM, err := u.repository.GetHMGet("usersm", u.Email)
		if err != nil {
			u.logger.Debug().Err(err).Msg("Сannot find user email")
			return nil, err
		}

		if userM == nil {
			newUser := map[string]interface{}{
				u.Email: password,
			}

			err = u.repository.HMSet("usersm", newUser)
			if err != nil {
				return nil, err
			}
		}

		if userM[0] != nil {
			result.Email = u.Email
			result.Password = ""
		}
	}

	return result, nil
}

// Registration
func (u *User) Registration() (*User, error) {
	u.logger.Info().Msg("Registration")

	userIsset, err := u.repository.GetHMGet("usersm", u.Email)
	if err != nil {
		return nil, err
	}

	if userIsset[0] != nil {
		u.logger.Debug().Str("user m", fmt.Sprint(userIsset)).Msg("User isset")
		return nil, err
	}
	newUser := map[string]interface{}{
		u.Email: u.Password,
	}

	err = u.repository.HMSet("usersm", newUser)
	if err != nil {
		return nil, err
	}
	return u, nil

}

// Changepass
func (u *User) Changepass() (*User, error) {
	result := &User{}
	u.logger.Info().Msg("Changepass")

	userM, err := u.repository.GetHMGet("usersm", u.Email)
	if err != nil {
		u.logger.Debug().Err(err).Msg("Сannot find user email")
		return nil, err
	}

	if userM == nil {
		return nil, fmt.Errorf("user not find")
	}

	userChangePass := u.repository.HSet("usersm", u.Email, string(u.NewPassword))

	u.logger.Debug().Bool("change pass", userChangePass).Msg("Changepass")

	return result, nil
}
