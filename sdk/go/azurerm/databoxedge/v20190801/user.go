// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190801

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Represents a user who has access to one or more shares on the Data Box Edge/Gateway device.
type User struct {
	pulumi.CustomResourceState

	// The object name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The storage account credential properties.
	Properties UserPropertiesResponseOutput `pulumi:"properties"`
	// The hierarchical type of the object.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewUser registers a new resource with the given unique name, arguments, and options.
func NewUser(ctx *pulumi.Context,
	name string, args *UserArgs, opts ...pulumi.ResourceOption) (*User, error) {
	if args == nil || args.DeviceName == nil {
		return nil, errors.New("missing required argument 'DeviceName'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.UserType == nil {
		return nil, errors.New("missing required argument 'UserType'")
	}
	if args == nil {
		args = &UserArgs{}
	}
	var resource User
	err := ctx.RegisterResource("azurerm:databoxedge/v20190801:User", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azurerm:databoxedge/v20190801:User", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering User resources.
type userState struct {
	// The object name.
	Name *string `pulumi:"name"`
	// The storage account credential properties.
	Properties *UserPropertiesResponse `pulumi:"properties"`
	// The hierarchical type of the object.
	Type *string `pulumi:"type"`
}

type UserState struct {
	// The object name.
	Name pulumi.StringPtrInput
	// The storage account credential properties.
	Properties UserPropertiesResponsePtrInput
	// The hierarchical type of the object.
	Type pulumi.StringPtrInput
}

func (UserState) ElementType() reflect.Type {
	return reflect.TypeOf((*userState)(nil)).Elem()
}

type userArgs struct {
	// The device name.
	DeviceName string `pulumi:"deviceName"`
	// The password details.
	EncryptedPassword *AsymmetricEncryptedSecret `pulumi:"encryptedPassword"`
	// The user name.
	Name string `pulumi:"name"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// List of shares that the user has rights on. This field should not be specified during user creation.
	ShareAccessRights []ShareAccessRight `pulumi:"shareAccessRights"`
	// Type of the user.
	UserType string `pulumi:"userType"`
}

// The set of arguments for constructing a User resource.
type UserArgs struct {
	// The device name.
	DeviceName pulumi.StringInput
	// The password details.
	EncryptedPassword AsymmetricEncryptedSecretPtrInput
	// The user name.
	Name pulumi.StringInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// List of shares that the user has rights on. This field should not be specified during user creation.
	ShareAccessRights ShareAccessRightArrayInput
	// Type of the user.
	UserType pulumi.StringInput
}

func (UserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userArgs)(nil)).Elem()
}
