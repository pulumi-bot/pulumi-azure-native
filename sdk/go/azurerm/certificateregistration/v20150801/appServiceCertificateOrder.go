// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150801

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// SSL certificate purchase order.
type AppServiceCertificateOrder struct {
	pulumi.CustomResourceState

	// Kind of resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource Location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// AppServiceCertificateOrder resource specific properties
	Properties AppServiceCertificateOrderResponsePropertiesOutput `pulumi:"properties"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAppServiceCertificateOrder registers a new resource with the given unique name, arguments, and options.
func NewAppServiceCertificateOrder(ctx *pulumi.Context,
	name string, args *AppServiceCertificateOrderArgs, opts ...pulumi.ResourceOption) (*AppServiceCertificateOrder, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ProductType == nil {
		return nil, errors.New("missing required argument 'ProductType'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &AppServiceCertificateOrderArgs{}
	}
	var resource AppServiceCertificateOrder
	err := ctx.RegisterResource("azurerm:certificateregistration/v20150801:AppServiceCertificateOrder", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppServiceCertificateOrder gets an existing AppServiceCertificateOrder resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppServiceCertificateOrder(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppServiceCertificateOrderState, opts ...pulumi.ResourceOption) (*AppServiceCertificateOrder, error) {
	var resource AppServiceCertificateOrder
	err := ctx.ReadResource("azurerm:certificateregistration/v20150801:AppServiceCertificateOrder", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppServiceCertificateOrder resources.
type appServiceCertificateOrderState struct {
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Location.
	Location *string `pulumi:"location"`
	// Resource Name.
	Name *string `pulumi:"name"`
	// AppServiceCertificateOrder resource specific properties
	Properties *AppServiceCertificateOrderResponseProperties `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type AppServiceCertificateOrderState struct {
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Resource Location.
	Location pulumi.StringPtrInput
	// Resource Name.
	Name pulumi.StringPtrInput
	// AppServiceCertificateOrder resource specific properties
	Properties AppServiceCertificateOrderResponsePropertiesPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (AppServiceCertificateOrderState) ElementType() reflect.Type {
	return reflect.TypeOf((*appServiceCertificateOrderState)(nil)).Elem()
}

type appServiceCertificateOrderArgs struct {
	// <code>true</code> if the certificate should be automatically renewed when it expires; otherwise, <code>false</code>.
	AutoRenew *bool `pulumi:"autoRenew"`
	// State of the Key Vault secret.
	Certificates map[string]AppServiceCertificate `pulumi:"certificates"`
	// Last CSR that was created for this order.
	Csr *string `pulumi:"csr"`
	// Certificate distinguished name.
	DistinguishedName *string `pulumi:"distinguishedName"`
	// Certificate key size.
	KeySize *int `pulumi:"keySize"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Location.
	Location string `pulumi:"location"`
	// Name of the certificate order.
	Name string `pulumi:"name"`
	// Certificate product type.
	ProductType string `pulumi:"productType"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Duration in years (must be between 1 and 3).
	ValidityInYears *int `pulumi:"validityInYears"`
}

// The set of arguments for constructing a AppServiceCertificateOrder resource.
type AppServiceCertificateOrderArgs struct {
	// <code>true</code> if the certificate should be automatically renewed when it expires; otherwise, <code>false</code>.
	AutoRenew pulumi.BoolPtrInput
	// State of the Key Vault secret.
	Certificates AppServiceCertificateMapInput
	// Last CSR that was created for this order.
	Csr pulumi.StringPtrInput
	// Certificate distinguished name.
	DistinguishedName pulumi.StringPtrInput
	// Certificate key size.
	KeySize pulumi.IntPtrInput
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Resource Location.
	Location pulumi.StringInput
	// Name of the certificate order.
	Name pulumi.StringInput
	// Certificate product type.
	ProductType pulumi.StringInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Duration in years (must be between 1 and 3).
	ValidityInYears pulumi.IntPtrInput
}

func (AppServiceCertificateOrderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appServiceCertificateOrderArgs)(nil)).Elem()
}
