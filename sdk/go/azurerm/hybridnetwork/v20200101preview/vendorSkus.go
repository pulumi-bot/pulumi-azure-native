// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200101preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Sku sub resource.
type VendorSkus struct {
	pulumi.CustomResourceState

	// Sku deployment mode.
	DeploymentMode pulumi.StringPtrOutput `pulumi:"deploymentMode"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The parameters for the managed application to be supplied by vendor.
	ManagedApplicationParameters pulumi.MapOutput `pulumi:"managedApplicationParameters"`
	// The template for the managed application deployment.
	ManagedApplicationTemplate pulumi.MapOutput `pulumi:"managedApplicationTemplate"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Indicates if the vendor sku is in preview mode.
	Preview pulumi.BoolPtrOutput `pulumi:"preview"`
	// The provisioning state of the vendor sku sub resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Sku type.
	SkuType pulumi.StringPtrOutput `pulumi:"skuType"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// The template definition of the virtual network function.
	VirtualNetworkFunctionTemplate VirtualNetworkFunctionTemplateResponsePtrOutput `pulumi:"virtualNetworkFunctionTemplate"`
}

// NewVendorSkus registers a new resource with the given unique name, arguments, and options.
func NewVendorSkus(ctx *pulumi.Context,
	name string, args *VendorSkusArgs, opts ...pulumi.ResourceOption) (*VendorSkus, error) {
	if args == nil || args.SkuName == nil {
		return nil, errors.New("missing required argument 'SkuName'")
	}
	if args == nil || args.VendorName == nil {
		return nil, errors.New("missing required argument 'VendorName'")
	}
	if args == nil {
		args = &VendorSkusArgs{}
	}
	var resource VendorSkus
	err := ctx.RegisterResource("azurerm:hybridnetwork/v20200101preview:VendorSkus", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVendorSkus gets an existing VendorSkus resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVendorSkus(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VendorSkusState, opts ...pulumi.ResourceOption) (*VendorSkus, error) {
	var resource VendorSkus
	err := ctx.ReadResource("azurerm:hybridnetwork/v20200101preview:VendorSkus", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VendorSkus resources.
type vendorSkusState struct {
	// Sku deployment mode.
	DeploymentMode *string `pulumi:"deploymentMode"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The parameters for the managed application to be supplied by vendor.
	ManagedApplicationParameters map[string]interface{} `pulumi:"managedApplicationParameters"`
	// The template for the managed application deployment.
	ManagedApplicationTemplate map[string]interface{} `pulumi:"managedApplicationTemplate"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Indicates if the vendor sku is in preview mode.
	Preview *bool `pulumi:"preview"`
	// The provisioning state of the vendor sku sub resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Sku type.
	SkuType *string `pulumi:"skuType"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
	// The template definition of the virtual network function.
	VirtualNetworkFunctionTemplate *VirtualNetworkFunctionTemplateResponse `pulumi:"virtualNetworkFunctionTemplate"`
}

type VendorSkusState struct {
	// Sku deployment mode.
	DeploymentMode pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The parameters for the managed application to be supplied by vendor.
	ManagedApplicationParameters pulumi.MapInput
	// The template for the managed application deployment.
	ManagedApplicationTemplate pulumi.MapInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Indicates if the vendor sku is in preview mode.
	Preview pulumi.BoolPtrInput
	// The provisioning state of the vendor sku sub resource.
	ProvisioningState pulumi.StringPtrInput
	// Sku type.
	SkuType pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
	// The template definition of the virtual network function.
	VirtualNetworkFunctionTemplate VirtualNetworkFunctionTemplateResponsePtrInput
}

func (VendorSkusState) ElementType() reflect.Type {
	return reflect.TypeOf((*vendorSkusState)(nil)).Elem()
}

type vendorSkusArgs struct {
	// Sku deployment mode.
	DeploymentMode *string `pulumi:"deploymentMode"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The parameters for the managed application to be supplied by vendor.
	ManagedApplicationParameters map[string]interface{} `pulumi:"managedApplicationParameters"`
	// The template for the managed application deployment.
	ManagedApplicationTemplate map[string]interface{} `pulumi:"managedApplicationTemplate"`
	// Indicates if the vendor sku is in preview mode.
	Preview *bool `pulumi:"preview"`
	// The name of the sku.
	SkuName string `pulumi:"skuName"`
	// Sku type.
	SkuType *string `pulumi:"skuType"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The name of the vendor.
	VendorName string `pulumi:"vendorName"`
	// The template definition of the virtual network function.
	VirtualNetworkFunctionTemplate *VirtualNetworkFunctionTemplate `pulumi:"virtualNetworkFunctionTemplate"`
}

// The set of arguments for constructing a VendorSkus resource.
type VendorSkusArgs struct {
	// Sku deployment mode.
	DeploymentMode pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The parameters for the managed application to be supplied by vendor.
	ManagedApplicationParameters pulumi.MapInput
	// The template for the managed application deployment.
	ManagedApplicationTemplate pulumi.MapInput
	// Indicates if the vendor sku is in preview mode.
	Preview pulumi.BoolPtrInput
	// The name of the sku.
	SkuName pulumi.StringInput
	// Sku type.
	SkuType pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The name of the vendor.
	VendorName pulumi.StringInput
	// The template definition of the virtual network function.
	VirtualNetworkFunctionTemplate VirtualNetworkFunctionTemplatePtrInput
}

func (VendorSkusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vendorSkusArgs)(nil)).Elem()
}
