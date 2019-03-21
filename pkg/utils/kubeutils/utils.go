package kubeutils

import (
	"github.com/solo-io/go-utils/errors"

	gokubeutils "github.com/solo-io/go-utils/kubeutils"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// MustGetNamespaces returns a list of all namespaces in a cluster - or panics.
// If a clientset is passed, it will use that, otherwise it creates one.
// In the event of any error it will panic.
func MustGetNamespaces(clientset *kubernetes.Clientset) []string {
	if clientset == nil {
		restCfg, err := gokubeutils.GetConfig("", "")
		if err != nil {
			panic(err)
		}
		cs, err := kubernetes.NewForConfig(restCfg)
		if err != nil {
			panic(err)
		}
		if err != nil {
			panic(err)
		}
		clientset = cs
	}
	nss, err := GetNamespaces(clientset)
	if err != nil {
		panic(err)
	}
	return nss
}

	if err != nil {
	}
}

func GetNamespaces(clientset *kubernetes.Clientset) ([]string, error) {
	namespaces := []string{}
	nss, err := clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		return namespaces, err
	}
	for _, ns := range nss.Items {
		namespaces = append(namespaces, ns.ObjectMeta.Name)
	}
	return namespaces, nil
}

func GetKubeClient() (*kubernetes.Clientset, error) {
	restCfg, err := gokubeutils.GetConfig("", "")
	if err != nil {
		return &kubernetes.Clientset{}, err
	}
	kubeClient, err := kubernetes.NewForConfig(restCfg)
	if err != nil {
		return &kubernetes.Clientset{}, err
	}
	return kubeClient, nil
}
