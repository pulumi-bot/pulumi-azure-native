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
    'GetReplicationStorageClassificationMappingResult',
    'AwaitableGetReplicationStorageClassificationMappingResult',
    'get_replication_storage_classification_mapping',
]

@pulumi.output_type
class GetReplicationStorageClassificationMappingResult:
    """
    Storage mapping object.
    """
    def __init__(__self__, location=None, name=None, properties=None, type=None):
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if properties and not isinstance(properties, dict):
            raise TypeError("Expected argument 'properties' to be a dict")
        pulumi.set(__self__, "properties", properties)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def location(self) -> Optional[str]:
        """
        Resource Location
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource Name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def properties(self) -> 'outputs.StorageClassificationMappingPropertiesResponse':
        """
        Properties of the storage mapping object.
        """
        return pulumi.get(self, "properties")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Resource Type
        """
        return pulumi.get(self, "type")


class AwaitableGetReplicationStorageClassificationMappingResult(GetReplicationStorageClassificationMappingResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetReplicationStorageClassificationMappingResult(
            location=self.location,
            name=self.name,
            properties=self.properties,
            type=self.type)


def get_replication_storage_classification_mapping(fabric_name: Optional[str] = None,
                                                   name: Optional[str] = None,
                                                   resource_group_name: Optional[str] = None,
                                                   resource_name: Optional[str] = None,
                                                   storage_classification_name: Optional[str] = None,
                                                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetReplicationStorageClassificationMappingResult:
    """
    Use this data source to access information about an existing resource.

    :param str fabric_name: Fabric name.
    :param str name: Storage classification mapping name.
    :param str resource_group_name: The name of the resource group where the recovery services vault is present.
    :param str resource_name: The name of the recovery services vault.
    :param str storage_classification_name: Storage classification name.
    """
    __args__ = dict()
    __args__['fabricName'] = fabric_name
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    __args__['resourceName'] = resource_name
    __args__['storageClassificationName'] = storage_classification_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:recoveryservices/v20160810:getReplicationStorageClassificationMapping', __args__, opts=opts, typ=GetReplicationStorageClassificationMappingResult).value

    return AwaitableGetReplicationStorageClassificationMappingResult(
        location=__ret__.location,
        name=__ret__.name,
        properties=__ret__.properties,
        type=__ret__.type)
