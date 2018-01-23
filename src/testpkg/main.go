package main

import (
	"fmt"

	// In both k8s versions.
	"k8s.io/client-go/rest"
	// Only in old k8s at e121606b0d09b2e1c467183ee46217fa85a6b672.
	"k8s.io/client-go/pkg/api/v1"
	"k8s.io/client-go/pkg/util/intstr"
)

func main() {
	fmt.Printf("vim-go: %v, %v, %v", v1.TLSCertKey, intstr.String, rest.RESTClient{})
}
