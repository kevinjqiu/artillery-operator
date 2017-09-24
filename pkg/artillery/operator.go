package artillery

import (
	"github.com/Sirupsen/logrus"
	"github.com/kevinjqiu/artillery-operator/pkg/client/v1alpha1"
	"github.com/kevinjqiu/artillery-operator/pkg/k8sutil"
	"k8s.io/client-go/kubernetes"
)

type Operator struct {
	kclient   kubernetes.Interface
	artclient v1alpha1.ArtilleryV1alpha1Interface
}

func createTPRs(ao *Operator) error {
	return nil
}

func createCRDs(ao *Operator) error {
	return nil
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
	artclient, err := v1alpha1.NewForConfig(k8scfg)
	if err != nil {
		return nil, err
	}

	op := &Operator{
		kclient:   k8sclient,
		artclient: artclient,
	}

	return op, nil
}

func (c *Operator) Run(stopc <-chan struct{}) error {
	logrus.Info("Operator is up and running...")
	<-stopc
	return nil
}
