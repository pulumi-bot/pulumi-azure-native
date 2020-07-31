// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150801

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Certificate signing request object
type CertificateCsr struct {
	pulumi.CustomResourceState

	// Kind of resource
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource Location
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource Name
	Name       pulumi.StringPtrOutput      `pulumi:"name"`
	Properties CsrResponsePropertiesOutput `pulumi:"properties"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewCertificateCsr registers a new resource with the given unique name, arguments, and options.
func NewCertificateCsr(ctx *pulumi.Context,
	name string, args *CertificateCsrArgs, opts ...pulumi.ResourceOption) (*CertificateCsr, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &CertificateCsrArgs{}
	}
	var resource CertificateCsr
	err := ctx.RegisterResource("azurerm:web/v20150801:CertificateCsr", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCertificateCsr gets an existing CertificateCsr resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertificateCsr(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateCsrState, opts ...pulumi.ResourceOption) (*CertificateCsr, error) {
	var resource CertificateCsr
	err := ctx.ReadResource("azurerm:web/v20150801:CertificateCsr", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CertificateCsr resources.
type certificateCsrState struct {
	// Kind of resource
	Kind *string `pulumi:"kind"`
	// Resource Location
	Location *string `pulumi:"location"`
	// Resource Name
	Name       *string                `pulumi:"name"`
	Properties *CsrResponseProperties `pulumi:"properties"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
}

type CertificateCsrState struct {
	// Kind of resource
	Kind pulumi.StringPtrInput
	// Resource Location
	Location pulumi.StringPtrInput
	// Resource Name
	Name       pulumi.StringPtrInput
	Properties CsrResponsePropertiesPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (CertificateCsrState) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateCsrState)(nil)).Elem()
}

type certificateCsrArgs struct {
	// Resource Id
	Id *string `pulumi:"id"`
	// Kind of resource
	Kind *string `pulumi:"kind"`
	// Resource Location
	Location string `pulumi:"location"`
	// Resource Name
	Name       string         `pulumi:"name"`
	Properties *CsrProperties `pulumi:"properties"`
	// Name of the resource group
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a CertificateCsr resource.
type CertificateCsrArgs struct {
	// Resource Id
	Id pulumi.StringPtrInput
	// Kind of resource
	Kind pulumi.StringPtrInput
	// Resource Location
	Location pulumi.StringInput
	// Resource Name
	Name       pulumi.StringInput
	Properties CsrPropertiesPtrInput
	// Name of the resource group
	ResourceGroupName pulumi.StringInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (CertificateCsrArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateCsrArgs)(nil)).Elem()
}