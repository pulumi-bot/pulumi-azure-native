# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class Profile(pulumi.CustomResource):
    location: pulumi.Output[str]
    """
    The Azure Region where the resource lives
    """
    name: pulumi.Output[str]
    """
    The name of the resource
    """
    properties: pulumi.Output[dict]
    """
    The properties of the Traffic Manager profile.
      * `dns_config` (`dict`) - The DNS settings of the Traffic Manager profile.
        * `fqdn` (`str`) - The fully-qualified domain name (FQDN) of the Traffic Manager profile. This is formed from the concatenation of the RelativeName with the DNS domain used by Azure Traffic Manager.
        * `relative_name` (`str`) - The relative DNS name provided by this Traffic Manager profile. This value is combined with the DNS domain name used by Azure Traffic Manager to form the fully-qualified domain name (FQDN) of the profile.
        * `ttl` (`float`) - The DNS Time-To-Live (TTL), in seconds. This informs the local DNS resolvers and DNS clients how long to cache DNS responses provided by this Traffic Manager profile.

      * `endpoints` (`list`) - The list of endpoints in the Traffic Manager profile.
        * `id` (`str`) - Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficManagerProfiles/{resourceName}
        * `name` (`str`) - The name of the resource
        * `properties` (`dict`) - The properties of the Traffic Manager endpoint.
          * `custom_headers` (`list`) - List of custom headers.
            * `name` (`str`) - Header name.
            * `value` (`str`) - Header value.

          * `endpoint_location` (`str`) - Specifies the location of the external or nested endpoints when using the ‘Performance’ traffic routing method.
          * `endpoint_monitor_status` (`str`) - The monitoring status of the endpoint.
          * `endpoint_status` (`str`) - The status of the endpoint. If the endpoint is Enabled, it is probed for endpoint health and is included in the traffic routing method.
          * `geo_mapping` (`list`) - The list of countries/regions mapped to this endpoint when using the ‘Geographic’ traffic routing method. Please consult Traffic Manager Geographic documentation for a full list of accepted values.
          * `min_child_endpoints` (`float`) - The minimum number of endpoints that must be available in the child profile in order for the parent profile to be considered available. Only applicable to endpoint of type 'NestedEndpoints'.
          * `priority` (`float`) - The priority of this endpoint when using the ‘Priority’ traffic routing method. Possible values are from 1 to 1000, lower values represent higher priority. This is an optional parameter.  If specified, it must be specified on all endpoints, and no two endpoints can share the same priority value.
          * `target` (`str`) - The fully-qualified DNS name or IP address of the endpoint. Traffic Manager returns this value in DNS responses to direct traffic to this endpoint.
          * `target_resource_id` (`str`) - The Azure Resource URI of the of the endpoint. Not applicable to endpoints of type 'ExternalEndpoints'.
          * `weight` (`float`) - The weight of this endpoint when using the 'Weighted' traffic routing method. Possible values are from 1 to 1000.

        * `type` (`str`) - The type of the resource. Ex- Microsoft.Network/trafficManagerProfiles.

      * `monitor_config` (`dict`) - The endpoint monitoring settings of the Traffic Manager profile.
        * `custom_headers` (`list`) - List of custom headers.
          * `name` (`str`) - Header name.
          * `value` (`str`) - Header value.

        * `expected_status_code_ranges` (`list`) - List of expected status code ranges.
        * `interval_in_seconds` (`float`) - The monitor interval for endpoints in this profile. This is the interval at which Traffic Manager will check the health of each endpoint in this profile.
        * `path` (`str`) - The path relative to the endpoint domain name used to probe for endpoint health.
        * `port` (`float`) - The TCP port used to probe for endpoint health.
        * `profile_monitor_status` (`str`) - The profile-level monitoring status of the Traffic Manager profile.
        * `protocol` (`str`) - The protocol (HTTP, HTTPS or TCP) used to probe for endpoint health.
        * `timeout_in_seconds` (`float`) - The monitor timeout for endpoints in this profile. This is the time that Traffic Manager allows endpoints in this profile to response to the health check.
        * `tolerated_number_of_failures` (`float`) - The number of consecutive failed health check that Traffic Manager tolerates before declaring an endpoint in this profile Degraded after the next failed health check.

      * `profile_status` (`str`) - The status of the Traffic Manager profile.
      * `traffic_routing_method` (`str`) - The traffic routing method of the Traffic Manager profile.
      * `traffic_view_enrollment_status` (`str`) - Indicates whether Traffic View is 'Enabled' or 'Disabled' for the Traffic Manager profile. Null, indicates 'Disabled'. Enabling this feature will increase the cost of the Traffic Manage profile.
    """
    tags: pulumi.Output[dict]
    """
    Resource tags.
    """
    type: pulumi.Output[str]
    """
    The type of the resource. Ex- Microsoft.Network/trafficManagerProfiles.
    """
    def __init__(__self__, resource_name, opts=None, id=None, location=None, name=None, properties=None, resource_group_name=None, tags=None, type=None, __props__=None, __name__=None, __opts__=None):
        """
        Class representing a Traffic Manager profile.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] id: Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficManagerProfiles/{resourceName}
        :param pulumi.Input[str] location: The Azure Region where the resource lives
        :param pulumi.Input[str] name: The name of the Traffic Manager profile.
        :param pulumi.Input[dict] properties: The properties of the Traffic Manager profile.
        :param pulumi.Input[str] resource_group_name: The name of the resource group containing the Traffic Manager profile.
        :param pulumi.Input[dict] tags: Resource tags.
        :param pulumi.Input[str] type: The type of the resource. Ex- Microsoft.Network/trafficManagerProfiles.

        The **properties** object supports the following:

          * `dns_config` (`pulumi.Input[dict]`) - The DNS settings of the Traffic Manager profile.
            * `relative_name` (`pulumi.Input[str]`) - The relative DNS name provided by this Traffic Manager profile. This value is combined with the DNS domain name used by Azure Traffic Manager to form the fully-qualified domain name (FQDN) of the profile.
            * `ttl` (`pulumi.Input[float]`) - The DNS Time-To-Live (TTL), in seconds. This informs the local DNS resolvers and DNS clients how long to cache DNS responses provided by this Traffic Manager profile.

          * `endpoints` (`pulumi.Input[list]`) - The list of endpoints in the Traffic Manager profile.
            * `name` (`pulumi.Input[str]`) - The name of the resource
            * `properties` (`pulumi.Input[dict]`) - The properties of the Traffic Manager endpoint.
              * `custom_headers` (`pulumi.Input[list]`) - List of custom headers.
                * `name` (`pulumi.Input[str]`) - Header name.
                * `value` (`pulumi.Input[str]`) - Header value.

              * `endpoint_location` (`pulumi.Input[str]`) - Specifies the location of the external or nested endpoints when using the ‘Performance’ traffic routing method.
              * `endpoint_monitor_status` (`pulumi.Input[str]`) - The monitoring status of the endpoint.
              * `endpoint_status` (`pulumi.Input[str]`) - The status of the endpoint. If the endpoint is Enabled, it is probed for endpoint health and is included in the traffic routing method.
              * `geo_mapping` (`pulumi.Input[list]`) - The list of countries/regions mapped to this endpoint when using the ‘Geographic’ traffic routing method. Please consult Traffic Manager Geographic documentation for a full list of accepted values.
              * `min_child_endpoints` (`pulumi.Input[float]`) - The minimum number of endpoints that must be available in the child profile in order for the parent profile to be considered available. Only applicable to endpoint of type 'NestedEndpoints'.
              * `priority` (`pulumi.Input[float]`) - The priority of this endpoint when using the ‘Priority’ traffic routing method. Possible values are from 1 to 1000, lower values represent higher priority. This is an optional parameter.  If specified, it must be specified on all endpoints, and no two endpoints can share the same priority value.
              * `target` (`pulumi.Input[str]`) - The fully-qualified DNS name or IP address of the endpoint. Traffic Manager returns this value in DNS responses to direct traffic to this endpoint.
              * `target_resource_id` (`pulumi.Input[str]`) - The Azure Resource URI of the of the endpoint. Not applicable to endpoints of type 'ExternalEndpoints'.
              * `weight` (`pulumi.Input[float]`) - The weight of this endpoint when using the 'Weighted' traffic routing method. Possible values are from 1 to 1000.

            * `type` (`pulumi.Input[str]`) - The type of the resource. Ex- Microsoft.Network/trafficManagerProfiles.

          * `monitor_config` (`pulumi.Input[dict]`) - The endpoint monitoring settings of the Traffic Manager profile.
            * `custom_headers` (`pulumi.Input[list]`) - List of custom headers.
              * `name` (`pulumi.Input[str]`) - Header name.
              * `value` (`pulumi.Input[str]`) - Header value.

            * `expected_status_code_ranges` (`pulumi.Input[list]`) - List of expected status code ranges.
            * `interval_in_seconds` (`pulumi.Input[float]`) - The monitor interval for endpoints in this profile. This is the interval at which Traffic Manager will check the health of each endpoint in this profile.
            * `path` (`pulumi.Input[str]`) - The path relative to the endpoint domain name used to probe for endpoint health.
            * `port` (`pulumi.Input[float]`) - The TCP port used to probe for endpoint health.
            * `profile_monitor_status` (`pulumi.Input[str]`) - The profile-level monitoring status of the Traffic Manager profile.
            * `protocol` (`pulumi.Input[str]`) - The protocol (HTTP, HTTPS or TCP) used to probe for endpoint health.
            * `timeout_in_seconds` (`pulumi.Input[float]`) - The monitor timeout for endpoints in this profile. This is the time that Traffic Manager allows endpoints in this profile to response to the health check.
            * `tolerated_number_of_failures` (`pulumi.Input[float]`) - The number of consecutive failed health check that Traffic Manager tolerates before declaring an endpoint in this profile Degraded after the next failed health check.

          * `profile_status` (`pulumi.Input[str]`) - The status of the Traffic Manager profile.
          * `traffic_routing_method` (`pulumi.Input[str]`) - The traffic routing method of the Traffic Manager profile.
          * `traffic_view_enrollment_status` (`pulumi.Input[str]`) - Indicates whether Traffic View is 'Enabled' or 'Disabled' for the Traffic Manager profile. Null, indicates 'Disabled'. Enabling this feature will increase the cost of the Traffic Manage profile.
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
            __props__['type'] = type
        super(Profile, __self__).__init__(
            'azurerm:network/v20180301:Profile',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing Profile resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Profile(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
