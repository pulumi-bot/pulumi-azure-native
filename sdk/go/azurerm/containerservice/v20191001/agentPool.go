// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191001

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Agent Pool.
//
// ## Example Usage
// ### Create/Update Agent Pool
//
// ```go
// package main
//
// import (
// 	containerservice "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/containerservice/v20191001"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := containerservice.NewAgentPool(ctx, "agentPool", &containerservice.AgentPoolArgs{
// 			AgentPoolName: pulumi.String("agentpool1"),
// 			Count:         pulumi.Int(3),
// 			NodeTaints: pulumi.StringArray{
// 				pulumi.String("Key1=Value1:NoSchedule"),
// 			},
// 			OrchestratorVersion:    pulumi.String(""),
// 			OsType:                 pulumi.String("Linux"),
// 			ResourceGroupName:      pulumi.String("rg1"),
// 			ResourceName:           pulumi.String("clustername1"),
// 			ScaleSetEvictionPolicy: pulumi.String("Delete"),
// 			ScaleSetPriority:       pulumi.String("Low"),
// 			VmSize:                 pulumi.String("Standard_DS1_v2"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
// ### Update Agent Pool
//
// ```go
// package main
//
// import (
// 	containerservice "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/containerservice/v20191001"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := containerservice.NewAgentPool(ctx, "agentPool", &containerservice.AgentPoolArgs{
// 			AgentPoolName:     pulumi.String("agentpool1"),
// 			Count:             pulumi.Int(3),
// 			EnableAutoScaling: pulumi.Bool(true),
// 			MaxCount:          pulumi.Int(2),
// 			MinCount:          pulumi.Int(2),
// 			NodeTaints: pulumi.StringArray{
// 				pulumi.String("Key1=Value1:NoSchedule"),
// 			},
// 			OrchestratorVersion:    pulumi.String(""),
// 			OsType:                 pulumi.String("Linux"),
// 			ResourceGroupName:      pulumi.String("rg1"),
// 			ResourceName:           pulumi.String("clustername1"),
// 			ScaleSetEvictionPolicy: pulumi.String("Delete"),
// 			ScaleSetPriority:       pulumi.String("Low"),
// 			VmSize:                 pulumi.String("Standard_DS1_v2"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type AgentPool struct {
	pulumi.CustomResourceState

	// (PREVIEW) Availability zones for nodes. Must use VirtualMachineScaleSets AgentPoolType.
	AvailabilityZones pulumi.StringArrayOutput `pulumi:"availabilityZones"`
	// Number of agents (VMs) to host docker containers. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 1.
	Count pulumi.IntPtrOutput `pulumi:"count"`
	// Whether to enable auto-scaler
	EnableAutoScaling pulumi.BoolPtrOutput `pulumi:"enableAutoScaling"`
	// Enable public IP for nodes
	EnableNodePublicIP pulumi.BoolPtrOutput `pulumi:"enableNodePublicIP"`
	// Maximum number of nodes for auto-scaling
	MaxCount pulumi.IntPtrOutput `pulumi:"maxCount"`
	// Maximum number of pods that can run on a node.
	MaxPods pulumi.IntPtrOutput `pulumi:"maxPods"`
	// Minimum number of nodes for auto-scaling
	MinCount pulumi.IntPtrOutput `pulumi:"minCount"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.
	NodeTaints pulumi.StringArrayOutput `pulumi:"nodeTaints"`
	// Version of orchestrator specified when creating the managed cluster.
	OrchestratorVersion pulumi.StringPtrOutput `pulumi:"orchestratorVersion"`
	// OS Disk Size in GB to be used to specify the disk size for every machine in this master/agent pool. If you specify 0, it will apply the default osDisk size according to the vmSize specified.
	OsDiskSizeGB pulumi.IntPtrOutput `pulumi:"osDiskSizeGB"`
	// OsType to be used to specify os type. Choose from Linux and Windows. Default to Linux.
	OsType pulumi.StringPtrOutput `pulumi:"osType"`
	// The current deployment or provisioning state, which only appears in the response.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// ScaleSetEvictionPolicy to be used to specify eviction policy for low priority virtual machine scale set. Default to Delete.
	ScaleSetEvictionPolicy pulumi.StringPtrOutput `pulumi:"scaleSetEvictionPolicy"`
	// ScaleSetPriority to be used to specify virtual machine scale set priority. Default to regular.
	ScaleSetPriority pulumi.StringPtrOutput `pulumi:"scaleSetPriority"`
	// AgentPoolType represents types of an agent pool
	Type pulumi.StringOutput `pulumi:"type"`
	// Size of agent VMs.
	VmSize pulumi.StringPtrOutput `pulumi:"vmSize"`
	// VNet SubnetID specifies the VNet's subnet identifier.
	VnetSubnetID pulumi.StringPtrOutput `pulumi:"vnetSubnetID"`
}

// NewAgentPool registers a new resource with the given unique name, arguments, and options.
func NewAgentPool(ctx *pulumi.Context,
	name string, args *AgentPoolArgs, opts ...pulumi.ResourceOption) (*AgentPool, error) {
	if args == nil || args.AgentPoolName == nil {
		return nil, errors.New("missing required argument 'AgentPoolName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ResourceName == nil {
		return nil, errors.New("missing required argument 'ResourceName'")
	}
	if args == nil {
		args = &AgentPoolArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:containerservice/latest:AgentPool"),
		},
		{
			Type: pulumi.String("azurerm:containerservice/v20190201:AgentPool"),
		},
		{
			Type: pulumi.String("azurerm:containerservice/v20190401:AgentPool"),
		},
		{
			Type: pulumi.String("azurerm:containerservice/v20190601:AgentPool"),
		},
		{
			Type: pulumi.String("azurerm:containerservice/v20190801:AgentPool"),
		},
		{
			Type: pulumi.String("azurerm:containerservice/v20191101:AgentPool"),
		},
		{
			Type: pulumi.String("azurerm:containerservice/v20200101:AgentPool"),
		},
		{
			Type: pulumi.String("azurerm:containerservice/v20200201:AgentPool"),
		},
		{
			Type: pulumi.String("azurerm:containerservice/v20200301:AgentPool"),
		},
		{
			Type: pulumi.String("azurerm:containerservice/v20200401:AgentPool"),
		},
		{
			Type: pulumi.String("azurerm:containerservice/v20200601:AgentPool"),
		},
		{
			Type: pulumi.String("azurerm:containerservice/v20200701:AgentPool"),
		},
		{
			Type: pulumi.String("azurerm:containerservice/v20200901:AgentPool"),
		},
	})
	opts = append(opts, aliases)
	var resource AgentPool
	err := ctx.RegisterResource("azurerm:containerservice/v20191001:AgentPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAgentPool gets an existing AgentPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAgentPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AgentPoolState, opts ...pulumi.ResourceOption) (*AgentPool, error) {
	var resource AgentPool
	err := ctx.ReadResource("azurerm:containerservice/v20191001:AgentPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AgentPool resources.
type agentPoolState struct {
	// (PREVIEW) Availability zones for nodes. Must use VirtualMachineScaleSets AgentPoolType.
	AvailabilityZones []string `pulumi:"availabilityZones"`
	// Number of agents (VMs) to host docker containers. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 1.
	Count *int `pulumi:"count"`
	// Whether to enable auto-scaler
	EnableAutoScaling *bool `pulumi:"enableAutoScaling"`
	// Enable public IP for nodes
	EnableNodePublicIP *bool `pulumi:"enableNodePublicIP"`
	// Maximum number of nodes for auto-scaling
	MaxCount *int `pulumi:"maxCount"`
	// Maximum number of pods that can run on a node.
	MaxPods *int `pulumi:"maxPods"`
	// Minimum number of nodes for auto-scaling
	MinCount *int `pulumi:"minCount"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// Taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.
	NodeTaints []string `pulumi:"nodeTaints"`
	// Version of orchestrator specified when creating the managed cluster.
	OrchestratorVersion *string `pulumi:"orchestratorVersion"`
	// OS Disk Size in GB to be used to specify the disk size for every machine in this master/agent pool. If you specify 0, it will apply the default osDisk size according to the vmSize specified.
	OsDiskSizeGB *int `pulumi:"osDiskSizeGB"`
	// OsType to be used to specify os type. Choose from Linux and Windows. Default to Linux.
	OsType *string `pulumi:"osType"`
	// The current deployment or provisioning state, which only appears in the response.
	ProvisioningState *string `pulumi:"provisioningState"`
	// ScaleSetEvictionPolicy to be used to specify eviction policy for low priority virtual machine scale set. Default to Delete.
	ScaleSetEvictionPolicy *string `pulumi:"scaleSetEvictionPolicy"`
	// ScaleSetPriority to be used to specify virtual machine scale set priority. Default to regular.
	ScaleSetPriority *string `pulumi:"scaleSetPriority"`
	// AgentPoolType represents types of an agent pool
	Type *string `pulumi:"type"`
	// Size of agent VMs.
	VmSize *string `pulumi:"vmSize"`
	// VNet SubnetID specifies the VNet's subnet identifier.
	VnetSubnetID *string `pulumi:"vnetSubnetID"`
}

