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

__all__ = ['AzureFirewall']


class AzureFirewall(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_rule_collections: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['AzureFirewallApplicationRuleCollectionArgs']]]]] = None,
                 azure_firewall_name: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 ip_configurations: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['AzureFirewallIPConfigurationArgs']]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 nat_rule_collections: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['AzureFirewallNatRuleCollectionArgs']]]]] = None,
                 network_rule_collections: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['AzureFirewallNetworkRuleCollectionArgs']]]]] = None,
                 provisioning_state: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 threat_intel_mode: Optional[pulumi.Input[str]] = None,
                 zones: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Azure Firewall resource.

        ## Example Usage
        ### Create Azure Firewall

        ```python
        import pulumi
        import pulumi_azurerm as azurerm

        azure_firewall = azurerm.network.v20190401.AzureFirewall("azureFirewall",
            application_rule_collections=[{
                "name": "apprulecoll",
            }],
            azure_firewall_name="azurefirewall",
            ip_configurations=[{
                "name": "azureFirewallIpConfiguration",
            }],
            location="West US",
            nat_rule_collections=[{
                "name": "natrulecoll",
            }],
            network_rule_collections=[{
                "name": "netrulecoll",
            }],
            resource_group_name="rg1",
            tags={
                "key1": "value1",
            },
            threat_intel_mode="Alert",
            zones=[])

        ```
        ### Create Azure Firewall With Zones

        ```python
        import pulumi
        import pulumi_azurerm as azurerm

        azure_firewall = azurerm.network.v20190401.AzureFirewall("azureFirewall",
            application_rule_collections=[{
                "name": "apprulecoll",
            }],
            azure_firewall_name="azurefirewall",
            ip_configurations=[{
                "name": "azureFirewallIpConfiguration",
            }],
            location="West US 2",
            nat_rule_collections=[{
                "name": "natrulecoll",
            }],
            network_rule_collections=[{
                "name": "netrulecoll",
            }],
            resource_group_name="rg1",
            tags={
                "key1": "value1",
            },
            threat_intel_mode="Alert",
            zones=[
                "1",
                "2",
                "3",
            ])

        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['AzureFirewallApplicationRuleCollectionArgs']]]] application_rule_collections: Collection of application rule collections used by Azure Firewall.
        :param pulumi.Input[str] azure_firewall_name: The name of the Azure Firewall.
        :param pulumi.Input[str] id: Resource ID.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['AzureFirewallIPConfigurationArgs']]]] ip_configurations: IP configuration of the Azure Firewall resource.
        :param pulumi.Input[str] location: Resource location.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['AzureFirewallNatRuleCollectionArgs']]]] nat_rule_collections: Collection of NAT rule collections used by Azure Firewall.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['AzureFirewallNetworkRuleCollectionArgs']]]] network_rule_collections: Collection of network rule collections used by Azure Firewall.
        :param pulumi.Input[str] provisioning_state: The provisioning state of the resource.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags.
        :param pulumi.Input[str] threat_intel_mode: The operation mode for Threat Intelligence.
        :param pulumi.Input[List[pulumi.Input[str]]] zones: A list of availability zones denoting where the resource needs to come from.
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

            __props__['application_rule_collections'] = application_rule_collections
            if azure_firewall_name is None:
                raise TypeError("Missing required property 'azure_firewall_name'")
            __props__['azure_firewall_name'] = azure_firewall_name
            __props__['id'] = id
            __props__['ip_configurations'] = ip_configurations
            __props__['location'] = location
            __props__['nat_rule_collections'] = nat_rule_collections
            __props__['network_rule_collections'] = network_rule_collections
            __props__['provisioning_state'] = provisioning_state
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['threat_intel_mode'] = threat_intel_mode
            __props__['zones'] = zones
            __props__['etag'] = None
            __props__['name'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:network/latest:AzureFirewall"), pulumi.Alias(type_="azurerm:network/v20180401:AzureFirewall"), pulumi.Alias(type_="azurerm:network/v20180601:AzureFirewall"), pulumi.Alias(type_="azurerm:network/v20180701:AzureFirewall"), pulumi.Alias(type_="azurerm:network/v20180801:AzureFirewall"), pulumi.Alias(type_="azurerm:network/v20181001:AzureFirewall"), pulumi.Alias(type_="azurerm:network/v20181101:AzureFirewall"), pulumi.Alias(type_="azurerm:network/v20181201:AzureFirewall"), pulumi.Alias(type_="azurerm:network/v20190201:AzureFirewall"), pulumi.Alias(type_="azurerm:network/v20190601:AzureFirewall"), pulumi.Alias(type_="azurerm:network/v20190701:AzureFirewall"), pulumi.Alias(type_="azurerm:network/v20190801:AzureFirewall"), pulumi.Alias(type_="azurerm:network/v20190901:AzureFirewall"), pulumi.Alias(type_="azurerm:network/v20191101:AzureFirewall"), pulumi.Alias(type_="azurerm:network/v20191201:AzureFirewall"), pulumi.Alias(type_="azurerm:network/v20200301:AzureFirewall"), pulumi.Alias(type_="azurerm:network/v20200401:AzureFirewall"), pulumi.Alias(type_="azurerm:network/v20200501:AzureFirewall"), pulumi.Alias(type_="azurerm:network/v20200601:AzureFirewall")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(AzureFirewall, __self__).__init__(
            'azurerm:network/v20190401:AzureFirewall',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'AzureFirewall':
        """
        Get an existing AzureFirewall resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return AzureFirewall(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="applicationRuleCollections")
    def application_rule_collections(self) -> pulumi.Output[Optional[List['outputs.AzureFirewallApplicationRuleCollectionResponse']]]:
        """
        Collection of application rule collections used by Azure Firewall.
        """
        return pulumi.get(self, "application_rule_collections")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        Gets a unique read-only string that changes whenever the resource is updated.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="ipConfigurations")
    def ip_configurations(self) -> pulumi.Output[Optional[List['outputs.AzureFirewallIPConfigurationResponse']]]:
        """
        IP configuration of the Azure Firewall resource.
        """
        return pulumi.get(self, "ip_configurations")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[Optional[str]]:
        """
        Resource location.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="natRuleCollections")
    def nat_rule_collections(self) -> pulumi.Output[Optional[List['outputs.AzureFirewallNatRuleCollectionResponse']]]:
        """
        Collection of NAT rule collections used by Azure Firewall.
        """
        return pulumi.get(self, "nat_rule_collections")

    @property
    @pulumi.getter(name="networkRuleCollections")
    def network_rule_collections(self) -> pulumi.Output[Optional[List['outputs.AzureFirewallNetworkRuleCollectionResponse']]]:
        """
        Collection of network rule collections used by Azure Firewall.
        """
        return pulumi.get(self, "network_rule_collections")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[Optional[str]]:
        """
        The provisioning state of the resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="threatIntelMode")
    def threat_intel_mode(self) -> pulumi.Output[Optional[str]]:
        """
        The operation mode for Threat Intelligence.
        """
        return pulumi.get(self, "threat_intel_mode")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Resource type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def zones(self) -> pulumi.Output[Optional[List[str]]]:
        """
        A list of availability zones denoting where the resource needs to come from.
        """
        return pulumi.get(self, "zones")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

