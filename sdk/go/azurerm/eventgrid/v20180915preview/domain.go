// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180915preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// EventGrid Domain
//
// ## Example Usage
// ### Domains_CreateOrUpdate
//
// ```go
// package main
//
// import (
// 	eventgrid "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/eventgrid/v20180915preview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := eventgrid.NewDomain(ctx, "domain", &eventgrid.DomainArgs{
// 			DomainName:        pulumi.String("exampledomain1"),
// 			Location:          pulumi.String("westus2"),
// 			ResourceGroupName: pulumi.String("examplerg"),
// 			Tags: pulumi.StringMap{
// 				"tag1": pulumi.String("value1"),
// 				"tag2": pulumi.String("value2"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type Domain struct {
	pulumi.CustomResourceState

	// Endpoint for the domain.
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// This determines the format that Event Grid should expect for incoming events published to the domain.
	InputSchema pulumi.StringPtrOutput `pulumi:"inputSchema"`
	// Information about the InputSchemaMapping which specified the info about mapping event payload.
	InputSchemaMapping JsonInputSchemaMappingResponsePtrOutput `pulumi:"inputSchemaMapping"`
	// Location of the resource
	Location pulumi.StringOutput `pulumi:"location"`
	// Name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning state of the domain.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Tags of the resource
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Type of the resource
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDomain registers a new resource with the given unique name, arguments, and options.
func NewDomain(ctx *pulumi.Context,
	name string, args *DomainArgs, opts ...pulumi.ResourceOption) (*Domain, error) {
	if args == nil || args.DomainName == nil {
		return nil, errors.New("missing required argument 'DomainName'")
	}
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &DomainArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:eventgrid/latest:Domain"),
		},
		{
			Type: pulumi.String("azurerm:eventgrid/v20190201preview:Domain"),
		},
		{
			Type: pulumi.String("azurerm:eventgrid/v20190601:Domain"),
		},
		{
			Type: pulumi.String("azurerm:eventgrid/v20200101preview:Domain"),
		},
		{
			Type: pulumi.String("azurerm:eventgrid/v20200401preview:Domain"),
		},
		{
			Type: pulumi.String("azurerm:eventgrid/v20200601:Domain"),
		},
	})
	opts = append(opts, aliases)
	var resource Domain
	err := ctx.RegisterResource("azurerm:eventgrid/v20180915preview:Domain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomain gets an existing Domain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainState, opts ...pulumi.ResourceOption) (*Domain, error) {
	var resource Domain
	err := ctx.ReadResource("azurerm:eventgrid/v20180915preview:Domain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Domain resources.
type domainState struct {
	// Endpoint for the domain.
	Endpoint *string `pulumi:"endpoint"`
	// This determines the format that Event Grid should expect for incoming events published to the domain.
	InputSchema *string `pulumi:"inputSchema"`
	// Information about the InputSchemaMapping which specified the info about mapping event payload.
	InputSchemaMapping *JsonInputSchemaMappingResponse `pulumi:"inputSchemaMapping"`
	// Location of the resource
	Location *string `pulumi:"location"`
	// Name of the resource
	Name *string `pulumi:"name"`
	// Provisioning state of the domain.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Tags of the resource
	Tags map[string]string `pulumi:"tags"`
	// Type of the resource
	Type *string `pulumi:"type"`
}

type DomainState struct {
	// Endpoint for the domain.
	Endpoint pulumi.StringPtrInput
	// This determines the format that Event Grid should expect for incoming events published to the domain.
	InputSchema pulumi.StringPtrInput
	// Information about the InputSchemaMapping which specified the info about mapping event payload.
	InputSchemaMapping JsonInputSchemaMappingResponsePtrInput
	// Location of the resource
	Location pulumi.StringPtrInput
	// Name of the resource
	Name pulumi.StringPtrInput
	// Provisioning state of the domain.
	ProvisioningState pulumi.StringPtrInput
	// Tags of the resource
	Tags pulumi.StringMapInput
	// Type of the resource
	Type pulumi.StringPtrInput
}

func (DomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainState)(nil)).Elem()
}

type domainArgs struct {
	// Name of the domain
	DomainName string `pulumi:"domainName"`
	// This determines the format that Event Grid should expect for incoming events published to the domain.
	InputSchema *string `pulumi:"inputSchema"`
	// Information about the InputSchemaMapping which specified the info about mapping event payload.
	InputSchemaMapping *JsonInputSchemaMapping `pulumi:"inputSchemaMapping"`
	// Location of the resource
	Location string `pulumi:"location"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Tags of the resource
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Domain resource.
type DomainArgs struct {
	// Name of the domain
	DomainName pulumi.StringInput
	// This determines the format that Event Grid should expect for incoming events published to the domain.
	InputSchema pulumi.StringPtrInput
	// Information about the InputSchemaMapping which specified the info about mapping event payload.
	InputSchemaMapping JsonInputSchemaMappingPtrInput
	// Location of the resource
	Location pulumi.StringInput
	// The name of the resource group within the user's subscription.
	ResourceGroupName pulumi.StringInput
	// Tags of the resource
	Tags pulumi.StringMapInput
}

func (DomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainArgs)(nil)).Elem()
}
