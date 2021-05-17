package gitlab

import (
	"github.com/xanzy/go-gitlab"
	"log"
)

func BuildProject(pjName string) (*gitlab.Project, error) {
	var pj *gitlab.Project
	pl := &gitlab.ListProjectsOptions{}
	pList, _, err := git.Projects.ListProjects(pl)
	if err == nil && len(pList) >= 0 {
		for _, rpj := range pList {
			if rpj.Name == pjName {
				return rpj, nil
			}
		}
	}
	p := &gitlab.CreateProjectOptions{
		Name:                 gitlab.String(pjName),
		Description:          gitlab.String("Unified framework template"),
		MergeRequestsEnabled: gitlab.Bool(true),
		SnippetsEnabled:      gitlab.Bool(true),
		Visibility:           gitlab.Visibility(gitlab.PublicVisibility),
	}
	pj, _, pjErr := git.Projects.CreateProject(p)
	if pjErr != nil {
		return nil, pjErr
	}
	return pj, nil
}

var (
	gitlab_token = "jZadCsFU8eFJt5NSRyFh"
	gitlab_url   = "http://work01.avengers-inc.ovopark.com/"
	git          *gitlab.Client
)

func init() {
	var err error
	git, err = gitlab.NewClient(gitlab_token, gitlab.WithBaseURL(gitlab_url))
	if err != nil {
		log.Println("GitLab 连接失败！")
		return
	}
	log.Println("GitLab 连接成功！")
}
