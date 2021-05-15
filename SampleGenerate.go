package main

import "log"

func Generate(args ...string) error {
	// 	生成模版
	buildErr := Execute("/scripts/build.sh", args[0], args[1], args[2], args[3])
	if buildErr != nil {
		log.Println("统一架构模版生成失败！")
		return buildErr
	}
	// gitlab创建项目
	project,buildProjectErr := BuildProject(args[1])
	if buildProjectErr != nil {
		log.Println("GitLab 创建项目失败！")
		return buildErr
	}
	// 上传git
	addErr := Execute("/scripts/add.sh")
	if addErr != nil {
		log.Println("统一架构模版上传 GitLab 失败！")
		return addErr
	}
	// 创建jenkins job
	buildJobErr := BuildJob(args[1],project.HTTPURLToRepo)

	if buildJobErr != nil {
		log.Println("创建 Jenkins Job 失败！")
		return buildJobErr
	}
	// 添加视图
	if len(args[3]) >= 0 {
		buildViewErr := BuildView(args[3], args[1])
		if buildViewErr != nil {
			log.Println("添加 Jenkins 视图失败！")
			return buildViewErr
		}
	}

	return nil
}
