// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200901

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
	case "azure-native:databoxedge/v20200901:Addon":
		r, err = NewAddon(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:databoxedge/v20200901:BandwidthSchedule":
		r, err = NewBandwidthSchedule(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:databoxedge/v20200901:Container":
		r, err = NewContainer(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:databoxedge/v20200901:Device":
		r, err = NewDevice(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:databoxedge/v20200901:MonitoringConfig":
		r, err = NewMonitoringConfig(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:databoxedge/v20200901:Order":
		r, err = NewOrder(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:databoxedge/v20200901:Role":
		r, err = NewRole(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:databoxedge/v20200901:Share":
		r, err = NewShare(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:databoxedge/v20200901:StorageAccount":
		r, err = NewStorageAccount(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:databoxedge/v20200901:StorageAccountCredential":
		r, err = NewStorageAccountCredential(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:databoxedge/v20200901:Trigger":
		r, err = NewTrigger(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:databoxedge/v20200901:User":
		r, err = NewUser(ctx, name, nil, pulumi.URN_(urn))
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
		"databoxedge/v20200901",
		&module{version},
	)
}
