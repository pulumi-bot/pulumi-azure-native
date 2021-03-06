// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package documentdb

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
	case "azure-native:documentdb:CassandraCluster":
		r, err = NewCassandraCluster(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:documentdb:CassandraDataCenter":
		r, err = NewCassandraDataCenter(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:documentdb:CassandraResourceCassandraKeyspace":
		r, err = NewCassandraResourceCassandraKeyspace(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:documentdb:CassandraResourceCassandraTable":
		r, err = NewCassandraResourceCassandraTable(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:documentdb:DatabaseAccount":
		r, err = NewDatabaseAccount(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:documentdb:DatabaseAccountCassandraKeyspace":
		r, err = NewDatabaseAccountCassandraKeyspace(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:documentdb:DatabaseAccountCassandraTable":
		r, err = NewDatabaseAccountCassandraTable(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:documentdb:DatabaseAccountGremlinDatabase":
		r, err = NewDatabaseAccountGremlinDatabase(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:documentdb:DatabaseAccountGremlinGraph":
		r, err = NewDatabaseAccountGremlinGraph(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:documentdb:DatabaseAccountMongoDBCollection":
		r, err = NewDatabaseAccountMongoDBCollection(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:documentdb:DatabaseAccountMongoDBDatabase":
		r, err = NewDatabaseAccountMongoDBDatabase(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:documentdb:DatabaseAccountSqlContainer":
		r, err = NewDatabaseAccountSqlContainer(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:documentdb:DatabaseAccountSqlDatabase":
		r, err = NewDatabaseAccountSqlDatabase(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:documentdb:DatabaseAccountTable":
		r, err = NewDatabaseAccountTable(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:documentdb:GremlinResourceGremlinDatabase":
		r, err = NewGremlinResourceGremlinDatabase(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:documentdb:GremlinResourceGremlinGraph":
		r, err = NewGremlinResourceGremlinGraph(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:documentdb:MongoDBResourceMongoDBCollection":
		r, err = NewMongoDBResourceMongoDBCollection(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:documentdb:MongoDBResourceMongoDBDatabase":
		r, err = NewMongoDBResourceMongoDBDatabase(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:documentdb:NotebookWorkspace":
		r, err = NewNotebookWorkspace(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:documentdb:PrivateEndpointConnection":
		r, err = NewPrivateEndpointConnection(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:documentdb:SqlResourceSqlContainer":
		r, err = NewSqlResourceSqlContainer(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:documentdb:SqlResourceSqlDatabase":
		r, err = NewSqlResourceSqlDatabase(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:documentdb:SqlResourceSqlRoleAssignment":
		r, err = NewSqlResourceSqlRoleAssignment(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:documentdb:SqlResourceSqlRoleDefinition":
		r, err = NewSqlResourceSqlRoleDefinition(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:documentdb:SqlResourceSqlStoredProcedure":
		r, err = NewSqlResourceSqlStoredProcedure(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:documentdb:SqlResourceSqlTrigger":
		r, err = NewSqlResourceSqlTrigger(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:documentdb:SqlResourceSqlUserDefinedFunction":
		r, err = NewSqlResourceSqlUserDefinedFunction(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:documentdb:TableResourceTable":
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
		"azure-native",
		"documentdb",
		&module{version},
	)
}
