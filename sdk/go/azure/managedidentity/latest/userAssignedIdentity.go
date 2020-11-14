// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Describes an identity resource.
type UserAssignedIdentity struct {
	pulumi.CustomResourceState

	// The id of the app associated with the identity. This is a random generated UUID by MSI.
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The id of the service principal object associated with the created identity.
	PrincipalId pulumi.StringOutput `pulumi:"principalId"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The id of the tenant which the identity belongs to.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewUserAssignedIdentity registers a new resource with the given unique name, arguments, and options.
func NewUserAssignedIdentity(ctx *pulumi.Context,
	name string, args *UserAssignedIdentityArgs, opts ...pulumi.ResourceOption) (*UserAssignedIdentity, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ResourceName == nil {
		return nil, errors.New("missing required argument 'ResourceName'")
	}
	if args == nil {
		args = &UserAssignedIdentityArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:managedidentity/v20150831preview:UserAssignedIdentity"),
		},
		{
			Type: pulumi.String("azure-nextgen:managedidentity/v20181130:UserAssignedIdentity"),
		},
	})
	opts = append(opts, aliases)
	var resource UserAssignedIdentity
	err := ctx.RegisterResource("azure-nextgen:managedidentity/latest:UserAssignedIdentity", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserAssignedIdentity gets an existing UserAssignedIdentity resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserAssignedIdentity(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserAssignedIdentityState, opts ...pulumi.ResourceOption) (*UserAssignedIdentity, error) {
	var resource UserAssignedIdentity
	err := ctx.ReadResource("azure-nextgen:managedidentity/latest:UserAssignedIdentity", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserAssignedIdentity resources.
type userAssignedIdentityState struct {
	// The id of the app associated with the identity. This is a random generated UUID by MSI.
	ClientId *string `pulumi:"clientId"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// The id of the service principal object associated with the created identity.
	PrincipalId *string `pulumi:"principalId"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The id of the tenant which the identity belongs to.
	TenantId *string `pulumi:"tenantId"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `pulumi:"type"`
}

type UserAssignedIdentityState struct {
	// The id of the app associated with the identity. This is a random generated UUID by MSI.
	ClientId pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// The id of the service principal object associated with the created identity.
	PrincipalId pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The id of the tenant which the identity belongs to.
	TenantId pulumi.StringPtrInput
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringPtrInput
}

func (UserAssignedIdentityState) ElementType() reflect.Type {
	return reflect.TypeOf((*userAssignedIdentityState)(nil)).Elem()
}

type userAssignedIdentityArgs struct {
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the Resource Group to which the identity belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the identity resource.
	ResourceName string `pulumi:"resourceName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a UserAssignedIdentity resource.
type UserAssignedIdentityArgs struct {
	// The geo-location where the resource lives
	Location pulumi.StringInput
	// The name of the Resource Group to which the identity belongs.
	ResourceGroupName pulumi.StringInput
	// The name of the identity resource.
	ResourceName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (UserAssignedIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userAssignedIdentityArgs)(nil)).Elem()
}

type UserAssignedIdentityInput interface {
	pulumi.Input

	ToUserAssignedIdentityOutput() UserAssignedIdentityOutput
	ToUserAssignedIdentityOutputWithContext(ctx context.Context) UserAssignedIdentityOutput
}

func (UserAssignedIdentity) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedIdentity)(nil)).Elem()
}

func (i UserAssignedIdentity) ToUserAssignedIdentityOutput() UserAssignedIdentityOutput {
	return i.ToUserAssignedIdentityOutputWithContext(context.Background())
}

func (i UserAssignedIdentity) ToUserAssignedIdentityOutputWithContext(ctx context.Context) UserAssignedIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserAssignedIdentityOutput)
}

type UserAssignedIdentityOutput struct {
	*pulumi.OutputState
}

func (UserAssignedIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedIdentityOutput)(nil)).Elem()
}

func (o UserAssignedIdentityOutput) ToUserAssignedIdentityOutput() UserAssignedIdentityOutput {
	return o
}

func (o UserAssignedIdentityOutput) ToUserAssignedIdentityOutputWithContext(ctx context.Context) UserAssignedIdentityOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(UserAssignedIdentityOutput{})
}
