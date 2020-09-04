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

__all__ = ['Project']


class Project(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 databases_info: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['DatabaseInfoArgs']]]]] = None,
                 group_name: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 source_connection_info: Optional[pulumi.Input[pulumi.InputType['ConnectionInfoArgs']]] = None,
                 source_platform: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 target_connection_info: Optional[pulumi.Input[pulumi.InputType['ConnectionInfoArgs']]] = None,
                 target_platform: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        A project resource

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['DatabaseInfoArgs']]]] databases_info: List of DatabaseInfo
        :param pulumi.Input[str] group_name: Name of the resource group
        :param pulumi.Input[str] location: Resource location.
        :param pulumi.Input[str] project_name: Name of the project
        :param pulumi.Input[str] service_name: Name of the service
        :param pulumi.Input[pulumi.InputType['ConnectionInfoArgs']] source_connection_info: Information for connecting to source
        :param pulumi.Input[str] source_platform: Source platform for the project
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags.
        :param pulumi.Input[pulumi.InputType['ConnectionInfoArgs']] target_connection_info: Information for connecting to target
        :param pulumi.Input[str] target_platform: Target platform for the project
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

            __props__['databases_info'] = databases_info
            if group_name is None:
                raise TypeError("Missing required property 'group_name'")
            __props__['group_name'] = group_name
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            if project_name is None:
                raise TypeError("Missing required property 'project_name'")
            __props__['project_name'] = project_name
            if service_name is None:
                raise TypeError("Missing required property 'service_name'")
            __props__['service_name'] = service_name
            __props__['source_connection_info'] = source_connection_info
            if source_platform is None:
                raise TypeError("Missing required property 'source_platform'")
            __props__['source_platform'] = source_platform
            __props__['tags'] = tags
            __props__['target_connection_info'] = target_connection_info
            if target_platform is None:
                raise TypeError("Missing required property 'target_platform'")
            __props__['target_platform'] = target_platform
            __props__['creation_time'] = None
            __props__['name'] = None
            __props__['provisioning_state'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:datamigration/latest:Project"), pulumi.Alias(type_="azurerm:datamigration/v20171115preview:Project"), pulumi.Alias(type_="azurerm:datamigration/v20180315preview:Project"), pulumi.Alias(type_="azurerm:datamigration/v20180331preview:Project"), pulumi.Alias(type_="azurerm:datamigration/v20180419:Project")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Project, __self__).__init__(
            'azurerm:datamigration/v20180715preview:Project',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Project':
        """
        Get an existing Project resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Project(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> pulumi.Output[str]:
        """
        UTC Date and time when project was created
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter(name="databasesInfo")
    def databases_info(self) -> pulumi.Output[Optional[List['outputs.DatabaseInfoResponse']]]:
        """
        List of DatabaseInfo
        """
        return pulumi.get(self, "databases_info")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Resource location.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        The project's provisioning state
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="sourceConnectionInfo")
    def source_connection_info(self) -> pulumi.Output[Optional['outputs.ConnectionInfoResponse']]:
        """
        Information for connecting to source
        """
        return pulumi.get(self, "source_connection_info")

    @property
    @pulumi.getter(name="sourcePlatform")
    def source_platform(self) -> pulumi.Output[str]:
        """
        Source platform for the project
        """
        return pulumi.get(self, "source_platform")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="targetConnectionInfo")
    def target_connection_info(self) -> pulumi.Output[Optional['outputs.ConnectionInfoResponse']]:
        """
        Information for connecting to target
        """
        return pulumi.get(self, "target_connection_info")

    @property
    @pulumi.getter(name="targetPlatform")
    def target_platform(self) -> pulumi.Output[str]:
        """
        Target platform for the project
        """
        return pulumi.get(self, "target_platform")

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

