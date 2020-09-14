# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = ['WorkloadClassifier']


class WorkloadClassifier(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 context: Optional[pulumi.Input[str]] = None,
                 database_name: Optional[pulumi.Input[str]] = None,
                 end_time: Optional[pulumi.Input[str]] = None,
                 importance: Optional[pulumi.Input[str]] = None,
                 label: Optional[pulumi.Input[str]] = None,
                 member_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 server_name: Optional[pulumi.Input[str]] = None,
                 start_time: Optional[pulumi.Input[str]] = None,
                 workload_classifier_name: Optional[pulumi.Input[str]] = None,
                 workload_group_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Workload classifier operations for a data warehouse

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] context: The workload classifier context.
        :param pulumi.Input[str] database_name: The name of the database.
        :param pulumi.Input[str] end_time: The workload classifier end time for classification.
        :param pulumi.Input[str] importance: The workload classifier importance.
        :param pulumi.Input[str] label: The workload classifier label.
        :param pulumi.Input[str] member_name: The workload classifier member name.
        :param pulumi.Input[str] resource_group_name: The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        :param pulumi.Input[str] server_name: The name of the server.
        :param pulumi.Input[str] start_time: The workload classifier start time for classification.
        :param pulumi.Input[str] workload_classifier_name: The name of the workload classifier to create/update.
        :param pulumi.Input[str] workload_group_name: The name of the workload group from which to receive the classifier from.
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

            __props__['context'] = context
            if database_name is None:
                raise TypeError("Missing required property 'database_name'")
            __props__['database_name'] = database_name
            __props__['end_time'] = end_time
            __props__['importance'] = importance
            __props__['label'] = label
            if member_name is None:
                raise TypeError("Missing required property 'member_name'")
            __props__['member_name'] = member_name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if server_name is None:
                raise TypeError("Missing required property 'server_name'")
            __props__['server_name'] = server_name
            __props__['start_time'] = start_time
            if workload_classifier_name is None:
                raise TypeError("Missing required property 'workload_classifier_name'")
            __props__['workload_classifier_name'] = workload_classifier_name
            if workload_group_name is None:
                raise TypeError("Missing required property 'workload_group_name'")
            __props__['workload_group_name'] = workload_group_name
            __props__['name'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:sql/latest:WorkloadClassifier")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(WorkloadClassifier, __self__).__init__(
            'azurerm:sql/v20190601preview:WorkloadClassifier',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'WorkloadClassifier':
        """
        Get an existing WorkloadClassifier resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return WorkloadClassifier(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def context(self) -> pulumi.Output[Optional[str]]:
        """
        The workload classifier context.
        """
        return pulumi.get(self, "context")

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> pulumi.Output[Optional[str]]:
        """
        The workload classifier end time for classification.
        """
        return pulumi.get(self, "end_time")

    @property
    @pulumi.getter
    def importance(self) -> pulumi.Output[Optional[str]]:
        """
        The workload classifier importance.
        """
        return pulumi.get(self, "importance")

    @property
    @pulumi.getter
    def label(self) -> pulumi.Output[Optional[str]]:
        """
        The workload classifier label.
        """
        return pulumi.get(self, "label")

    @property
    @pulumi.getter(name="memberName")
    def member_name(self) -> pulumi.Output[str]:
        """
        The workload classifier member name.
        """
        return pulumi.get(self, "member_name")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> pulumi.Output[Optional[str]]:
        """
        The workload classifier start time for classification.
        """
        return pulumi.get(self, "start_time")

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

