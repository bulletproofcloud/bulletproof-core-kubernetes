/*
Copyright 2022 EscherCloud.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package util

import (
	"context"
	"errors"
	"fmt"

	"github.com/eschercloudai/unikorn/generated/clientset/unikorn"
	"github.com/eschercloudai/unikorn/pkg/constants"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
)

var (
	// ErrUnavailable is for when the resource status reports unready.
	ErrUnavailable = errors.New("resource unavailable")

	// ErrNamespace is for when the resource status doesn't contain a namespace.
	ErrNamespace = errors.New("namespace error")
)

// GetProjectNamespace figures out the namespace associated with a project.
func GetProjectNamespace(ctx context.Context, client unikorn.Interface, project string) (string, error) {
	p, err := client.UnikornV1alpha1().Projects().Get(ctx, project, metav1.GetOptions{})
	if err != nil {
		return "", err
	}

	namespace := p.Status.Namespace

	if namespace == "" {
		return "", fmt.Errorf("%w: project namespace unset", ErrNamespace)
	}

	return namespace, nil
}

// GetControlPlaneNamespace figures out the namespace associated with a project's control plane.
func GetControlPlaneNamespace(ctx context.Context, client unikorn.Interface, project, controlPlane string) (string, error) {
	p, err := client.UnikornV1alpha1().Projects().Get(ctx, project, metav1.GetOptions{})
	if err != nil {
		return "", err
	}

	namespace := p.Status.Namespace

	if namespace == "" {
		return "", fmt.Errorf("%w: project namespace unset", ErrNamespace)
	}

	cp, err := client.UnikornV1alpha1().ControlPlanes(namespace).Get(ctx, controlPlane, metav1.GetOptions{})
	if err != nil {
		return "", err
	}

	namespace = cp.Status.Namespace

	if namespace == "" {
		return "", fmt.Errorf("%w: control plane namespace unset", ErrNamespace)
	}

	return namespace, nil
}

// GetClusterWorkloadPools gets workload pools linked to a cluster in a project's control plane.
func GetClusterWorkloadPools(ctx context.Context, client unikorn.Interface, project, controlPlane, cluster string) ([]string, error) {
	namespace, err := GetControlPlaneNamespace(ctx, client, project, controlPlane)
	if err != nil {
		return nil, err
	}

	selector := labels.Everything()

	if cluster != "" {
		clusterLabel, err := labels.NewRequirement(constants.KubernetesClusterLabel, selection.Equals, []string{cluster})
		if err != nil {
			return nil, err
		}

		selector = selector.Add(*clusterLabel)
	}

	pools, err := client.UnikornV1alpha1().KubernetesWorkloadPools(namespace).List(ctx, metav1.ListOptions{LabelSelector: selector.String()})
	if err != nil {
		return nil, err
	}

	names := make([]string, len(pools.Items))

	for i, pool := range pools.Items {
		names[i] = pool.Name
	}

	return names, nil
}
