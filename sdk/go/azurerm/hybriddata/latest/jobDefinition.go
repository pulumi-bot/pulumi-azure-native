// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Job Definition.
//
// ## Example Usage
// ### JobDefinitions_CreateOrUpdatePUT83
//
// ```go
// package main
//
// import (
// 	hybriddata "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/hybriddata/latest"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := hybriddata.NewJobDefinition(ctx, "jobDefinition", &hybriddata.JobDefinitionArgs{
// 			DataManagerName: pulumi.String("TestAzureSDKOperations"),
// 			DataServiceInput: pulumi.Map{
// 				"AzureStorageType": pulumi.String("Blob"),
// 				"BackupChoice":     pulumi.String("UseExistingLatest"),
// 				"ContainerName":    pulumi.String("containerfromtest"),
// 				"DeviceName":       pulumi.String("8600-SHG0997877L71FC"),
// 				"FileNameFilter":   pulumi.String("*"),
// 				"IsDirectoryMode":  pulumi.Bool(false),
// 				"RootDirectories": pulumi.StringArray{
// 					pulumi.String("\\"),
// 				},
// 				"VolumeNames": pulumi.StringArray{
// 					pulumi.String("TestAutomation"),
// 				},
// 			},
// 			DataServiceName:   pulumi.String("DataTransformation"),
// 			DataSinkId:        pulumi.String("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/TestAzureSDKOperations/dataStores/TestAzureStorage1"),
// 			DataSourceId:      pulumi.String("/subscriptions/6e0219f5-327a-4365-904f-05eed4227ad7/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.HybridData/dataManagers/TestAzureSDKOperations/dataStores/TestStorSimpleSource1"),
// 			JobDefinitionName: pulumi.String("jobdeffromtestcode1"),
// 			ResourceGroupName: pulumi.String("ResourceGroupForSDKTest"),
// 			RunLocation:       pulumi.String("westus"),
// 			State:             pulumi.String("Enabled"),
// 			UserConfirmation:  pulumi.String("Required"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type JobDefinition struct {
	pulumi.CustomResourceState

	// List of customer secrets containing a key identifier and key value. The key identifier is a way for the specific data source to understand the key. Value contains customer secret encrypted by the encryptionKeys.
	CustomerSecrets CustomerSecretResponseArrayOutput `pulumi:"customerSecrets"`
	// A generic json used differently by each data service type.
	DataServiceInput pulumi.MapOutput `pulumi:"dataServiceInput"`
	// Data Sink Id associated to the job definition.
	DataSinkId pulumi.StringOutput `pulumi:"dataSinkId"`
	// Data Source Id associated to the job definition.
	DataSourceId pulumi.StringOutput `pulumi:"dataSourceId"`
	// Last modified time of the job definition.
	LastModifiedTime pulumi.StringPtrOutput `pulumi:"lastModifiedTime"`
	// Name of the object.
	Name pulumi.StringOutput `pulumi:"name"`
	// This is the preferred geo location for the job to run.
	RunLocation pulumi.StringPtrOutput `pulumi:"runLocation"`
	// Schedule for running the job definition
	Schedules ScheduleResponseArrayOutput `pulumi:"schedules"`
	// State of the job definition.
	State pulumi.StringOutput `pulumi:"state"`
	// Type of the object.
	Type pulumi.StringOutput `pulumi:"type"`
	// Enum to detect if user confirmation is required. If not passed will default to NotRequired.
	UserConfirmation pulumi.StringPtrOutput `pulumi:"userConfirmation"`
}

