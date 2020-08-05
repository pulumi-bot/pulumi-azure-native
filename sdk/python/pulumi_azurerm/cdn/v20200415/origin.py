# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class Origin(pulumi.CustomResource):
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
    The JSON object that contains the properties of the origin.
      * `enabled` (`bool`) - Origin is enabled for load balancing or not
      * `host_name` (`str`) - The address of the origin. Domain names, IPv4 addresses, and IPv6 addresses are supported.This should be unique across all origins in an endpoint.
      * `http_port` (`float`) - The value of the HTTP port. Must be between 1 and 65535.
      * `https_port` (`float`) - The value of the HTTPS port. Must be between 1 and 65535.
      * `origin_host_header` (`str`) - The host header value sent to the origin with each request. If you leave this blank, the request hostname determines this value. Azure CDN origins, such as Web Apps, Blob Storage, and Cloud Services require this host header value to match the origin hostname by default. This overrides the host header defined at Endpoint
      * `priority` (`float`) - Priority of origin in given origin group for load balancing. Higher priorities will not be used for load balancing if any lower priority origin is healthy.Must be between 1 and 5
      * `private_endpoint_status` (`str`) - The approval status for the connection to the Private Link
      * `private_link_alias` (`str`) - The Alias of the Private Link resource. Populating this optional field indicates that this origin is 'Private'
      * `private_link_approval_message` (`str`) - A custom message to be included in the approval request to connect to the Private Link.
      * `private_link_location` (`str`) - The location of the Private Link resource. Required only if 'privateLinkResourceId' is populated
      * `private_link_resource_id` (`str`) - The Resource Id of the Private Link resource. Populating this optional field indicates that this backend is 'Private'
      * `provisioning_state` (`str`) - Provisioning status of the origin.
      * `resource_state` (`str`) - Resource status of the origin.
      * `weight` (`float`) - Weight of the origin in given origin group for load balancing. Must be between 1 and 1000
    """
    tags: pulumi.Output[dict]
    """
    Resource tags.
    """
    type: pulumi.Output[str]
    """
    Resource type.
    """
    def __init__(__self__, resource_name, opts=None, enabled=None, endpoint_name=None, host_name=None, http_port=None, https_port=None, location=None, name=None, origin_host_header=None, priority=None, private_link_alias=None, private_link_approval_message=None, private_link_location=None, private_link_resource_id=None, profile_name=None, resource_group_name=None, tags=None, weight=None, __props__=None, __name__=None, __opts__=None):
        """
        CDN origin is the source of the content being delivered via CDN. When the edge nodes represented by an endpoint do not have the requested content cached, they attempt to fetch it from one or more of the configured origins.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enabled: Origin is enabled for load balancing or not
        :param pulumi.Input[str] endpoint_name: Name of the endpoint under the profile which is unique globally.
        :param pulumi.Input[str] host_name: The address of the origin. Domain names, IPv4 addresses, and IPv6 addresses are supported.This should be unique across all origins in an endpoint.
        :param pulumi.Input[float] http_port: The value of the HTTP port. Must be between 1 and 65535.
        :param pulumi.Input[float] https_port: The value of the HTTPS port. Must be between 1 and 65535.
        :param pulumi.Input[str] location: Resource location.
        :param pulumi.Input[str] name: Name of the origin that is unique within the endpoint.
        :param pulumi.Input[str] origin_host_header: The host header value sent to the origin with each request. If you leave this blank, the request hostname determines this value. Azure CDN origins, such as Web Apps, Blob Storage, and Cloud Services require this host header value to match the origin hostname by default. This overrides the host header defined at Endpoint
        :param pulumi.Input[float] priority: Priority of origin in given origin group for load balancing. Higher priorities will not be used for load balancing if any lower priority origin is healthy.Must be between 1 and 5
        :param pulumi.Input[str] private_link_alias: The Alias of the Private Link resource. Populating this optional field indicates that this origin is 'Private'
        :param pulumi.Input[str] private_link_approval_message: A custom message to be included in the approval request to connect to the Private Link.
        :param pulumi.Input[str] private_link_location: The location of the Private Link resource. Required only if 'privateLinkResourceId' is populated
        :param pulumi.Input[str] private_link_resource_id: The Resource Id of the Private Link resource. Populating this optional field indicates that this backend is 'Private'
        :param pulumi.Input[str] profile_name: Name of the CDN profile which is unique within the resource group.
        :param pulumi.Input[str] resource_group_name: Name of the Resource group within the Azure subscription.
        :param pulumi.Input[dict] tags: Resource tags.
        :param pulumi.Input[float] weight: Weight of the origin in given origin group for load balancing. Must be between 1 and 1000
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

            __props__['enabled'] = enabled
            if endpoint_name is None:
                raise TypeError("Missing required property 'endpoint_name'")
            __props__['endpoint_name'] = endpoint_name
            if host_name is None:
                raise TypeError("Missing required property 'host_name'")
            __props__['host_name'] = host_name
            __props__['http_port'] = http_port
            __props__['https_port'] = https_port
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['origin_host_header'] = origin_host_header
            __props__['priority'] = priority
            __props__['private_link_alias'] = private_link_alias
            __props__['private_link_approval_message'] = private_link_approval_message
            __props__['private_link_location'] = private_link_location
            __props__['private_link_resource_id'] = private_link_resource_id
            if profile_name is None:
                raise TypeError("Missing required property 'profile_name'")
            __props__['profile_name'] = profile_name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['weight'] = weight
            __props__['properties'] = None
            __props__['type'] = None
        super(Origin, __self__).__init__(
            'azurerm:cdn/v20200415:Origin',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing Origin resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Origin(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
