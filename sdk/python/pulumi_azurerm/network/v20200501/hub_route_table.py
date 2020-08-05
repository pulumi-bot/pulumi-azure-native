# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class HubRouteTable(pulumi.CustomResource):
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
    Properties of the RouteTable resource.
      * `associated_connections` (`list`) - List of all connections associated with this route table.
      * `labels` (`list`) - List of labels associated with this route table.
      * `propagating_connections` (`list`) - List of all connections that advertise to this route table.
      * `provisioning_state` (`str`) - The provisioning state of the RouteTable resource.
      * `routes` (`list`) - List of all routes.
        * `destination_type` (`str`) - The type of destinations (eg: CIDR, ResourceId, Service).
        * `destinations` (`list`) - List of all destinations.
        * `name` (`str`) - The name of the Route that is unique within a RouteTable. This name can be used to access this route.
        * `next_hop` (`str`) - NextHop resource ID.
        * `next_hop_type` (`str`) - The type of next hop (eg: ResourceId).
    """
    type: pulumi.Output[str]
    """
    Resource type.
    """
    def __init__(__self__, resource_name, opts=None, id=None, labels=None, name=None, resource_group_name=None, routes=None, virtual_hub_name=None, __props__=None, __name__=None, __opts__=None):
        """
        RouteTable resource in a virtual hub.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] id: Resource ID.
        :param pulumi.Input[list] labels: List of labels associated with this route table.
        :param pulumi.Input[str] name: The name of the RouteTable.
        :param pulumi.Input[str] resource_group_name: The resource group name of the VirtualHub.
        :param pulumi.Input[list] routes: List of all routes.
        :param pulumi.Input[str] virtual_hub_name: The name of the VirtualHub.

        The **routes** object supports the following:

          * `destination_type` (`pulumi.Input[str]`) - The type of destinations (eg: CIDR, ResourceId, Service).
          * `destinations` (`pulumi.Input[list]`) - List of all destinations.
          * `name` (`pulumi.Input[str]`) - The name of the Route that is unique within a RouteTable. This name can be used to access this route.
          * `next_hop` (`pulumi.Input[str]`) - NextHop resource ID.
          * `next_hop_type` (`pulumi.Input[str]`) - The type of next hop (eg: ResourceId).
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
            __props__['labels'] = labels
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['routes'] = routes
            if virtual_hub_name is None:
                raise TypeError("Missing required property 'virtual_hub_name'")
            __props__['virtual_hub_name'] = virtual_hub_name
            __props__['etag'] = None
            __props__['properties'] = None
            __props__['type'] = None
        super(HubRouteTable, __self__).__init__(
            'azurerm:network/v20200501:HubRouteTable',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing HubRouteTable resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return HubRouteTable(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
