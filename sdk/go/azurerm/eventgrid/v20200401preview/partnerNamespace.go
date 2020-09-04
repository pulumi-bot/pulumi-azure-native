// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200401preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// EventGrid Partner Namespace.
type PartnerNamespace struct {
	pulumi.CustomResourceState

	// Endpoint for the partner namespace.
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// Location of the resource.
	Location pulumi.StringOutput `pulumi:"location"`
	// Name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The fully qualified ARM Id of the partner registration that should be associated with this partner namespace. This takes the following format:
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerRegistrations/{partnerRegistrationName}.
	PartnerRegistrationFullyQualifiedId pulumi.StringPtrOutput `pulumi:"partnerRegistrationFullyQualifiedId"`
	// Provisioning state of the partner namespace.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Tags of the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Type of the resource
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPartnerNamespace registers a new resource with the given unique name, arguments, and options.
func NewPartnerNamespace(ctx *pulumi.Context,
	name string, args *PartnerNamespaceArgs, opts ...pulumi.ResourceOption) (*PartnerNamespace, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.PartnerNamespaceName == nil {
		return nil, errors.New("missing required argument 'PartnerNamespaceName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &PartnerNamespaceArgs{}
	}
	var resource PartnerNamespace
	err := ctx.RegisterResource("azurerm:eventgrid/v20200401preview:PartnerNamespace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPartnerNamespace gets an existing PartnerNamespace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPartnerNamespace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PartnerNamespaceState, opts ...pulumi.ResourceOption) (*PartnerNamespace, error) {
	var resource PartnerNamespace
	err := ctx.ReadResource("azurerm:eventgrid/v20200401preview:PartnerNamespace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PartnerNamespace resources.
type partnerNamespaceState struct {
	// Endpoint for the partner namespace.
	Endpoint *string `pulumi:"endpoint"`
	// Location of the resource.
	Location *string `pulumi:"location"`
	// Name of the resource
	Name *string `pulumi:"name"`
	// The fully qualified ARM Id of the partner registration that should be associated with this partner namespace. This takes the following format:
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerRegistrations/{partnerRegistrationName}.
	PartnerRegistrationFullyQualifiedId *string `pulumi:"partnerRegistrationFullyQualifiedId"`
	// Provisioning state of the partner namespace.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// Type of the resource
	Type *string `pulumi:"type"`
}

type PartnerNamespaceState struct {
	// Endpoint for the partner namespace.
	Endpoint pulumi.StringPtrInput
	// Location of the resource.
	Location pulumi.StringPtrInput
	// Name of the resource
	Name pulumi.StringPtrInput
	// The fully qualified ARM Id of the partner registration that should be associated with this partner namespace. This takes the following format:
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerRegistrations/{partnerRegistrationName}.
	PartnerRegistrationFullyQualifiedId pulumi.StringPtrInput
	// Provisioning state of the partner namespace.
	ProvisioningState pulumi.StringPtrInput
	// Tags of the resource.
	Tags pulumi.StringMapInput
	// Type of the resource
	Type pulumi.StringPtrInput
}

func (PartnerNamespaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*partnerNamespaceState)(nil)).Elem()
}

type partnerNamespaceArgs struct {
	// Location of the resource.
	Location string `pulumi:"location"`
	// Name of the partner namespace.
	PartnerNamespaceName string `pulumi:"partnerNamespaceName"`
	// The fully qualified ARM Id of the partner registration that should be associated with this partner namespace. This takes the following format:
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerRegistrations/{partnerRegistrationName}.
	PartnerRegistrationFullyQualifiedId *string `pulumi:"partnerRegistrationFullyQualifiedId"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Tags of the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a PartnerNamespace resource.
type PartnerNamespaceArgs struct {
	// Location of the resource.
	Location pulumi.StringInput
	// Name of the partner namespace.
	PartnerNamespaceName pulumi.StringInput
	// The fully qualified ARM Id of the partner registration that should be associated with this partner namespace. This takes the following format:
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerRegistrations/{partnerRegistrationName}.
	PartnerRegistrationFullyQualifiedId pulumi.StringPtrInput
	// The name of the resource group within the user's subscription.
	ResourceGroupName pulumi.StringInput
	// Tags of the resource.
	Tags pulumi.StringMapInput
}

func (PartnerNamespaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*partnerNamespaceArgs)(nil)).Elem()
}
