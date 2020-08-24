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
    'ListGlobalUserLabsResult',
    'AwaitableListGlobalUserLabsResult',
    'list_global_user_labs',
]

@pulumi.output_type
class ListGlobalUserLabsResult:
    """
    Lists the labs owned by a user
    """
    def __init__(__self__, labs=None):
        if labs and not isinstance(labs, list):
            raise TypeError("Expected argument 'labs' to be a list")
        pulumi.set(__self__, "labs", labs)

    @property
    @pulumi.getter
    def labs(self) -> Optional[List['outputs.LabDetailsResponseResult']]:
        """
        List of all the labs
        """
        return pulumi.get(self, "labs")


class AwaitableListGlobalUserLabsResult(ListGlobalUserLabsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListGlobalUserLabsResult(
            labs=self.labs)


def list_global_user_labs(user_name: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableListGlobalUserLabsResult:
    """
    Use this data source to access information about an existing resource.

    :param str user_name: The name of the user.
    """
    __args__ = dict()
    __args__['userName'] = user_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:labservices/v20181015:listGlobalUserLabs', __args__, opts=opts, typ=ListGlobalUserLabsResult).value

    return AwaitableListGlobalUserLabsResult(
        labs=__ret__.labs)
