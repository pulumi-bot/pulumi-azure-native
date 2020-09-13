// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170301preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A database resource.
//
// ## Example Usage
// ### Creates a database as a copy.
//
// ```go
// package main
//
// import (
// 	sql "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/sql/v20170301preview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := sql.NewDatabase(ctx, "database", &sql.DatabaseArgs{
// 			CreateMode:        pulumi.String("Copy"),
// 			DatabaseName:      pulumi.String("dbcopy"),
// 			Location:          pulumi.String("southeastasia"),
// 			ResourceGroupName: pulumi.String("Default-SQL-SouthEastAsia"),
// 			ServerName:        pulumi.String("testsvr"),
// 			Sku: &sql.SkuArgs{
// 				Name: pulumi.String("S0"),
// 				Tier: pulumi.String("Standard"),
// 			},
// 			SourceDatabaseId: pulumi.String("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/servers/testsvr/databases/testdb"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
// ### Creates a database as an on-line secondary.
//
// ```go
// package main
//
// import (
// 	sql "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/sql/v20170301preview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := sql.NewDatabase(ctx, "database", &sql.DatabaseArgs{
// 			CreateMode:        pulumi.String("Secondary"),
// 			DatabaseName:      pulumi.String("testdb"),
// 			Location:          pulumi.String("southeastasia"),
// 			ResourceGroupName: pulumi.String("Default-SQL-SouthEastAsia"),
// 			ServerName:        pulumi.String("testsvr"),
// 			Sku: &sql.SkuArgs{
// 				Name: pulumi.String("S0"),
// 				Tier: pulumi.String("Standard"),
// 			},
// 			SourceDatabaseId: pulumi.String("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-NorthEurope/providers/Microsoft.Sql/servers/testsvr1/databases/testdb"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
// ### Creates a database from PointInTimeRestore.
//
// ```go
// package main
//
// import (
// 	sql "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/sql/v20170301preview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := sql.NewDatabase(ctx, "database", &sql.DatabaseArgs{
// 			CreateMode:         pulumi.String("PointInTimeRestore"),
// 			DatabaseName:       pulumi.String("dbpitr"),
// 			Location:           pulumi.String("southeastasia"),
// 			ResourceGroupName:  pulumi.String("Default-SQL-SouthEastAsia"),
// 			RestorePointInTime: pulumi.String("2017-07-14T05:35:31.503Z"),
// 			ServerName:         pulumi.String("testsvr"),
// 			Sku: &sql.SkuArgs{
// 				Name: pulumi.String("S0"),
// 				Tier: pulumi.String("Standard"),
// 			},
// 			SourceDatabaseId: pulumi.String("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/servers/testsvr/databases/testdb"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
// ### Creates a database from recoverable DatabaseId.
//
// ```go
// package main
//
// import (
// 	sql "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/sql/v20170301preview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := sql.NewDatabase(ctx, "database", &sql.DatabaseArgs{
// 			CreateMode:                  pulumi.String("Restore"),
// 			DatabaseName:                pulumi.String("dbrestore"),
// 			Location:                    pulumi.String("southeastasia"),
// 			ResourceGroupName:           pulumi.String("Default-SQL-SouthEastAsia"),
// 			RestorableDroppedDatabaseId: pulumi.String("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/servers/testsvr/restorableDroppedDatabases/testdb2,131444841315030000"),
// 			ServerName:                  pulumi.String("testsvr"),
// 			Sku: &sql.SkuArgs{
// 				Name: pulumi.String("S0"),
// 				Tier: pulumi.String("Standard"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
// ### Creates a database from restore with database deletion time.
//
// ```go
// package main
//
// import (
// 	sql "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/sql/v20170301preview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := sql.NewDatabase(ctx, "database", &sql.DatabaseArgs{
// 			CreateMode:        pulumi.String("Restore"),
// 			DatabaseName:      pulumi.String("dbrestore"),
// 			Location:          pulumi.String("southeastasia"),
// 			ResourceGroupName: pulumi.String("Default-SQL-SouthEastAsia"),
// 			ServerName:        pulumi.String("testsvr"),
// 			Sku: &sql.SkuArgs{
// 				Name: pulumi.String("S0"),
// 				Tier: pulumi.String("Standard"),
// 			},
// 			SourceDatabaseDeletionDate: pulumi.String("2017-07-14T06:41:06.613Z"),
// 			SourceDatabaseId:           pulumi.String("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/servers/testsvr/databases/testdb"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
// ### Creates a database from restore with restorableDroppedDatabaseId.
//
// ```go
// package main
//
// import (
// 	sql "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/sql/v20170301preview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := sql.NewDatabase(ctx, "database", &sql.DatabaseArgs{
// 			CreateMode:        pulumi.String("Copy"),
// 			DatabaseName:      pulumi.String("dbcopy"),
// 			Location:          pulumi.String("southeastasia"),
// 			ResourceGroupName: pulumi.String("Default-SQL-SouthEastAsia"),
// 			ServerName:        pulumi.String("testsvr"),
// 			Sku: &sql.SkuArgs{
// 				Name: pulumi.String("S0"),
// 				Tier: pulumi.String("Standard"),
// 			},
// 			SourceDatabaseId: pulumi.String("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/servers/testsvr/databases/testdb"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
// ### Creates a database with default mode.
//
// ```go
// package main
//
// import (
// 	sql "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/sql/v20170301preview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := sql.NewDatabase(ctx, "database", &sql.DatabaseArgs{
// 			Collation:         pulumi.String("SQL_Latin1_General_CP1_CI_AS"),
// 			CreateMode:        pulumi.String("Default"),
// 			DatabaseName:      pulumi.String("testdb"),
// 			Location:          pulumi.String("southeastasia"),
// 			MaxSizeBytes:      pulumi.Int(1073741824),
// 			ResourceGroupName: pulumi.String("Default-SQL-SouthEastAsia"),
// 			ServerName:        pulumi.String("testsvr"),
// 			Sku: &sql.SkuArgs{
// 				Name: pulumi.String("S0"),
// 				Tier: pulumi.String("Standard"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
// ### Creates a database with minimum number of parameters.
//
// ```go
// package main
//
// import (
// 	sql "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/sql/v20170301preview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := sql.NewDatabase(ctx, "database", &sql.DatabaseArgs{
// 			DatabaseName:      pulumi.String("testdb"),
// 			Location:          pulumi.String("southeastasia"),
// 			ResourceGroupName: pulumi.String("Default-SQL-SouthEastAsia"),
// 			ServerName:        pulumi.String("testsvr"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type Database struct {
	pulumi.CustomResourceState

	// Collation of the metadata catalog.
	CatalogCollation pulumi.StringPtrOutput `pulumi:"catalogCollation"`
	// The collation of the database.
	Collation pulumi.StringPtrOutput `pulumi:"collation"`
	// Specifies the mode of database creation.
	//
	// Default: regular database creation.
	//
	// Copy: creates a database as a copy of an existing database. sourceDatabaseId must be specified as the resource ID of the source database.
	//
	// Secondary: creates a database as a secondary replica of an existing database. sourceDatabaseId must be specified as the resource ID of the existing primary database.
	//
	// PointInTimeRestore: Creates a database by restoring a point in time backup of an existing database. sourceDatabaseId must be specified as the resource ID of the existing database, and restorePointInTime must be specified.
	//
	// Recovery: Creates a database by restoring a geo-replicated backup. sourceDatabaseId must be specified as the recoverable database resource ID to restore.
	//
	// Restore: Creates a database by restoring a backup of a deleted database. sourceDatabaseId must be specified. If sourceDatabaseId is the database's original resource ID, then sourceDatabaseDeletionDate must be specified. Otherwise sourceDatabaseId must be the restorable dropped database resource ID and sourceDatabaseDeletionDate is ignored. restorePointInTime may also be specified to restore from an earlier point in time.
	//
	// RestoreLongTermRetentionBackup: Creates a database by restoring from a long term retention vault. recoveryServicesRecoveryPointResourceId must be specified as the recovery point resource ID.
	//
	// Copy, Secondary, and RestoreLongTermRetentionBackup are not supported for DataWarehouse edition.
	CreateMode pulumi.StringPtrOutput `pulumi:"createMode"`
	// The creation date of the database (ISO8601 format).
	CreationDate pulumi.StringOutput `pulumi:"creationDate"`
	// The current service level objective name of the database.
	CurrentServiceObjectiveName pulumi.StringOutput `pulumi:"currentServiceObjectiveName"`
	// The ID of the database.
	DatabaseId pulumi.StringOutput `pulumi:"databaseId"`
	// The default secondary region for this database.
	DefaultSecondaryLocation pulumi.StringOutput `pulumi:"defaultSecondaryLocation"`
	// The resource identifier of the elastic pool containing this database.
	ElasticPoolId pulumi.StringPtrOutput `pulumi:"elasticPoolId"`
	// Failover Group resource identifier that this database belongs to.
	FailoverGroupId pulumi.StringOutput `pulumi:"failoverGroupId"`
	// Kind of database. This is metadata used for the Azure portal experience.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// The resource identifier of the long term retention backup associated with create operation of this database.
	LongTermRetentionBackupResourceId pulumi.StringPtrOutput `pulumi:"longTermRetentionBackupResourceId"`
	// The max size of the database expressed in bytes.
	MaxSizeBytes pulumi.IntPtrOutput `pulumi:"maxSizeBytes"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource identifier of the recoverable database associated with create operation of this database.
	RecoverableDatabaseId pulumi.StringPtrOutput `pulumi:"recoverableDatabaseId"`
	// The resource identifier of the recovery point associated with create operation of this database.
	RecoveryServicesRecoveryPointId pulumi.StringPtrOutput `pulumi:"recoveryServicesRecoveryPointId"`
	// The resource identifier of the restorable dropped database associated with create operation of this database.
	RestorableDroppedDatabaseId pulumi.StringPtrOutput `pulumi:"restorableDroppedDatabaseId"`
	// Specifies the point in time (ISO8601 format) of the source database that will be restored to create the new database.
	RestorePointInTime pulumi.StringPtrOutput `pulumi:"restorePointInTime"`
	// The name of the sample schema to apply when creating this database.
	SampleName pulumi.StringPtrOutput `pulumi:"sampleName"`
	// The name and tier of the SKU.
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// Specifies the time that the database was deleted.
	SourceDatabaseDeletionDate pulumi.StringPtrOutput `pulumi:"sourceDatabaseDeletionDate"`
	// The resource identifier of the source database associated with create operation of this database.
	SourceDatabaseId pulumi.StringPtrOutput `pulumi:"sourceDatabaseId"`
	// The status of the database.
	Status pulumi.StringOutput `pulumi:"status"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones.
	ZoneRedundant pulumi.BoolPtrOutput `pulumi:"zoneRedundant"`
}

