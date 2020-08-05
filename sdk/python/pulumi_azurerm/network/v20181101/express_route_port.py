# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class ExpressRoutePort(pulumi.CustomResource):
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
    ExpressRoutePort properties
      * `allocation_date` (`str`) - Date of the physical port allocation to be used in Letter of Authorization.
      * `bandwidth_in_gbps` (`float`) - Bandwidth of procured ports in Gbps
      * `circuits` (`list`) - Reference the ExpressRoute circuit(s) that are provisioned on this ExpressRoutePort resource.
        * `id` (`str`) - Resource ID.

      * `encapsulation` (`str`) - Encapsulation method on physical ports.
      * `ether_type` (`str`) - Ether type of the physical port.
      * `links` (`list`) - The set of physical links of the ExpressRoutePort resource
        * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
        * `id` (`str`) - Resource ID.
        * `name` (`str`) - Name of child port resource that is unique among child port resources of the parent.
        * `properties` (`dict`) - ExpressRouteLink properties
          * `admin_state` (`str`) - Administrative state of the physical port
          * `connector_type` (`str`) - Physical fiber port type.
          * `interface_name` (`str`) - Name of Azure router interface.
          * `patch_panel_id` (`str`) - Mapping between physical port to patch panel port.
          * `provisioning_state` (`str`) - The provisioning state of the ExpressRouteLink resource. Possible values are: 'Succeeded', 'Updating', 'Deleting', and 'Failed'.
          * `rack_id` (`str`) - Mapping of physical patch panel to rack.
          * `router_name` (`str`) - Name of Azure router associated with physical port.

      * `mtu` (`str`) - Maximum transmission unit of the physical port pair(s)
      * `peering_location` (`str`) - The name of the peering location that the ExpressRoutePort is mapped to physically.
      * `provisioned_bandwidth_in_gbps` (`float`) - Aggregate Gbps of associated circuit bandwidths.
      * `provisioning_state` (`str`) - The provisioning state of the ExpressRoutePort resource. Possible values are: 'Succeeded', 'Updating', 'Deleting', and 'Failed'.
      * `resource_guid` (`str`) - The resource GUID property of the ExpressRoutePort resource.
    """
    tags: pulumi.Output[dict]
    """
    Resource tags.
    """
    type: pulumi.Output[str]
    """
    Resource type.
    """
    def __init__(__self__, resource_name, opts=None, bandwidth_in_gbps=None, encapsulation=None, id=None, links=None, location=None, name=None, peering_location=None, resource_group_name=None, resource_guid=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        ExpressRoutePort resource definition.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] bandwidth_in_gbps: Bandwidth of procured ports in Gbps
        :param pulumi.Input[str] encapsulation: Encapsulation method on physical ports.
        :param pulumi.Input[str] id: Resource ID.
        :param pulumi.Input[list] links: The set of physical links of the ExpressRoutePort resource
        :param pulumi.Input[str] location: Resource location.
        :param pulumi.Input[str] name: The name of the ExpressRoutePort resource.
        :param pulumi.Input[str] peering_location: The name of the peering location that the ExpressRoutePort is mapped to physically.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[str] resource_guid: The resource GUID property of the ExpressRoutePort resource.
        :param pulumi.Input[dict] tags: Resource tags.

        The **links** object supports the following:

          * `admin_state` (`pulumi.Input[str]`) - Administrative state of the physical port
          * `id` (`pulumi.Input[str]`) - Resource ID.
          * `name` (`pulumi.Input[str]`) - Name of child port resource that is unique among child port resources of the parent.
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

            __props__['bandwidth_in_gbps'] = bandwidth_in_gbps
            __props__['encapsulation'] = encapsulation
            __props__['id'] = id
            __props__['links'] = links
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['peering_location'] = peering_location
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['resource_guid'] = resource_guid
            __props__['tags'] = tags
            __props__['etag'] = None
            __props__['properties'] = None
            __props__['type'] = None
        super(ExpressRoutePort, __self__).__init__(
            'azurerm:network/v20181101:ExpressRoutePort',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing ExpressRoutePort resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ExpressRoutePort(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
