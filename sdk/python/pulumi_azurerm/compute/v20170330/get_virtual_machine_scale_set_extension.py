# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class GetVirtualMachineScaleSetExtensionResult:
    """
    Describes a Virtual Machine Scale Set Extension.
    """
    def __init__(__self__, name=None, properties=None):
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        The name of the extension.
        """
        if properties and not isinstance(properties, dict):
            raise TypeError("Expected argument 'properties' to be a dict")
        __self__.properties = properties
        """
        Describes the properties of a Virtual Machine Scale Set Extension.
        """


class AwaitableGetVirtualMachineScaleSetExtensionResult(GetVirtualMachineScaleSetExtensionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVirtualMachineScaleSetExtensionResult(
            name=self.name,
            properties=self.properties)


def get_virtual_machine_scale_set_extension(name=None, resource_group_name=None, vm_scale_set_name=None, opts=None):
    """
    Use this data source to access information about an existing resource.

    :param str name: The name of the VM scale set extension.
    :param str resource_group_name: The name of the resource group.
    :param str vm_scale_set_name: The name of the VM scale set containing the extension.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    __args__['vmScaleSetName'] = vm_scale_set_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:compute/v20170330:getVirtualMachineScaleSetExtension', __args__, opts=opts).value

    return AwaitableGetVirtualMachineScaleSetExtensionResult(
        name=__ret__.get('name'),
        properties=__ret__.get('properties'))