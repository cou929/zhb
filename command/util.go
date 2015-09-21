package command

import (
	"errors"
	"fmt"
	"github.com/codegangsta/cli"
	zenhub "github.com/cou929/zenhub-client"
	"os"
)

func Log(msg string) {
	fmt.Printf("[ERROR] %v\n", msg)
}

func DieIfError(e error) {
	if e != nil {
		Log(e.Error())
		os.Exit(1)
	}
}

func validateRequredArgs(authToken, org, repo string) error {
	if authToken == "" {
		return errors.New("Auth Token (ZHB_AUTH_TOKEN) is required")
	}

	if org == "" {
		return errors.New("Organization (ZHB_ORG) is required")
	}

	if repo == "" {
		return errors.New("Repository (ZHB_REPO) is required")
	}

	return nil
}

func NewZenhubClient(c *cli.Context) *zenhub.Client {
	authToken := c.GlobalString("authtoken")
	org := c.GlobalString("org")
	repo := c.GlobalString("repo")

	err := validateRequredArgs(authToken, org, repo)
	DieIfError(err)

	client := zenhub.NewClient(authToken, org, repo)
	client.Verbose = c.GlobalBool("verbose")

	return client
}