// NewDatabase registers a new resource with the given unique name, arguments, and options.
func NewDatabase(ctx *pulumi.Context,
	name string, args *DatabaseArgs, opts ...pulumi.ResourceOption) (*Database, error) {
	if args == nil || args.DatabaseName == nil {
		return nil, errors.New("missing required argument 'DatabaseName'")
	}
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServerName == nil {
		return nil, errors.New("missing required argument 'ServerName'")
	}
	if args == nil {
		args = &DatabaseArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:sql/latest:Database"),
		},
		{
			Type: pulumi.String("azurerm:sql/v20140401:Database"),
		},
		{
			Type: pulumi.String("azurerm:sql/v20171001preview:Database"),
		},
		{
			Type: pulumi.String("azurerm:sql/v20190601preview:Database"),
		},
	})
	opts = append(opts, aliases)
	var resource Database
	err := ctx.RegisterResource("azurerm:sql/v20170301preview:Database", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabase gets an existing Database resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseState, opts ...pulumi.ResourceOption) (*Database, error) {
	var resource Database
	err := ctx.ReadResource("azurerm:sql/v20170301preview:Database", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Database resources.
type databaseState struct {
	// Collation of the metadata catalog.
	CatalogCollation *string `pulumi:"catalogCollation"`
	// The collation of the database.
	Collation *string `pulumi:"collation"`
	// Specifies the mode of database creation.
	//
	// Default: regular database creation.
	//
	// Copy: creates a database as a copy of an existing database. sourceDatabaseId must be specified as the resource ID of the source database.
	//
	// Secondary: creates a database as a secondary replica of an existing database. sourceDatabaseId must be specified as the resource ID of the existing primary database.
	//
	// PointInTimeRestore: Creates a database by restoring a point in time backup of an existing database. sourceDatabaseId must be specified as the resource ID of the existing database, and restorePointInTime must be specified.
	//
	// Recovery: Creates a database by restoring a geo-replicated backup. sourceDatabaseId must be specified as the recoverable database resource ID to restore.
	//
	// Restore: Creates a database by restoring a backup of a deleted database. sourceDatabaseId must be specified. If sourceDatabaseId is the database's original resource ID, then sourceDatabaseDeletionDate must be specified. Otherwise sourceDatabaseId must be the restorable dropped database resource ID and sourceDatabaseDeletionDate is ignored. restorePointInTime may also be specified to restore from an earlier point in time.
	//
	// RestoreLongTermRetentionBackup: Creates a database by restoring from a long term retention vault. recoveryServicesRecoveryPointResourceId must be specified as the recovery point resource ID.
	//
	// Copy, Secondary, and RestoreLongTermRetentionBackup are not supported for DataWarehouse edition.
	CreateMode *string `pulumi:"createMode"`
	// The creation date of the database (ISO8601 format).
	CreationDate *string `pulumi:"creationDate"`
	// The current service level objective name of the database.
	CurrentServiceObjectiveName *string `pulumi:"currentServiceObjectiveName"`
	// The ID of the database.
	DatabaseId *string `pulumi:"databaseId"`
	// The default secondary region for this database.
	DefaultSecondaryLocation *string `pulumi:"defaultSecondaryLocation"`
	// The resource identifier of the elastic pool containing this database.
	ElasticPoolId *string `pulumi:"elasticPoolId"`
	// Failover Group resource identifier that this database belongs to.
	FailoverGroupId *string `pulumi:"failoverGroupId"`
	// Kind of database. This is metadata used for the Azure portal experience.
	Kind *string `pulumi:"kind"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The resource identifier of the long term retention backup associated with create operation of this database.
	LongTermRetentionBackupResourceId *string `pulumi:"longTermRetentionBackupResourceId"`
	// The max size of the database expressed in bytes.
	MaxSizeBytes *int `pulumi:"maxSizeBytes"`
	// Resource name.
	Name *string `pulumi:"name"`
	// The resource identifier of the recoverable database associated with create operation of this database.
	RecoverableDatabaseId *string `pulumi:"recoverableDatabaseId"`
	// The resource identifier of the recovery point associated with create operation of this database.
	RecoveryServicesRecoveryPointId *string `pulumi:"recoveryServicesRecoveryPointId"`
	// The resource identifier of the restorable dropped database associated with create operation of this database.
	RestorableDroppedDatabaseId *string `pulumi:"restorableDroppedDatabaseId"`
	// Specifies the point in time (ISO8601 format) of the source database that will be restored to create the new database.
	RestorePointInTime *string `pulumi:"restorePointInTime"`
	// The name of the sample schema to apply when creating this database.
	SampleName *string `pulumi:"sampleName"`
	// The name and tier of the SKU.
	Sku *SkuResponse `pulumi:"sku"`
	// Specifies the time that the database was deleted.
	SourceDatabaseDeletionDate *string `pulumi:"sourceDatabaseDeletionDate"`
	// The resource identifier of the source database associated with create operation of this database.
	SourceDatabaseId *string `pulumi:"sourceDatabaseId"`
	// The status of the database.
	Status *string `pulumi:"status"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
	// Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones.
	ZoneRedundant *bool `pulumi:"zoneRedundant"`
}

type DatabaseState struct {
	// Collation of the metadata catalog.
	CatalogCollation pulumi.StringPtrInput
	// The collation of the database.
	Collation pulumi.StringPtrInput
	// Specifies the mode of database creation.
	//
	// Default: regular database creation.
	//
	// Copy: creates a database as a copy of an existing database. sourceDatabaseId must be specified as the resource ID of the source database.
	//
	// Secondary: creates a database as a secondary replica of an existing database. sourceDatabaseId must be specified as the resource ID of the existing primary database.
	//
	// PointInTimeRestore: Creates a database by restoring a point in time backup of an existing database. sourceDatabaseId must be specified as the resource ID of the existing database, and restorePointInTime must be specified.
	//
	// Recovery: Creates a database by restoring a geo-replicated backup. sourceDatabaseId must be specified as the recoverable database resource ID to restore.
	//
	// Restore: Creates a database by restoring a backup of a deleted database. sourceDatabaseId must be specified. If sourceDatabaseId is the database's original resource ID, then sourceDatabaseDeletionDate must be specified. Otherwise sourceDatabaseId must be the restorable dropped database resource ID and sourceDatabaseDeletionDate is ignored. restorePointInTime may also be specified to restore from an earlier point in time.
	//
	// RestoreLongTermRetentionBackup: Creates a database by restoring from a long term retention vault. recoveryServicesRecoveryPointResourceId must be specified as the recovery point resource ID.
	//
	// Copy, Secondary, and RestoreLongTermRetentionBackup are not supported for DataWarehouse edition.
	CreateMode pulumi.StringPtrInput
	// The creation date of the database (ISO8601 format).
	CreationDate pulumi.StringPtrInput
	// The current service level objective name of the database.
	CurrentServiceObjectiveName pulumi.StringPtrInput
	// The ID of the database.
	DatabaseId pulumi.StringPtrInput
	// The default secondary region for this database.
	DefaultSecondaryLocation pulumi.StringPtrInput
	// The resource identifier of the elastic pool containing this database.
	ElasticPoolId pulumi.StringPtrInput
	// Failover Group resource identifier that this database belongs to.
	FailoverGroupId pulumi.StringPtrInput
	// Kind of database. This is metadata used for the Azure portal experience.
	Kind pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The resource identifier of the long term retention backup associated with create operation of this database.
	LongTermRetentionBackupResourceId pulumi.StringPtrInput
	// The max size of the database expressed in bytes.
	MaxSizeBytes pulumi.IntPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// The resource identifier of the recoverable database associated with create operation of this database.
	RecoverableDatabaseId pulumi.StringPtrInput
	// The resource identifier of the recovery point associated with create operation of this database.
	RecoveryServicesRecoveryPointId pulumi.StringPtrInput
	// The resource identifier of the restorable dropped database associated with create operation of this database.
	RestorableDroppedDatabaseId pulumi.StringPtrInput
	// Specifies the point in time (ISO8601 format) of the source database that will be restored to create the new database.
	RestorePointInTime pulumi.StringPtrInput
	// The name of the sample schema to apply when creating this database.
	SampleName pulumi.StringPtrInput
	// The name and tier of the SKU.
	Sku SkuResponsePtrInput
	// Specifies the time that the database was deleted.
	SourceDatabaseDeletionDate pulumi.StringPtrInput
	// The resource identifier of the source database associated with create operation of this database.
	SourceDatabaseId pulumi.StringPtrInput
	// The status of the database.
	Status pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
	// Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones.
	ZoneRedundant pulumi.BoolPtrInput
}

func (DatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseState)(nil)).Elem()
}

