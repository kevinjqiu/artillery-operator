package v1alpha1

import (
	"k8s.io/client-go/pkg/api"
	"k8s.io/client-go/pkg/api/unversioned"
	"k8s.io/client-go/pkg/runtime/serializer"
	"k8s.io/client-go/rest"
)

const (
	Group   = "artillery.idempotent.ca"
	Version = "v1alpha1"
)

type ArtilleryV1alpha1Interface interface {
	RESTClient() rest.Interface
}

type ArtilleryV1alpha1Client struct {
	restClient rest.Interface
}

func (ac *ArtilleryV1alpha1Client) RESTClient() rest.Interface {
	return ac.restClient
}

func NewForConfig(c *rest.Config) (*ArtilleryV1alpha1Client, error) {
	config := *c
	config.GroupVersion = &unversioned.GroupVersion{
		Group:   Group,
		Version: Version,
	}
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: api.Codecs}

	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}

	return &ArtilleryV1alpha1Client{client}, nil
}
