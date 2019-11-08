package controller

import (
	"github.com/appvia/kube-operator/pkg/controller/namespaceclaim"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, namespaceclaim.Add)
}
