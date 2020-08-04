# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class ExpressRouteConnection(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    The name of the resource.
    """
    properties: pulumi.Output[dict]
    """
    Properties of the ExpressRouteConnection subresource.
      * `authorization_key` (`str`) - Authorization key to establish the connection.
      * `express_route_circuit_peering` (`dict`) - The ExpressRoute circuit peering.
        * `id` (`str`) - The ID of the ExpressRoute circuit peering.

      * `provisioning_state` (`str`) - The provisioning state of the resource.
      * `routing_weight` (`float`) - The routing weight associated to the connection.
    """
    def __init__(__self__, resource_name, opts=None, express_route_gateway_name=None, id=None, name=None, properties=None, resource_group_name=None, __props__=None, __name__=None, __opts__=None):
        """
        ExpressRouteConnection resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] express_route_gateway_name: The name of the ExpressRoute gateway.
        :param pulumi.Input[str] id: Resource ID.
        :param pulumi.Input[str] name: The name of the connection subresource.
        :param pulumi.Input[dict] properties: Properties of the ExpressRouteConnection subresource.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.

        The **properties** object supports the following:

          * `authorization_key` (`pulumi.Input[str]`) - Authorization key to establish the connection.
          * `express_route_circuit_peering` (`pulumi.Input[dict]`) - The ExpressRoute circuit peering.
            * `id` (`pulumi.Input[str]`) - The ID of the ExpressRoute circuit peering.

          * `routing_weight` (`pulumi.Input[float]`) - The routing weight associated to the connection.
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

            if express_route_gateway_name is None:
                raise TypeError("Missing required property 'express_route_gateway_name'")
            __props__['express_route_gateway_name'] = express_route_gateway_name
            __props__['id'] = id
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
        super(ExpressRouteConnection, __self__).__init__(
            'azurerm:network/v20181201:ExpressRouteConnection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing ExpressRouteConnection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ExpressRouteConnection(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
