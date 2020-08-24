# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'GetMigrationConfigResult',
    'AwaitableGetMigrationConfigResult',
    'get_migration_config',
]

@pulumi.output_type
class GetMigrationConfigResult:
    """
    Single item in List or Get Migration Config operation
    """
    def __init__(__self__, migration_state=None, name=None, pending_replication_operations_count=None, post_migration_name=None, provisioning_state=None, target_namespace=None, type=None):
        if migration_state and not isinstance(migration_state, str):
            raise TypeError("Expected argument 'migration_state' to be a str")
        pulumi.set(__self__, "migration_state", migration_state)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if pending_replication_operations_count and not isinstance(pending_replication_operations_count, float):
            raise TypeError("Expected argument 'pending_replication_operations_count' to be a float")
        pulumi.set(__self__, "pending_replication_operations_count", pending_replication_operations_count)
        if post_migration_name and not isinstance(post_migration_name, str):
            raise TypeError("Expected argument 'post_migration_name' to be a str")
        pulumi.set(__self__, "post_migration_name", post_migration_name)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if target_namespace and not isinstance(target_namespace, str):
            raise TypeError("Expected argument 'target_namespace' to be a str")
        pulumi.set(__self__, "target_namespace", target_namespace)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="migrationState")
    def migration_state(self) -> str:
        """
        State in which Standard to Premium Migration is, possible values : Unknown, Reverting, Completing, Initiating, Syncing, Active
        """
        return pulumi.get(self, "migration_state")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="pendingReplicationOperationsCount")
    def pending_replication_operations_count(self) -> float:
        """
        Number of entities pending to be replicated.
        """
        return pulumi.get(self, "pending_replication_operations_count")

    @property
    @pulumi.getter(name="postMigrationName")
    def post_migration_name(self) -> str:
        """
        Name to access Standard Namespace after migration
        """
        return pulumi.get(self, "post_migration_name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        Provisioning state of Migration Configuration 
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="targetNamespace")
    def target_namespace(self) -> str:
        """
        Existing premium Namespace ARM Id name which has no entities, will be used for migration
        """
        return pulumi.get(self, "target_namespace")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Resource type
        """
        return pulumi.get(self, "type")


class AwaitableGetMigrationConfigResult(GetMigrationConfigResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMigrationConfigResult(
            migration_state=self.migration_state,
            name=self.name,
            pending_replication_operations_count=self.pending_replication_operations_count,
            post_migration_name=self.post_migration_name,
            provisioning_state=self.provisioning_state,
            target_namespace=self.target_namespace,
            type=self.type)


def get_migration_config(name: Optional[str] = None,
                         namespace_name: Optional[str] = None,
                         resource_group_name: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMigrationConfigResult:
    """
    Use this data source to access information about an existing resource.

    :param str name: The configuration name. Should always be "$default".
    :param str namespace_name: The namespace name
    :param str resource_group_name: Name of the Resource group within the Azure subscription.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['namespaceName'] = namespace_name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:servicebus/v20170401:getMigrationConfig', __args__, opts=opts, typ=GetMigrationConfigResult).value

    return AwaitableGetMigrationConfigResult(
        migration_state=__ret__.migration_state,
        name=__ret__.name,
        pending_replication_operations_count=__ret__.pending_replication_operations_count,
        post_migration_name=__ret__.post_migration_name,
        provisioning_state=__ret__.provisioning_state,
        target_namespace=__ret__.target_namespace,
        type=__ret__.type)
