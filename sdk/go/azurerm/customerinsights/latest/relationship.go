// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The relationship resource format.
//
// ## Example Usage
// ### Relationships_CreateOrUpdate
//
// ```go
// package main
//
// import (
// 	customerinsights "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/customerinsights/latest"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := customerinsights.NewRelationship(ctx, "relationship", &customerinsights.RelationshipArgs{
// 			Cardinality: pulumi.String("OneToOne"),
// 			Description: pulumi.StringMap{
// 				"en-us": pulumi.String("Relationship Description"),
// 			},
// 			DisplayName: pulumi.StringMap{
// 				"en-us": pulumi.String("Relationship DisplayName"),
// 			},
// 			Fields:             customerinsights.PropertyDefinitionArray{},
// 			HubName:            pulumi.String("sdkTestHub"),
// 			ProfileType:        pulumi.String("testProfile2326994"),
// 			RelatedProfileType: pulumi.String("testProfile2326994"),
// 			RelationshipName:   pulumi.String("SomeRelationship"),
// 			ResourceGroupName:  pulumi.String("TestHubRG"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type Relationship struct {
	pulumi.CustomResourceState

	// The Relationship Cardinality.
	Cardinality pulumi.StringPtrOutput `pulumi:"cardinality"`
	// Localized descriptions for the Relationship.
	Description pulumi.StringMapOutput `pulumi:"description"`
	// Localized display name for the Relationship.
	DisplayName pulumi.StringMapOutput `pulumi:"displayName"`
	// The expiry date time in UTC.
	ExpiryDateTimeUtc pulumi.StringPtrOutput `pulumi:"expiryDateTimeUtc"`
	// The properties of the Relationship.
	Fields PropertyDefinitionResponseArrayOutput `pulumi:"fields"`
	// Optional property to be used to map fields in profile to their strong ids in related profile.
	LookupMappings RelationshipTypeMappingResponseArrayOutput `pulumi:"lookupMappings"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Profile type.
	ProfileType pulumi.StringOutput `pulumi:"profileType"`
	// Provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Related profile being referenced.
	RelatedProfileType pulumi.StringOutput `pulumi:"relatedProfileType"`
	// The relationship guid id.
	RelationshipGuidId pulumi.StringOutput `pulumi:"relationshipGuidId"`
	// The Relationship name.
	RelationshipName pulumi.StringOutput `pulumi:"relationshipName"`
	// The hub name.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewRelationship registers a new resource with the given unique name, arguments, and options.
func NewRelationship(ctx *pulumi.Context,
	name string, args *RelationshipArgs, opts ...pulumi.ResourceOption) (*Relationship, error) {
	if args == nil || args.HubName == nil {
		return nil, errors.New("missing required argument 'HubName'")
	}
	if args == nil || args.ProfileType == nil {
		return nil, errors.New("missing required argument 'ProfileType'")
	}
	if args == nil || args.RelatedProfileType == nil {
		return nil, errors.New("missing required argument 'RelatedProfileType'")
	}
	if args == nil || args.RelationshipName == nil {
		return nil, errors.New("missing required argument 'RelationshipName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &RelationshipArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:customerinsights/v20170101:Relationship"),
		},
		{
			Type: pulumi.String("azurerm:customerinsights/v20170426:Relationship"),
		},
	})
	opts = append(opts, aliases)
	var resource Relationship
	err := ctx.RegisterResource("azurerm:customerinsights/latest:Relationship", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRelationship gets an existing Relationship resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRelationship(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RelationshipState, opts ...pulumi.ResourceOption) (*Relationship, error) {
	var resource Relationship
	err := ctx.ReadResource("azurerm:customerinsights/latest:Relationship", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Relationship resources.
type relationshipState struct {
	// The Relationship Cardinality.
	Cardinality *string `pulumi:"cardinality"`
	// Localized descriptions for the Relationship.
	Description map[string]string `pulumi:"description"`
	// Localized display name for the Relationship.
	DisplayName map[string]string `pulumi:"displayName"`
	// The expiry date time in UTC.
	ExpiryDateTimeUtc *string `pulumi:"expiryDateTimeUtc"`
	// The properties of the Relationship.
	Fields []PropertyDefinitionResponse `pulumi:"fields"`
	// Optional property to be used to map fields in profile to their strong ids in related profile.
	LookupMappings []RelationshipTypeMappingResponse `pulumi:"lookupMappings"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Profile type.
	ProfileType *string `pulumi:"profileType"`
	// Provisioning state.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Related profile being referenced.
	RelatedProfileType *string `pulumi:"relatedProfileType"`
	// The relationship guid id.
	RelationshipGuidId *string `pulumi:"relationshipGuidId"`
	// The Relationship name.
	RelationshipName *string `pulumi:"relationshipName"`
	// The hub name.
	TenantId *string `pulumi:"tenantId"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type RelationshipState struct {
	// The Relationship Cardinality.
	Cardinality pulumi.StringPtrInput
	// Localized descriptions for the Relationship.
	Description pulumi.StringMapInput
	// Localized display name for the Relationship.
	DisplayName pulumi.StringMapInput
	// The expiry date time in UTC.
	ExpiryDateTimeUtc pulumi.StringPtrInput
	// The properties of the Relationship.
	Fields PropertyDefinitionResponseArrayInput
	// Optional property to be used to map fields in profile to their strong ids in related profile.
	LookupMappings RelationshipTypeMappingResponseArrayInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Profile type.
	ProfileType pulumi.StringPtrInput
	// Provisioning state.
	ProvisioningState pulumi.StringPtrInput
	// Related profile being referenced.
	RelatedProfileType pulumi.StringPtrInput
	// The relationship guid id.
	RelationshipGuidId pulumi.StringPtrInput
	// The Relationship name.
	RelationshipName pulumi.StringPtrInput
	// The hub name.
	TenantId pulumi.StringPtrInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (RelationshipState) ElementType() reflect.Type {
	return reflect.TypeOf((*relationshipState)(nil)).Elem()
}

type relationshipArgs struct {
	// The Relationship Cardinality.
	Cardinality *string `pulumi:"cardinality"`
	// Localized descriptions for the Relationship.
	Description map[string]string `pulumi:"description"`
	// Localized display name for the Relationship.
	DisplayName map[string]string `pulumi:"displayName"`
	// The expiry date time in UTC.
	ExpiryDateTimeUtc *string `pulumi:"expiryDateTimeUtc"`
	// The properties of the Relationship.
	Fields []PropertyDefinition `pulumi:"fields"`
	// The name of the hub.
	HubName string `pulumi:"hubName"`
	// Optional property to be used to map fields in profile to their strong ids in related profile.
	LookupMappings []RelationshipTypeMapping `pulumi:"lookupMappings"`
	// Profile type.
	ProfileType string `pulumi:"profileType"`
	// Related profile being referenced.
	RelatedProfileType string `pulumi:"relatedProfileType"`
	// The name of the Relationship.
	RelationshipName string `pulumi:"relationshipName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a Relationship resource.
type RelationshipArgs struct {
	// The Relationship Cardinality.
	Cardinality pulumi.StringPtrInput
	// Localized descriptions for the Relationship.
	Description pulumi.StringMapInput
	// Localized display name for the Relationship.
	DisplayName pulumi.StringMapInput
	// The expiry date time in UTC.
	ExpiryDateTimeUtc pulumi.StringPtrInput
	// The properties of the Relationship.
	Fields PropertyDefinitionArrayInput
	// The name of the hub.
	HubName pulumi.StringInput
	// Optional property to be used to map fields in profile to their strong ids in related profile.
	LookupMappings RelationshipTypeMappingArrayInput
	// Profile type.
	ProfileType pulumi.StringInput
	// Related profile being referenced.
	RelatedProfileType pulumi.StringInput
	// The name of the Relationship.
	RelationshipName pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
}

func (RelationshipArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*relationshipArgs)(nil)).Elem()
}
