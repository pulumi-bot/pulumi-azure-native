# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'ProtectedItemArgs',
]

@pulumi.input_type
class ProtectedItemArgs:
    def __init__(__self__, *,
                 backup_management_type: Optional[pulumi.Input[str]] = None,
                 backup_set_name: Optional[pulumi.Input[str]] = None,
                 container_name: Optional[pulumi.Input[str]] = None,
                 create_mode: Optional[pulumi.Input[str]] = None,
                 deferred_delete_time_in_utc: Optional[pulumi.Input[str]] = None,
                 deferred_delete_time_remaining: Optional[pulumi.Input[str]] = None,
                 is_deferred_delete_schedule_upcoming: Optional[pulumi.Input[bool]] = None,
                 is_rehydrate: Optional[pulumi.Input[bool]] = None,
                 is_scheduled_for_deferred_delete: Optional[pulumi.Input[bool]] = None,
                 last_recovery_point: Optional[pulumi.Input[str]] = None,
                 policy_id: Optional[pulumi.Input[str]] = None,
                 protected_item_type: Optional[pulumi.Input[str]] = None,
                 source_resource_id: Optional[pulumi.Input[str]] = None,
                 workload_type: Optional[pulumi.Input[str]] = None):
        """
        Base class for backup items.
        :param pulumi.Input[str] backup_management_type: Type of backup management for the backed up item.
        :param pulumi.Input[str] backup_set_name: Name of the backup set the backup item belongs to
        :param pulumi.Input[str] container_name: Unique name of container
        :param pulumi.Input[str] create_mode: Create mode to indicate recovery of existing soft deleted data source or creation of new data source.
        :param pulumi.Input[str] deferred_delete_time_in_utc: Time for deferred deletion in UTC
        :param pulumi.Input[str] deferred_delete_time_remaining: Time remaining before the DS marked for deferred delete is permanently deleted
        :param pulumi.Input[bool] is_deferred_delete_schedule_upcoming: Flag to identify whether the deferred deleted DS is to be purged soon
        :param pulumi.Input[bool] is_rehydrate: Flag to identify that deferred deleted DS is to be moved into Pause state
        :param pulumi.Input[bool] is_scheduled_for_deferred_delete: Flag to identify whether the DS is scheduled for deferred delete
        :param pulumi.Input[str] last_recovery_point: Timestamp when the last (latest) backup copy was created for this backup item.
        :param pulumi.Input[str] policy_id: ID of the backup policy with which this item is backed up.
        :param pulumi.Input[str] protected_item_type: backup item type.
        :param pulumi.Input[str] source_resource_id: ARM ID of the resource to be backed up.
        :param pulumi.Input[str] workload_type: Type of workload this item represents.
        """
        if backup_management_type is not None:
            pulumi.set(__self__, "backup_management_type", backup_management_type)
        if backup_set_name is not None:
            pulumi.set(__self__, "backup_set_name", backup_set_name)
        if container_name is not None:
            pulumi.set(__self__, "container_name", container_name)
        if create_mode is not None:
            pulumi.set(__self__, "create_mode", create_mode)
        if deferred_delete_time_in_utc is not None:
            pulumi.set(__self__, "deferred_delete_time_in_utc", deferred_delete_time_in_utc)
        if deferred_delete_time_remaining is not None:
            pulumi.set(__self__, "deferred_delete_time_remaining", deferred_delete_time_remaining)
        if is_deferred_delete_schedule_upcoming is not None:
            pulumi.set(__self__, "is_deferred_delete_schedule_upcoming", is_deferred_delete_schedule_upcoming)
        if is_rehydrate is not None:
            pulumi.set(__self__, "is_rehydrate", is_rehydrate)
        if is_scheduled_for_deferred_delete is not None:
            pulumi.set(__self__, "is_scheduled_for_deferred_delete", is_scheduled_for_deferred_delete)
        if last_recovery_point is not None:
            pulumi.set(__self__, "last_recovery_point", last_recovery_point)
        if policy_id is not None:
            pulumi.set(__self__, "policy_id", policy_id)
        if protected_item_type is not None:
            pulumi.set(__self__, "protected_item_type", protected_item_type)
        if source_resource_id is not None:
            pulumi.set(__self__, "source_resource_id", source_resource_id)
        if workload_type is not None:
            pulumi.set(__self__, "workload_type", workload_type)

    @property
    @pulumi.getter(name="backupManagementType")
    def backup_management_type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of backup management for the backed up item.
        """
        return pulumi.get(self, "backup_management_type")

    @backup_management_type.setter
    def backup_management_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "backup_management_type", value)

    @property
    @pulumi.getter(name="backupSetName")
    def backup_set_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the backup set the backup item belongs to
        """
        return pulumi.get(self, "backup_set_name")

    @backup_set_name.setter
    def backup_set_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "backup_set_name", value)

    @property
    @pulumi.getter(name="containerName")
    def container_name(self) -> Optional[pulumi.Input[str]]:
        """
        Unique name of container
        """
        return pulumi.get(self, "container_name")

    @container_name.setter
    def container_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "container_name", value)

    @property
    @pulumi.getter(name="createMode")
    def create_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Create mode to indicate recovery of existing soft deleted data source or creation of new data source.
        """
        return pulumi.get(self, "create_mode")

    @create_mode.setter
    def create_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_mode", value)

    @property
    @pulumi.getter(name="deferredDeleteTimeInUTC")
    def deferred_delete_time_in_utc(self) -> Optional[pulumi.Input[str]]:
        """
        Time for deferred deletion in UTC
        """
        return pulumi.get(self, "deferred_delete_time_in_utc")

    @deferred_delete_time_in_utc.setter
    def deferred_delete_time_in_utc(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "deferred_delete_time_in_utc", value)

    @property
    @pulumi.getter(name="deferredDeleteTimeRemaining")
    def deferred_delete_time_remaining(self) -> Optional[pulumi.Input[str]]:
        """
        Time remaining before the DS marked for deferred delete is permanently deleted
        """
        return pulumi.get(self, "deferred_delete_time_remaining")

    @deferred_delete_time_remaining.setter
    def deferred_delete_time_remaining(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "deferred_delete_time_remaining", value)

    @property
    @pulumi.getter(name="isDeferredDeleteScheduleUpcoming")
    def is_deferred_delete_schedule_upcoming(self) -> Optional[pulumi.Input[bool]]:
        """
        Flag to identify whether the deferred deleted DS is to be purged soon
        """
        return pulumi.get(self, "is_deferred_delete_schedule_upcoming")

    @is_deferred_delete_schedule_upcoming.setter
    def is_deferred_delete_schedule_upcoming(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_deferred_delete_schedule_upcoming", value)

    @property
    @pulumi.getter(name="isRehydrate")
    def is_rehydrate(self) -> Optional[pulumi.Input[bool]]:
        """
        Flag to identify that deferred deleted DS is to be moved into Pause state
        """
        return pulumi.get(self, "is_rehydrate")

    @is_rehydrate.setter
    def is_rehydrate(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_rehydrate", value)

    @property
    @pulumi.getter(name="isScheduledForDeferredDelete")
    def is_scheduled_for_deferred_delete(self) -> Optional[pulumi.Input[bool]]:
        """
        Flag to identify whether the DS is scheduled for deferred delete
        """
        return pulumi.get(self, "is_scheduled_for_deferred_delete")

    @is_scheduled_for_deferred_delete.setter
    def is_scheduled_for_deferred_delete(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_scheduled_for_deferred_delete", value)

    @property
    @pulumi.getter(name="lastRecoveryPoint")
    def last_recovery_point(self) -> Optional[pulumi.Input[str]]:
        """
        Timestamp when the last (latest) backup copy was created for this backup item.
        """
        return pulumi.get(self, "last_recovery_point")

    @last_recovery_point.setter
    def last_recovery_point(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "last_recovery_point", value)

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the backup policy with which this item is backed up.
        """
        return pulumi.get(self, "policy_id")

    @policy_id.setter
    def policy_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_id", value)

    @property
    @pulumi.getter(name="protectedItemType")
    def protected_item_type(self) -> Optional[pulumi.Input[str]]:
        """
        backup item type.
        """
        return pulumi.get(self, "protected_item_type")

    @protected_item_type.setter
    def protected_item_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protected_item_type", value)

    @property
    @pulumi.getter(name="sourceResourceId")
    def source_resource_id(self) -> Optional[pulumi.Input[str]]:
        """
        ARM ID of the resource to be backed up.
        """
        return pulumi.get(self, "source_resource_id")

    @source_resource_id.setter
    def source_resource_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_resource_id", value)

    @property
    @pulumi.getter(name="workloadType")
    def workload_type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of workload this item represents.
        """
        return pulumi.get(self, "workload_type")

    @workload_type.setter
    def workload_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "workload_type", value)


