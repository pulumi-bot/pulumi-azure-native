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
    'GetUserSettingsWithLocationResult',
    'AwaitableGetUserSettingsWithLocationResult',
    'get_user_settings_with_location',
]

@pulumi.output_type
class GetUserSettingsWithLocationResult:
    """
    Response to get user settings
    """
    def __init__(__self__, properties=None):
        if properties and not isinstance(properties, dict):
            raise TypeError("Expected argument 'properties' to be a dict")
        pulumi.set(__self__, "properties", properties)

    @property
    @pulumi.getter
    def properties(self) -> 'outputs.UserPropertiesResponse':
        """
        The cloud shell user settings properties.
        """
        return pulumi.get(self, "properties")


class AwaitableGetUserSettingsWithLocationResult(GetUserSettingsWithLocationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetUserSettingsWithLocationResult(
            properties=self.properties)


def get_user_settings_with_location(location: Optional[str] = None,
                                    user_settings_name: Optional[str] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetUserSettingsWithLocationResult:
    """
    Use this data source to access information about an existing resource.

    :param str location: The provider location
    :param str user_settings_name: The name of the user settings
    """
    __args__ = dict()
    __args__['location'] = location
    __args__['userSettingsName'] = user_settings_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:portal/v20181001:getUserSettingsWithLocation', __args__, opts=opts, typ=GetUserSettingsWithLocationResult).value

    return AwaitableGetUserSettingsWithLocationResult(
        properties=__ret__.properties)
