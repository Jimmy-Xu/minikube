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

package options

import (
	"runtime"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kubernetes/pkg/kubelet/config"
	kubetypes "k8s.io/kubernetes/pkg/kubelet/types"
)

const (
	// When these values are updated, also update test/e2e/framework/util.go
	defaultPodSandboxImageName    = "registry.cn-hangzhou.aliyuncs.com/google_containers/pause"
	defaultPodSandboxImageVersion = "3.1"
	// From pkg/kubelet/rkt/rkt.go to avoid circular import
	defaultRktAPIServiceEndpoint = "localhost:15441"
)

var (
	defaultPodSandboxImage = defaultPodSandboxImageName +
		"-" + runtime.GOARCH + ":" +
		defaultPodSandboxImageVersion
)

// NewContainerRuntimeOptions will create a new ContainerRuntimeOptions with
// default values.
func NewContainerRuntimeOptions() *config.ContainerRuntimeOptions {
	dockerEndpoint := ""
	if runtime.GOOS != "windows" {
		dockerEndpoint = "unix:///var/run/docker.sock"
	}

	return &config.ContainerRuntimeOptions{
		ContainerRuntime:          kubetypes.DockerContainerRuntime,
		DockerEndpoint:            dockerEndpoint,
		DockershimRootDirectory:   "/var/lib/dockershim",
		DockerDisableSharedPID:    true,
		PodSandboxImage:           defaultPodSandboxImage,
		ImagePullProgressDeadline: metav1.Duration{Duration: 1 * time.Minute},
		RktAPIEndpoint:            defaultRktAPIServiceEndpoint,
		ExperimentalDockershim:    false,
	}
}
