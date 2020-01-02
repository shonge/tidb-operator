// Copyright 2020 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/pingcap/tidb-operator/pkg/apis/pingcap/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// TidbMonitorLister helps list TidbMonitors.
type TidbMonitorLister interface {
	// List lists all TidbMonitors in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.TidbMonitor, err error)
	// TidbMonitors returns an object that can list and get TidbMonitors.
	TidbMonitors(namespace string) TidbMonitorNamespaceLister
	TidbMonitorListerExpansion
}

// tidbMonitorLister implements the TidbMonitorLister interface.
type tidbMonitorLister struct {
	indexer cache.Indexer
}

// NewTidbMonitorLister returns a new TidbMonitorLister.
func NewTidbMonitorLister(indexer cache.Indexer) TidbMonitorLister {
	return &tidbMonitorLister{indexer: indexer}
}

// List lists all TidbMonitors in the indexer.
func (s *tidbMonitorLister) List(selector labels.Selector) (ret []*v1alpha1.TidbMonitor, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TidbMonitor))
	})
	return ret, err
}

// TidbMonitors returns an object that can list and get TidbMonitors.
func (s *tidbMonitorLister) TidbMonitors(namespace string) TidbMonitorNamespaceLister {
	return tidbMonitorNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TidbMonitorNamespaceLister helps list and get TidbMonitors.
type TidbMonitorNamespaceLister interface {
	// List lists all TidbMonitors in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.TidbMonitor, err error)
	// Get retrieves the TidbMonitor from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.TidbMonitor, error)
	TidbMonitorNamespaceListerExpansion
}

// tidbMonitorNamespaceLister implements the TidbMonitorNamespaceLister
// interface.
type tidbMonitorNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all TidbMonitors in the indexer for a given namespace.
func (s tidbMonitorNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.TidbMonitor, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TidbMonitor))
	})
	return ret, err
}

// Get retrieves the TidbMonitor from the indexer for a given namespace and name.
func (s tidbMonitorNamespaceLister) Get(name string) (*v1alpha1.TidbMonitor, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("tidbmonitor"), name)
	}
	return obj.(*v1alpha1.TidbMonitor), nil
}
