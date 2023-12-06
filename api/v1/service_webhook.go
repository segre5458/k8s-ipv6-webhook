/*
Copyright 2023.

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

package v1

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	corev1 "k8s.io/api/core/v1"
)

// log is for logging in this package.
var servicelog = logf.Log.WithName("service-resource")

// // SetupWebhookWithManager will setup the manager to manage the webhooks
// func (r *Service) SetupWebhookWithManager(mgr ctrl.Manager) error {
// 	return ctrl.NewWebhookManagedBy(mgr).
// 		For(r).
// 		Complete()
// }

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//+kubebuilder:webhook:path=/mutate-core-v1-service,mutating=true,failurePolicy=fail,sideEffects=None,groups=core,resources=services,verbs=create;update,versions=v1,name=mservice.kb.io,admissionReviewVersions=v1

type ServiceNetwork struct {}

var _ admission.CustomDefaulter = &ServiceNetwork{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (*ServiceNetwork) Default(ctx context.Context, obj runtime.Object) {
	service := obj.(*corev1.Service)
	servicelog.Info("default", "name", service.Name)

	// TODO(user): fill in your defaulting logic.
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-core-v1-service,mutating=false,failurePolicy=fail,sideEffects=None,groups=core,resources=services,verbs=create;update,versions=v1,name=vservice.kb.io,admissionReviewVersions=v1

// var _ webhook.Validator = &Service{}

// // ValidateCreate implements webhook.Validator so a webhook will be registered for the type
// func (r *Service) ValidateCreate() (admission.Warnings, error) {
// 	servicelog.Info("validate create", "name", r.Name)

// 	// TODO(user): fill in your validation logic upon object creation.
// 	return nil, nil
// }

// // ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
// func (r *Service) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
// 	servicelog.Info("validate update", "name", r.Name)

// 	// TODO(user): fill in your validation logic upon object update.
// 	return nil, nil
// }

// // ValidateDelete implements webhook.Validator so a webhook will be registered for the type
// func (r *Service) ValidateDelete() (admission.Warnings, error) {
// 	servicelog.Info("validate delete", "name", r.Name)

// 	// TODO(user): fill in your validation logic upon object deletion.
// 	return nil, nil
// }
