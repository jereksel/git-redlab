package main

import (
	"fmt"
	"os"
)

type io interface {
	ScanString(a *string) (n int, err error)
	ScanInt(a *int) (n int, err error)
	Print(a ...interface{}) (n int, err error)
	Printf(format string, a ...interface{}) (n int, err error)
}

type ioIpml struct {
}

func (io ioIpml) ScanString(a *string) (n int, err error) {
	return fmt.Fscan(os.Stdin, a)
}

func (io ioIpml) ScanInt(a *int) (n int, err error) {
	return fmt.Fscan(os.Stdin, a)
}

func (io ioIpml) Print(a ...interface{}) (n int, err error) {
	return fmt.Fprint(os.Stdout, a...)
}

func (io ioIpml) Printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(os.Stdout, format, a...)
}

var CONFIG_REDMINE_URL = "redlab.redmine.url"
var CONFIG_REDMINE_API_KEY = "redlab.redmine.apikey"
var CONFIG_REDMINE_PROJECT_ID = "redlab.redmine.projectid"

var CONFIG_GITLAB_URL = "redlab.gitlab.url"
var CONFIG_GITLAB_API_KEY = "redlab.gitlab.apikey"
var CONFIG_GITLAB_PROJECT_ID = "redlab.gitlab.projectid"

func main() {

	firstArg := os.Args[1]

	g := git{"."}

	if !g.IsThisGitDir() {
		fmt.Println("fatal: Not a git repository")
		os.Exit(1)
	}

	switch firstArg {
	case "init":
		repoInit(g, ioIpml{})
		break
	case "clean":
		cleanRepo(g)
		break
	case "feature":
		ensureInit(g)
		newFeature(g)
		break
	default:
		fmt.Printf("Unknown command: %s\n", firstArg)
		os.Exit(1)
	}

}

func ensureInit(g git) {
	if !isInitialized(g) {
		fmt.Println("Repo is not initialized. Please run \"git redlab init\" first")
		os.Exit(1)
	}
}

func isInitialized(g git) bool {
	return g.getConfig(CONFIG_GITLAB_API_KEY) != "" &&
		g.getConfig(CONFIG_GITLAB_PROJECT_ID) != "" &&
		g.getConfig(CONFIG_GITLAB_URL) != "" &&
		g.getConfig(CONFIG_REDMINE_API_KEY) != "" &&
		g.getConfig(CONFIG_REDMINE_PROJECT_ID) != "" &&
		g.getConfig(CONFIG_REDMINE_URL) != ""
}

func cleanRepo(g git) {
	g.clearConfig(CONFIG_GITLAB_API_KEY)
	g.clearConfig(CONFIG_GITLAB_PROJECT_ID)
	g.clearConfig(CONFIG_GITLAB_URL)
	g.clearConfig(CONFIG_REDMINE_API_KEY)
	g.clearConfig(CONFIG_REDMINE_PROJECT_ID)
	g.clearConfig(CONFIG_REDMINE_URL)
	fmt.Println("Repo cleared")
}

func usage() {

}
