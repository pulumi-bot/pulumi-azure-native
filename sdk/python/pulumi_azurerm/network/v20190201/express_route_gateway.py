# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class ExpressRouteGateway(pulumi.CustomResource):
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
    Properties of the express route gateway.
      * `auto_scale_configuration` (`dict`) - Configuration for auto scaling.
        * `bounds` (`dict`) - Minimum and maximum number of scale units to deploy.

      * `express_route_connections` (`list`) - List of ExpressRoute connections to the ExpressRoute gateway.
        * `id` (`str`) - Resource ID.
        * `name` (`str`) - The name of the resource.
        * `properties` (`dict`) - Properties of the express route connection.
          * `authorization_key` (`str`) - Authorization key to establish the connection.
          * `express_route_circuit_peering` (`dict`) - The ExpressRoute circuit peering.
            * `id` (`str`) - The ID of the ExpressRoute circuit peering.

          * `provisioning_state` (`str`) - The provisioning state of the resource.
          * `routing_weight` (`float`) - The routing weight associated to the connection.

      * `provisioning_state` (`str`) - The provisioning state of the resource.
      * `virtual_hub` (`dict`) - The Virtual Hub where the ExpressRoute gateway is or will be deployed.
        * `id` (`str`) - The resource URI for the Virtual Hub where the ExpressRoute gateway is or will be deployed. The Virtual Hub resource and the ExpressRoute gateway resource reside in the same subscription.
    """
    tags: pulumi.Output[dict]
    """
    Resource tags.
    """
    type: pulumi.Output[str]
    """
    Resource type.
    """
    def __init__(__self__, resource_name, opts=None, id=None, location=None, name=None, properties=None, resource_group_name=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        ExpressRoute gateway resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] id: Resource ID.
        :param pulumi.Input[str] location: Resource location.
        :param pulumi.Input[str] name: The name of the ExpressRoute gateway.
        :param pulumi.Input[dict] properties: Properties of the express route gateway.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[dict] tags: Resource tags.

        The **properties** object supports the following:

          * `auto_scale_configuration` (`pulumi.Input[dict]`) - Configuration for auto scaling.
            * `bounds` (`pulumi.Input[dict]`) - Minimum and maximum number of scale units to deploy.

          * `provisioning_state` (`pulumi.Input[str]`) - The provisioning state of the resource.
          * `virtual_hub` (`pulumi.Input[dict]`) - The Virtual Hub where the ExpressRoute gateway is or will be deployed.
            * `id` (`pulumi.Input[str]`) - The resource URI for the Virtual Hub where the ExpressRoute gateway is or will be deployed. The Virtual Hub resource and the ExpressRoute gateway resource reside in the same subscription.
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

            __props__['id'] = id
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['etag'] = None
            __props__['type'] = None
        super(ExpressRouteGateway, __self__).__init__(
            'azurerm:network/v20190201:ExpressRouteGateway',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing ExpressRouteGateway resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ExpressRouteGateway(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
