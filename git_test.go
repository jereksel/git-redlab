package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGitExistence(t *testing.T) {

	assert := assert.New(t)

	g := git{"/tmp/git-redlab-tests/git-here/"}

	assert.True(g.IsThisGitDir(), "First dir should be git")

}

func TestGitNonExistence(t *testing.T) {

	assert := assert.New(t)

	g := git{"/tmp/git-redlab-tests/no-git-here/"}

	assert.False(g.IsThisGitDir(), "Second dir shouldn't be git")

}

func TestGitGetConfig(t *testing.T) {

	assert := assert.New(t)

	g := git{"/tmp/git-redlab-tests/git-here/"}

	assert.Equal("xyz", g.getConfig("test.abc"), "Config test.abc should be xyz")

}

func TestGitPutConfig(t *testing.T) {

	assert := assert.New(t)

	g := git{"/tmp/git-redlab-tests/git-here/"}

	g.putConfig("test.xyz", "abc")

	assert.Equal("abc", g.getConfig("test.xyz"), "Config test.xyz should be abc")

}

func TestGitCleanConfig(t *testing.T) {

	assert := assert.New(t)

	g := git{"/tmp/git-redlab-tests/git-here/"}

	g.putConfig("test.qwe", "qwe")
	g.clearConfig("test.qwe")

	assert.Equal("", g.getConfig("test.123"), "Config test.xyz should be empty")

}

func TestGitRepoClean(t *testing.T) {

	assert := assert.New(t)

	g := git{"/tmp/git-redlab-tests/clean-git"}

	assert.True(g.isRepoClean(), "Repo should be clean")

}

func TestGitRepoDirty(t *testing.T) {

	assert := assert.New(t)

	g := git{"/tmp/git-redlab-tests/dirty-git"}

	assert.False(g.isRepoClean(), "Repo should be dirty")

}

func TestGitBranchName(t *testing.T) {

	assert := assert.New(t)

	g := git{"/tmp/git-redlab-tests/qwerty-branch"}

	assert.Equal("qwerty", g.getCurrentBranch(), "Branch name should be qwerty")

}

func TestGitLocalBranches(t *testing.T) {

	assert := assert.New(t)

	g := git{"/tmp/git-redlab-tests/multiple-branches"}

	assert.Equal([]string{"branch1", "branch2", "branch3", "feature/#1", "master"}, g.getLocalBranches(), "Branches names")

}

func TestGitRemoteBranches(t *testing.T) {

	assert := assert.New(t)

	g := git{"/tmp/git-redlab-tests/remote-branches"}

	assert.Equal([]string{"#123", "branch1", "branch2", "branch3", "feature/#1", "master"}, g.getRemoteBranches(), "Branches names")

}
