# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class Account(pulumi.CustomResource):
    identity: pulumi.Output[dict]
    """
    The Key Vault encryption identity, if any.
      * `principal_id` (`str`) - The principal identifier associated with the encryption.
      * `tenant_id` (`str`) - The tenant identifier associated with the encryption.
      * `type` (`str`) - The type of encryption being used. Currently the only supported type is 'SystemAssigned'.
    """
    location: pulumi.Output[str]
    """
    The resource location.
    """
    name: pulumi.Output[str]
    """
    The resource name.
    """
    properties: pulumi.Output[dict]
    """
    The Data Lake Store account properties.
      * `account_id` (`str`) - The unique identifier associated with this Data Lake Store account.
      * `creation_time` (`str`) - The account creation time.
      * `current_tier` (`str`) - The commitment tier in use for the current month.
      * `default_group` (`str`) - The default owner group for all new folders and files created in the Data Lake Store account.
      * `encryption_config` (`dict`) - The Key Vault encryption configuration.
        * `key_vault_meta_info` (`dict`) - The Key Vault information for connecting to user managed encryption keys.
          * `encryption_key_name` (`str`) - The name of the user managed encryption key.
          * `encryption_key_version` (`str`) - The version of the user managed encryption key.
          * `key_vault_resource_id` (`str`) - The resource identifier for the user managed Key Vault being used to encrypt.

        * `type` (`str`) - The type of encryption configuration being used. Currently the only supported types are 'UserManaged' and 'ServiceManaged'.

      * `encryption_provisioning_state` (`str`) - The current state of encryption provisioning for this Data Lake Store account.
      * `encryption_state` (`str`) - The current state of encryption for this Data Lake Store account.
      * `endpoint` (`str`) - The full CName endpoint for this account.
      * `firewall_allow_azure_ips` (`str`) - The current state of allowing or disallowing IPs originating within Azure through the firewall. If the firewall is disabled, this is not enforced.
      * `firewall_rules` (`list`) - The list of firewall rules associated with this Data Lake Store account.
        * `id` (`str`) - The resource identifier.
        * `name` (`str`) - The resource name.
        * `properties` (`dict`) - The firewall rule properties.
          * `end_ip_address` (`str`) - The end IP address for the firewall rule. This can be either ipv4 or ipv6. Start and End should be in the same protocol.
          * `start_ip_address` (`str`) - The start IP address for the firewall rule. This can be either ipv4 or ipv6. Start and End should be in the same protocol.

        * `type` (`str`) - The resource type.

      * `firewall_state` (`str`) - The current state of the IP address firewall for this Data Lake Store account.
      * `last_modified_time` (`str`) - The account last modified time.
      * `new_tier` (`str`) - The commitment tier to use for next month.
      * `provisioning_state` (`str`) - The provisioning status of the Data Lake Store account.
      * `state` (`str`) - The state of the Data Lake Store account.
      * `trusted_id_provider_state` (`str`) - The current state of the trusted identity provider feature for this Data Lake Store account.
      * `trusted_id_providers` (`list`) - The list of trusted identity providers associated with this Data Lake Store account.
        * `id` (`str`) - The resource identifier.
        * `name` (`str`) - The resource name.
        * `properties` (`dict`) - The trusted identity provider properties.
          * `id_provider` (`str`) - The URL of this trusted identity provider.

        * `type` (`str`) - The resource type.

      * `virtual_network_rules` (`list`) - The list of virtual network rules associated with this Data Lake Store account.
        * `id` (`str`) - The resource identifier.
        * `name` (`str`) - The resource name.
        * `properties` (`dict`) - The virtual network rule properties.
          * `subnet_id` (`str`) - The resource identifier for the subnet.

        * `type` (`str`) - The resource type.
    """
    tags: pulumi.Output[dict]
    """
    The resource tags.
    """
    type: pulumi.Output[str]
    """
    The resource type.
    """
    def __init__(__self__, resource_name, opts=None, default_group=None, encryption_config=None, encryption_state=None, firewall_allow_azure_ips=None, firewall_rules=None, firewall_state=None, identity=None, location=None, name=None, new_tier=None, resource_group_name=None, tags=None, trusted_id_provider_state=None, trusted_id_providers=None, virtual_network_rules=None, __props__=None, __name__=None, __opts__=None):
        """
        Data Lake Store account information.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] default_group: The default owner group for all new folders and files created in the Data Lake Store account.
        :param pulumi.Input[dict] encryption_config: The Key Vault encryption configuration.
        :param pulumi.Input[str] encryption_state: The current state of encryption for this Data Lake Store account.
        :param pulumi.Input[str] firewall_allow_azure_ips: The current state of allowing or disallowing IPs originating within Azure through the firewall. If the firewall is disabled, this is not enforced.
        :param pulumi.Input[list] firewall_rules: The list of firewall rules associated with this Data Lake Store account.
        :param pulumi.Input[str] firewall_state: The current state of the IP address firewall for this Data Lake Store account.
        :param pulumi.Input[dict] identity: The Key Vault encryption identity, if any.
        :param pulumi.Input[str] location: The resource location.
        :param pulumi.Input[str] name: The name of the Data Lake Store account.
        :param pulumi.Input[str] new_tier: The commitment tier to use for next month.
        :param pulumi.Input[str] resource_group_name: The name of the Azure resource group.
        :param pulumi.Input[dict] tags: The resource tags.
        :param pulumi.Input[str] trusted_id_provider_state: The current state of the trusted identity provider feature for this Data Lake Store account.
        :param pulumi.Input[list] trusted_id_providers: The list of trusted identity providers associated with this Data Lake Store account.
        :param pulumi.Input[list] virtual_network_rules: The list of virtual network rules associated with this Data Lake Store account.

        The **encryption_config** object supports the following:

          * `key_vault_meta_info` (`pulumi.Input[dict]`) - The Key Vault information for connecting to user managed encryption keys.
            * `encryption_key_name` (`pulumi.Input[str]`) - The name of the user managed encryption key.
            * `encryption_key_version` (`pulumi.Input[str]`) - The version of the user managed encryption key.
            * `key_vault_resource_id` (`pulumi.Input[str]`) - The resource identifier for the user managed Key Vault being used to encrypt.

          * `type` (`pulumi.Input[str]`) - The type of encryption configuration being used. Currently the only supported types are 'UserManaged' and 'ServiceManaged'.

        The **firewall_rules** object supports the following:

          * `end_ip_address` (`pulumi.Input[str]`) - The end IP address for the firewall rule. This can be either ipv4 or ipv6. Start and End should be in the same protocol.
          * `name` (`pulumi.Input[str]`) - The unique name of the firewall rule to create.
          * `start_ip_address` (`pulumi.Input[str]`) - The start IP address for the firewall rule. This can be either ipv4 or ipv6. Start and End should be in the same protocol.

        The **identity** object supports the following:

          * `type` (`pulumi.Input[str]`) - The type of encryption being used. Currently the only supported type is 'SystemAssigned'.

        The **trusted_id_providers** object supports the following:

          * `id_provider` (`pulumi.Input[str]`) - The URL of this trusted identity provider.
          * `name` (`pulumi.Input[str]`) - The unique name of the trusted identity provider to create.

        The **virtual_network_rules** object supports the following:

          * `name` (`pulumi.Input[str]`) - The unique name of the virtual network rule to create.
          * `subnet_id` (`pulumi.Input[str]`) - The resource identifier for the subnet.
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

            __props__['default_group'] = default_group
            __props__['encryption_config'] = encryption_config
            __props__['encryption_state'] = encryption_state
            __props__['firewall_allow_azure_ips'] = firewall_allow_azure_ips
            __props__['firewall_rules'] = firewall_rules
            __props__['firewall_state'] = firewall_state
            __props__['identity'] = identity
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['new_tier'] = new_tier
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['trusted_id_provider_state'] = trusted_id_provider_state
            __props__['trusted_id_providers'] = trusted_id_providers
            __props__['virtual_network_rules'] = virtual_network_rules
            __props__['properties'] = None
            __props__['type'] = None
        super(Account, __self__).__init__(
            'azurerm:datalakestore/v20161101:Account',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing Account resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Account(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop