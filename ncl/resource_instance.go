package ncl

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

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
	resp, err := nclClient.RunInstances(&opts)

	if err != nil {
		return fmt.Errorf("Error completing tasks: %#v", err)
	}
	d.SetId(resp.Instances[0].InstanceId + "," + resp.Instances[0].InstanceUniqueId)
	//	d.Get("external_ip").(string) + ":" + portString + " > " + d.Get("internal_ip").(string) + ":" + translatedPortString)
	return nil
}

func resourceInstanceDelete(d *schema.ResourceData, meta interface{}) error {

	return nil
}

func resourceInstanceUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceInstanceRead(d *schema.ResourceData, meta interface{}) error {
	filePath := "/home/koike/go/src/github.com/jinshikoike/terraform-provider-ncl/examples/mylog"
	file, err := os.Create(filePath)
	if err != nil {
		// Openエラー処理
	}
	defer file.Close()
	nclClient := meta.(*NclClient)
	terraformId := d.Id()

	if len(strings.Split(terraformId, ",")) < 2 {
		file.Write(([]byte)("strings split < 2 ; " + terraformId + "\n"))

		d.SetId("")
		return nil
	}

	instanceId := strings.Split(terraformId, ",")[0]
	instanceUniqueId := strings.Split(terraformId, ",")[1]

	resp, err := nclClient.DescribeInstances([]string{instanceId}, nil)
	if err != nil {
		return fmt.Errorf("DescribeInstances request error: %s", err)
	}

	var instance *compute.Instance
	_ = instance
	for _, reservation := range resp.Reservations {
		for _, instanceItem := range reservation.Instances {
			if instanceItem.InstanceUniqueId == instanceUniqueId {
				instance = &instanceItem
				break
			}
		}
	}

	if instance == nil {
		file.Write(([]byte)("instance is nil"))
		d.SetId("")
		return nil
	}

	d.Set("image_id", instance.ImageId)
	d.Set("key_name", instance.KeyName)
	d.Set("instance_type", instance.InstanceType)
	d.Set("avail_zone", instance.AvailZone)
	d.Set("accounting_type", instance.AccountingType)
	d.Set("instance_id", instance.InstanceId)

	data, _ := json.Marshal(resp.Reservations)
	file.Write(data)

	return nil
}
