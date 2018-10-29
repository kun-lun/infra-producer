# Artifacts
[![Build Status](https://xplaceholderci.gugagaga.fun/buildStatus/icon?job=xplaceholder/artifacts/draft)](https://xplaceholderci.gugagaga.fun/job/xplaceholder/job/artifacts/job/draft/)

[![GoDoc](https://godoc.org/github.com/xplaceholder/artifacts?status.svg)](https://godoc.org/github.com/xplaceholder/artifacts)

[![Go Report Card](https://goreportcard.com/badge/xplaceholder/artifacts)](https://goreportcard.com/report/xplaceholder/artifacts)

# Example usage

```go
func main() {
	// Below code will be executed by checker, the checker determines what resource will be created
	lb := LoadBalancer{
		SKU: "standard",
	}
	vmGroups := []VMGroup{
		{
			Name:  "group1",
			Count: 2,
			SKU:   "Standard_D2_v3",
			Type:  "VM",
		},
		{
			Name:  "group2",
			Count: 3,
			SKU:   "Standard_D2_v2",
		},
	}
	// The checker add needed resource to manifest
	manifest := InfraManifest{
		LoadBalancer: lb,
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
	im, err := NewInfraManifestFromYAML(b)
	if err != nil {
		fmt.Println(err)
		return
	}
	// The infra module access field in the yaml object as needed
	fmt.Println(im.LoadBalancer)
	fmt.Println(im.VMGroups)
	fmt.Println(im.Region)
}
```

