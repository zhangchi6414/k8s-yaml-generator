package model

import (
	v12 "k8s.io/client-go/applyconfigurations/core/v1"
	v1 "k8s.io/client-go/applyconfigurations/rbac/v1"
)

type Role struct {
	v1.RoleApplyConfiguration
}

type ClusterRole struct {
	v1.ClusterRoleApplyConfiguration
}

type RoleBinding struct {
	v1.RoleBindingApplyConfiguration
}

type ClusterRoleBinding struct {
	v1.ClusterRoleBindingApplyConfiguration
}

type ServiceAccount struct {
	v12.ServiceAccountApplyConfiguration
}
