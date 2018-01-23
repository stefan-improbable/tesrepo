package main

import (
	"fmt"

	"k8s.io/client-go/pkg/api/v1"
)

func main() {
	fmt.Printf("vim-go: %v", v1.TLSCertKey)
}
