# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables

__all__ = ['TrafficManagerUserMetricsKey']

warnings.warn("""The 'latest' version is deprecated. Please migrate to the resource in the top-level module: 'azure-native:network:TrafficManagerUserMetricsKey'.""", DeprecationWarning)


class TrafficManagerUserMetricsKey(pulumi.CustomResource):
    warnings.warn("""The 'latest' version is deprecated. Please migrate to the resource in the top-level module: 'azure-native:network:TrafficManagerUserMetricsKey'.""", DeprecationWarning)

    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Class representing Traffic Manager User Metrics.
        Latest API Version: 2018-04-01.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        pulumi.log.warn("""TrafficManagerUserMetricsKey is deprecated: The 'latest' version is deprecated. Please migrate to the resource in the top-level module: 'azure-native:network:TrafficManagerUserMetricsKey'.""")
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

            __props__['key'] = None
            __props__['name'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:network/latest:TrafficManagerUserMetricsKey"), pulumi.Alias(type_="azure-native:network:TrafficManagerUserMetricsKey"), pulumi.Alias(type_="azure-nextgen:network:TrafficManagerUserMetricsKey"), pulumi.Alias(type_="azure-native:network/v20180401:TrafficManagerUserMetricsKey"), pulumi.Alias(type_="azure-nextgen:network/v20180401:TrafficManagerUserMetricsKey")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(TrafficManagerUserMetricsKey, __self__).__init__(
            'azure-native:network/latest:TrafficManagerUserMetricsKey',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'TrafficManagerUserMetricsKey':
        """
        Get an existing TrafficManagerUserMetricsKey resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["key"] = None
        __props__["name"] = None
        __props__["type"] = None
        return TrafficManagerUserMetricsKey(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Output[Optional[str]]:
        """
        The key returned by the User Metrics operation.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[Optional[str]]:
        """
        The type of the resource. Ex- Microsoft.Network/trafficManagerProfiles.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

