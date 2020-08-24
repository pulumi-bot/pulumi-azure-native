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
    'GetNetworkVirtualApplianceResult',
    'AwaitableGetNetworkVirtualApplianceResult',
    'get_network_virtual_appliance',
]

@pulumi.output_type
class GetNetworkVirtualApplianceResult:
    """
    NetworkVirtualAppliance Resource.
    """
    def __init__(__self__, boot_strap_configuration_blobs=None, cloud_init_configuration=None, cloud_init_configuration_blobs=None, etag=None, identity=None, location=None, name=None, nva_sku=None, provisioning_state=None, tags=None, type=None, virtual_appliance_asn=None, virtual_appliance_nics=None, virtual_appliance_sites=None, virtual_hub=None):
        if boot_strap_configuration_blobs and not isinstance(boot_strap_configuration_blobs, list):
            raise TypeError("Expected argument 'boot_strap_configuration_blobs' to be a list")
        pulumi.set(__self__, "boot_strap_configuration_blobs", boot_strap_configuration_blobs)
        if cloud_init_configuration and not isinstance(cloud_init_configuration, str):
            raise TypeError("Expected argument 'cloud_init_configuration' to be a str")
        pulumi.set(__self__, "cloud_init_configuration", cloud_init_configuration)
        if cloud_init_configuration_blobs and not isinstance(cloud_init_configuration_blobs, list):
            raise TypeError("Expected argument 'cloud_init_configuration_blobs' to be a list")
        pulumi.set(__self__, "cloud_init_configuration_blobs", cloud_init_configuration_blobs)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if identity and not isinstance(identity, dict):
            raise TypeError("Expected argument 'identity' to be a dict")
        pulumi.set(__self__, "identity", identity)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if nva_sku and not isinstance(nva_sku, dict):
            raise TypeError("Expected argument 'nva_sku' to be a dict")
        pulumi.set(__self__, "nva_sku", nva_sku)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if virtual_appliance_asn and not isinstance(virtual_appliance_asn, float):
            raise TypeError("Expected argument 'virtual_appliance_asn' to be a float")
        pulumi.set(__self__, "virtual_appliance_asn", virtual_appliance_asn)
        if virtual_appliance_nics and not isinstance(virtual_appliance_nics, list):
            raise TypeError("Expected argument 'virtual_appliance_nics' to be a list")
        pulumi.set(__self__, "virtual_appliance_nics", virtual_appliance_nics)
        if virtual_appliance_sites and not isinstance(virtual_appliance_sites, list):
            raise TypeError("Expected argument 'virtual_appliance_sites' to be a list")
        pulumi.set(__self__, "virtual_appliance_sites", virtual_appliance_sites)
        if virtual_hub and not isinstance(virtual_hub, dict):
            raise TypeError("Expected argument 'virtual_hub' to be a dict")
        pulumi.set(__self__, "virtual_hub", virtual_hub)

    @property
    @pulumi.getter(name="bootStrapConfigurationBlobs")
    def boot_strap_configuration_blobs(self) -> Optional[List[str]]:
        """
        BootStrapConfigurationBlobs storage URLs.
        """
        return pulumi.get(self, "boot_strap_configuration_blobs")

    @property
    @pulumi.getter(name="cloudInitConfiguration")
    def cloud_init_configuration(self) -> Optional[str]:
        """
        CloudInitConfiguration string in plain text.
        """
        return pulumi.get(self, "cloud_init_configuration")

    @property
    @pulumi.getter(name="cloudInitConfigurationBlobs")
    def cloud_init_configuration_blobs(self) -> Optional[List[str]]:
        """
        CloudInitConfigurationBlob storage URLs.
        """
        return pulumi.get(self, "cloud_init_configuration_blobs")

    @property
    @pulumi.getter
    def etag(self) -> str:
        """
        A unique read-only string that changes whenever the resource is updated.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def identity(self) -> Optional['outputs.ManagedServiceIdentityResponse']:
        """
        The service principal that has read access to cloud-init and config blob.
        """
        return pulumi.get(self, "identity")

    @property
    @pulumi.getter
    def location(self) -> Optional[str]:
        """
        Resource location.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nvaSku")
    def nva_sku(self) -> Optional['outputs.VirtualApplianceSkuPropertiesResponse']:
        """
        Network Virtual Appliance SKU.
        """
        return pulumi.get(self, "nva_sku")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        The provisioning state of the resource.
        """
        return pulumi.get(self, "provisioning_state")

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
        Resource type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="virtualApplianceAsn")
    def virtual_appliance_asn(self) -> Optional[float]:
        """
        VirtualAppliance ASN.
        """
        return pulumi.get(self, "virtual_appliance_asn")

    @property
    @pulumi.getter(name="virtualApplianceNics")
    def virtual_appliance_nics(self) -> List['outputs.VirtualApplianceNicPropertiesResponse']:
        """
        List of Virtual Appliance Network Interfaces.
        """
        return pulumi.get(self, "virtual_appliance_nics")

    @property
    @pulumi.getter(name="virtualApplianceSites")
    def virtual_appliance_sites(self) -> List['outputs.SubResourceResponse']:
        """
        List of references to VirtualApplianceSite.
        """
        return pulumi.get(self, "virtual_appliance_sites")

    @property
    @pulumi.getter(name="virtualHub")
    def virtual_hub(self) -> Optional['outputs.SubResourceResponse']:
        """
        The Virtual Hub where Network Virtual Appliance is being deployed.
        """
        return pulumi.get(self, "virtual_hub")


class AwaitableGetNetworkVirtualApplianceResult(GetNetworkVirtualApplianceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNetworkVirtualApplianceResult(
            boot_strap_configuration_blobs=self.boot_strap_configuration_blobs,
            cloud_init_configuration=self.cloud_init_configuration,
            cloud_init_configuration_blobs=self.cloud_init_configuration_blobs,
            etag=self.etag,
            identity=self.identity,
            location=self.location,
            name=self.name,
            nva_sku=self.nva_sku,
            provisioning_state=self.provisioning_state,
            tags=self.tags,
            type=self.type,
            virtual_appliance_asn=self.virtual_appliance_asn,
            virtual_appliance_nics=self.virtual_appliance_nics,
            virtual_appliance_sites=self.virtual_appliance_sites,
            virtual_hub=self.virtual_hub)


def get_network_virtual_appliance(expand: Optional[str] = None,
                                  name: Optional[str] = None,
                                  resource_group_name: Optional[str] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNetworkVirtualApplianceResult:
    """
    Use this data source to access information about an existing resource.

    :param str expand: Expands referenced resources.
    :param str name: The name of Network Virtual Appliance.
    :param str resource_group_name: The name of the resource group.
    """
    __args__ = dict()
    __args__['expand'] = expand
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:network/v20200501:getNetworkVirtualAppliance', __args__, opts=opts, typ=GetNetworkVirtualApplianceResult).value

    return AwaitableGetNetworkVirtualApplianceResult(
        boot_strap_configuration_blobs=__ret__.boot_strap_configuration_blobs,
        cloud_init_configuration=__ret__.cloud_init_configuration,
        cloud_init_configuration_blobs=__ret__.cloud_init_configuration_blobs,
        etag=__ret__.etag,
        identity=__ret__.identity,
        location=__ret__.location,
        name=__ret__.name,
        nva_sku=__ret__.nva_sku,
        provisioning_state=__ret__.provisioning_state,
        tags=__ret__.tags,
        type=__ret__.type,
        virtual_appliance_asn=__ret__.virtual_appliance_asn,
        virtual_appliance_nics=__ret__.virtual_appliance_nics,
        virtual_appliance_sites=__ret__.virtual_appliance_sites,
        virtual_hub=__ret__.virtual_hub)
