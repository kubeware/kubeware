/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	kubebindv1alpha1 "go.bytebuilders.dev/kube-bind/apis/kubebind/v1alpha1"
	versioned "go.bytebuilders.dev/kube-bind/client/clientset/versioned"
	internalinterfaces "go.bytebuilders.dev/kube-bind/client/informers/externalversions/internalinterfaces"
	v1alpha1 "go.bytebuilders.dev/kube-bind/client/listers/kubebind/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// APIServiceNamespaceInformer provides access to a shared informer and lister for
// APIServiceNamespaces.
type APIServiceNamespaceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.APIServiceNamespaceLister
}

type aPIServiceNamespaceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewAPIServiceNamespaceInformer constructs a new informer for APIServiceNamespace type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAPIServiceNamespaceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredAPIServiceNamespaceInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredAPIServiceNamespaceInformer constructs a new informer for APIServiceNamespace type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAPIServiceNamespaceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubeBindV1alpha1().APIServiceNamespaces(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubeBindV1alpha1().APIServiceNamespaces(namespace).Watch(context.TODO(), options)
			},
		},
		&kubebindv1alpha1.APIServiceNamespace{},
		resyncPeriod,
		indexers,
	)
}

func (f *aPIServiceNamespaceInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredAPIServiceNamespaceInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *aPIServiceNamespaceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&kubebindv1alpha1.APIServiceNamespace{}, f.defaultInformer)
}

func (f *aPIServiceNamespaceInformer) Lister() v1alpha1.APIServiceNamespaceLister {
	return v1alpha1.NewAPIServiceNamespaceLister(f.Informer().GetIndexer())
}