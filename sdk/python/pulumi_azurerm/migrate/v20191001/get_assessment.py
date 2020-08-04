# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class GetAssessmentResult:
    """
    An assessment created for a group in the Migration project.
    """
    def __init__(__self__, e_tag=None, name=None, properties=None, type=None):
        if e_tag and not isinstance(e_tag, str):
            raise TypeError("Expected argument 'e_tag' to be a str")
        __self__.e_tag = e_tag
        """
        For optimistic concurrency control.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        Unique name of an assessment.
        """
        if properties and not isinstance(properties, dict):
            raise TypeError("Expected argument 'properties' to be a dict")
        __self__.properties = properties
        """
        Properties of the assessment.
        """
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        __self__.type = type
        """
        Type of the object = [Microsoft.Migrate/assessmentProjects/groups/assessments].
        """


class AwaitableGetAssessmentResult(GetAssessmentResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAssessmentResult(
            e_tag=self.e_tag,
            name=self.name,
            properties=self.properties,
            type=self.type)


def get_assessment(group_name=None, name=None, project_name=None, resource_group_name=None, opts=None):
    """
    Use this data source to access information about an existing resource.

    :param str group_name: Unique name of a group within a project.
    :param str name: Unique name of an assessment within a project.
    :param str project_name: Name of the Azure Migrate project.
    :param str resource_group_name: Name of the Azure Resource Group that project is part of.
    """
    __args__ = dict()
    __args__['groupName'] = group_name
    __args__['name'] = name
    __args__['projectName'] = project_name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:migrate/v20191001:getAssessment', __args__, opts=opts).value

    return AwaitableGetAssessmentResult(
        e_tag=__ret__.get('eTag'),
        name=__ret__.get('name'),
        properties=__ret__.get('properties'),
        type=__ret__.get('type'))