package tools_test

import (
	"testing"

	"github.com/aixoio/speed-notes/server/tools"
)

func TestIsUsernameValid(t *testing.T) {
	if !tools.IsUsernameValid("Validusername0123465798") {
		t.Fail()
	}

	if tools.IsUsernameValid("IN-VALID_USERNAME@$#%^&*") {
		t.Fail()
	}

}