type AgentPoolState struct {
	// (PREVIEW) Availability zones for nodes. Must use VirtualMachineScaleSets AgentPoolType.
	AvailabilityZones pulumi.StringArrayInput
	// Number of agents (VMs) to host docker containers. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 1.
	Count pulumi.IntPtrInput
	// Whether to enable auto-scaler
	EnableAutoScaling pulumi.BoolPtrInput
	// Enable public IP for nodes
	EnableNodePublicIP pulumi.BoolPtrInput
	// Maximum number of nodes for auto-scaling
	MaxCount pulumi.IntPtrInput
	// Maximum number of pods that can run on a node.
	MaxPods pulumi.IntPtrInput
	// Minimum number of nodes for auto-scaling
	MinCount pulumi.IntPtrInput
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrInput
	// Taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.
	NodeTaints pulumi.StringArrayInput
	// Version of orchestrator specified when creating the managed cluster.
	OrchestratorVersion pulumi.StringPtrInput
	// OS Disk Size in GB to be used to specify the disk size for every machine in this master/agent pool. If you specify 0, it will apply the default osDisk size according to the vmSize specified.
	OsDiskSizeGB pulumi.IntPtrInput
	// OsType to be used to specify os type. Choose from Linux and Windows. Default to Linux.
	OsType pulumi.StringPtrInput
	// The current deployment or provisioning state, which only appears in the response.
	ProvisioningState pulumi.StringPtrInput
	// ScaleSetEvictionPolicy to be used to specify eviction policy for low priority virtual machine scale set. Default to Delete.
	ScaleSetEvictionPolicy pulumi.StringPtrInput
	// ScaleSetPriority to be used to specify virtual machine scale set priority. Default to regular.
	ScaleSetPriority pulumi.StringPtrInput
	// AgentPoolType represents types of an agent pool
	Type pulumi.StringPtrInput
	// Size of agent VMs.
	VmSize pulumi.StringPtrInput
	// VNet SubnetID specifies the VNet's subnet identifier.
	VnetSubnetID pulumi.StringPtrInput
}

