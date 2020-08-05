// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20160901

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// App Service plan.
type AppServicePlan struct {
	pulumi.CustomResourceState

	// Kind of resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource Location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// AppServicePlan resource specific properties
	Properties AppServicePlanResponsePropertiesOutput `pulumi:"properties"`
	// Description of a SKU for a scalable resource.
	Sku SkuDescriptionResponsePtrOutput `pulumi:"sku"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAppServicePlan registers a new resource with the given unique name, arguments, and options.
func NewAppServicePlan(ctx *pulumi.Context,
	name string, args *AppServicePlanArgs, opts ...pulumi.ResourceOption) (*AppServicePlan, error) {
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
		args = &AppServicePlanArgs{}
	}
	var resource AppServicePlan
	err := ctx.RegisterResource("azurerm:web/v20160901:AppServicePlan", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppServicePlan gets an existing AppServicePlan resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppServicePlan(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppServicePlanState, opts ...pulumi.ResourceOption) (*AppServicePlan, error) {
	var resource AppServicePlan
	err := ctx.ReadResource("azurerm:web/v20160901:AppServicePlan", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppServicePlan resources.
type appServicePlanState struct {
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Location.
	Location *string `pulumi:"location"`
	// Resource Name.
	Name *string `pulumi:"name"`
	// AppServicePlan resource specific properties
	Properties *AppServicePlanResponseProperties `pulumi:"properties"`
	// Description of a SKU for a scalable resource.
	Sku *SkuDescriptionResponse `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type AppServicePlanState struct {
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Resource Location.
	Location pulumi.StringPtrInput
	// Resource Name.
	Name pulumi.StringPtrInput
	// AppServicePlan resource specific properties
	Properties AppServicePlanResponsePropertiesPtrInput
	// Description of a SKU for a scalable resource.
	Sku SkuDescriptionResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (AppServicePlanState) ElementType() reflect.Type {
	return reflect.TypeOf((*appServicePlanState)(nil)).Elem()
}

type appServicePlanArgs struct {
	// App Service plan administration site.
	AdminSiteName *string `pulumi:"adminSiteName"`
	// Specification for the App Service Environment to use for the App Service plan.
	HostingEnvironmentProfile *HostingEnvironmentProfile `pulumi:"hostingEnvironmentProfile"`
	// If <code>true</code>, this App Service Plan owns spot instances.
	IsSpot *bool `pulumi:"isSpot"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Location.
	Location string `pulumi:"location"`
	// Name for the App Service plan.
	Name string `pulumi:"name"`
	// If <code>true</code>, apps assigned to this App Service plan can be scaled independently.
	// If <code>false</code>, apps assigned to this App Service plan will scale to all instances of the plan.
	PerSiteScaling *bool `pulumi:"perSiteScaling"`
	// If Linux app service plan <code>true</code>, <code>false</code> otherwise.
	Reserved *bool `pulumi:"reserved"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Description of a SKU for a scalable resource.
	Sku *SkuDescription `pulumi:"sku"`
	// The time when the server farm expires. Valid only if it is a spot server farm.
	SpotExpirationTime *string `pulumi:"spotExpirationTime"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Scaling worker count.
	TargetWorkerCount *int `pulumi:"targetWorkerCount"`
	// Scaling worker size ID.
	TargetWorkerSizeId *int `pulumi:"targetWorkerSizeId"`
	// Target worker tier assigned to the App Service plan.
	WorkerTierName *string `pulumi:"workerTierName"`
}

// The set of arguments for constructing a AppServicePlan resource.
type AppServicePlanArgs struct {
	// App Service plan administration site.
	AdminSiteName pulumi.StringPtrInput
	// Specification for the App Service Environment to use for the App Service plan.
	HostingEnvironmentProfile HostingEnvironmentProfilePtrInput
	// If <code>true</code>, this App Service Plan owns spot instances.
	IsSpot pulumi.BoolPtrInput
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Resource Location.
	Location pulumi.StringInput
	// Name for the App Service plan.
	Name pulumi.StringInput
	// If <code>true</code>, apps assigned to this App Service plan can be scaled independently.
	// If <code>false</code>, apps assigned to this App Service plan will scale to all instances of the plan.
	PerSiteScaling pulumi.BoolPtrInput
	// If Linux app service plan <code>true</code>, <code>false</code> otherwise.
	Reserved pulumi.BoolPtrInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
	// Description of a SKU for a scalable resource.
	Sku SkuDescriptionPtrInput
	// The time when the server farm expires. Valid only if it is a spot server farm.
	SpotExpirationTime pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Scaling worker count.
	TargetWorkerCount pulumi.IntPtrInput
	// Scaling worker size ID.
	TargetWorkerSizeId pulumi.IntPtrInput
	// Target worker tier assigned to the App Service plan.
	WorkerTierName pulumi.StringPtrInput
}

func (AppServicePlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appServicePlanArgs)(nil)).Elem()
}
