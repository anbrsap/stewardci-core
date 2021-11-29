/*
#########################
#  SAP Steward-CI       #
#########################

THIS CODE IS GENERATED! DO NOT TOUCH!

Copyright SAP SE.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	time "time"

	versioned "github.com/SAP/stewardci-core/pkg/tektonclient/clientset/versioned"
	internalinterfaces "github.com/SAP/stewardci-core/pkg/tektonclient/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/SAP/stewardci-core/pkg/tektonclient/listers/pipeline/v1beta1"
	pipelinev1beta1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// PipelineRunInformer provides access to a shared informer and lister for
// PipelineRuns.
type PipelineRunInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.PipelineRunLister
}

type pipelineRunInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewPipelineRunInformer constructs a new informer for PipelineRun type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPipelineRunInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPipelineRunInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredPipelineRunInformer constructs a new informer for PipelineRun type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPipelineRunInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TektonV1beta1().PipelineRuns(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TektonV1beta1().PipelineRuns(namespace).Watch(context.TODO(), options)
			},
		},
		&pipelinev1beta1.PipelineRun{},
		resyncPeriod,
		indexers,
	)
}

func (f *pipelineRunInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPipelineRunInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *pipelineRunInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&pipelinev1beta1.PipelineRun{}, f.defaultInformer)
}

func (f *pipelineRunInformer) Lister() v1beta1.PipelineRunLister {
	return v1beta1.NewPipelineRunLister(f.Informer().GetIndexer())
}
