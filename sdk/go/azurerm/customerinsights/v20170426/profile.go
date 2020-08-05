// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170426

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The profile resource format.
type Profile struct {
	pulumi.CustomResourceState

	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The profile type definition.
	Properties ProfileTypeDefinitionResponseOutput `pulumi:"properties"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewProfile registers a new resource with the given unique name, arguments, and options.
func NewProfile(ctx *pulumi.Context,
	name string, args *ProfileArgs, opts ...pulumi.ResourceOption) (*Profile, error) {
	if args == nil || args.HubName == nil {
		return nil, errors.New("missing required argument 'HubName'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &ProfileArgs{}
	}
	var resource Profile
	err := ctx.RegisterResource("azurerm:customerinsights/v20170426:Profile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProfile gets an existing Profile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProfileState, opts ...pulumi.ResourceOption) (*Profile, error) {
	var resource Profile
	err := ctx.ReadResource("azurerm:customerinsights/v20170426:Profile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Profile resources.
type profileState struct {
	// Resource name.
	Name *string `pulumi:"name"`
	// The profile type definition.
	Properties *ProfileTypeDefinitionResponse `pulumi:"properties"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type ProfileState struct {
	// Resource name.
	Name pulumi.StringPtrInput
	// The profile type definition.
	Properties ProfileTypeDefinitionResponsePtrInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (ProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*profileState)(nil)).Elem()
}

type profileArgs struct {
	// The api entity set name. This becomes the odata entity set name for the entity Type being referred in this object.
	ApiEntitySetName *string `pulumi:"apiEntitySetName"`
	// The attributes for the Type.
	Attributes map[string][]string `pulumi:"attributes"`
	// Localized descriptions for the property.
	Description map[string]string `pulumi:"description"`
	// Localized display names for the property.
	DisplayName map[string]string `pulumi:"displayName"`
	// Type of entity.
	EntityType *string `pulumi:"entityType"`
	// The properties of the Profile.
	Fields []PropertyDefinition `pulumi:"fields"`
	// The name of the hub.
	HubName string `pulumi:"hubName"`
	// The instance count.
	InstancesCount *int `pulumi:"instancesCount"`
	// Large Image associated with the Property or EntityType.
	LargeImage *string `pulumi:"largeImage"`
	// Any custom localized attributes for the Type.
	LocalizedAttributes map[string]map[string]string `pulumi:"localizedAttributes"`
	// Medium Image associated with the Property or EntityType.
	MediumImage *string `pulumi:"mediumImage"`
	// The name of the profile.
	Name string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The schema org link. This helps ACI identify and suggest semantic models.
	SchemaItemTypeLink *string `pulumi:"schemaItemTypeLink"`
	// Small Image associated with the Property or EntityType.
	SmallImage *string `pulumi:"smallImage"`
	// The strong IDs.
	StrongIds []StrongId `pulumi:"strongIds"`
	// The timestamp property name. Represents the time when the interaction or profile update happened.
	TimestampFieldName *string `pulumi:"timestampFieldName"`
	// The name of the entity.
	TypeName *string `pulumi:"typeName"`
}

// The set of arguments for constructing a Profile resource.
type ProfileArgs struct {
	// The api entity set name. This becomes the odata entity set name for the entity Type being referred in this object.
	ApiEntitySetName pulumi.StringPtrInput
	// The attributes for the Type.
	Attributes pulumi.StringArrayMapInput
	// Localized descriptions for the property.
	Description pulumi.StringMapInput
	// Localized display names for the property.
	DisplayName pulumi.StringMapInput
	// Type of entity.
	EntityType pulumi.StringPtrInput
	// The properties of the Profile.
	Fields PropertyDefinitionArrayInput
	// The name of the hub.
	HubName pulumi.StringInput
	// The instance count.
	InstancesCount pulumi.IntPtrInput
	// Large Image associated with the Property or EntityType.
	LargeImage pulumi.StringPtrInput
	// Any custom localized attributes for the Type.
	LocalizedAttributes pulumi.StringMapMapInput
	// Medium Image associated with the Property or EntityType.
	MediumImage pulumi.StringPtrInput
	// The name of the profile.
	Name pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The schema org link. This helps ACI identify and suggest semantic models.
	SchemaItemTypeLink pulumi.StringPtrInput
	// Small Image associated with the Property or EntityType.
	SmallImage pulumi.StringPtrInput
	// The strong IDs.
	StrongIds StrongIdArrayInput
	// The timestamp property name. Represents the time when the interaction or profile update happened.
	TimestampFieldName pulumi.StringPtrInput
	// The name of the entity.
	TypeName pulumi.StringPtrInput
}

func (ProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*profileArgs)(nil)).Elem()
}
