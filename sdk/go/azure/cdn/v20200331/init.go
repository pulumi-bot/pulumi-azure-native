// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200331

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
	case "azure-nextgen:cdn/v20200331:CustomDomain":
		r, err = NewCustomDomain(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:cdn/v20200331:Endpoint":
		r, err = NewEndpoint(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:cdn/v20200331:Origin":
		r, err = NewOrigin(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:cdn/v20200331:OriginGroup":
		r, err = NewOriginGroup(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:cdn/v20200331:Policy":
		r, err = NewPolicy(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:cdn/v20200331:Profile":
		r, err = NewProfile(ctx, name, nil, pulumi.URN_(urn))
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
		"cdn/v20200331",
		&module{version},
	)
}
