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
	time "time"

	versioned "github.com/automationbroker/broker-client-go/client/clientset/versioned"
	internalinterfaces "github.com/automationbroker/broker-client-go/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/automationbroker/broker-client-go/client/listers/automationbroker/v1alpha1"
	automationbroker_v1alpha1 "github.com/automationbroker/broker-client-go/pkg/apis/automationbroker/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// BundleBindingInformer provides access to a shared informer and lister for
// BundleBindings.
type BundleBindingInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.BundleBindingLister
}

type bundleBindingInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewBundleBindingInformer constructs a new informer for BundleBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewBundleBindingInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredBundleBindingInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredBundleBindingInformer constructs a new informer for BundleBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredBundleBindingInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AutomationbrokerV1alpha1().BundleBindings(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AutomationbrokerV1alpha1().BundleBindings(namespace).Watch(options)
			},
		},
		&automationbroker_v1alpha1.BundleBinding{},
		resyncPeriod,
		indexers,
	)
}

func (f *bundleBindingInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredBundleBindingInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *bundleBindingInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&automationbroker_v1alpha1.BundleBinding{}, f.defaultInformer)
}

func (f *bundleBindingInformer) Lister() v1alpha1.BundleBindingLister {
	return v1alpha1.NewBundleBindingLister(f.Informer().GetIndexer())
}
