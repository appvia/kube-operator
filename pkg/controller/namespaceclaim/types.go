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

const (
	// ClusterRoleName is the name of the cluster role we will bind the
	// teams users to
	ClusterRoleName = "hub:system:ns-admin"
	// RoleBindingName is the name namespace role binding
	RoleBindingName = "hub:team"
	// FinalizerName is our finalizer name
	FinalizerName = "kube.hub.appvia.io"
)

const (
	// PhaseDeleting indicates the resource is being deleted
	PhaseDeleting = "Deleting"
	// PhaseInstalled indicates the resource was installed
	PhaseInstalled = "Installed"
	// PhaseFailure indicates some inreconcilaable situation
	PhaseFailure = "Failed"
)
