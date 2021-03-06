/*
Copyright (c) 2018 Red Hat, Inc.

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

// This file was automatically generated by informer-gen

package v1alpha1

import (
	internalinterfaces "github.com/automationbroker/broker-client-go/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Bundles returns a BundleInformer.
	Bundles() BundleInformer
	// BundleBindings returns a BundleBindingInformer.
	BundleBindings() BundleBindingInformer
	// BundleInstances returns a BundleInstanceInformer.
	BundleInstances() BundleInstanceInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Bundles returns a BundleInformer.
func (v *version) Bundles() BundleInformer {
	return &bundleInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// BundleBindings returns a BundleBindingInformer.
func (v *version) BundleBindings() BundleBindingInformer {
	return &bundleBindingInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// BundleInstances returns a BundleInstanceInformer.
func (v *version) BundleInstances() BundleInstanceInformer {
	return &bundleInstanceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
