package ncl

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jinshikoike/go-niftycloud/compute"
)

func resourceFireWall() *schema.Resource {
	return &schema.Resource{
		Create: resourceFireWallCreate,
		Read:   resourceFireWallRead,
		Update: resourceFireWallUpdate,
		Delete: resourceFireWallDelete,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
      "description": &schema.Schema{
				Type:     schema.TypeString,
        Description: "description",
				Optional: true,
			},
  	},
	}
}

func resourceFireWallCreate(d *schema.ResourceData, meta interface{}) error {
	nclClient := meta.(*NclClient)

  opts := compute.SecurityGroup{
    Name: d.Get("name").(string),
    Description: d.Get("description").(string),
  }

	resp, err := nclClient.CreateSecurityGroup(opts)
	if err != nil {
		return fmt.Errorf("Error completing tasks: %#v", err)
	}
	d.SetId(resp.Name)
	return nil
}

func resourceFireWallDelete(d *schema.ResourceData, meta interface{}) error {
	nclClient := meta.(*NclClient)
	if err := resourceFireWallRead(d, meta); err != nil {
		return err
	}

	groupName := d.Id()
	if groupName == "" {
		return nil
	}

  // TODO: Before delete security groups, you must remove applied server from securityg group apply

  opts := compute.SecurityGroup{
    Name: groupName,
  }
  _, deleteErr := nclClient.DeleteSecurityGroup(opts)
	if deleteErr != nil {
		return fmt.Errorf("Error Delete Security Group error: %s", deleteErr)
	}

	return resourceFireWallRead(d, meta)
}

func resourceFireWallUpdate(d *schema.ResourceData, meta interface{}) error {
  // nclClient := meta.(*NclClient)
	keyName := d.Id()
	if keyName == "" {
		return nil
	}

	return resourceFireWallRead(d, meta)
}

func resourceFireWallRead(d *schema.ResourceData, meta interface{}) error {
	nclClient := meta.(*NclClient)

	groupName := d.Id()
	if groupName == "" {
		groupName = d.Get("name").(string)
	}

  group := compute.SecurityGroup{
    Name: groupName,
  }

  resp, err := nclClient.SecurityGroups([]compute.SecurityGroup{group}, nil)
	if err != nil {
    return fmt.Errorf("Error SecurityGroups request : %#v", err)
  }

  if len(resp.Groups) <= 0 {
    d.SetId("")
    return nil
  }

  //fmt.Println(resp.Groups[0].IPPerms)
	d.Set("name", resp.Groups[0].Name)
  d.Set("description", resp.Groups[0].Description)

  d.SetId(resp.Groups[0].Name)

	return nil
}
