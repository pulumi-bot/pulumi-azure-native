# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class ManagedCluster(pulumi.CustomResource):
    identity: pulumi.Output[dict]
    """
    The identity of the managed cluster, if configured.
      * `principal_id` (`str`) - The principal id of the system assigned identity which is used by master components.
      * `tenant_id` (`str`) - The tenant id of the system assigned identity which is used by master components.
      * `type` (`str`) - The type of identity used for the managed cluster. Type 'SystemAssigned' will use an implicitly created identity in master components and an auto-created user assigned identity in MC_ resource group in agent nodes. Type 'None' will not use MSI for the managed cluster, service principal will be used instead.
    """
    location: pulumi.Output[str]
    """
    Resource location
    """
    name: pulumi.Output[str]
    """
    Resource name
    """
    properties: pulumi.Output[dict]
    """
    Properties of a managed cluster.
      * `aad_profile` (`dict`) - Profile of Azure Active Directory configuration.
        * `client_app_id` (`str`) - The client AAD application ID.
        * `server_app_id` (`str`) - The server AAD application ID.
        * `server_app_secret` (`str`) - The server AAD application secret.
        * `tenant_id` (`str`) - The AAD tenant ID to use for authentication. If not specified, will use the tenant of the deployment subscription.

      * `addon_profiles` (`dict`) - Profile of managed cluster add-on.
      * `agent_pool_profiles` (`list`) - Properties of the agent pool.
        * `availability_zones` (`list`) - Availability zones for nodes. Must use VirtualMachineScaleSets AgentPoolType.
        * `count` (`float`) - Number of agents (VMs) to host docker containers. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 1.
        * `enable_auto_scaling` (`bool`) - Whether to enable auto-scaler
        * `enable_node_public_ip` (`bool`) - Enable public IP for nodes
        * `max_count` (`float`) - Maximum number of nodes for auto-scaling
        * `max_pods` (`float`) - Maximum number of pods that can run on a node.
        * `min_count` (`float`) - Minimum number of nodes for auto-scaling
        * `name` (`str`) - Unique name of the agent pool profile in the context of the subscription and resource group.
        * `node_labels` (`dict`) - Agent pool node labels to be persisted across all nodes in agent pool.
        * `node_taints` (`list`) - Taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.
        * `orchestrator_version` (`str`) - Version of orchestrator specified when creating the managed cluster.
        * `os_disk_size_gb` (`float`) - OS Disk Size in GB to be used to specify the disk size for every machine in this master/agent pool. If you specify 0, it will apply the default osDisk size according to the vmSize specified.
        * `os_type` (`str`) - OsType to be used to specify os type. Choose from Linux and Windows. Default to Linux.
        * `provisioning_state` (`str`) - The current deployment or provisioning state, which only appears in the response.
        * `scale_set_eviction_policy` (`str`) - ScaleSetEvictionPolicy to be used to specify eviction policy for low priority virtual machine scale set. Default to Delete.
        * `scale_set_priority` (`str`) - ScaleSetPriority to be used to specify virtual machine scale set priority. Default to regular.
        * `tags` (`dict`) - Agent pool tags to be persisted on the agent pool virtual machine scale set.
        * `type` (`str`) - AgentPoolType represents types of an agent pool
        * `vm_size` (`str`) - Size of agent VMs.
        * `vnet_subnet_id` (`str`) - VNet SubnetID specifies the VNet's subnet identifier.

      * `api_server_access_profile` (`dict`) - Access profile for managed cluster API server.
        * `authorized_ip_ranges` (`list`) - Authorized IP Ranges to kubernetes API server.
        * `enable_private_cluster` (`bool`) - Whether to create the cluster as a private cluster or not.

      * `disk_encryption_set_id` (`str`) - ResourceId of the disk encryption set to use for enabling encryption at rest.
      * `dns_prefix` (`str`) - DNS prefix specified when creating the managed cluster.
      * `enable_pod_security_policy` (`bool`) - (PREVIEW) Whether to enable Kubernetes Pod security policy.
      * `enable_rbac` (`bool`) - Whether to enable Kubernetes Role-Based Access Control.
      * `fqdn` (`str`) - FQDN for the master pool.
      * `identity_profile` (`dict`) - Identities associated with the cluster.
      * `kubernetes_version` (`str`) - Version of Kubernetes specified when creating the managed cluster.
      * `linux_profile` (`dict`) - Profile for Linux VMs in the container service cluster.
        * `admin_username` (`str`) - The administrator username to use for Linux VMs.
        * `ssh` (`dict`) - SSH configuration for Linux-based VMs running on Azure.
          * `public_keys` (`list`) - The list of SSH public keys used to authenticate with Linux-based VMs. Only expect one key specified.
            * `key_data` (`str`) - Certificate public key used to authenticate with VMs through SSH. The certificate must be in PEM format with or without headers.

      * `max_agent_pools` (`float`) - The max number of agent pools for the managed cluster.
      * `network_profile` (`dict`) - Profile of network configuration.
        * `dns_service_ip` (`str`) - An IP address assigned to the Kubernetes DNS service. It must be within the Kubernetes service address range specified in serviceCidr.
        * `docker_bridge_cidr` (`str`) - A CIDR notation IP range assigned to the Docker bridge network. It must not overlap with any Subnet IP ranges or the Kubernetes service address range.
        * `load_balancer_profile` (`dict`) - Profile of the cluster load balancer.
          * `allocated_outbound_ports` (`float`) - Desired number of allocated SNAT ports per VM. Allowed values must be in the range of 0 to 64000 (inclusive). The default value is 0 which results in Azure dynamically allocating ports.
          * `effective_outbound_i_ps` (`list`) - The effective outbound IP resources of the cluster load balancer.
            * `id` (`str`) - The fully qualified Azure resource id.

          * `idle_timeout_in_minutes` (`float`) - Desired outbound flow idle timeout in minutes. Allowed values must be in the range of 4 to 120 (inclusive). The default value is 30 minutes.
          * `managed_outbound_i_ps` (`dict`) - Desired managed outbound IPs for the cluster load balancer.
            * `count` (`float`) - Desired number of outbound IP created/managed by Azure for the cluster load balancer. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 1. 

          * `outbound_ip_prefixes` (`dict`) - Desired outbound IP Prefix resources for the cluster load balancer.
          * `outbound_i_ps` (`dict`) - Desired outbound IP resources for the cluster load balancer.

        * `load_balancer_sku` (`str`) - The load balancer sku for the managed cluster.
        * `network_plugin` (`str`) - Network plugin used for building Kubernetes network.
        * `network_policy` (`str`) - Network policy used for building Kubernetes network.
        * `outbound_type` (`str`) - The outbound (egress) routing method.
        * `pod_cidr` (`str`) - A CIDR notation IP range from which to assign pod IPs when kubenet is used.
        * `service_cidr` (`str`) - A CIDR notation IP range from which to assign service cluster IPs. It must not overlap with any Subnet IP ranges.

      * `node_resource_group` (`str`) - Name of the resource group containing agent pool nodes.
      * `private_fqdn` (`str`) - FQDN of private cluster.
      * `provisioning_state` (`str`) - The current deployment or provisioning state, which only appears in the response.
      * `service_principal_profile` (`dict`) - Information about a service principal identity for the cluster to use for manipulating Azure APIs.
        * `client_id` (`str`) - The ID for the service principal.
        * `secret` (`str`) - The secret password associated with the service principal in plain text.

      * `windows_profile` (`dict`) - Profile for Windows VMs in the container service cluster.
        * `admin_password` (`str`) - The administrator password to use for Windows VMs.
        * `admin_username` (`str`) - The administrator username to use for Windows VMs.
    """
    tags: pulumi.Output[dict]
    """
    Resource tags
    """
    type: pulumi.Output[str]
    """
    Resource type
    """
    def __init__(__self__, resource_name, opts=None, aad_profile=None, addon_profiles=None, agent_pool_profiles=None, api_server_access_profile=None, disk_encryption_set_id=None, dns_prefix=None, enable_pod_security_policy=None, enable_rbac=None, identity=None, identity_profile=None, kubernetes_version=None, linux_profile=None, location=None, name=None, network_profile=None, node_resource_group=None, resource_group_name=None, service_principal_profile=None, tags=None, windows_profile=None, __props__=None, __name__=None, __opts__=None):
        """
        Managed cluster.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] aad_profile: Profile of Azure Active Directory configuration.
        :param pulumi.Input[dict] addon_profiles: Profile of managed cluster add-on.
        :param pulumi.Input[list] agent_pool_profiles: Properties of the agent pool.
        :param pulumi.Input[dict] api_server_access_profile: Access profile for managed cluster API server.
        :param pulumi.Input[str] disk_encryption_set_id: ResourceId of the disk encryption set to use for enabling encryption at rest.
        :param pulumi.Input[str] dns_prefix: DNS prefix specified when creating the managed cluster.
        :param pulumi.Input[bool] enable_pod_security_policy: (PREVIEW) Whether to enable Kubernetes Pod security policy.
        :param pulumi.Input[bool] enable_rbac: Whether to enable Kubernetes Role-Based Access Control.
        :param pulumi.Input[dict] identity: The identity of the managed cluster, if configured.
        :param pulumi.Input[dict] identity_profile: Identities associated with the cluster.
        :param pulumi.Input[str] kubernetes_version: Version of Kubernetes specified when creating the managed cluster.
        :param pulumi.Input[dict] linux_profile: Profile for Linux VMs in the container service cluster.
        :param pulumi.Input[str] location: Resource location
        :param pulumi.Input[str] name: The name of the managed cluster resource.
        :param pulumi.Input[dict] network_profile: Profile of network configuration.
        :param pulumi.Input[str] node_resource_group: Name of the resource group containing agent pool nodes.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[dict] service_principal_profile: Information about a service principal identity for the cluster to use for manipulating Azure APIs.
        :param pulumi.Input[dict] tags: Resource tags
        :param pulumi.Input[dict] windows_profile: Profile for Windows VMs in the container service cluster.

        The **aad_profile** object supports the following:

          * `client_app_id` (`pulumi.Input[str]`) - The client AAD application ID.
          * `server_app_id` (`pulumi.Input[str]`) - The server AAD application ID.
          * `server_app_secret` (`pulumi.Input[str]`) - The server AAD application secret.
          * `tenant_id` (`pulumi.Input[str]`) - The AAD tenant ID to use for authentication. If not specified, will use the tenant of the deployment subscription.

        The **agent_pool_profiles** object supports the following:

          * `availability_zones` (`pulumi.Input[list]`) - Availability zones for nodes. Must use VirtualMachineScaleSets AgentPoolType.
          * `count` (`pulumi.Input[float]`) - Number of agents (VMs) to host docker containers. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 1.
          * `enable_auto_scaling` (`pulumi.Input[bool]`) - Whether to enable auto-scaler
          * `enable_node_public_ip` (`pulumi.Input[bool]`) - Enable public IP for nodes
          * `max_count` (`pulumi.Input[float]`) - Maximum number of nodes for auto-scaling
          * `max_pods` (`pulumi.Input[float]`) - Maximum number of pods that can run on a node.
          * `min_count` (`pulumi.Input[float]`) - Minimum number of nodes for auto-scaling
          * `name` (`pulumi.Input[str]`) - Unique name of the agent pool profile in the context of the subscription and resource group.
          * `node_labels` (`pulumi.Input[dict]`) - Agent pool node labels to be persisted across all nodes in agent pool.
          * `node_taints` (`pulumi.Input[list]`) - Taints added to new nodes during node pool create and scale. For example, key=value:NoSchedule.
          * `orchestrator_version` (`pulumi.Input[str]`) - Version of orchestrator specified when creating the managed cluster.
          * `os_disk_size_gb` (`pulumi.Input[float]`) - OS Disk Size in GB to be used to specify the disk size for every machine in this master/agent pool. If you specify 0, it will apply the default osDisk size according to the vmSize specified.
          * `os_type` (`pulumi.Input[str]`) - OsType to be used to specify os type. Choose from Linux and Windows. Default to Linux.
          * `scale_set_eviction_policy` (`pulumi.Input[str]`) - ScaleSetEvictionPolicy to be used to specify eviction policy for low priority virtual machine scale set. Default to Delete.
          * `scale_set_priority` (`pulumi.Input[str]`) - ScaleSetPriority to be used to specify virtual machine scale set priority. Default to regular.
          * `tags` (`pulumi.Input[dict]`) - Agent pool tags to be persisted on the agent pool virtual machine scale set.
          * `type` (`pulumi.Input[str]`) - AgentPoolType represents types of an agent pool
          * `vm_size` (`pulumi.Input[str]`) - Size of agent VMs.
          * `vnet_subnet_id` (`pulumi.Input[str]`) - VNet SubnetID specifies the VNet's subnet identifier.

        The **api_server_access_profile** object supports the following:

          * `authorized_ip_ranges` (`pulumi.Input[list]`) - Authorized IP Ranges to kubernetes API server.
          * `enable_private_cluster` (`pulumi.Input[bool]`) - Whether to create the cluster as a private cluster or not.

        The **identity** object supports the following:

          * `type` (`pulumi.Input[str]`) - The type of identity used for the managed cluster. Type 'SystemAssigned' will use an implicitly created identity in master components and an auto-created user assigned identity in MC_ resource group in agent nodes. Type 'None' will not use MSI for the managed cluster, service principal will be used instead.

        The **linux_profile** object supports the following:

          * `admin_username` (`pulumi.Input[str]`) - The administrator username to use for Linux VMs.
          * `ssh` (`pulumi.Input[dict]`) - SSH configuration for Linux-based VMs running on Azure.
            * `public_keys` (`pulumi.Input[list]`) - The list of SSH public keys used to authenticate with Linux-based VMs. Only expect one key specified.
              * `key_data` (`pulumi.Input[str]`) - Certificate public key used to authenticate with VMs through SSH. The certificate must be in PEM format with or without headers.

        The **network_profile** object supports the following:

          * `dns_service_ip` (`pulumi.Input[str]`) - An IP address assigned to the Kubernetes DNS service. It must be within the Kubernetes service address range specified in serviceCidr.
          * `docker_bridge_cidr` (`pulumi.Input[str]`) - A CIDR notation IP range assigned to the Docker bridge network. It must not overlap with any Subnet IP ranges or the Kubernetes service address range.
          * `load_balancer_profile` (`pulumi.Input[dict]`) - Profile of the cluster load balancer.
            * `allocated_outbound_ports` (`pulumi.Input[float]`) - Desired number of allocated SNAT ports per VM. Allowed values must be in the range of 0 to 64000 (inclusive). The default value is 0 which results in Azure dynamically allocating ports.
            * `effective_outbound_i_ps` (`pulumi.Input[list]`) - The effective outbound IP resources of the cluster load balancer.
              * `id` (`pulumi.Input[str]`) - The fully qualified Azure resource id.

            * `idle_timeout_in_minutes` (`pulumi.Input[float]`) - Desired outbound flow idle timeout in minutes. Allowed values must be in the range of 4 to 120 (inclusive). The default value is 30 minutes.
            * `managed_outbound_i_ps` (`pulumi.Input[dict]`) - Desired managed outbound IPs for the cluster load balancer.
              * `count` (`pulumi.Input[float]`) - Desired number of outbound IP created/managed by Azure for the cluster load balancer. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 1. 

            * `outbound_ip_prefixes` (`pulumi.Input[dict]`) - Desired outbound IP Prefix resources for the cluster load balancer.
            * `outbound_i_ps` (`pulumi.Input[dict]`) - Desired outbound IP resources for the cluster load balancer.

          * `load_balancer_sku` (`pulumi.Input[str]`) - The load balancer sku for the managed cluster.
          * `network_plugin` (`pulumi.Input[str]`) - Network plugin used for building Kubernetes network.
          * `network_policy` (`pulumi.Input[str]`) - Network policy used for building Kubernetes network.
          * `outbound_type` (`pulumi.Input[str]`) - The outbound (egress) routing method.
          * `pod_cidr` (`pulumi.Input[str]`) - A CIDR notation IP range from which to assign pod IPs when kubenet is used.
          * `service_cidr` (`pulumi.Input[str]`) - A CIDR notation IP range from which to assign service cluster IPs. It must not overlap with any Subnet IP ranges.

        The **service_principal_profile** object supports the following:

          * `client_id` (`pulumi.Input[str]`) - The ID for the service principal.
          * `secret` (`pulumi.Input[str]`) - The secret password associated with the service principal in plain text.

        The **windows_profile** object supports the following:

          * `admin_password` (`pulumi.Input[str]`) - The administrator password to use for Windows VMs.
          * `admin_username` (`pulumi.Input[str]`) - The administrator username to use for Windows VMs.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['aad_profile'] = aad_profile
            __props__['addon_profiles'] = addon_profiles
            __props__['agent_pool_profiles'] = agent_pool_profiles
            __props__['api_server_access_profile'] = api_server_access_profile
            __props__['disk_encryption_set_id'] = disk_encryption_set_id
            __props__['dns_prefix'] = dns_prefix
            __props__['enable_pod_security_policy'] = enable_pod_security_policy
            __props__['enable_rbac'] = enable_rbac
            __props__['identity'] = identity
            __props__['identity_profile'] = identity_profile
            __props__['kubernetes_version'] = kubernetes_version
            __props__['linux_profile'] = linux_profile
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['network_profile'] = network_profile
            __props__['node_resource_group'] = node_resource_group
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['service_principal_profile'] = service_principal_profile
            __props__['tags'] = tags
            __props__['windows_profile'] = windows_profile
            __props__['properties'] = None
            __props__['type'] = None
        super(ManagedCluster, __self__).__init__(
            'azurerm:containerservice/v20200101:ManagedCluster',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing ManagedCluster resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ManagedCluster(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
