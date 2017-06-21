package log

import (

	"github.com/Nivenly/kamp/local"
	"github.com/Nivenly/kamp/k8s"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/pkg/api/v1"
)


func GetLogs(local *local.KampConfig, namespace string) error {
	client, err := k8s.LoadClient()
	if err != nil {
		return err
	}
	listOpts := v1.ListOptions{
		LabelSelector: "kamp=" + local.ProjectName,
	}
	pods, err := client.CoreV1().Pods(namespace).List(listOpts)
	if err != nil {
		return err
	}
	for _, pod := range pods.Items{
		go tailPod(pod, client)
	}
	return nil
}

func tailPod(p v1.Pod, cli *kubernetes.Clientset) {
	cli.CoreV1().Pods(p.Namespace).GetLogs(p.Name, &v1.PodLogOptions{Follow: true})
}
