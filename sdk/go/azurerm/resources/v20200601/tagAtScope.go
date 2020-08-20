// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200601

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Wrapper resource for tags API requests and responses.
type TagAtScope struct {
	pulumi.CustomResourceState

	// The name of the tags wrapper resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The set of tags.
	Properties TagsResponseOutput `pulumi:"properties"`
	// The type of the tags wrapper resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewTagAtScope registers a new resource with the given unique name, arguments, and options.
func NewTagAtScope(ctx *pulumi.Context,
	name string, args *TagAtScopeArgs, opts ...pulumi.ResourceOption) (*TagAtScope, error) {
	if args == nil || args.Properties == nil {
		return nil, errors.New("missing required argument 'Properties'")
	}
	if args == nil || args.Scope == nil {
		return nil, errors.New("missing required argument 'Scope'")
	}
	if args == nil {
		args = &TagAtScopeArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:resources/v20191001:TagAtScope"),
		},
	})
	opts = append(opts, aliases)
	var resource TagAtScope
	err := ctx.RegisterResource("azurerm:resources/v20200601:TagAtScope", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTagAtScope gets an existing TagAtScope resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTagAtScope(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TagAtScopeState, opts ...pulumi.ResourceOption) (*TagAtScope, error) {
	var resource TagAtScope
	err := ctx.ReadResource("azurerm:resources/v20200601:TagAtScope", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TagAtScope resources.
type tagAtScopeState struct {
	// The name of the tags wrapper resource.
	Name *string `pulumi:"name"`
	// The set of tags.
	Properties *TagsResponse `pulumi:"properties"`
	// The type of the tags wrapper resource.
	Type *string `pulumi:"type"`
}

type TagAtScopeState struct {
	// The name of the tags wrapper resource.
	Name pulumi.StringPtrInput
	// The set of tags.
	Properties TagsResponsePtrInput
	// The type of the tags wrapper resource.
	Type pulumi.StringPtrInput
}

func (TagAtScopeState) ElementType() reflect.Type {
	return reflect.TypeOf((*tagAtScopeState)(nil)).Elem()
}

type tagAtScopeArgs struct {
	// The set of tags.
	Properties Tags `pulumi:"properties"`
	// The resource scope.
	Scope string `pulumi:"scope"`
}

// The set of arguments for constructing a TagAtScope resource.
type TagAtScopeArgs struct {
	// The set of tags.
	Properties TagsInput
	// The resource scope.
	Scope pulumi.StringInput
}

func (TagAtScopeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tagAtScopeArgs)(nil)).Elem()
}