type databaseArgs struct {
	// Collation of the metadata catalog.
	CatalogCollation *string `pulumi:"catalogCollation"`
	// The collation of the database.
	Collation *string `pulumi:"collation"`
	// Specifies the mode of database creation.
	//
	// Default: regular database creation.
	//
	// Copy: creates a database as a copy of an existing database. sourceDatabaseId must be specified as the resource ID of the source database.
	//
	// Secondary: creates a database as a secondary replica of an existing database. sourceDatabaseId must be specified as the resource ID of the existing primary database.
	//
	// PointInTimeRestore: Creates a database by restoring a point in time backup of an existing database. sourceDatabaseId must be specified as the resource ID of the existing database, and restorePointInTime must be specified.
	//
	// Recovery: Creates a database by restoring a geo-replicated backup. sourceDatabaseId must be specified as the recoverable database resource ID to restore.
	//
	// Restore: Creates a database by restoring a backup of a deleted database. sourceDatabaseId must be specified. If sourceDatabaseId is the database's original resource ID, then sourceDatabaseDeletionDate must be specified. Otherwise sourceDatabaseId must be the restorable dropped database resource ID and sourceDatabaseDeletionDate is ignored. restorePointInTime may also be specified to restore from an earlier point in time.
	//
	// RestoreLongTermRetentionBackup: Creates a database by restoring from a long term retention vault. recoveryServicesRecoveryPointResourceId must be specified as the recovery point resource ID.
	//
	// Copy, Secondary, and RestoreLongTermRetentionBackup are not supported for DataWarehouse edition.
	CreateMode *string `pulumi:"createMode"`
	// The name of the database.
	DatabaseName string `pulumi:"databaseName"`
	// The resource identifier of the elastic pool containing this database.
	ElasticPoolId *string `pulumi:"elasticPoolId"`
	// Resource location.
	Location string `pulumi:"location"`
	// The resource identifier of the long term retention backup associated with create operation of this database.
	LongTermRetentionBackupResourceId *string `pulumi:"longTermRetentionBackupResourceId"`
	// The max size of the database expressed in bytes.
	MaxSizeBytes *int `pulumi:"maxSizeBytes"`
	// The resource identifier of the recoverable database associated with create operation of this database.
	RecoverableDatabaseId *string `pulumi:"recoverableDatabaseId"`
	// The resource identifier of the recovery point associated with create operation of this database.
	RecoveryServicesRecoveryPointId *string `pulumi:"recoveryServicesRecoveryPointId"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The resource identifier of the restorable dropped database associated with create operation of this database.
	RestorableDroppedDatabaseId *string `pulumi:"restorableDroppedDatabaseId"`
	// Specifies the point in time (ISO8601 format) of the source database that will be restored to create the new database.
	RestorePointInTime *string `pulumi:"restorePointInTime"`
	// The name of the sample schema to apply when creating this database.
	SampleName *string `pulumi:"sampleName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
	// The name and tier of the SKU.
	Sku *Sku `pulumi:"sku"`
	// Specifies the time that the database was deleted.
	SourceDatabaseDeletionDate *string `pulumi:"sourceDatabaseDeletionDate"`
	// The resource identifier of the source database associated with create operation of this database.
	SourceDatabaseId *string `pulumi:"sourceDatabaseId"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones.
	ZoneRedundant *bool `pulumi:"zoneRedundant"`
}

// The set of arguments for constructing a Database resource.
type DatabaseArgs struct {
	// Collation of the metadata catalog.
	CatalogCollation pulumi.StringPtrInput
	// The collation of the database.
	Collation pulumi.StringPtrInput
	// Specifies the mode of database creation.
	//
	// Default: regular database creation.
	//
	// Copy: creates a database as a copy of an existing database. sourceDatabaseId must be specified as the resource ID of the source database.
	//
	// Secondary: creates a database as a secondary replica of an existing database. sourceDatabaseId must be specified as the resource ID of the existing primary database.
	//
	// PointInTimeRestore: Creates a database by restoring a point in time backup of an existing database. sourceDatabaseId must be specified as the resource ID of the existing database, and restorePointInTime must be specified.
	//
	// Recovery: Creates a database by restoring a geo-replicated backup. sourceDatabaseId must be specified as the recoverable database resource ID to restore.
	//
	// Restore: Creates a database by restoring a backup of a deleted database. sourceDatabaseId must be specified. If sourceDatabaseId is the database's original resource ID, then sourceDatabaseDeletionDate must be specified. Otherwise sourceDatabaseId must be the restorable dropped database resource ID and sourceDatabaseDeletionDate is ignored. restorePointInTime may also be specified to restore from an earlier point in time.
	//
	// RestoreLongTermRetentionBackup: Creates a database by restoring from a long term retention vault. recoveryServicesRecoveryPointResourceId must be specified as the recovery point resource ID.
	//
	// Copy, Secondary, and RestoreLongTermRetentionBackup are not supported for DataWarehouse edition.
	CreateMode pulumi.StringPtrInput
	// The name of the database.
	DatabaseName pulumi.StringInput
	// The resource identifier of the elastic pool containing this database.
	ElasticPoolId pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringInput
	// The resource identifier of the long term retention backup associated with create operation of this database.
	LongTermRetentionBackupResourceId pulumi.StringPtrInput
	// The max size of the database expressed in bytes.
	MaxSizeBytes pulumi.IntPtrInput
	// The resource identifier of the recoverable database associated with create operation of this database.
	RecoverableDatabaseId pulumi.StringPtrInput
	// The resource identifier of the recovery point associated with create operation of this database.
	RecoveryServicesRecoveryPointId pulumi.StringPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The resource identifier of the restorable dropped database associated with create operation of this database.
	RestorableDroppedDatabaseId pulumi.StringPtrInput
	// Specifies the point in time (ISO8601 format) of the source database that will be restored to create the new database.
	RestorePointInTime pulumi.StringPtrInput
	// The name of the sample schema to apply when creating this database.
	SampleName pulumi.StringPtrInput
	// The name of the server.
	ServerName pulumi.StringInput
	// The name and tier of the SKU.
	Sku SkuPtrInput
	// Specifies the time that the database was deleted.
	SourceDatabaseDeletionDate pulumi.StringPtrInput
	// The resource identifier of the source database associated with create operation of this database.
	SourceDatabaseId pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones.
	ZoneRedundant pulumi.BoolPtrInput
}

func (DatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseArgs)(nil)).Elem()
}
