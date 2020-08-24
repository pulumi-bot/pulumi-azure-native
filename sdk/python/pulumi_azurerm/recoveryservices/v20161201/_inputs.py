# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'ProtectionContainerArgs',
]

@pulumi.input_type
class ProtectionContainerArgs:
    def __init__(__self__, *,
                 backup_management_type: Optional[pulumi.Input[str]] = None,
                 container_type: Optional[pulumi.Input[str]] = None,
                 friendly_name: Optional[pulumi.Input[str]] = None,
                 health_status: Optional[pulumi.Input[str]] = None,
                 registration_status: Optional[pulumi.Input[str]] = None):
        """
        Base class for container with backup items. Containers with specific workloads are derived from this class.
        :param pulumi.Input[str] backup_management_type: Type of backup management for the container.
        :param pulumi.Input[str] container_type: Type of the container. The value of this property for: 1. Compute Azure VM is Microsoft.Compute/virtualMachines 2.
               Classic Compute Azure VM is Microsoft.ClassicCompute/virtualMachines 3. Windows machines (like MAB, DPM etc) is
               Windows 4. Azure SQL instance is AzureSqlContainer. 5. Storage containers is StorageContainer. 6. Azure workload
               Backup is VMAppContainer
        :param pulumi.Input[str] friendly_name: Friendly name of the container.
        :param pulumi.Input[str] health_status: Status of health of the container.
        :param pulumi.Input[str] registration_status: Status of registration of the container with the Recovery Services Vault.
        """
        if backup_management_type is not None:
            pulumi.set(__self__, "backup_management_type", backup_management_type)
        if container_type is not None:
            pulumi.set(__self__, "container_type", container_type)
        if friendly_name is not None:
            pulumi.set(__self__, "friendly_name", friendly_name)
        if health_status is not None:
            pulumi.set(__self__, "health_status", health_status)
        if registration_status is not None:
            pulumi.set(__self__, "registration_status", registration_status)

    @property
    @pulumi.getter(name="backupManagementType")
    def backup_management_type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of backup management for the container.
        """
        return pulumi.get(self, "backup_management_type")

    @backup_management_type.setter
    def backup_management_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "backup_management_type", value)

    @property
    @pulumi.getter(name="containerType")
    def container_type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of the container. The value of this property for: 1. Compute Azure VM is Microsoft.Compute/virtualMachines 2.
        Classic Compute Azure VM is Microsoft.ClassicCompute/virtualMachines 3. Windows machines (like MAB, DPM etc) is
        Windows 4. Azure SQL instance is AzureSqlContainer. 5. Storage containers is StorageContainer. 6. Azure workload
        Backup is VMAppContainer
        """
        return pulumi.get(self, "container_type")

    @container_type.setter
    def container_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "container_type", value)

    @property
    @pulumi.getter(name="friendlyName")
    def friendly_name(self) -> Optional[pulumi.Input[str]]:
        """
        Friendly name of the container.
        """
        return pulumi.get(self, "friendly_name")

    @friendly_name.setter
    def friendly_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "friendly_name", value)

    @property
    @pulumi.getter(name="healthStatus")
    def health_status(self) -> Optional[pulumi.Input[str]]:
        """
        Status of health of the container.
        """
        return pulumi.get(self, "health_status")

    @health_status.setter
    def health_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "health_status", value)

    @property
    @pulumi.getter(name="registrationStatus")
    def registration_status(self) -> Optional[pulumi.Input[str]]:
        """
        Status of registration of the container with the Recovery Services Vault.
        """
        return pulumi.get(self, "registration_status")

    @registration_status.setter
    def registration_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "registration_status", value)


