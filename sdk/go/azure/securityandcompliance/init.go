// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package securityandcompliance

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
	case "azure-native:securityandcompliance:PrivateEndpointConnectionsAdtAPI":
		r, err = NewPrivateEndpointConnectionsAdtAPI(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:securityandcompliance:PrivateEndpointConnectionsComp":
		r, err = NewPrivateEndpointConnectionsComp(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:securityandcompliance:PrivateEndpointConnectionsForEDM":
		r, err = NewPrivateEndpointConnectionsForEDM(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:securityandcompliance:PrivateEndpointConnectionsForSCCPowershell":
		r, err = NewPrivateEndpointConnectionsForSCCPowershell(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:securityandcompliance:PrivateEndpointConnectionsSec":
		r, err = NewPrivateEndpointConnectionsSec(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:securityandcompliance:privateLinkServicesForEDMUpload":
		r, err = NewPrivateLinkServicesForEDMUpload(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:securityandcompliance:privateLinkServicesForM365ComplianceCenter":
		r, err = NewPrivateLinkServicesForM365ComplianceCenter(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:securityandcompliance:privateLinkServicesForM365SecurityCenter":
		r, err = NewPrivateLinkServicesForM365SecurityCenter(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:securityandcompliance:privateLinkServicesForO365ManagementActivityAPI":
		r, err = NewPrivateLinkServicesForO365ManagementActivityAPI(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:securityandcompliance:privateLinkServicesForSCCPowershell":
		r, err = NewPrivateLinkServicesForSCCPowershell(ctx, name, nil, pulumi.URN_(urn))
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
		"securityandcompliance",
		&module{version},
	)
}
