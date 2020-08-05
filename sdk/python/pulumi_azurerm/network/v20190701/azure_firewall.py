# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class AzureFirewall(pulumi.CustomResource):
    etag: pulumi.Output[str]
    """
    A unique read-only string that changes whenever the resource is updated.
    """
    location: pulumi.Output[str]
    """
    Resource location.
    """
    name: pulumi.Output[str]
    """
    Resource name.
    """
    properties: pulumi.Output[dict]
    """
    Properties of the azure firewall.
      * `application_rule_collections` (`list`) - Collection of application rule collections used by Azure Firewall.
        * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
        * `id` (`str`) - Resource ID.
        * `name` (`str`) - The name of the resource that is unique within the Azure firewall. This name can be used to access the resource.
        * `properties` (`dict`) - Properties of the azure firewall application rule collection.
          * `action` (`dict`) - The action type of a rule collection.
            * `type` (`str`) - The type of action.

          * `priority` (`float`) - Priority of the application rule collection resource.
          * `provisioning_state` (`str`) - The provisioning state of the application rule collection resource.
          * `rules` (`list`) - Collection of rules used by a application rule collection.
            * `description` (`str`) - Description of the rule.
            * `fqdn_tags` (`list`) - List of FQDN Tags for this rule.
            * `name` (`str`) - Name of the application rule.
            * `protocols` (`list`) - Array of ApplicationRuleProtocols.
              * `port` (`float`) - Port number for the protocol, cannot be greater than 64000. This field is optional.
              * `protocol_type` (`str`) - Protocol type.

            * `source_addresses` (`list`) - List of source IP addresses for this rule.
            * `target_fqdns` (`list`) - List of FQDNs for this rule.

      * `firewall_policy` (`dict`) - The firewallPolicy associated with this azure firewall.
        * `id` (`str`) - Resource ID.

      * `hub_ip_addresses` (`dict`) - IP addresses associated with AzureFirewall.
        * `private_ip_address` (`str`) - Private IP Address associated with azure firewall.
        * `public_ip_addresses` (`list`) - List of Public IP addresses associated with azure firewall.
          * `address` (`str`) - Public IP Address value.

      * `ip_configurations` (`list`) - IP configuration of the Azure Firewall resource.
        * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
        * `id` (`str`) - Resource ID.
        * `name` (`str`) - Name of the resource that is unique within a resource group. This name can be used to access the resource.
        * `properties` (`dict`) - Properties of the azure firewall IP configuration.
          * `private_ip_address` (`str`) - The Firewall Internal Load Balancer IP to be used as the next hop in User Defined Routes.
          * `provisioning_state` (`str`) - The provisioning state of the Azure firewall IP configuration resource.
          * `public_ip_address` (`dict`) - Reference of the PublicIP resource. This field is a mandatory input if subnet is not null.
          * `subnet` (`dict`) - Reference of the subnet resource. This resource must be named 'AzureFirewallSubnet'.

      * `nat_rule_collections` (`list`) - Collection of NAT rule collections used by Azure Firewall.
        * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
        * `id` (`str`) - Resource ID.
        * `name` (`str`) - The name of the resource that is unique within the Azure firewall. This name can be used to access the resource.
        * `properties` (`dict`) - Properties of the azure firewall NAT rule collection.
          * `action` (`dict`) - The action type of a NAT rule collection.
            * `type` (`str`) - The type of action.

          * `priority` (`float`) - Priority of the NAT rule collection resource.
          * `provisioning_state` (`str`) - The provisioning state of the NAT rule collection resource.
          * `rules` (`list`) - Collection of rules used by a NAT rule collection.
            * `description` (`str`) - Description of the rule.
            * `destination_addresses` (`list`) - List of destination IP addresses for this rule. Supports IP ranges, prefixes, and service tags.
            * `destination_ports` (`list`) - List of destination ports.
            * `name` (`str`) - Name of the NAT rule.
            * `protocols` (`list`) - Array of AzureFirewallNetworkRuleProtocols applicable to this NAT rule.
            * `source_addresses` (`list`) - List of source IP addresses for this rule.
            * `translated_address` (`str`) - The translated address for this NAT rule.
            * `translated_port` (`str`) - The translated port for this NAT rule.

      * `network_rule_collections` (`list`) - Collection of network rule collections used by Azure Firewall.
        * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
        * `id` (`str`) - Resource ID.
        * `name` (`str`) - The name of the resource that is unique within the Azure firewall. This name can be used to access the resource.
        * `properties` (`dict`) - Properties of the azure firewall network rule collection.
          * `action` (`dict`) - The action type of a rule collection.
          * `priority` (`float`) - Priority of the network rule collection resource.
          * `provisioning_state` (`str`) - The provisioning state of the network rule collection resource.
          * `rules` (`list`) - Collection of rules used by a network rule collection.
            * `description` (`str`) - Description of the rule.
            * `destination_addresses` (`list`) - List of destination IP addresses.
            * `destination_ports` (`list`) - List of destination ports.
            * `name` (`str`) - Name of the network rule.
            * `protocols` (`list`) - Array of AzureFirewallNetworkRuleProtocols.
            * `source_addresses` (`list`) - List of source IP addresses for this rule.

      * `provisioning_state` (`str`) - The provisioning state of the Azure firewall resource.
      * `threat_intel_mode` (`str`) - The operation mode for Threat Intelligence.
      * `virtual_hub` (`dict`) - The virtualHub to which the firewall belongs.
    """
    tags: pulumi.Output[dict]
    """
    Resource tags.
    """
    type: pulumi.Output[str]
    """
    Resource type.
    """
    zones: pulumi.Output[list]
    """
    A list of availability zones denoting where the resource needs to come from.
    """
    def __init__(__self__, resource_name, opts=None, application_rule_collections=None, firewall_policy=None, id=None, ip_configurations=None, location=None, name=None, nat_rule_collections=None, network_rule_collections=None, provisioning_state=None, resource_group_name=None, tags=None, threat_intel_mode=None, virtual_hub=None, zones=None, __props__=None, __name__=None, __opts__=None):
        """
        Azure Firewall resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] application_rule_collections: Collection of application rule collections used by Azure Firewall.
        :param pulumi.Input[dict] firewall_policy: The firewallPolicy associated with this azure firewall.
        :param pulumi.Input[str] id: Resource ID.
        :param pulumi.Input[list] ip_configurations: IP configuration of the Azure Firewall resource.
        :param pulumi.Input[str] location: Resource location.
        :param pulumi.Input[str] name: The name of the Azure Firewall.
        :param pulumi.Input[list] nat_rule_collections: Collection of NAT rule collections used by Azure Firewall.
        :param pulumi.Input[list] network_rule_collections: Collection of network rule collections used by Azure Firewall.
        :param pulumi.Input[str] provisioning_state: The provisioning state of the Azure firewall resource.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[dict] tags: Resource tags.
        :param pulumi.Input[str] threat_intel_mode: The operation mode for Threat Intelligence.
        :param pulumi.Input[dict] virtual_hub: The virtualHub to which the firewall belongs.
        :param pulumi.Input[list] zones: A list of availability zones denoting where the resource needs to come from.

        The **application_rule_collections** object supports the following:

          * `action` (`pulumi.Input[dict]`) - The action type of a rule collection.
            * `type` (`pulumi.Input[str]`) - The type of action.

          * `id` (`pulumi.Input[str]`) - Resource ID.
          * `name` (`pulumi.Input[str]`) - The name of the resource that is unique within the Azure firewall. This name can be used to access the resource.
          * `priority` (`pulumi.Input[float]`) - Priority of the application rule collection resource.
          * `provisioning_state` (`pulumi.Input[str]`) - The provisioning state of the application rule collection resource.
          * `rules` (`pulumi.Input[list]`) - Collection of rules used by a application rule collection.
            * `description` (`pulumi.Input[str]`) - Description of the rule.
            * `fqdn_tags` (`pulumi.Input[list]`) - List of FQDN Tags for this rule.
            * `name` (`pulumi.Input[str]`) - Name of the application rule.
            * `protocols` (`pulumi.Input[list]`) - Array of ApplicationRuleProtocols.
              * `port` (`pulumi.Input[float]`) - Port number for the protocol, cannot be greater than 64000. This field is optional.
              * `protocol_type` (`pulumi.Input[str]`) - Protocol type.

            * `source_addresses` (`pulumi.Input[list]`) - List of source IP addresses for this rule.
            * `target_fqdns` (`pulumi.Input[list]`) - List of FQDNs for this rule.

        The **firewall_policy** object supports the following:

          * `id` (`pulumi.Input[str]`) - Resource ID.

        The **ip_configurations** object supports the following:

          * `id` (`pulumi.Input[str]`) - Resource ID.
          * `name` (`pulumi.Input[str]`) - Name of the resource that is unique within a resource group. This name can be used to access the resource.
          * `provisioning_state` (`pulumi.Input[str]`) - The provisioning state of the Azure firewall IP configuration resource.
          * `public_ip_address` (`pulumi.Input[dict]`) - Reference of the PublicIP resource. This field is a mandatory input if subnet is not null.
          * `subnet` (`pulumi.Input[dict]`) - Reference of the subnet resource. This resource must be named 'AzureFirewallSubnet'.

        The **nat_rule_collections** object supports the following:

          * `action` (`pulumi.Input[dict]`) - The action type of a NAT rule collection.
            * `type` (`pulumi.Input[str]`) - The type of action.

          * `id` (`pulumi.Input[str]`) - Resource ID.
          * `name` (`pulumi.Input[str]`) - The name of the resource that is unique within the Azure firewall. This name can be used to access the resource.
          * `priority` (`pulumi.Input[float]`) - Priority of the NAT rule collection resource.
          * `provisioning_state` (`pulumi.Input[str]`) - The provisioning state of the NAT rule collection resource.
          * `rules` (`pulumi.Input[list]`) - Collection of rules used by a NAT rule collection.
            * `description` (`pulumi.Input[str]`) - Description of the rule.
            * `destination_addresses` (`pulumi.Input[list]`) - List of destination IP addresses for this rule. Supports IP ranges, prefixes, and service tags.
            * `destination_ports` (`pulumi.Input[list]`) - List of destination ports.
            * `name` (`pulumi.Input[str]`) - Name of the NAT rule.
            * `protocols` (`pulumi.Input[list]`) - Array of AzureFirewallNetworkRuleProtocols applicable to this NAT rule.
            * `source_addresses` (`pulumi.Input[list]`) - List of source IP addresses for this rule.
            * `translated_address` (`pulumi.Input[str]`) - The translated address for this NAT rule.
            * `translated_port` (`pulumi.Input[str]`) - The translated port for this NAT rule.

        The **network_rule_collections** object supports the following:

          * `action` (`pulumi.Input[dict]`) - The action type of a rule collection.
          * `id` (`pulumi.Input[str]`) - Resource ID.
          * `name` (`pulumi.Input[str]`) - The name of the resource that is unique within the Azure firewall. This name can be used to access the resource.
          * `priority` (`pulumi.Input[float]`) - Priority of the network rule collection resource.
          * `provisioning_state` (`pulumi.Input[str]`) - The provisioning state of the network rule collection resource.
          * `rules` (`pulumi.Input[list]`) - Collection of rules used by a network rule collection.
            * `description` (`pulumi.Input[str]`) - Description of the rule.
            * `destination_addresses` (`pulumi.Input[list]`) - List of destination IP addresses.
            * `destination_ports` (`pulumi.Input[list]`) - List of destination ports.
            * `name` (`pulumi.Input[str]`) - Name of the network rule.
            * `protocols` (`pulumi.Input[list]`) - Array of AzureFirewallNetworkRuleProtocols.
            * `source_addresses` (`pulumi.Input[list]`) - List of source IP addresses for this rule.
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
            __props__['firewall_policy'] = firewall_policy
            __props__['id'] = id
            __props__['ip_configurations'] = ip_configurations
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['nat_rule_collections'] = nat_rule_collections
            __props__['network_rule_collections'] = network_rule_collections
            __props__['provisioning_state'] = provisioning_state
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['threat_intel_mode'] = threat_intel_mode
            __props__['virtual_hub'] = virtual_hub
            __props__['zones'] = zones
            __props__['etag'] = None
            __props__['properties'] = None
            __props__['type'] = None
        super(AzureFirewall, __self__).__init__(
            'azurerm:network/v20190701:AzureFirewall',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing AzureFirewall resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return AzureFirewall(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
