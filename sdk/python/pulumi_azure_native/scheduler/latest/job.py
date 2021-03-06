# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from . import outputs
from ._enums import *
from ._inputs import *

__all__ = ['Job']

warnings.warn("""The 'latest' version is deprecated. Please migrate to the resource in the top-level module: 'azure-native:scheduler:Job'.""", DeprecationWarning)


class Job(pulumi.CustomResource):
    warnings.warn("""The 'latest' version is deprecated. Please migrate to the resource in the top-level module: 'azure-native:scheduler:Job'.""", DeprecationWarning)

    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 job_collection_name: Optional[pulumi.Input[str]] = None,
                 job_name: Optional[pulumi.Input[str]] = None,
                 properties: Optional[pulumi.Input[pulumi.InputType['JobPropertiesArgs']]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Latest API Version: 2016-03-01.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] job_collection_name: The job collection name.
        :param pulumi.Input[str] job_name: The job name.
        :param pulumi.Input[pulumi.InputType['JobPropertiesArgs']] properties: Gets or sets the job properties.
        :param pulumi.Input[str] resource_group_name: The resource group name.
        """
        pulumi.log.warn("""Job is deprecated: The 'latest' version is deprecated. Please migrate to the resource in the top-level module: 'azure-native:scheduler:Job'.""")
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

            if job_collection_name is None and not opts.urn:
                raise TypeError("Missing required property 'job_collection_name'")
            __props__['job_collection_name'] = job_collection_name
            __props__['job_name'] = job_name
            __props__['properties'] = properties
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['name'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:scheduler/latest:Job"), pulumi.Alias(type_="azure-native:scheduler:Job"), pulumi.Alias(type_="azure-nextgen:scheduler:Job"), pulumi.Alias(type_="azure-native:scheduler/v20140801preview:Job"), pulumi.Alias(type_="azure-nextgen:scheduler/v20140801preview:Job"), pulumi.Alias(type_="azure-native:scheduler/v20160101:Job"), pulumi.Alias(type_="azure-nextgen:scheduler/v20160101:Job"), pulumi.Alias(type_="azure-native:scheduler/v20160301:Job"), pulumi.Alias(type_="azure-nextgen:scheduler/v20160301:Job")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Job, __self__).__init__(
            'azure-native:scheduler/latest:Job',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Job':
        """
        Get an existing Job resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["name"] = None
        __props__["properties"] = None
        __props__["type"] = None
        return Job(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Gets the job resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def properties(self) -> pulumi.Output['outputs.JobPropertiesResponse']:
        """
        Gets or sets the job properties.
        """
        return pulumi.get(self, "properties")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Gets the job resource type.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

