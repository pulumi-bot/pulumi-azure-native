// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200901

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// App Service plan.
type AppServicePlan struct {
	pulumi.CustomResourceState

	// The time when the server farm free offer expires.
	FreeOfferExpirationTime pulumi.StringPtrOutput `pulumi:"freeOfferExpirationTime"`
	// Geographical location for the App Service plan.
	GeoRegion pulumi.StringOutput `pulumi:"geoRegion"`
	// Specification for the App Service Environment to use for the App Service plan.
	HostingEnvironmentProfile HostingEnvironmentProfileResponsePtrOutput `pulumi:"hostingEnvironmentProfile"`
	// If Hyper-V container app service plan <code>true</code>, <code>false</code> otherwise.
	HyperV pulumi.BoolPtrOutput `pulumi:"hyperV"`
	// If <code>true</code>, this App Service Plan owns spot instances.
	IsSpot pulumi.BoolPtrOutput `pulumi:"isSpot"`
	// Obsolete: If Hyper-V container app service plan <code>true</code>, <code>false</code> otherwise.
	IsXenon pulumi.BoolPtrOutput `pulumi:"isXenon"`
	// Kind of resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource Location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Maximum number of total workers allowed for this ElasticScaleEnabled App Service Plan
	MaximumElasticWorkerCount pulumi.IntPtrOutput `pulumi:"maximumElasticWorkerCount"`
	// Maximum number of instances that can be assigned to this App Service plan.
	MaximumNumberOfWorkers pulumi.IntOutput `pulumi:"maximumNumberOfWorkers"`
	// Resource Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Number of apps assigned to this App Service plan.
	NumberOfSites pulumi.IntOutput `pulumi:"numberOfSites"`
	// If <code>true</code>, apps assigned to this App Service plan can be scaled independently.
	// If <code>false</code>, apps assigned to this App Service plan will scale to all instances of the plan.
	PerSiteScaling pulumi.BoolPtrOutput `pulumi:"perSiteScaling"`
	// Provisioning state of the App Service Environment.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// If Linux app service plan <code>true</code>, <code>false</code> otherwise.
	Reserved pulumi.BoolPtrOutput `pulumi:"reserved"`
	// Resource group of the App Service plan.
	ResourceGroup pulumi.StringOutput `pulumi:"resourceGroup"`
	// Description of a SKU for a scalable resource.
	Sku SkuDescriptionResponsePtrOutput `pulumi:"sku"`
	// The time when the server farm expires. Valid only if it is a spot server farm.
	SpotExpirationTime pulumi.StringPtrOutput `pulumi:"spotExpirationTime"`
	// App Service plan status.
	Status pulumi.StringOutput `pulumi:"status"`
	// App Service plan subscription.
	Subscription pulumi.StringOutput `pulumi:"subscription"`
	// The system metadata relating to this resource.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Scaling worker count.
	TargetWorkerCount pulumi.IntPtrOutput `pulumi:"targetWorkerCount"`
	// Scaling worker size ID.
	TargetWorkerSizeId pulumi.IntPtrOutput `pulumi:"targetWorkerSizeId"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// Target worker tier assigned to the App Service plan.
	WorkerTierName pulumi.StringPtrOutput `pulumi:"workerTierName"`
}

// NewAppServicePlan registers a new resource with the given unique name, arguments, and options.
func NewAppServicePlan(ctx *pulumi.Context,
	name string, args *AppServicePlanArgs, opts ...pulumi.ResourceOption) (*AppServicePlan, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.HyperV == nil {
		args.HyperV = pulumi.BoolPtr(false)
	}
	if args.IsXenon == nil {
		args.IsXenon = pulumi.BoolPtr(false)
	}
	if args.PerSiteScaling == nil {
		args.PerSiteScaling = pulumi.BoolPtr(false)
	}
	if args.Reserved == nil {
		args.Reserved = pulumi.BoolPtr(false)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:web/latest:AppServicePlan"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20150801:AppServicePlan"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20160901:AppServicePlan"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20180201:AppServicePlan"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20190801:AppServicePlan"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20200601:AppServicePlan"),
		},
	})
	opts = append(opts, aliases)
	var resource AppServicePlan
	err := ctx.RegisterResource("azure-nextgen:web/v20200901:AppServicePlan", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-nextgen:web/v20200901:AppServicePlan", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppServicePlan resources.
type appServicePlanState struct {
	// The time when the server farm free offer expires.
	FreeOfferExpirationTime *string `pulumi:"freeOfferExpirationTime"`
	// Geographical location for the App Service plan.
	GeoRegion *string `pulumi:"geoRegion"`
	// Specification for the App Service Environment to use for the App Service plan.
	HostingEnvironmentProfile *HostingEnvironmentProfileResponse `pulumi:"hostingEnvironmentProfile"`
	// If Hyper-V container app service plan <code>true</code>, <code>false</code> otherwise.
	HyperV *bool `pulumi:"hyperV"`
	// If <code>true</code>, this App Service Plan owns spot instances.
	IsSpot *bool `pulumi:"isSpot"`
	// Obsolete: If Hyper-V container app service plan <code>true</code>, <code>false</code> otherwise.
	IsXenon *bool `pulumi:"isXenon"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Location.
	Location *string `pulumi:"location"`
	// Maximum number of total workers allowed for this ElasticScaleEnabled App Service Plan
	MaximumElasticWorkerCount *int `pulumi:"maximumElasticWorkerCount"`
	// Maximum number of instances that can be assigned to this App Service plan.
	MaximumNumberOfWorkers *int `pulumi:"maximumNumberOfWorkers"`
	// Resource Name.
	Name *string `pulumi:"name"`
	// Number of apps assigned to this App Service plan.
	NumberOfSites *int `pulumi:"numberOfSites"`
	// If <code>true</code>, apps assigned to this App Service plan can be scaled independently.
	// If <code>false</code>, apps assigned to this App Service plan will scale to all instances of the plan.
	PerSiteScaling *bool `pulumi:"perSiteScaling"`
	// Provisioning state of the App Service Environment.
	ProvisioningState *string `pulumi:"provisioningState"`
	// If Linux app service plan <code>true</code>, <code>false</code> otherwise.
	Reserved *bool `pulumi:"reserved"`
	// Resource group of the App Service plan.
	ResourceGroup *string `pulumi:"resourceGroup"`
	// Description of a SKU for a scalable resource.
	Sku *SkuDescriptionResponse `pulumi:"sku"`
	// The time when the server farm expires. Valid only if it is a spot server farm.
	SpotExpirationTime *string `pulumi:"spotExpirationTime"`
	// App Service plan status.
	Status *string `pulumi:"status"`
	// App Service plan subscription.
	Subscription *string `pulumi:"subscription"`
	// The system metadata relating to this resource.
	SystemData *SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Scaling worker count.
	TargetWorkerCount *int `pulumi:"targetWorkerCount"`
	// Scaling worker size ID.
	TargetWorkerSizeId *int `pulumi:"targetWorkerSizeId"`
	// Resource type.
	Type *string `pulumi:"type"`
	// Target worker tier assigned to the App Service plan.
	WorkerTierName *string `pulumi:"workerTierName"`
}

