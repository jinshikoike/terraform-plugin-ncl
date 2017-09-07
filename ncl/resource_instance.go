package ncl

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jinshikoike/go-niftycloud/compute"
)

const noState = -1
const statePending = 0
const stateActive = 16
const stateStop = 80

const maxLoopCount = 20
const elapsedSecond = 10

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
			"force_destroy": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"code": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	nclClient := meta.(*NclClient)

	if d.Get("force_destroy").(bool) {
		return nil
	}

	securityGroupList := []compute.SecurityGroup{}
	if v, ok := d.GetOk("security_groups"); ok {
		securityGroups := v.([]interface{})
		for _, securityGroupSchema := range securityGroups {
			securityGroup := compute.SecurityGroup{Name: securityGroupSchema.(map[string]interface{})["name"].(string)}
			securityGroupList = append(securityGroupList, securityGroup)
		}
	}

	opts := compute.RunInstancesOptions{
		ImageId:               d.Get("image_id").(string),
		KeyName:               d.Get("key_name").(string),
		InstanceType:          d.Get("instance_type").(string),
		SecurityGroups:        securityGroupList,
		AvailZone:             d.Get("avail_zone").(string),
		AccountingType:        d.Get("accounting_type").(string),
		InstanceId:            d.Get("instance_id").(string),
		DisableAPITermination: false,
	}
	resp, err := nclClient.RunInstances(&opts)

	if err != nil {
		return fmt.Errorf("Error completing tasks: %#v", err)
	}
	//d.SetId(resp.Instances[0].InstanceId + "," + resp.Instances[0].InstanceUniqueId)
	// TODO: https://github.com/higebu/go-niftycloud/blob/master/compute/compute.go#L328
	// can not set false to DisableApiTermination

	//for loopCount := 0; loopCount <= 10; {
	//	result, _, _ := checkStatusCode(nclClient, instanceID, stateActive)
	//	if result {
	//		break
	//	}
	//	time.Sleep(10 * time.Second)
	//	loopCount++
	//}

	result := polling(nclClient, resp.Instances[0].InstanceId, stateActive)
	if result {
		d.Set("code", stateActive)
	}

	d.SetId(resp.Instances[0].InstanceId)
	return nil
}

func resourceInstanceDelete(d *schema.ResourceData, meta interface{}) error {
	nclClient := meta.(*NclClient)
	if err := resourceInstanceRead(d, meta); err != nil {
		return err
	}

	instanceID := d.Id()
	if instanceID == "" {
		return nil
	}

	instanceIsActive, currentState, _ := checkStatusCode(nclClient, instanceID, stateActive, stateStop)
	if !instanceIsActive {
		return nil
	}

	if currentState != stateStop {
		opts := compute.StopInstancesOptions{
			//Force: true,
			InstanceIds: []string{
				instanceID,
			},
		}

		_, respErr := nclClient.StopInstances(&opts)
		if respErr != nil {
			return fmt.Errorf("Error Stop Instance: %s", respErr)
		}

		// wait instance state change from pending to stop
		//for loopCount := 0; loopCount <= 10; {
		//	result, _, _ := checkStatusCode(nclClient, instanceID, stateStop)
		//	if result {
		//		break
		//	}
		//	time.Sleep(10 * time.Second)
		//	loopCount++
		//}
		polling(nclClient, instanceID, stateStop)
	}

	if d.Get("force_destroy").(bool) {
		_, terminateErr := nclClient.TerminateInstances([]string{instanceID})
		if terminateErr != nil {
			return fmt.Errorf("Error Terminate Instance: %s", terminateErr)
		}
	}

	return resourceInstanceRead(d, meta)
}

func resourceInstanceUpdate(d *schema.ResourceData, meta interface{}) error {
	nclClient := meta.(*NclClient)
	instanceID := d.Id()
	if instanceID == "" {
		return nil
	}

	result, _, _ := checkStatusCode(nclClient, instanceID, stateStop)

	if result {
		// Start Instance
		_, err := nclClient.StartInstances(instanceID)
		if err != nil {
			return fmt.Errorf("Error start Instances")
		}

		if polling(nclClient, instanceID, stateActive) {
			d.Set("code", stateActive)
		}

	}

	return resourceInstanceRead(d, meta)
}

func resourceInstanceRead(d *schema.ResourceData, meta interface{}) error {
	filePath := "/Users/shinji/go/src/github.com/jinshikoike/terraform-provider-ncl/examples/mylog"
	//filePath := "/home/koike/go/src/github.com/jinshikoike/terraform-provider-ncl/examples/mylog"
	file, err := os.Create(filePath)
	if err != nil {
		// Openエラー処理
	}
	defer file.Close()

	nclClient := meta.(*NclClient)

	instanceID := d.Id()
	if instanceID == "" {
		instanceID = d.Get("image_id").(string)
	}

	resp, err := nclClient.DescribeInstances([]string{instanceID}, nil)
	if err != nil {
		if err.Error() == "500 Internal Server Error" {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("DescribeInstances request error: %s", err)
	}

	if len(resp.Reservations) < 1 || len(resp.Reservations[0].Instances) < 1 {
		d.SetId("")
		return nil
	}

	instance := resp.Reservations[0].Instances[0]

	d.Set("image_id", instance.ImageId)
	d.Set("key_name", instance.KeyName)
	d.Set("instance_type", instance.InstanceType)
	d.Set("avail_zone", instance.AvailZone)
	d.Set("accounting_type", instance.AccountingType)
	d.Set("instance_id", instance.InstanceId)
	d.Set("code", instance.State.Code)
	data, _ := json.Marshal(resp.Reservations)
	file.Write(data)

	return nil
}

func checkStatusCode(client *NclClient, instanceID string, expectedState ...int) (bool, int, error) {
	resp, err := client.DescribeInstances([]string{instanceID}, nil)
	if err != nil {
		if err.Error() == "500 Internal Server Error" {
			// May be Instance is not found.
			return false, -1, nil
		}
		return false, -1, err
	}
	currentState := resp.Reservations[0].Instances[0].State.Code
	result := false

	for _, state := range expectedState {
		if currentState == state {
			result = true
		}
	}
	return result, currentState, nil
}

func polling(client *NclClient, instanceID string, expectedState ...int) (result bool) {
	for loopCount := 0; loopCount <= maxLoopCount; {
		result, _, _ = checkStatusCode(client, instanceID, expectedState...)
		if result {
			break
		}
		time.Sleep(elapsedSecond * time.Second)
		loopCount++
	}
	return
}
