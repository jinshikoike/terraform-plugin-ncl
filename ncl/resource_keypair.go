package ncl

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jinshikoike/go-niftycloud/compute"
)

func resourceKeyPair() *schema.Resource {
	return &schema.Resource{
		Create: resourceKeyPairCreate,
		Read:   resourceKeyPairRead,
		Update: resourceKeyPairUpdate,
		Delete: resourceKeyPairDelete,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"public_key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceKeyPairCreate(d *schema.ResourceData, meta interface{}) error {
	//nclClient := meta.(*NclClient)

	//opts := compute.KeyPair{
	//	Name:        d.Get("name").(string),
	//	Fingerprint: d.Get("finger_print").(string),
	//}
	//resp, err := nclClient.RunInstances(&opts)

	//if err != nil {
	//	return fmt.Errorf("Error completing tasks: %#v", err)
	//}

	//result, _ := polling(nclClient, resp.Instances[0].InstanceId, stateActive)
	//if result {
	//	d.Set("code", stateActive)
	//}

	//d.SetId(resp.Instances[0].InstanceId)
	return nil
}

func resourceKeyPairDelete(d *schema.ResourceData, meta interface{}) error {
	nclClient := meta.(*NclClient)
	if err := resourceKeyPairRead(d, meta); err != nil {
		return err
	}

	instanceID := d.Id()
	if instanceID == "" {
		return nil
	}
	return resourceKeyPairRead(d, meta)
	stateCheck, currentState, _ := checkStatusCode(nclClient, instanceID, stateActive, stateStop)
	if !stateCheck {
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

		stateCheck, currentState = polling(nclClient, instanceID, stateStop)

	}
	if stateCheck {
		d.Set("code", stateStop)
	}

	if d.Get("force_destroy").(bool) {
		_, terminateErr := nclClient.TerminateInstances([]string{instanceID})
		if terminateErr != nil {
			return fmt.Errorf("Error Terminate Instance: %s", terminateErr)
		}
	}
	return resourceKeyPairRead(d, meta)
}

func resourceKeyPairUpdate(d *schema.ResourceData, meta interface{}) error {
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

		if result, _ = polling(nclClient, instanceID, stateActive); result {
			d.SetId("")
		}
	}

	return resourceKeyPairRead(d, meta)
}

func resourceKeyPairRead(d *schema.ResourceData, meta interface{}) error {
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

	return nil
}
