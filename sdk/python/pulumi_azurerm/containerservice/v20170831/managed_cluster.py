# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class ManagedCluster(pulumi.CustomResource):
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
      * `agent_pool_profiles` (`list`) - Properties of the agent pool.
        * `count` (`float`) - Number of agents (VMs) to host docker containers. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 1. 
        * `dns_prefix` (`str`) - DNS prefix to be used to create the FQDN for the agent pool.
        * `fqdn` (`str`) - FQDN for the agent pool.
        * `name` (`str`) - Unique name of the agent pool profile in the context of the subscription and resource group.
        * `os_disk_size_gb` (`float`) - OS Disk Size in GB to be used to specify the disk size for every machine in this master/agent pool. If you specify 0, it will apply the default osDisk size according to the vmSize specified.
        * `os_type` (`str`) - OsType to be used to specify os type. Choose from Linux and Windows. Default to Linux.
        * `ports` (`list`) - Ports number array used to expose on this agent pool. The default opened ports are different based on your choice of orchestrator.
        * `storage_profile` (`str`) - Storage profile specifies what kind of storage used. Choose from StorageAccount and ManagedDisks. Leave it empty, we will choose for you based on the orchestrator choice.
        * `vm_size` (`str`) - Size of agent VMs.
        * `vnet_subnet_id` (`str`) - VNet SubnetID specifies the VNet's subnet identifier.

      * `dns_prefix` (`str`) - DNS prefix specified when creating the managed cluster.
      * `fqdn` (`str`) - FQDN for the master pool.
      * `kubernetes_version` (`str`) - Version of Kubernetes specified when creating the managed cluster.
      * `linux_profile` (`dict`) - Profile for Linux VMs in the container service cluster.
        * `admin_username` (`str`) - The administrator username to use for Linux VMs.
        * `ssh` (`dict`) - SSH configuration for Linux-based VMs running on Azure.
          * `public_keys` (`list`) - The list of SSH public keys used to authenticate with Linux-based VMs. Only expect one key specified.
            * `key_data` (`str`) - Certificate public key used to authenticate with VMs through SSH. The certificate must be in PEM format with or without headers.

      * `provisioning_state` (`str`) - The current deployment or provisioning state, which only appears in the response.
      * `service_principal_profile` (`dict`) - Information about a service principal identity for the cluster to use for manipulating Azure APIs. Either secret or keyVaultSecretRef must be specified.
        * `client_id` (`str`) - The ID for the service principal.
        * `key_vault_secret_ref` (`dict`) - Reference to a secret stored in Azure Key Vault.
          * `secret_name` (`str`) - The secret name.
          * `vault_id` (`str`) - Key vault identifier.
          * `version` (`str`) - The secret version.

        * `secret` (`str`) - The secret password associated with the service principal in plain text.
    """
    tags: pulumi.Output[dict]
    """
    Resource tags
    """
    type: pulumi.Output[str]
    """
    Resource type
    """
    def __init__(__self__, resource_name, opts=None, agent_pool_profiles=None, dns_prefix=None, kubernetes_version=None, linux_profile=None, location=None, name=None, resource_group_name=None, service_principal_profile=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Managed cluster.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] agent_pool_profiles: Properties of the agent pool.
        :param pulumi.Input[str] dns_prefix: DNS prefix specified when creating the managed cluster.
        :param pulumi.Input[str] kubernetes_version: Version of Kubernetes specified when creating the managed cluster.
        :param pulumi.Input[dict] linux_profile: Profile for Linux VMs in the container service cluster.
        :param pulumi.Input[str] location: Resource location
        :param pulumi.Input[str] name: The name of the managed cluster resource.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[dict] service_principal_profile: Information about a service principal identity for the cluster to use for manipulating Azure APIs. Either secret or keyVaultSecretRef must be specified.
        :param pulumi.Input[dict] tags: Resource tags

        The **agent_pool_profiles** object supports the following:

          * `count` (`pulumi.Input[float]`) - Number of agents (VMs) to host docker containers. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 1. 
          * `dns_prefix` (`pulumi.Input[str]`) - DNS prefix to be used to create the FQDN for the agent pool.
          * `name` (`pulumi.Input[str]`) - Unique name of the agent pool profile in the context of the subscription and resource group.
          * `os_disk_size_gb` (`pulumi.Input[float]`) - OS Disk Size in GB to be used to specify the disk size for every machine in this master/agent pool. If you specify 0, it will apply the default osDisk size according to the vmSize specified.
          * `os_type` (`pulumi.Input[str]`) - OsType to be used to specify os type. Choose from Linux and Windows. Default to Linux.
          * `ports` (`pulumi.Input[list]`) - Ports number array used to expose on this agent pool. The default opened ports are different based on your choice of orchestrator.
          * `storage_profile` (`pulumi.Input[str]`) - Storage profile specifies what kind of storage used. Choose from StorageAccount and ManagedDisks. Leave it empty, we will choose for you based on the orchestrator choice.
          * `vm_size` (`pulumi.Input[str]`) - Size of agent VMs.
          * `vnet_subnet_id` (`pulumi.Input[str]`) - VNet SubnetID specifies the VNet's subnet identifier.

        The **linux_profile** object supports the following:

          * `admin_username` (`pulumi.Input[str]`) - The administrator username to use for Linux VMs.
          * `ssh` (`pulumi.Input[dict]`) - SSH configuration for Linux-based VMs running on Azure.
            * `public_keys` (`pulumi.Input[list]`) - The list of SSH public keys used to authenticate with Linux-based VMs. Only expect one key specified.
              * `key_data` (`pulumi.Input[str]`) - Certificate public key used to authenticate with VMs through SSH. The certificate must be in PEM format with or without headers.

        The **service_principal_profile** object supports the following:

          * `client_id` (`pulumi.Input[str]`) - The ID for the service principal.
          * `key_vault_secret_ref` (`pulumi.Input[dict]`) - Reference to a secret stored in Azure Key Vault.
            * `secret_name` (`pulumi.Input[str]`) - The secret name.
            * `vault_id` (`pulumi.Input[str]`) - Key vault identifier.
            * `version` (`pulumi.Input[str]`) - The secret version.

          * `secret` (`pulumi.Input[str]`) - The secret password associated with the service principal in plain text.
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

            __props__['agent_pool_profiles'] = agent_pool_profiles
            __props__['dns_prefix'] = dns_prefix
            __props__['kubernetes_version'] = kubernetes_version
            __props__['linux_profile'] = linux_profile
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['service_principal_profile'] = service_principal_profile
            __props__['tags'] = tags
            __props__['properties'] = None
            __props__['type'] = None
        super(ManagedCluster, __self__).__init__(
            'azurerm:containerservice/v20170831:ManagedCluster',
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
