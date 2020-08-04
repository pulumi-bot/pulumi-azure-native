# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class GetVirtualMachineScaleSetResult:
    """
    Describes a Virtual Machine Scale Set.
    """
    def __init__(__self__, identity=None, location=None, name=None, plan=None, properties=None, sku=None, tags=None, type=None, zones=None):
        if identity and not isinstance(identity, dict):
            raise TypeError("Expected argument 'identity' to be a dict")
        __self__.identity = identity
        """
        The identity of the virtual machine scale set, if configured.
        """
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        __self__.location = location
        """
        Resource location
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        Resource name
        """
        if plan and not isinstance(plan, dict):
            raise TypeError("Expected argument 'plan' to be a dict")
        __self__.plan = plan
        """
        Specifies information about the marketplace image used to create the virtual machine. This element is only used for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic use.  In the Azure portal, find the marketplace image that you want to use and then click **Want to deploy programmatically, Get Started ->**. Enter any required information and then click **Save**.
        """
        if properties and not isinstance(properties, dict):
            raise TypeError("Expected argument 'properties' to be a dict")
        __self__.properties = properties
        """
        Describes the properties of a Virtual Machine Scale Set.
        """
        if sku and not isinstance(sku, dict):
            raise TypeError("Expected argument 'sku' to be a dict")
        __self__.sku = sku
        """
        The virtual machine scale set sku.
        """
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
        """
        Resource tags
        """
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        __self__.type = type
        """
        Resource type
        """
        if zones and not isinstance(zones, list):
            raise TypeError("Expected argument 'zones' to be a list")
        __self__.zones = zones
        """
        The virtual machine scale set zones. NOTE: Availability zones can only be set when you create the scale set.
        """


class AwaitableGetVirtualMachineScaleSetResult(GetVirtualMachineScaleSetResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVirtualMachineScaleSetResult(
            identity=self.identity,
            location=self.location,
            name=self.name,
            plan=self.plan,
            properties=self.properties,
            sku=self.sku,
            tags=self.tags,
            type=self.type,
            zones=self.zones)


def get_virtual_machine_scale_set(name=None, resource_group_name=None, opts=None):
    """
    Use this data source to access information about an existing resource.

    :param str name: The name of the VM scale set.
    :param str resource_group_name: The name of the resource group.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:compute/v20170330:getVirtualMachineScaleSet', __args__, opts=opts).value

    return AwaitableGetVirtualMachineScaleSetResult(
        identity=__ret__.get('identity'),
        location=__ret__.get('location'),
        name=__ret__.get('name'),
        plan=__ret__.get('plan'),
        properties=__ret__.get('properties'),
        sku=__ret__.get('sku'),
        tags=__ret__.get('tags'),
        type=__ret__.get('type'),
        zones=__ret__.get('zones'))
