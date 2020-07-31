// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190601

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The Get Storage Account ManagementPolicies operation response.
type ManagementPolicy struct {
	pulumi.CustomResourceState

	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Returns the Storage Account Data Policies Rules.
	Properties ManagementPolicyPropertiesResponseOutput `pulumi:"properties"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewManagementPolicy registers a new resource with the given unique name, arguments, and options.
func NewManagementPolicy(ctx *pulumi.Context,
	name string, args *ManagementPolicyArgs, opts ...pulumi.ResourceOption) (*ManagementPolicy, error) {
	if args == nil || args.AccountName == nil {
		return nil, errors.New("missing required argument 'AccountName'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &ManagementPolicyArgs{}
	}
	var resource ManagementPolicy
	err := ctx.RegisterResource("azurerm:storage/v20190601:ManagementPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagementPolicy gets an existing ManagementPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagementPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagementPolicyState, opts ...pulumi.ResourceOption) (*ManagementPolicy, error) {
	var resource ManagementPolicy
	err := ctx.ReadResource("azurerm:storage/v20190601:ManagementPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagementPolicy resources.
type managementPolicyState struct {
	// The name of the resource
	Name *string `pulumi:"name"`
	// Returns the Storage Account Data Policies Rules.
	Properties *ManagementPolicyPropertiesResponse `pulumi:"properties"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `pulumi:"type"`
}

type ManagementPolicyState struct {
	// The name of the resource
	Name pulumi.StringPtrInput
	// Returns the Storage Account Data Policies Rules.
	Properties ManagementPolicyPropertiesResponsePtrInput
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringPtrInput
}

func (ManagementPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*managementPolicyState)(nil)).Elem()
}

type managementPolicyArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName string `pulumi:"accountName"`
	// The name of the Storage Account Management Policy. It should always be 'default'
	Name string `pulumi:"name"`
	// Returns the Storage Account Data Policies Rules.
	Properties *ManagementPolicyProperties `pulumi:"properties"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ManagementPolicy resource.
type ManagementPolicyArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName pulumi.StringInput
	// The name of the Storage Account Management Policy. It should always be 'default'
	Name pulumi.StringInput
	// Returns the Storage Account Data Policies Rules.
	Properties ManagementPolicyPropertiesPtrInput
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
}

func (ManagementPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managementPolicyArgs)(nil)).Elem()
}