// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20181101

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
	case "azure-nextgen:web/v20181101:Certificate":
		r, err = NewCertificate(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/v20181101:WebApp":
		r, err = NewWebApp(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/v20181101:WebAppDeployment":
		r, err = NewWebAppDeployment(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/v20181101:WebAppDeploymentSlot":
		r, err = NewWebAppDeploymentSlot(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/v20181101:WebAppDiagnosticLogsConfiguration":
		r, err = NewWebAppDiagnosticLogsConfiguration(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/v20181101:WebAppDomainOwnershipIdentifier":
		r, err = NewWebAppDomainOwnershipIdentifier(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/v20181101:WebAppDomainOwnershipIdentifierSlot":
		r, err = NewWebAppDomainOwnershipIdentifierSlot(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/v20181101:WebAppFunction":
		r, err = NewWebAppFunction(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/v20181101:WebAppHostNameBinding":
		r, err = NewWebAppHostNameBinding(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/v20181101:WebAppHostNameBindingSlot":
		r, err = NewWebAppHostNameBindingSlot(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/v20181101:WebAppHybridConnection":
		r, err = NewWebAppHybridConnection(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/v20181101:WebAppHybridConnectionSlot":
		r, err = NewWebAppHybridConnectionSlot(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/v20181101:WebAppInstanceFunctionSlot":
		r, err = NewWebAppInstanceFunctionSlot(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/v20181101:WebAppPremierAddOn":
		r, err = NewWebAppPremierAddOn(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/v20181101:WebAppPremierAddOnSlot":
		r, err = NewWebAppPremierAddOnSlot(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/v20181101:WebAppPublicCertificate":
		r, err = NewWebAppPublicCertificate(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/v20181101:WebAppPublicCertificateSlot":
		r, err = NewWebAppPublicCertificateSlot(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/v20181101:WebAppRelayServiceConnection":
		r, err = NewWebAppRelayServiceConnection(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/v20181101:WebAppRelayServiceConnectionSlot":
		r, err = NewWebAppRelayServiceConnectionSlot(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/v20181101:WebAppSiteExtension":
		r, err = NewWebAppSiteExtension(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/v20181101:WebAppSiteExtensionSlot":
		r, err = NewWebAppSiteExtensionSlot(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/v20181101:WebAppSlot":
		r, err = NewWebAppSlot(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/v20181101:WebAppSlotConfigurationNames":
		r, err = NewWebAppSlotConfigurationNames(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/v20181101:WebAppSourceControl":
		r, err = NewWebAppSourceControl(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/v20181101:WebAppSourceControlSlot":
		r, err = NewWebAppSourceControlSlot(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/v20181101:WebAppSwiftVirtualNetworkConnection":
		r, err = NewWebAppSwiftVirtualNetworkConnection(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/v20181101:WebAppSwiftVirtualNetworkConnectionSlot":
		r, err = NewWebAppSwiftVirtualNetworkConnectionSlot(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/v20181101:WebAppVnetConnection":
		r, err = NewWebAppVnetConnection(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/v20181101:WebAppVnetConnectionSlot":
		r, err = NewWebAppVnetConnectionSlot(ctx, name, nil, pulumi.URN_(urn))
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
		"web/v20181101",
		&module{version},
	)
}
