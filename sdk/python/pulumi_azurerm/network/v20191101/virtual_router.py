# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class VirtualRouter(pulumi.CustomResource):
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
    Properties of the Virtual Router.
      * `hosted_gateway` (`dict`) - The Gateway on which VirtualRouter is hosted.
        * `id` (`str`) - Resource ID.

      * `hosted_subnet` (`dict`) - The Subnet on which VirtualRouter is hosted.
      * `peerings` (`list`) - List of references to VirtualRouterPeerings.
      * `provisioning_state` (`str`) - The provisioning state of the resource.
      * `virtual_router_asn` (`float`) - VirtualRouter ASN.
      * `virtual_router_ips` (`list`) - VirtualRouter IPs.
    """
    tags: pulumi.Output[dict]
    """
    Resource tags.
    """
    type: pulumi.Output[str]
    """
    Resource type.
    """
    def __init__(__self__, resource_name, opts=None, hosted_gateway=None, hosted_subnet=None, id=None, location=None, name=None, resource_group_name=None, tags=None, virtual_router_asn=None, virtual_router_ips=None, __props__=None, __name__=None, __opts__=None):
        """
        VirtualRouter Resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] hosted_gateway: The Gateway on which VirtualRouter is hosted.
        :param pulumi.Input[dict] hosted_subnet: The Subnet on which VirtualRouter is hosted.
        :param pulumi.Input[str] id: Resource ID.
        :param pulumi.Input[str] location: Resource location.
        :param pulumi.Input[str] name: The name of the Virtual Router.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[dict] tags: Resource tags.
        :param pulumi.Input[float] virtual_router_asn: VirtualRouter ASN.
        :param pulumi.Input[list] virtual_router_ips: VirtualRouter IPs.

        The **hosted_gateway** object supports the following:

          * `id` (`pulumi.Input[str]`) - Resource ID.
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

            __props__['hosted_gateway'] = hosted_gateway
            __props__['hosted_subnet'] = hosted_subnet
            __props__['id'] = id
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['virtual_router_asn'] = virtual_router_asn
            __props__['virtual_router_ips'] = virtual_router_ips
            __props__['etag'] = None
            __props__['properties'] = None
            __props__['type'] = None
        super(VirtualRouter, __self__).__init__(
            'azurerm:network/v20191101:VirtualRouter',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing VirtualRouter resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return VirtualRouter(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
