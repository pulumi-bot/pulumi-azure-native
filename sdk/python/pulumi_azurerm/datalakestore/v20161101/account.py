# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Account']


class Account(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default_group: Optional[pulumi.Input[str]] = None,
                 encryption_config: Optional[pulumi.Input[pulumi.InputType['EncryptionConfigArgs']]] = None,
                 encryption_state: Optional[pulumi.Input[str]] = None,
                 firewall_allow_azure_ips: Optional[pulumi.Input[str]] = None,
                 firewall_rules: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['CreateFirewallRuleWithAccountParametersArgs']]]]] = None,
                 firewall_state: Optional[pulumi.Input[str]] = None,
                 identity: Optional[pulumi.Input[pulumi.InputType['EncryptionIdentityArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 new_tier: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 trusted_id_provider_state: Optional[pulumi.Input[str]] = None,
                 trusted_id_providers: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['CreateTrustedIdProviderWithAccountParametersArgs']]]]] = None,
                 virtual_network_rules: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['CreateVirtualNetworkRuleWithAccountParametersArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Data Lake Store account information.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] default_group: The default owner group for all new folders and files created in the Data Lake Store account.
        :param pulumi.Input[pulumi.InputType['EncryptionConfigArgs']] encryption_config: The Key Vault encryption configuration.
        :param pulumi.Input[str] encryption_state: The current state of encryption for this Data Lake Store account.
        :param pulumi.Input[str] firewall_allow_azure_ips: The current state of allowing or disallowing IPs originating within Azure through the firewall. If the firewall is disabled, this is not enforced.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['CreateFirewallRuleWithAccountParametersArgs']]]] firewall_rules: The list of firewall rules associated with this Data Lake Store account.
        :param pulumi.Input[str] firewall_state: The current state of the IP address firewall for this Data Lake Store account.
        :param pulumi.Input[pulumi.InputType['EncryptionIdentityArgs']] identity: The Key Vault encryption identity, if any.
        :param pulumi.Input[str] location: The resource location.
        :param pulumi.Input[str] name: The name of the Data Lake Store account.
        :param pulumi.Input[str] new_tier: The commitment tier to use for next month.
        :param pulumi.Input[str] resource_group_name: The name of the Azure resource group.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: The resource tags.
        :param pulumi.Input[str] trusted_id_provider_state: The current state of the trusted identity provider feature for this Data Lake Store account.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['CreateTrustedIdProviderWithAccountParametersArgs']]]] trusted_id_providers: The list of trusted identity providers associated with this Data Lake Store account.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['CreateVirtualNetworkRuleWithAccountParametersArgs']]]] virtual_network_rules: The list of virtual network rules associated with this Data Lake Store account.
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
            __props__['account_id'] = None
            __props__['creation_time'] = None
            __props__['current_tier'] = None
            __props__['encryption_provisioning_state'] = None
            __props__['endpoint'] = None
            __props__['last_modified_time'] = None
            __props__['provisioning_state'] = None
            __props__['state'] = None
            __props__['type'] = None
        super(Account, __self__).__init__(
            'azurerm:datalakestore/v20161101:Account',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Account':
        """
        Get an existing Account resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Account(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> str:
        """
        The unique identifier associated with this Data Lake Store account.
        """
        return pulumi.get(self, "account_id")

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> str:
        """
        The account creation time.
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter(name="currentTier")
    def current_tier(self) -> str:
        """
        The commitment tier in use for the current month.
        """
        return pulumi.get(self, "current_tier")

    @property
    @pulumi.getter(name="defaultGroup")
    def default_group(self) -> str:
        """
        The default owner group for all new folders and files created in the Data Lake Store account.
        """
        return pulumi.get(self, "default_group")

    @property
    @pulumi.getter(name="encryptionConfig")
    def encryption_config(self) -> 'outputs.EncryptionConfigResponse':
        """
        The Key Vault encryption configuration.
        """
        return pulumi.get(self, "encryption_config")

    @property
    @pulumi.getter(name="encryptionProvisioningState")
    def encryption_provisioning_state(self) -> str:
        """
        The current state of encryption provisioning for this Data Lake Store account.
        """
        return pulumi.get(self, "encryption_provisioning_state")

    @property
    @pulumi.getter(name="encryptionState")
    def encryption_state(self) -> str:
        """
        The current state of encryption for this Data Lake Store account.
        """
        return pulumi.get(self, "encryption_state")

    @property
    @pulumi.getter
    def endpoint(self) -> str:
        """
        The full CName endpoint for this account.
        """
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter(name="firewallAllowAzureIps")
    def firewall_allow_azure_ips(self) -> str:
        """
        The current state of allowing or disallowing IPs originating within Azure through the firewall. If the firewall is disabled, this is not enforced.
        """
        return pulumi.get(self, "firewall_allow_azure_ips")

    @property
    @pulumi.getter(name="firewallRules")
    def firewall_rules(self) -> List['outputs.FirewallRuleResponse']:
        """
        The list of firewall rules associated with this Data Lake Store account.
        """
        return pulumi.get(self, "firewall_rules")

    @property
    @pulumi.getter(name="firewallState")
    def firewall_state(self) -> str:
        """
        The current state of the IP address firewall for this Data Lake Store account.
        """
        return pulumi.get(self, "firewall_state")

    @property
    @pulumi.getter
    def identity(self) -> 'outputs.EncryptionIdentityResponse':
        """
        The Key Vault encryption identity, if any.
        """
        return pulumi.get(self, "identity")

    @property
    @pulumi.getter(name="lastModifiedTime")
    def last_modified_time(self) -> str:
        """
        The account last modified time.
        """
        return pulumi.get(self, "last_modified_time")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        The resource location.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="newTier")
    def new_tier(self) -> str:
        """
        The commitment tier to use for next month.
        """
        return pulumi.get(self, "new_tier")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        The provisioning status of the Data Lake Store account.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The state of the Data Lake Store account.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        """
        The resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="trustedIdProviderState")
    def trusted_id_provider_state(self) -> str:
        """
        The current state of the trusted identity provider feature for this Data Lake Store account.
        """
        return pulumi.get(self, "trusted_id_provider_state")

    @property
    @pulumi.getter(name="trustedIdProviders")
    def trusted_id_providers(self) -> List['outputs.TrustedIdProviderResponse']:
        """
        The list of trusted identity providers associated with this Data Lake Store account.
        """
        return pulumi.get(self, "trusted_id_providers")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The resource type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="virtualNetworkRules")
    def virtual_network_rules(self) -> List['outputs.VirtualNetworkRuleResponse']:
        """
        The list of virtual network rules associated with this Data Lake Store account.
        """
        return pulumi.get(self, "virtual_network_rules")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

