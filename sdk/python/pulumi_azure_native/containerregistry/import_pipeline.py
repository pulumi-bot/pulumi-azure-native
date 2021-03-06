# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._enums import *
from ._inputs import *

__all__ = ['ImportPipeline']


class ImportPipeline(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 identity: Optional[pulumi.Input[pulumi.InputType['IdentityPropertiesArgs']]] = None,
                 import_pipeline_name: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 options: Optional[pulumi.Input[Sequence[pulumi.Input[Union[str, 'PipelineOptions']]]]] = None,
                 registry_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[pulumi.InputType['ImportPipelineSourcePropertiesArgs']]] = None,
                 trigger: Optional[pulumi.Input[pulumi.InputType['PipelineTriggerPropertiesArgs']]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        An object that represents an import pipeline for a container registry.
        API Version: 2020-11-01-preview.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['IdentityPropertiesArgs']] identity: The identity of the import pipeline.
        :param pulumi.Input[str] import_pipeline_name: The name of the import pipeline.
        :param pulumi.Input[str] location: The location of the import pipeline.
        :param pulumi.Input[Sequence[pulumi.Input[Union[str, 'PipelineOptions']]]] options: The list of all options configured for the pipeline.
        :param pulumi.Input[str] registry_name: The name of the container registry.
        :param pulumi.Input[str] resource_group_name: The name of the resource group to which the container registry belongs.
        :param pulumi.Input[pulumi.InputType['ImportPipelineSourcePropertiesArgs']] source: The source properties of the import pipeline.
        :param pulumi.Input[pulumi.InputType['PipelineTriggerPropertiesArgs']] trigger: The properties that describe the trigger of the import pipeline.
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

            __props__['identity'] = identity
            __props__['import_pipeline_name'] = import_pipeline_name
            __props__['location'] = location
            __props__['options'] = options
            if registry_name is None and not opts.urn:
                raise TypeError("Missing required property 'registry_name'")
            __props__['registry_name'] = registry_name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if source is None and not opts.urn:
                raise TypeError("Missing required property 'source'")
            __props__['source'] = source
            __props__['trigger'] = trigger
            __props__['name'] = None
            __props__['provisioning_state'] = None
            __props__['system_data'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:containerregistry:ImportPipeline"), pulumi.Alias(type_="azure-native:containerregistry/v20191201preview:ImportPipeline"), pulumi.Alias(type_="azure-nextgen:containerregistry/v20191201preview:ImportPipeline"), pulumi.Alias(type_="azure-native:containerregistry/v20201101preview:ImportPipeline"), pulumi.Alias(type_="azure-nextgen:containerregistry/v20201101preview:ImportPipeline")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(ImportPipeline, __self__).__init__(
            'azure-native:containerregistry:ImportPipeline',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ImportPipeline':
        """
        Get an existing ImportPipeline resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["identity"] = None
        __props__["location"] = None
        __props__["name"] = None
        __props__["options"] = None
        __props__["provisioning_state"] = None
        __props__["source"] = None
        __props__["system_data"] = None
        __props__["trigger"] = None
        __props__["type"] = None
        return ImportPipeline(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def identity(self) -> pulumi.Output[Optional['outputs.IdentityPropertiesResponse']]:
        """
        The identity of the import pipeline.
        """
        return pulumi.get(self, "identity")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[Optional[str]]:
        """
        The location of the import pipeline.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def options(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The list of all options configured for the pipeline.
        """
        return pulumi.get(self, "options")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        The provisioning state of the pipeline at the time the operation was called.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def source(self) -> pulumi.Output['outputs.ImportPipelineSourcePropertiesResponse']:
        """
        The source properties of the import pipeline.
        """
        return pulumi.get(self, "source")

    @property
    @pulumi.getter(name="systemData")
    def system_data(self) -> pulumi.Output['outputs.SystemDataResponse']:
        """
        Metadata pertaining to creation and last modification of the resource.
        """
        return pulumi.get(self, "system_data")

    @property
    @pulumi.getter
    def trigger(self) -> pulumi.Output[Optional['outputs.PipelineTriggerPropertiesResponse']]:
        """
        The properties that describe the trigger of the import pipeline.
        """
        return pulumi.get(self, "trigger")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of the resource.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

