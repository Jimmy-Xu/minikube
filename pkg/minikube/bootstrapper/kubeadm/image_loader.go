/*
Copyright 2016 The Kubernetes Authors All rights reserved.

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

package kubeadm

import "fmt"

var imageMap = map[string]string{
	"k8s.gcr.io/pause-amd64:3.1":                    "registry.cn-hangzhou.aliyuncs.com/google_containers/pause-amd64:3.1",
	"k8s.gcr.io/k8s-dns-dnsmasq-nanny-amd64:1.14.8": "registry.cn-hangzhou.aliyuncs.com/google_containers/k8s-dns-dnsmasq-nanny-amd64:1.14.8",
	"k8s.gcr.io/k8s-dns-sidecar-amd64:1.14.8":       "registry.cn-hangzhou.aliyuncs.com/google_containers/k8s-dns-sidecar-amd64:1.14.8",
	"k8s.gcr.io/k8s-dns-kube-dns-amd64:1.14.8":      "registry.cn-hangzhou.aliyuncs.com/google_containers/k8s-dns-kube-dns-amd64:1.14.8",
}

func loadImageScripts() string {

	var scripts string

	for src, dest := range imageMap {

		script := fmt.Sprintf("sudo docker pull %s && sudo docker tag %s %s && ", dest, dest, src)
		scripts += script

	}

	return scripts
}
