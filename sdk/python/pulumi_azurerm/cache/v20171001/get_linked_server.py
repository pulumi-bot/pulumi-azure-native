# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'GetLinkedServerResult',
    'AwaitableGetLinkedServerResult',
    'get_linked_server',
]

@pulumi.output_type
class GetLinkedServerResult:
    """
    Response to put/get linked server (with properties) for Redis cache.
    """
    def __init__(__self__, linked_redis_cache_id=None, linked_redis_cache_location=None, name=None, provisioning_state=None, server_role=None, type=None):
        if linked_redis_cache_id and not isinstance(linked_redis_cache_id, str):
            raise TypeError("Expected argument 'linked_redis_cache_id' to be a str")
        pulumi.set(__self__, "linked_redis_cache_id", linked_redis_cache_id)
        if linked_redis_cache_location and not isinstance(linked_redis_cache_location, str):
            raise TypeError("Expected argument 'linked_redis_cache_location' to be a str")
        pulumi.set(__self__, "linked_redis_cache_location", linked_redis_cache_location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if server_role and not isinstance(server_role, str):
            raise TypeError("Expected argument 'server_role' to be a str")
        pulumi.set(__self__, "server_role", server_role)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="linkedRedisCacheId")
    def linked_redis_cache_id(self) -> str:
        """
        Fully qualified resourceId of the linked redis cache.
        """
        return pulumi.get(self, "linked_redis_cache_id")

    @property
    @pulumi.getter(name="linkedRedisCacheLocation")
    def linked_redis_cache_location(self) -> str:
        """
        Location of the linked redis cache.
        """
        return pulumi.get(self, "linked_redis_cache_location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        Terminal state of the link between primary and secondary redis cache.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="serverRole")
    def server_role(self) -> str:
        """
        Role of the linked server.
        """
        return pulumi.get(self, "server_role")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Resource type.
        """
        return pulumi.get(self, "type")


class AwaitableGetLinkedServerResult(GetLinkedServerResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLinkedServerResult(
            linked_redis_cache_id=self.linked_redis_cache_id,
            linked_redis_cache_location=self.linked_redis_cache_location,
            name=self.name,
            provisioning_state=self.provisioning_state,
            server_role=self.server_role,
            type=self.type)


def get_linked_server(name: Optional[str] = None,
                      resource_group_name: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetLinkedServerResult:
    """
    Use this data source to access information about an existing resource.

    :param str name: The name of the linked server.
    :param str resource_group_name: The name of the resource group.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:cache/v20171001:getLinkedServer', __args__, opts=opts, typ=GetLinkedServerResult).value

    return AwaitableGetLinkedServerResult(
        linked_redis_cache_id=__ret__.linked_redis_cache_id,
        linked_redis_cache_location=__ret__.linked_redis_cache_location,
        name=__ret__.name,
        provisioning_state=__ret__.provisioning_state,
        server_role=__ret__.server_role,
        type=__ret__.type)
