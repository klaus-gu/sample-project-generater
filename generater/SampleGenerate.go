package generater

import (
	"errors"
	"log"
	"sample-project-generater/gitlab"
	"sample-project-generater/jenkins"
	"sample-project-generater/scripts"
)

func Generate(args ...string) error {
	// 	生成模版
	buildErr := scripts.Execute("/scripts/build.sh", args[0], args[1], args[2], args[3])
	if buildErr != nil {
		log.Printf("GitLab 统一架构模版生成失败！%v\n", buildErr.Error())
		return errors.New("统一架构模版生成失败!")
	}
	// gitlab创建项目
	project, buildProjectErr := gitlab.BuildProject(args[1])
	if buildProjectErr != nil {
		log.Printf("GitLab 创建项目失败！%v\n", buildProjectErr.Error())
		return errors.New("GitLab 创建项目失败！")
	}
	// 上传git
	addErr := scripts.Execute("/scripts/add.sh", args[1], project.HTTPURLToRepo, "", "", "", "")
	if addErr != nil {
		log.Printf("统一架构模版上传 GitLab 失败！%v\n", addErr.Error())
		return errors.New("统一架构模版上传 GitLab 失败！")
	}
	log.Println("统一架构模版上传 GitLab 成功！开始创建 Jenkins Job")
	// 创建jenkins job
	buildJobErr := jenkins.BuildJob(args[1], project.HTTPURLToRepo)
	if buildJobErr != nil {
		log.Printf("创建 Jenkins Job 失败！%v\n", buildJobErr.Error())
		return errors.New("创建 Jenkins Job 失败！")
	}
	log.Println("创建 Jenkins Job 成功！开始创建 Jenkins View")
	// 添加视图
	if len(args[3]) >= 0 {
		buildViewErr := jenkins.BuildView(args[3], args[1])
		if buildViewErr != nil {
			log.Printf("添加 Jenkins 视图失败！%v\n", buildViewErr.Error())
			return errors.New("添加 Jenkins 视图失败！")
		}
	}
	return nil
}
