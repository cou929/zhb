package command

import (
	"testing"
)

func TestValidateRequredArgs(t *testing.T) {
	func() {
		// valid case
		authToken := "test-token"
		org := "org"
		repo := "repo"
		got := validateRequredArgs(authToken, org, repo)
		if got != nil {
			t.Error("should be no error if all args are passed")
		}
	}()

	func() {
		// lacks some params
		authToken := "test-token"
		org := "org"
		repo := ""
		got := validateRequredArgs(authToken, org, repo)
		if got == nil {
			t.Error("should be error if repo is empty")
		}
	}()
}
