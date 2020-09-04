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
    'GetManagedNetworkGroupResult',
    'AwaitableGetManagedNetworkGroupResult',
    'get_managed_network_group',
]

@pulumi.output_type
class GetManagedNetworkGroupResult:
    """
    The Managed Network Group resource
    """
    def __init__(__self__, etag=None, kind=None, location=None, management_groups=None, name=None, provisioning_state=None, subnets=None, subscriptions=None, type=None, virtual_networks=None):
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if management_groups and not isinstance(management_groups, list):
            raise TypeError("Expected argument 'management_groups' to be a list")
        pulumi.set(__self__, "management_groups", management_groups)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if subnets and not isinstance(subnets, list):
            raise TypeError("Expected argument 'subnets' to be a list")
        pulumi.set(__self__, "subnets", subnets)
        if subscriptions and not isinstance(subscriptions, list):
            raise TypeError("Expected argument 'subscriptions' to be a list")
        pulumi.set(__self__, "subscriptions", subscriptions)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if virtual_networks and not isinstance(virtual_networks, list):
            raise TypeError("Expected argument 'virtual_networks' to be a list")
        pulumi.set(__self__, "virtual_networks", virtual_networks)

    @property
    @pulumi.getter
    def etag(self) -> str:
        """
        A unique read-only string that changes whenever the resource is updated.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def kind(self) -> Optional[str]:
        """
        Responsibility role under which this Managed Network Group will be created
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def location(self) -> Optional[str]:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="managementGroups")
    def management_groups(self) -> Optional[List['outputs.ResourceIdResponse']]:
        """
        The collection of management groups covered by the Managed Network
        """
        return pulumi.get(self, "management_groups")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        Provisioning state of the ManagedNetwork resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def subnets(self) -> Optional[List['outputs.ResourceIdResponse']]:
        """
        The collection of  subnets covered by the Managed Network
        """
        return pulumi.get(self, "subnets")

    @property
    @pulumi.getter
    def subscriptions(self) -> Optional[List['outputs.ResourceIdResponse']]:
        """
        The collection of subscriptions covered by the Managed Network
        """
        return pulumi.get(self, "subscriptions")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="virtualNetworks")
    def virtual_networks(self) -> Optional[List['outputs.ResourceIdResponse']]:
        """
        The collection of virtual nets covered by the Managed Network
        """
        return pulumi.get(self, "virtual_networks")


class AwaitableGetManagedNetworkGroupResult(GetManagedNetworkGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetManagedNetworkGroupResult(
            etag=self.etag,
            kind=self.kind,
            location=self.location,
            management_groups=self.management_groups,
            name=self.name,
            provisioning_state=self.provisioning_state,
            subnets=self.subnets,
            subscriptions=self.subscriptions,
            type=self.type,
            virtual_networks=self.virtual_networks)


def get_managed_network_group(managed_network_group_name: Optional[str] = None,
                              managed_network_name: Optional[str] = None,
                              resource_group_name: Optional[str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetManagedNetworkGroupResult:
    """
    Use this data source to access information about an existing resource.

    :param str managed_network_group_name: The name of the Managed Network Group.
    :param str managed_network_name: The name of the Managed Network.
    :param str resource_group_name: The name of the resource group.
    """
    __args__ = dict()
    __args__['managedNetworkGroupName'] = managed_network_group_name
    __args__['managedNetworkName'] = managed_network_name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:managednetwork/v20190601preview:getManagedNetworkGroup', __args__, opts=opts, typ=GetManagedNetworkGroupResult).value

    return AwaitableGetManagedNetworkGroupResult(
        etag=__ret__.etag,
        kind=__ret__.kind,
        location=__ret__.location,
        management_groups=__ret__.management_groups,
        name=__ret__.name,
        provisioning_state=__ret__.provisioning_state,
        subnets=__ret__.subnets,
        subscriptions=__ret__.subscriptions,
        type=__ret__.type,
        virtual_networks=__ret__.virtual_networks)
