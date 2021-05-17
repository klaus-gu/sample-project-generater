package jenkins

import (
	"context"
	"github.com/bndr/gojenkins"
	"log"
)

func BuildJob(jobName string, gitPath string) error {
	jobConfig := getConfig(gitPath)
	ctx := context.Background()
	jb, jbErr := jenkins.GetJob(ctx, jobName)
	if jbErr != nil {
		return jbErr
	}
	if jb == nil {
		_, jobErr := jenkins.CreateJob(ctx, jobConfig, jobName)
		if jobErr != nil {
			return jobErr
		}
	}
	return nil
}

func BuildView(viewName string, jobName string) error {
	var view *gojenkins.View
	var viewErr error
	ctx := context.Background()
	view, viewErr = jenkins.GetView(ctx, viewName)
	if viewErr != nil {
		return viewErr
	}
	if view == nil {
		view, viewErr = jenkins.CreateView(ctx, viewName, LIST_VIEW)
		if viewErr != nil {
			return viewErr
		}
	}
	vb, viewErr := view.AddJob(ctx, jobName)
	if vb {
		return nil
	}
	return viewErr

}

var (
	jenkins_user   = "guyue"
	jenkins_pwd    = "guyue375"
	jenkins_url    = "http://work01.avengers-inc.ovopark.com:8080/"
	jenkins        *gojenkins.Jenkins
	LIST_VIEW      = "hudson.model.ListView"
	NESTED_VIEW    = "hudson.plugins.nested_view.NestedView"
	MY_VIEW        = "hudson.model.MyView"
	DASHBOARD_VIEW = "hudson.plugins.view.dashboard.Dashboard"
	PIPELINE_VIEW  = "au.com.centrumsystems.hudson.plugin.buildpipeline.BuildPipelineView"
)

func init() {
	jenkins = gojenkins.CreateJenkins(nil, jenkins_url, jenkins_user, jenkins_pwd)
	ctx := context.Background()
	_, err := jenkins.Init(ctx)
	if err != nil {
		log.Printf("连接Jenkins失败!, %v\n", err)
		return
	}
	log.Println("Jenkins 连接成功!")
}

