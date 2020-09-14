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

__all__ = ['BigDataPool']


class BigDataPool(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_pause: Optional[pulumi.Input[pulumi.InputType['AutoPausePropertiesArgs']]] = None,
                 auto_scale: Optional[pulumi.Input[pulumi.InputType['AutoScalePropertiesArgs']]] = None,
                 big_data_pool_name: Optional[pulumi.Input[str]] = None,
                 creation_date: Optional[pulumi.Input[str]] = None,
                 default_spark_log_folder: Optional[pulumi.Input[str]] = None,
                 force: Optional[pulumi.Input[bool]] = None,
                 library_requirements: Optional[pulumi.Input[pulumi.InputType['LibraryRequirementsArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 node_count: Optional[pulumi.Input[float]] = None,
                 node_size: Optional[pulumi.Input[str]] = None,
                 node_size_family: Optional[pulumi.Input[str]] = None,
                 provisioning_state: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 spark_events_folder: Optional[pulumi.Input[str]] = None,
                 spark_version: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 workspace_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        A Big Data pool

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['AutoPausePropertiesArgs']] auto_pause: Auto-pausing properties
        :param pulumi.Input[pulumi.InputType['AutoScalePropertiesArgs']] auto_scale: Auto-scaling properties
        :param pulumi.Input[str] big_data_pool_name: Big Data pool name
        :param pulumi.Input[str] creation_date: The time when the Big Data pool was created.
        :param pulumi.Input[str] default_spark_log_folder: The default folder where Spark logs will be written.
        :param pulumi.Input[bool] force: Whether to stop any running jobs in the Big Data pool
        :param pulumi.Input[pulumi.InputType['LibraryRequirementsArgs']] library_requirements: Library version requirements
        :param pulumi.Input[str] location: The geo-location where the resource lives
        :param pulumi.Input[float] node_count: The number of nodes in the Big Data pool.
        :param pulumi.Input[str] node_size: The level of compute power that each node in the Big Data pool has.
        :param pulumi.Input[str] node_size_family: The kind of nodes that the Big Data pool provides.
        :param pulumi.Input[str] provisioning_state: The state of the Big Data pool.
        :param pulumi.Input[str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[str] spark_events_folder: The Spark events folder
        :param pulumi.Input[str] spark_version: The Apache Spark version.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags.
        :param pulumi.Input[str] workspace_name: The name of the workspace
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

            __props__['auto_pause'] = auto_pause
            __props__['auto_scale'] = auto_scale
            if big_data_pool_name is None:
                raise TypeError("Missing required property 'big_data_pool_name'")
            __props__['big_data_pool_name'] = big_data_pool_name
            __props__['creation_date'] = creation_date
            __props__['default_spark_log_folder'] = default_spark_log_folder
            __props__['force'] = force
            __props__['library_requirements'] = library_requirements
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            __props__['node_count'] = node_count
            __props__['node_size'] = node_size
            __props__['node_size_family'] = node_size_family
            __props__['provisioning_state'] = provisioning_state
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['spark_events_folder'] = spark_events_folder
            __props__['spark_version'] = spark_version
            __props__['tags'] = tags
            if workspace_name is None:
                raise TypeError("Missing required property 'workspace_name'")
            __props__['workspace_name'] = workspace_name
            __props__['name'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:synapse/latest:BigDataPool")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(BigDataPool, __self__).__init__(
            'azurerm:synapse/v20190601preview:BigDataPool',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'BigDataPool':
        """
        Get an existing BigDataPool resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return BigDataPool(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoPause")
    def auto_pause(self) -> pulumi.Output[Optional['outputs.AutoPausePropertiesResponse']]:
        """
        Auto-pausing properties
        """
        return pulumi.get(self, "auto_pause")

    @property
    @pulumi.getter(name="autoScale")
    def auto_scale(self) -> pulumi.Output[Optional['outputs.AutoScalePropertiesResponse']]:
        """
        Auto-scaling properties
        """
        return pulumi.get(self, "auto_scale")

    @property
    @pulumi.getter(name="creationDate")
    def creation_date(self) -> pulumi.Output[Optional[str]]:
        """
        The time when the Big Data pool was created.
        """
        return pulumi.get(self, "creation_date")

    @property
    @pulumi.getter(name="defaultSparkLogFolder")
    def default_spark_log_folder(self) -> pulumi.Output[Optional[str]]:
        """
        The default folder where Spark logs will be written.
        """
        return pulumi.get(self, "default_spark_log_folder")

    @property
    @pulumi.getter(name="libraryRequirements")
    def library_requirements(self) -> pulumi.Output[Optional['outputs.LibraryRequirementsResponse']]:
        """
        Library version requirements
        """
        return pulumi.get(self, "library_requirements")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nodeCount")
    def node_count(self) -> pulumi.Output[Optional[float]]:
        """
        The number of nodes in the Big Data pool.
        """
        return pulumi.get(self, "node_count")

    @property
    @pulumi.getter(name="nodeSize")
    def node_size(self) -> pulumi.Output[Optional[str]]:
        """
        The level of compute power that each node in the Big Data pool has.
        """
        return pulumi.get(self, "node_size")

    @property
    @pulumi.getter(name="nodeSizeFamily")
    def node_size_family(self) -> pulumi.Output[Optional[str]]:
        """
        The kind of nodes that the Big Data pool provides.
        """
        return pulumi.get(self, "node_size_family")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[Optional[str]]:
        """
        The state of the Big Data pool.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="sparkEventsFolder")
    def spark_events_folder(self) -> pulumi.Output[Optional[str]]:
        """
        The Spark events folder
        """
        return pulumi.get(self, "spark_events_folder")

    @property
    @pulumi.getter(name="sparkVersion")
    def spark_version(self) -> pulumi.Output[Optional[str]]:
        """
        The Apache Spark version.
        """
        return pulumi.get(self, "spark_version")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

