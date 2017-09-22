package artillery

import "k8s.io/client-go/kubernetes"

type Operator struct {
	kclient kubernetes.Interface
}
