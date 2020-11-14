// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Description of an hostingEnvironment (App Service Environment)
type ManagedHostingEnvironment struct {
	pulumi.CustomResourceState

	// List of comma separated strings describing which VM sizes are allowed for front-ends
	AllowedMultiSizes pulumi.StringPtrOutput `pulumi:"allowedMultiSizes"`
	// List of comma separated strings describing which VM sizes are allowed for workers
	AllowedWorkerSizes pulumi.StringPtrOutput `pulumi:"allowedWorkerSizes"`
	// Api Management Account associated with this Hosting Environment
	ApiManagementAccountId pulumi.StringPtrOutput `pulumi:"apiManagementAccountId"`
	// Custom settings for changing the behavior of the hosting environment
	ClusterSettings NameValuePairResponseArrayOutput `pulumi:"clusterSettings"`
	// Edition of the metadata database for the hostingEnvironment (App Service Environment) e.g. "Standard"
	DatabaseEdition pulumi.StringPtrOutput `pulumi:"databaseEdition"`
	// Service objective of the metadata database for the hostingEnvironment (App Service Environment) e.g. "S0"
	DatabaseServiceObjective pulumi.StringPtrOutput `pulumi:"databaseServiceObjective"`
	// DNS suffix of the hostingEnvironment (App Service Environment)
	DnsSuffix pulumi.StringPtrOutput `pulumi:"dnsSuffix"`
	// Current total, used, and available worker capacities
	EnvironmentCapacities StampCapacityResponseArrayOutput `pulumi:"environmentCapacities"`
	// True/false indicating whether the hostingEnvironment (App Service Environment) is healthy
	EnvironmentIsHealthy pulumi.BoolPtrOutput `pulumi:"environmentIsHealthy"`
	// Detailed message about with results of the last check of the hostingEnvironment (App Service Environment)
	EnvironmentStatus pulumi.StringPtrOutput `pulumi:"environmentStatus"`
	// Specifies which endpoints to serve internally in the hostingEnvironment's (App Service Environment) VNET
	InternalLoadBalancingMode pulumi.StringPtrOutput `pulumi:"internalLoadBalancingMode"`
	// Number of IP SSL addresses reserved for this hostingEnvironment (App Service Environment)
	IpsslAddressCount pulumi.IntPtrOutput `pulumi:"ipsslAddressCount"`
	// Kind of resource
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Last deployment action on this hostingEnvironment (App Service Environment)
	LastAction pulumi.StringPtrOutput `pulumi:"lastAction"`
	// Result of the last deployment action on this hostingEnvironment (App Service Environment)
	LastActionResult pulumi.StringPtrOutput `pulumi:"lastActionResult"`
	// Resource Location
	Location pulumi.StringOutput `pulumi:"location"`
	// Maximum number of VMs in this hostingEnvironment (App Service Environment)
	MaximumNumberOfMachines pulumi.IntPtrOutput `pulumi:"maximumNumberOfMachines"`
	// Number of front-end instances
	MultiRoleCount pulumi.IntPtrOutput `pulumi:"multiRoleCount"`
	// Front-end VM size, e.g. "Medium", "Large"
	MultiSize pulumi.StringPtrOutput `pulumi:"multiSize"`
	// Resource Name
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// Access control list for controlling traffic to the hostingEnvironment (App Service Environment)
	NetworkAccessControlList NetworkAccessControlEntryResponseArrayOutput `pulumi:"networkAccessControlList"`
	// Provisioning state of the hostingEnvironment (App Service Environment)
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// Resource group of the hostingEnvironment (App Service Environment)
	ResourceGroup pulumi.StringPtrOutput `pulumi:"resourceGroup"`
	// Current status of the hostingEnvironment (App Service Environment)
	Status pulumi.StringOutput `pulumi:"status"`
	// Subscription of the hostingEnvironment (App Service Environment)
	SubscriptionId pulumi.StringPtrOutput `pulumi:"subscriptionId"`
	// True/false indicating whether the hostingEnvironment is suspended. The environment can be suspended e.g. when the management endpoint is no longer available
	//             (most likely because NSG blocked the incoming traffic)
	Suspended pulumi.BoolPtrOutput `pulumi:"suspended"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringPtrOutput `pulumi:"type"`
	// Number of upgrade domains of this hostingEnvironment (App Service Environment)
	UpgradeDomains pulumi.IntPtrOutput `pulumi:"upgradeDomains"`
	// Description of IP SSL mapping for this hostingEnvironment (App Service Environment)
	VipMappings VirtualIPMappingResponseArrayOutput `pulumi:"vipMappings"`
	// Description of the hostingEnvironment's (App Service Environment) virtual network
	VirtualNetwork VirtualNetworkProfileResponsePtrOutput `pulumi:"virtualNetwork"`
	// Name of the hostingEnvironment's (App Service Environment) virtual network
	VnetName pulumi.StringPtrOutput `pulumi:"vnetName"`
	// Resource group of the hostingEnvironment's (App Service Environment) virtual network
	VnetResourceGroupName pulumi.StringPtrOutput `pulumi:"vnetResourceGroupName"`
	// Subnet of the hostingEnvironment's (App Service Environment) virtual network
	VnetSubnetName pulumi.StringPtrOutput `pulumi:"vnetSubnetName"`
	// Description of worker pools with worker size ids, VM sizes, and number of workers in each pool
	WorkerPools WorkerPoolResponseArrayOutput `pulumi:"workerPools"`
}

// NewManagedHostingEnvironment registers a new resource with the given unique name, arguments, and options.
func NewManagedHostingEnvironment(ctx *pulumi.Context,
	name string, args *ManagedHostingEnvironmentArgs, opts ...pulumi.ResourceOption) (*ManagedHostingEnvironment, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Status == nil {
		return nil, errors.New("missing required argument 'Status'")
	}
	if args == nil {
		args = &ManagedHostingEnvironmentArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:web/v20150801:ManagedHostingEnvironment"),
		},
	})
	opts = append(opts, aliases)
	var resource ManagedHostingEnvironment
	err := ctx.RegisterResource("azure-nextgen:web/latest:ManagedHostingEnvironment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagedHostingEnvironment gets an existing ManagedHostingEnvironment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagedHostingEnvironment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedHostingEnvironmentState, opts ...pulumi.ResourceOption) (*ManagedHostingEnvironment, error) {
	var resource ManagedHostingEnvironment
	err := ctx.ReadResource("azure-nextgen:web/latest:ManagedHostingEnvironment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagedHostingEnvironment resources.
type managedHostingEnvironmentState struct {
	// List of comma separated strings describing which VM sizes are allowed for front-ends
	AllowedMultiSizes *string `pulumi:"allowedMultiSizes"`
	// List of comma separated strings describing which VM sizes are allowed for workers
	AllowedWorkerSizes *string `pulumi:"allowedWorkerSizes"`
	// Api Management Account associated with this Hosting Environment
	ApiManagementAccountId *string `pulumi:"apiManagementAccountId"`
	// Custom settings for changing the behavior of the hosting environment
	ClusterSettings []NameValuePairResponse `pulumi:"clusterSettings"`
	// Edition of the metadata database for the hostingEnvironment (App Service Environment) e.g. "Standard"
	DatabaseEdition *string `pulumi:"databaseEdition"`
	// Service objective of the metadata database for the hostingEnvironment (App Service Environment) e.g. "S0"
	DatabaseServiceObjective *string `pulumi:"databaseServiceObjective"`
	// DNS suffix of the hostingEnvironment (App Service Environment)
	DnsSuffix *string `pulumi:"dnsSuffix"`
	// Current total, used, and available worker capacities
	EnvironmentCapacities []StampCapacityResponse `pulumi:"environmentCapacities"`
	// True/false indicating whether the hostingEnvironment (App Service Environment) is healthy
	EnvironmentIsHealthy *bool `pulumi:"environmentIsHealthy"`
	// Detailed message about with results of the last check of the hostingEnvironment (App Service Environment)
	EnvironmentStatus *string `pulumi:"environmentStatus"`
	// Specifies which endpoints to serve internally in the hostingEnvironment's (App Service Environment) VNET
	InternalLoadBalancingMode *string `pulumi:"internalLoadBalancingMode"`
	// Number of IP SSL addresses reserved for this hostingEnvironment (App Service Environment)
	IpsslAddressCount *int `pulumi:"ipsslAddressCount"`
	// Kind of resource
	Kind *string `pulumi:"kind"`
	// Last deployment action on this hostingEnvironment (App Service Environment)
	LastAction *string `pulumi:"lastAction"`
	// Result of the last deployment action on this hostingEnvironment (App Service Environment)
	LastActionResult *string `pulumi:"lastActionResult"`
	// Resource Location
	Location *string `pulumi:"location"`
	// Maximum number of VMs in this hostingEnvironment (App Service Environment)
	MaximumNumberOfMachines *int `pulumi:"maximumNumberOfMachines"`
	// Number of front-end instances
	MultiRoleCount *int `pulumi:"multiRoleCount"`
	// Front-end VM size, e.g. "Medium", "Large"
	MultiSize *string `pulumi:"multiSize"`
	// Resource Name
	Name *string `pulumi:"name"`
	// Access control list for controlling traffic to the hostingEnvironment (App Service Environment)
	NetworkAccessControlList []NetworkAccessControlEntryResponse `pulumi:"networkAccessControlList"`
	// Provisioning state of the hostingEnvironment (App Service Environment)
	ProvisioningState *string `pulumi:"provisioningState"`
	// Resource group of the hostingEnvironment (App Service Environment)
	ResourceGroup *string `pulumi:"resourceGroup"`
	// Current status of the hostingEnvironment (App Service Environment)
	Status *string `pulumi:"status"`
	// Subscription of the hostingEnvironment (App Service Environment)
	SubscriptionId *string `pulumi:"subscriptionId"`
	// True/false indicating whether the hostingEnvironment is suspended. The environment can be suspended e.g. when the management endpoint is no longer available
	//             (most likely because NSG blocked the incoming traffic)
	Suspended *bool `pulumi:"suspended"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
	// Number of upgrade domains of this hostingEnvironment (App Service Environment)
	UpgradeDomains *int `pulumi:"upgradeDomains"`
	// Description of IP SSL mapping for this hostingEnvironment (App Service Environment)
	VipMappings []VirtualIPMappingResponse `pulumi:"vipMappings"`
	// Description of the hostingEnvironment's (App Service Environment) virtual network
	VirtualNetwork *VirtualNetworkProfileResponse `pulumi:"virtualNetwork"`
	// Name of the hostingEnvironment's (App Service Environment) virtual network
	VnetName *string `pulumi:"vnetName"`
	// Resource group of the hostingEnvironment's (App Service Environment) virtual network
	VnetResourceGroupName *string `pulumi:"vnetResourceGroupName"`
	// Subnet of the hostingEnvironment's (App Service Environment) virtual network
	VnetSubnetName *string `pulumi:"vnetSubnetName"`
	// Description of worker pools with worker size ids, VM sizes, and number of workers in each pool
	WorkerPools []WorkerPoolResponse `pulumi:"workerPools"`
}

