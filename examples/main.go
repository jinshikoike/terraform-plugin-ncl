package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/higebu/go-niftycloud/compute"
	"github.com/higebu/go-niftycloud/niftycloud"
)

func main() {
	auth, err := niftycloud.EnvAuth()
	if err != nil {
		log.Fatal(err)
	}
	client := compute.New(auth, niftycloud.JPWest)

	//opts := compute.RunInstancesOptions{
	//	ImageId: "68", // CentOS 6.4 64bit Plain
	//	//		InstanceId:     "test",
	//	KeyName:      "OjtKoike",

	//	InstanceType: "mini",
	//	//SecurityGroups: []compute.SecurityGroup{{Name: "examplegroup"}},
	//	AvailZone:      "west-12",
	//	AccountingType: "2",
	//}
	//resp, err := client.RunInstances(&opts)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//instances := resp.Instances
	//j, err := json.Marshal(instances)
	//os.Stdout.Write(j)

	//targetInstances := []string{resp.Instances[0].InstanceId}
	//targetInstances := []string{}
	resp2, err2 := client.DescribeInstances([]string{"OjtKoike"}, nil)
	fmt.Println(err2)
	//resp2, err2 := client.DescribeInstances([]string{"i-0bpbqidh"}, nil)
	// fmt.Println(len(resp2.Reservations))
	// fmt.Println(len(resp2.Reservations[0].Instances))
	fmt.Println("\n")
	a, _ := json.Marshal(resp2.Reservations)
	os.Stdout.Write(a)
	fmt.Println("\n")
	fmt.Println(err2)
	fmt.Println("--------------------------------------------------------------------")

	terraformId := "aaaaaa"
	fmt.Println(len(strings.Split(terraformId, ",")))
	instanceId := strings.Split(terraformId, ",")[0]
	//uniqueInstanceId := strings.Split(terraformId, ",")[1]
	//fmt.Println("terraform id : " + terraformId)
	fmt.Println("instance id : " + instanceId)
	//fmt.Println("unique instance id : " + uniqueInstanceId)

	fmt.Println("----------------------------------------------------")
	_, err3 := client.TerminateInstances([]string{"OjtTerra"})
	fmt.Println(err3)

}
