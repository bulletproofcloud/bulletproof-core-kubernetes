/*
Copyright 2022-2024 EscherCloud.
Copyright 2024 the Unikorn Authors.

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

package openstack

/*
import (
	"context"
	goerrors "errors"
	"fmt"
	"reflect"
	"slices"
	"sort"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/servergroups"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/applicationcredentials"
	lru "github.com/hashicorp/golang-lru/v2"

	"github.com/unikorn-cloud/identity/pkg/oauth2"
	unikornv1 "github.com/unikorn-cloud/unikorn/pkg/apis/unikorn/v1alpha1"
	"github.com/unikorn-cloud/unikorn/pkg/providers/openstack"
	"github.com/unikorn-cloud/unikorn/pkg/server/errors"
	"github.com/unikorn-cloud/unikorn/pkg/openapi"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	ErrResourceNotFound = goerrors.New("resource not found")
)

// covertError takes a generic gophercloud error and converts it into something
// more useful to our clients.
// / NOTE: at present we're only concerned with 401s because it reacts badly with
// the UI if we return a 500, when a 401 would cause a reauthentication and make
// the bad behaviour go away.
func covertError(err error) error {
	var err401 gophercloud.ErrDefault401

	if goerrors.As(err, &err401) {
		return errors.OAuth2AccessDenied("provider request denied").WithError(err)
	}

	var err403 gophercloud.ErrDefault403

	if goerrors.As(err, &err403) {
		return errors.HTTPForbidden("provider request forbidden, ensure you have the correct roles assigned to your user")
	}

	v := reflect.ValueOf(err)

	return errors.OAuth2ServerError("provider error unhandled: " + v.Type().Name()).WithError(err)
}

// findApplicationCredential, in the spirit of making the platform usable, allows
// a client to use names, rather than IDs for lookups.
func findApplicationCredential(in []applicationcredentials.ApplicationCredential, name string) (*applicationcredentials.ApplicationCredential, error) {
	for i, c := range in {
		if c.Name == name {
			return &in[i], nil
		}
	}

	return nil, errors.HTTPNotFound().WithError(fmt.Errorf("%w: application credential %s", ErrResourceNotFound, name))
}

func (o *Openstack) GetApplicationCredential(ctx context.Context, name string) (*applicationcredentials.ApplicationCredential, error) {
	user, err := getUser(ctx)
	if err != nil {
		return nil, err
	}

	client, err := o.IdentityClient(ctx)
	if err != nil {
		return nil, errors.OAuth2ServerError("failed get identity client").WithError(err)
	}

	result, err := client.ListApplicationCredentials(ctx, user)
	if err != nil {
		return nil, covertError(err)
	}

	match, err := findApplicationCredential(result, name)
	if err != nil {
		return nil, err
	}

	return match, nil
}

func (o *Openstack) CreateApplicationCredential(ctx context.Context, name string) (*applicationcredentials.ApplicationCredential, error) {
	user, err := getUser(ctx)
	if err != nil {
		return nil, err
	}

	client, err := o.IdentityClient(ctx)
	if err != nil {
		return nil, errors.OAuth2ServerError("failed get identity client").WithError(err)
	}

	description := "Automatically generated by platform service [DO NOT DELETE]."

	roles := []string{
		"member",
		"load-balancer_member",
	}

	if o.options.Identity != nil && o.options.Identity.ClusterRoles != nil {
		roles = o.options.Identity.ClusterRoles
	}

	result, err := client.CreateApplicationCredential(ctx, user, name, description, roles)
	if err != nil {
		return nil, covertError(err)
	}

	return result, nil
}

func (o *Openstack) DeleteApplicationCredential(ctx context.Context, name string) error {
	user, err := getUser(ctx)
	if err != nil {
		return err
	}

	client, err := o.IdentityClient(ctx)
	if err != nil {
		return errors.OAuth2ServerError("failed get identity client").WithError(err)
	}

	result, err := client.ListApplicationCredentials(ctx, user)
	if err != nil {
		return covertError(err)
	}

	match, err := findApplicationCredential(result, name)
	if err != nil {
		return err
	}

	if err := client.DeleteApplicationCredential(ctx, user, match.ID); err != nil {
		return errors.OAuth2ServerError("failed delete application credentials").WithError(err)
	}

	return nil
}

func (o *Openstack) GetServerGroup(ctx context.Context, name string) (*servergroups.ServerGroup, error) {
	client, err := o.ComputeClient(ctx)
	if err != nil {
		return nil, errors.OAuth2ServerError("failed get compute client").WithError(err)
	}

	result, err := client.ListServerGroups(ctx)
	if err != nil {
		return nil, covertError(err)
	}

	filtered := slices.DeleteFunc(result, func(group servergroups.ServerGroup) bool {
		return group.Name != name
	})

	switch len(filtered) {
	case 0:
		return nil, errors.HTTPNotFound().WithError(fmt.Errorf("%w: server group %s", ErrResourceNotFound, name))
	case 1:
		return &filtered[0], nil
	default:
		return nil, errors.OAuth2ServerError("multiple server groups matched name")
	}
}

func (o *Openstack) CreateServerGroup(ctx context.Context, name string) (*servergroups.ServerGroup, error) {
	client, err := o.ComputeClient(ctx)
	if err != nil {
		return nil, errors.OAuth2ServerError("failed get compute client").WithError(err)
	}

	result, err := client.CreateServerGroup(ctx, name)
	if err != nil {
		return nil, covertError(err)
	}

	return result, nil
}
*/
