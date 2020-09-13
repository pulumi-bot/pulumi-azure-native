// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200401preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A sql database resource.
//
// ## Example Usage
// ### Creates a sqldatabase.
//
// ```go
// package main
//
// import (
// 	synapse "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/synapse/v20200401preview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := synapse.NewSqlDatabase(ctx, "sqlDatabase", &synapse.SqlDatabaseArgs{
// 			Location:          pulumi.String("westus"),
// 			ResourceGroupName: pulumi.String("rg_88eed28c-5243-4f02-8251-bfa1e69e7057"),
// 			SqlDatabaseName:   pulumi.String("73164535-f22d-4daa-b014-1f23bbe361b7"),
// 			Tags: pulumi.StringMap{
// 				"tag1": pulumi.String("val1"),
// 			},
// 			WorkspaceName: pulumi.String("srv_85d31a1e-718b-4bda-9e53-a91230dd4ce5"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type SqlDatabase struct {
	pulumi.CustomResourceState

	// The collation of the database.
	Collation pulumi.StringPtrOutput `pulumi:"collation"`
	// The Guid of the database.
	DatabaseGuid pulumi.StringOutput `pulumi:"databaseGuid"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The max size of the database expressed in bytes.
	MaxSizeBytes pulumi.IntPtrOutput `pulumi:"maxSizeBytes"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// SystemData of SqlDatabase.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSqlDatabase registers a new resource with the given unique name, arguments, and options.
func NewSqlDatabase(ctx *pulumi.Context,
	name string, args *SqlDatabaseArgs, opts ...pulumi.ResourceOption) (*SqlDatabase, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.SqlDatabaseName == nil {
		return nil, errors.New("missing required argument 'SqlDatabaseName'")
	}
	if args == nil || args.WorkspaceName == nil {
		return nil, errors.New("missing required argument 'WorkspaceName'")
	}
	if args == nil {
		args = &SqlDatabaseArgs{}
	}
	var resource SqlDatabase
	err := ctx.RegisterResource("azurerm:synapse/v20200401preview:SqlDatabase", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSqlDatabase gets an existing SqlDatabase resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSqlDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlDatabaseState, opts ...pulumi.ResourceOption) (*SqlDatabase, error) {
	var resource SqlDatabase
	err := ctx.ReadResource("azurerm:synapse/v20200401preview:SqlDatabase", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SqlDatabase resources.
type sqlDatabaseState struct {
	// The collation of the database.
	Collation *string `pulumi:"collation"`
	// The Guid of the database.
	DatabaseGuid *string `pulumi:"databaseGuid"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The max size of the database expressed in bytes.
	MaxSizeBytes *int `pulumi:"maxSizeBytes"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// SystemData of SqlDatabase.
	SystemData *SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `pulumi:"type"`
}

type SqlDatabaseState struct {
	// The collation of the database.
	Collation pulumi.StringPtrInput
	// The Guid of the database.
	DatabaseGuid pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The max size of the database expressed in bytes.
	MaxSizeBytes pulumi.IntPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// SystemData of SqlDatabase.
	SystemData SystemDataResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringPtrInput
}

func (SqlDatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlDatabaseState)(nil)).Elem()
}

type sqlDatabaseArgs struct {
	// The collation of the database.
	Collation *string `pulumi:"collation"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The max size of the database expressed in bytes.
	MaxSizeBytes *int `pulumi:"maxSizeBytes"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the sql database.
	SqlDatabaseName string `pulumi:"sqlDatabaseName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a SqlDatabase resource.
type SqlDatabaseArgs struct {
	// The collation of the database.
	Collation pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringInput
	// The max size of the database expressed in bytes.
	MaxSizeBytes pulumi.IntPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the sql database.
	SqlDatabaseName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (SqlDatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlDatabaseArgs)(nil)).Elem()
}
