# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables
from . import outputs

__all__ = [
    'ContainerServiceLinuxProfileResponse',
    'ContainerServiceNetworkProfileResponse',
    'ContainerServiceSshConfigurationResponse',
    'ContainerServiceSshPublicKeyResponse',
    'CredentialResultResponseResult',
    'ManagedClusterAADProfileResponse',
    'ManagedClusterAPIServerAccessProfileResponse',
    'ManagedClusterAddonProfileResponse',
    'ManagedClusterAgentPoolProfileResponse',
    'ManagedClusterIdentityResponse',
    'ManagedClusterLoadBalancerProfileResponse',
    'ManagedClusterLoadBalancerProfileResponseManagedOutboundIPs',
    'ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixes',
    'ManagedClusterLoadBalancerProfileResponseOutboundIPs',
    'ManagedClusterServicePrincipalProfileResponse',
    'ManagedClusterWindowsProfileResponse',
    'ResourceReferenceResponse',
]

@pulumi.output_type
class ContainerServiceLinuxProfileResponse(dict):
    """
    Profile for Linux VMs in the container service cluster.
    """
    def __init__(__self__, *,
                 admin_username: str,
                 ssh: 'outputs.ContainerServiceSshConfigurationResponse'):
        """
        Profile for Linux VMs in the container service cluster.
        :param str admin_username: The administrator username to use for Linux VMs.
        :param 'ContainerServiceSshConfigurationResponseArgs' ssh: SSH configuration for Linux-based VMs running on Azure.
        """
        pulumi.set(__self__, "admin_username", admin_username)
        pulumi.set(__self__, "ssh", ssh)

    @property
    @pulumi.getter(name="adminUsername")
    def admin_username(self) -> str:
        """
        The administrator username to use for Linux VMs.
        """
        return pulumi.get(self, "admin_username")

    @property
    @pulumi.getter
    def ssh(self) -> 'outputs.ContainerServiceSshConfigurationResponse':
        """
        SSH configuration for Linux-based VMs running on Azure.
        """
        return pulumi.get(self, "ssh")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ContainerServiceNetworkProfileResponse(dict):
    """
    Profile of network configuration.
    """
    def __init__(__self__, *,
                 dns_service_ip: Optional[str] = None,
                 docker_bridge_cidr: Optional[str] = None,
                 load_balancer_profile: Optional['outputs.ManagedClusterLoadBalancerProfileResponse'] = None,
                 load_balancer_sku: Optional[str] = None,
                 network_plugin: Optional[str] = None,
                 network_policy: Optional[str] = None,
                 pod_cidr: Optional[str] = None,
                 service_cidr: Optional[str] = None):
        """
        Profile of network configuration.
        :param str dns_service_ip: An IP address assigned to the Kubernetes DNS service. It must be within the Kubernetes service address range specified in serviceCidr.
        :param str docker_bridge_cidr: A CIDR notation IP range assigned to the Docker bridge network. It must not overlap with any Subnet IP ranges or the Kubernetes service address range.
        :param 'ManagedClusterLoadBalancerProfileResponseArgs' load_balancer_profile: Profile of the cluster load balancer.
        :param str load_balancer_sku: The load balancer sku for the managed cluster.
        :param str network_plugin: Network plugin used for building Kubernetes network.
        :param str network_policy: Network policy used for building Kubernetes network.
        :param str pod_cidr: A CIDR notation IP range from which to assign pod IPs when kubenet is used.
        :param str service_cidr: A CIDR notation IP range from which to assign service cluster IPs. It must not overlap with any Subnet IP ranges.
        """
        if dns_service_ip is not None:
            pulumi.set(__self__, "dns_service_ip", dns_service_ip)
        if docker_bridge_cidr is not None:
            pulumi.set(__self__, "docker_bridge_cidr", docker_bridge_cidr)
        if load_balancer_profile is not None:
            pulumi.set(__self__, "load_balancer_profile", load_balancer_profile)
        if load_balancer_sku is not None:
            pulumi.set(__self__, "load_balancer_sku", load_balancer_sku)
        if network_plugin is not None:
            pulumi.set(__self__, "network_plugin", network_plugin)
        if network_policy is not None:
            pulumi.set(__self__, "network_policy", network_policy)
        if pod_cidr is not None:
            pulumi.set(__self__, "pod_cidr", pod_cidr)
        if service_cidr is not None:
            pulumi.set(__self__, "service_cidr", service_cidr)

    @property
    @pulumi.getter(name="dnsServiceIP")
    def dns_service_ip(self) -> Optional[str]:
        """
        An IP address assigned to the Kubernetes DNS service. It must be within the Kubernetes service address range specified in serviceCidr.
        """
        return pulumi.get(self, "dns_service_ip")

    @property
    @pulumi.getter(name="dockerBridgeCidr")
    def docker_bridge_cidr(self) -> Optional[str]:
        """
        A CIDR notation IP range assigned to the Docker bridge network. It must not overlap with any Subnet IP ranges or the Kubernetes service address range.
        """
        return pulumi.get(self, "docker_bridge_cidr")

    @property
    @pulumi.getter(name="loadBalancerProfile")
    def load_balancer_profile(self) -> Optional['outputs.ManagedClusterLoadBalancerProfileResponse']:
        """
        Profile of the cluster load balancer.
        """
        return pulumi.get(self, "load_balancer_profile")

    @property
    @pulumi.getter(name="loadBalancerSku")
    def load_balancer_sku(self) -> Optional[str]:
        """
        The load balancer sku for the managed cluster.
        """
        return pulumi.get(self, "load_balancer_sku")

    @property
    @pulumi.getter(name="networkPlugin")
    def network_plugin(self) -> Optional[str]:
        """
        Network plugin used for building Kubernetes network.
        """
        return pulumi.get(self, "network_plugin")

    @property
    @pulumi.getter(name="networkPolicy")
    def network_policy(self) -> Optional[str]:
        """
        Network policy used for building Kubernetes network.
        """
        return pulumi.get(self, "network_policy")

    @property
    @pulumi.getter(name="podCidr")
    def pod_cidr(self) -> Optional[str]:
        """
        A CIDR notation IP range from which to assign pod IPs when kubenet is used.
        """
        return pulumi.get(self, "pod_cidr")

    @property
    @pulumi.getter(name="serviceCidr")
    def service_cidr(self) -> Optional[str]:
        """
        A CIDR notation IP range from which to assign service cluster IPs. It must not overlap with any Subnet IP ranges.
        """
        return pulumi.get(self, "service_cidr")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ContainerServiceSshConfigurationResponse(dict):
    """
    SSH configuration for Linux-based VMs running on Azure.
    """
    def __init__(__self__, *,
                 public_keys: List['outputs.ContainerServiceSshPublicKeyResponse']):
        """
        SSH configuration for Linux-based VMs running on Azure.
        :param List['ContainerServiceSshPublicKeyResponseArgs'] public_keys: The list of SSH public keys used to authenticate with Linux-based VMs. Only expect one key specified.
        """
        pulumi.set(__self__, "public_keys", public_keys)

    @property
    @pulumi.getter(name="publicKeys")
    def public_keys(self) -> List['outputs.ContainerServiceSshPublicKeyResponse']:
        """
        The list of SSH public keys used to authenticate with Linux-based VMs. Only expect one key specified.
        """
        return pulumi.get(self, "public_keys")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ContainerServiceSshPublicKeyResponse(dict):
    """
    Contains information about SSH certificate public key data.
    """
    def __init__(__self__, *,
                 key_data: str):
        """
        Contains information about SSH certificate public key data.
        :param str key_data: Certificate public key used to authenticate with VMs through SSH. The certificate must be in PEM format with or without headers.
        """
        pulumi.set(__self__, "key_data", key_data)

    @property
    @pulumi.getter(name="keyData")
    def key_data(self) -> str:
        """
        Certificate public key used to authenticate with VMs through SSH. The certificate must be in PEM format with or without headers.
        """
        return pulumi.get(self, "key_data")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class CredentialResultResponseResult(dict):
    """
    The credential result response.
    """
    def __init__(__self__, *,
                 name: str,
                 value: str):
        """
        The credential result response.
        :param str name: The name of the credential.
        :param str value: Base64-encoded Kubernetes configuration file.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the credential.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        Base64-encoded Kubernetes configuration file.
        """
        return pulumi.get(self, "value")


