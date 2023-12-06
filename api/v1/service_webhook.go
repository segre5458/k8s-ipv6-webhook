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

//+kubebuilder:webhook:path=/mutate--v1-service,mutating=true,failurePolicy=fail,sideEffects=None,groups=core,resources=services,verbs=create;update,versions=v1,name=mutate.service.webhook.marokiki.net,admissionReviewVersions=v1

type ServiceNetwork struct{}

var _ admission.CustomDefaulter = &ServiceNetwork{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (*ServiceNetwork) Default(ctx context.Context, obj runtime.Object) error {
	service := obj.(*corev1.Service)
	servicelog.Info("default", "name", service.Name)

	if service.Spec.IPFamilyPolicy == nil {
		service.Spec.IPFamilyPolicy = new(corev1.IPFamilyPolicy)
	}
	*service.Spec.IPFamilyPolicy = corev1.IPFamilyPolicyPreferDualStack
	*&service.Spec.IPFamilies = []corev1.IPFamily{corev1.IPv4Protocol, corev1.IPv6Protocol}

	return nil
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate--v1-service,mutating=false,failurePolicy=fail,sideEffects=None,groups=core,resources=services,verbs=create;update,versions=v1,name=validate.service.webhook.marokiki.net,admissionReviewVersions=v1

var _ admission.CustomValidator = &ServiceNetwork{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (*ServiceNetwork) ValidateCreate(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	service := obj.(*corev1.Service)
	servicelog.Info("validate create", "name", service.Name)

	return nil, nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (*ServiceNetwork) ValidateUpdate(ctx context.Context, oldObj, newObj runtime.Object) (admission.Warnings, error) {
	service := newObj.(*corev1.Service)
	servicelog.Info("validate update", "name", service.Name)

	return nil, nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (*ServiceNetwork) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	return nil, nil
}
