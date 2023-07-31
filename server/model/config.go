package model

import v1 "k8s.io/client-go/applyconfigurations/core/v1"

type Configmap struct {
	v1.ConfigMapApplyConfiguration
}

type Secret struct {
	v1.SecretApplyConfiguration
}