type ManagedHostingEnvironmentState struct {
	// List of comma separated strings describing which VM sizes are allowed for front-ends
	AllowedMultiSizes pulumi.StringPtrInput
	// List of comma separated strings describing which VM sizes are allowed for workers
	AllowedWorkerSizes pulumi.StringPtrInput
	// Api Management Account associated with this Hosting Environment
	ApiManagementAccountId pulumi.StringPtrInput
	// Custom settings for changing the behavior of the hosting environment
	ClusterSettings NameValuePairResponseArrayInput
	// Edition of the metadata database for the hostingEnvironment (App Service Environment) e.g. "Standard"
	DatabaseEdition pulumi.StringPtrInput
	// Service objective of the metadata database for the hostingEnvironment (App Service Environment) e.g. "S0"
	DatabaseServiceObjective pulumi.StringPtrInput
	// DNS suffix of the hostingEnvironment (App Service Environment)
	DnsSuffix pulumi.StringPtrInput
	// Current total, used, and available worker capacities
	EnvironmentCapacities StampCapacityResponseArrayInput
	// True/false indicating whether the hostingEnvironment (App Service Environment) is healthy
	EnvironmentIsHealthy pulumi.BoolPtrInput
	// Detailed message about with results of the last check of the hostingEnvironment (App Service Environment)
	EnvironmentStatus pulumi.StringPtrInput
	// Specifies which endpoints to serve internally in the hostingEnvironment's (App Service Environment) VNET
	InternalLoadBalancingMode pulumi.StringPtrInput
	// Number of IP SSL addresses reserved for this hostingEnvironment (App Service Environment)
	IpsslAddressCount pulumi.IntPtrInput
	// Kind of resource
	Kind pulumi.StringPtrInput
	// Last deployment action on this hostingEnvironment (App Service Environment)
	LastAction pulumi.StringPtrInput
	// Result of the last deployment action on this hostingEnvironment (App Service Environment)
	LastActionResult pulumi.StringPtrInput
	// Resource Location
	Location pulumi.StringPtrInput
	// Maximum number of VMs in this hostingEnvironment (App Service Environment)
	MaximumNumberOfMachines pulumi.IntPtrInput
	// Number of front-end instances
	MultiRoleCount pulumi.IntPtrInput
	// Front-end VM size, e.g. "Medium", "Large"
	MultiSize pulumi.StringPtrInput
	// Resource Name
	Name pulumi.StringPtrInput
	// Access control list for controlling traffic to the hostingEnvironment (App Service Environment)
	NetworkAccessControlList NetworkAccessControlEntryResponseArrayInput
	// Provisioning state of the hostingEnvironment (App Service Environment)
	ProvisioningState pulumi.StringPtrInput
	// Resource group of the hostingEnvironment (App Service Environment)
	ResourceGroup pulumi.StringPtrInput
	// Current status of the hostingEnvironment (App Service Environment)
	Status pulumi.StringPtrInput
	// Subscription of the hostingEnvironment (App Service Environment)
	SubscriptionId pulumi.StringPtrInput
	// True/false indicating whether the hostingEnvironment is suspended. The environment can be suspended e.g. when the management endpoint is no longer available
	//             (most likely because NSG blocked the incoming traffic)
	Suspended pulumi.BoolPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Resource type
	Type pulumi.StringPtrInput
	// Number of upgrade domains of this hostingEnvironment (App Service Environment)
	UpgradeDomains pulumi.IntPtrInput
	// Description of IP SSL mapping for this hostingEnvironment (App Service Environment)
	VipMappings VirtualIPMappingResponseArrayInput
	// Description of the hostingEnvironment's (App Service Environment) virtual network
	VirtualNetwork VirtualNetworkProfileResponsePtrInput
	// Name of the hostingEnvironment's (App Service Environment) virtual network
	VnetName pulumi.StringPtrInput
	// Resource group of the hostingEnvironment's (App Service Environment) virtual network
	VnetResourceGroupName pulumi.StringPtrInput
	// Subnet of the hostingEnvironment's (App Service Environment) virtual network
	VnetSubnetName pulumi.StringPtrInput
	// Description of worker pools with worker size ids, VM sizes, and number of workers in each pool
	WorkerPools WorkerPoolResponseArrayInput
}

