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
    def __init__(__self__, boot_strap_configuration_blob=None, cloud_init_configuration_blob=None, etag=None, identity=None, location=None, name=None, provisioning_state=None, sku=None, tags=None, type=None, virtual_appliance_asn=None, virtual_appliance_nics=None, virtual_hub=None):
        if boot_strap_configuration_blob and not isinstance(boot_strap_configuration_blob, list):
            raise TypeError("Expected argument 'boot_strap_configuration_blob' to be a list")
        pulumi.set(__self__, "boot_strap_configuration_blob", boot_strap_configuration_blob)
        if cloud_init_configuration_blob and not isinstance(cloud_init_configuration_blob, list):
            raise TypeError("Expected argument 'cloud_init_configuration_blob' to be a list")
        pulumi.set(__self__, "cloud_init_configuration_blob", cloud_init_configuration_blob)
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
        if virtual_appliance_asn and not isinstance(virtual_appliance_asn, float):
            raise TypeError("Expected argument 'virtual_appliance_asn' to be a float")
        pulumi.set(__self__, "virtual_appliance_asn", virtual_appliance_asn)
        if virtual_appliance_nics and not isinstance(virtual_appliance_nics, list):
            raise TypeError("Expected argument 'virtual_appliance_nics' to be a list")
        pulumi.set(__self__, "virtual_appliance_nics", virtual_appliance_nics)
        if virtual_hub and not isinstance(virtual_hub, dict):
            raise TypeError("Expected argument 'virtual_hub' to be a dict")
        pulumi.set(__self__, "virtual_hub", virtual_hub)

    @property
    @pulumi.getter(name="bootStrapConfigurationBlob")
    def boot_strap_configuration_blob(self) -> Optional[List[str]]:
        """
        BootStrapConfigurationBlob storage URLs.
        """
        return pulumi.get(self, "boot_strap_configuration_blob")

    @property
    @pulumi.getter(name="cloudInitConfigurationBlob")
    def cloud_init_configuration_blob(self) -> Optional[List[str]]:
        """
        CloudInitConfigurationBlob storage URLs.
        """
        return pulumi.get(self, "cloud_init_configuration_blob")

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
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        The provisioning state of the resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def sku(self) -> Optional['outputs.VirtualApplianceSkuPropertiesResponse']:
        """
        Network Virtual Appliance SKU.
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
            boot_strap_configuration_blob=self.boot_strap_configuration_blob,
            cloud_init_configuration_blob=self.cloud_init_configuration_blob,
            etag=self.etag,
            identity=self.identity,
            location=self.location,
            name=self.name,
            provisioning_state=self.provisioning_state,
            sku=self.sku,
            tags=self.tags,
            type=self.type,
            virtual_appliance_asn=self.virtual_appliance_asn,
            virtual_appliance_nics=self.virtual_appliance_nics,
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
    __ret__ = pulumi.runtime.invoke('azurerm:network/v20200301:getNetworkVirtualAppliance', __args__, opts=opts, typ=GetNetworkVirtualApplianceResult).value

    return AwaitableGetNetworkVirtualApplianceResult(
        boot_strap_configuration_blob=__ret__.boot_strap_configuration_blob,
        cloud_init_configuration_blob=__ret__.cloud_init_configuration_blob,
        etag=__ret__.etag,
        identity=__ret__.identity,
        location=__ret__.location,
        name=__ret__.name,
        provisioning_state=__ret__.provisioning_state,
        sku=__ret__.sku,
        tags=__ret__.tags,
        type=__ret__.type,
        virtual_appliance_asn=__ret__.virtual_appliance_asn,
        virtual_appliance_nics=__ret__.virtual_appliance_nics,
        virtual_hub=__ret__.virtual_hub)
