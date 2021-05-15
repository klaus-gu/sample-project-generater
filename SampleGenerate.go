package main

import (
	"errors"
	"log"
)

func Generate(args ...string) error {
	// 	生成模版
	buildErr := Execute("/scripts/build.sh", args[0], args[1], args[2], args[3])
	if buildErr != nil {
		log.Printf("GitLab 统一架构模版生成失败！%v\n", buildErr.Error())
		return errors.New("统一架构模版生成失败!")
	}
	// gitlab创建项目
	project, buildProjectErr := BuildProject(args[1])
	if buildProjectErr != nil {
		log.Printf("GitLab 创建项目失败！%v\n", buildProjectErr.Error())
		return errors.New("GitLab 创建项目失败！")
	}
	// 上传git
	addErr := Execute("/scripts/add.sh")
	if addErr != nil {
		log.Printf("统一架构模版上传 GitLab 失败！%v\n", addErr.Error())
		return errors.New("统一架构模版上传 GitLab 失败！")
	}
	// 创建jenkins job
	buildJobErr := BuildJob(args[1], project.HTTPURLToRepo)

	if buildJobErr != nil {
		log.Printf("创建 Jenkins Job 失败！%v\n", buildJobErr.Error())
		return errors.New("创建 Jenkins Job 失败！")
	}
	// 添加视图
	if len(args[3]) >= 0 {
		buildViewErr := BuildView(args[3], args[1])
		if buildViewErr != nil {
			log.Printf("添加 Jenkins 视图失败！%v\n", buildViewErr.Error())
			return errors.New("添加 Jenkins 视图失败！")
		}
	}
	return nil
}
