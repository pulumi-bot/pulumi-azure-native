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
    def __init__(__self__, resource_name, opts=None, access=None, description=None, destination_address_prefix=None, destination_port_range=None, direction=None, etag=None, id=None, name=None, network_security_group_name=None, priority=None, protocol=None, provisioning_state=None, resource_group_name=None, source_address_prefix=None, source_port_range=None, __props__=None, __name__=None, __opts__=None):
        """
        Network security rule.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access: The network traffic is allowed or denied. Possible values are: 'Allow' and 'Deny'.
        :param pulumi.Input[str] description: A description for this rule. Restricted to 140 chars.
        :param pulumi.Input[str] destination_address_prefix: The destination address prefix. CIDR or source IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used.
        :param pulumi.Input[str] destination_port_range: The destination port or range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.
        :param pulumi.Input[str] direction: The direction of the rule. The direction specifies if rule will be evaluated on incoming or outgoing traffic. Possible values are: 'Inbound' and 'Outbound'.
        :param pulumi.Input[str] etag: A unique read-only string that changes whenever the resource is updated.
        :param pulumi.Input[str] id: Resource ID.
        :param pulumi.Input[str] name: The name of the security rule.
        :param pulumi.Input[str] network_security_group_name: The name of the network security group.
        :param pulumi.Input[float] priority: The priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
        :param pulumi.Input[str] protocol: Network protocol this rule applies to. Possible values are 'Tcp', 'Udp', and '*'.
        :param pulumi.Input[str] provisioning_state: The provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[str] source_address_prefix: The CIDR or source IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used. If this is an ingress rule, specifies where network traffic originates from. 
        :param pulumi.Input[str] source_port_range: The source port or range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.
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

            if access is None:
                raise TypeError("Missing required property 'access'")
            __props__['access'] = access
            __props__['description'] = description
            if destination_address_prefix is None:
                raise TypeError("Missing required property 'destination_address_prefix'")
            __props__['destination_address_prefix'] = destination_address_prefix
            __props__['destination_port_range'] = destination_port_range
            if direction is None:
                raise TypeError("Missing required property 'direction'")
            __props__['direction'] = direction
            __props__['etag'] = etag
            __props__['id'] = id
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if network_security_group_name is None:
                raise TypeError("Missing required property 'network_security_group_name'")
            __props__['network_security_group_name'] = network_security_group_name
            __props__['priority'] = priority
            if protocol is None:
                raise TypeError("Missing required property 'protocol'")
            __props__['protocol'] = protocol
            __props__['provisioning_state'] = provisioning_state
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if source_address_prefix is None:
                raise TypeError("Missing required property 'source_address_prefix'")
            __props__['source_address_prefix'] = source_address_prefix
            __props__['source_port_range'] = source_port_range
            __props__['properties'] = None
        super(SecurityRule, __self__).__init__(
            'azurerm:network/v20161201:SecurityRule',
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
