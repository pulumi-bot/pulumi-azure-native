# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables

__all__ = ['WorkloadGroup']


class WorkloadGroup(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 database_name: Optional[pulumi.Input[str]] = None,
                 importance: Optional[pulumi.Input[str]] = None,
                 max_resource_percent: Optional[pulumi.Input[int]] = None,
                 max_resource_percent_per_request: Optional[pulumi.Input[float]] = None,
                 min_resource_percent: Optional[pulumi.Input[int]] = None,
                 min_resource_percent_per_request: Optional[pulumi.Input[float]] = None,
                 query_execution_timeout: Optional[pulumi.Input[int]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 server_name: Optional[pulumi.Input[str]] = None,
                 workload_group_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Workload group operations for a data warehouse

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] database_name: The name of the database.
        :param pulumi.Input[str] importance: The workload group importance level.
        :param pulumi.Input[int] max_resource_percent: The workload group cap percentage resource.
        :param pulumi.Input[float] max_resource_percent_per_request: The workload group request maximum grant percentage.
        :param pulumi.Input[int] min_resource_percent: The workload group minimum percentage resource.
        :param pulumi.Input[float] min_resource_percent_per_request: The workload group request minimum grant percentage.
        :param pulumi.Input[int] query_execution_timeout: The workload group query execution timeout.
        :param pulumi.Input[str] resource_group_name: The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        :param pulumi.Input[str] server_name: The name of the server.
        :param pulumi.Input[str] workload_group_name: The name of the workload group.
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

            if database_name is None and not opts.urn:
                raise TypeError("Missing required property 'database_name'")
            __props__['database_name'] = database_name
            __props__['importance'] = importance
            if max_resource_percent is None and not opts.urn:
                raise TypeError("Missing required property 'max_resource_percent'")
            __props__['max_resource_percent'] = max_resource_percent
            __props__['max_resource_percent_per_request'] = max_resource_percent_per_request
            if min_resource_percent is None and not opts.urn:
                raise TypeError("Missing required property 'min_resource_percent'")
            __props__['min_resource_percent'] = min_resource_percent
            if min_resource_percent_per_request is None and not opts.urn:
                raise TypeError("Missing required property 'min_resource_percent_per_request'")
            __props__['min_resource_percent_per_request'] = min_resource_percent_per_request
            __props__['query_execution_timeout'] = query_execution_timeout
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if server_name is None and not opts.urn:
                raise TypeError("Missing required property 'server_name'")
            __props__['server_name'] = server_name
            __props__['workload_group_name'] = workload_group_name
            __props__['name'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:sql/v20190601preview:WorkloadGroup"), pulumi.Alias(type_="azure-native:sql:WorkloadGroup"), pulumi.Alias(type_="azure-nextgen:sql:WorkloadGroup"), pulumi.Alias(type_="azure-native:sql/v20200202preview:WorkloadGroup"), pulumi.Alias(type_="azure-nextgen:sql/v20200202preview:WorkloadGroup"), pulumi.Alias(type_="azure-native:sql/v20200801preview:WorkloadGroup"), pulumi.Alias(type_="azure-nextgen:sql/v20200801preview:WorkloadGroup")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(WorkloadGroup, __self__).__init__(
            'azure-native:sql/v20190601preview:WorkloadGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'WorkloadGroup':
        """
        Get an existing WorkloadGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["importance"] = None
        __props__["max_resource_percent"] = None
        __props__["max_resource_percent_per_request"] = None
        __props__["min_resource_percent"] = None
        __props__["min_resource_percent_per_request"] = None
        __props__["name"] = None
        __props__["query_execution_timeout"] = None
        __props__["type"] = None
        return WorkloadGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def importance(self) -> pulumi.Output[Optional[str]]:
        """
        The workload group importance level.
        """
        return pulumi.get(self, "importance")

    @property
    @pulumi.getter(name="maxResourcePercent")
    def max_resource_percent(self) -> pulumi.Output[int]:
        """
        The workload group cap percentage resource.
        """
        return pulumi.get(self, "max_resource_percent")

    @property
    @pulumi.getter(name="maxResourcePercentPerRequest")
    def max_resource_percent_per_request(self) -> pulumi.Output[Optional[float]]:
        """
        The workload group request maximum grant percentage.
        """
        return pulumi.get(self, "max_resource_percent_per_request")

    @property
    @pulumi.getter(name="minResourcePercent")
    def min_resource_percent(self) -> pulumi.Output[int]:
        """
        The workload group minimum percentage resource.
        """
        return pulumi.get(self, "min_resource_percent")

    @property
    @pulumi.getter(name="minResourcePercentPerRequest")
    def min_resource_percent_per_request(self) -> pulumi.Output[float]:
        """
        The workload group request minimum grant percentage.
        """
        return pulumi.get(self, "min_resource_percent_per_request")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="queryExecutionTimeout")
    def query_execution_timeout(self) -> pulumi.Output[Optional[int]]:
        """
        The workload group query execution timeout.
        """
        return pulumi.get(self, "query_execution_timeout")

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

