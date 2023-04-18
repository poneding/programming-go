package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	var (
		kubeClient, _ = KubeClient()
	)

	// 直接从Api Server查询资源，造成Api Server压力。
	p, e := kubeClient.CoreV1().Pods("default").Get(context.TODO(), "nginx", v1.GetOptions{})
	if e != nil {
		fmt.Printf("get pod error: %s\n", e.Error())
		return
	}
	fmt.Printf("get pod success: %s\n", p.GetName())

	// 从本地存储中查询，并监听Api Server资源操作的事件，然后反馈到本地存储，并触发controller的回调函数。
	sharedInformerFactory := informers.NewSharedInformerFactory(kubeClient, time.Second*30)

	podInformer := sharedInformerFactory.Core().V1().Pods()
	podInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			p, ok := obj.(*corev1.Pod)
			if ok {
				fmt.Printf("pod [%s] added\n", p.GetName())
			}
		},
		DeleteFunc: func(obj interface{}) {
			p, ok := obj.(*corev1.Pod)
			if ok {
				fmt.Printf("pod [%s] deleted\n", p.GetName())
			}
		},
		//UpdateFunc: func(oldObj, newObj interface{}) {
		//	p, ok := oldObj.(corev1.Pod)
		//	if ok {
		//		fmt.Printf("pod [%s] updated", p.GetName())
		//	}
		//},
	})

	//stopCh := make(chan struct{})
	//sharedInformerFactory.Start(stopCh)
	//sharedInformerFactory.WaitForCacheSync(stopCh)
	sharedInformerFactory.Start(wait.NeverStop)
	sharedInformerFactory.WaitForCacheSync(wait.NeverStop)

	pods, _ := podInformer.Lister().Pods("default").List(labels.SelectorFromSet(labels.Set{
		"app": "nginx",
	}))
	fmt.Println("name=nginx pod count:", len(pods))
	for _, pod := range pods {
		fmt.Println("pod name:", pod.GetName())
	}

	nginxPod, err := podInformer.Lister().Pods("default").Get("nginx")
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}

	fmt.Println(nginxPod.GetName())

	<-wait.NeverStop
	//serviceLister := serviceInformer.Lister()
	//sListrSynced := serviceInformer.Informer().HasSynced()

}

// KubeConfig read kubernetes config from kubeconfig file
func KubeClient() (*kubernetes.Clientset, error) {
	kubeconfig := flag.String("kubeconfig", "C:/Users/dp/.kube/config", "location to your kubeconfig file")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Printf("error %s building config from flags\n", err.Error())
		config, err = rest.InClusterConfig()
		if err != nil {
			fmt.Printf("error %s getting inclusterconfig\n", err.Error())
			return nil, err
		}
	}
	return kubernetes.NewForConfig(config)
}
