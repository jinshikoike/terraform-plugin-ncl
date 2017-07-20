package ncl

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/higebu/go-niftycloud/compute"
)

func resourceInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceInstanceCreate,
		//Delete: resourceVcdDNATDelete,
		//Read:   resourceVcdDNATRead,
		Schema: map[string]*schema.Schema{
			"image_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"key_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"instance_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"security_groups": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				//MiniItms: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"user_data": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"avail_zone": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"disable_api_termination": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"accounting_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"instance_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: false,
			},
			"admin": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"public_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"agreement": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

//ImageId               string
//	KeyName               string
//	InstanceType          string
//	SecurityGroups        []SecurityGroup
//	UserData              []byte
//	AvailZone             string
//	DisableAPITermination bool
//	AccountingType        string
//	InstanceId            string
//	Admin                 string
//	Password              string
//	IpType                string
//	PublicIp              string
//Agreement string

func resourceInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	nclClient := meta.(*NclClient)

	opts := compute.RunInstances{
		ImageId:      d.Get("image_id").(string),
		KeyName:      d.Get("key_name").(string),
		InstanceType: d.Get("instance_type").(string),
		//SecurityGroups: []compute.SecurityGroup{{Name: "examplegroup"}},
		AvailZone:      d.Get("avail_zone").(string),
		AccountingType: d.Get("accounting_type").(string),
	}
	resp, err := nclClient.RunInstances(&opts)
	if err != nil {
		return fmt.Errorf("Error completing tasks: %#v", err)
	}

	//	d.SetId(d.Get("external_ip").(string) + ":" + portString + " > " + d.Get("internal_ip").(string) + ":" + translatedPortString)
	return nil
}
