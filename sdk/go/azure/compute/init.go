// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

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
	case "azure-native:compute:AvailabilitySet":
		r, err = NewAvailabilitySet(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:compute:CloudService":
		r, err = NewCloudService(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:compute:DedicatedHost":
		r, err = NewDedicatedHost(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:compute:DedicatedHostGroup":
		r, err = NewDedicatedHostGroup(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:compute:Disk":
		r, err = NewDisk(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:compute:DiskAccess":
		r, err = NewDiskAccess(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:compute:DiskAccessAPrivateEndpointConnection":
		r, err = NewDiskAccessAPrivateEndpointConnection(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:compute:DiskEncryptionSet":
		r, err = NewDiskEncryptionSet(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:compute:Gallery":
		r, err = NewGallery(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:compute:GalleryApplication":
		r, err = NewGalleryApplication(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:compute:GalleryApplicationVersion":
		r, err = NewGalleryApplicationVersion(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:compute:GalleryImage":
		r, err = NewGalleryImage(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:compute:GalleryImageVersion":
		r, err = NewGalleryImageVersion(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:compute:Image":
		r, err = NewImage(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:compute:ProximityPlacementGroup":
		r, err = NewProximityPlacementGroup(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:compute:Snapshot":
		r, err = NewSnapshot(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:compute:SshPublicKey":
		r, err = NewSshPublicKey(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:compute:VirtualMachine":
		r, err = NewVirtualMachine(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:compute:VirtualMachineExtension":
		r, err = NewVirtualMachineExtension(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:compute:VirtualMachineRunCommandByVirtualMachine":
		r, err = NewVirtualMachineRunCommandByVirtualMachine(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:compute:VirtualMachineScaleSet":
		r, err = NewVirtualMachineScaleSet(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:compute:VirtualMachineScaleSetExtension":
		r, err = NewVirtualMachineScaleSetExtension(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:compute:VirtualMachineScaleSetVM":
		r, err = NewVirtualMachineScaleSetVM(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:compute:VirtualMachineScaleSetVMExtension":
		r, err = NewVirtualMachineScaleSetVMExtension(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:compute:VirtualMachineScaleSetVMRunCommand":
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
		"azure-native",
		"compute",
		&module{version},
	)
}
