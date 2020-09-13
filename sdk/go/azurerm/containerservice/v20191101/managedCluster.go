// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191101

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Managed cluster.
//
// ## Example Usage
// ### Create/Update Managed Cluster
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	containerservice "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/containerservice/v20191101"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := containerservice.NewManagedCluster(ctx, "managedCluster", &containerservice.ManagedClusterArgs{
// 			AddonProfiles: nil,
// 			AgentPoolProfiles: containerservice.ManagedClusterAgentPoolProfileArray{
// 				&containerservice.ManagedClusterAgentPoolProfileArgs{
// 					AvailabilityZones: pulumi.StringArray{
// 						pulumi.String("1"),
// 						pulumi.String("2"),
// 						pulumi.String("3"),
// 					},
// 					Count:              pulumi.Int(3),
// 					EnableNodePublicIP: pulumi.Bool(true),
// 					Name:               pulumi.String("nodepool1"),
// 					OsType:             pulumi.String("Linux"),
// 					Type:               pulumi.String("VirtualMachineScaleSets"),
// 					VmSize:             pulumi.String("Standard_DS1_v2"),
// 				},
// 			},
// 			DnsPrefix:               pulumi.String("dnsprefix1"),
// 			EnablePodSecurityPolicy: pulumi.Bool(true),
// 			EnableRBAC:              pulumi.Bool(true),
// 			KubernetesVersion:       pulumi.String(""),
// 			LinuxProfile: &containerservice.ContainerServiceLinuxProfileArgs{
// 				AdminUsername: pulumi.String("azureuser"),
// 				Ssh: &containerservice.ContainerServiceSshConfigurationArgs{
// 					PublicKeys: containerservice.ContainerServiceSshPublicKeyArray{
// 						&containerservice.ContainerServiceSshPublicKeyArgs{
// 							KeyData: pulumi.String("keydata"),
// 						},
// 					},
// 				},
// 			},
// 			Location: pulumi.String("location1"),
// 			NetworkProfile: &containerservice.ContainerServiceNetworkProfileArgs{
// 				LoadBalancerProfile: &containerservice.ManagedClusterLoadBalancerProfileArgs{
// 					ManagedOutboundIPs: &containerservice.ManagedClusterLoadBalancerProfileManagedOutboundIPsArgs{
// 						Count: pulumi.Int(2),
// 					},
// 				},
// 				LoadBalancerSku: pulumi.String("standard"),
// 				OutboundType:    pulumi.String("loadBalancer"),
// 			},
// 			ResourceGroupName: pulumi.String("rg1"),
// 			ResourceName:      pulumi.String("clustername1"),
// 			ServicePrincipalProfile: &containerservice.ManagedClusterServicePrincipalProfileArgs{
// 				ClientId: pulumi.String("clientid"),
// 				Secret:   pulumi.String("secret"),
// 			},
// 			Tags: pulumi.StringMap{
// 				"archv2": pulumi.String(""),
// 				"tier":   pulumi.String("production"),
// 			},
// 			WindowsProfile: &containerservice.ManagedClusterWindowsProfileArgs{
// 				AdminPassword: pulumi.String(fmt.Sprintf("%v%v", "replacePassword1234", "$")),
// 				AdminUsername: pulumi.String("azureuser"),
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
type ManagedCluster struct {
	pulumi.CustomResourceState

	// Profile of Azure Active Directory configuration.
	AadProfile ManagedClusterAADProfileResponsePtrOutput `pulumi:"aadProfile"`
	// Profile of managed cluster add-on.
	AddonProfiles ManagedClusterAddonProfileResponseMapOutput `pulumi:"addonProfiles"`
	// Properties of the agent pool.
	AgentPoolProfiles ManagedClusterAgentPoolProfileResponseArrayOutput `pulumi:"agentPoolProfiles"`
	// Access profile for managed cluster API server.
	ApiServerAccessProfile ManagedClusterAPIServerAccessProfileResponsePtrOutput `pulumi:"apiServerAccessProfile"`
	// DNS prefix specified when creating the managed cluster.
	DnsPrefix pulumi.StringPtrOutput `pulumi:"dnsPrefix"`
	// (DEPRECATING) Whether to enable Kubernetes pod security policy (preview). This feature is set for removal on October 15th, 2020. Learn more at aka.ms/aks/azpodpolicy.
	EnablePodSecurityPolicy pulumi.BoolPtrOutput `pulumi:"enablePodSecurityPolicy"`
	// Whether to enable Kubernetes Role-Based Access Control.
	EnableRBAC pulumi.BoolPtrOutput `pulumi:"enableRBAC"`
	// FQDN for the master pool.
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// The identity of the managed cluster, if configured.
	Identity ManagedClusterIdentityResponsePtrOutput `pulumi:"identity"`
	// Identities associated with the cluster.
	IdentityProfile ManagedClusterPropertiesResponseIdentityProfileMapOutput `pulumi:"identityProfile"`
	// Version of Kubernetes specified when creating the managed cluster.
	KubernetesVersion pulumi.StringPtrOutput `pulumi:"kubernetesVersion"`
	// Profile for Linux VMs in the container service cluster.
	LinuxProfile ContainerServiceLinuxProfileResponsePtrOutput `pulumi:"linuxProfile"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// The max number of agent pools for the managed cluster.
	MaxAgentPools pulumi.IntOutput `pulumi:"maxAgentPools"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Profile of network configuration.
	NetworkProfile ContainerServiceNetworkProfileResponsePtrOutput `pulumi:"networkProfile"`
	// Name of the resource group containing agent pool nodes.
	NodeResourceGroup pulumi.StringPtrOutput `pulumi:"nodeResourceGroup"`
	// FQDN of private cluster.
	PrivateFQDN pulumi.StringOutput `pulumi:"privateFQDN"`
	// The current deployment or provisioning state, which only appears in the response.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Information about a service principal identity for the cluster to use for manipulating Azure APIs.
	ServicePrincipalProfile ManagedClusterServicePrincipalProfileResponsePtrOutput `pulumi:"servicePrincipalProfile"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// Profile for Windows VMs in the container service cluster.
	WindowsProfile ManagedClusterWindowsProfileResponsePtrOutput `pulumi:"windowsProfile"`
}

// NewManagedCluster registers a new resource with the given unique name, arguments, and options.
func NewManagedCluster(ctx *pulumi.Context,
	name string, args *ManagedClusterArgs, opts ...pulumi.ResourceOption) (*ManagedCluster, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ResourceName == nil {
		return nil, errors.New("missing required argument 'ResourceName'")
	}
	if args == nil {
		args = &ManagedClusterArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:containerservice/latest:ManagedCluster"),
		},
		{
			Type: pulumi.String("azurerm:containerservice/v20170831:ManagedCluster"),
		},
		{
			Type: pulumi.String("azurerm:containerservice/v20180331:ManagedCluster"),
		},
		{
			Type: pulumi.String("azurerm:containerservice/v20180801preview:ManagedCluster"),
		},
		{
			Type: pulumi.String("azurerm:containerservice/v20190201:ManagedCluster"),
		},
		{
			Type: pulumi.String("azurerm:containerservice/v20190401:ManagedCluster"),
		},
		{
			Type: pulumi.String("azurerm:containerservice/v20190601:ManagedCluster"),
		},
		{
			Type: pulumi.String("azurerm:containerservice/v20190801:ManagedCluster"),
		},
		{
			Type: pulumi.String("azurerm:containerservice/v20191001:ManagedCluster"),
		},
		{
			Type: pulumi.String("azurerm:containerservice/v20200101:ManagedCluster"),
		},
		{
			Type: pulumi.String("azurerm:containerservice/v20200201:ManagedCluster"),
		},
		{
			Type: pulumi.String("azurerm:containerservice/v20200301:ManagedCluster"),
		},
		{
			Type: pulumi.String("azurerm:containerservice/v20200401:ManagedCluster"),
		},
		{
			Type: pulumi.String("azurerm:containerservice/v20200601:ManagedCluster"),
		},
		{
			Type: pulumi.String("azurerm:containerservice/v20200701:ManagedCluster"),
		},
		{
			Type: pulumi.String("azurerm:containerservice/v20200901:ManagedCluster"),
		},
	})
	opts = append(opts, aliases)
	var resource ManagedCluster
	err := ctx.RegisterResource("azurerm:containerservice/v20191101:ManagedCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagedCluster gets an existing ManagedCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagedCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedClusterState, opts ...pulumi.ResourceOption) (*ManagedCluster, error) {
	var resource ManagedCluster
	err := ctx.ReadResource("azurerm:containerservice/v20191101:ManagedCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagedCluster resources.
type managedClusterState struct {
	// Profile of Azure Active Directory configuration.
	AadProfile *ManagedClusterAADProfileResponse `pulumi:"aadProfile"`
	// Profile of managed cluster add-on.
	AddonProfiles map[string]ManagedClusterAddonProfileResponse `pulumi:"addonProfiles"`
	// Properties of the agent pool.
	AgentPoolProfiles []ManagedClusterAgentPoolProfileResponse `pulumi:"agentPoolProfiles"`
	// Access profile for managed cluster API server.
	ApiServerAccessProfile *ManagedClusterAPIServerAccessProfileResponse `pulumi:"apiServerAccessProfile"`
	// DNS prefix specified when creating the managed cluster.
	DnsPrefix *string `pulumi:"dnsPrefix"`
	// (DEPRECATING) Whether to enable Kubernetes pod security policy (preview). This feature is set for removal on October 15th, 2020. Learn more at aka.ms/aks/azpodpolicy.
	EnablePodSecurityPolicy *bool `pulumi:"enablePodSecurityPolicy"`
	// Whether to enable Kubernetes Role-Based Access Control.
	EnableRBAC *bool `pulumi:"enableRBAC"`
	// FQDN for the master pool.
	Fqdn *string `pulumi:"fqdn"`
	// The identity of the managed cluster, if configured.
	Identity *ManagedClusterIdentityResponse `pulumi:"identity"`
	// Identities associated with the cluster.
	IdentityProfile map[string]ManagedClusterPropertiesResponseIdentityProfile `pulumi:"identityProfile"`
	// Version of Kubernetes specified when creating the managed cluster.
	KubernetesVersion *string `pulumi:"kubernetesVersion"`
	// Profile for Linux VMs in the container service cluster.
	LinuxProfile *ContainerServiceLinuxProfileResponse `pulumi:"linuxProfile"`
	// Resource location
	Location *string `pulumi:"location"`
	// The max number of agent pools for the managed cluster.
	MaxAgentPools *int `pulumi:"maxAgentPools"`
	// Resource name
	Name *string `pulumi:"name"`
	// Profile of network configuration.
	NetworkProfile *ContainerServiceNetworkProfileResponse `pulumi:"networkProfile"`
	// Name of the resource group containing agent pool nodes.
	NodeResourceGroup *string `pulumi:"nodeResourceGroup"`
	// FQDN of private cluster.
	PrivateFQDN *string `pulumi:"privateFQDN"`
	// The current deployment or provisioning state, which only appears in the response.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Information about a service principal identity for the cluster to use for manipulating Azure APIs.
	ServicePrincipalProfile *ManagedClusterServicePrincipalProfileResponse `pulumi:"servicePrincipalProfile"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
	// Profile for Windows VMs in the container service cluster.
	WindowsProfile *ManagedClusterWindowsProfileResponse `pulumi:"windowsProfile"`
}

type ManagedClusterState struct {
	// Profile of Azure Active Directory configuration.
	AadProfile ManagedClusterAADProfileResponsePtrInput
	// Profile of managed cluster add-on.
	AddonProfiles ManagedClusterAddonProfileResponseMapInput
	// Properties of the agent pool.
	AgentPoolProfiles ManagedClusterAgentPoolProfileResponseArrayInput
	// Access profile for managed cluster API server.
	ApiServerAccessProfile ManagedClusterAPIServerAccessProfileResponsePtrInput
	// DNS prefix specified when creating the managed cluster.
	DnsPrefix pulumi.StringPtrInput
	// (DEPRECATING) Whether to enable Kubernetes pod security policy (preview). This feature is set for removal on October 15th, 2020. Learn more at aka.ms/aks/azpodpolicy.
	EnablePodSecurityPolicy pulumi.BoolPtrInput
	// Whether to enable Kubernetes Role-Based Access Control.
	EnableRBAC pulumi.BoolPtrInput
	// FQDN for the master pool.
	Fqdn pulumi.StringPtrInput
	// The identity of the managed cluster, if configured.
	Identity ManagedClusterIdentityResponsePtrInput
	// Identities associated with the cluster.
	IdentityProfile ManagedClusterPropertiesResponseIdentityProfileMapInput
	// Version of Kubernetes specified when creating the managed cluster.
	KubernetesVersion pulumi.StringPtrInput
	// Profile for Linux VMs in the container service cluster.
	LinuxProfile ContainerServiceLinuxProfileResponsePtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// The max number of agent pools for the managed cluster.
	MaxAgentPools pulumi.IntPtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// Profile of network configuration.
	NetworkProfile ContainerServiceNetworkProfileResponsePtrInput
	// Name of the resource group containing agent pool nodes.
	NodeResourceGroup pulumi.StringPtrInput
	// FQDN of private cluster.
	PrivateFQDN pulumi.StringPtrInput
	// The current deployment or provisioning state, which only appears in the response.
	ProvisioningState pulumi.StringPtrInput
	// Information about a service principal identity for the cluster to use for manipulating Azure APIs.
	ServicePrincipalProfile ManagedClusterServicePrincipalProfileResponsePtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Resource type
	Type pulumi.StringPtrInput
	// Profile for Windows VMs in the container service cluster.
	WindowsProfile ManagedClusterWindowsProfileResponsePtrInput
}

func (ManagedClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedClusterState)(nil)).Elem()
}

