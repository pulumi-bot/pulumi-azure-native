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
    'GetServerResult',
    'AwaitableGetServerResult',
    'get_server',
]

@pulumi.output_type
class GetServerResult:
    """
    Represents a server.
    """
    def __init__(__self__, administrator_login=None, earliest_restore_date=None, fully_qualified_domain_name=None, identity=None, location=None, master_server_id=None, name=None, replica_capacity=None, replication_role=None, sku=None, ssl_enforcement=None, storage_profile=None, tags=None, type=None, user_visible_state=None, version=None):
        if administrator_login and not isinstance(administrator_login, str):
            raise TypeError("Expected argument 'administrator_login' to be a str")
        pulumi.set(__self__, "administrator_login", administrator_login)
        if earliest_restore_date and not isinstance(earliest_restore_date, str):
            raise TypeError("Expected argument 'earliest_restore_date' to be a str")
        pulumi.set(__self__, "earliest_restore_date", earliest_restore_date)
        if fully_qualified_domain_name and not isinstance(fully_qualified_domain_name, str):
            raise TypeError("Expected argument 'fully_qualified_domain_name' to be a str")
        pulumi.set(__self__, "fully_qualified_domain_name", fully_qualified_domain_name)
        if identity and not isinstance(identity, dict):
            raise TypeError("Expected argument 'identity' to be a dict")
        pulumi.set(__self__, "identity", identity)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if master_server_id and not isinstance(master_server_id, str):
            raise TypeError("Expected argument 'master_server_id' to be a str")
        pulumi.set(__self__, "master_server_id", master_server_id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if replica_capacity and not isinstance(replica_capacity, float):
            raise TypeError("Expected argument 'replica_capacity' to be a float")
        pulumi.set(__self__, "replica_capacity", replica_capacity)
        if replication_role and not isinstance(replication_role, str):
            raise TypeError("Expected argument 'replication_role' to be a str")
        pulumi.set(__self__, "replication_role", replication_role)
        if sku and not isinstance(sku, dict):
            raise TypeError("Expected argument 'sku' to be a dict")
        pulumi.set(__self__, "sku", sku)
        if ssl_enforcement and not isinstance(ssl_enforcement, str):
            raise TypeError("Expected argument 'ssl_enforcement' to be a str")
        pulumi.set(__self__, "ssl_enforcement", ssl_enforcement)
        if storage_profile and not isinstance(storage_profile, dict):
            raise TypeError("Expected argument 'storage_profile' to be a dict")
        pulumi.set(__self__, "storage_profile", storage_profile)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if user_visible_state and not isinstance(user_visible_state, str):
            raise TypeError("Expected argument 'user_visible_state' to be a str")
        pulumi.set(__self__, "user_visible_state", user_visible_state)
        if version and not isinstance(version, str):
            raise TypeError("Expected argument 'version' to be a str")
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="administratorLogin")
    def administrator_login(self) -> Optional[str]:
        """
        The administrator's login name of a server. Can only be specified when the server is being created (and is required for creation).
        """
        return pulumi.get(self, "administrator_login")

    @property
    @pulumi.getter(name="earliestRestoreDate")
    def earliest_restore_date(self) -> Optional[str]:
        """
        Earliest restore point creation time (ISO8601 format)
        """
        return pulumi.get(self, "earliest_restore_date")

    @property
    @pulumi.getter(name="fullyQualifiedDomainName")
    def fully_qualified_domain_name(self) -> Optional[str]:
        """
        The fully qualified domain name of a server.
        """
        return pulumi.get(self, "fully_qualified_domain_name")

    @property
    @pulumi.getter
    def identity(self) -> Optional['outputs.ResourceIdentityResponse']:
        """
        The Azure Active Directory identity of the server.
        """
        return pulumi.get(self, "identity")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="masterServerId")
    def master_server_id(self) -> Optional[str]:
        """
        The master server id of a replica server.
        """
        return pulumi.get(self, "master_server_id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="replicaCapacity")
    def replica_capacity(self) -> Optional[float]:
        """
        The maximum number of replicas that a master server can have.
        """
        return pulumi.get(self, "replica_capacity")

    @property
    @pulumi.getter(name="replicationRole")
    def replication_role(self) -> Optional[str]:
        """
        The replication role of the server.
        """
        return pulumi.get(self, "replication_role")

    @property
    @pulumi.getter
    def sku(self) -> Optional['outputs.SkuResponse']:
        """
        The SKU (pricing tier) of the server.
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter(name="sslEnforcement")
    def ssl_enforcement(self) -> Optional[str]:
        """
        Enable ssl enforcement or not when connect to server.
        """
        return pulumi.get(self, "ssl_enforcement")

    @property
    @pulumi.getter(name="storageProfile")
    def storage_profile(self) -> Optional['outputs.StorageProfileResponse']:
        """
        Storage profile of a server.
        """
        return pulumi.get(self, "storage_profile")

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
        The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="userVisibleState")
    def user_visible_state(self) -> Optional[str]:
        """
        A state of a server that is visible to user.
        """
        return pulumi.get(self, "user_visible_state")

    @property
    @pulumi.getter
    def version(self) -> Optional[str]:
        """
        Server version.
        """
        return pulumi.get(self, "version")


class AwaitableGetServerResult(GetServerResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetServerResult(
            administrator_login=self.administrator_login,
            earliest_restore_date=self.earliest_restore_date,
            fully_qualified_domain_name=self.fully_qualified_domain_name,
            identity=self.identity,
            location=self.location,
            master_server_id=self.master_server_id,
            name=self.name,
            replica_capacity=self.replica_capacity,
            replication_role=self.replication_role,
            sku=self.sku,
            ssl_enforcement=self.ssl_enforcement,
            storage_profile=self.storage_profile,
            tags=self.tags,
            type=self.type,
            user_visible_state=self.user_visible_state,
            version=self.version)


def get_server(resource_group_name: Optional[str] = None,
               server_name: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetServerResult:
    """
    Use this data source to access information about an existing resource.

    :param str resource_group_name: The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
    :param str server_name: The name of the server.
    """
    __args__ = dict()
    __args__['resourceGroupName'] = resource_group_name
    __args__['serverName'] = server_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:dbformariadb/latest:getServer', __args__, opts=opts, typ=GetServerResult).value

    return AwaitableGetServerResult(
        administrator_login=__ret__.administrator_login,
        earliest_restore_date=__ret__.earliest_restore_date,
        fully_qualified_domain_name=__ret__.fully_qualified_domain_name,
        identity=__ret__.identity,
        location=__ret__.location,
        master_server_id=__ret__.master_server_id,
        name=__ret__.name,
        replica_capacity=__ret__.replica_capacity,
        replication_role=__ret__.replication_role,
        sku=__ret__.sku,
        ssl_enforcement=__ret__.ssl_enforcement,
        storage_profile=__ret__.storage_profile,
        tags=__ret__.tags,
        type=__ret__.type,
        user_visible_state=__ret__.user_visible_state,
        version=__ret__.version)
