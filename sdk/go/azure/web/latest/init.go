// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

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
	case "azure-nextgen:web/latest:AppServiceEnvironment":
		r, err = NewAppServiceEnvironment(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/latest:AppServicePlan":
		r, err = NewAppServicePlan(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/latest:AppServicePlanRouteForVnet":
		r, err = NewAppServicePlanRouteForVnet(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/latest:Certificate":
		r, err = NewCertificate(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/latest:CertificateCsr":
		r, err = NewCertificateCsr(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/latest:Connection":
		r, err = NewConnection(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/latest:ConnectionGateway":
		r, err = NewConnectionGateway(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/latest:CustomApi":
		r, err = NewCustomApi(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/latest:ManagedHostingEnvironment":
		r, err = NewManagedHostingEnvironment(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/latest:SiteInstanceDeployment":
		r, err = NewSiteInstanceDeployment(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/latest:SiteInstanceDeploymentSlot":
		r, err = NewSiteInstanceDeploymentSlot(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/latest:StaticSite":
		r, err = NewStaticSite(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/latest:WebApp":
		r, err = NewWebApp(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/latest:WebAppDeployment":
		r, err = NewWebAppDeployment(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/latest:WebAppDeploymentSlot":
		r, err = NewWebAppDeploymentSlot(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/latest:WebAppDiagnosticLogsConfiguration":
		r, err = NewWebAppDiagnosticLogsConfiguration(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/latest:WebAppDomainOwnershipIdentifier":
		r, err = NewWebAppDomainOwnershipIdentifier(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/latest:WebAppDomainOwnershipIdentifierSlot":
		r, err = NewWebAppDomainOwnershipIdentifierSlot(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/latest:WebAppFunction":
		r, err = NewWebAppFunction(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/latest:WebAppHostNameBinding":
		r, err = NewWebAppHostNameBinding(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/latest:WebAppHostNameBindingSlot":
		r, err = NewWebAppHostNameBindingSlot(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/latest:WebAppHybridConnection":
		r, err = NewWebAppHybridConnection(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/latest:WebAppHybridConnectionSlot":
		r, err = NewWebAppHybridConnectionSlot(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/latest:WebAppInstanceFunctionSlot":
		r, err = NewWebAppInstanceFunctionSlot(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/latest:WebAppPremierAddOn":
		r, err = NewWebAppPremierAddOn(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/latest:WebAppPremierAddOnSlot":
		r, err = NewWebAppPremierAddOnSlot(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/latest:WebAppPrivateEndpointConnection":
		r, err = NewWebAppPrivateEndpointConnection(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/latest:WebAppPublicCertificate":
		r, err = NewWebAppPublicCertificate(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/latest:WebAppPublicCertificateSlot":
		r, err = NewWebAppPublicCertificateSlot(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/latest:WebAppRelayServiceConnection":
		r, err = NewWebAppRelayServiceConnection(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/latest:WebAppRelayServiceConnectionSlot":
		r, err = NewWebAppRelayServiceConnectionSlot(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/latest:WebAppSiteExtension":
		r, err = NewWebAppSiteExtension(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/latest:WebAppSiteExtensionSlot":
		r, err = NewWebAppSiteExtensionSlot(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/latest:WebAppSlot":
		r, err = NewWebAppSlot(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/latest:WebAppSlotConfigurationNames":
		r, err = NewWebAppSlotConfigurationNames(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/latest:WebAppSourceControl":
		r, err = NewWebAppSourceControl(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/latest:WebAppSourceControlSlot":
		r, err = NewWebAppSourceControlSlot(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/latest:WebAppSwiftVirtualNetworkConnection":
		r, err = NewWebAppSwiftVirtualNetworkConnection(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/latest:WebAppSwiftVirtualNetworkConnectionSlot":
		r, err = NewWebAppSwiftVirtualNetworkConnectionSlot(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/latest:WebAppVnetConnection":
		r, err = NewWebAppVnetConnection(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:web/latest:WebAppVnetConnectionSlot":
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
		"web/latest",
		&module{version},
	)
}
