// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200801

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-azure-native/sdk/go/azure"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "azure-native:machinelearningservices/v20200801:MachineLearningCompute":
		r, err = NewMachineLearningCompute(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:machinelearningservices/v20200801:PrivateEndpointConnection":
		r, err = NewPrivateEndpointConnection(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:machinelearningservices/v20200801:Workspace":
		r, err = NewWorkspace(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:machinelearningservices/v20200801:WorkspaceConnection":
		r, err = NewWorkspaceConnection(ctx, name, nil, pulumi.URN_(urn))
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	return
}

func init() {
	version, err := azure.PkgVersion()
	if err != nil {
		fmt.Println("failed to determine package version. defaulting to v1: %v", err)
	}
	pulumi.RegisterResourceModule(
		"azure-native",
		"machinelearningservices/v20200801",
		&module{version},
	)
}
