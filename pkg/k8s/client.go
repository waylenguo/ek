package k8s

import (
	"context"
	"flag"
	"fmt"
	set "github.com/ek/pkg/util"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
)

type Client struct {
	*kubernetes.Clientset
}

func NewClient() *Client {

	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return &Client{
		clientset,
	}
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

func (k8sClient *Client) FetchImages(namespace string) []string {
	pods, err := k8sClient.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	images := set.New()

	for _, item := range pods.Items {
		containers := item.Spec.Containers
		initContainers := item.Spec.InitContainers

		for _, container := range containers {
			image := container.Image
			images.Add(image)
		}

		for _, container := range initContainers {
			image := container.Image
			images.Add(image)
		}
	}

	return images.List();
}

func (k8sClient *Client) FetchPods(namespace string) *v1.PodList {
	pods, err := k8sClient.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	return pods
}

func (k8sClient *Client) FetchNamespaces() *v1.NamespaceList {
	namespaces, err := k8sClient.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	return namespaces
}

func (k8sClient *Client) DescribePod(pod v1.Pod) {
	name := pod.GetName()
	namespace := pod.GetNamespace()
	labels := pod.GetLabels()
	annotations := pod.GetAnnotations()

	fmt.Println(name, namespace, labels, annotations)

}