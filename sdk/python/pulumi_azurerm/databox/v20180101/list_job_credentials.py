# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables
from . import outputs

__all__ = [
    'ListJobCredentialsResult',
    'AwaitableListJobCredentialsResult',
    'list_job_credentials',
]

@pulumi.output_type
class ListJobCredentialsResult:
    """
    List of unencrypted credentials for accessing device.
    """
    def __init__(__self__, next_link=None, value=None):
        if next_link and not isinstance(next_link, str):
            raise TypeError("Expected argument 'next_link' to be a str")
        pulumi.set(__self__, "next_link", next_link)
        if value and not isinstance(value, list):
            raise TypeError("Expected argument 'value' to be a list")
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter(name="nextLink")
    def next_link(self) -> Optional[str]:
        """
        Link for the next set of unencrypted credentials.
        """
        return pulumi.get(self, "next_link")

    @property
    @pulumi.getter
    def value(self) -> Optional[List['outputs.UnencryptedCredentialsResponseResult']]:
        """
        List of unencrypted credentials.
        """
        return pulumi.get(self, "value")


class AwaitableListJobCredentialsResult(ListJobCredentialsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListJobCredentialsResult(
            next_link=self.next_link,
            value=self.value)


def list_job_credentials(job_name: Optional[str] = None,
                         resource_group_name: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListJobCredentialsResult:
    """
    Use this data source to access information about an existing resource.

    :param str job_name: The name of the job Resource within the specified resource group. job names must be between 3 and 24 characters in length and use any alphanumeric and underscore only
    :param str resource_group_name: The Resource Group Name
    """
    __args__ = dict()
    __args__['jobName'] = job_name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:databox/v20180101:listJobCredentials', __args__, opts=opts, typ=ListJobCredentialsResult).value

    return AwaitableListJobCredentialsResult(
        next_link=__ret__.next_link,
        value=__ret__.value)
