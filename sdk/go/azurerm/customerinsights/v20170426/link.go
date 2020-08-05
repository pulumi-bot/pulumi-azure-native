// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170426

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The link resource format.
type Link struct {
	pulumi.CustomResourceState

	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The definition of Link.
	Properties LinkDefinitionResponseOutput `pulumi:"properties"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewLink registers a new resource with the given unique name, arguments, and options.
func NewLink(ctx *pulumi.Context,
	name string, args *LinkArgs, opts ...pulumi.ResourceOption) (*Link, error) {
	if args == nil || args.HubName == nil {
		return nil, errors.New("missing required argument 'HubName'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ParticipantPropertyReferences == nil {
		return nil, errors.New("missing required argument 'ParticipantPropertyReferences'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.SourceEntityType == nil {
		return nil, errors.New("missing required argument 'SourceEntityType'")
	}
	if args == nil || args.SourceEntityTypeName == nil {
		return nil, errors.New("missing required argument 'SourceEntityTypeName'")
	}
	if args == nil || args.TargetEntityType == nil {
		return nil, errors.New("missing required argument 'TargetEntityType'")
	}
	if args == nil || args.TargetEntityTypeName == nil {
		return nil, errors.New("missing required argument 'TargetEntityTypeName'")
	}
	if args == nil {
		args = &LinkArgs{}
	}
	var resource Link
	err := ctx.RegisterResource("azurerm:customerinsights/v20170426:Link", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLink gets an existing Link resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LinkState, opts ...pulumi.ResourceOption) (*Link, error) {
	var resource Link
	err := ctx.ReadResource("azurerm:customerinsights/v20170426:Link", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Link resources.
type linkState struct {
	// Resource name.
	Name *string `pulumi:"name"`
	// The definition of Link.
	Properties *LinkDefinitionResponse `pulumi:"properties"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type LinkState struct {
	// Resource name.
	Name pulumi.StringPtrInput
	// The definition of Link.
	Properties LinkDefinitionResponsePtrInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (LinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*linkState)(nil)).Elem()
}

type linkArgs struct {
	// Localized descriptions for the Link.
	Description map[string]string `pulumi:"description"`
	// Localized display name for the Link.
	DisplayName map[string]string `pulumi:"displayName"`
	// The name of the hub.
	HubName string `pulumi:"hubName"`
	// The set of properties mappings between the source and target Types.
	Mappings []TypePropertiesMapping `pulumi:"mappings"`
	// The name of the link.
	Name string `pulumi:"name"`
	// Determines whether this link is supposed to create or delete instances if Link is NOT Reference Only.
	OperationType *string `pulumi:"operationType"`
	// The properties that represent the participating profile.
	ParticipantPropertyReferences []ParticipantPropertyReference `pulumi:"participantPropertyReferences"`
	// Indicating whether the link is reference only link. This flag is ignored if the Mappings are defined. If the mappings are not defined and it is set to true, links processing will not create or update profiles.
	ReferenceOnly *bool `pulumi:"referenceOnly"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Type of source entity.
	SourceEntityType string `pulumi:"sourceEntityType"`
	// Name of the source Entity Type.
	SourceEntityTypeName string `pulumi:"sourceEntityTypeName"`
	// Type of target entity.
	TargetEntityType string `pulumi:"targetEntityType"`
	// Name of the target Entity Type.
	TargetEntityTypeName string `pulumi:"targetEntityTypeName"`
}

// The set of arguments for constructing a Link resource.
type LinkArgs struct {
	// Localized descriptions for the Link.
	Description pulumi.StringMapInput
	// Localized display name for the Link.
	DisplayName pulumi.StringMapInput
	// The name of the hub.
	HubName pulumi.StringInput
	// The set of properties mappings between the source and target Types.
	Mappings TypePropertiesMappingArrayInput
	// The name of the link.
	Name pulumi.StringInput
	// Determines whether this link is supposed to create or delete instances if Link is NOT Reference Only.
	OperationType pulumi.StringPtrInput
	// The properties that represent the participating profile.
	ParticipantPropertyReferences ParticipantPropertyReferenceArrayInput
	// Indicating whether the link is reference only link. This flag is ignored if the Mappings are defined. If the mappings are not defined and it is set to true, links processing will not create or update profiles.
	ReferenceOnly pulumi.BoolPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Type of source entity.
	SourceEntityType pulumi.StringInput
	// Name of the source Entity Type.
	SourceEntityTypeName pulumi.StringInput
	// Type of target entity.
	TargetEntityType pulumi.StringInput
	// Name of the target Entity Type.
	TargetEntityTypeName pulumi.StringInput
}

func (LinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*linkArgs)(nil)).Elem()
}
