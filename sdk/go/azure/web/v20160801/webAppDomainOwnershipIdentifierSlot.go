// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20160801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A domain specific resource identifier.
type WebAppDomainOwnershipIdentifierSlot struct {
	pulumi.CustomResourceState

	// Kind of resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWebAppDomainOwnershipIdentifierSlot registers a new resource with the given unique name, arguments, and options.
func NewWebAppDomainOwnershipIdentifierSlot(ctx *pulumi.Context,
	name string, args *WebAppDomainOwnershipIdentifierSlotArgs, opts ...pulumi.ResourceOption) (*WebAppDomainOwnershipIdentifierSlot, error) {
	if args == nil || args.DomainOwnershipIdentifierName == nil {
		return nil, errors.New("missing required argument 'DomainOwnershipIdentifierName'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Slot == nil {
		return nil, errors.New("missing required argument 'Slot'")
	}
	if args == nil {
		args = &WebAppDomainOwnershipIdentifierSlotArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:web/latest:WebAppDomainOwnershipIdentifierSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20180201:WebAppDomainOwnershipIdentifierSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20181101:WebAppDomainOwnershipIdentifierSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20190801:WebAppDomainOwnershipIdentifierSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20200601:WebAppDomainOwnershipIdentifierSlot"),
		},
	})
	opts = append(opts, aliases)
	var resource WebAppDomainOwnershipIdentifierSlot
	err := ctx.RegisterResource("azure-nextgen:web/v20160801:WebAppDomainOwnershipIdentifierSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebAppDomainOwnershipIdentifierSlot gets an existing WebAppDomainOwnershipIdentifierSlot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebAppDomainOwnershipIdentifierSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppDomainOwnershipIdentifierSlotState, opts ...pulumi.ResourceOption) (*WebAppDomainOwnershipIdentifierSlot, error) {
	var resource WebAppDomainOwnershipIdentifierSlot
	err := ctx.ReadResource("azure-nextgen:web/v20160801:WebAppDomainOwnershipIdentifierSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebAppDomainOwnershipIdentifierSlot resources.
type webAppDomainOwnershipIdentifierSlotState struct {
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Name.
	Name *string `pulumi:"name"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type WebAppDomainOwnershipIdentifierSlotState struct {
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Resource Name.
	Name pulumi.StringPtrInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (WebAppDomainOwnershipIdentifierSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppDomainOwnershipIdentifierSlotState)(nil)).Elem()
}

type webAppDomainOwnershipIdentifierSlotArgs struct {
	// Name of domain ownership identifier.
	DomainOwnershipIdentifierName string `pulumi:"domainOwnershipIdentifierName"`
	// String representation of the identity.
	Id *string `pulumi:"id"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Name of the app.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the deployment slot. If a slot is not specified, the API will delete the binding for the production slot.
	Slot string `pulumi:"slot"`
}

// The set of arguments for constructing a WebAppDomainOwnershipIdentifierSlot resource.
type WebAppDomainOwnershipIdentifierSlotArgs struct {
	// Name of domain ownership identifier.
	DomainOwnershipIdentifierName pulumi.StringInput
	// String representation of the identity.
	Id pulumi.StringPtrInput
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Name of the app.
	Name pulumi.StringInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
	// Name of the deployment slot. If a slot is not specified, the API will delete the binding for the production slot.
	Slot pulumi.StringInput
}

func (WebAppDomainOwnershipIdentifierSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppDomainOwnershipIdentifierSlotArgs)(nil)).Elem()
}

type WebAppDomainOwnershipIdentifierSlotInput interface {
	pulumi.Input

	ToWebAppDomainOwnershipIdentifierSlotOutput() WebAppDomainOwnershipIdentifierSlotOutput
	ToWebAppDomainOwnershipIdentifierSlotOutputWithContext(ctx context.Context) WebAppDomainOwnershipIdentifierSlotOutput
}

func (WebAppDomainOwnershipIdentifierSlot) ElementType() reflect.Type {
	return reflect.TypeOf((*WebAppDomainOwnershipIdentifierSlot)(nil)).Elem()
}

func (i WebAppDomainOwnershipIdentifierSlot) ToWebAppDomainOwnershipIdentifierSlotOutput() WebAppDomainOwnershipIdentifierSlotOutput {
	return i.ToWebAppDomainOwnershipIdentifierSlotOutputWithContext(context.Background())
}

func (i WebAppDomainOwnershipIdentifierSlot) ToWebAppDomainOwnershipIdentifierSlotOutputWithContext(ctx context.Context) WebAppDomainOwnershipIdentifierSlotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppDomainOwnershipIdentifierSlotOutput)
}

type WebAppDomainOwnershipIdentifierSlotOutput struct {
	*pulumi.OutputState
}

func (WebAppDomainOwnershipIdentifierSlotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebAppDomainOwnershipIdentifierSlotOutput)(nil)).Elem()
}

func (o WebAppDomainOwnershipIdentifierSlotOutput) ToWebAppDomainOwnershipIdentifierSlotOutput() WebAppDomainOwnershipIdentifierSlotOutput {
	return o
}

func (o WebAppDomainOwnershipIdentifierSlotOutput) ToWebAppDomainOwnershipIdentifierSlotOutputWithContext(ctx context.Context) WebAppDomainOwnershipIdentifierSlotOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WebAppDomainOwnershipIdentifierSlotOutput{})
}
