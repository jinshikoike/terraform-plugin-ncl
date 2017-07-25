package ncl

import (
	"github.com/hashicorp/terraform/helper/schema"
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
			//	"key_name": &schema.Schema{
			//		Type:     schema.TypeString,
			//		Required: true,
			//	},
			//	"instance_type": &schema.Schema{
			//		Type:     schema.TypeString,
			//		Required: true,
			//	},
			//	"avail_zone": &schema.Schema{
			//		Type:     schema.TypeString,
			//		Required: true,
			//	},
			//			"security_groups": &schema.Schema{
			//				Type: schema.TypeList,
			//				//Required: true,
			//				Optional: true,
			//				//MiniItms: 1,
			//				Elem: &schema.Resource{
			//					Schema: map[string]*schema.Schema{
			//						"id": &schema.Schema{
			//							Type:     schema.TypeString,
			//							Required: true,
			//						},
			//						"name": &schema.Schema{
			//							Type:     schema.TypeString,
			//							Optional: true,
			//						},
			//					},
			//				},
			//			},
			//			"user_data": &schema.Schema{
			//				Type:     schema.TypeString,
			//				Optional: true,
			//			},
			//			"disable_api_termination": &schema.Schema{
			//				Type:     schema.TypeBool,
			//				Optional: true,
			//				Default:  false,
			//			},
			//			"accounting_type": &schema.Schema{
			//				Type:     schema.TypeString,
			//				Required: true,
			//			},
			//			"instance_id": &schema.Schema{
			//				Type:     schema.TypeString,
			//				Required: false,
			//			},
			//			"admin": &schema.Schema{
			//				Type:     schema.TypeString,
			//				Optional: true,
			//			},
			//			"password": &schema.Schema{
			//				Type:     schema.TypeString,
			//				Optional: true,
			//			},
			//			"ip_type": &schema.Schema{
			//				Type:     schema.TypeString,
			//				Optional: true,
			//			},
			//			"public_ip": &schema.Schema{
			//				Type:     schema.TypeString,
			//				Optional: true,
			//			},
			//			"agreement": &schema.Schema{
			//				Type:     schema.TypeString,
			//				Optional: true,
			//			},
		},
	}
}

func resourceInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	//nclClient := meta.(*NclClient)

	//opts := compute.RunInstancesOptions{
	//	ImageId:      d.Get("image_id").(string),
	//	KeyName:      d.Get("key_name").(string),
	//	InstanceType: d.Get("instance_type").(string),
	//	//SecurityGroups: []compute.SecurityGroup{{Name: "examplegroup"}},
	//	AvailZone:      d.Get("avail_zone").(string),
	//	AccountingType: d.Get("accounting_type").(string),
	//}
	//_, err := nclClient.RunInstances(&opts)

	//if err != nil {
	//	return fmt.Errorf("Error completing tasks: %#v", err)
	//}

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
