# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class Registry(pulumi.CustomResource):
    location: pulumi.Output[str]
    """
    The location of the resource. This cannot be changed after the resource is created.
    """
    name: pulumi.Output[str]
    """
    The name of the resource.
    """
    properties: pulumi.Output[dict]
    """
    The properties of the container registry.
      * `admin_user_enabled` (`bool`) - The value that indicates whether the admin user is enabled.
      * `creation_date` (`str`) - The creation date of the container registry in ISO8601 format.
      * `login_server` (`str`) - The URL that can be used to log into the container registry.
      * `network_rule_set` (`dict`) - The network rule set for a container registry.
        * `default_action` (`str`) - The default action of allow or deny when no other rules match.
        * `ip_rules` (`list`) - The IP ACL rules.
          * `action` (`str`) - The action of IP ACL rule.
          * `value` (`str`) - Specifies the IP or IP range in CIDR format. Only IPV4 address is allowed.

        * `virtual_network_rules` (`list`) - The virtual network rules.
          * `action` (`str`) - The action of virtual network rule.
          * `id` (`str`) - Resource ID of a subnet, for example: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{vnetName}/subnets/{subnetName}.

      * `policies` (`dict`) - The policies for a container registry.
        * `quarantine_policy` (`dict`) - The quarantine policy for a container registry.
          * `status` (`str`) - The value that indicates whether the policy is enabled or not.

        * `retention_policy` (`dict`) - The retention policy for a container registry.
          * `days` (`float`) - The number of days to retain an untagged manifest after which it gets purged.
          * `last_updated_time` (`str`) - The timestamp when the policy was last updated.
          * `status` (`str`) - The value that indicates whether the policy is enabled or not.

        * `trust_policy` (`dict`) - The content trust policy for a container registry.
          * `status` (`str`) - The value that indicates whether the policy is enabled or not.
          * `type` (`str`) - The type of trust policy.

      * `provisioning_state` (`str`) - The provisioning state of the container registry at the time the operation was called.
      * `status` (`dict`) - The status of the container registry at the time the operation was called.
        * `display_status` (`str`) - The short label for the status.
        * `message` (`str`) - The detailed message for the status, including alerts and error messages.
        * `timestamp` (`str`) - The timestamp when the status was changed to the current value.

      * `storage_account` (`dict`) - The properties of the storage account for the container registry. Only applicable to Classic SKU.
        * `id` (`str`) - The resource ID of the storage account.
    """
    sku: pulumi.Output[dict]
    """
    The SKU of the container registry.
      * `name` (`str`) - The SKU name of the container registry. Required for registry creation.
      * `tier` (`str`) - The SKU tier based on the SKU name.
    """
    tags: pulumi.Output[dict]
    """
    The tags of the resource.
    """
    type: pulumi.Output[str]
    """
    The type of the resource.
    """
    def __init__(__self__, resource_name, opts=None, admin_user_enabled=None, location=None, name=None, network_rule_set=None, policies=None, resource_group_name=None, sku=None, storage_account=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        An object that represents a container registry.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] admin_user_enabled: The value that indicates whether the admin user is enabled.
        :param pulumi.Input[str] location: The location of the resource. This cannot be changed after the resource is created.
        :param pulumi.Input[str] name: The name of the container registry.
        :param pulumi.Input[dict] network_rule_set: The network rule set for a container registry.
        :param pulumi.Input[dict] policies: The policies for a container registry.
        :param pulumi.Input[str] resource_group_name: The name of the resource group to which the container registry belongs.
        :param pulumi.Input[dict] sku: The SKU of the container registry.
        :param pulumi.Input[dict] storage_account: The properties of the storage account for the container registry. Only applicable to Classic SKU.
        :param pulumi.Input[dict] tags: The tags of the resource.

        The **network_rule_set** object supports the following:

          * `default_action` (`pulumi.Input[str]`) - The default action of allow or deny when no other rules match.
          * `ip_rules` (`pulumi.Input[list]`) - The IP ACL rules.
            * `action` (`pulumi.Input[str]`) - The action of IP ACL rule.
            * `i_p_address_or_range` (`pulumi.Input[str]`) - Specifies the IP or IP range in CIDR format. Only IPV4 address is allowed.

          * `virtual_network_rules` (`pulumi.Input[list]`) - The virtual network rules.
            * `action` (`pulumi.Input[str]`) - The action of virtual network rule.
            * `virtual_network_resource_id` (`pulumi.Input[str]`) - Resource ID of a subnet, for example: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{vnetName}/subnets/{subnetName}.

        The **policies** object supports the following:

          * `quarantine_policy` (`pulumi.Input[dict]`) - The quarantine policy for a container registry.
            * `status` (`pulumi.Input[str]`) - The value that indicates whether the policy is enabled or not.

          * `retention_policy` (`pulumi.Input[dict]`) - The retention policy for a container registry.
            * `days` (`pulumi.Input[float]`) - The number of days to retain an untagged manifest after which it gets purged.
            * `status` (`pulumi.Input[str]`) - The value that indicates whether the policy is enabled or not.

          * `trust_policy` (`pulumi.Input[dict]`) - The content trust policy for a container registry.
            * `status` (`pulumi.Input[str]`) - The value that indicates whether the policy is enabled or not.
            * `type` (`pulumi.Input[str]`) - The type of trust policy.

        The **sku** object supports the following:

          * `name` (`pulumi.Input[str]`) - The SKU name of the container registry. Required for registry creation.

        The **storage_account** object supports the following:

          * `id` (`pulumi.Input[str]`) - The resource ID of the storage account.
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

            __props__['admin_user_enabled'] = admin_user_enabled
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['network_rule_set'] = network_rule_set
            __props__['policies'] = policies
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if sku is None:
                raise TypeError("Missing required property 'sku'")
            __props__['sku'] = sku
            __props__['storage_account'] = storage_account
            __props__['tags'] = tags
            __props__['properties'] = None
            __props__['type'] = None
        super(Registry, __self__).__init__(
            'azurerm:containerregistry/v20190501:Registry',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing Registry resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Registry(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