type AppServicePlanState struct {
	// The time when the server farm free offer expires.
	FreeOfferExpirationTime pulumi.StringPtrInput
	// Geographical location for the App Service plan.
	GeoRegion pulumi.StringPtrInput
	// Specification for the App Service Environment to use for the App Service plan.
	HostingEnvironmentProfile HostingEnvironmentProfileResponsePtrInput
	// If Hyper-V container app service plan <code>true</code>, <code>false</code> otherwise.
	HyperV pulumi.BoolPtrInput
	// If <code>true</code>, this App Service Plan owns spot instances.
	IsSpot pulumi.BoolPtrInput
	// Obsolete: If Hyper-V container app service plan <code>true</code>, <code>false</code> otherwise.
	IsXenon pulumi.BoolPtrInput
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Resource Location.
	Location pulumi.StringPtrInput
	// Maximum number of total workers allowed for this ElasticScaleEnabled App Service Plan
	MaximumElasticWorkerCount pulumi.IntPtrInput
	// Maximum number of instances that can be assigned to this App Service plan.
	MaximumNumberOfWorkers pulumi.IntPtrInput
	// Resource Name.
	Name pulumi.StringPtrInput
	// Number of apps assigned to this App Service plan.
	NumberOfSites pulumi.IntPtrInput
	// If <code>true</code>, apps assigned to this App Service plan can be scaled independently.
	// If <code>false</code>, apps assigned to this App Service plan will scale to all instances of the plan.
	PerSiteScaling pulumi.BoolPtrInput
	// Provisioning state of the App Service Environment.
	ProvisioningState pulumi.StringPtrInput
	// If Linux app service plan <code>true</code>, <code>false</code> otherwise.
	Reserved pulumi.BoolPtrInput
	// Resource group of the App Service plan.
	ResourceGroup pulumi.StringPtrInput
	// Description of a SKU for a scalable resource.
	Sku SkuDescriptionResponsePtrInput
	// The time when the server farm expires. Valid only if it is a spot server farm.
	SpotExpirationTime pulumi.StringPtrInput
	// App Service plan status.
	Status pulumi.StringPtrInput
	// App Service plan subscription.
	Subscription pulumi.StringPtrInput
	// The system metadata relating to this resource.
	SystemData SystemDataResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Scaling worker count.
	TargetWorkerCount pulumi.IntPtrInput
	// Scaling worker size ID.
	TargetWorkerSizeId pulumi.IntPtrInput
	// Resource type.
	Type pulumi.StringPtrInput
	// Target worker tier assigned to the App Service plan.
	WorkerTierName pulumi.StringPtrInput
}

