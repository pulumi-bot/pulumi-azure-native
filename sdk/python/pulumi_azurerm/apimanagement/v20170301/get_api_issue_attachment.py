# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class GetApiIssueAttachmentResult:
    """
    Issue Attachment Contract details.
    """
    def __init__(__self__, name=None, properties=None, type=None):
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        Resource name.
        """
        if properties and not isinstance(properties, dict):
            raise TypeError("Expected argument 'properties' to be a dict")
        __self__.properties = properties
        """
        Properties of the Issue Attachment.
        """
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        __self__.type = type
        """
        Resource type for API Management resource.
        """


class AwaitableGetApiIssueAttachmentResult(GetApiIssueAttachmentResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetApiIssueAttachmentResult(
            name=self.name,
            properties=self.properties,
            type=self.type)


def get_api_issue_attachment(api_id=None, issue_id=None, name=None, resource_group_name=None, service_name=None, opts=None):
    """
    Use this data source to access information about an existing resource.

    :param str api_id: API identifier. Must be unique in the current API Management service instance.
    :param str issue_id: Issue identifier. Must be unique in the current API Management service instance.
    :param str name: Attachment identifier within an Issue. Must be unique in the current Issue.
    :param str resource_group_name: The name of the resource group.
    :param str service_name: The name of the API Management service.
    """
    __args__ = dict()
    __args__['apiId'] = api_id
    __args__['issueId'] = issue_id
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    __args__['serviceName'] = service_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:apimanagement/v20170301:getApiIssueAttachment', __args__, opts=opts).value

    return AwaitableGetApiIssueAttachmentResult(
        name=__ret__.get('name'),
        properties=__ret__.get('properties'),
        type=__ret__.get('type'))
