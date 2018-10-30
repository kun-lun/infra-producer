package main

import (
	"fmt"

	"github.com/xplaceholder/artifacts/pkg/apis/manifests"
	"github.com/xplaceholder/artifacts/pkg/apis/resources"
)

func main() {
	// Below code will be executed by checker, the checker determines what resource will be created
	lb := resources.LoadBalancer{
		SKU: resources.LoadBalancerStandardSKU,
	}
	vmGroups := []resources.VMGroup{
		{
			Name:  "group1",
			Count: 2,
			SKU:   resources.VMStandardD2V3,
			Type:  "VM",
		},
		{
			Name:  "group2",
			Count: 3,
			SKU:   resources.VMStandardD4V3,
			Type:  "VMSS",
		},
	}
	// The checker add needed resource to manifest
	manifest := manifests.InfraManifest{
		LoadBalancer: &lb,
		VMGroups:     vmGroups,
	}
	// The checker convert the object to yaml bytes
	b, _ := manifest.ToYAML()
	fmt.Println(string(b))
	// ...
	// The checker write yaml bytes to the disk
	// ...

	// Below code will be executed by later on modules, here we take infra module as an example

	// ...
	// The infra module read yaml file from disk and get yaml bytes
	// ...

	// The infra module new a manifest object using yaml bytes
	im, err := manifests.NewInfraManifestFromYAML(b)
	if err != nil {
		fmt.Println(err)
		return
	}
	// The infra module access field in the yaml object as needed
	fmt.Println(im.LoadBalancer)
	fmt.Println(im.VMGroups)
	fmt.Println(im.Region)
}
