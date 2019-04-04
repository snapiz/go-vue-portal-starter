package common

import "fmt"

// GetAvatarURL get avatar url from email hash
func GetAvatarURL(emailHash string) string {
	return fmt.Sprintf("https://www.gravatar.com/avatar/%s?d=identicon&s=68", emailHash)
}
