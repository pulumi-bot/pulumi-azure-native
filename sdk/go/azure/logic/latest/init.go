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
	case "azure-nextgen:logic/latest:IntegrationAccount":
		r, err = NewIntegrationAccount(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:logic/latest:IntegrationAccountAgreement":
		r, err = NewIntegrationAccountAgreement(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:logic/latest:IntegrationAccountAssembly":
		r, err = NewIntegrationAccountAssembly(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:logic/latest:IntegrationAccountBatchConfiguration":
		r, err = NewIntegrationAccountBatchConfiguration(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:logic/latest:IntegrationAccountCertificate":
		r, err = NewIntegrationAccountCertificate(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:logic/latest:IntegrationAccountMap":
		r, err = NewIntegrationAccountMap(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:logic/latest:IntegrationAccountPartner":
		r, err = NewIntegrationAccountPartner(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:logic/latest:IntegrationAccountSchema":
		r, err = NewIntegrationAccountSchema(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:logic/latest:IntegrationAccountSession":
		r, err = NewIntegrationAccountSession(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:logic/latest:IntegrationServiceEnvironment":
		r, err = NewIntegrationServiceEnvironment(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:logic/latest:IntegrationServiceEnvironmentManagedApi":
		r, err = NewIntegrationServiceEnvironmentManagedApi(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:logic/latest:RosettaNetProcessConfiguration":
		r, err = NewRosettaNetProcessConfiguration(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:logic/latest:Workflow":
		r, err = NewWorkflow(ctx, name, nil, pulumi.URN_(urn))
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
		"logic/latest",
		&module{version},
	)
}
