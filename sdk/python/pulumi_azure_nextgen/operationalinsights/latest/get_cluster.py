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
    'GetClusterResult',
    'AwaitableGetClusterResult',
    'get_cluster',
]

@pulumi.output_type
class GetClusterResult:
    """
    The top level Log Analytics cluster resource container.
    """
    def __init__(__self__, billing_type=None, cluster_id=None, identity=None, is_availability_zones_enabled=None, is_double_encryption_enabled=None, key_vault_properties=None, location=None, name=None, provisioning_state=None, sku=None, tags=None, type=None):
        if billing_type and not isinstance(billing_type, str):
            raise TypeError("Expected argument 'billing_type' to be a str")
        pulumi.set(__self__, "billing_type", billing_type)
        if cluster_id and not isinstance(cluster_id, str):
            raise TypeError("Expected argument 'cluster_id' to be a str")
        pulumi.set(__self__, "cluster_id", cluster_id)
        if identity and not isinstance(identity, dict):
            raise TypeError("Expected argument 'identity' to be a dict")
        pulumi.set(__self__, "identity", identity)
        if is_availability_zones_enabled and not isinstance(is_availability_zones_enabled, bool):
            raise TypeError("Expected argument 'is_availability_zones_enabled' to be a bool")
        pulumi.set(__self__, "is_availability_zones_enabled", is_availability_zones_enabled)
        if is_double_encryption_enabled and not isinstance(is_double_encryption_enabled, bool):
            raise TypeError("Expected argument 'is_double_encryption_enabled' to be a bool")
        pulumi.set(__self__, "is_double_encryption_enabled", is_double_encryption_enabled)
        if key_vault_properties and not isinstance(key_vault_properties, dict):
            raise TypeError("Expected argument 'key_vault_properties' to be a dict")
        pulumi.set(__self__, "key_vault_properties", key_vault_properties)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if sku and not isinstance(sku, dict):
            raise TypeError("Expected argument 'sku' to be a dict")
        pulumi.set(__self__, "sku", sku)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="billingType")
    def billing_type(self) -> Optional[str]:
        """
        Configures whether billing will be only on the cluster or each workspace will be billed by its proportional use. This does not change the overall billing, only how it will be distributed. Default value is 'Cluster'
        """
        return pulumi.get(self, "billing_type")

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> str:
        """
        The ID associated with the cluster.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter
    def identity(self) -> Optional['outputs.IdentityResponse']:
        """
        The identity of the resource.
        """
        return pulumi.get(self, "identity")

    @property
    @pulumi.getter(name="isAvailabilityZonesEnabled")
    def is_availability_zones_enabled(self) -> Optional[bool]:
        """
        Sets whether the cluster will support availability zones. This can be set as true only in regions where Azure Data Explorer support Availability Zones. This Property can not be modified after cluster creation. Default value is 'true' if region supports Availability Zones.
        """
        return pulumi.get(self, "is_availability_zones_enabled")

    @property
    @pulumi.getter(name="isDoubleEncryptionEnabled")
    def is_double_encryption_enabled(self) -> Optional[bool]:
        """
        Configures whether cluster will use double encryption. This Property can not be modified after cluster creation. Default value is 'true'
        """
        return pulumi.get(self, "is_double_encryption_enabled")

    @property
    @pulumi.getter(name="keyVaultProperties")
    def key_vault_properties(self) -> Optional['outputs.KeyVaultPropertiesResponse']:
        """
        The associated key properties.
        """
        return pulumi.get(self, "key_vault_properties")

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
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        The provisioning state of the cluster.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def sku(self) -> Optional['outputs.ClusterSkuResponse']:
        """
        The sku properties.
        """
        return pulumi.get(self, "sku")

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


class AwaitableGetClusterResult(GetClusterResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetClusterResult(
            billing_type=self.billing_type,
            cluster_id=self.cluster_id,
            identity=self.identity,
            is_availability_zones_enabled=self.is_availability_zones_enabled,
            is_double_encryption_enabled=self.is_double_encryption_enabled,
            key_vault_properties=self.key_vault_properties,
            location=self.location,
            name=self.name,
            provisioning_state=self.provisioning_state,
            sku=self.sku,
            tags=self.tags,
            type=self.type)


def get_cluster(cluster_name: Optional[str] = None,
                resource_group_name: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetClusterResult:
    """
    Use this data source to access information about an existing resource.

    :param str cluster_name: Name of the Log Analytics Cluster.
    :param str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['clusterName'] = cluster_name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure-nextgen:operationalinsights/latest:getCluster', __args__, opts=opts, typ=GetClusterResult).value

    return AwaitableGetClusterResult(
        billing_type=__ret__.billing_type,
        cluster_id=__ret__.cluster_id,
        identity=__ret__.identity,
        is_availability_zones_enabled=__ret__.is_availability_zones_enabled,
        is_double_encryption_enabled=__ret__.is_double_encryption_enabled,
        key_vault_properties=__ret__.key_vault_properties,
        location=__ret__.location,
        name=__ret__.name,
        provisioning_state=__ret__.provisioning_state,
        sku=__ret__.sku,
        tags=__ret__.tags,
        type=__ret__.type)
