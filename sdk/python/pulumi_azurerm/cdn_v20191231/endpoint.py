# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class Endpoint(pulumi.CustomResource):
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
    The JSON object that contains the properties required to create an endpoint.
      * `content_types_to_compress` (`list`) - List of content types on which compression applies. The value should be a valid MIME type.
      * `default_origin_group` (`dict`) - A reference to the origin group.
        * `id` (`str`) - Resource ID.

      * `delivery_policy` (`dict`) - A policy that specifies the delivery rules to be used for an endpoint.
        * `description` (`str`) - User-friendly description of the policy.
        * `rules` (`list`) - A list of the delivery rules.
          * `actions` (`list`) - A list of actions that are executed when all the conditions of a rule are satisfied.
            * `name` (`str`) - The name of the action for the delivery rule.

          * `conditions` (`list`) - A list of conditions that must be matched for the actions to be executed
            * `name` (`str`) - The name of the condition for the delivery rule.

          * `name` (`str`) - Name of the rule
          * `order` (`float`) - The order in which the rules are applied for the endpoint. Possible values {0,1,2,3,………}. A rule with a lesser order will be applied before a rule with a greater order. Rule with order 0 is a special rule. It does not require any condition and actions listed in it will always be applied.

      * `geo_filters` (`list`) - List of rules defining the user's geo access within a CDN endpoint. Each geo filter defines an access rule to a specified path or content, e.g. block APAC for path /pictures/
        * `action` (`str`) - Action of the geo filter, i.e. allow or block access.
        * `country_codes` (`list`) - Two letter country codes defining user country access in a geo filter, e.g. AU, MX, US.
        * `relative_path` (`str`) - Relative path applicable to geo filter. (e.g. '/mypictures', '/mypicture/kitty.jpg', and etc.)

      * `host_name` (`str`) - The host name of the endpoint structured as {endpointName}.{DNSZone}, e.g. contoso.azureedge.net
      * `is_compression_enabled` (`bool`) - Indicates whether content compression is enabled on CDN. Default value is false. If compression is enabled, content will be served as compressed if user requests for a compressed version. Content won't be compressed on CDN when requested content is smaller than 1 byte or larger than 1 MB.
      * `is_http_allowed` (`bool`) - Indicates whether HTTP traffic is allowed on the endpoint. Default value is true. At least one protocol (HTTP or HTTPS) must be allowed.
      * `is_https_allowed` (`bool`) - Indicates whether HTTPS traffic is allowed on the endpoint. Default value is true. At least one protocol (HTTP or HTTPS) must be allowed.
      * `optimization_type` (`str`) - Specifies what scenario the customer wants this CDN endpoint to optimize for, e.g. Download, Media services. With this information, CDN can apply scenario driven optimization.
      * `origin_groups` (`list`) - The origin groups comprising of origins that are used for load balancing the traffic based on availability.
        * `name` (`str`) - Origin group name which must be unique within the endpoint.
        * `properties` (`dict`) - Properties of the origin group created on the CDN endpoint.
          * `health_probe_settings` (`dict`) - Health probe settings to the origin that is used to determine the health of the origin.
            * `probe_interval_in_seconds` (`float`) - The number of seconds between health probes.Default is 240sec.
            * `probe_path` (`str`) - The path relative to the origin that is used to determine the health of the origin.
            * `probe_protocol` (`str`) - Protocol to use for health probe.
            * `probe_request_type` (`str`) - The type of health probe request that is made.

          * `origins` (`list`) - The source of the content being delivered via CDN within given origin group.
          * `response_based_origin_error_detection_settings` (`dict`) - The JSON object that contains the properties to determine origin health using real requests/responses.This property is currently not supported.
            * `http_error_ranges` (`list`) - The list of Http status code ranges that are considered as server errors for origin and it is marked as unhealthy.
              * `begin` (`float`) - The inclusive start of the http status code range.
              * `end` (`float`) - The inclusive end of the http status code range.

            * `response_based_detected_error_types` (`str`) - Type of response errors for real user requests for which origin will be deemed unhealthy
            * `response_based_failover_threshold_percentage` (`float`) - The percentage of failed requests in the sample where failover should trigger.

          * `traffic_restoration_time_to_healed_or_new_endpoints_in_minutes` (`float`) - Time in minutes to shift the traffic to the endpoint gradually when an unhealthy endpoint comes healthy or a new endpoint is added. Default is 10 mins. This property is currently not supported.

      * `origin_host_header` (`str`) - The host header value sent to the origin with each request. This property at Endpoint can only be set allowed when endpoint uses single origin. If you leave this blank, the request hostname determines this value. Azure CDN origins, such as Web Apps, Blob Storage, and Cloud Services require this host header value to match the origin hostname by default.
      * `origin_path` (`str`) - A directory path on the origin that CDN can use to retrieve content from, e.g. contoso.cloudapp.net/originpath.
      * `origins` (`list`) - The source of the content being delivered via CDN.
        * `name` (`str`) - Origin name which must be unique within the endpoint. 
        * `properties` (`dict`) - Properties of the origin created on the CDN endpoint.
          * `enabled` (`bool`) - Origin is enabled for load balancing or not. By default, origin is always enabled.
          * `host_name` (`str`) - The address of the origin. It can be a domain name, IPv4 address, or IPv6 address. This should be unique across all origins in an endpoint.
          * `http_port` (`float`) - The value of the HTTP port. Must be between 1 and 65535.
          * `https_port` (`float`) - The value of the HTTPS port. Must be between 1 and 65535.
          * `origin_host_header` (`str`) - The host header value sent to the origin with each request. If you leave this blank, the request hostname determines this value. Azure CDN origins, such as Web Apps, Blob Storage, and Cloud Services require this host header value to match the origin hostname by default. If endpoint uses multiple origins for load balancing, then the host header at endpoint is ignored and this one is considered.
          * `priority` (`float`) - Priority of origin in given origin group for load balancing. Higher priorities will not be used for load balancing if any lower priority origin is healthy.Must be between 1 and 5.
          * `weight` (`float`) - Weight of the origin in given origin group for load balancing. Must be between 1 and 1000

      * `probe_path` (`str`) - Path to a file hosted on the origin which helps accelerate delivery of the dynamic content and calculate the most optimal routes for the CDN. This is relative to the origin path. This property is only relevant when using a single origin.
      * `provisioning_state` (`str`) - Provisioning status of the endpoint.
      * `query_string_caching_behavior` (`str`) - Defines how CDN caches requests that include query strings. You can ignore any query strings when caching, bypass caching to prevent requests that contain query strings from being cached, or cache every request with a unique URL.
      * `resource_state` (`str`) - Resource status of the endpoint.
    """
    tags: pulumi.Output[dict]
    """
    Resource tags.
    """
    type: pulumi.Output[str]
    """
    Resource type.
    """
    def __init__(__self__, resource_name, opts=None, location=None, name=None, profile_name=None, properties=None, resource_group_name=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        CDN endpoint is the entity within a CDN profile containing configuration information such as origin, protocol, content caching and delivery behavior. The CDN endpoint uses the URL format <endpointname>.azureedge.net.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: Resource location.
        :param pulumi.Input[str] name: Name of the endpoint under the profile which is unique globally.
        :param pulumi.Input[str] profile_name: Name of the CDN profile which is unique within the resource group.
        :param pulumi.Input[dict] properties: The JSON object that contains the properties required to create an endpoint.
        :param pulumi.Input[str] resource_group_name: Name of the Resource group within the Azure subscription.
        :param pulumi.Input[dict] tags: Resource tags.

        The **properties** object supports the following:

          * `content_types_to_compress` (`pulumi.Input[list]`) - List of content types on which compression applies. The value should be a valid MIME type.
          * `default_origin_group` (`pulumi.Input[dict]`) - A reference to the origin group.
            * `id` (`pulumi.Input[str]`) - Resource ID.

          * `delivery_policy` (`pulumi.Input[dict]`) - A policy that specifies the delivery rules to be used for an endpoint.
            * `description` (`pulumi.Input[str]`) - User-friendly description of the policy.
            * `rules` (`pulumi.Input[list]`) - A list of the delivery rules.
              * `actions` (`pulumi.Input[list]`) - A list of actions that are executed when all the conditions of a rule are satisfied.
                * `name` (`pulumi.Input[str]`) - The name of the action for the delivery rule.

              * `conditions` (`pulumi.Input[list]`) - A list of conditions that must be matched for the actions to be executed
                * `name` (`pulumi.Input[str]`) - The name of the condition for the delivery rule.

              * `name` (`pulumi.Input[str]`) - Name of the rule
              * `order` (`pulumi.Input[float]`) - The order in which the rules are applied for the endpoint. Possible values {0,1,2,3,………}. A rule with a lesser order will be applied before a rule with a greater order. Rule with order 0 is a special rule. It does not require any condition and actions listed in it will always be applied.

          * `geo_filters` (`pulumi.Input[list]`) - List of rules defining the user's geo access within a CDN endpoint. Each geo filter defines an access rule to a specified path or content, e.g. block APAC for path /pictures/
            * `action` (`pulumi.Input[str]`) - Action of the geo filter, i.e. allow or block access.
            * `country_codes` (`pulumi.Input[list]`) - Two letter country codes defining user country access in a geo filter, e.g. AU, MX, US.
            * `relative_path` (`pulumi.Input[str]`) - Relative path applicable to geo filter. (e.g. '/mypictures', '/mypicture/kitty.jpg', and etc.)

          * `is_compression_enabled` (`pulumi.Input[bool]`) - Indicates whether content compression is enabled on CDN. Default value is false. If compression is enabled, content will be served as compressed if user requests for a compressed version. Content won't be compressed on CDN when requested content is smaller than 1 byte or larger than 1 MB.
          * `is_http_allowed` (`pulumi.Input[bool]`) - Indicates whether HTTP traffic is allowed on the endpoint. Default value is true. At least one protocol (HTTP or HTTPS) must be allowed.
          * `is_https_allowed` (`pulumi.Input[bool]`) - Indicates whether HTTPS traffic is allowed on the endpoint. Default value is true. At least one protocol (HTTP or HTTPS) must be allowed.
          * `optimization_type` (`pulumi.Input[str]`) - Specifies what scenario the customer wants this CDN endpoint to optimize for, e.g. Download, Media services. With this information, CDN can apply scenario driven optimization.
          * `origin_groups` (`pulumi.Input[list]`) - The origin groups comprising of origins that are used for load balancing the traffic based on availability.
            * `name` (`pulumi.Input[str]`) - Origin group name which must be unique within the endpoint.
            * `properties` (`pulumi.Input[dict]`) - Properties of the origin group created on the CDN endpoint.
              * `health_probe_settings` (`pulumi.Input[dict]`) - Health probe settings to the origin that is used to determine the health of the origin.
                * `probe_interval_in_seconds` (`pulumi.Input[float]`) - The number of seconds between health probes.Default is 240sec.
                * `probe_path` (`pulumi.Input[str]`) - The path relative to the origin that is used to determine the health of the origin.
                * `probe_protocol` (`pulumi.Input[str]`) - Protocol to use for health probe.
                * `probe_request_type` (`pulumi.Input[str]`) - The type of health probe request that is made.

              * `origins` (`pulumi.Input[list]`) - The source of the content being delivered via CDN within given origin group.
              * `response_based_origin_error_detection_settings` (`pulumi.Input[dict]`) - The JSON object that contains the properties to determine origin health using real requests/responses.This property is currently not supported.
                * `http_error_ranges` (`pulumi.Input[list]`) - The list of Http status code ranges that are considered as server errors for origin and it is marked as unhealthy.
                  * `begin` (`pulumi.Input[float]`) - The inclusive start of the http status code range.
                  * `end` (`pulumi.Input[float]`) - The inclusive end of the http status code range.

                * `response_based_detected_error_types` (`pulumi.Input[str]`) - Type of response errors for real user requests for which origin will be deemed unhealthy
                * `response_based_failover_threshold_percentage` (`pulumi.Input[float]`) - The percentage of failed requests in the sample where failover should trigger.

              * `traffic_restoration_time_to_healed_or_new_endpoints_in_minutes` (`pulumi.Input[float]`) - Time in minutes to shift the traffic to the endpoint gradually when an unhealthy endpoint comes healthy or a new endpoint is added. Default is 10 mins. This property is currently not supported.

          * `origin_host_header` (`pulumi.Input[str]`) - The host header value sent to the origin with each request. This property at Endpoint can only be set allowed when endpoint uses single origin. If you leave this blank, the request hostname determines this value. Azure CDN origins, such as Web Apps, Blob Storage, and Cloud Services require this host header value to match the origin hostname by default.
          * `origin_path` (`pulumi.Input[str]`) - A directory path on the origin that CDN can use to retrieve content from, e.g. contoso.cloudapp.net/originpath.
          * `origins` (`pulumi.Input[list]`) - The source of the content being delivered via CDN.
            * `name` (`pulumi.Input[str]`) - Origin name which must be unique within the endpoint. 
            * `properties` (`pulumi.Input[dict]`) - Properties of the origin created on the CDN endpoint.
              * `enabled` (`pulumi.Input[bool]`) - Origin is enabled for load balancing or not. By default, origin is always enabled.
              * `host_name` (`pulumi.Input[str]`) - The address of the origin. It can be a domain name, IPv4 address, or IPv6 address. This should be unique across all origins in an endpoint.
              * `http_port` (`pulumi.Input[float]`) - The value of the HTTP port. Must be between 1 and 65535.
              * `https_port` (`pulumi.Input[float]`) - The value of the HTTPS port. Must be between 1 and 65535.
              * `origin_host_header` (`pulumi.Input[str]`) - The host header value sent to the origin with each request. If you leave this blank, the request hostname determines this value. Azure CDN origins, such as Web Apps, Blob Storage, and Cloud Services require this host header value to match the origin hostname by default. If endpoint uses multiple origins for load balancing, then the host header at endpoint is ignored and this one is considered.
              * `priority` (`pulumi.Input[float]`) - Priority of origin in given origin group for load balancing. Higher priorities will not be used for load balancing if any lower priority origin is healthy.Must be between 1 and 5.
              * `weight` (`pulumi.Input[float]`) - Weight of the origin in given origin group for load balancing. Must be between 1 and 1000

          * `probe_path` (`pulumi.Input[str]`) - Path to a file hosted on the origin which helps accelerate delivery of the dynamic content and calculate the most optimal routes for the CDN. This is relative to the origin path. This property is only relevant when using a single origin.
          * `query_string_caching_behavior` (`pulumi.Input[str]`) - Defines how CDN caches requests that include query strings. You can ignore any query strings when caching, bypass caching to prevent requests that contain query strings from being cached, or cache every request with a unique URL.
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

            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if profile_name is None:
                raise TypeError("Missing required property 'profile_name'")
            __props__['profile_name'] = profile_name
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['type'] = None
        super(Endpoint, __self__).__init__(
            'azurerm:cdn/v20191231:Endpoint',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing Endpoint resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Endpoint(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop