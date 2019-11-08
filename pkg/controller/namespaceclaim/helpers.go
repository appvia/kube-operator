/*
 * Copyright (C) 2019  Rohith Jayawardene <gambol99@gmail.com>
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

package namespaceclaim

import (
	"context"
	"fmt"

	clusterv1 "github.com/appvia/hub-apis/pkg/apis/clusters/v1"
	orgv1 "github.com/appvia/hub-apis/pkg/apis/org/v1"
	kubev1 "github.com/appvia/kube-operator/pkg/apis/kube/v1"
	k8s "github.com/gambol99/hub-utils/pkg/kubernetes"

	kerrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// HubNamespaceType returns a reference to the hub namespaced type
func HubNamespaceType(name string) types.NamespacedName {
	return types.NamespacedName{
		Name:      name,
		Namespace: "hub",
	}
}

// IsTeam checks the team exists in the hub
// @TODO we need to bring these methods together under a single repo - but i cna't be asked at the
// moment, while so many things are influx
func IsTeam(ctx context.Context, c client.Client, name string) (*orgv1.Team, bool, error) {
	team := &orgv1.Team{}

	if err := c.Get(ctx, HubNamespaceType(name), team); err != nil {
		if kerrors.IsNotFound(err) {
			return nil, false, nil
		}

		return nil, false, err
	}

	return team, true, nil
}

// MakeTeamMembersList is responsible for retrieving the teams members
func MakeTeamMembersList(ctx context.Context, c client.Client, team string) (*orgv1.TeamMembershipList, error) {
	users := &orgv1.TeamMembershipList{}
	namespace := fmt.Sprintf("team-%s", team)

	if err := c.List(ctx, users, client.InNamespace(namespace)); err != nil {
		return users, err
	}

	return users, nil
}

// MakeClusterKubeClient is responisble for creating a kubernetes client from a cluster reference
func MakeClusterKubeClient(ctx context.Context, c client.Client, reference types.NamespacedName) (kubernetes.Interface, error) {
	// @step: retrieve the cluster resource
	cluster := &clusterv1.Kubernetes{}
	if err := c.Get(ctx, reference, cluster); err != nil {
		return nil, err
	}

	// @step: create a kubernetes client to this cluster
	return k8s.NewFromToken(cluster.Spec.Endpoint, cluster.Spec.Token, "")
}

// NamespaceClaimsToRequests converts a collection of claims to requests
func NamespaceClaimsToRequests(items []kubev1.NamespaceClaim) []reconcile.Request {
	requests := make([]reconcile.Request, len(items))

	// @step: trigger the namespaceclaims to reconcile
	for i := 0; i < len(items); i++ {
		requests[i] = reconcile.Request{
			NamespacedName: types.NamespacedName{
				Name:      items[i].Name,
				Namespace: items[i].Namespace,
			},
		}
	}

	return requests
}

// ListTeamNamespaceClaims returns a list of namespaceclaims in a namespace (team in our case)
func ListTeamNamespaceClaims(ctx context.Context, c client.Client, team string) ([]kubev1.NamespaceClaim, error) {
	list := &kubev1.NamespaceClaimList{}

	labels := map[string]string{
		"hub.appvia.io/team": team,
	}

	if err := c.List(ctx, list, client.MatchingLabels(labels), client.InNamespace("")); err != nil {
		return []kubev1.NamespaceClaim{}, err
	}

	return list.Items, nil
}
