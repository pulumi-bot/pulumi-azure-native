// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package web

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A domain specific resource identifier.
type AppServiceDomainOwnershipIdentifier struct {
	pulumi.CustomResourceState

	// Kind of resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Identifier resource specific properties
	Properties IdentifierResponsePropertiesOutput `pulumi:"properties"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAppServiceDomainOwnershipIdentifier registers a new resource with the given unique name, arguments, and options.
func NewAppServiceDomainOwnershipIdentifier(ctx *pulumi.Context,
	name string, args *AppServiceDomainOwnershipIdentifierArgs, opts ...pulumi.ResourceOption) (*AppServiceDomainOwnershipIdentifier, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &AppServiceDomainOwnershipIdentifierArgs{}
	}
	var resource AppServiceDomainOwnershipIdentifier
	err := ctx.RegisterResource("azurerm:web:AppServiceDomainOwnershipIdentifier", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppServiceDomainOwnershipIdentifier gets an existing AppServiceDomainOwnershipIdentifier resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppServiceDomainOwnershipIdentifier(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppServiceDomainOwnershipIdentifierState, opts ...pulumi.ResourceOption) (*AppServiceDomainOwnershipIdentifier, error) {
	var resource AppServiceDomainOwnershipIdentifier
	err := ctx.ReadResource("azurerm:web:AppServiceDomainOwnershipIdentifier", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppServiceDomainOwnershipIdentifier resources.
type appServiceDomainOwnershipIdentifierState struct {
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Name.
	Name *string `pulumi:"name"`
	// Identifier resource specific properties
	Properties *IdentifierResponseProperties `pulumi:"properties"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type AppServiceDomainOwnershipIdentifierState struct {
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Resource Name.
	Name pulumi.StringPtrInput
	// Identifier resource specific properties
	Properties IdentifierResponsePropertiesPtrInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (AppServiceDomainOwnershipIdentifierState) ElementType() reflect.Type {
	return reflect.TypeOf((*appServiceDomainOwnershipIdentifierState)(nil)).Elem()
}

type appServiceDomainOwnershipIdentifierArgs struct {
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Name of domain ownership identifier.
	Name string `pulumi:"name"`
	// Identifier resource specific properties
	Properties *IdentifierProperties `pulumi:"properties"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a AppServiceDomainOwnershipIdentifier resource.
type AppServiceDomainOwnershipIdentifierArgs struct {
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Name of domain ownership identifier.
	Name pulumi.StringInput
	// Identifier resource specific properties
	Properties IdentifierPropertiesPtrInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
}

func (AppServiceDomainOwnershipIdentifierArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appServiceDomainOwnershipIdentifierArgs)(nil)).Elem()
}
