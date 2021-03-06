// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azuredata

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
	case "azure-native:azuredata:DataController":
		r, err = NewDataController(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:azuredata:PostgresInstance":
		r, err = NewPostgresInstance(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:azuredata:SqlManagedInstance":
		r, err = NewSqlManagedInstance(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:azuredata:SqlServer":
		r, err = NewSqlServer(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:azuredata:SqlServerInstance":
		r, err = NewSqlServerInstance(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:azuredata:SqlServerRegistration":
		r, err = NewSqlServerRegistration(ctx, name, nil, pulumi.URN_(urn))
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
		"azuredata",
		&module{version},
	)
}
