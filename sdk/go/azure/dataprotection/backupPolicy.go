// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dataprotection

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// BaseBackupPolicy resource
// API Version: 2021-02-01-preview.
type BackupPolicy struct {
	pulumi.CustomResourceState

	// Type of datasource for the backup management
	DatasourceTypes pulumi.StringArrayOutput `pulumi:"datasourceTypes"`
	// Resource name associated with the resource.
	Name       pulumi.StringOutput `pulumi:"name"`
	ObjectType pulumi.StringOutput `pulumi:"objectType"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewBackupPolicy registers a new resource with the given unique name, arguments, and options.
func NewBackupPolicy(ctx *pulumi.Context,
	name string, args *BackupPolicyArgs, opts ...pulumi.ResourceOption) (*BackupPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatasourceTypes == nil {
		return nil, errors.New("invalid value for required argument 'DatasourceTypes'")
	}
	if args.ObjectType == nil {
		return nil, errors.New("invalid value for required argument 'ObjectType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VaultName == nil {
		return nil, errors.New("invalid value for required argument 'VaultName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:dataprotection:BackupPolicy"),
		},
		{
			Type: pulumi.String("azure-native:dataprotection/v20210201preview:BackupPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:dataprotection/v20210201preview:BackupPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource BackupPolicy
	err := ctx.RegisterResource("azure-native:dataprotection:BackupPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackupPolicy gets an existing BackupPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackupPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackupPolicyState, opts ...pulumi.ResourceOption) (*BackupPolicy, error) {
	var resource BackupPolicy
	err := ctx.ReadResource("azure-native:dataprotection:BackupPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BackupPolicy resources.
type backupPolicyState struct {
	// Type of datasource for the backup management
	DatasourceTypes []string `pulumi:"datasourceTypes"`
	// Resource name associated with the resource.
	Name       *string `pulumi:"name"`
	ObjectType *string `pulumi:"objectType"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemDataResponse `pulumi:"systemData"`
	// Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
	Type *string `pulumi:"type"`
}

type BackupPolicyState struct {
	// Type of datasource for the backup management
	DatasourceTypes pulumi.StringArrayInput
	// Resource name associated with the resource.
	Name       pulumi.StringPtrInput
	ObjectType pulumi.StringPtrInput
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponsePtrInput
	// Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
	Type pulumi.StringPtrInput
}

func (BackupPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*backupPolicyState)(nil)).Elem()
}

type backupPolicyArgs struct {
	BackupPolicyName *string `pulumi:"backupPolicyName"`
	// Type of datasource for the backup management
	DatasourceTypes []string `pulumi:"datasourceTypes"`
	ObjectType      string   `pulumi:"objectType"`
	// The name of the resource group where the backup vault is present.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the backup vault.
	VaultName string `pulumi:"vaultName"`
}

// The set of arguments for constructing a BackupPolicy resource.
type BackupPolicyArgs struct {
	BackupPolicyName pulumi.StringPtrInput
	// Type of datasource for the backup management
	DatasourceTypes pulumi.StringArrayInput
	ObjectType      pulumi.StringInput
	// The name of the resource group where the backup vault is present.
	ResourceGroupName pulumi.StringInput
	// The name of the backup vault.
	VaultName pulumi.StringInput
}

func (BackupPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backupPolicyArgs)(nil)).Elem()
}

type BackupPolicyInput interface {
	pulumi.Input

	ToBackupPolicyOutput() BackupPolicyOutput
	ToBackupPolicyOutputWithContext(ctx context.Context) BackupPolicyOutput
}

func (*BackupPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupPolicy)(nil))
}

func (i *BackupPolicy) ToBackupPolicyOutput() BackupPolicyOutput {
	return i.ToBackupPolicyOutputWithContext(context.Background())
}

func (i *BackupPolicy) ToBackupPolicyOutputWithContext(ctx context.Context) BackupPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupPolicyOutput)
}

type BackupPolicyOutput struct {
	*pulumi.OutputState
}

func (BackupPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupPolicy)(nil))
}

func (o BackupPolicyOutput) ToBackupPolicyOutput() BackupPolicyOutput {
	return o
}

func (o BackupPolicyOutput) ToBackupPolicyOutputWithContext(ctx context.Context) BackupPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(BackupPolicyOutput{})
}