type managedClusterArgs struct {
	// Profile of Azure Active Directory configuration.
	AadProfile *ManagedClusterAADProfile `pulumi:"aadProfile"`
	// Profile of managed cluster add-on.
	AddonProfiles map[string]ManagedClusterAddonProfile `pulumi:"addonProfiles"`
	// Properties of the agent pool.
	AgentPoolProfiles []ManagedClusterAgentPoolProfile `pulumi:"agentPoolProfiles"`
	// Access profile for managed cluster API server.
	ApiServerAccessProfile *ManagedClusterAPIServerAccessProfile `pulumi:"apiServerAccessProfile"`
	// DNS prefix specified when creating the managed cluster.
	DnsPrefix *string `pulumi:"dnsPrefix"`
	// (DEPRECATING) Whether to enable Kubernetes pod security policy (preview). This feature is set for removal on October 15th, 2020. Learn more at aka.ms/aks/azpodpolicy.
	EnablePodSecurityPolicy *bool `pulumi:"enablePodSecurityPolicy"`
	// Whether to enable Kubernetes Role-Based Access Control.
	EnableRBAC *bool `pulumi:"enableRBAC"`
	// The identity of the managed cluster, if configured.
	Identity *ManagedClusterIdentity `pulumi:"identity"`
	// Identities associated with the cluster.
	IdentityProfile map[string]ManagedClusterPropertiesIdentityProfile `pulumi:"identityProfile"`
	// Version of Kubernetes specified when creating the managed cluster.
	KubernetesVersion *string `pulumi:"kubernetesVersion"`
	// Profile for Linux VMs in the container service cluster.
	LinuxProfile *ContainerServiceLinuxProfile `pulumi:"linuxProfile"`
	// Resource location
	Location string `pulumi:"location"`
	// Profile of network configuration.
	NetworkProfile *ContainerServiceNetworkProfile `pulumi:"networkProfile"`
	// Name of the resource group containing agent pool nodes.
	NodeResourceGroup *string `pulumi:"nodeResourceGroup"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the managed cluster resource.
	ResourceName string `pulumi:"resourceName"`
	// Information about a service principal identity for the cluster to use for manipulating Azure APIs.
	ServicePrincipalProfile *ManagedClusterServicePrincipalProfile `pulumi:"servicePrincipalProfile"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Profile for Windows VMs in the container service cluster.
	WindowsProfile *ManagedClusterWindowsProfile `pulumi:"windowsProfile"`
}

