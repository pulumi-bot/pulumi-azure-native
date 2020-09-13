// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20160601

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// API connection
//
// ## Example Usage
// ### Replace a connection
//
// ```go
// package main
//
// import (
// 	web "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/web/v20160601"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := web.NewConnection(ctx, "connection", &web.ConnectionArgs{
// 			ConnectionName:    pulumi.String("testManagedApi"),
// 			ResourceGroupName: pulumi.String("testResourceGroup"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type Connection struct {
	pulumi.CustomResourceState

	// Resource ETag
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Resource location
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name
	Name       pulumi.StringOutput                             `pulumi:"name"`
	Properties ApiConnectionDefinitionResponsePropertiesOutput `pulumi:"properties"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewConnection registers a new resource with the given unique name, arguments, and options.
func NewConnection(ctx *pulumi.Context,
	name string, args *ConnectionArgs, opts ...pulumi.ResourceOption) (*Connection, error) {
	if args == nil || args.ConnectionName == nil {
		return nil, errors.New("missing required argument 'ConnectionName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &ConnectionArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:web/latest:Connection"),
		},
		{
			Type: pulumi.String("azurerm:web/v20150801preview:Connection"),
		},
	})
	opts = append(opts, aliases)
	var resource Connection
	err := ctx.RegisterResource("azurerm:web/v20160601:Connection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConnection gets an existing Connection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectionState, opts ...pulumi.ResourceOption) (*Connection, error) {
	var resource Connection
	err := ctx.ReadResource("azurerm:web/v20160601:Connection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Connection resources.
type connectionState struct {
	// Resource ETag
	Etag *string `pulumi:"etag"`
	// Resource location
	Location *string `pulumi:"location"`
	// Resource name
	Name       *string                                    `pulumi:"name"`
	Properties *ApiConnectionDefinitionResponseProperties `pulumi:"properties"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
}

type ConnectionState struct {
	// Resource ETag
	Etag pulumi.StringPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// Resource name
	Name       pulumi.StringPtrInput
	Properties ApiConnectionDefinitionResponsePropertiesPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (ConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionState)(nil)).Elem()
}

type connectionArgs struct {
	// Connection name
	ConnectionName string `pulumi:"connectionName"`
	// Resource ETag
	Etag *string `pulumi:"etag"`
	// Resource location
	Location   *string                            `pulumi:"location"`
	Properties *ApiConnectionDefinitionProperties `pulumi:"properties"`
	// The resource group
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Connection resource.
type ConnectionArgs struct {
	// Connection name
	ConnectionName pulumi.StringInput
	// Resource ETag
	Etag pulumi.StringPtrInput
	// Resource location
	Location   pulumi.StringPtrInput
	Properties ApiConnectionDefinitionPropertiesPtrInput
	// The resource group
	ResourceGroupName pulumi.StringInput
	// Resource tags
	Tags pulumi.StringMapInput
}

func (ConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionArgs)(nil)).Elem()
}