func (ManagedHostingEnvironmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedHostingEnvironmentState)(nil)).Elem()
}

type managedHostingEnvironmentArgs struct {
	// List of comma separated strings describing which VM sizes are allowed for front-ends
	AllowedMultiSizes *string `pulumi:"allowedMultiSizes"`
	// List of comma separated strings describing which VM sizes are allowed for workers
	AllowedWorkerSizes *string `pulumi:"allowedWorkerSizes"`
	// Api Management Account associated with this Hosting Environment
	ApiManagementAccountId *string `pulumi:"apiManagementAccountId"`
	// Custom settings for changing the behavior of the hosting environment
	ClusterSettings []NameValuePair `pulumi:"clusterSettings"`
	// Edition of the metadata database for the hostingEnvironment (App Service Environment) e.g. "Standard"
	DatabaseEdition *string `pulumi:"databaseEdition"`
	// Service objective of the metadata database for the hostingEnvironment (App Service Environment) e.g. "S0"
	DatabaseServiceObjective *string `pulumi:"databaseServiceObjective"`
	// DNS suffix of the hostingEnvironment (App Service Environment)
	DnsSuffix *string `pulumi:"dnsSuffix"`
	// Current total, used, and available worker capacities
	EnvironmentCapacities []StampCapacity `pulumi:"environmentCapacities"`
	// True/false indicating whether the hostingEnvironment (App Service Environment) is healthy
	EnvironmentIsHealthy *bool `pulumi:"environmentIsHealthy"`
	// Detailed message about with results of the last check of the hostingEnvironment (App Service Environment)
	EnvironmentStatus *string `pulumi:"environmentStatus"`
	// Resource Id
	Id *string `pulumi:"id"`
	// Specifies which endpoints to serve internally in the hostingEnvironment's (App Service Environment) VNET
	InternalLoadBalancingMode *string `pulumi:"internalLoadBalancingMode"`
	// Number of IP SSL addresses reserved for this hostingEnvironment (App Service Environment)
	IpsslAddressCount *int `pulumi:"ipsslAddressCount"`
	// Kind of resource
	Kind *string `pulumi:"kind"`
	// Last deployment action on this hostingEnvironment (App Service Environment)
	LastAction *string `pulumi:"lastAction"`
	// Result of the last deployment action on this hostingEnvironment (App Service Environment)
	LastActionResult *string `pulumi:"lastActionResult"`
	// Resource Location
	Location string `pulumi:"location"`
	// Maximum number of VMs in this hostingEnvironment (App Service Environment)
	MaximumNumberOfMachines *int `pulumi:"maximumNumberOfMachines"`
	// Number of front-end instances
	MultiRoleCount *int `pulumi:"multiRoleCount"`
	// Front-end VM size, e.g. "Medium", "Large"
	MultiSize *string `pulumi:"multiSize"`
	// Resource Name
	Name string `pulumi:"name"`
	// Access control list for controlling traffic to the hostingEnvironment (App Service Environment)
	NetworkAccessControlList []NetworkAccessControlEntry `pulumi:"networkAccessControlList"`
	// Provisioning state of the hostingEnvironment (App Service Environment)
	ProvisioningState *string `pulumi:"provisioningState"`
	// Resource group of the hostingEnvironment (App Service Environment)
	ResourceGroup *string `pulumi:"resourceGroup"`
	// Name of resource group
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Current status of the hostingEnvironment (App Service Environment)
	Status string `pulumi:"status"`
	// Subscription of the hostingEnvironment (App Service Environment)
	SubscriptionId *string `pulumi:"subscriptionId"`
	// True/false indicating whether the hostingEnvironment is suspended. The environment can be suspended e.g. when the management endpoint is no longer available
	//             (most likely because NSG blocked the incoming traffic)
	Suspended *bool `pulumi:"suspended"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
	// Number of upgrade domains of this hostingEnvironment (App Service Environment)
	UpgradeDomains *int `pulumi:"upgradeDomains"`
	// Description of IP SSL mapping for this hostingEnvironment (App Service Environment)
	VipMappings []VirtualIPMapping `pulumi:"vipMappings"`
	// Description of the hostingEnvironment's (App Service Environment) virtual network
	VirtualNetwork *VirtualNetworkProfile `pulumi:"virtualNetwork"`
	// Name of the hostingEnvironment's (App Service Environment) virtual network
	VnetName *string `pulumi:"vnetName"`
	// Resource group of the hostingEnvironment's (App Service Environment) virtual network
	VnetResourceGroupName *string `pulumi:"vnetResourceGroupName"`
	// Subnet of the hostingEnvironment's (App Service Environment) virtual network
	VnetSubnetName *string `pulumi:"vnetSubnetName"`
	// Description of worker pools with worker size ids, VM sizes, and number of workers in each pool
	WorkerPools []WorkerPool `pulumi:"workerPools"`
}

// The set of arguments for constructing a ManagedHostingEnvironment resource.
type ManagedHostingEnvironmentArgs struct {
	// List of comma separated strings describing which VM sizes are allowed for front-ends
	AllowedMultiSizes pulumi.StringPtrInput
	// List of comma separated strings describing which VM sizes are allowed for workers
	AllowedWorkerSizes pulumi.StringPtrInput
	// Api Management Account associated with this Hosting Environment
	ApiManagementAccountId pulumi.StringPtrInput
	// Custom settings for changing the behavior of the hosting environment
	ClusterSettings NameValuePairArrayInput
	// Edition of the metadata database for the hostingEnvironment (App Service Environment) e.g. "Standard"
	DatabaseEdition pulumi.StringPtrInput
	// Service objective of the metadata database for the hostingEnvironment (App Service Environment) e.g. "S0"
	DatabaseServiceObjective pulumi.StringPtrInput
	// DNS suffix of the hostingEnvironment (App Service Environment)
	DnsSuffix pulumi.StringPtrInput
	// Current total, used, and available worker capacities
	EnvironmentCapacities StampCapacityArrayInput
	// True/false indicating whether the hostingEnvironment (App Service Environment) is healthy
	EnvironmentIsHealthy pulumi.BoolPtrInput
	// Detailed message about with results of the last check of the hostingEnvironment (App Service Environment)
	EnvironmentStatus pulumi.StringPtrInput
	// Resource Id
	Id pulumi.StringPtrInput
	// Specifies which endpoints to serve internally in the hostingEnvironment's (App Service Environment) VNET
	InternalLoadBalancingMode pulumi.StringPtrInput
	// Number of IP SSL addresses reserved for this hostingEnvironment (App Service Environment)
	IpsslAddressCount pulumi.IntPtrInput
	// Kind of resource
	Kind pulumi.StringPtrInput
	// Last deployment action on this hostingEnvironment (App Service Environment)
	LastAction pulumi.StringPtrInput
	// Result of the last deployment action on this hostingEnvironment (App Service Environment)
	LastActionResult pulumi.StringPtrInput
	// Resource Location
	Location pulumi.StringInput
	// Maximum number of VMs in this hostingEnvironment (App Service Environment)
	MaximumNumberOfMachines pulumi.IntPtrInput
	// Number of front-end instances
	MultiRoleCount pulumi.IntPtrInput
	// Front-end VM size, e.g. "Medium", "Large"
	MultiSize pulumi.StringPtrInput
	// Resource Name
	Name pulumi.StringInput
	// Access control list for controlling traffic to the hostingEnvironment (App Service Environment)
	NetworkAccessControlList NetworkAccessControlEntryArrayInput
	// Provisioning state of the hostingEnvironment (App Service Environment)
	ProvisioningState pulumi.StringPtrInput
	// Resource group of the hostingEnvironment (App Service Environment)
	ResourceGroup pulumi.StringPtrInput
	// Name of resource group
	ResourceGroupName pulumi.StringInput
	// Current status of the hostingEnvironment (App Service Environment)
	Status pulumi.StringInput
	// Subscription of the hostingEnvironment (App Service Environment)
	SubscriptionId pulumi.StringPtrInput
	// True/false indicating whether the hostingEnvironment is suspended. The environment can be suspended e.g. when the management endpoint is no longer available
	//             (most likely because NSG blocked the incoming traffic)
	Suspended pulumi.BoolPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Resource type
	Type pulumi.StringPtrInput
	// Number of upgrade domains of this hostingEnvironment (App Service Environment)
	UpgradeDomains pulumi.IntPtrInput
	// Description of IP SSL mapping for this hostingEnvironment (App Service Environment)
	VipMappings VirtualIPMappingArrayInput
	// Description of the hostingEnvironment's (App Service Environment) virtual network
	VirtualNetwork VirtualNetworkProfilePtrInput
	// Name of the hostingEnvironment's (App Service Environment) virtual network
	VnetName pulumi.StringPtrInput
	// Resource group of the hostingEnvironment's (App Service Environment) virtual network
	VnetResourceGroupName pulumi.StringPtrInput
	// Subnet of the hostingEnvironment's (App Service Environment) virtual network
	VnetSubnetName pulumi.StringPtrInput
	// Description of worker pools with worker size ids, VM sizes, and number of workers in each pool
	WorkerPools WorkerPoolArrayInput
}

func (ManagedHostingEnvironmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedHostingEnvironmentArgs)(nil)).Elem()
}

type ManagedHostingEnvironmentInput interface {
	pulumi.Input

	ToManagedHostingEnvironmentOutput() ManagedHostingEnvironmentOutput
	ToManagedHostingEnvironmentOutputWithContext(ctx context.Context) ManagedHostingEnvironmentOutput
}

func (ManagedHostingEnvironment) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedHostingEnvironment)(nil)).Elem()
}

func (i ManagedHostingEnvironment) ToManagedHostingEnvironmentOutput() ManagedHostingEnvironmentOutput {
	return i.ToManagedHostingEnvironmentOutputWithContext(context.Background())
}

func (i ManagedHostingEnvironment) ToManagedHostingEnvironmentOutputWithContext(ctx context.Context) ManagedHostingEnvironmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedHostingEnvironmentOutput)
}

type ManagedHostingEnvironmentOutput struct {
	*pulumi.OutputState
}

func (ManagedHostingEnvironmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedHostingEnvironmentOutput)(nil)).Elem()
}

func (o ManagedHostingEnvironmentOutput) ToManagedHostingEnvironmentOutput() ManagedHostingEnvironmentOutput {
	return o
}

func (o ManagedHostingEnvironmentOutput) ToManagedHostingEnvironmentOutputWithContext(ctx context.Context) ManagedHostingEnvironmentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ManagedHostingEnvironmentOutput{})
}
