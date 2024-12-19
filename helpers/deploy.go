package helpers

import (
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func PrepareDeployment(
	id string,
	imageNameWithTag string,
	port int32,
	replicas int32,
	envVars map[string]string,
) *appsv1.Deployment {
	return &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Deployment",
			APIVersion: "apps/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: id,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": id,
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": id,
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  id,
							Image: imageNameWithTag,
							Ports: []apiv1.ContainerPort{
								{
									ContainerPort: port,
								},
							},
							Env: func() []apiv1.EnvVar {
								var envVarsSlice []apiv1.EnvVar
								for key, value := range envVars {
									envVarsSlice = append(envVarsSlice, apiv1.EnvVar{
										Name:  key,
										Value: value,
									})
								}
								return envVarsSlice
							}(),
						},
					},
				},
			},
		},
	}
}
