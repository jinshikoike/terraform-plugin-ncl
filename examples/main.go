package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/higebu/go-niftycloud/compute"
	"github.com/higebu/go-niftycloud/niftycloud"
)

func main() {
	auth, err := niftycloud.EnvAuth()
	if err != nil {
		log.Fatal(err)
	}
	client := compute.New(auth, niftycloud.JPWest)
	opts := compute.RunInstancesOptions{
		ImageId: "68", // CentOS 6.4 64bit Plain
		//		InstanceId:     "test",
		KeyName:      "tfplugin",
		InstanceType: "mini",
		//SecurityGroups: []compute.SecurityGroup{{Name: "examplegroup"}},
		AvailZone:      "west-11",
		AccountingType: "2",
	}
	resp, err := client.RunInstances(&opts)
	if err != nil {
		log.Fatal(err)
	}

	instances := resp.Instances
	j, err := json.Marshal(instances)
	os.Stdout.Write(j)
}