@pulumi.output_type
class ManagedClusterAADProfileResponse(dict):
    """
    AADProfile specifies attributes for Azure Active Directory integration.
    """
    def __init__(__self__, *,
                 client_app_id: str,
                 server_app_id: str,
                 server_app_secret: Optional[str] = None,
                 tenant_id: Optional[str] = None):
        """
        AADProfile specifies attributes for Azure Active Directory integration.
        :param str client_app_id: The client AAD application ID.
        :param str server_app_id: The server AAD application ID.
        :param str server_app_secret: The server AAD application secret.
        :param str tenant_id: The AAD tenant ID to use for authentication. If not specified, will use the tenant of the deployment subscription.
        """
        pulumi.set(__self__, "client_app_id", client_app_id)
        pulumi.set(__self__, "server_app_id", server_app_id)
        if server_app_secret is not None:
            pulumi.set(__self__, "server_app_secret", server_app_secret)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter(name="clientAppID")
    def client_app_id(self) -> str:
        """
        The client AAD application ID.
        """
        return pulumi.get(self, "client_app_id")

    @property
    @pulumi.getter(name="serverAppID")
    def server_app_id(self) -> str:
        """
        The server AAD application ID.
        """
        return pulumi.get(self, "server_app_id")

    @property
    @pulumi.getter(name="serverAppSecret")
    def server_app_secret(self) -> Optional[str]:
        """
        The server AAD application secret.
        """
        return pulumi.get(self, "server_app_secret")

    @property
    @pulumi.getter(name="tenantID")
    def tenant_id(self) -> Optional[str]:
        """
        The AAD tenant ID to use for authentication. If not specified, will use the tenant of the deployment subscription.
        """
        return pulumi.get(self, "tenant_id")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ManagedClusterAPIServerAccessProfileResponse(dict):
    """
    Access profile for managed cluster API server.
    """
    def __init__(__self__, *,
                 authorized_ip_ranges: Optional[List[str]] = None,
                 enable_private_cluster: Optional[bool] = None):
        """
        Access profile for managed cluster API server.
        :param List[str] authorized_ip_ranges: Authorized IP Ranges to kubernetes API server.
        :param bool enable_private_cluster: Whether to create the cluster as a private cluster or not.
        """
        if authorized_ip_ranges is not None:
            pulumi.set(__self__, "authorized_ip_ranges", authorized_ip_ranges)
        if enable_private_cluster is not None:
            pulumi.set(__self__, "enable_private_cluster", enable_private_cluster)

    @property
    @pulumi.getter(name="authorizedIPRanges")
    def authorized_ip_ranges(self) -> Optional[List[str]]:
        """
        Authorized IP Ranges to kubernetes API server.
        """
        return pulumi.get(self, "authorized_ip_ranges")

    @property
    @pulumi.getter(name="enablePrivateCluster")
    def enable_private_cluster(self) -> Optional[bool]:
        """
        Whether to create the cluster as a private cluster or not.
        """
        return pulumi.get(self, "enable_private_cluster")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ManagedClusterAddonProfileResponse(dict):
    """
    A Kubernetes add-on profile for a managed cluster.
    """
    def __init__(__self__, *,
                 enabled: bool,
                 config: Optional[Mapping[str, str]] = None):
        """
        A Kubernetes add-on profile for a managed cluster.
        :param bool enabled: Whether the add-on is enabled or not.
        :param Mapping[str, str] config: Key-value pairs for configuring an add-on.
        """
        pulumi.set(__self__, "enabled", enabled)
        if config is not None:
            pulumi.set(__self__, "config", config)

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        """
        Whether the add-on is enabled or not.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def config(self) -> Optional[Mapping[str, str]]:
        """
        Key-value pairs for configuring an add-on.
        """
        return pulumi.get(self, "config")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ManagedClusterAgentPoolProfileResponse(dict):
    """
    Profile for the container service agent pool.
    """
    def __init__(__self__, *,
                 name: str,
                 provisioning_state: str,
                 availability_zones: Optional[List[str]] = None,
                 count: Optional[float] = None,
                 enable_auto_scaling: Optional[bool] = None,
                 enable_node_public_ip: Optional[bool] = None,
                 max_count: Optional[float] = None,
                 max_pods: Optional[float] = None,
                 min_count: Optional[float] = None,
                 node_taints: Optional[List[str]] = None,
                 orchestrator_version: Optional[str] = None,
                 os_disk_size_gb: Optional[float] = None,
                 os_type: Optional[str] = None,
                 scale_set_eviction_policy: Optional[str] = None,
                 scale_set_priority: Optional[str] = None,
                 type: Optional[str] = None,
                 vm_size: Optional[str] = None,
                 vnet_subnet_id: Optional[str] = None):
        """
        Profile for the container service agent pool.
        :param str name: Unique name of the agent pool profile in the context of the subscription and resource group.
        :param str provisioning_state: The current deployment or provisioning state, which only appears in the response.
        :param List[str] availability_zones: (PREVIEW) Availability zones for nodes. Must use VirtualMachineScaleSets AgentPoolType.
        :param float count: Number of agents (VMs) to host docker containers. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 1.
        :param bool enable_auto_scaling: Whether to enable auto-scaler
        :param bool enable_node_public_ip: Enable public IP for nodes
        :param float max_count: Maximum number of nodes for auto-scaling
        :param float max_pods: Maximum number of pods that can run on a node.
        :param float min_count: Minimum number of nodes for auto-scaling
        :param List[str] node_taints: Taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.
        :param str orchestrator_version: Version of orchestrator specified when creating the managed cluster.
        :param float os_disk_size_gb: OS Disk Size in GB to be used to specify the disk size for every machine in this master/agent pool. If you specify 0, it will apply the default osDisk size according to the vmSize specified.
        :param str os_type: OsType to be used to specify os type. Choose from Linux and Windows. Default to Linux.
        :param str scale_set_eviction_policy: ScaleSetEvictionPolicy to be used to specify eviction policy for low priority virtual machine scale set. Default to Delete.
        :param str scale_set_priority: ScaleSetPriority to be used to specify virtual machine scale set priority. Default to regular.
        :param str type: AgentPoolType represents types of an agent pool
        :param str vm_size: Size of agent VMs.
        :param str vnet_subnet_id: VNet SubnetID specifies the VNet's subnet identifier.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if availability_zones is not None:
            pulumi.set(__self__, "availability_zones", availability_zones)
        if count is not None:
            pulumi.set(__self__, "count", count)
        if enable_auto_scaling is not None:
            pulumi.set(__self__, "enable_auto_scaling", enable_auto_scaling)
        if enable_node_public_ip is not None:
            pulumi.set(__self__, "enable_node_public_ip", enable_node_public_ip)
        if max_count is not None:
            pulumi.set(__self__, "max_count", max_count)
        if max_pods is not None:
            pulumi.set(__self__, "max_pods", max_pods)
        if min_count is not None:
            pulumi.set(__self__, "min_count", min_count)
        if node_taints is not None:
            pulumi.set(__self__, "node_taints", node_taints)
        if orchestrator_version is not None:
            pulumi.set(__self__, "orchestrator_version", orchestrator_version)
        if os_disk_size_gb is not None:
            pulumi.set(__self__, "os_disk_size_gb", os_disk_size_gb)
        if os_type is not None:
            pulumi.set(__self__, "os_type", os_type)
        if scale_set_eviction_policy is not None:
            pulumi.set(__self__, "scale_set_eviction_policy", scale_set_eviction_policy)
        if scale_set_priority is not None:
            pulumi.set(__self__, "scale_set_priority", scale_set_priority)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if vm_size is not None:
            pulumi.set(__self__, "vm_size", vm_size)
        if vnet_subnet_id is not None:
            pulumi.set(__self__, "vnet_subnet_id", vnet_subnet_id)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Unique name of the agent pool profile in the context of the subscription and resource group.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        The current deployment or provisioning state, which only appears in the response.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="availabilityZones")
    def availability_zones(self) -> Optional[List[str]]:
        """
        (PREVIEW) Availability zones for nodes. Must use VirtualMachineScaleSets AgentPoolType.
        """
        return pulumi.get(self, "availability_zones")

    @property
    @pulumi.getter
    def count(self) -> Optional[float]:
        """
        Number of agents (VMs) to host docker containers. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 1.
        """
        return pulumi.get(self, "count")

    @property
    @pulumi.getter(name="enableAutoScaling")
    def enable_auto_scaling(self) -> Optional[bool]:
        """
        Whether to enable auto-scaler
        """
        return pulumi.get(self, "enable_auto_scaling")

    @property
    @pulumi.getter(name="enableNodePublicIP")
    def enable_node_public_ip(self) -> Optional[bool]:
        """
        Enable public IP for nodes
        """
        return pulumi.get(self, "enable_node_public_ip")

    @property
    @pulumi.getter(name="maxCount")
    def max_count(self) -> Optional[float]:
        """
        Maximum number of nodes for auto-scaling
        """
        return pulumi.get(self, "max_count")

    @property
    @pulumi.getter(name="maxPods")
    def max_pods(self) -> Optional[float]:
        """
        Maximum number of pods that can run on a node.
        """
        return pulumi.get(self, "max_pods")

    @property
    @pulumi.getter(name="minCount")
    def min_count(self) -> Optional[float]:
        """
        Minimum number of nodes for auto-scaling
        """
        return pulumi.get(self, "min_count")

    @property
    @pulumi.getter(name="nodeTaints")
    def node_taints(self) -> Optional[List[str]]:
        """
        Taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.
        """
        return pulumi.get(self, "node_taints")

    @property
    @pulumi.getter(name="orchestratorVersion")
    def orchestrator_version(self) -> Optional[str]:
        """
        Version of orchestrator specified when creating the managed cluster.
        """
        return pulumi.get(self, "orchestrator_version")

    @property
    @pulumi.getter(name="osDiskSizeGB")
    def os_disk_size_gb(self) -> Optional[float]:
        """
        OS Disk Size in GB to be used to specify the disk size for every machine in this master/agent pool. If you specify 0, it will apply the default osDisk size according to the vmSize specified.
        """
        return pulumi.get(self, "os_disk_size_gb")

    @property
    @pulumi.getter(name="osType")
    def os_type(self) -> Optional[str]:
        """
        OsType to be used to specify os type. Choose from Linux and Windows. Default to Linux.
        """
        return pulumi.get(self, "os_type")

    @property
    @pulumi.getter(name="scaleSetEvictionPolicy")
    def scale_set_eviction_policy(self) -> Optional[str]:
        """
        ScaleSetEvictionPolicy to be used to specify eviction policy for low priority virtual machine scale set. Default to Delete.
        """
        return pulumi.get(self, "scale_set_eviction_policy")

    @property
    @pulumi.getter(name="scaleSetPriority")
    def scale_set_priority(self) -> Optional[str]:
        """
        ScaleSetPriority to be used to specify virtual machine scale set priority. Default to regular.
        """
        return pulumi.get(self, "scale_set_priority")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        AgentPoolType represents types of an agent pool
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="vmSize")
    def vm_size(self) -> Optional[str]:
        """
        Size of agent VMs.
        """
        return pulumi.get(self, "vm_size")

    @property
    @pulumi.getter(name="vnetSubnetID")
    def vnet_subnet_id(self) -> Optional[str]:
        """
        VNet SubnetID specifies the VNet's subnet identifier.
        """
        return pulumi.get(self, "vnet_subnet_id")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ManagedClusterIdentityResponse(dict):
    """
    Identity for the managed cluster.
    """
    def __init__(__self__, *,
                 principal_id: str,
                 tenant_id: str,
                 type: Optional[str] = None):
        """
        Identity for the managed cluster.
        :param str principal_id: The principal id of the system assigned identity which is used by master components.
        :param str tenant_id: The tenant id of the system assigned identity which is used by master components.
        :param str type: The type of identity used for the managed cluster. Type 'SystemAssigned' will use an implicitly created identity in master components and an auto-created user assigned identity in MC_ resource group in agent nodes. Type 'None' will not use MSI for the managed cluster, service principal will be used instead.
        """
        pulumi.set(__self__, "principal_id", principal_id)
        pulumi.set(__self__, "tenant_id", tenant_id)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> str:
        """
        The principal id of the system assigned identity which is used by master components.
        """
        return pulumi.get(self, "principal_id")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> str:
        """
        The tenant id of the system assigned identity which is used by master components.
        """
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        The type of identity used for the managed cluster. Type 'SystemAssigned' will use an implicitly created identity in master components and an auto-created user assigned identity in MC_ resource group in agent nodes. Type 'None' will not use MSI for the managed cluster, service principal will be used instead.
        """
        return pulumi.get(self, "type")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ManagedClusterLoadBalancerProfileResponse(dict):
    """
    Profile of the managed cluster load balancer
    """
    def __init__(__self__, *,
                 effective_outbound_ips: Optional[List['outputs.ResourceReferenceResponse']] = None,
                 managed_outbound_ips: Optional['outputs.ManagedClusterLoadBalancerProfileResponseManagedOutboundIPs'] = None,
                 outbound_ip_prefixes: Optional['outputs.ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixes'] = None,
                 outbound_ips: Optional['outputs.ManagedClusterLoadBalancerProfileResponseOutboundIPs'] = None):
        """
        Profile of the managed cluster load balancer
        :param List['ResourceReferenceResponseArgs'] effective_outbound_ips: The effective outbound IP resources of the cluster load balancer.
        :param 'ManagedClusterLoadBalancerProfileResponseManagedOutboundIPsArgs' managed_outbound_ips: Desired managed outbound IPs for the cluster load balancer.
        :param 'ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixesArgs' outbound_ip_prefixes: Desired outbound IP Prefix resources for the cluster load balancer.
        :param 'ManagedClusterLoadBalancerProfileResponseOutboundIPsArgs' outbound_ips: Desired outbound IP resources for the cluster load balancer.
        """
        if effective_outbound_ips is not None:
            pulumi.set(__self__, "effective_outbound_ips", effective_outbound_ips)
        if managed_outbound_ips is not None:
            pulumi.set(__self__, "managed_outbound_ips", managed_outbound_ips)
        if outbound_ip_prefixes is not None:
            pulumi.set(__self__, "outbound_ip_prefixes", outbound_ip_prefixes)
        if outbound_ips is not None:
            pulumi.set(__self__, "outbound_ips", outbound_ips)

    @property
    @pulumi.getter(name="effectiveOutboundIPs")
    def effective_outbound_ips(self) -> Optional[List['outputs.ResourceReferenceResponse']]:
        """
        The effective outbound IP resources of the cluster load balancer.
        """
        return pulumi.get(self, "effective_outbound_ips")

    @property
    @pulumi.getter(name="managedOutboundIPs")
    def managed_outbound_ips(self) -> Optional['outputs.ManagedClusterLoadBalancerProfileResponseManagedOutboundIPs']:
        """
        Desired managed outbound IPs for the cluster load balancer.
        """
        return pulumi.get(self, "managed_outbound_ips")

    @property
    @pulumi.getter(name="outboundIPPrefixes")
    def outbound_ip_prefixes(self) -> Optional['outputs.ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixes']:
        """
        Desired outbound IP Prefix resources for the cluster load balancer.
        """
        return pulumi.get(self, "outbound_ip_prefixes")

    @property
    @pulumi.getter(name="outboundIPs")
    def outbound_ips(self) -> Optional['outputs.ManagedClusterLoadBalancerProfileResponseOutboundIPs']:
        """
        Desired outbound IP resources for the cluster load balancer.
        """
        return pulumi.get(self, "outbound_ips")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ManagedClusterLoadBalancerProfileResponseManagedOutboundIPs(dict):
    """
    Desired managed outbound IPs for the cluster load balancer.
    """
    def __init__(__self__, *,
                 count: Optional[float] = None):
        """
        Desired managed outbound IPs for the cluster load balancer.
        :param float count: Desired number of outbound IP created/managed by Azure for the cluster load balancer. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 1. 
        """
        if count is not None:
            pulumi.set(__self__, "count", count)

    @property
    @pulumi.getter
    def count(self) -> Optional[float]:
        """
        Desired number of outbound IP created/managed by Azure for the cluster load balancer. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 1. 
        """
        return pulumi.get(self, "count")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ManagedClusterLoadBalancerProfileResponseOutboundIPPrefixes(dict):
    """
    Desired outbound IP Prefix resources for the cluster load balancer.
    """
    def __init__(__self__, *,
                 public_ip_prefixes: Optional[List['outputs.ResourceReferenceResponse']] = None):
        """
        Desired outbound IP Prefix resources for the cluster load balancer.
        :param List['ResourceReferenceResponseArgs'] public_ip_prefixes: A list of public IP prefix resources.
        """
        if public_ip_prefixes is not None:
            pulumi.set(__self__, "public_ip_prefixes", public_ip_prefixes)

    @property
    @pulumi.getter(name="publicIPPrefixes")
    def public_ip_prefixes(self) -> Optional[List['outputs.ResourceReferenceResponse']]:
        """
        A list of public IP prefix resources.
        """
        return pulumi.get(self, "public_ip_prefixes")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ManagedClusterLoadBalancerProfileResponseOutboundIPs(dict):
    """
    Desired outbound IP resources for the cluster load balancer.
    """
    def __init__(__self__, *,
                 public_ips: Optional[List['outputs.ResourceReferenceResponse']] = None):
        """
        Desired outbound IP resources for the cluster load balancer.
        :param List['ResourceReferenceResponseArgs'] public_ips: A list of public IP resources.
        """
        if public_ips is not None:
            pulumi.set(__self__, "public_ips", public_ips)

    @property
    @pulumi.getter(name="publicIPs")
    def public_ips(self) -> Optional[List['outputs.ResourceReferenceResponse']]:
        """
        A list of public IP resources.
        """
        return pulumi.get(self, "public_ips")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ManagedClusterServicePrincipalProfileResponse(dict):
    """
    Information about a service principal identity for the cluster to use for manipulating Azure APIs.
    """
    def __init__(__self__, *,
                 client_id: str,
                 secret: Optional[str] = None):
        """
        Information about a service principal identity for the cluster to use for manipulating Azure APIs.
        :param str client_id: The ID for the service principal.
        :param str secret: The secret password associated with the service principal in plain text.
        """
        pulumi.set(__self__, "client_id", client_id)
        if secret is not None:
            pulumi.set(__self__, "secret", secret)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> str:
        """
        The ID for the service principal.
        """
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter
    def secret(self) -> Optional[str]:
        """
        The secret password associated with the service principal in plain text.
        """
        return pulumi.get(self, "secret")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ManagedClusterWindowsProfileResponse(dict):
    """
    Profile for Windows VMs in the container service cluster.
    """
    def __init__(__self__, *,
                 admin_username: str,
                 admin_password: Optional[str] = None):
        """
        Profile for Windows VMs in the container service cluster.
        :param str admin_username: The administrator username to use for Windows VMs.
        :param str admin_password: The administrator password to use for Windows VMs.
        """
        pulumi.set(__self__, "admin_username", admin_username)
        if admin_password is not None:
            pulumi.set(__self__, "admin_password", admin_password)

    @property
    @pulumi.getter(name="adminUsername")
    def admin_username(self) -> str:
        """
        The administrator username to use for Windows VMs.
        """
        return pulumi.get(self, "admin_username")

    @property
    @pulumi.getter(name="adminPassword")
    def admin_password(self) -> Optional[str]:
        """
        The administrator password to use for Windows VMs.
        """
        return pulumi.get(self, "admin_password")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ResourceReferenceResponse(dict):
    """
    A reference to an Azure resource.
    """
    def __init__(__self__, *,
                 id: Optional[str] = None):
        """
        A reference to an Azure resource.
        :param str id: The fully qualified Azure resource id.
        """
        if id is not None:
            pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        The fully qualified Azure resource id.
        """
        return pulumi.get(self, "id")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


