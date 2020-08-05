// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190601

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Managed cluster.
type ManagedCluster struct {
	pulumi.CustomResourceState

	// The identity of the managed cluster, if configured.
	Identity ManagedClusterIdentityResponsePtrOutput `pulumi:"identity"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of a managed cluster.
	Properties ManagedClusterPropertiesResponseOutput `pulumi:"properties"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewManagedCluster registers a new resource with the given unique name, arguments, and options.
func NewManagedCluster(ctx *pulumi.Context,
	name string, args *ManagedClusterArgs, opts ...pulumi.ResourceOption) (*ManagedCluster, error) {
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
		args = &ManagedClusterArgs{}
	}
	var resource ManagedCluster
	err := ctx.RegisterResource("azurerm:containerservice/v20190601:ManagedCluster", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azurerm:containerservice/v20190601:ManagedCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagedCluster resources.
type managedClusterState struct {
	// The identity of the managed cluster, if configured.
	Identity *ManagedClusterIdentityResponse `pulumi:"identity"`
	// Resource location
	Location *string `pulumi:"location"`
	// Resource name
	Name *string `pulumi:"name"`
	// Properties of a managed cluster.
	Properties *ManagedClusterPropertiesResponse `pulumi:"properties"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
}

type ManagedClusterState struct {
	// The identity of the managed cluster, if configured.
	Identity ManagedClusterIdentityResponsePtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// Properties of a managed cluster.
	Properties ManagedClusterPropertiesResponsePtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Resource type
	Type pulumi.StringPtrInput
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
	// (PREVIEW) Authorized IP Ranges to kubernetes API server.
	ApiServerAuthorizedIPRanges []string `pulumi:"apiServerAuthorizedIPRanges"`
	// DNS prefix specified when creating the managed cluster.
	DnsPrefix *string `pulumi:"dnsPrefix"`
	// (PREVIEW) Whether to enable Kubernetes Pod security policy.
	EnablePodSecurityPolicy *bool `pulumi:"enablePodSecurityPolicy"`
	// Whether to enable Kubernetes Role-Based Access Control.
	EnableRBAC *bool `pulumi:"enableRBAC"`
	// The identity of the managed cluster, if configured.
	Identity *ManagedClusterIdentity `pulumi:"identity"`
	// Version of Kubernetes specified when creating the managed cluster.
	KubernetesVersion *string `pulumi:"kubernetesVersion"`
	// Profile for Linux VMs in the container service cluster.
	LinuxProfile *ContainerServiceLinuxProfile `pulumi:"linuxProfile"`
	// Resource location
	Location string `pulumi:"location"`
	// The name of the managed cluster resource.
	Name string `pulumi:"name"`
	// Profile of network configuration.
	NetworkProfile *ContainerServiceNetworkProfile `pulumi:"networkProfile"`
	// Name of the resource group containing agent pool nodes.
	NodeResourceGroup *string `pulumi:"nodeResourceGroup"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
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
	// (PREVIEW) Authorized IP Ranges to kubernetes API server.
	ApiServerAuthorizedIPRanges pulumi.StringArrayInput
	// DNS prefix specified when creating the managed cluster.
	DnsPrefix pulumi.StringPtrInput
	// (PREVIEW) Whether to enable Kubernetes Pod security policy.
	EnablePodSecurityPolicy pulumi.BoolPtrInput
	// Whether to enable Kubernetes Role-Based Access Control.
	EnableRBAC pulumi.BoolPtrInput
	// The identity of the managed cluster, if configured.
	Identity ManagedClusterIdentityPtrInput
	// Version of Kubernetes specified when creating the managed cluster.
	KubernetesVersion pulumi.StringPtrInput
	// Profile for Linux VMs in the container service cluster.
	LinuxProfile ContainerServiceLinuxProfilePtrInput
	// Resource location
	Location pulumi.StringInput
	// The name of the managed cluster resource.
	Name pulumi.StringInput
	// Profile of network configuration.
	NetworkProfile ContainerServiceNetworkProfilePtrInput
	// Name of the resource group containing agent pool nodes.
	NodeResourceGroup pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
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
