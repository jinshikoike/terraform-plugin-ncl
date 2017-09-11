package main

import (
	"fmt"
	"log"

	"github.com/jinshikoike/go-niftycloud/compute"
	"github.com/jinshikoike/go-niftycloud/niftycloud"
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
	//	resp2, err2 := client.DescribeInstances([]string{"OjtKoike"}, nil)
	//	fmt.Println(err2)
	//	//resp2, err2 := client.DescribeInstances([]string{"i-0bpbqidh"}, nil)
	//	// fmt.Println(len(resp2.Reservations))
	//	// fmt.Println(len(resp2.Reservations[0].Instances))
	//	fmt.Println("\n")
	//	a, _ := json.Marshal(resp2.Reservations)
	//	os.Stdout.Write(a)
	//	fmt.Println("\n")
	//	fmt.Println(err2)
	//	fmt.Println("--------------------------------------------------------------------")
	//
	//	fmt.Println("----------------------------------------------------")
	//	_, err3 := client.TerminateInstances([]string{"OjtTerra"})
	//	fmt.Println(err3)
	//

	//modifyInstanceOpts := compute.ModifyInstance{
	//	InstanceType:          "mini",
	//	DisableAPITermination: false,
	//}
	//_, err2 := client.ModifyInstance("OjtTerra2", &modifyInstanceOpts)
	//if err2 != nil {
	//	fmt.Errorf("Error ModifyInstance request %#v", err2)
	//}

	resp, err := client.ImportKeyPair("keygote", "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCztl4uCOS3M+JMcDVJyWn2HjyLTVEOWWS5Fm5573iMFVF9y/XcPiXqNdVnxkWqjaxycnmyLOXYWMKurZnRF8qvLVl+MqzUYxypjcQKGySo5MxYfayUd53TWv2p+ZpykJ6omg+HBD2CEtV+4XRGb+/Q5OC40qD8d9T1XdZu6f/jUSO3RNeqRWARKmaFcVfoKYzA8p0RjLRmdJus2ir9kH3OYfSzglqmtw5m8Cj8ikgfs9C99M2KAQUflBcMeHNbIdHhTvuclA86ESRnZNyi3hUCLCme2EaClgl3wMKUxfmqTAHZvnaRs4BhOvi3BFPQXzM8dk+frtCNa+4Ut9yZZSAuKyddGcJeOGNp7ev0752JZtiG+QLwCMZ30aibImFQYhAInhRxSGq0b6UYMgETUXHwj3uJ4pm/ts8r4EODRs2PLbMQjcy41Gnnf52DgIHppNYC8zmrVfZ9wzJtuNdlp/XgiTJJlvUx1Ng+b86WkbGvVIHXaD+hokKKy5KqYF5YmlQcM0GesErvJ9iA8OKsE2t8Yt3UYooG6BMr5zwu6YntVdI2yxzYkbym5zFEMHFiu12hscp9EKo87D9Q2fdyQgBWJj4mMrNirRwFqq+rOOiZujU777Leu+fxJLnKliCo56bBIGhYh7/LUU7Eopjm9VdYeNX3mn0oP+WzPDxuWbSuxw== koike@ubuntu")

	//resp, err := client.KeyPairs([]string{"OjtKoike"}, &compute.Filter{})
	if err != nil {
		fmt.Println("Import Key error %s", err)
	}
	fmt.Print(resp)

}
