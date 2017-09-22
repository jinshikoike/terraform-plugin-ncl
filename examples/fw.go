package main

import (
	"fmt"
	"log"
  //"net/url"
	//"time"

	"github.com/jinshikoike/go-niftycloud/compute"
	"github.com/jinshikoike/go-niftycloud/niftycloud"
)

func main() {
	auth, err := niftycloud.EnvAuth()
	if err != nil {
		log.Fatal(err)
	}
	client := compute.New(auth, niftycloud.JPWest)
	_ = client
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

	// ------------------------------
	group := compute.SecurityGroup{Name: "terraform51", Description: "powered by terraform"}
	resp, err := client.CreateSecurityGroup(group)

	if err != nil {
		fmt.Println("Import Key error %s", err)
	} else {
		fmt.Print(resp.SecurityGroup.Id)
	}

	resp2, err2 := client.SecurityGroups([]compute.SecurityGroup{group}, nil)
	if err2 != nil {
		fmt.Println("error occur ", err2)
	} else {
		fmt.Println(resp2.Groups[0].IPPerms)
	}
  fmt.Println(resp2.Groups[0].Name)

	//ip1 := compute.IPPerm{
	//	Protocol:  "tcp",
	//	FromPort:  1111,
	//	InOut:     "OUT",
	//	ToPort:    1112,
	//	SourceIPs: []string{},
	//	SourceGroups: []compute.UserSecurityGroup{
	//		compute.UserSecurityGroup{Name: "terraform4"},
	//	},
	//}

	//resp3, err3 := client.AuthorizeSecurityGroup(group, []compute.IPPerm{ip1})
	//if err3 != nil {
	//	fmt.Println("error occur ", err3)
	//} else {
	//	fmt.Println("ip1 = ", resp3)
	//}

}
