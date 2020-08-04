# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class Subnet(pulumi.CustomResource):
    etag: pulumi.Output[str]
    """
    A unique read-only string that changes whenever the resource is updated
    """
    name: pulumi.Output[str]
    """
    Gets or sets the name of the resource that is unique within a resource group. This name can be used to access the resource
    """
    properties: pulumi.Output[dict]
    def __init__(__self__, resource_name, opts=None, etag=None, id=None, name=None, properties=None, resource_group_name=None, virtual_network_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Subnet in a VirtualNetwork resource

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] etag: A unique read-only string that changes whenever the resource is updated
        :param pulumi.Input[str] id: Resource Id
        :param pulumi.Input[str] name: The name of the subnet.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[str] virtual_network_name: The name of the virtual network.

        The **properties** object supports the following:

          * `address_prefix` (`pulumi.Input[str]`) - Gets or sets Address prefix for the subnet.
          * `network_security_group` (`pulumi.Input[dict]`) - Gets or sets the reference of the NetworkSecurityGroup resource
            * `etag` (`pulumi.Input[str]`) - Gets a unique read-only string that changes whenever the resource is updated
            * `id` (`pulumi.Input[str]`) - Resource Id
            * `location` (`pulumi.Input[str]`) - Resource location
            * `properties` (`pulumi.Input[dict]`) - Network Security Group resource
              * `default_security_rules` (`pulumi.Input[list]`) - Gets or default security rules of network security group
                * `etag` (`pulumi.Input[str]`) - A unique read-only string that changes whenever the resource is updated
                * `id` (`pulumi.Input[str]`) - Resource Id
                * `name` (`pulumi.Input[str]`) - Gets name of the resource that is unique within a resource group. This name can be used to access the resource
                * `properties` (`pulumi.Input[dict]`)
                  * `access` (`pulumi.Input[str]`) - Gets or sets network traffic is allowed or denied. Possible values are 'Allow' and 'Deny'
                  * `description` (`pulumi.Input[str]`) - Gets or sets a description for this rule. Restricted to 140 chars.
                  * `destination_address_prefix` (`pulumi.Input[str]`) - Gets or sets destination address prefix. CIDR or source IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used. 
                  * `destination_port_range` (`pulumi.Input[str]`) - Gets or sets Destination Port or Range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.
                  * `direction` (`pulumi.Input[str]`) - Gets or sets the direction of the rule.InBound or Outbound. The direction specifies if rule will be evaluated on incoming or outgoing traffic.
                  * `priority` (`pulumi.Input[float]`) - Gets or sets the priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
                  * `protocol` (`pulumi.Input[str]`) - Gets or sets Network protocol this rule applies to. Can be Tcp, Udp or All(*).
                  * `provisioning_state` (`pulumi.Input[str]`) - Gets provisioning state of the PublicIP resource Updating/Deleting/Failed
                  * `source_address_prefix` (`pulumi.Input[str]`) - Gets or sets source address prefix. CIDR or source IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used. If this is an ingress rule, specifies where network traffic originates from. 
                  * `source_port_range` (`pulumi.Input[str]`) - Gets or sets Source Port or Range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.

              * `provisioning_state` (`pulumi.Input[str]`) - Gets provisioning state of the PublicIP resource Updating/Deleting/Failed
              * `resource_guid` (`pulumi.Input[str]`) - Gets or sets resource guid property of the network security group resource
              * `security_rules` (`pulumi.Input[list]`) - Gets or sets security rules of network security group

            * `tags` (`pulumi.Input[dict]`) - Resource tags

          * `provisioning_state` (`pulumi.Input[str]`) - Gets provisioning state of the resource
          * `resource_navigation_links` (`pulumi.Input[list]`) - Gets array of references to the external resources using subnet
            * `id` (`pulumi.Input[str]`) - Resource Id
            * `name` (`pulumi.Input[str]`) - Name of the resource that is unique within a resource group. This name can be used to access the resource
            * `properties` (`pulumi.Input[dict]`) - Properties of ResourceNavigationLink
              * `link` (`pulumi.Input[str]`) - Link to the external resource
              * `linked_resource_type` (`pulumi.Input[str]`) - Resource type of the linked resource

          * `route_table` (`pulumi.Input[dict]`) - Gets or sets the reference of the RouteTable resource
            * `etag` (`pulumi.Input[str]`) - Gets a unique read-only string that changes whenever the resource is updated
            * `id` (`pulumi.Input[str]`) - Resource Id
            * `location` (`pulumi.Input[str]`) - Resource location
            * `properties` (`pulumi.Input[dict]`) - Route Table resource
              * `provisioning_state` (`pulumi.Input[str]`) - Gets provisioning state of the resource Updating/Deleting/Failed
              * `routes` (`pulumi.Input[list]`) - Gets or sets Routes in a Route Table
                * `etag` (`pulumi.Input[str]`) - A unique read-only string that changes whenever the resource is updated
                * `id` (`pulumi.Input[str]`) - Resource Id
                * `name` (`pulumi.Input[str]`) - Gets name of the resource that is unique within a resource group. This name can be used to access the resource
                * `properties` (`pulumi.Input[dict]`) - Route resource
                  * `address_prefix` (`pulumi.Input[str]`) - Gets or sets the destination CIDR to which the route applies.
                  * `next_hop_ip_address` (`pulumi.Input[str]`) - Gets or sets the IP address packets should be forwarded to. Next hop values are only allowed in routes where the next hop type is VirtualAppliance.
                  * `next_hop_type` (`pulumi.Input[str]`) - Gets or sets the type of Azure hop the packet should be sent to.
                  * `provisioning_state` (`pulumi.Input[str]`) - Gets provisioning state of the resource Updating/Deleting/Failed

            * `tags` (`pulumi.Input[dict]`) - Resource tags
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
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if virtual_network_name is None:
                raise TypeError("Missing required property 'virtual_network_name'")
            __props__['virtual_network_name'] = virtual_network_name
        super(Subnet, __self__).__init__(
            'azurerm:network/v20160601:Subnet',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing Subnet resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Subnet(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
