// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190601preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Response on GET of a hybrid use benefit
type HybridUseBenefit struct {
	pulumi.CustomResourceState

	// Created date
	CreatedDate pulumi.StringOutput `pulumi:"createdDate"`
	// Indicates the revision of the hybrid use benefit
	Etag pulumi.IntOutput `pulumi:"etag"`
	// Last updated date
	LastUpdatedDate pulumi.StringOutput `pulumi:"lastUpdatedDate"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning state
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Hybrid use benefit SKU
	Sku SkuResponseOutput `pulumi:"sku"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewHybridUseBenefit registers a new resource with the given unique name, arguments, and options.
func NewHybridUseBenefit(ctx *pulumi.Context,
	name string, args *HybridUseBenefitArgs, opts ...pulumi.ResourceOption) (*HybridUseBenefit, error) {
	if args == nil || args.PlanId == nil {
		return nil, errors.New("missing required argument 'PlanId'")
	}
	if args == nil || args.Scope == nil {
		return nil, errors.New("missing required argument 'Scope'")
	}
	if args == nil || args.Sku == nil {
		return nil, errors.New("missing required argument 'Sku'")
	}
	if args == nil {
		args = &HybridUseBenefitArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:softwareplan/latest:HybridUseBenefit"),
		},
		{
			Type: pulumi.String("azurerm:softwareplan/v20191201:HybridUseBenefit"),
		},
	})
	opts = append(opts, aliases)
	var resource HybridUseBenefit
	err := ctx.RegisterResource("azurerm:softwareplan/v20190601preview:HybridUseBenefit", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHybridUseBenefit gets an existing HybridUseBenefit resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHybridUseBenefit(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HybridUseBenefitState, opts ...pulumi.ResourceOption) (*HybridUseBenefit, error) {
	var resource HybridUseBenefit
	err := ctx.ReadResource("azurerm:softwareplan/v20190601preview:HybridUseBenefit", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HybridUseBenefit resources.
type hybridUseBenefitState struct {
	// Created date
	CreatedDate *string `pulumi:"createdDate"`
	// Indicates the revision of the hybrid use benefit
	Etag *int `pulumi:"etag"`
	// Last updated date
	LastUpdatedDate *string `pulumi:"lastUpdatedDate"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// Provisioning state
	ProvisioningState *string `pulumi:"provisioningState"`
	// Hybrid use benefit SKU
	Sku *SkuResponse `pulumi:"sku"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `pulumi:"type"`
}

type HybridUseBenefitState struct {
	// Created date
	CreatedDate pulumi.StringPtrInput
	// Indicates the revision of the hybrid use benefit
	Etag pulumi.IntPtrInput
	// Last updated date
	LastUpdatedDate pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// Provisioning state
	ProvisioningState pulumi.StringPtrInput
	// Hybrid use benefit SKU
	Sku SkuResponsePtrInput
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringPtrInput
}

func (HybridUseBenefitState) ElementType() reflect.Type {
	return reflect.TypeOf((*hybridUseBenefitState)(nil)).Elem()
}

type hybridUseBenefitArgs struct {
	// This is a unique identifier for a plan. Should be a guid.
	PlanId string `pulumi:"planId"`
	// The scope at which the operation is performed. This is limited to Microsoft.Compute/virtualMachines and Microsoft.Compute/hostGroups/hosts for now
	Scope string `pulumi:"scope"`
	// Hybrid use benefit SKU
	Sku Sku `pulumi:"sku"`
}

// The set of arguments for constructing a HybridUseBenefit resource.
type HybridUseBenefitArgs struct {
	// This is a unique identifier for a plan. Should be a guid.
	PlanId pulumi.StringInput
	// The scope at which the operation is performed. This is limited to Microsoft.Compute/virtualMachines and Microsoft.Compute/hostGroups/hosts for now
	Scope pulumi.StringInput
	// Hybrid use benefit SKU
	Sku SkuInput
}

func (HybridUseBenefitArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hybridUseBenefitArgs)(nil)).Elem()
}
