package model

import (
	v1 "k8s.io/client-go/applyconfigurations/core/v1"
	v12 "k8s.io/client-go/applyconfigurations/networking/v1"
)

type Service struct {
	v1.ServiceApplyConfiguration
}

type Endpoint struct {
	v1.EndpointsApplyConfiguration
}

type Ingress struct {
	v12.IngressApplyConfiguration
}
