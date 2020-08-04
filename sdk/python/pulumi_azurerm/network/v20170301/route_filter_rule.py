# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class RouteFilterRule(pulumi.CustomResource):
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
    The name of the resource that is unique within a resource group. This name can be used to access the resource.
    """
    properties: pulumi.Output[dict]
    """
    Route Filter Rule Resource
      * `access` (`str`) - The access type of the rule. Valid values are: 'Allow', 'Deny'
      * `communities` (`list`) - The collection for bgp community values to filter on. e.g. ['12076:5010','12076:5020']
      * `provisioning_state` (`str`) - The provisioning state of the resource. Possible values are: 'Updating', 'Deleting', 'Succeeded' and 'Failed'.
      * `route_filter_rule_type` (`str`) - The rule type of the rule. Valid value is: 'Community'
    """
    tags: pulumi.Output[dict]
    """
    Resource tags.
    """
    def __init__(__self__, resource_name, opts=None, id=None, location=None, name=None, properties=None, resource_group_name=None, route_filter_name=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Route Filter Rule Resource

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] id: Resource ID.
        :param pulumi.Input[str] location: Resource location.
        :param pulumi.Input[str] name: The name of the route filter rule.
        :param pulumi.Input[dict] properties: Route Filter Rule Resource
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[str] route_filter_name: The name of the route filter.
        :param pulumi.Input[dict] tags: Resource tags.

        The **properties** object supports the following:

          * `access` (`pulumi.Input[str]`) - The access type of the rule. Valid values are: 'Allow', 'Deny'
          * `communities` (`pulumi.Input[list]`) - The collection for bgp community values to filter on. e.g. ['12076:5010','12076:5020']
          * `route_filter_rule_type` (`pulumi.Input[str]`) - The rule type of the rule. Valid value is: 'Community'
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
            if route_filter_name is None:
                raise TypeError("Missing required property 'route_filter_name'")
            __props__['route_filter_name'] = route_filter_name
            __props__['tags'] = tags
            __props__['etag'] = None
        super(RouteFilterRule, __self__).__init__(
            'azurerm:network/v20170301:RouteFilterRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing RouteFilterRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return RouteFilterRule(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
