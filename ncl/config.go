package ncl

import (
	"fmt"

	"github.com/higebu/go-niftycloud/compute"
	"github.com/higebu/go-niftycloud/niftycloud"
)

type Config struct {
	AccessKey string
	SecretKey string
	Region    string
}

type NclClient struct {
	*compute.Compute
	// TODO: Mutex?
}

func (c *Config) Client() (*NclClient, error) {
	if c.AccessKey == "" {
		return nil, fmt.Errorf("[Err] No Access key for NiftyCloud")
	}

	if c.SecretKey == "" {
		return nil, fmt.Errorf("[Err] No Secret key for NiftyCloud")
	}

	if c.Region == "" {
		return nil, fmt.Errorf("[Err] No Region Name for NiftyCloud")
	}

	region := niftycloud.Regions[c.Region]
	auth, authError := niftycloud.GetAuth(c.AccessKey, c.SecretKey)

	if authError != nil {
		return authError, fmt.Errorf("[Err] auth error")
	}

	client := &NclClient{compute.New(auth, region)}

	return client, nil
}
