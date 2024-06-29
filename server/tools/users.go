package tools

import "regexp"

func IsUsernameValid(un string) bool {
	return !regexp.MustCompile("[^A-z0-9]+").MatchString(un)
}
