# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables

__all__ = [
    'GetDatabaseSecurityAlertPolicyResult',
    'AwaitableGetDatabaseSecurityAlertPolicyResult',
    'get_database_security_alert_policy',
]

@pulumi.output_type
class GetDatabaseSecurityAlertPolicyResult:
    """
    A database security alert policy.
    """
    def __init__(__self__, creation_time=None, id=None, name=None, state=None, type=None):
        if creation_time and not isinstance(creation_time, str):
            raise TypeError("Expected argument 'creation_time' to be a str")
        pulumi.set(__self__, "creation_time", creation_time)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> str:
        """
        Specifies the UTC creation time of the policy.
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        Resource ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        Specifies the state of the policy, whether it is enabled or disabled or a policy has not been applied yet on the specific database.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Resource type.
        """
        return pulumi.get(self, "type")


class AwaitableGetDatabaseSecurityAlertPolicyResult(GetDatabaseSecurityAlertPolicyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDatabaseSecurityAlertPolicyResult(
            creation_time=self.creation_time,
            id=self.id,
            name=self.name,
            state=self.state,
            type=self.type)


def get_database_security_alert_policy(database_name: Optional[str] = None,
                                       resource_group_name: Optional[str] = None,
                                       security_alert_policy_name: Optional[str] = None,
                                       server_name: Optional[str] = None,
                                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDatabaseSecurityAlertPolicyResult:
    """
    Use this data source to access information about an existing resource.

    :param str database_name: The name of the  database for which the security alert policy is defined.
    :param str resource_group_name: The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
    :param str security_alert_policy_name: The name of the security alert policy.
    :param str server_name: The name of the  server.
    """
    __args__ = dict()
    __args__['databaseName'] = database_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['securityAlertPolicyName'] = security_alert_policy_name
    __args__['serverName'] = server_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure-nextgen:sql/v20200801preview:getDatabaseSecurityAlertPolicy', __args__, opts=opts, typ=GetDatabaseSecurityAlertPolicyResult).value

    return AwaitableGetDatabaseSecurityAlertPolicyResult(
        creation_time=__ret__.creation_time,
        id=__ret__.id,
        name=__ret__.name,
        state=__ret__.state,
        type=__ret__.type)
