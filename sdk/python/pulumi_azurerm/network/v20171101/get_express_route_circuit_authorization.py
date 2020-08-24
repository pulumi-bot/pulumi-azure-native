# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'GetExpressRouteCircuitAuthorizationResult',
    'AwaitableGetExpressRouteCircuitAuthorizationResult',
    'get_express_route_circuit_authorization',
]

@pulumi.output_type
class GetExpressRouteCircuitAuthorizationResult:
    """
    Authorization in an ExpressRouteCircuit resource.
    """
    def __init__(__self__, authorization_key=None, authorization_use_status=None, etag=None, name=None, provisioning_state=None):
        if authorization_key and not isinstance(authorization_key, str):
            raise TypeError("Expected argument 'authorization_key' to be a str")
        pulumi.set(__self__, "authorization_key", authorization_key)
        if authorization_use_status and not isinstance(authorization_use_status, str):
            raise TypeError("Expected argument 'authorization_use_status' to be a str")
        pulumi.set(__self__, "authorization_use_status", authorization_use_status)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)

    @property
    @pulumi.getter(name="authorizationKey")
    def authorization_key(self) -> Optional[str]:
        """
        The authorization key.
        """
        return pulumi.get(self, "authorization_key")

    @property
    @pulumi.getter(name="authorizationUseStatus")
    def authorization_use_status(self) -> Optional[str]:
        """
        AuthorizationUseStatus. Possible values are: 'Available' and 'InUse'.
        """
        return pulumi.get(self, "authorization_use_status")

    @property
    @pulumi.getter
    def etag(self) -> str:
        """
        A unique read-only string that changes whenever the resource is updated.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Gets name of the resource that is unique within a resource group. This name can be used to access the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> Optional[str]:
        """
        Gets the provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        """
        return pulumi.get(self, "provisioning_state")


class AwaitableGetExpressRouteCircuitAuthorizationResult(GetExpressRouteCircuitAuthorizationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetExpressRouteCircuitAuthorizationResult(
            authorization_key=self.authorization_key,
            authorization_use_status=self.authorization_use_status,
            etag=self.etag,
            name=self.name,
            provisioning_state=self.provisioning_state)


def get_express_route_circuit_authorization(circuit_name: Optional[str] = None,
                                            name: Optional[str] = None,
                                            resource_group_name: Optional[str] = None,
                                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetExpressRouteCircuitAuthorizationResult:
    """
    Use this data source to access information about an existing resource.

    :param str circuit_name: The name of the express route circuit.
    :param str name: The name of the authorization.
    :param str resource_group_name: The name of the resource group.
    """
    __args__ = dict()
    __args__['circuitName'] = circuit_name
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:network/v20171101:getExpressRouteCircuitAuthorization', __args__, opts=opts, typ=GetExpressRouteCircuitAuthorizationResult).value

    return AwaitableGetExpressRouteCircuitAuthorizationResult(
        authorization_key=__ret__.authorization_key,
        authorization_use_status=__ret__.authorization_use_status,
        etag=__ret__.etag,
        name=__ret__.name,
        provisioning_state=__ret__.provisioning_state)
