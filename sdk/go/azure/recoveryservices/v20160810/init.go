// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20160810

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
	case "azure-native:recoveryservices/v20160810:ReplicationFabric":
		r, err = NewReplicationFabric(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:recoveryservices/v20160810:ReplicationNetworkMapping":
		r, err = NewReplicationNetworkMapping(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:recoveryservices/v20160810:ReplicationPolicy":
		r, err = NewReplicationPolicy(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:recoveryservices/v20160810:ReplicationProtectedItem":
		r, err = NewReplicationProtectedItem(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:recoveryservices/v20160810:ReplicationProtectionContainerMapping":
		r, err = NewReplicationProtectionContainerMapping(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:recoveryservices/v20160810:ReplicationRecoveryPlan":
		r, err = NewReplicationRecoveryPlan(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:recoveryservices/v20160810:ReplicationStorageClassificationMapping":
		r, err = NewReplicationStorageClassificationMapping(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:recoveryservices/v20160810:ReplicationvCenter":
		r, err = NewReplicationvCenter(ctx, name, nil, pulumi.URN_(urn))
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
		"recoveryservices/v20160810",
		&module{version},
	)
}
