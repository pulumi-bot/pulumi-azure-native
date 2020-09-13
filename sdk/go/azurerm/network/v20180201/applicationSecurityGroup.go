// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180201

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// An application security group in a resource group.
//
// ## Example Usage
// ### Create application security group
//
// ```go
// package main
//
// import (
// 	network "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/network/v20180201"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := network.NewApplicationSecurityGroup(ctx, "applicationSecurityGroup", &network.ApplicationSecurityGroupArgs{
// 			ApplicationSecurityGroupName: pulumi.String("test-asg"),
// 			Location:                     pulumi.String("westus"),
// 			ResourceGroupName:            pulumi.String("rg1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type ApplicationSecurityGroup struct {
	pulumi.CustomResourceState

	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the application security group resource. Possible values are: 'Succeeded', 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The resource GUID property of the application security group resource. It uniquely identifies a resource, even if the user changes its name or migrate the resource across subscriptions or resource groups.
	ResourceGuid pulumi.StringOutput `pulumi:"resourceGuid"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewApplicationSecurityGroup registers a new resource with the given unique name, arguments, and options.
func NewApplicationSecurityGroup(ctx *pulumi.Context,
	name string, args *ApplicationSecurityGroupArgs, opts ...pulumi.ResourceOption) (*ApplicationSecurityGroup, error) {
	if args == nil || args.ApplicationSecurityGroupName == nil {
		return nil, errors.New("missing required argument 'ApplicationSecurityGroupName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &ApplicationSecurityGroupArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:network/latest:ApplicationSecurityGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20170901:ApplicationSecurityGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20171001:ApplicationSecurityGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20171101:ApplicationSecurityGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180101:ApplicationSecurityGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180401:ApplicationSecurityGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180601:ApplicationSecurityGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180701:ApplicationSecurityGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20180801:ApplicationSecurityGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20181001:ApplicationSecurityGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20181101:ApplicationSecurityGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20181201:ApplicationSecurityGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190201:ApplicationSecurityGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190401:ApplicationSecurityGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190601:ApplicationSecurityGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190701:ApplicationSecurityGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190801:ApplicationSecurityGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20190901:ApplicationSecurityGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20191101:ApplicationSecurityGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20191201:ApplicationSecurityGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200301:ApplicationSecurityGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200401:ApplicationSecurityGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200501:ApplicationSecurityGroup"),
		},
		{
			Type: pulumi.String("azurerm:network/v20200601:ApplicationSecurityGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource ApplicationSecurityGroup
	err := ctx.RegisterResource("azurerm:network/v20180201:ApplicationSecurityGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationSecurityGroup gets an existing ApplicationSecurityGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationSecurityGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationSecurityGroupState, opts ...pulumi.ResourceOption) (*ApplicationSecurityGroup, error) {
	var resource ApplicationSecurityGroup
	err := ctx.ReadResource("azurerm:network/v20180201:ApplicationSecurityGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationSecurityGroup resources.
type applicationSecurityGroupState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name *string `pulumi:"name"`
	// The provisioning state of the application security group resource. Possible values are: 'Succeeded', 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The resource GUID property of the application security group resource. It uniquely identifies a resource, even if the user changes its name or migrate the resource across subscriptions or resource groups.
	ResourceGuid *string `pulumi:"resourceGuid"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type ApplicationSecurityGroupState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// The provisioning state of the application security group resource. Possible values are: 'Succeeded', 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState pulumi.StringPtrInput
	// The resource GUID property of the application security group resource. It uniquely identifies a resource, even if the user changes its name or migrate the resource across subscriptions or resource groups.
	ResourceGuid pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (ApplicationSecurityGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationSecurityGroupState)(nil)).Elem()
}

type applicationSecurityGroupArgs struct {
	// The name of the application security group.
	ApplicationSecurityGroupName string `pulumi:"applicationSecurityGroupName"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ApplicationSecurityGroup resource.
type ApplicationSecurityGroupArgs struct {
	// The name of the application security group.
	ApplicationSecurityGroupName pulumi.StringInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (ApplicationSecurityGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationSecurityGroupArgs)(nil)).Elem()
}
