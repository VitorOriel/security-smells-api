package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
)

// var manifestURL = "https://raw.githubusercontent.com/infracloudio/botkube/develop/deploy-all-in-one.yaml"
var manifestURL = "https://github.com/fission/fission/releases/download/1.12.0/fission-all-1.12.0.yaml"

func main() {
	resp, err := http.Get(manifestURL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	manifests := string(data)
	fmt.Printf("Workloads:\n\n")
	decode := scheme.Codecs.UniversalDeserializer().Decode
	for _, spec := range strings.Split(manifests, "---") {
		if len(spec) == 0 {
			continue
		}
		obj, _, err := decode([]byte(spec), nil, nil)
		if err != nil {
			continue
		}
		switch obj.(type) {
		case *corev1.Pod:
			pods := obj.(*corev1.Pod)
			fmt.Println("Name:", pods.GetName())
			fmt.Println("Namespace:", pods.GetNamespace())
			fmt.Println("Kind:", pods.GetResourceVersion())
			fmt.Println("---")

		case *appsv1.Deployment:
			d := obj.(*appsv1.Deployment)
			fmt.Println("Name:", d.GetName())
			fmt.Println("Namespace:", d.GetNamespace())
			fmt.Println("GVK:", d.GroupVersionKind())
			fmt.Println("Containers IMAGEMS", d.Spec.Template.Spec.Containers[0].Image)
			fmt.Println("---")

		case *appsv1.StatefulSet:
			ss := obj.(*appsv1.StatefulSet)
			fmt.Println("Name:", ss.GetName())
			fmt.Println("Namespace:", ss.GetNamespace())
			fmt.Println("GVK:", ss.GroupVersionKind())
			fmt.Println("---")

		case *appsv1.DaemonSet:
			ds := obj.(*appsv1.DaemonSet)
			fmt.Println("Name:", ds.GetName())
			fmt.Println("Namespace:", ds.GetNamespace())
			fmt.Println("GVK:", ds.GroupVersionKind())
			fmt.Println("---")
		}
	}
}
