// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200601preview

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
	case "azure-nextgen:documentdb/v20200601preview:CassandraResourceCassandraKeyspace":
		r, err = NewCassandraResourceCassandraKeyspace(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:documentdb/v20200601preview:CassandraResourceCassandraTable":
		r, err = NewCassandraResourceCassandraTable(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:documentdb/v20200601preview:DatabaseAccount":
		r, err = NewDatabaseAccount(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:documentdb/v20200601preview:GremlinResourceGremlinDatabase":
		r, err = NewGremlinResourceGremlinDatabase(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:documentdb/v20200601preview:GremlinResourceGremlinGraph":
		r, err = NewGremlinResourceGremlinGraph(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:documentdb/v20200601preview:MongoDBResourceMongoDBCollection":
		r, err = NewMongoDBResourceMongoDBCollection(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:documentdb/v20200601preview:MongoDBResourceMongoDBDatabase":
		r, err = NewMongoDBResourceMongoDBDatabase(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:documentdb/v20200601preview:NotebookWorkspace":
		r, err = NewNotebookWorkspace(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:documentdb/v20200601preview:SqlResourceSqlContainer":
		r, err = NewSqlResourceSqlContainer(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:documentdb/v20200601preview:SqlResourceSqlDatabase":
		r, err = NewSqlResourceSqlDatabase(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:documentdb/v20200601preview:SqlResourceSqlRoleAssignment":
		r, err = NewSqlResourceSqlRoleAssignment(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:documentdb/v20200601preview:SqlResourceSqlRoleDefinition":
		r, err = NewSqlResourceSqlRoleDefinition(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:documentdb/v20200601preview:SqlResourceSqlStoredProcedure":
		r, err = NewSqlResourceSqlStoredProcedure(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:documentdb/v20200601preview:SqlResourceSqlTrigger":
		r, err = NewSqlResourceSqlTrigger(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:documentdb/v20200601preview:SqlResourceSqlUserDefinedFunction":
		r, err = NewSqlResourceSqlUserDefinedFunction(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:documentdb/v20200601preview:TableResourceTable":
		r, err = NewTableResourceTable(ctx, name, nil, pulumi.URN_(urn))
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
		"documentdb/v20200601preview",
		&module{version},
	)
}
