package main

import (
	"fmt"
	"os"

	cli "github.com/jawher/mow.cli"
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

	g := git{"."}

	app := cli.App("git redlab", "")
	//app.Before = func() { ensureInit(g) }

	app.Command("init", "Initialize a git redlab support in existing git repo", func(cmd *cli.Cmd) { cmd.Action = func() { repoInit(g, ioIpml{}) } })
	app.Command("feature", "Manage features (redmine)", func(cmd *cli.Cmd) {
		cmd.Before = func() { ensureInit(g) }
		cmd.Command("start", "Start new feature", func(cmd *cli.Cmd) { cmd.Action = func() { newFeature(g) } })
		cmd.Command("pull", "Pull existing feature", func(cmd *cli.Cmd) {

			cmd.Spec = "--id"

			var (
				issueID = cmd.IntOpt("id", 0, "Issue id")
			)

			cmd.Action = func() {
				pullIssue(g, *issueID)
			}
		})
		cmd.Command("list", "List opened features", func(cmd *cli.Cmd) { cmd.Action = func() { listIssues(g) } })
		cmd.Command("publish", "Push changes from current feature", func(cmd *cli.Cmd) { cmd.Action = func() { publishFeature(g) } })
	})
	app.Command("clean", "Remove all redlab related configs from repo", func(cmd *cli.Cmd) { cmd.Action = func() { cleanRepo(g) } })

	app.Run(os.Args)

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
	fmt.Println("git redlab <subcommand>")
	fmt.Println("")
	fmt.Println("Available subcommand are:")
	fmt.Println("   init     Initialize a git redlab support in existing git repo")

}
