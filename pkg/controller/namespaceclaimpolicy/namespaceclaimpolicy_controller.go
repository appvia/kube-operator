package namespaceclaimpolicy

import (
	"context"

	kubev1 "github.com/appvia/kube-operator/pkg/apis/kube/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_namespaceclaimpolicy")

// Add creates a new NamespaceClaimPolicy Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileNamespaceClaimPolicy{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("namespaceclaimpolicy-controller", mgr, controller.Options{
		MaxConcurrentReconciles: 10,
		Reconciler:              r,
	})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource NamespaceClaimPolicy
	err = c.Watch(&source.Kind{Type: &kubev1.NamespaceClaimPolicy{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	return nil
}

// blank assignment to verify that ReconcileNamespaceClaimPolicy implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileNamespaceClaimPolicy{}

// ReconcileNamespaceClaimPolicy reconciles a NamespaceClaimPolicy object
type ReconcileNamespaceClaimPolicy struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a NamespaceClaimPolicy object and makes changes based on the state read
// and what is in the NamespaceClaimPolicy.Spec
func (r *ReconcileNamespaceClaimPolicy) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues(
		"resource.namespace", request.Namespace,
		"resource.name", request.Name)

	// @step: fetch the NamespaceClaimPolicy instance
	instance := &kubev1.NamespaceClaimPolicy{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			return reconcile.Result{}, nil
		}
		return reconcile.Result{}, err
	}

	// Set NamespaceClaimPolicy instance as the owner and controller
	if err := controllerutil.SetControllerReference(instance, pod, r.scheme); err != nil {
		return reconcile.Result{}, err
	}

	return reconcile.Result{}, nil
}