// The set of arguments for constructing a ManagedCluster resource.
type ManagedClusterArgs struct {
	// Profile of Azure Active Directory configuration.
	AadProfile ManagedClusterAADProfilePtrInput
	// Profile of managed cluster add-on.
	AddonProfiles ManagedClusterAddonProfileMapInput
	// Properties of the agent pool.
	AgentPoolProfiles ManagedClusterAgentPoolProfileArrayInput
	// Access profile for managed cluster API server.
	ApiServerAccessProfile ManagedClusterAPIServerAccessProfilePtrInput
	// DNS prefix specified when creating the managed cluster.
	DnsPrefix pulumi.StringPtrInput
	// (DEPRECATING) Whether to enable Kubernetes pod security policy (preview). This feature is set for removal on October 15th, 2020. Learn more at aka.ms/aks/azpodpolicy.
	EnablePodSecurityPolicy pulumi.BoolPtrInput
	// Whether to enable Kubernetes Role-Based Access Control.
	EnableRBAC pulumi.BoolPtrInput
	// The identity of the managed cluster, if configured.
	Identity ManagedClusterIdentityPtrInput
	// Identities associated with the cluster.
	IdentityProfile ManagedClusterPropertiesIdentityProfileMapInput
	// Version of Kubernetes specified when creating the managed cluster.
	KubernetesVersion pulumi.StringPtrInput
	// Profile for Linux VMs in the container service cluster.
	LinuxProfile ContainerServiceLinuxProfilePtrInput
	// Resource location
	Location pulumi.StringInput
	// Profile of network configuration.
	NetworkProfile ContainerServiceNetworkProfilePtrInput
	// Name of the resource group containing agent pool nodes.
	NodeResourceGroup pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the managed cluster resource.
	ResourceName pulumi.StringInput
	// Information about a service principal identity for the cluster to use for manipulating Azure APIs.
	ServicePrincipalProfile ManagedClusterServicePrincipalProfilePtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Profile for Windows VMs in the container service cluster.
	WindowsProfile ManagedClusterWindowsProfilePtrInput
}

func (ManagedClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedClusterArgs)(nil)).Elem()
}
