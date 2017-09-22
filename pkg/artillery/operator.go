package artillery

import (
	"github.com/Sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
)

type Operator struct {
	kclient kubernetes.Interface
}

func (c *Operator) Run(stopc <-chan struct{}) error {
	logrus.Info("Operator is up and running...")
	<-stopc
	return nil
}
