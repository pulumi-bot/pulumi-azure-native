# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Endpoint']


class Endpoint(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 content_types_to_compress: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 geo_filters: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['GeoFilterArgs']]]]] = None,
                 is_compression_enabled: Optional[pulumi.Input[bool]] = None,
                 is_http_allowed: Optional[pulumi.Input[bool]] = None,
                 is_https_allowed: Optional[pulumi.Input[bool]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 optimization_type: Optional[pulumi.Input[str]] = None,
                 origin_host_header: Optional[pulumi.Input[str]] = None,
                 origin_path: Optional[pulumi.Input[str]] = None,
                 origins: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['DeepCreatedOriginArgs']]]]] = None,
                 profile_name: Optional[pulumi.Input[str]] = None,
                 query_string_caching_behavior: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        CDN endpoint is the entity within a CDN profile containing configuration information such as origin, protocol, content caching and delivery behavior. The CDN endpoint uses the URL format <endpointname>.azureedge.net.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[str]]] content_types_to_compress: List of content types on which compression applies. The value should be a valid MIME type.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['GeoFilterArgs']]]] geo_filters: List of rules defining user geo access within a CDN endpoint. Each geo filter defines an access rule to a specified path or content, e.g. block APAC for path /pictures/
        :param pulumi.Input[bool] is_compression_enabled: Indicates whether content compression is enabled on CDN. Default value is false. If compression is enabled, content will be served as compressed if user requests for a compressed version. Content won't be compressed on CDN when requested content is smaller than 1 byte or larger than 1 MB.
        :param pulumi.Input[bool] is_http_allowed: Indicates whether HTTP traffic is allowed on the endpoint. Default value is true. At least one protocol (HTTP or HTTPS) must be allowed.
        :param pulumi.Input[bool] is_https_allowed: Indicates whether HTTPS traffic is allowed on the endpoint. Default value is true. At least one protocol (HTTP or HTTPS) must be allowed.
        :param pulumi.Input[str] location: Resource location.
        :param pulumi.Input[str] name: Name of the endpoint under the profile which is unique globally.
        :param pulumi.Input[str] optimization_type: Customer can specify what scenario they want this CDN endpoint to optimize, e.g. Download, Media services. With this information we can apply scenario driven optimization.
        :param pulumi.Input[str] origin_host_header: The host header CDN sends along with content requests to origin. The default value is the host name of the origin.
        :param pulumi.Input[str] origin_path: The path used when CDN sends request to origin.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['DeepCreatedOriginArgs']]]] origins: The source of the content being delivered via CDN.
        :param pulumi.Input[str] profile_name: Name of the CDN profile which is unique within the resource group.
        :param pulumi.Input[str] query_string_caching_behavior: Defines the query string caching behavior
        :param pulumi.Input[str] resource_group_name: Name of the Resource group within the Azure subscription.
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

            __props__['content_types_to_compress'] = content_types_to_compress
            __props__['geo_filters'] = geo_filters
            __props__['is_compression_enabled'] = is_compression_enabled
            __props__['is_http_allowed'] = is_http_allowed
            __props__['is_https_allowed'] = is_https_allowed
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['optimization_type'] = optimization_type
            __props__['origin_host_header'] = origin_host_header
            __props__['origin_path'] = origin_path
            if origins is None:
                raise TypeError("Missing required property 'origins'")
            __props__['origins'] = origins
            if profile_name is None:
                raise TypeError("Missing required property 'profile_name'")
            __props__['profile_name'] = profile_name
            __props__['query_string_caching_behavior'] = query_string_caching_behavior
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['host_name'] = None
            __props__['provisioning_state'] = None
            __props__['resource_state'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:cdn/v20150601:Endpoint"), pulumi.Alias(type_="azurerm:cdn/v20160402:Endpoint"), pulumi.Alias(type_="azurerm:cdn/v20170402:Endpoint"), pulumi.Alias(type_="azurerm:cdn/v20171012:Endpoint"), pulumi.Alias(type_="azurerm:cdn/v20190415:Endpoint"), pulumi.Alias(type_="azurerm:cdn/v20190615:Endpoint"), pulumi.Alias(type_="azurerm:cdn/v20191231:Endpoint"), pulumi.Alias(type_="azurerm:cdn/v20200331:Endpoint"), pulumi.Alias(type_="azurerm:cdn/v20200415:Endpoint")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Endpoint, __self__).__init__(
            'azurerm:cdn/v20161002:Endpoint',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Endpoint':
        """
        Get an existing Endpoint resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Endpoint(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="contentTypesToCompress")
    def content_types_to_compress(self) -> Optional[List[str]]:
        """
        List of content types on which compression applies. The value should be a valid MIME type.
        """
        return pulumi.get(self, "content_types_to_compress")

    @property
    @pulumi.getter(name="geoFilters")
    def geo_filters(self) -> Optional[List['outputs.GeoFilterResponse']]:
        """
        List of rules defining user geo access within a CDN endpoint. Each geo filter defines an access rule to a specified path or content, e.g. block APAC for path /pictures/
        """
        return pulumi.get(self, "geo_filters")

    @property
    @pulumi.getter(name="hostName")
    def host_name(self) -> str:
        """
        The host name of the endpoint structured as {endpointName}.{DNSZone}, e.g. contoso.azureedge.net
        """
        return pulumi.get(self, "host_name")

    @property
    @pulumi.getter(name="isCompressionEnabled")
    def is_compression_enabled(self) -> Optional[bool]:
        """
        Indicates whether content compression is enabled on CDN. Default value is false. If compression is enabled, content will be served as compressed if user requests for a compressed version. Content won't be compressed on CDN when requested content is smaller than 1 byte or larger than 1 MB.
        """
        return pulumi.get(self, "is_compression_enabled")

    @property
    @pulumi.getter(name="isHttpAllowed")
    def is_http_allowed(self) -> Optional[bool]:
        """
        Indicates whether HTTP traffic is allowed on the endpoint. Default value is true. At least one protocol (HTTP or HTTPS) must be allowed.
        """
        return pulumi.get(self, "is_http_allowed")

    @property
    @pulumi.getter(name="isHttpsAllowed")
    def is_https_allowed(self) -> Optional[bool]:
        """
        Indicates whether HTTPS traffic is allowed on the endpoint. Default value is true. At least one protocol (HTTP or HTTPS) must be allowed.
        """
        return pulumi.get(self, "is_https_allowed")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        Resource location.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="optimizationType")
    def optimization_type(self) -> Optional[str]:
        """
        Customer can specify what scenario they want this CDN endpoint to optimize, e.g. Download, Media services. With this information we can apply scenario driven optimization.
        """
        return pulumi.get(self, "optimization_type")

    @property
    @pulumi.getter(name="originHostHeader")
    def origin_host_header(self) -> Optional[str]:
        """
        The host header CDN sends along with content requests to origin. The default value is the host name of the origin.
        """
        return pulumi.get(self, "origin_host_header")

    @property
    @pulumi.getter(name="originPath")
    def origin_path(self) -> Optional[str]:
        """
        The path used when CDN sends request to origin.
        """
        return pulumi.get(self, "origin_path")

    @property
    @pulumi.getter
    def origins(self) -> List['outputs.DeepCreatedOriginResponse']:
        """
        The source of the content being delivered via CDN.
        """
        return pulumi.get(self, "origins")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        Provisioning status of the endpoint.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="queryStringCachingBehavior")
    def query_string_caching_behavior(self) -> Optional[str]:
        """
        Defines the query string caching behavior
        """
        return pulumi.get(self, "query_string_caching_behavior")

    @property
    @pulumi.getter(name="resourceState")
    def resource_state(self) -> str:
        """
        Resource status of the endpoint.
        """
        return pulumi.get(self, "resource_state")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Resource type.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

