package artillery

import (
	"github.com/Sirupsen/logrus"
	"github.com/kevinjqiu/artillery-operator/pkg/k8sutil"
	"k8s.io/client-go/kubernetes"
)

type Operator struct {
	kclient kubernetes.Interface
}

func New(config Config) (*Operator, error) {
	k8scfg, err := k8sutil.NewClusterConfig(config.Host, config.TLSInsecure, &config.TLSConfig)
	if err != nil {
		return nil, err
	}
	k8sclient, err := kubernetes.NewForConfig(k8scfg)
	if err != nil {
		return nil, err
	}

	op := &Operator{
		kclient: k8sclient,
	}

	return op, nil
}

func (c *Operator) Run(stopc <-chan struct{}) error {
	logrus.Info("Operator is up and running...")
	<-stopc
	return nil
}
