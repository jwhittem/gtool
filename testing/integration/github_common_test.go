package integration

import (
	"context"
	"os"
	"testing"

	"github.com/jwhittem/gtool/cmd"
	"github.com/stretchr/testify/assert"
)

func Test_NewGithubClient(t *testing.T) {
	assert := assert.New(t)

	//recommend a separate token with limited privs for integration testing.
	token := os.Getenv("GITHUB_AUTH_TOKEN")
	client := cmd.NewGithubClient(token)

	//Zen returns random philosophical strings....
	Zen, _, _ := client.Zen(context.Background())
	assert.NotEmpty(Zen, "Should get a 'Zen' message from github api.")
}

func Test_GetAllRepos(t *testing.T) {
	assert := assert.New(t)

	//recommend a separate token with limited privs for integration testing.
	token := os.Getenv("GITHUB_AUTH_TOKEN")
	client := cmd.NewGithubClient(token)

	checkOrg := "jwhittem"
	allRepos := cmd.GetAllRepos(client, checkOrg)

	assert.NotNil(allRepos, "Can't find any repos for: %s", checkOrg)
}
