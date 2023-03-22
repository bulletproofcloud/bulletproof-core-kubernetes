/*
Copyright 2022-2023 EscherCloud.

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

package metricsserver

import (
	unikornv1 "github.com/eschercloudai/unikorn/pkg/apis/unikorn/v1alpha1"
	"github.com/eschercloudai/unikorn/pkg/provisioners"
	"github.com/eschercloudai/unikorn/pkg/provisioners/application"
	"github.com/eschercloudai/unikorn/pkg/provisioners/util"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	// applicationName is the unique name of the application.
	applicationName = "metrics-server"
)

type Provisioner struct{}

// Ensure the Provisioner interface is implemented.
var _ application.ValuesGenerator = &Provisioner{}

// New returns a new initialized provisioner object.
func New(client client.Client, resource application.MutuallyExclusiveResource, helm *unikornv1.HelmApplication) provisioners.Provisioner {
	p := &Provisioner{}

	return application.New(client, applicationName, resource, helm).WithGenerator(p).InNamespace("kube-system")
}

// Generate implements the application.Generator interface.
// This forces the server onto the control plane rather than take up a
// worker node (and thus incur the ire of users).
func (p *Provisioner) Values(version *string) (interface{}, error) {
	values := map[string]interface{}{
		"tolerations":  util.ControlPlaneTolerations(),
		"nodeSelector": util.ControlPlaneNodeSelector(),
	}

	return values, nil
}