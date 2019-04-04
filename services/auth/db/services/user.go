package services

import (
	"time"

	"github.com/snapiz/go-vue-portal-starter/services/auth/db/models"
	common "github.com/snapiz/go-vue-portal-starter/common/go"
	"github.com/volatiletech/null"
)

// SetUserPassword set user password
func SetUserPassword(u *models.User, p string) error {
	hash, err := common.Hash(p)

	if err != nil {
		return err
	}

	u.Password = null.StringFrom(hash)
	u.TokenVersion = null.Int64From(time.Now().Unix())

	return err
}

// VerifyUserPassword verify user password
func VerifyUserPassword(u *models.User, p string) bool {
	ok, err := common.Verify(p, *u.Password.Ptr())

	return ok && err == nil
}
