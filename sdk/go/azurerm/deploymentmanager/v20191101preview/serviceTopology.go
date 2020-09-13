// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191101preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The resource representation of a service topology.
//
// ## Example Usage
// ### Create a topology with Artifact Source
//
// ```go
// package main
//
// import (
// 	deploymentmanager "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/deploymentmanager/v20191101preview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := deploymentmanager.NewServiceTopology(ctx, "serviceTopology", &deploymentmanager.ServiceTopologyArgs{
// 			ArtifactSourceId:    pulumi.String("Microsoft.DeploymentManager/artifactSources/myArtifactSource"),
// 			Location:            pulumi.String("centralus"),
// 			ResourceGroupName:   pulumi.String("myResourceGroup"),
// 			ServiceTopologyName: pulumi.String("myTopology"),
// 			Tags:                nil,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
// ### Create a topology without Artifact Source
//
// ```go
// package main
//
// import (
// 	deploymentmanager "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/deploymentmanager/v20191101preview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := deploymentmanager.NewServiceTopology(ctx, "serviceTopology", &deploymentmanager.ServiceTopologyArgs{
// 			Location:            pulumi.String("centralus"),
// 			ResourceGroupName:   pulumi.String("myResourceGroup"),
// 			ServiceTopologyName: pulumi.String("myTopology"),
// 			Tags:                nil,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type ServiceTopology struct {
	pulumi.CustomResourceState

	// The resource Id of the artifact source that contains the artifacts that can be referenced in the service units.
	ArtifactSourceId pulumi.StringPtrOutput `pulumi:"artifactSourceId"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewServiceTopology registers a new resource with the given unique name, arguments, and options.
func NewServiceTopology(ctx *pulumi.Context,
	name string, args *ServiceTopologyArgs, opts ...pulumi.ResourceOption) (*ServiceTopology, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServiceTopologyName == nil {
		return nil, errors.New("missing required argument 'ServiceTopologyName'")
	}
	if args == nil {
		args = &ServiceTopologyArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:deploymentmanager/v20180901preview:ServiceTopology"),
		},
	})
	opts = append(opts, aliases)
	var resource ServiceTopology
	err := ctx.RegisterResource("azurerm:deploymentmanager/v20191101preview:ServiceTopology", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceTopology gets an existing ServiceTopology resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceTopology(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceTopologyState, opts ...pulumi.ResourceOption) (*ServiceTopology, error) {
	var resource ServiceTopology
	err := ctx.ReadResource("azurerm:deploymentmanager/v20191101preview:ServiceTopology", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceTopology resources.
type serviceTopologyState struct {
	// The resource Id of the artifact source that contains the artifacts that can be referenced in the service units.
	ArtifactSourceId *string `pulumi:"artifactSourceId"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `pulumi:"type"`
}

type ServiceTopologyState struct {
	// The resource Id of the artifact source that contains the artifacts that can be referenced in the service units.
	ArtifactSourceId pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringPtrInput
}

func (ServiceTopologyState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceTopologyState)(nil)).Elem()
}

type serviceTopologyArgs struct {
	// The resource Id of the artifact source that contains the artifacts that can be referenced in the service units.
	ArtifactSourceId *string `pulumi:"artifactSourceId"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the service topology .
	ServiceTopologyName string `pulumi:"serviceTopologyName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ServiceTopology resource.
type ServiceTopologyArgs struct {
	// The resource Id of the artifact source that contains the artifacts that can be referenced in the service units.
	ArtifactSourceId pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the service topology .
	ServiceTopologyName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (ServiceTopologyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceTopologyArgs)(nil)).Elem()
}
