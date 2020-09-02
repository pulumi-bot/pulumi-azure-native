# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'ImageTemplateCustomizerArgs',
    'ImageTemplateDistributorArgs',
    'ImageTemplateSourceArgs',
]

@pulumi.input_type
class ImageTemplateCustomizerArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] type: The type of customization tool you want to use on the Image. For example, "shell" can be shellCustomizer
        :param pulumi.Input[str] name: Friendly Name to provide context on what this customization step does
        """
        pulumi.set(__self__, "type", type)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The type of customization tool you want to use on the Image. For example, "shell" can be shellCustomizer
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Friendly Name to provide context on what this customization step does
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class ImageTemplateDistributorArgs:
    def __init__(__self__, *,
                 run_output_name: pulumi.Input[str],
                 type: pulumi.Input[str],
                 artifact_tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Generic distribution object
        :param pulumi.Input[str] run_output_name: The name to be used for the associated RunOutput.
        :param pulumi.Input[str] type: Type of distribution.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] artifact_tags: Tags that will be applied to the artifact once it has been created/updated by the distributor.
        """
        pulumi.set(__self__, "run_output_name", run_output_name)
        pulumi.set(__self__, "type", type)
        if artifact_tags is not None:
            pulumi.set(__self__, "artifact_tags", artifact_tags)

    @property
    @pulumi.getter(name="runOutputName")
    def run_output_name(self) -> pulumi.Input[str]:
        """
        The name to be used for the associated RunOutput.
        """
        return pulumi.get(self, "run_output_name")

    @run_output_name.setter
    def run_output_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "run_output_name", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Type of distribution.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="artifactTags")
    def artifact_tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Tags that will be applied to the artifact once it has been created/updated by the distributor.
        """
        return pulumi.get(self, "artifact_tags")

    @artifact_tags.setter
    def artifact_tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "artifact_tags", value)


@pulumi.input_type
class ImageTemplateSourceArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str]):
        """
        :param pulumi.Input[str] type: Specifies the type of source image you want to start with.
        """
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Specifies the type of source image you want to start with.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)