func getConfig(gitPath string) string {
	var prex = "<?xml version='1.1' encoding='UTF-8'?>\n<project>\n  <actions/>\n  <description></description>\n  <keepDependencies>false</keepDependencies>\n  <properties/>\n  <scm class=\"hudson.plugins.git.GitSCM\" plugin=\"git@4.7.1\">\n    <configVersion>2</configVersion>\n    <userRemoteConfigs>\n      <hudson.plugins.git.UserRemoteConfig>\n        <url>"
	var next = "</url>\n      </hudson.plugins.git.UserRemoteConfig>\n    </userRemoteConfigs>\n    <branches>\n      <hudson.plugins.git.BranchSpec>\n        <name>*/master</name>\n      </hudson.plugins.git.BranchSpec>\n    </branches>\n    <doGenerateSubmoduleConfigurations>false</doGenerateSubmoduleConfigurations>\n    <submoduleCfg class=\"empty-list\"/>\n    <extensions/>\n  </scm>\n  <canRoam>true</canRoam>\n  <disabled>false</disabled>\n  <blockBuildWhenDownstreamBuilding>false</blockBuildWhenDownstreamBuilding>\n  <blockBuildWhenUpstreamBuilding>false</blockBuildWhenUpstreamBuilding>\n  <triggers/>\n  <concurrentBuild>false</concurrentBuild>\n  <builders>\n    <hudson.tasks.Shell>\n      <command>sh /usr/local/jenkins/buildCheck.sh</command>\n      <configuredLocalRules/>\n    </hudson.tasks.Shell>\n    <hudson.tasks.Maven>\n      <targets>clean install -Dmaven.test.skip=true</targets>\n      <mavenName>apache-maven-3.6.3</mavenName>\n      <usePrivateRepository>false</usePrivateRepository>\n      <settings class=\"jenkins.mvn.DefaultSettingsProvider\"/>\n      <globalSettings class=\"jenkins.mvn.DefaultGlobalSettingsProvider\"/>\n      <injectBuildVariables>false</injectBuildVariables>\n    </hudson.tasks.Maven>\n  </builders>\n  <publishers>\n    <jenkins.plugins.publish__over__ssh.BapSshPublisherPlugin plugin=\"publish-over-ssh@1.22\">\n      <consolePrefix>SSH: </consolePrefix>\n      <delegate plugin=\"publish-over@0.22\">\n        <publishers>\n          <jenkins.plugins.publish__over__ssh.BapSshPublisher plugin=\"publish-over-ssh@1.22\">\n            <configName>172.16.22.37-SSH</configName>\n            <verbose>true</verbose>\n            <transfers>\n              <jenkins.plugins.publish__over__ssh.BapSshTransfer>\n                <remoteDirectory>/remote</remoteDirectory>\n                <sourceFiles>consumer/target/*.jar</sourceFiles>\n                <excludes></excludes>\n                <removePrefix>/consumer/target/</removePrefix>\n                <remoteDirectorySDF>false</remoteDirectorySDF>\n                <flatten>false</flatten>\n                <cleanRemote>false</cleanRemote>\n                <noDefaultExcludes>false</noDefaultExcludes>\n                <makeEmptyDirs>false</makeEmptyDirs>\n                <patternSeparator>[, ]+</patternSeparator>\n                <execCommand>sh /usr/local/java/restart.sh\n</execCommand>\n                <execTimeout>120000</execTimeout>\n                <usePty>true</usePty>\n                <useAgentForwarding>false</useAgentForwarding>\n                <useSftpForExec>false</useSftpForExec>\n              </jenkins.plugins.publish__over__ssh.BapSshTransfer>\n            </transfers>\n            <useWorkspaceInPromotion>false</useWorkspaceInPromotion>\n            <usePromotionTimestamp>false</usePromotionTimestamp>\n          </jenkins.plugins.publish__over__ssh.BapSshPublisher>\n          <jenkins.plugins.publish__over__ssh.BapSshPublisher plugin=\"publish-over-ssh@1.22\">\n            <configName>172.16.22.37-SSH</configName>\n            <verbose>true</verbose>\n            <transfers>\n              <jenkins.plugins.publish__over__ssh.BapSshTransfer>\n                <remoteDirectory>/remote</remoteDirectory>\n                <sourceFiles>provider/target/*.jar</sourceFiles>\n                <excludes></excludes>\n                <removePrefix>/provider/target/</removePrefix>\n                <remoteDirectorySDF>false</remoteDirectorySDF>\n                <flatten>false</flatten>\n                <cleanRemote>false</cleanRemote>\n                <noDefaultExcludes>false</noDefaultExcludes>\n                <makeEmptyDirs>false</makeEmptyDirs>\n                <patternSeparator>[, ]+</patternSeparator>\n                <execCommand>sh /usr/local/java/restart.sh\n</execCommand>\n                <execTimeout>120000</execTimeout>\n                <usePty>true</usePty>\n                <useAgentForwarding>false</useAgentForwarding>\n                <useSftpForExec>false</useSftpForExec>\n              </jenkins.plugins.publish__over__ssh.BapSshTransfer>\n            </transfers>\n            <useWorkspaceInPromotion>false</useWorkspaceInPromotion>\n            <usePromotionTimestamp>false</usePromotionTimestamp>\n          </jenkins.plugins.publish__over__ssh.BapSshPublisher>\n        </publishers>\n        <continueOnError>false</continueOnError>\n        <failOnError>false</failOnError>\n        <alwaysPublishFromMaster>false</alwaysPublishFromMaster>\n        <hostConfigurationAccess class=\"jenkins.plugins.publish_over_ssh.BapSshPublisherPlugin\" reference=\"../..\"/>\n      </delegate>\n    </jenkins.plugins.publish__over__ssh.BapSshPublisherPlugin>\n  </publishers>\n  <buildWrappers>\n    <hudson.plugins.timestamper.TimestamperBuildWrapper plugin=\"timestamper@1.13\"/>\n  </buildWrappers>\n</project>"
	return prex + gitPath + next
}
