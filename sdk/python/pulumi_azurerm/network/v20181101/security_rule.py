# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class SecurityRule(pulumi.CustomResource):
    etag: pulumi.Output[str]
    """
    A unique read-only string that changes whenever the resource is updated.
    """
    name: pulumi.Output[str]
    """
    The name of the resource that is unique within a resource group. This name can be used to access the resource.
    """
    properties: pulumi.Output[dict]
    """
    Properties of the security rule
      * `access` (`str`) - The network traffic is allowed or denied. Possible values are: 'Allow' and 'Deny'.
      * `description` (`str`) - A description for this rule. Restricted to 140 chars.
      * `destination_address_prefix` (`str`) - The destination address prefix. CIDR or destination IP range. Asterisks '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used.
      * `destination_address_prefixes` (`list`) - The destination address prefixes. CIDR or destination IP ranges.
      * `destination_application_security_groups` (`list`) - The application security group specified as destination.
        * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
        * `id` (`str`) - Resource ID.
        * `location` (`str`) - Resource location.
        * `name` (`str`) - Resource name.
        * `properties` (`dict`) - Properties of the application security group.
          * `provisioning_state` (`str`) - The provisioning state of the application security group resource. Possible values are: 'Succeeded', 'Updating', 'Deleting', and 'Failed'.
          * `resource_guid` (`str`) - The resource GUID property of the application security group resource. It uniquely identifies a resource, even if the user changes its name or migrate the resource across subscriptions or resource groups.

        * `tags` (`dict`) - Resource tags.
        * `type` (`str`) - Resource type.

      * `destination_port_range` (`str`) - The destination port or range. Integer or range between 0 and 65535. Asterisks '*' can also be used to match all ports.
      * `destination_port_ranges` (`list`) - The destination port ranges.
      * `direction` (`str`) - The direction of the rule. The direction specifies if rule will be evaluated on incoming or outgoing traffic. Possible values are: 'Inbound' and 'Outbound'.
      * `priority` (`float`) - The priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
      * `protocol` (`str`) - Network protocol this rule applies to. Possible values are 'Tcp', 'Udp', and '*'.
      * `provisioning_state` (`str`) - The provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
      * `source_address_prefix` (`str`) - The CIDR or source IP range. Asterisks '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used. If this is an ingress rule, specifies where network traffic originates from. 
      * `source_address_prefixes` (`list`) - The CIDR or source IP ranges.
      * `source_application_security_groups` (`list`) - The application security group specified as source.
      * `source_port_range` (`str`) - The source port or range. Integer or range between 0 and 65535. Asterisks '*' can also be used to match all ports.
      * `source_port_ranges` (`list`) - The source port ranges.
    """
    def __init__(__self__, resource_name, opts=None, etag=None, id=None, name=None, network_security_group_name=None, properties=None, resource_group_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Network security rule.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] etag: A unique read-only string that changes whenever the resource is updated.
        :param pulumi.Input[str] id: Resource ID.
        :param pulumi.Input[str] name: The name of the security rule.
        :param pulumi.Input[str] network_security_group_name: The name of the network security group.
        :param pulumi.Input[dict] properties: Properties of the security rule
        :param pulumi.Input[str] resource_group_name: The name of the resource group.

        The **properties** object supports the following:

          * `access` (`pulumi.Input[str]`) - The network traffic is allowed or denied. Possible values are: 'Allow' and 'Deny'.
          * `description` (`pulumi.Input[str]`) - A description for this rule. Restricted to 140 chars.
          * `destination_address_prefix` (`pulumi.Input[str]`) - The destination address prefix. CIDR or destination IP range. Asterisks '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used.
          * `destination_address_prefixes` (`pulumi.Input[list]`) - The destination address prefixes. CIDR or destination IP ranges.
          * `destination_application_security_groups` (`pulumi.Input[list]`) - The application security group specified as destination.
            * `id` (`pulumi.Input[str]`) - Resource ID.
            * `location` (`pulumi.Input[str]`) - Resource location.
            * `properties` (`pulumi.Input[dict]`) - Properties of the application security group.
            * `tags` (`pulumi.Input[dict]`) - Resource tags.

          * `destination_port_range` (`pulumi.Input[str]`) - The destination port or range. Integer or range between 0 and 65535. Asterisks '*' can also be used to match all ports.
          * `destination_port_ranges` (`pulumi.Input[list]`) - The destination port ranges.
          * `direction` (`pulumi.Input[str]`) - The direction of the rule. The direction specifies if rule will be evaluated on incoming or outgoing traffic. Possible values are: 'Inbound' and 'Outbound'.
          * `priority` (`pulumi.Input[float]`) - The priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
          * `protocol` (`pulumi.Input[str]`) - Network protocol this rule applies to. Possible values are 'Tcp', 'Udp', and '*'.
          * `provisioning_state` (`pulumi.Input[str]`) - The provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
          * `source_address_prefix` (`pulumi.Input[str]`) - The CIDR or source IP range. Asterisks '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used. If this is an ingress rule, specifies where network traffic originates from. 
          * `source_address_prefixes` (`pulumi.Input[list]`) - The CIDR or source IP ranges.
          * `source_application_security_groups` (`pulumi.Input[list]`) - The application security group specified as source.
          * `source_port_range` (`pulumi.Input[str]`) - The source port or range. Integer or range between 0 and 65535. Asterisks '*' can also be used to match all ports.
          * `source_port_ranges` (`pulumi.Input[list]`) - The source port ranges.
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

            __props__['etag'] = etag
            __props__['id'] = id
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if network_security_group_name is None:
                raise TypeError("Missing required property 'network_security_group_name'")
            __props__['network_security_group_name'] = network_security_group_name
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
        super(SecurityRule, __self__).__init__(
            'azurerm:network/v20181101:SecurityRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing SecurityRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return SecurityRule(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
