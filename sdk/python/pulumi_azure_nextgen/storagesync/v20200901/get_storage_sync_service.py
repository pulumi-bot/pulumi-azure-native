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
    'GetStorageSyncServiceResult',
    'AwaitableGetStorageSyncServiceResult',
    'get_storage_sync_service',
]

@pulumi.output_type
class GetStorageSyncServiceResult:
    """
    Storage Sync Service object.
    """
    def __init__(__self__, incoming_traffic_policy=None, last_operation_name=None, last_workflow_id=None, location=None, name=None, private_endpoint_connections=None, provisioning_state=None, storage_sync_service_status=None, storage_sync_service_uid=None, tags=None, type=None):
        if incoming_traffic_policy and not isinstance(incoming_traffic_policy, str):
            raise TypeError("Expected argument 'incoming_traffic_policy' to be a str")
        pulumi.set(__self__, "incoming_traffic_policy", incoming_traffic_policy)
        if last_operation_name and not isinstance(last_operation_name, str):
            raise TypeError("Expected argument 'last_operation_name' to be a str")
        pulumi.set(__self__, "last_operation_name", last_operation_name)
        if last_workflow_id and not isinstance(last_workflow_id, str):
            raise TypeError("Expected argument 'last_workflow_id' to be a str")
        pulumi.set(__self__, "last_workflow_id", last_workflow_id)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if private_endpoint_connections and not isinstance(private_endpoint_connections, list):
            raise TypeError("Expected argument 'private_endpoint_connections' to be a list")
        pulumi.set(__self__, "private_endpoint_connections", private_endpoint_connections)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if storage_sync_service_status and not isinstance(storage_sync_service_status, int):
            raise TypeError("Expected argument 'storage_sync_service_status' to be a int")
        pulumi.set(__self__, "storage_sync_service_status", storage_sync_service_status)
        if storage_sync_service_uid and not isinstance(storage_sync_service_uid, str):
            raise TypeError("Expected argument 'storage_sync_service_uid' to be a str")
        pulumi.set(__self__, "storage_sync_service_uid", storage_sync_service_uid)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="incomingTrafficPolicy")
    def incoming_traffic_policy(self) -> Optional[str]:
        """
        Incoming Traffic Policy
        """
        return pulumi.get(self, "incoming_traffic_policy")

    @property
    @pulumi.getter(name="lastOperationName")
    def last_operation_name(self) -> str:
        """
        Resource Last Operation Name
        """
        return pulumi.get(self, "last_operation_name")

    @property
    @pulumi.getter(name="lastWorkflowId")
    def last_workflow_id(self) -> str:
        """
        StorageSyncService lastWorkflowId
        """
        return pulumi.get(self, "last_workflow_id")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="privateEndpointConnections")
    def private_endpoint_connections(self) -> Sequence['outputs.PrivateEndpointConnectionResponse']:
        """
        List of private endpoint connection associated with the specified storage sync service
        """
        return pulumi.get(self, "private_endpoint_connections")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        StorageSyncService Provisioning State
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="storageSyncServiceStatus")
    def storage_sync_service_status(self) -> int:
        """
        Storage Sync service status.
        """
        return pulumi.get(self, "storage_sync_service_status")

    @property
    @pulumi.getter(name="storageSyncServiceUid")
    def storage_sync_service_uid(self) -> str:
        """
        Storage Sync service Uid
        """
        return pulumi.get(self, "storage_sync_service_uid")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")


class AwaitableGetStorageSyncServiceResult(GetStorageSyncServiceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetStorageSyncServiceResult(
            incoming_traffic_policy=self.incoming_traffic_policy,
            last_operation_name=self.last_operation_name,
            last_workflow_id=self.last_workflow_id,
            location=self.location,
            name=self.name,
            private_endpoint_connections=self.private_endpoint_connections,
            provisioning_state=self.provisioning_state,
            storage_sync_service_status=self.storage_sync_service_status,
            storage_sync_service_uid=self.storage_sync_service_uid,
            tags=self.tags,
            type=self.type)


def get_storage_sync_service(resource_group_name: Optional[str] = None,
                             storage_sync_service_name: Optional[str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetStorageSyncServiceResult:
    """
    Use this data source to access information about an existing resource.

    :param str resource_group_name: The name of the resource group. The name is case insensitive.
    :param str storage_sync_service_name: Name of Storage Sync Service resource.
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['storageSyncServiceName'] = storage_sync_service_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure-nextgen:storagesync/v20200901:getStorageSyncService', __args__, opts=opts, typ=GetStorageSyncServiceResult).value

    return AwaitableGetStorageSyncServiceResult(
        incoming_traffic_policy=__ret__.incoming_traffic_policy,
        last_operation_name=__ret__.last_operation_name,
        last_workflow_id=__ret__.last_workflow_id,
        location=__ret__.location,
        name=__ret__.name,
        private_endpoint_connections=__ret__.private_endpoint_connections,
        provisioning_state=__ret__.provisioning_state,
        storage_sync_service_status=__ret__.storage_sync_service_status,
        storage_sync_service_uid=__ret__.storage_sync_service_uid,
        tags=__ret__.tags,
        type=__ret__.type)
