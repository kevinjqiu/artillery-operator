package artillery

import (
	"k8s.io/client-go/rest"
)

type Config struct {
	Host        string
	TLSConfig   rest.TLSClientConfig
	TLSInsecure bool
}
