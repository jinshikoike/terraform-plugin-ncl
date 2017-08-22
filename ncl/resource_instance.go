package ncl

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/higebu/go-niftycloud/compute"
)

func resourceInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceInstanceCreate,
		Read:   resourceInstanceRead,
		Update: resourceInstanceUpdate,
		Delete: resourceInstanceDelete,
		Schema: map[string]*schema.Schema{
			"image_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				//ForceNew: true,
			},
			"key_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"instance_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"avail_zone": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"security_groups": &schema.Schema{ // TypeString is better to understand.
				Type:     schema.TypeList,
				Required: true,
				MinItems: 1,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"user_data": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
				Optional: true,
				ValidateFunc: func(data interface{}, name string) (ws []string, errors []error) {
					input_data := data.(string)
					if len(input_data) > 15 {
						errors = append(errors, fmt.Errorf("instance_id length must be <= 15"))
					}
					return
				},
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

func resourceInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	nclClient := meta.(*NclClient)

	securityGroupList := []compute.SecurityGroup{}
	if v, ok := d.GetOk("security_groups"); ok {
		securityGroups := v.([]interface{})
		for _, securityGroupSchema := range securityGroups {
			securityGroup := compute.SecurityGroup{Name: securityGroupSchema.(map[string]interface{})["name"].(string)}
			securityGroupList = append(securityGroupList, securityGroup)
		}
	}

	opts := compute.RunInstancesOptions{
		ImageId:        d.Get("image_id").(string),
		KeyName:        d.Get("key_name").(string),
		InstanceType:   d.Get("instance_type").(string),
		SecurityGroups: securityGroupList,
		AvailZone:      d.Get("avail_zone").(string),
		AccountingType: d.Get("accounting_type").(string),
		InstanceId:     d.Get("instance_id").(string),
	}
	_, err := nclClient.RunInstances(&opts)

	if err != nil {
		return fmt.Errorf("Error completing tasks: %#v", err)
	}

	//	d.SetId(d.Get("external_ip").(string) + ":" + portString + " > " + d.Get("internal_ip").(string) + ":" + translatedPortString)
	return nil
}
func resourceInstanceDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
func resourceInstanceUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}
func resourceInstanceRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}