func (AppServicePlanState) ElementType() reflect.Type {
	return reflect.TypeOf((*appServicePlanState)(nil)).Elem()
}

type appServicePlanArgs struct {
	// The time when the server farm free offer expires.
	FreeOfferExpirationTime *string `pulumi:"freeOfferExpirationTime"`
	// Specification for the App Service Environment to use for the App Service plan.
	HostingEnvironmentProfile *HostingEnvironmentProfile `pulumi:"hostingEnvironmentProfile"`
	// If Hyper-V container app service plan <code>true</code>, <code>false</code> otherwise.
	HyperV *bool `pulumi:"hyperV"`
	// If <code>true</code>, this App Service Plan owns spot instances.
	IsSpot *bool `pulumi:"isSpot"`
	// Obsolete: If Hyper-V container app service plan <code>true</code>, <code>false</code> otherwise.
	IsXenon *bool `pulumi:"isXenon"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Location.
	Location *string `pulumi:"location"`
	// Maximum number of total workers allowed for this ElasticScaleEnabled App Service Plan
	MaximumElasticWorkerCount *int `pulumi:"maximumElasticWorkerCount"`
	// Name of the App Service plan.
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
	// The time when the server farm free offer expires.
	FreeOfferExpirationTime pulumi.StringPtrInput
	// Specification for the App Service Environment to use for the App Service plan.
	HostingEnvironmentProfile HostingEnvironmentProfilePtrInput
	// If Hyper-V container app service plan <code>true</code>, <code>false</code> otherwise.
	HyperV pulumi.BoolPtrInput
	// If <code>true</code>, this App Service Plan owns spot instances.
	IsSpot pulumi.BoolPtrInput
	// Obsolete: If Hyper-V container app service plan <code>true</code>, <code>false</code> otherwise.
	IsXenon pulumi.BoolPtrInput
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Resource Location.
	Location pulumi.StringPtrInput
	// Maximum number of total workers allowed for this ElasticScaleEnabled App Service Plan
	MaximumElasticWorkerCount pulumi.IntPtrInput
	// Name of the App Service plan.
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

type AppServicePlanInput interface {
	pulumi.Input

	ToAppServicePlanOutput() AppServicePlanOutput
	ToAppServicePlanOutputWithContext(ctx context.Context) AppServicePlanOutput
}

func (*AppServicePlan) ElementType() reflect.Type {
	return reflect.TypeOf((*AppServicePlan)(nil))
}

func (i *AppServicePlan) ToAppServicePlanOutput() AppServicePlanOutput {
	return i.ToAppServicePlanOutputWithContext(context.Background())
}

func (i *AppServicePlan) ToAppServicePlanOutputWithContext(ctx context.Context) AppServicePlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppServicePlanOutput)
}

type AppServicePlanOutput struct {
	*pulumi.OutputState
}

func (AppServicePlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppServicePlan)(nil))
}

func (o AppServicePlanOutput) ToAppServicePlanOutput() AppServicePlanOutput {
	return o
}

func (o AppServicePlanOutput) ToAppServicePlanOutputWithContext(ctx context.Context) AppServicePlanOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AppServicePlanOutput{})
}
