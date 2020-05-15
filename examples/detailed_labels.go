package main

import (
	"github.com/xanzy/go-gitlab"
	"log"
)

func detailedLabels() {

	git, _ := gitlab.NewClient("yourTokenhere")

	issuesOptions := &gitlab.ListGroupIssuesOptions{
		ListOptions: gitlab.ListOptions{
			PerPage: 100,
			Page:    1,
		},
		State:            gitlab.String("opened"),
		Scope:            gitlab.String("all"),
		Milestone:        gitlab.String("YourSprint"),
		WithLabelDetails: gitlab.Bool(true),
	}

	issues, response, err := git.Issues.ListGroupIssues("YourGroup", issuesOptions)
	if err != nil {
		log.Print(err)
	}

	log.Print(issues)
	log.Print(response)

}
