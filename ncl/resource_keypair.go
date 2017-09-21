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
			"public_key_material": &schema.Schema{
				Type:     schema.TypeString,
        Description: "public key string",
				Required: true,
			},
      "description": &schema.Schema{
				Type:     schema.TypeString,
        Description: "description",
				Optional: true,
			},
  	},
	}
}

func resourceKeyPairCreate(d *schema.ResourceData, meta interface{}) error {
	nclClient := meta.(*NclClient)

  opts := compute.ImportKeyOpts{
    KeyName: d.Get("name").(string),
    PublicKeyMaterial: d.Get("public_key_material").(string),
    Description: d.Get("description").(string),
  }

	resp, err := nclClient.ImportKeyPair(&opts)
	if err != nil {
		return fmt.Errorf("Error completing tasks: %#v", err)
	}
	d.SetId(resp.KeyName)
	return nil
}

func resourceKeyPairDelete(d *schema.ResourceData, meta interface{}) error {
	nclClient := meta.(*NclClient)
	if err := resourceKeyPairRead(d, meta); err != nil {
		return err
	}

	keyName := d.Id()
	if keyName == "" {
		return nil
	}

  _, deleteErr := nclClient.DeleteKeyPair(keyName)
	if deleteErr != nil {
		return fmt.Errorf("Error Delete Key Pair error: %s", deleteErr)
	}
	return resourceKeyPairRead(d, meta)
}

func resourceKeyPairUpdate(d *schema.ResourceData, meta interface{}) error {
  //nclClient := meta.(*NclClient)
	keyName := d.Id()
	if keyName == "" {
		return nil
	}

	return resourceKeyPairRead(d, meta)
}

func resourceKeyPairRead(d *schema.ResourceData, meta interface{}) error {
	nclClient := meta.(*NclClient)

	keyName := d.Id()
	if keyName == "" {
		keyName = d.Get("name").(string)
	}

  resp, err := nclClient.KeyPairs([]string{keyName}, nil)
  if err != nil {
    fmt.Println("KeyPairs request failed")
  }

  var keyPairObj *compute.KeyPair
  for _, val := range resp.Keys {
    if val.Name == keyName {
        keyPairObj = &val
        break
    }
  }

  if keyPairObj == nil {
    d.SetId("")
    return nil
  }

	d.Set("name", keyPairObj.Name)

	return nil
}
