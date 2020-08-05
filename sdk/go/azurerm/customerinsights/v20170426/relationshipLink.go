// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170426

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The relationship link resource format.
type RelationshipLink struct {
	pulumi.CustomResourceState

	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The definition of relationship link.
	Properties RelationshipLinkDefinitionResponseOutput `pulumi:"properties"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewRelationshipLink registers a new resource with the given unique name, arguments, and options.
func NewRelationshipLink(ctx *pulumi.Context,
	name string, args *RelationshipLinkArgs, opts ...pulumi.ResourceOption) (*RelationshipLink, error) {
	if args == nil || args.HubName == nil {
		return nil, errors.New("missing required argument 'HubName'")
	}
	if args == nil || args.InteractionType == nil {
		return nil, errors.New("missing required argument 'InteractionType'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ProfilePropertyReferences == nil {
		return nil, errors.New("missing required argument 'ProfilePropertyReferences'")
	}
	if args == nil || args.RelatedProfilePropertyReferences == nil {
		return nil, errors.New("missing required argument 'RelatedProfilePropertyReferences'")
	}
	if args == nil || args.RelationshipName == nil {
		return nil, errors.New("missing required argument 'RelationshipName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &RelationshipLinkArgs{}
	}
	var resource RelationshipLink
	err := ctx.RegisterResource("azurerm:customerinsights/v20170426:RelationshipLink", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRelationshipLink gets an existing RelationshipLink resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRelationshipLink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RelationshipLinkState, opts ...pulumi.ResourceOption) (*RelationshipLink, error) {
	var resource RelationshipLink
	err := ctx.ReadResource("azurerm:customerinsights/v20170426:RelationshipLink", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RelationshipLink resources.
type relationshipLinkState struct {
	// Resource name.
	Name *string `pulumi:"name"`
	// The definition of relationship link.
	Properties *RelationshipLinkDefinitionResponse `pulumi:"properties"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type RelationshipLinkState struct {
	// Resource name.
	Name pulumi.StringPtrInput
	// The definition of relationship link.
	Properties RelationshipLinkDefinitionResponsePtrInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (RelationshipLinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*relationshipLinkState)(nil)).Elem()
}

type relationshipLinkArgs struct {
	// Localized descriptions for the Relationship Link.
	Description map[string]string `pulumi:"description"`
	// Localized display name for the Relationship Link.
	DisplayName map[string]string `pulumi:"displayName"`
	// The name of the hub.
	HubName string `pulumi:"hubName"`
	// The InteractionType associated with the Relationship Link.
	InteractionType string `pulumi:"interactionType"`
	// The mappings between Interaction and Relationship fields.
	Mappings []RelationshipLinkFieldMapping `pulumi:"mappings"`
	// The name of the relationship link.
	Name string `pulumi:"name"`
	// The property references for the Profile of the Relationship.
	ProfilePropertyReferences []ParticipantProfilePropertyReference `pulumi:"profilePropertyReferences"`
	// The property references for the Related Profile of the Relationship.
	RelatedProfilePropertyReferences []ParticipantProfilePropertyReference `pulumi:"relatedProfilePropertyReferences"`
	// The Relationship associated with the Link.
	RelationshipName string `pulumi:"relationshipName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a RelationshipLink resource.
type RelationshipLinkArgs struct {
	// Localized descriptions for the Relationship Link.
	Description pulumi.StringMapInput
	// Localized display name for the Relationship Link.
	DisplayName pulumi.StringMapInput
	// The name of the hub.
	HubName pulumi.StringInput
	// The InteractionType associated with the Relationship Link.
	InteractionType pulumi.StringInput
	// The mappings between Interaction and Relationship fields.
	Mappings RelationshipLinkFieldMappingArrayInput
	// The name of the relationship link.
	Name pulumi.StringInput
	// The property references for the Profile of the Relationship.
	ProfilePropertyReferences ParticipantProfilePropertyReferenceArrayInput
	// The property references for the Related Profile of the Relationship.
	RelatedProfilePropertyReferences ParticipantProfilePropertyReferenceArrayInput
	// The Relationship associated with the Link.
	RelationshipName pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
}

func (RelationshipLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*relationshipLinkArgs)(nil)).Elem()
}
