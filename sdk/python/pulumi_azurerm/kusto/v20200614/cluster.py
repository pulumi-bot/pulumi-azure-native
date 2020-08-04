# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class Cluster(pulumi.CustomResource):
    identity: pulumi.Output[dict]
    """
    The identity of the cluster, if configured.
      * `principal_id` (`str`) - The principal ID of resource identity.
      * `tenant_id` (`str`) - The tenant ID of resource.
      * `type` (`str`) - The identity type.
      * `user_assigned_identities` (`dict`) - The list of user identities associated with the Kusto cluster. The user identity dictionary key references will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'.
    """
    location: pulumi.Output[str]
    """
    The geo-location where the resource lives
    """
    name: pulumi.Output[str]
    """
    The name of the resource
    """
    properties: pulumi.Output[dict]
    """
    The cluster properties.
      * `data_ingestion_uri` (`str`) - The cluster data ingestion URI.
      * `enable_disk_encryption` (`bool`) - A boolean value that indicates if the cluster's disks are encrypted.
      * `enable_double_encryption` (`bool`) - A boolean value that indicates if double encryption is enabled.
      * `enable_purge` (`bool`) - A boolean value that indicates if the purge operations are enabled.
      * `enable_streaming_ingest` (`bool`) - A boolean value that indicates if the streaming ingest is enabled.
      * `key_vault_properties` (`dict`) - KeyVault properties for the cluster encryption.
        * `key_name` (`str`) - The name of the key vault key.
        * `key_vault_uri` (`str`) - The Uri of the key vault.
        * `key_version` (`str`) - The version of the key vault key.

      * `language_extensions` (`dict`) - List of the cluster's language extensions.
        * `value` (`list`) - The list of language extensions.
          * `language_extension_name` (`str`) - The language extension name.

      * `optimized_autoscale` (`dict`) - Optimized auto scale definition.
        * `is_enabled` (`bool`) - A boolean value that indicate if the optimized autoscale feature is enabled or not.
        * `maximum` (`float`) - Maximum allowed instances count.
        * `minimum` (`float`) - Minimum allowed instances count.
        * `version` (`float`) - The version of the template defined, for instance 1.

      * `provisioning_state` (`str`) - The provisioned state of the resource.
      * `state` (`str`) - The state of the resource.
      * `state_reason` (`str`) - The reason for the cluster's current state.
      * `trusted_external_tenants` (`list`) - The cluster's external tenants.
        * `value` (`str`) - GUID representing an external tenant.

      * `uri` (`str`) - The cluster URI.
      * `virtual_network_configuration` (`dict`) - Virtual network definition.
        * `data_management_public_ip_id` (`str`) - Data management's service public IP address resource id.
        * `engine_public_ip_id` (`str`) - Engine service's public IP address resource id.
        * `subnet_id` (`str`) - The subnet resource id.
    """
    sku: pulumi.Output[dict]
    """
    The SKU of the cluster.
      * `capacity` (`float`) - The number of instances of the cluster.
      * `name` (`str`) - SKU name.
      * `tier` (`str`) - SKU tier.
    """
    tags: pulumi.Output[dict]
    """
    Resource tags.
    """
    type: pulumi.Output[str]
    """
    The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
    """
    zones: pulumi.Output[dict]
    """
    The availability zones of the cluster.
    """
    def __init__(__self__, resource_name, opts=None, identity=None, location=None, name=None, properties=None, resource_group_name=None, sku=None, tags=None, zones=None, __props__=None, __name__=None, __opts__=None):
        """
        Class representing a Kusto cluster.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] identity: The identity of the cluster, if configured.
        :param pulumi.Input[str] location: The geo-location where the resource lives
        :param pulumi.Input[str] name: The name of the Kusto cluster.
        :param pulumi.Input[dict] properties: The cluster properties.
        :param pulumi.Input[str] resource_group_name: The name of the resource group containing the Kusto cluster.
        :param pulumi.Input[dict] sku: The SKU of the cluster.
        :param pulumi.Input[dict] tags: Resource tags.
        :param pulumi.Input[dict] zones: The availability zones of the cluster.

        The **identity** object supports the following:

          * `type` (`pulumi.Input[str]`) - The identity type.
          * `user_assigned_identities` (`pulumi.Input[dict]`) - The list of user identities associated with the Kusto cluster. The user identity dictionary key references will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'.

        The **properties** object supports the following:

          * `enable_disk_encryption` (`pulumi.Input[bool]`) - A boolean value that indicates if the cluster's disks are encrypted.
          * `enable_double_encryption` (`pulumi.Input[bool]`) - A boolean value that indicates if double encryption is enabled.
          * `enable_purge` (`pulumi.Input[bool]`) - A boolean value that indicates if the purge operations are enabled.
          * `enable_streaming_ingest` (`pulumi.Input[bool]`) - A boolean value that indicates if the streaming ingest is enabled.
          * `key_vault_properties` (`pulumi.Input[dict]`) - KeyVault properties for the cluster encryption.
            * `key_name` (`pulumi.Input[str]`) - The name of the key vault key.
            * `key_vault_uri` (`pulumi.Input[str]`) - The Uri of the key vault.
            * `key_version` (`pulumi.Input[str]`) - The version of the key vault key.

          * `optimized_autoscale` (`pulumi.Input[dict]`) - Optimized auto scale definition.
            * `is_enabled` (`pulumi.Input[bool]`) - A boolean value that indicate if the optimized autoscale feature is enabled or not.
            * `maximum` (`pulumi.Input[float]`) - Maximum allowed instances count.
            * `minimum` (`pulumi.Input[float]`) - Minimum allowed instances count.
            * `version` (`pulumi.Input[float]`) - The version of the template defined, for instance 1.

          * `trusted_external_tenants` (`pulumi.Input[list]`) - The cluster's external tenants.
            * `value` (`pulumi.Input[str]`) - GUID representing an external tenant.

          * `virtual_network_configuration` (`pulumi.Input[dict]`) - Virtual network definition.
            * `data_management_public_ip_id` (`pulumi.Input[str]`) - Data management's service public IP address resource id.
            * `engine_public_ip_id` (`pulumi.Input[str]`) - Engine service's public IP address resource id.
            * `subnet_id` (`pulumi.Input[str]`) - The subnet resource id.

        The **sku** object supports the following:

          * `capacity` (`pulumi.Input[float]`) - The number of instances of the cluster.
          * `name` (`pulumi.Input[str]`) - SKU name.
          * `tier` (`pulumi.Input[str]`) - SKU tier.
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

            __props__['identity'] = identity
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if sku is None:
                raise TypeError("Missing required property 'sku'")
            __props__['sku'] = sku
            __props__['tags'] = tags
            __props__['zones'] = zones
            __props__['type'] = None
        super(Cluster, __self__).__init__(
            'azurerm:kusto/v20200614:Cluster',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing Cluster resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Cluster(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
