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
    'ListBitLockerKeyResult',
    'AwaitableListBitLockerKeyResult',
    'list_bit_locker_key',
]

@pulumi.output_type
class ListBitLockerKeyResult:
    """
    GetBitLockerKeys response
    """
    def __init__(__self__, value=None):
        if value and not isinstance(value, list):
            raise TypeError("Expected argument 'value' to be a list")
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[List['outputs.DriveBitLockerKeyResponseResult']]:
        """
        drive status
        """
        return pulumi.get(self, "value")


class AwaitableListBitLockerKeyResult(ListBitLockerKeyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListBitLockerKeyResult(
            value=self.value)


def list_bit_locker_key(job_name: Optional[str] = None,
                        resource_group_name: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListBitLockerKeyResult:
    """
    Use this data source to access information about an existing resource.

    :param str job_name: The name of the import/export job.
    :param str resource_group_name: The resource group name uniquely identifies the resource group within the user subscription.
    """
    __args__ = dict()
    __args__['jobName'] = job_name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:importexport/v20161101:listBitLockerKey', __args__, opts=opts, typ=ListBitLockerKeyResult).value

    return AwaitableListBitLockerKeyResult(
        value=__ret__.value)
