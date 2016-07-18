package main

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
)

type git struct {
	location string
}

func (g git) IsThisGitDir() bool {

	goToDir(g.location)

	cmd := exec.Command("git", "status")

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	if err := cmd.Start(); err != nil {
		panic(err)
	}

	if err := cmd.Wait(); err != nil {
		//It means that command returns non-zero error code - that means
		//there is no git repo
		return false
	}

	return true

}

func (g git) getConfig(config string) string {

	goToDir(g.location)

	cmd := exec.Command("git", "config", "--get", config)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	if err := cmd.Start(); err != nil {
		panic(err)
	}

	if err := cmd.Wait(); err != nil {
		//It means that command returns non-zero error code - that means there is no
		//such config
		return ""
	}

	//We remove last character (\n)
	return string(out.Bytes()[:len(out.Bytes())-1])

}

func (g git) putConfig(configName string, configValue string) {

	goToDir(g.location)

	cmd := exec.Command("git", "config", "--add", configName, configValue)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	if err := cmd.Start(); err != nil {
		panic(err)
	}

	if err := cmd.Wait(); err != nil {
		panic(err)
	}

}

func (g git) clearConfig(configName string) {

	if g.getConfig(configName) == "" {
		return
	}

	goToDir(g.location)

	cmd := exec.Command("git", "config", "--unset-all", configName)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	if err := cmd.Start(); err != nil {
		panic(err)
	}

	if err := cmd.Wait(); err != nil {
		panic(err)
	}
}

func (g git) isRepoClean() bool {

	goToDir(g.location)

	cmd := exec.Command("git", "diff", "--no-ext-diff", "--ignore-submodules", "--quiet", "--exit-code")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	if err := cmd.Start(); err != nil {
		panic(err)
	}

	if err := cmd.Wait(); err != nil {
		//It means that command returns non-zero error code - that means there is no
		//such config
		return false
	}

	//We remove last character (\n)
	return true
}

func (g git) getCurrentBranch() string {

	goToDir(g.location)

	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	if err := cmd.Start(); err != nil {
		panic(err)
	}

	if err := cmd.Wait(); err != nil {
		//It means that command returns non-zero error code - that means there is no
		//such config
		panic(err)
	}

	//We remove last character (\n)
	return string(out.Bytes()[:len(out.Bytes())-1])

}

func (g git) getLocalBranches() []string {

	goToDir(g.location)

	cmd := exec.Command("git", "branch", "--no-color")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	if err := cmd.Start(); err != nil {
		panic(err)
	}

	if err := cmd.Wait(); err != nil {
		//It means that command returns non-zero error code - that means there is no
		//such config
		panic(err)
	}

	var branches []string

	//We remove last character (\n)
	ret := string(out.Bytes()[:len(out.Bytes())-1])

	for _, s := range strings.Split(ret, "\n") {
		//FIXME
		branches = append(branches, strings.Replace(strings.Replace(s, "  ", "", 1), "* ", "", 1))
	}

	return branches

}

func (g git) getRemoteBranches() []string {

	goToDir(g.location)

	cmd := exec.Command("git", "branch", "-r")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	if err := cmd.Start(); err != nil {
		panic(err)
	}

	if err := cmd.Wait(); err != nil {
		panic(err)
	}

	var branches []string

	//We remove last character (\n)
	ret := string(out.Bytes()[:len(out.Bytes())-1])

	for _, s := range strings.Split(ret, "\n") {

		branch := strings.SplitN(s, "/", 2)[1]

		if strings.HasPrefix(branch, "HEAD ->") {
			continue
		}

		branches = append(branches, branch)

	}

	return branches

}

func getCurrentDir() string {
	str, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return str
}

func goToDir(dir string) {
	if err := os.Chdir(dir); err != nil {
		panic(err)
	}
}
