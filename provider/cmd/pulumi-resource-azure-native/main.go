// Copyright 2016-2020, Pulumi Corporation.

package main

import (
	"github.com/pulumi/pulumi-azure-native/provider/pkg/provider"
	"github.com/pulumi/pulumi-azure-native/provider/pkg/version"
)

var providerName = "azure-native"

func main() {
	provider.Serve(providerName, version.Version, pulumiSchema, azureApiResources)
}
