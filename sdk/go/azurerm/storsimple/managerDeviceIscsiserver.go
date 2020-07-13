// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storsimple

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The iSCSI server.
type ManagerDeviceIscsiserver struct {
	pulumi.CustomResourceState

	// The name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The properties.
	Properties ISCSIServerPropertiesResponseOutput `pulumi:"properties"`
	// The type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewManagerDeviceIscsiserver registers a new resource with the given unique name, arguments, and options.
func NewManagerDeviceIscsiserver(ctx *pulumi.Context,
	name string, args *ManagerDeviceIscsiserverArgs, opts ...pulumi.ResourceOption) (*ManagerDeviceIscsiserver, error) {
	if args == nil || args.DeviceName == nil {
		return nil, errors.New("missing required argument 'DeviceName'")
	}
	if args == nil || args.ManagerName == nil {
		return nil, errors.New("missing required argument 'ManagerName'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.Properties == nil {
		return nil, errors.New("missing required argument 'Properties'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &ManagerDeviceIscsiserverArgs{}
	}
	var resource ManagerDeviceIscsiserver
	err := ctx.RegisterResource("azurerm:storsimple:ManagerDeviceIscsiserver", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagerDeviceIscsiserver gets an existing ManagerDeviceIscsiserver resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagerDeviceIscsiserver(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagerDeviceIscsiserverState, opts ...pulumi.ResourceOption) (*ManagerDeviceIscsiserver, error) {
	var resource ManagerDeviceIscsiserver
	err := ctx.ReadResource("azurerm:storsimple:ManagerDeviceIscsiserver", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagerDeviceIscsiserver resources.
type managerDeviceIscsiserverState struct {
	// The name.
	Name *string `pulumi:"name"`
	// The properties.
	Properties *ISCSIServerPropertiesResponse `pulumi:"properties"`
	// The type.
	Type *string `pulumi:"type"`
}

type ManagerDeviceIscsiserverState struct {
	// The name.
	Name pulumi.StringPtrInput
	// The properties.
	Properties ISCSIServerPropertiesResponsePtrInput
	// The type.
	Type pulumi.StringPtrInput
}

func (ManagerDeviceIscsiserverState) ElementType() reflect.Type {
	return reflect.TypeOf((*managerDeviceIscsiserverState)(nil)).Elem()
}

type managerDeviceIscsiserverArgs struct {
	// The device name.
	DeviceName string `pulumi:"deviceName"`
	// The manager name
	ManagerName string `pulumi:"managerName"`
	// The iSCSI server name.
	Name string `pulumi:"name"`
	// The properties.
	Properties ISCSIServerProperties `pulumi:"properties"`
	// The resource group name
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ManagerDeviceIscsiserver resource.
type ManagerDeviceIscsiserverArgs struct {
	// The device name.
	DeviceName pulumi.StringInput
	// The manager name
	ManagerName pulumi.StringInput
	// The iSCSI server name.
	Name pulumi.StringInput
	// The properties.
	Properties ISCSIServerPropertiesInput
	// The resource group name
	ResourceGroupName pulumi.StringInput
}

func (ManagerDeviceIscsiserverArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managerDeviceIscsiserverArgs)(nil)).Elem()
}