/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	model "github.com/upbound/provider-gcp/internal/controller/mlengine/model"
)

// Setup_mlengine creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_mlengine(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		model.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
