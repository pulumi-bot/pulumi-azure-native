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
    'GetManagerResult',
    'AwaitableGetManagerResult',
    'get_manager',
]

@pulumi.output_type
class GetManagerResult:
    """
    The StorSimple Manager.
    """
    def __init__(__self__, cis_intrinsic_settings=None, etag=None, location=None, name=None, provisioning_state=None, sku=None, tags=None, type=None):
        if cis_intrinsic_settings and not isinstance(cis_intrinsic_settings, dict):
            raise TypeError("Expected argument 'cis_intrinsic_settings' to be a dict")
        pulumi.set(__self__, "cis_intrinsic_settings", cis_intrinsic_settings)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if sku and not isinstance(sku, dict):
            raise TypeError("Expected argument 'sku' to be a dict")
        pulumi.set(__self__, "sku", sku)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="cisIntrinsicSettings")
    def cis_intrinsic_settings(self) -> Optional['outputs.ManagerIntrinsicSettingsResponse']:
        """
        Represents the type of StorSimple Manager.
        """
        return pulumi.get(self, "cis_intrinsic_settings")

    @property
    @pulumi.getter
    def etag(self) -> Optional[str]:
        """
        The etag of the manager.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        The geo location of the resource.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> Optional[str]:
        """
        Specifies the state of the resource as it is getting provisioned. Value of "Succeeded" means the Manager was successfully created.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def sku(self) -> Optional['outputs.ManagerSkuResponse']:
        """
        Specifies the Sku.
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        The tags attached to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The resource type.
        """
        return pulumi.get(self, "type")


class AwaitableGetManagerResult(GetManagerResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetManagerResult(
            cis_intrinsic_settings=self.cis_intrinsic_settings,
            etag=self.etag,
            location=self.location,
            name=self.name,
            provisioning_state=self.provisioning_state,
            sku=self.sku,
            tags=self.tags,
            type=self.type)


def get_manager(name: Optional[str] = None,
                resource_group_name: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetManagerResult:
    """
    Use this data source to access information about an existing resource.

    :param str name: The manager name
    :param str resource_group_name: The resource group name
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:storsimple/v20170601:getManager', __args__, opts=opts, typ=GetManagerResult).value

    return AwaitableGetManagerResult(
        cis_intrinsic_settings=__ret__.cis_intrinsic_settings,
        etag=__ret__.etag,
        location=__ret__.location,
        name=__ret__.name,
        provisioning_state=__ret__.provisioning_state,
        sku=__ret__.sku,
        tags=__ret__.tags,
        type=__ret__.type)
