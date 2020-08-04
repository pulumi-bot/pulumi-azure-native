# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class Endpoint(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    The name of the resource
    """
    properties: pulumi.Output[dict]
    """
    The properties of the Traffic Manager endpoint.
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
    """
    type: pulumi.Output[str]
    """
    The type of the resource. Ex- Microsoft.Network/trafficManagerProfiles.
    """
    def __init__(__self__, resource_name, opts=None, endpoint_type=None, id=None, name=None, profile_name=None, properties=None, resource_group_name=None, type=None, __props__=None, __name__=None, __opts__=None):
        """
        Class representing a Traffic Manager endpoint.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] endpoint_type: The type of the Traffic Manager endpoint to be created or updated.
        :param pulumi.Input[str] id: Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficManagerProfiles/{resourceName}
        :param pulumi.Input[str] name: The name of the Traffic Manager endpoint to be created or updated.
        :param pulumi.Input[str] profile_name: The name of the Traffic Manager profile.
        :param pulumi.Input[dict] properties: The properties of the Traffic Manager endpoint.
        :param pulumi.Input[str] resource_group_name: The name of the resource group containing the Traffic Manager endpoint to be created or updated.
        :param pulumi.Input[str] type: The type of the resource. Ex- Microsoft.Network/trafficManagerProfiles.

        The **properties** object supports the following:

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

            if endpoint_type is None:
                raise TypeError("Missing required property 'endpoint_type'")
            __props__['endpoint_type'] = endpoint_type
            __props__['id'] = id
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
            __props__['type'] = type
        super(Endpoint, __self__).__init__(
            'azurerm:network/v20180301:Endpoint',
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
