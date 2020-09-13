// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200501

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A Transform encapsulates the rules or instructions for generating desired outputs from input media, such as by transcoding or by extracting insights. After the Transform is created, it can be applied to input media by creating Jobs.
//
// ## Example Usage
// ### Create or update a Transform
//
// ```go
//
// ```
type Transform struct {
	pulumi.CustomResourceState

	// The UTC date and time when the Transform was created, in 'YYYY-MM-DDThh:mm:ssZ' format.
	Created pulumi.StringOutput `pulumi:"created"`
	// An optional verbose description of the Transform.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The UTC date and time when the Transform was last updated, in 'YYYY-MM-DDThh:mm:ssZ' format.
	LastModified pulumi.StringOutput `pulumi:"lastModified"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// An array of one or more TransformOutputs that the Transform should generate.
	Outputs TransformOutputResponseArrayOutput `pulumi:"outputs"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewTransform registers a new resource with the given unique name, arguments, and options.
func NewTransform(ctx *pulumi.Context,
	name string, args *TransformArgs, opts ...pulumi.ResourceOption) (*Transform, error) {
	if args == nil || args.AccountName == nil {
		return nil, errors.New("missing required argument 'AccountName'")
	}
	if args == nil || args.Outputs == nil {
		return nil, errors.New("missing required argument 'Outputs'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.TransformName == nil {
		return nil, errors.New("missing required argument 'TransformName'")
	}
	if args == nil {
		args = &TransformArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:media/latest:Transform"),
		},
		{
			Type: pulumi.String("azurerm:media/v20180330preview:Transform"),
		},
		{
			Type: pulumi.String("azurerm:media/v20180601preview:Transform"),
		},
		{
			Type: pulumi.String("azurerm:media/v20180701:Transform"),
		},
	})
	opts = append(opts, aliases)
	var resource Transform
	err := ctx.RegisterResource("azurerm:media/v20200501:Transform", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTransform gets an existing Transform resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTransform(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TransformState, opts ...pulumi.ResourceOption) (*Transform, error) {
	var resource Transform
	err := ctx.ReadResource("azurerm:media/v20200501:Transform", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Transform resources.
type transformState struct {
	// The UTC date and time when the Transform was created, in 'YYYY-MM-DDThh:mm:ssZ' format.
	Created *string `pulumi:"created"`
	// An optional verbose description of the Transform.
	Description *string `pulumi:"description"`
	// The UTC date and time when the Transform was last updated, in 'YYYY-MM-DDThh:mm:ssZ' format.
	LastModified *string `pulumi:"lastModified"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// An array of one or more TransformOutputs that the Transform should generate.
	Outputs []TransformOutputResponse `pulumi:"outputs"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `pulumi:"type"`
}

type TransformState struct {
	// The UTC date and time when the Transform was created, in 'YYYY-MM-DDThh:mm:ssZ' format.
	Created pulumi.StringPtrInput
	// An optional verbose description of the Transform.
	Description pulumi.StringPtrInput
	// The UTC date and time when the Transform was last updated, in 'YYYY-MM-DDThh:mm:ssZ' format.
	LastModified pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// An array of one or more TransformOutputs that the Transform should generate.
	Outputs TransformOutputResponseArrayInput
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringPtrInput
}

func (TransformState) ElementType() reflect.Type {
	return reflect.TypeOf((*transformState)(nil)).Elem()
}

type transformArgs struct {
	// The Media Services account name.
	AccountName string `pulumi:"accountName"`
	// An optional verbose description of the Transform.
	Description *string `pulumi:"description"`
	// An array of one or more TransformOutputs that the Transform should generate.
	Outputs []TransformOutput `pulumi:"outputs"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Transform name.
	TransformName string `pulumi:"transformName"`
}

// The set of arguments for constructing a Transform resource.
type TransformArgs struct {
	// The Media Services account name.
	AccountName pulumi.StringInput
	// An optional verbose description of the Transform.
	Description pulumi.StringPtrInput
	// An array of one or more TransformOutputs that the Transform should generate.
	Outputs TransformOutputArrayInput
	// The name of the resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// The Transform name.
	TransformName pulumi.StringInput
}

func (TransformArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*transformArgs)(nil)).Elem()
}