// NewJobDefinition registers a new resource with the given unique name, arguments, and options.
func NewJobDefinition(ctx *pulumi.Context,
	name string, args *JobDefinitionArgs, opts ...pulumi.ResourceOption) (*JobDefinition, error) {
	if args == nil || args.DataManagerName == nil {
		return nil, errors.New("missing required argument 'DataManagerName'")
	}
	if args == nil || args.DataServiceName == nil {
		return nil, errors.New("missing required argument 'DataServiceName'")
	}
	if args == nil || args.DataSinkId == nil {
		return nil, errors.New("missing required argument 'DataSinkId'")
	}
	if args == nil || args.DataSourceId == nil {
		return nil, errors.New("missing required argument 'DataSourceId'")
	}
	if args == nil || args.JobDefinitionName == nil {
		return nil, errors.New("missing required argument 'JobDefinitionName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.State == nil {
		return nil, errors.New("missing required argument 'State'")
	}
	if args == nil {
		args = &JobDefinitionArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:hybriddata/v20160601:JobDefinition"),
		},
		{
			Type: pulumi.String("azurerm:hybriddata/v20190601:JobDefinition"),
		},
	})
	opts = append(opts, aliases)
	var resource JobDefinition
	err := ctx.RegisterResource("azurerm:hybriddata/latest:JobDefinition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJobDefinition gets an existing JobDefinition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJobDefinition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobDefinitionState, opts ...pulumi.ResourceOption) (*JobDefinition, error) {
	var resource JobDefinition
	err := ctx.ReadResource("azurerm:hybriddata/latest:JobDefinition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering JobDefinition resources.
type jobDefinitionState struct {
	// List of customer secrets containing a key identifier and key value. The key identifier is a way for the specific data source to understand the key. Value contains customer secret encrypted by the encryptionKeys.
	CustomerSecrets []CustomerSecretResponse `pulumi:"customerSecrets"`
	// A generic json used differently by each data service type.
	DataServiceInput map[string]interface{} `pulumi:"dataServiceInput"`
	// Data Sink Id associated to the job definition.
	DataSinkId *string `pulumi:"dataSinkId"`
	// Data Source Id associated to the job definition.
	DataSourceId *string `pulumi:"dataSourceId"`
	// Last modified time of the job definition.
	LastModifiedTime *string `pulumi:"lastModifiedTime"`
	// Name of the object.
	Name *string `pulumi:"name"`
	// This is the preferred geo location for the job to run.
	RunLocation *string `pulumi:"runLocation"`
	// Schedule for running the job definition
	Schedules []ScheduleResponse `pulumi:"schedules"`
	// State of the job definition.
	State *string `pulumi:"state"`
	// Type of the object.
	Type *string `pulumi:"type"`
	// Enum to detect if user confirmation is required. If not passed will default to NotRequired.
	UserConfirmation *string `pulumi:"userConfirmation"`
}

type JobDefinitionState struct {
	// List of customer secrets containing a key identifier and key value. The key identifier is a way for the specific data source to understand the key. Value contains customer secret encrypted by the encryptionKeys.
	CustomerSecrets CustomerSecretResponseArrayInput
	// A generic json used differently by each data service type.
	DataServiceInput pulumi.MapInput
	// Data Sink Id associated to the job definition.
	DataSinkId pulumi.StringPtrInput
	// Data Source Id associated to the job definition.
	DataSourceId pulumi.StringPtrInput
	// Last modified time of the job definition.
	LastModifiedTime pulumi.StringPtrInput
	// Name of the object.
	Name pulumi.StringPtrInput
	// This is the preferred geo location for the job to run.
	RunLocation pulumi.StringPtrInput
	// Schedule for running the job definition
	Schedules ScheduleResponseArrayInput
	// State of the job definition.
	State pulumi.StringPtrInput
	// Type of the object.
	Type pulumi.StringPtrInput
	// Enum to detect if user confirmation is required. If not passed will default to NotRequired.
	UserConfirmation pulumi.StringPtrInput
}

func (JobDefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobDefinitionState)(nil)).Elem()
}

type jobDefinitionArgs struct {
	// List of customer secrets containing a key identifier and key value. The key identifier is a way for the specific data source to understand the key. Value contains customer secret encrypted by the encryptionKeys.
	CustomerSecrets []CustomerSecret `pulumi:"customerSecrets"`
	// The name of the DataManager Resource within the specified resource group. DataManager names must be between 3 and 24 characters in length and use any alphanumeric and underscore only
	DataManagerName string `pulumi:"dataManagerName"`
	// A generic json used differently by each data service type.
	DataServiceInput map[string]interface{} `pulumi:"dataServiceInput"`
	// The data service type of the job definition.
	DataServiceName string `pulumi:"dataServiceName"`
	// Data Sink Id associated to the job definition.
	DataSinkId string `pulumi:"dataSinkId"`
	// Data Source Id associated to the job definition.
	DataSourceId string `pulumi:"dataSourceId"`
	// The job definition name to be created or updated.
	JobDefinitionName string `pulumi:"jobDefinitionName"`
	// Last modified time of the job definition.
	LastModifiedTime *string `pulumi:"lastModifiedTime"`
	// The Resource Group Name
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// This is the preferred geo location for the job to run.
	RunLocation *string `pulumi:"runLocation"`
	// Schedule for running the job definition
	Schedules []Schedule `pulumi:"schedules"`
	// State of the job definition.
	State string `pulumi:"state"`
	// Enum to detect if user confirmation is required. If not passed will default to NotRequired.
	UserConfirmation *string `pulumi:"userConfirmation"`
}

// The set of arguments for constructing a JobDefinition resource.
type JobDefinitionArgs struct {
	// List of customer secrets containing a key identifier and key value. The key identifier is a way for the specific data source to understand the key. Value contains customer secret encrypted by the encryptionKeys.
	CustomerSecrets CustomerSecretArrayInput
	// The name of the DataManager Resource within the specified resource group. DataManager names must be between 3 and 24 characters in length and use any alphanumeric and underscore only
	DataManagerName pulumi.StringInput
	// A generic json used differently by each data service type.
	DataServiceInput pulumi.MapInput
	// The data service type of the job definition.
	DataServiceName pulumi.StringInput
	// Data Sink Id associated to the job definition.
	DataSinkId pulumi.StringInput
	// Data Source Id associated to the job definition.
	DataSourceId pulumi.StringInput
	// The job definition name to be created or updated.
	JobDefinitionName pulumi.StringInput
	// Last modified time of the job definition.
	LastModifiedTime pulumi.StringPtrInput
	// The Resource Group Name
	ResourceGroupName pulumi.StringInput
	// This is the preferred geo location for the job to run.
	RunLocation pulumi.StringPtrInput
	// Schedule for running the job definition
	Schedules ScheduleArrayInput
	// State of the job definition.
	State pulumi.StringInput
	// Enum to detect if user confirmation is required. If not passed will default to NotRequired.
	UserConfirmation pulumi.StringPtrInput
}

func (JobDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobDefinitionArgs)(nil)).Elem()
}
