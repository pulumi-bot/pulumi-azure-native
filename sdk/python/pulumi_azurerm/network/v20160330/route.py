# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class Route(pulumi.CustomResource):
    etag: pulumi.Output[str]
    """
    A unique read-only string that changes whenever the resource is updated
    """
    name: pulumi.Output[str]
    """
    Gets name of the resource that is unique within a resource group. This name can be used to access the resource
    """
    properties: pulumi.Output[dict]
    """
    Route resource
      * `address_prefix` (`str`) - Gets or sets the destination CIDR to which the route applies.
      * `next_hop_ip_address` (`str`) - Gets or sets the IP address packets should be forwarded to. Next hop values are only allowed in routes where the next hop type is VirtualAppliance.
      * `next_hop_type` (`str`) - Gets or sets the type of Azure hop the packet should be sent to.
      * `provisioning_state` (`str`) - Gets or sets Provisioning state of the resource Updating/Deleting/Failed
    """
    def __init__(__self__, resource_name, opts=None, etag=None, id=None, name=None, properties=None, resource_group_name=None, route_table_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Route resource

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] etag: A unique read-only string that changes whenever the resource is updated
        :param pulumi.Input[str] id: Resource Id
        :param pulumi.Input[str] name: The name of the route.
        :param pulumi.Input[dict] properties: Route resource
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[str] route_table_name: The name of the route table.

        The **properties** object supports the following:

          * `address_prefix` (`pulumi.Input[str]`) - Gets or sets the destination CIDR to which the route applies.
          * `next_hop_ip_address` (`pulumi.Input[str]`) - Gets or sets the IP address packets should be forwarded to. Next hop values are only allowed in routes where the next hop type is VirtualAppliance.
          * `next_hop_type` (`pulumi.Input[str]`) - Gets or sets the type of Azure hop the packet should be sent to.
          * `provisioning_state` (`pulumi.Input[str]`) - Gets or sets Provisioning state of the resource Updating/Deleting/Failed
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
            if route_table_name is None:
                raise TypeError("Missing required property 'route_table_name'")
            __props__['route_table_name'] = route_table_name
        super(Route, __self__).__init__(
            'azurerm:network/v20160330:Route',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing Route resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Route(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
