// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191101

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A DDoS custom policy in a resource group.
type DdosCustomPolicy struct {
	pulumi.CustomResourceState

	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the DDoS custom policy.
	Properties DdosCustomPolicyPropertiesFormatResponseOutput `pulumi:"properties"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDdosCustomPolicy registers a new resource with the given unique name, arguments, and options.
func NewDdosCustomPolicy(ctx *pulumi.Context,
	name string, args *DdosCustomPolicyArgs, opts ...pulumi.ResourceOption) (*DdosCustomPolicy, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &DdosCustomPolicyArgs{}
	}
	var resource DdosCustomPolicy
	err := ctx.RegisterResource("azurerm:network/v20191101:DdosCustomPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDdosCustomPolicy gets an existing DdosCustomPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDdosCustomPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DdosCustomPolicyState, opts ...pulumi.ResourceOption) (*DdosCustomPolicy, error) {
	var resource DdosCustomPolicy
	err := ctx.ReadResource("azurerm:network/v20191101:DdosCustomPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DdosCustomPolicy resources.
type ddosCustomPolicyState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Properties of the DDoS custom policy.
	Properties *DdosCustomPolicyPropertiesFormatResponse `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type DdosCustomPolicyState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Properties of the DDoS custom policy.
	Properties DdosCustomPolicyPropertiesFormatResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (DdosCustomPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*ddosCustomPolicyState)(nil)).Elem()
}

type ddosCustomPolicyArgs struct {
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the DDoS custom policy.
	Name string `pulumi:"name"`
	// The protocol-specific DDoS policy customization parameters.
	ProtocolCustomSettings []ProtocolCustomSettingsFormat `pulumi:"protocolCustomSettings"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a DdosCustomPolicy resource.
type DdosCustomPolicyArgs struct {
	// Resource ID.
	Id pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the DDoS custom policy.
	Name pulumi.StringInput
	// The protocol-specific DDoS policy customization parameters.
	ProtocolCustomSettings ProtocolCustomSettingsFormatArrayInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (DdosCustomPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ddosCustomPolicyArgs)(nil)).Elem()
}
