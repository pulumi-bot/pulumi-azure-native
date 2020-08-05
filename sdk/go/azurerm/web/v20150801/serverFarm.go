// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150801

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// App Service Plan Model
type ServerFarm struct {
	pulumi.CustomResourceState

	// Kind of resource
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource Location
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource Name
	Name       pulumi.StringPtrOutput                        `pulumi:"name"`
	Properties ServerFarmWithRichSkuResponsePropertiesOutput `pulumi:"properties"`
	// Describes a sku for a scalable resource
	Sku SkuDescriptionResponsePtrOutput `pulumi:"sku"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewServerFarm registers a new resource with the given unique name, arguments, and options.
func NewServerFarm(ctx *pulumi.Context,
	name string, args *ServerFarmArgs, opts ...pulumi.ResourceOption) (*ServerFarm, error) {
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
		args = &ServerFarmArgs{}
	}
	var resource ServerFarm
	err := ctx.RegisterResource("azurerm:web/v20150801:ServerFarm", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServerFarm gets an existing ServerFarm resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServerFarm(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerFarmState, opts ...pulumi.ResourceOption) (*ServerFarm, error) {
	var resource ServerFarm
	err := ctx.ReadResource("azurerm:web/v20150801:ServerFarm", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServerFarm resources.
type serverFarmState struct {
	// Kind of resource
	Kind *string `pulumi:"kind"`
	// Resource Location
	Location *string `pulumi:"location"`
	// Resource Name
	Name       *string                                  `pulumi:"name"`
	Properties *ServerFarmWithRichSkuResponseProperties `pulumi:"properties"`
	// Describes a sku for a scalable resource
	Sku *SkuDescriptionResponse `pulumi:"sku"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
}

type ServerFarmState struct {
	// Kind of resource
	Kind pulumi.StringPtrInput
	// Resource Location
	Location pulumi.StringPtrInput
	// Resource Name
	Name       pulumi.StringPtrInput
	Properties ServerFarmWithRichSkuResponsePropertiesPtrInput
	// Describes a sku for a scalable resource
	Sku SkuDescriptionResponsePtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (ServerFarmState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverFarmState)(nil)).Elem()
}

type serverFarmArgs struct {
	// App Service Plan administration site
	AdminSiteName *string `pulumi:"adminSiteName"`
	// OBSOLETE: If true, allow pending state for App Service Plan
	AllowPendingState *bool `pulumi:"allowPendingState"`
	// Specification for the hosting environment (App Service Environment) to use for the App Service Plan
	HostingEnvironmentProfile *HostingEnvironmentProfile `pulumi:"hostingEnvironmentProfile"`
	// Resource Id
	Id *string `pulumi:"id"`
	// Kind of resource
	Kind *string `pulumi:"kind"`
	// Resource Location
	Location string `pulumi:"location"`
	// Maximum number of instances that can be assigned to this App Service Plan
	MaximumNumberOfWorkers *int `pulumi:"maximumNumberOfWorkers"`
	// Resource Name
	Name string `pulumi:"name"`
	// If True apps assigned to this App Service Plan can be scaled independently
	//             If False apps assigned to this App Service Plan will scale to all instances of the plan
	PerSiteScaling *bool `pulumi:"perSiteScaling"`
	// Enables creation of a Linux App Service Plan
	Reserved *bool `pulumi:"reserved"`
	// Name of resource group
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Describes a sku for a scalable resource
	Sku *SkuDescription `pulumi:"sku"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
	// Target worker tier assigned to the App Service Plan
	WorkerTierName *string `pulumi:"workerTierName"`
}

// The set of arguments for constructing a ServerFarm resource.
type ServerFarmArgs struct {
	// App Service Plan administration site
	AdminSiteName pulumi.StringPtrInput
	// OBSOLETE: If true, allow pending state for App Service Plan
	AllowPendingState pulumi.BoolPtrInput
	// Specification for the hosting environment (App Service Environment) to use for the App Service Plan
	HostingEnvironmentProfile HostingEnvironmentProfilePtrInput
	// Resource Id
	Id pulumi.StringPtrInput
	// Kind of resource
	Kind pulumi.StringPtrInput
	// Resource Location
	Location pulumi.StringInput
	// Maximum number of instances that can be assigned to this App Service Plan
	MaximumNumberOfWorkers pulumi.IntPtrInput
	// Resource Name
	Name pulumi.StringInput
	// If True apps assigned to this App Service Plan can be scaled independently
	//             If False apps assigned to this App Service Plan will scale to all instances of the plan
	PerSiteScaling pulumi.BoolPtrInput
	// Enables creation of a Linux App Service Plan
	Reserved pulumi.BoolPtrInput
	// Name of resource group
	ResourceGroupName pulumi.StringInput
	// Describes a sku for a scalable resource
	Sku SkuDescriptionPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Resource type
	Type pulumi.StringPtrInput
	// Target worker tier assigned to the App Service Plan
	WorkerTierName pulumi.StringPtrInput
}

func (ServerFarmArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverFarmArgs)(nil)).Elem()
}
