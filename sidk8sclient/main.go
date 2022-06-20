package main

import (
	"context"
	"flag"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "/home/siddharth/.kube/config", "location for your kubeconfig file")
	// fmt.Println(*kubeconfig)
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Printf("error %s building config from flags\n", err.Error())
		config, err = rest.InClusterConfig()
		if err != nil {
			fmt.Printf("error %s getting inclusterconfig", err.Error())
		}
	}
	// fmt.Printf("%+v", config)
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("err %s creating clientset", err.Error())
	}
	ctx := context.Background()
	pods, err := clientset.CoreV1().Pods("default").List(ctx, metav1.ListOptions{})
	if err != nil {
		fmt.Printf("err %s,while listing all the pods from kube-system namespace\n", err.Error())
	}
	fmt.Println("Print the number of Pod")
	i := 1
	for _, pod := range pods.Items {
		fmt.Printf("%d: -- %s\n", i, pod.Name)
		i++
	}

	deployment, err := clientset.AppsV1().Deployments("default").List(ctx, metav1.ListOptions{})
	if err != nil {
		fmt.Printf("error %s listing the deployments", err.Error())
	}
	fmt.Println("List of deployment")
	for _, d := range deployment.Items {
		fmt.Printf("%s\n", d.Name)
	}
}
