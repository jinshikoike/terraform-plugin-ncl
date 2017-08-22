package ncl

import (
	"fmt"
	"strings"

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
		regionNameList := []string{}

		for k := range niftycloud.Regions {
			regionNameList = append(regionNameList, k)
		}
		return nil, fmt.Errorf("[Err] No Region Name for NiftyCloud, You Can Choose one from ", strings.Join(regionNameList, ","))
	}

	region := niftycloud.Regions[c.Region]
	auth, authError := niftycloud.GetAuth(c.AccessKey, c.SecretKey)

	if authError != nil {
		return nil, fmt.Errorf("[Err] auth error")
	}

	client := &NclClient{compute.New(auth, region)}

	return client, nil
}