func (AgentPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*agentPoolState)(nil)).Elem()
}

type agentPoolArgs struct {
	// The name of the agent pool.
	AgentPoolName string `pulumi:"agentPoolName"`
	// (PREVIEW) Availability zones for nodes. Must use VirtualMachineScaleSets AgentPoolType.
	AvailabilityZones []string `pulumi:"availabilityZones"`
	// Number of agents (VMs) to host docker containers. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 1.
	Count *int `pulumi:"count"`
	// Whether to enable auto-scaler
	EnableAutoScaling *bool `pulumi:"enableAutoScaling"`
	// Enable public IP for nodes
	EnableNodePublicIP *bool `pulumi:"enableNodePublicIP"`
	// Maximum number of nodes for auto-scaling
	MaxCount *int `pulumi:"maxCount"`
	// Maximum number of pods that can run on a node.
	MaxPods *int `pulumi:"maxPods"`
	// Minimum number of nodes for auto-scaling
	MinCount *int `pulumi:"minCount"`
	// Taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.
	NodeTaints []string `pulumi:"nodeTaints"`
	// Version of orchestrator specified when creating the managed cluster.
	OrchestratorVersion *string `pulumi:"orchestratorVersion"`
	// OS Disk Size in GB to be used to specify the disk size for every machine in this master/agent pool. If you specify 0, it will apply the default osDisk size according to the vmSize specified.
	OsDiskSizeGB *int `pulumi:"osDiskSizeGB"`
	// OsType to be used to specify os type. Choose from Linux and Windows. Default to Linux.
	OsType *string `pulumi:"osType"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the managed cluster resource.
	ResourceName string `pulumi:"resourceName"`
	// ScaleSetEvictionPolicy to be used to specify eviction policy for low priority virtual machine scale set. Default to Delete.
	ScaleSetEvictionPolicy *string `pulumi:"scaleSetEvictionPolicy"`
	// ScaleSetPriority to be used to specify virtual machine scale set priority. Default to regular.
	ScaleSetPriority *string `pulumi:"scaleSetPriority"`
	// AgentPoolType represents types of an agent pool
	Type *string `pulumi:"type"`
	// Size of agent VMs.
	VmSize *string `pulumi:"vmSize"`
	// VNet SubnetID specifies the VNet's subnet identifier.
	VnetSubnetID *string `pulumi:"vnetSubnetID"`
}

// The set of arguments for constructing a AgentPool resource.
type AgentPoolArgs struct {
	// The name of the agent pool.
	AgentPoolName pulumi.StringInput
	// (PREVIEW) Availability zones for nodes. Must use VirtualMachineScaleSets AgentPoolType.
	AvailabilityZones pulumi.StringArrayInput
	// Number of agents (VMs) to host docker containers. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 1.
	Count pulumi.IntPtrInput
	// Whether to enable auto-scaler
	EnableAutoScaling pulumi.BoolPtrInput
	// Enable public IP for nodes
	EnableNodePublicIP pulumi.BoolPtrInput
	// Maximum number of nodes for auto-scaling
	MaxCount pulumi.IntPtrInput
	// Maximum number of pods that can run on a node.
	MaxPods pulumi.IntPtrInput
	// Minimum number of nodes for auto-scaling
	MinCount pulumi.IntPtrInput
	// Taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.
	NodeTaints pulumi.StringArrayInput
	// Version of orchestrator specified when creating the managed cluster.
	OrchestratorVersion pulumi.StringPtrInput
	// OS Disk Size in GB to be used to specify the disk size for every machine in this master/agent pool. If you specify 0, it will apply the default osDisk size according to the vmSize specified.
	OsDiskSizeGB pulumi.IntPtrInput
	// OsType to be used to specify os type. Choose from Linux and Windows. Default to Linux.
	OsType pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the managed cluster resource.
	ResourceName pulumi.StringInput
	// ScaleSetEvictionPolicy to be used to specify eviction policy for low priority virtual machine scale set. Default to Delete.
	ScaleSetEvictionPolicy pulumi.StringPtrInput
	// ScaleSetPriority to be used to specify virtual machine scale set priority. Default to regular.
	ScaleSetPriority pulumi.StringPtrInput
	// AgentPoolType represents types of an agent pool
	Type pulumi.StringPtrInput
	// Size of agent VMs.
	VmSize pulumi.StringPtrInput
	// VNet SubnetID specifies the VNet's subnet identifier.
	VnetSubnetID pulumi.StringPtrInput
}

func (AgentPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*agentPoolArgs)(nil)).Elem()
}
