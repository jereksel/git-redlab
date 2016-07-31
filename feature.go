package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	redmine "github.com/mattn/go-redmine"
	"github.com/ryanuber/columnize"
)

func newFeature(g git) {

	redmineURL := g.getConfig(CONFIG_REDMINE_URL)
	redmineAPIKey := g.getConfig(CONFIG_REDMINE_API_KEY)
	redmineProjectID := g.getConfig(CONFIG_REDMINE_PROJECT_ID)

	redmineClient := redmine.NewClient(redmineURL, redmineAPIKey)

	editor := os.Getenv("EDITOR")

	if editor == "" {
		if runtime.GOOS == "windows" {
			editor = "notepad"
		} else {
			editor = "vim"
		}
	}

	binary, err := exec.LookPath(editor)
	if err != nil {
		panic(err)
	}

	dir, err := ioutil.TempDir("", "redlab")
	if err != nil {
		log.Fatal(err)
	}

	issueMsg := filepath.Join(dir, "ISSUE_MESSAGE")

	if _, err = os.Create(issueMsg); err != nil {
		panic(err)
	}

	cmd := exec.Command(binary, issueMsg)

	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Env = os.Environ()

	if err = cmd.Start(); err != nil {
		panic(err)
	}

	if err = cmd.Wait(); err != nil {
		panic(err)
	}

	_ = redmineClient

	topic, desc := readFile(issueMsg)

	if topic == "" && desc == "" {
		fmt.Println("No topic and description given. Aborting")
		return
	}

	redmineProjectIDInt, _ := strconv.Atoi(redmineProjectID)

	issue, err := redmineClient.CreateIssue(redmine.Issue{Subject: topic, Description: desc, ProjectId: redmineProjectIDInt, StatusId: 1, TrackerId: 2})

	if err != nil {
		panic(err)
	}

	fmt.Printf("Issue ID: %d\n", issue.Id)

}

func listIssues(g git) {

	redmineURL := g.getConfig(CONFIG_REDMINE_URL)
	redmineAPIKey := g.getConfig(CONFIG_REDMINE_API_KEY)
	redmineProjectID, _ := strconv.Atoi(g.getConfig(CONFIG_REDMINE_PROJECT_ID))

	redmineClient := redmine.NewClient(redmineURL, redmineAPIKey)

	_, _ = redmineProjectID, redmineClient

	fmt.Println("Issues:")

	issues, err := redmineClient.IssuesOf(redmineProjectID)

	if err != nil {
		panic(err)
	}

	output := make([]string, len(issues)+2)

	output[0] = "Status | Title"
	output[1] = "  |  "

	for id, issue := range issues {
		output[id+2] = fmt.Sprintf("%s | %s", issue.Status.Name, issue.GetTitle())
	}

	result := columnize.SimpleFormat(output)
	fmt.Println(result)

}

func pullIssue(g git, id int) {

	if !g.isRepoClean() {
		fmt.Println("This repository is not clean")
		os.Exit(1)
	}

	g.fetch()

	featureBranch := "feature/#" + strconv.Itoa(id)

	if featureBranch == g.getCurrentBranch() {
		return
	} else if stringInArray(featureBranch, g.getLocalBranches()) || stringInArray(featureBranch, g.getRemoteBranches()) {
		g.changeToBranch(featureBranch)
	} else {
		g.createNewBranch(featureBranch)
	}

}

func stringInArray(s string, arr []string) bool {

	for _, str := range arr {
		if s == str {
			return true
		}
	}
	return false
}

func readFile(file string) (string, string) {

	b, err := ioutil.ReadFile(file)

	if err != nil {
		panic(err)
	}

	str := string(b)

	strs := strings.SplitN(str, "\n", 2)

	if len(strs) == 1 {
		return strings.TrimSpace(strs[0]), ""
	}

	return strings.TrimSpace(strs[0]), strings.TrimSpace(strs[1])
}
