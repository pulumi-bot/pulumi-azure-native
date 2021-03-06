# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from . import outputs

__all__ = [
    'GetConsoleResult',
    'AwaitableGetConsoleResult',
    'get_console',
]

warnings.warn("""The 'latest' version is deprecated. Please migrate to the function in the top-level module: 'azure-native:portal:getConsole'.""", DeprecationWarning)

@pulumi.output_type
class GetConsoleResult:
    """
    Cloud shell console
    """
    def __init__(__self__, properties=None):
        if properties and not isinstance(properties, dict):
            raise TypeError("Expected argument 'properties' to be a dict")
        pulumi.set(__self__, "properties", properties)

    @property
    @pulumi.getter
    def properties(self) -> 'outputs.ConsolePropertiesResponse':
        """
        Cloud shell console properties.
        """
        return pulumi.get(self, "properties")


class AwaitableGetConsoleResult(GetConsoleResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetConsoleResult(
            properties=self.properties)


def get_console(console_name: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetConsoleResult:
    """
    Cloud shell console
    Latest API Version: 2018-10-01.


    :param str console_name: The name of the console
    """
    pulumi.log.warn("""get_console is deprecated: The 'latest' version is deprecated. Please migrate to the function in the top-level module: 'azure-native:portal:getConsole'.""")
    __args__ = dict()
    __args__['consoleName'] = console_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure-native:portal/latest:getConsole', __args__, opts=opts, typ=GetConsoleResult).value

    return AwaitableGetConsoleResult(
        properties=__ret__.properties)
