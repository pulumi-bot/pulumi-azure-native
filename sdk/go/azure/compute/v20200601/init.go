// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200601

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-azure-nextgen/sdk/go/azure"
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
	case "azure-nextgen:compute/v20200601:AvailabilitySet":
		r, err = NewAvailabilitySet(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:compute/v20200601:DedicatedHost":
		r, err = NewDedicatedHost(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:compute/v20200601:DedicatedHostGroup":
		r, err = NewDedicatedHostGroup(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:compute/v20200601:Image":
		r, err = NewImage(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:compute/v20200601:ProximityPlacementGroup":
		r, err = NewProximityPlacementGroup(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:compute/v20200601:SshPublicKey":
		r, err = NewSshPublicKey(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:compute/v20200601:VirtualMachine":
		r, err = NewVirtualMachine(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:compute/v20200601:VirtualMachineExtension":
		r, err = NewVirtualMachineExtension(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:compute/v20200601:VirtualMachineRunCommandByVirtualMachine":
		r, err = NewVirtualMachineRunCommandByVirtualMachine(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:compute/v20200601:VirtualMachineScaleSet":
		r, err = NewVirtualMachineScaleSet(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:compute/v20200601:VirtualMachineScaleSetExtension":
		r, err = NewVirtualMachineScaleSetExtension(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:compute/v20200601:VirtualMachineScaleSetVM":
		r, err = NewVirtualMachineScaleSetVM(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:compute/v20200601:VirtualMachineScaleSetVMExtension":
		r, err = NewVirtualMachineScaleSetVMExtension(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:compute/v20200601:VirtualMachineScaleSetVMRunCommand":
		r, err = NewVirtualMachineScaleSetVMRunCommand(ctx, name, nil, pulumi.URN_(urn))
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
		"azure-nextgen",
		"compute/v20200601",
		&module{version},
	)
}
