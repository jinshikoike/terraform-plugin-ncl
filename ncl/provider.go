package ncl

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"access_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"NIFTY_CLOUD_ACCESS_KEY",
				}, nil),
				Description: "NiftyCloud API Key ",
			},
			"secret_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"NIFTY_CLOUD_SECRET_KEY",
				}, nil),
				Description: "NiftyCloud API Key ",
			},
			"region": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"NIFTY_CLOUD_REGION",
				}, nil),
				Description: "NiftyCloud Region",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"ncl_instance": resourceInstance(),
      "ncl_keypair": resourceKeyPair(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := &Config{
		AccessKey: d.Get("access_key").(string),
		SecretKey: d.Get("secret_key").(string),
		Region:    d.Get("region").(string),
	}
	client, err := config.Client()
	return client, err
}
