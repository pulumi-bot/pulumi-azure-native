// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200101

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type AdaptiveApplicationControl struct {
	pulumi.CustomResourceState

	// Location where the resource is stored
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Represents a VM/server group and set of rules to be allowed running on a machine
	Properties AppWhitelistingGroupDataResponseOutput `pulumi:"properties"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAdaptiveApplicationControl registers a new resource with the given unique name, arguments, and options.
func NewAdaptiveApplicationControl(ctx *pulumi.Context,
	name string, args *AdaptiveApplicationControlArgs, opts ...pulumi.ResourceOption) (*AdaptiveApplicationControl, error) {
	if args == nil || args.AscLocation == nil {
		return nil, errors.New("missing required argument 'AscLocation'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil {
		args = &AdaptiveApplicationControlArgs{}
	}
	var resource AdaptiveApplicationControl
	err := ctx.RegisterResource("azurerm:security/v20200101:AdaptiveApplicationControl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAdaptiveApplicationControl gets an existing AdaptiveApplicationControl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAdaptiveApplicationControl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AdaptiveApplicationControlState, opts ...pulumi.ResourceOption) (*AdaptiveApplicationControl, error) {
	var resource AdaptiveApplicationControl
	err := ctx.ReadResource("azurerm:security/v20200101:AdaptiveApplicationControl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AdaptiveApplicationControl resources.
type adaptiveApplicationControlState struct {
	// Location where the resource is stored
	Location *string `pulumi:"location"`
	// Resource name
	Name *string `pulumi:"name"`
	// Represents a VM/server group and set of rules to be allowed running on a machine
	Properties *AppWhitelistingGroupDataResponse `pulumi:"properties"`
	// Resource type
	Type *string `pulumi:"type"`
}

type AdaptiveApplicationControlState struct {
	// Location where the resource is stored
	Location pulumi.StringPtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// Represents a VM/server group and set of rules to be allowed running on a machine
	Properties AppWhitelistingGroupDataResponsePtrInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (AdaptiveApplicationControlState) ElementType() reflect.Type {
	return reflect.TypeOf((*adaptiveApplicationControlState)(nil)).Elem()
}

type adaptiveApplicationControlArgs struct {
	// The location where ASC stores the data of the subscription. can be retrieved from Get locations
	AscLocation string `pulumi:"ascLocation"`
	// The application control policy enforcement/protection mode of the VM/server group
	EnforcementMode *string `pulumi:"enforcementMode"`
	// Name of an application control VM/server group
	Name                string               `pulumi:"name"`
	PathRecommendations *PathRecommendations `pulumi:"pathRecommendations"`
	// The protection mode of the collection/file types. Exe/Msi/Script are used for Windows, Executable is used for Linux.
	ProtectionMode    *ProtectionMode    `pulumi:"protectionMode"`
	VmRecommendations *VmRecommendations `pulumi:"vmRecommendations"`
}

// The set of arguments for constructing a AdaptiveApplicationControl resource.
type AdaptiveApplicationControlArgs struct {
	// The location where ASC stores the data of the subscription. can be retrieved from Get locations
	AscLocation pulumi.StringInput
	// The application control policy enforcement/protection mode of the VM/server group
	EnforcementMode pulumi.StringPtrInput
	// Name of an application control VM/server group
	Name                pulumi.StringInput
	PathRecommendations PathRecommendationsPtrInput
	// The protection mode of the collection/file types. Exe/Msi/Script are used for Windows, Executable is used for Linux.
	ProtectionMode    ProtectionModePtrInput
	VmRecommendations VmRecommendationsPtrInput
}

func (AdaptiveApplicationControlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*adaptiveApplicationControlArgs)(nil)).Elem()
}
