package helpers

import (
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func PrepareService(
	id string,
	port int32,
) *apiv1.Service {
	return &apiv1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: id,
		},
		Spec: apiv1.ServiceSpec{
			Selector: map[string]string{
				"app": id,
			},
			Ports: []apiv1.ServicePort{
				{
					Port: port,
				},
			},
			Type: apiv1.ServiceTypeNodePort,
		},
	}
}
