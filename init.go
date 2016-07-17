package main

import (
	"strconv"

	redmine "github.com/mattn/go-redmine"
	gitlab "github.com/xanzy/go-gitlab"
)

type name func(interface{}) string

func repoInit(g git, system io) {

	if g.getConfig(CONFIG_REDMINE_URL) == "" ||
		g.getConfig(CONFIG_REDMINE_PROJECT_ID) == "" ||
		g.getConfig(CONFIG_REDMINE_API_KEY) == "" {
		initRedmine(g, system)
	}

	if g.getConfig(CONFIG_GITLAB_URL) == "" ||
		g.getConfig(CONFIG_GITLAB_PROJECT_ID) == "" ||
		g.getConfig(CONFIG_GITLAB_API_KEY) == "" {
		initGitlab(g, system)
	}

}

func initRedmine(g git, system io) {

	var redmineURL, redmineAPIKey string

	system.Print("Type redmine's url (e.g https://www.redmine.org): ")

	system.ScanString(&redmineURL)

	system.Printf("Paste redmine's api key (you can find it here: %s/my/account ): ", redmineURL)

	system.ScanString(&redmineAPIKey)

	system.Print("Redmine projects\n")

	redmineClient := redmine.NewClient(redmineURL, redmineAPIKey)

	projects, err := redmineClient.Projects()

	if err != nil {
		panic(err)
	}

	t := projects
	s := make([]interface{}, len(t))
	for i, v := range t {
		s[i] = v
	}

	redmineProjectID := selection(s, system, func(i interface{}) string {
		return i.(redmine.Project).Name
	})

	system.Printf("Selected redmine project: %s\n", projects[redmineProjectID].Name)

	_ = redmineProjectID

	g.putConfig(CONFIG_REDMINE_URL, redmineURL)
	g.putConfig(CONFIG_REDMINE_API_KEY, redmineAPIKey)
	g.putConfig(CONFIG_REDMINE_PROJECT_ID, strconv.Itoa(projects[redmineProjectID].Id))

}

func initGitlab(g git, system io) {

	var gitlabURL, gitlabAPIKey string

	system.Print("Type gitlab's url (e.g https://gitlab.com): ")

	system.ScanString(&gitlabURL)

	system.Printf("Paste gitlab's api key (you can find it here: %s/profile/personal_access_tokens ): ", gitlabURL)

	system.ScanString(&gitlabAPIKey)

	system.Print("Gitlab projects\n")

	git := gitlab.NewClient(nil, gitlabAPIKey)
	git.SetBaseURL(gitlabURL + "/api/v3")

	projects, _, err := git.Projects.ListProjects(nil)

	if err != nil {
		panic(err)
	}

	t := projects
	s := make([]interface{}, len(t))
	for i, v := range t {
		s[i] = v
	}

	gitlabProjectID := selection(s, system, func(i interface{}) string {
		return *i.(*gitlab.Project).NameWithNamespace
	})

	system.Printf("Selected redmine project: %s\n", *projects[gitlabProjectID].NameWithNamespace)

	g.putConfig(CONFIG_GITLAB_URL, gitlabURL)
	g.putConfig(CONFIG_GITLAB_API_KEY, gitlabAPIKey)
	g.putConfig(CONFIG_GITLAB_PROJECT_ID, strconv.Itoa(*projects[gitlabProjectID].ID))

}

func selection(projects []interface{}, system io, getName name) int {

	for i := 0; i < len(projects); i++ {
		project := projects[i]
		system.Printf("[%d]: %s\n", i, getName(project))
	}

	for {

		system.Print("Please select project index: ")

		var projectID int

		system.ScanInt(&projectID)

		if projectID <= 0 || projectID >= len(projects) {
			system.Print("Wrong index given\n")
		} else {
			return projectID
		}

	}
}
