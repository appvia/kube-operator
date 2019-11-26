/*
 * Copyright (C) 2019 Rohith Jayawardene <gambol99@gmail.com>
 *
 * This program is free software; you can redistribute it and/or
 * modify it under the terms of the GNU General Public License
 * as published by the Free Software Foundation; either version 2
 * of the License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package main

import (
	"context"
	"fmt"

	kubev1 "github.com/appvia/kube-operator/pkg/apis/kube/v1"
	"github.com/appvia/kube-operator/pkg/schema"

	clusterv1 "github.com/appvia/hub-apis/pkg/apis/clusters/v1"
	configv1 "github.com/appvia/hub-apis/pkg/apis/config/v1"
	"github.com/appvia/hub-apis/pkg/publish"
	hschema "github.com/appvia/hub-apis/pkg/schema"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

var (
	kubeClass = configv1.Class{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "kubernetes",
			Namespace: "hub",
		},
		Spec: configv1.ClassSpec{
			APIVersion:    kubev1.SchemeGroupVersion.String(),
			AutoProvision: true,
			Category:      "kubernetes",
			Description:   "Provides the ability to create a number of resouce types a kubernetes cluster",
			DisplayName:   "Kubernetes Provideir",
			Requires: metav1.GroupVersionKind{
				Group:   clusterv1.SchemeGroupVersion.Group,
				Kind:    "Kubernetes",
				Version: clusterv1.SchemeGroupVersion.Version,
			},
			Plans: []string{},
			Resources: []configv1.ClassResource{
				{
					Group:            kubev1.SchemeGroupVersion.Group,
					Kind:             "NamespaceClaim",
					Plans:            []string{},
					DisplayName:      "Kubernetes Namespace ",
					ShortDescription: "Provisions a kubernetes namespace on a cluster",
					LongDescription:  "Provisions a kubernetes namespace, the role and binding for members of the team",
					Scope:            configv1.WorkspaceScope,
					Version:          kubev1.SchemeGroupVersion.Version,
				},
			},
			Schemas: schema.ConvertToJSON(),
		},
	}
)

// publishOperator is responsible for injecting the classes configuration
// into the hub-api and crds
func publishOperator(cfg *rest.Config) error {
	// @step: publish the CRDs in the hub
	ac, err := publish.NewExtentionsAPIClient(cfg)
	if err != nil {
		return err
	}

	if err := publish.ApplyCustomResourceDefinitions(ac, schema.GetCustomResourceDefinitions()); err != nil {
		return fmt.Errorf("failed to register the operator crds: %s", err)
	}

	c, err := hschema.NewClient(cfg)
	if err != nil {
		return err
	}

	return hschema.PublishAll(context.TODO(), c, kubeClass, []configv1.Plan{})
}
