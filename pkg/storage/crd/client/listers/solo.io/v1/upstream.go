/*
Copyright 2018 The Kubernetes Authors.

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

// This file was automatically generated by lister-gen

package v1

import (
	v1 "github.com/solo-io/glue-storage/pkg/storage/crd/solo.io/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// UpstreamLister helps list Upstreams.
type UpstreamLister interface {
	// List lists all Upstreams in the indexer.
	List(selector labels.Selector) (ret []*v1.Upstream, err error)
	// Upstreams returns an object that can list and get Upstreams.
	Upstreams(namespace string) UpstreamNamespaceLister
	UpstreamListerExpansion
}

// upstreamLister implements the UpstreamLister interface.
type upstreamLister struct {
	indexer cache.Indexer
}

// NewUpstreamLister returns a new UpstreamLister.
func NewUpstreamLister(indexer cache.Indexer) UpstreamLister {
	return &upstreamLister{indexer: indexer}
}

// List lists all Upstreams in the indexer.
func (s *upstreamLister) List(selector labels.Selector) (ret []*v1.Upstream, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Upstream))
	})
	return ret, err
}

// Upstreams returns an object that can list and get Upstreams.
func (s *upstreamLister) Upstreams(namespace string) UpstreamNamespaceLister {
	return upstreamNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// UpstreamNamespaceLister helps list and get Upstreams.
type UpstreamNamespaceLister interface {
	// List lists all Upstreams in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.Upstream, err error)
	// Get retrieves the Upstream from the indexer for a given namespace and name.
	Get(name string) (*v1.Upstream, error)
	UpstreamNamespaceListerExpansion
}

// upstreamNamespaceLister implements the UpstreamNamespaceLister
// interface.
type upstreamNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Upstreams in the indexer for a given namespace.
func (s upstreamNamespaceLister) List(selector labels.Selector) (ret []*v1.Upstream, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Upstream))
	})
	return ret, err
}

// Get retrieves the Upstream from the indexer for a given namespace and name.
func (s upstreamNamespaceLister) Get(name string) (*v1.Upstream, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("upstream"), name)
	}
	return obj.(*v1.Upstream), nil
}
