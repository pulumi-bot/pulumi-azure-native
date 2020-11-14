// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20160515

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Profile of a lab user.
type User struct {
	pulumi.CustomResourceState

	// The creation date of the user profile.
	CreatedDate pulumi.StringOutput `pulumi:"createdDate"`
	// The identity of the user.
	Identity UserIdentityResponsePtrOutput `pulumi:"identity"`
	// The location of the resource.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning status of the resource.
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// The secret store of the user.
	SecretStore UserSecretStoreResponsePtrOutput `pulumi:"secretStore"`
	// The tags of the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
	// The unique immutable identifier of a resource (Guid).
	UniqueIdentifier pulumi.StringPtrOutput `pulumi:"uniqueIdentifier"`
}

// NewUser registers a new resource with the given unique name, arguments, and options.
func NewUser(ctx *pulumi.Context,
	name string, args *UserArgs, opts ...pulumi.ResourceOption) (*User, error) {
	if args == nil || args.LabName == nil {
		return nil, errors.New("missing required argument 'LabName'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &UserArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:devtestlab/latest:User"),
		},
		{
			Type: pulumi.String("azure-nextgen:devtestlab/v20180915:User"),
		},
	})
	opts = append(opts, aliases)
	var resource User
	err := ctx.RegisterResource("azure-nextgen:devtestlab/v20160515:User", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUser gets an existing User resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserState, opts ...pulumi.ResourceOption) (*User, error) {
	var resource User
	err := ctx.ReadResource("azure-nextgen:devtestlab/v20160515:User", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering User resources.
type userState struct {
	// The creation date of the user profile.
	CreatedDate *string `pulumi:"createdDate"`
	// The identity of the user.
	Identity *UserIdentityResponse `pulumi:"identity"`
	// The location of the resource.
	Location *string `pulumi:"location"`
	// The name of the resource.
	Name *string `pulumi:"name"`
	// The provisioning status of the resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The secret store of the user.
	SecretStore *UserSecretStoreResponse `pulumi:"secretStore"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type *string `pulumi:"type"`
	// The unique immutable identifier of a resource (Guid).
	UniqueIdentifier *string `pulumi:"uniqueIdentifier"`
}

type UserState struct {
	// The creation date of the user profile.
	CreatedDate pulumi.StringPtrInput
	// The identity of the user.
	Identity UserIdentityResponsePtrInput
	// The location of the resource.
	Location pulumi.StringPtrInput
	// The name of the resource.
	Name pulumi.StringPtrInput
	// The provisioning status of the resource.
	ProvisioningState pulumi.StringPtrInput
	// The secret store of the user.
	SecretStore UserSecretStoreResponsePtrInput
	// The tags of the resource.
	Tags pulumi.StringMapInput
	// The type of the resource.
	Type pulumi.StringPtrInput
	// The unique immutable identifier of a resource (Guid).
	UniqueIdentifier pulumi.StringPtrInput
}

func (UserState) ElementType() reflect.Type {
	return reflect.TypeOf((*userState)(nil)).Elem()
}

type userArgs struct {
	// The identity of the user.
	Identity *UserIdentity `pulumi:"identity"`
	// The name of the lab.
	LabName string `pulumi:"labName"`
	// The location of the resource.
	Location *string `pulumi:"location"`
	// The name of the user profile.
	Name string `pulumi:"name"`
	// The provisioning status of the resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The secret store of the user.
	SecretStore *UserSecretStore `pulumi:"secretStore"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The unique immutable identifier of a resource (Guid).
	UniqueIdentifier *string `pulumi:"uniqueIdentifier"`
}

// The set of arguments for constructing a User resource.
type UserArgs struct {
	// The identity of the user.
	Identity UserIdentityPtrInput
	// The name of the lab.
	LabName pulumi.StringInput
	// The location of the resource.
	Location pulumi.StringPtrInput
	// The name of the user profile.
	Name pulumi.StringInput
	// The provisioning status of the resource.
	ProvisioningState pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The secret store of the user.
	SecretStore UserSecretStorePtrInput
	// The tags of the resource.
	Tags pulumi.StringMapInput
	// The unique immutable identifier of a resource (Guid).
	UniqueIdentifier pulumi.StringPtrInput
}

func (UserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userArgs)(nil)).Elem()
}

type UserInput interface {
	pulumi.Input

	ToUserOutput() UserOutput
	ToUserOutputWithContext(ctx context.Context) UserOutput
}

func (User) ElementType() reflect.Type {
	return reflect.TypeOf((*User)(nil)).Elem()
}

func (i User) ToUserOutput() UserOutput {
	return i.ToUserOutputWithContext(context.Background())
}

func (i User) ToUserOutputWithContext(ctx context.Context) UserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserOutput)
}

type UserOutput struct {
	*pulumi.OutputState
}

func (UserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserOutput)(nil)).Elem()
}

func (o UserOutput) ToUserOutput() UserOutput {
	return o
}

func (o UserOutput) ToUserOutputWithContext(ctx context.Context) UserOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(UserOutput{})
}
