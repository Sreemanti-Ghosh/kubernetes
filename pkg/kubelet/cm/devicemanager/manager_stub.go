/*
Copyright 2017 The Kubernetes Authors.

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

package devicemanager

import (
	"k8s.io/api/core/v1"
	pluginapi "k8s.io/kubernetes/pkg/kubelet/apis/deviceplugin/v1beta1"
	podresourcesapi "k8s.io/kubernetes/pkg/kubelet/apis/podresources/v1alpha1"
	"k8s.io/kubernetes/pkg/kubelet/cm/topologymanager"
	"k8s.io/kubernetes/pkg/kubelet/config"
	"k8s.io/kubernetes/pkg/kubelet/lifecycle"
	"k8s.io/kubernetes/pkg/kubelet/util/pluginwatcher"
	schedulernodeinfo "k8s.io/kubernetes/pkg/scheduler/nodeinfo"
)

// ManagerStub provides a simple stub implementation for the Device Manager.
type ManagerStub struct{}

// NewManagerStub creates a ManagerStub.
func NewManagerStub() (*ManagerStub, error) {
	return &ManagerStub{}, nil
}

// Start simply returns nil.
func (h *ManagerStub) Start(activePods ActivePodsFunc, sourcesReady config.SourcesReady) error {
	return nil
}

// Stop simply returns nil.
func (h *ManagerStub) Stop() error {
	return nil
}

// Devices returns an empty map.
func (h *ManagerStub) Devices() map[string][]pluginapi.Device {
	return make(map[string][]pluginapi.Device)
}

// Allocate simply returns nil.
func (h *ManagerStub) Allocate(node *schedulernodeinfo.NodeInfo, attrs *lifecycle.PodAdmitAttributes) error {
	return nil
}

// GetDeviceRunContainerOptions simply returns nil.
func (h *ManagerStub) GetDeviceRunContainerOptions(pod *v1.Pod, container *v1.Container) (*DeviceRunContainerOptions, error) {
	return nil, nil
}

// GetCapacity simply returns nil capacity and empty removed resource list.
func (h *ManagerStub) GetCapacity() (v1.ResourceList, v1.ResourceList, []string) {
	return nil, nil, []string{}
}

// GetWatcherHandler returns plugin watcher interface
func (h *ManagerStub) GetWatcherHandler() pluginwatcher.PluginHandler {
	return nil
}

// GetTopologyHints returns an empty TopologyHints struct
func (h *ManagerStub) GetTopologyHints(pod v1.Pod, container v1.Container) topologymanager.TopologyHints {
	return topologymanager.TopologyHints{}
}

// GetDevices returns nil
func (h *ManagerStub) GetDevices(_, _ string) []*podresourcesapi.ContainerDevices {
	return nil
}
