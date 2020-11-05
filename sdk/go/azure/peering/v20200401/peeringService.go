// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200401

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Peering Service
type PeeringService struct {
	pulumi.CustomResourceState

	// The location of the resource.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The PeeringServiceLocation of the Customer.
	PeeringServiceLocation pulumi.StringPtrOutput `pulumi:"peeringServiceLocation"`
	// The MAPS Provider Name.
	PeeringServiceProvider pulumi.StringPtrOutput `pulumi:"peeringServiceProvider"`
	// The provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The SKU that defines the type of the peering service.
	Sku PeeringServiceSkuResponsePtrOutput `pulumi:"sku"`
	// The resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPeeringService registers a new resource with the given unique name, arguments, and options.
func NewPeeringService(ctx *pulumi.Context,
	name string, args *PeeringServiceArgs, opts ...pulumi.ResourceOption) (*PeeringService, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.PeeringServiceName == nil {
		return nil, errors.New("missing required argument 'PeeringServiceName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &PeeringServiceArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:peering/latest:PeeringService"),
		},
		{
			Type: pulumi.String("azure-nextgen:peering/v20190801preview:PeeringService"),
		},
		{
			Type: pulumi.String("azure-nextgen:peering/v20190901preview:PeeringService"),
		},
		{
			Type: pulumi.String("azure-nextgen:peering/v20200101preview:PeeringService"),
		},
		{
			Type: pulumi.String("azure-nextgen:peering/v20201001:PeeringService"),
		},
	})
	opts = append(opts, aliases)
	var resource PeeringService
	err := ctx.RegisterResource("azure-nextgen:peering/v20200401:PeeringService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPeeringService gets an existing PeeringService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPeeringService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PeeringServiceState, opts ...pulumi.ResourceOption) (*PeeringService, error) {
	var resource PeeringService
	err := ctx.ReadResource("azure-nextgen:peering/v20200401:PeeringService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PeeringService resources.
type peeringServiceState struct {
	// The location of the resource.
	Location *string `pulumi:"location"`
	// The name of the resource.
	Name *string `pulumi:"name"`
	// The PeeringServiceLocation of the Customer.
	PeeringServiceLocation *string `pulumi:"peeringServiceLocation"`
	// The MAPS Provider Name.
	PeeringServiceProvider *string `pulumi:"peeringServiceProvider"`
	// The provisioning state of the resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The SKU that defines the type of the peering service.
	Sku *PeeringServiceSkuResponse `pulumi:"sku"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type *string `pulumi:"type"`
}

type PeeringServiceState struct {
	// The location of the resource.
	Location pulumi.StringPtrInput
	// The name of the resource.
	Name pulumi.StringPtrInput
	// The PeeringServiceLocation of the Customer.
	PeeringServiceLocation pulumi.StringPtrInput
	// The MAPS Provider Name.
	PeeringServiceProvider pulumi.StringPtrInput
	// The provisioning state of the resource.
	ProvisioningState pulumi.StringPtrInput
	// The SKU that defines the type of the peering service.
	Sku PeeringServiceSkuResponsePtrInput
	// The resource tags.
	Tags pulumi.StringMapInput
	// The type of the resource.
	Type pulumi.StringPtrInput
}

func (PeeringServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*peeringServiceState)(nil)).Elem()
}

type peeringServiceArgs struct {
	// The location of the resource.
	Location string `pulumi:"location"`
	// The PeeringServiceLocation of the Customer.
	PeeringServiceLocation *string `pulumi:"peeringServiceLocation"`
	// The name of the peering service.
	PeeringServiceName string `pulumi:"peeringServiceName"`
	// The MAPS Provider Name.
	PeeringServiceProvider *string `pulumi:"peeringServiceProvider"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The SKU that defines the type of the peering service.
	Sku *PeeringServiceSku `pulumi:"sku"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a PeeringService resource.
type PeeringServiceArgs struct {
	// The location of the resource.
	Location pulumi.StringInput
	// The PeeringServiceLocation of the Customer.
	PeeringServiceLocation pulumi.StringPtrInput
	// The name of the peering service.
	PeeringServiceName pulumi.StringInput
	// The MAPS Provider Name.
	PeeringServiceProvider pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The SKU that defines the type of the peering service.
	Sku PeeringServiceSkuPtrInput
	// The resource tags.
	Tags pulumi.StringMapInput
}

func (PeeringServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*peeringServiceArgs)(nil)).Elem()
}
