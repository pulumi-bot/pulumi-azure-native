// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190801

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Public certificate object
type WebAppPublicCertificate struct {
	pulumi.CustomResourceState

	// Kind of resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// PublicCertificate resource specific properties
	Properties PublicCertificateResponsePropertiesOutput `pulumi:"properties"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWebAppPublicCertificate registers a new resource with the given unique name, arguments, and options.
func NewWebAppPublicCertificate(ctx *pulumi.Context,
	name string, args *WebAppPublicCertificateArgs, opts ...pulumi.ResourceOption) (*WebAppPublicCertificate, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &WebAppPublicCertificateArgs{}
	}
	var resource WebAppPublicCertificate
	err := ctx.RegisterResource("azurerm:web/v20190801:WebAppPublicCertificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebAppPublicCertificate gets an existing WebAppPublicCertificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebAppPublicCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppPublicCertificateState, opts ...pulumi.ResourceOption) (*WebAppPublicCertificate, error) {
	var resource WebAppPublicCertificate
	err := ctx.ReadResource("azurerm:web/v20190801:WebAppPublicCertificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebAppPublicCertificate resources.
type webAppPublicCertificateState struct {
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Name.
	Name *string `pulumi:"name"`
	// PublicCertificate resource specific properties
	Properties *PublicCertificateResponseProperties `pulumi:"properties"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type WebAppPublicCertificateState struct {
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Resource Name.
	Name pulumi.StringPtrInput
	// PublicCertificate resource specific properties
	Properties PublicCertificateResponsePropertiesPtrInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (WebAppPublicCertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppPublicCertificateState)(nil)).Elem()
}

type webAppPublicCertificateArgs struct {
	// Public Certificate byte array
	Blob *string `pulumi:"blob"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Public certificate name.
	Name string `pulumi:"name"`
	// Public Certificate Location
	PublicCertificateLocation *string `pulumi:"publicCertificateLocation"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a WebAppPublicCertificate resource.
type WebAppPublicCertificateArgs struct {
	// Public Certificate byte array
	Blob pulumi.StringPtrInput
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Public certificate name.
	Name pulumi.StringInput
	// Public Certificate Location
	PublicCertificateLocation pulumi.StringPtrInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
}

func (WebAppPublicCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppPublicCertificateArgs)(nil)).Elem()
}
