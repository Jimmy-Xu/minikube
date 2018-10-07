#!/bin/bash
export K8SRELEASE=v1.10.8

rm -fr temp
mkdir temp
cd temp

wget https://storage.googleapis.com/kubernetes-release/release/$K8SRELEASE/bin/linux/amd64/kubeadm
ossutil cp kubeadm oss://kubernetes/kubernetes-release/release/$K8SRELEASE/bin/linux/amd64/kubeadm
wget https://storage.googleapis.com/kubernetes-release/release/$K8SRELEASE/bin/linux/amd64/kubelet
ossutil cp kubelet oss://kubernetes/kubernetes-release/release/$K8SRELEASE/bin/linux/amd64/kubelet

wget https://storage.googleapis.com/kubernetes-release/release/$K8SRELEASE/bin/linux/amd64/kubeadm.sha1
ossutil cp kubeadm.sha1 oss://kubernetes/kubernetes-release/release/$K8SRELEASE/bin/linux/amd64/kubeadm.sha1
wget https://storage.googleapis.com/kubernetes-release/release/$K8SRELEASE/bin/linux/amd64/kubelet.sha1
ossutil cp kubelet.sha1 oss://kubernetes/kubernetes-release/release/$K8SRELEASE/bin/linux/amd64/kubelet.sha1

wget https://storage.googleapis.com/kubernetes-release/release/stable-1.txt
ossutil cp -f stable-1.txt oss://kubernetes/kubernetes-release/release/stable-1.txt
cd ..
rm -fr temp

