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
    'GetStorageAccountResult',
    'AwaitableGetStorageAccountResult',
    'get_storage_account',
]

@pulumi.output_type
class GetStorageAccountResult:
    """
    The storage account.
    """
    def __init__(__self__, access_tier=None, creation_time=None, custom_domain=None, enable_https_traffic_only=None, encryption=None, identity=None, kind=None, last_geo_failover_time=None, location=None, name=None, network_rule_set=None, primary_endpoints=None, primary_location=None, provisioning_state=None, secondary_endpoints=None, secondary_location=None, sku=None, status_of_primary=None, status_of_secondary=None, tags=None, type=None):
        if access_tier and not isinstance(access_tier, str):
            raise TypeError("Expected argument 'access_tier' to be a str")
        pulumi.set(__self__, "access_tier", access_tier)
        if creation_time and not isinstance(creation_time, str):
            raise TypeError("Expected argument 'creation_time' to be a str")
        pulumi.set(__self__, "creation_time", creation_time)
        if custom_domain and not isinstance(custom_domain, dict):
            raise TypeError("Expected argument 'custom_domain' to be a dict")
        pulumi.set(__self__, "custom_domain", custom_domain)
        if enable_https_traffic_only and not isinstance(enable_https_traffic_only, bool):
            raise TypeError("Expected argument 'enable_https_traffic_only' to be a bool")
        pulumi.set(__self__, "enable_https_traffic_only", enable_https_traffic_only)
        if encryption and not isinstance(encryption, dict):
            raise TypeError("Expected argument 'encryption' to be a dict")
        pulumi.set(__self__, "encryption", encryption)
        if identity and not isinstance(identity, dict):
            raise TypeError("Expected argument 'identity' to be a dict")
        pulumi.set(__self__, "identity", identity)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if last_geo_failover_time and not isinstance(last_geo_failover_time, str):
            raise TypeError("Expected argument 'last_geo_failover_time' to be a str")
        pulumi.set(__self__, "last_geo_failover_time", last_geo_failover_time)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if network_rule_set and not isinstance(network_rule_set, dict):
            raise TypeError("Expected argument 'network_rule_set' to be a dict")
        pulumi.set(__self__, "network_rule_set", network_rule_set)
        if primary_endpoints and not isinstance(primary_endpoints, dict):
            raise TypeError("Expected argument 'primary_endpoints' to be a dict")
        pulumi.set(__self__, "primary_endpoints", primary_endpoints)
        if primary_location and not isinstance(primary_location, str):
            raise TypeError("Expected argument 'primary_location' to be a str")
        pulumi.set(__self__, "primary_location", primary_location)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if secondary_endpoints and not isinstance(secondary_endpoints, dict):
            raise TypeError("Expected argument 'secondary_endpoints' to be a dict")
        pulumi.set(__self__, "secondary_endpoints", secondary_endpoints)
        if secondary_location and not isinstance(secondary_location, str):
            raise TypeError("Expected argument 'secondary_location' to be a str")
        pulumi.set(__self__, "secondary_location", secondary_location)
        if sku and not isinstance(sku, dict):
            raise TypeError("Expected argument 'sku' to be a dict")
        pulumi.set(__self__, "sku", sku)
        if status_of_primary and not isinstance(status_of_primary, str):
            raise TypeError("Expected argument 'status_of_primary' to be a str")
        pulumi.set(__self__, "status_of_primary", status_of_primary)
        if status_of_secondary and not isinstance(status_of_secondary, str):
            raise TypeError("Expected argument 'status_of_secondary' to be a str")
        pulumi.set(__self__, "status_of_secondary", status_of_secondary)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="accessTier")
    def access_tier(self) -> str:
        """
        Required for storage accounts where kind = BlobStorage. The access tier used for billing.
        """
        return pulumi.get(self, "access_tier")

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> str:
        """
        Gets the creation date and time of the storage account in UTC.
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter(name="customDomain")
    def custom_domain(self) -> 'outputs.CustomDomainResponse':
        """
        Gets the custom domain the user assigned to this storage account.
        """
        return pulumi.get(self, "custom_domain")

    @property
    @pulumi.getter(name="enableHttpsTrafficOnly")
    def enable_https_traffic_only(self) -> Optional[bool]:
        """
        Allows https traffic only to storage service if sets to true.
        """
        return pulumi.get(self, "enable_https_traffic_only")

    @property
    @pulumi.getter
    def encryption(self) -> 'outputs.EncryptionResponse':
        """
        Gets the encryption settings on the account. If unspecified, the account is unencrypted.
        """
        return pulumi.get(self, "encryption")

    @property
    @pulumi.getter
    def identity(self) -> Optional['outputs.IdentityResponse']:
        """
        The identity of the resource.
        """
        return pulumi.get(self, "identity")

    @property
    @pulumi.getter
    def kind(self) -> str:
        """
        Gets the Kind.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="lastGeoFailoverTime")
    def last_geo_failover_time(self) -> str:
        """
        Gets the timestamp of the most recent instance of a failover to the secondary location. Only the most recent timestamp is retained. This element is not returned if there has never been a failover instance. Only available if the accountType is Standard_GRS or Standard_RAGRS.
        """
        return pulumi.get(self, "last_geo_failover_time")

    @property
    @pulumi.getter
    def location(self) -> Optional[str]:
        """
        Resource location
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkRuleSet")
    def network_rule_set(self) -> 'outputs.NetworkRuleSetResponse':
        """
        Network rule set
        """
        return pulumi.get(self, "network_rule_set")

    @property
    @pulumi.getter(name="primaryEndpoints")
    def primary_endpoints(self) -> 'outputs.EndpointsResponse':
        """
        Gets the URLs that are used to perform a retrieval of a public blob, queue, or table object. Note that Standard_ZRS and Premium_LRS accounts only return the blob endpoint.
        """
        return pulumi.get(self, "primary_endpoints")

    @property
    @pulumi.getter(name="primaryLocation")
    def primary_location(self) -> str:
        """
        Gets the location of the primary data center for the storage account.
        """
        return pulumi.get(self, "primary_location")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        Gets the status of the storage account at the time the operation was called.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="secondaryEndpoints")
    def secondary_endpoints(self) -> 'outputs.EndpointsResponse':
        """
        Gets the URLs that are used to perform a retrieval of a public blob, queue, or table object from the secondary location of the storage account. Only available if the SKU name is Standard_RAGRS.
        """
        return pulumi.get(self, "secondary_endpoints")

    @property
    @pulumi.getter(name="secondaryLocation")
    def secondary_location(self) -> str:
        """
        Gets the location of the geo-replicated secondary for the storage account. Only available if the accountType is Standard_GRS or Standard_RAGRS.
        """
        return pulumi.get(self, "secondary_location")

    @property
    @pulumi.getter
    def sku(self) -> 'outputs.SkuResponse':
        """
        Gets the SKU.
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter(name="statusOfPrimary")
    def status_of_primary(self) -> str:
        """
        Gets the status indicating whether the primary location of the storage account is available or unavailable.
        """
        return pulumi.get(self, "status_of_primary")

    @property
    @pulumi.getter(name="statusOfSecondary")
    def status_of_secondary(self) -> str:
        """
        Gets the status indicating whether the secondary location of the storage account is available or unavailable. Only available if the SKU name is Standard_GRS or Standard_RAGRS.
        """
        return pulumi.get(self, "status_of_secondary")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        Tags assigned to a resource; can be used for viewing and grouping a resource (across resource groups).
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Resource type
        """
        return pulumi.get(self, "type")


class AwaitableGetStorageAccountResult(GetStorageAccountResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetStorageAccountResult(
            access_tier=self.access_tier,
            creation_time=self.creation_time,
            custom_domain=self.custom_domain,
            enable_https_traffic_only=self.enable_https_traffic_only,
            encryption=self.encryption,
            identity=self.identity,
            kind=self.kind,
            last_geo_failover_time=self.last_geo_failover_time,
            location=self.location,
            name=self.name,
            network_rule_set=self.network_rule_set,
            primary_endpoints=self.primary_endpoints,
            primary_location=self.primary_location,
            provisioning_state=self.provisioning_state,
            secondary_endpoints=self.secondary_endpoints,
            secondary_location=self.secondary_location,
            sku=self.sku,
            status_of_primary=self.status_of_primary,
            status_of_secondary=self.status_of_secondary,
            tags=self.tags,
            type=self.type)


def get_storage_account(name: Optional[str] = None,
                        resource_group_name: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetStorageAccountResult:
    """
    Use this data source to access information about an existing resource.

    :param str name: The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.  
    :param str resource_group_name: The name of the resource group within the user's subscription. The name is case insensitive.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:storage/v20171001:getStorageAccount', __args__, opts=opts, typ=GetStorageAccountResult).value

    return AwaitableGetStorageAccountResult(
        access_tier=__ret__.access_tier,
        creation_time=__ret__.creation_time,
        custom_domain=__ret__.custom_domain,
        enable_https_traffic_only=__ret__.enable_https_traffic_only,
        encryption=__ret__.encryption,
        identity=__ret__.identity,
        kind=__ret__.kind,
        last_geo_failover_time=__ret__.last_geo_failover_time,
        location=__ret__.location,
        name=__ret__.name,
        network_rule_set=__ret__.network_rule_set,
        primary_endpoints=__ret__.primary_endpoints,
        primary_location=__ret__.primary_location,
        provisioning_state=__ret__.provisioning_state,
        secondary_endpoints=__ret__.secondary_endpoints,
        secondary_location=__ret__.secondary_location,
        sku=__ret__.sku,
        status_of_primary=__ret__.status_of_primary,
        status_of_secondary=__ret__.status_of_secondary,
        tags=__ret__.tags,
        type=__ret__.type)
