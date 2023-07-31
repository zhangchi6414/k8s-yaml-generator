package model

import (
	"k8s.io/client-go/applyconfigurations/apps/v1beta1"
	"k8s.io/client-go/applyconfigurations/apps/v1beta2"
	v12 "k8s.io/client-go/applyconfigurations/batch/v1"
	v1 "k8s.io/client-go/applyconfigurations/core/v1"
)

type Pod struct {
	v1.PodApplyConfiguration
}

type Deployment struct {
	v1beta1.DeploymentApplyConfiguration
}

type StatefulSets struct {
	v1beta1.StatefulSetApplyConfiguration
}

type DaemonSet struct {
	v1beta2.DaemonSetApplyConfiguration
}

type Job struct {
	v12.JobApplyConfiguration
}

type CronJob struct {
	v12.CronJobApplyConfiguration
}
