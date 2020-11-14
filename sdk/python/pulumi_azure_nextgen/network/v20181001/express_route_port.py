# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['ExpressRoutePort']


class ExpressRoutePort(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bandwidth_in_gbps: Optional[pulumi.Input[int]] = None,
                 encapsulation: Optional[pulumi.Input[str]] = None,
                 express_route_port_name: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 links: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ExpressRouteLinkArgs']]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 peering_location: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 resource_guid: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        ExpressRoutePort resource definition.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] bandwidth_in_gbps: Bandwidth of procured ports in Gbps
        :param pulumi.Input[str] encapsulation: Encapsulation method on physical ports.
        :param pulumi.Input[str] express_route_port_name: The name of the ExpressRoutePort resource.
        :param pulumi.Input[str] id: Resource ID.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ExpressRouteLinkArgs']]]] links: The set of physical links of the ExpressRoutePort resource
        :param pulumi.Input[str] location: Resource location.
        :param pulumi.Input[str] peering_location: The name of the peering location that the ExpressRoutePort is mapped to physically.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[str] resource_guid: The resource GUID property of the ExpressRoutePort resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags.
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
            if express_route_port_name is None:
                raise TypeError("Missing required property 'express_route_port_name'")
            __props__['express_route_port_name'] = express_route_port_name
            __props__['id'] = id
            __props__['links'] = links
            __props__['location'] = location
            __props__['peering_location'] = peering_location
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['resource_guid'] = resource_guid
            __props__['tags'] = tags
            __props__['allocation_date'] = None
            __props__['circuits'] = None
            __props__['etag'] = None
            __props__['ether_type'] = None
            __props__['mtu'] = None
            __props__['name'] = None
            __props__['provisioned_bandwidth_in_gbps'] = None
            __props__['provisioning_state'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:network/latest:ExpressRoutePort"), pulumi.Alias(type_="azure-nextgen:network/v20180801:ExpressRoutePort"), pulumi.Alias(type_="azure-nextgen:network/v20181101:ExpressRoutePort"), pulumi.Alias(type_="azure-nextgen:network/v20181201:ExpressRoutePort"), pulumi.Alias(type_="azure-nextgen:network/v20190201:ExpressRoutePort"), pulumi.Alias(type_="azure-nextgen:network/v20190401:ExpressRoutePort"), pulumi.Alias(type_="azure-nextgen:network/v20190601:ExpressRoutePort"), pulumi.Alias(type_="azure-nextgen:network/v20190701:ExpressRoutePort"), pulumi.Alias(type_="azure-nextgen:network/v20190801:ExpressRoutePort"), pulumi.Alias(type_="azure-nextgen:network/v20190901:ExpressRoutePort"), pulumi.Alias(type_="azure-nextgen:network/v20191101:ExpressRoutePort"), pulumi.Alias(type_="azure-nextgen:network/v20191201:ExpressRoutePort"), pulumi.Alias(type_="azure-nextgen:network/v20200301:ExpressRoutePort"), pulumi.Alias(type_="azure-nextgen:network/v20200401:ExpressRoutePort"), pulumi.Alias(type_="azure-nextgen:network/v20200501:ExpressRoutePort"), pulumi.Alias(type_="azure-nextgen:network/v20200601:ExpressRoutePort"), pulumi.Alias(type_="azure-nextgen:network/v20200701:ExpressRoutePort")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(ExpressRoutePort, __self__).__init__(
            'azure-nextgen:network/v20181001:ExpressRoutePort',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ExpressRoutePort':
        """
        Get an existing ExpressRoutePort resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ExpressRoutePort(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allocationDate")
    def allocation_date(self) -> pulumi.Output[str]:
        """
        Date of the physical port allocation to be used in Letter of Authorization.
        """
        return pulumi.get(self, "allocation_date")

    @property
    @pulumi.getter(name="bandwidthInGbps")
    def bandwidth_in_gbps(self) -> pulumi.Output[Optional[int]]:
        """
        Bandwidth of procured ports in Gbps
        """
        return pulumi.get(self, "bandwidth_in_gbps")

    @property
    @pulumi.getter
    def circuits(self) -> pulumi.Output[Sequence['outputs.SubResourceResponse']]:
        """
        Reference the ExpressRoute circuit(s) that are provisioned on this ExpressRoutePort resource.
        """
        return pulumi.get(self, "circuits")

    @property
    @pulumi.getter
    def encapsulation(self) -> pulumi.Output[Optional[str]]:
        """
        Encapsulation method on physical ports.
        """
        return pulumi.get(self, "encapsulation")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        A unique read-only string that changes whenever the resource is updated.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="etherType")
    def ether_type(self) -> pulumi.Output[str]:
        """
        Ether type of the physical port.
        """
        return pulumi.get(self, "ether_type")

    @property
    @pulumi.getter
    def links(self) -> pulumi.Output[Optional[Sequence['outputs.ExpressRouteLinkResponse']]]:
        """
        The set of physical links of the ExpressRoutePort resource
        """
        return pulumi.get(self, "links")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[Optional[str]]:
        """
        Resource location.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def mtu(self) -> pulumi.Output[str]:
        """
        Maximum transmission unit of the physical port pair(s)
        """
        return pulumi.get(self, "mtu")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="peeringLocation")
    def peering_location(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the peering location that the ExpressRoutePort is mapped to physically.
        """
        return pulumi.get(self, "peering_location")

    @property
    @pulumi.getter(name="provisionedBandwidthInGbps")
    def provisioned_bandwidth_in_gbps(self) -> pulumi.Output[float]:
        """
        Aggregate Gbps of associated circuit bandwidths.
        """
        return pulumi.get(self, "provisioned_bandwidth_in_gbps")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        The provisioning state of the ExpressRoutePort resource. Possible values are: 'Succeeded', 'Updating', 'Deleting', and 'Failed'.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="resourceGuid")
    def resource_guid(self) -> pulumi.Output[Optional[str]]:
        """
        The resource GUID property of the ExpressRoutePort resource.
        """
        return pulumi.get(self, "resource_guid")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Resource type.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

