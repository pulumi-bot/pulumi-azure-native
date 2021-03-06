# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from . import outputs
from ._enums import *
from ._inputs import *

__all__ = ['NetworkVirtualAppliance']

warnings.warn("""The 'latest' version is deprecated. Please migrate to the resource in the top-level module: 'azure-native:network:NetworkVirtualAppliance'.""", DeprecationWarning)


class NetworkVirtualAppliance(pulumi.CustomResource):
    warnings.warn("""The 'latest' version is deprecated. Please migrate to the resource in the top-level module: 'azure-native:network:NetworkVirtualAppliance'.""", DeprecationWarning)

    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 boot_strap_configuration_blobs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 cloud_init_configuration: Optional[pulumi.Input[str]] = None,
                 cloud_init_configuration_blobs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 identity: Optional[pulumi.Input[pulumi.InputType['ManagedServiceIdentityArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 network_virtual_appliance_name: Optional[pulumi.Input[str]] = None,
                 nva_sku: Optional[pulumi.Input[pulumi.InputType['VirtualApplianceSkuPropertiesArgs']]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 virtual_appliance_asn: Optional[pulumi.Input[float]] = None,
                 virtual_hub: Optional[pulumi.Input[pulumi.InputType['SubResourceArgs']]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        NetworkVirtualAppliance Resource.
        Latest API Version: 2020-08-01.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] boot_strap_configuration_blobs: BootStrapConfigurationBlobs storage URLs.
        :param pulumi.Input[str] cloud_init_configuration: CloudInitConfiguration string in plain text.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] cloud_init_configuration_blobs: CloudInitConfigurationBlob storage URLs.
        :param pulumi.Input[str] id: Resource ID.
        :param pulumi.Input[pulumi.InputType['ManagedServiceIdentityArgs']] identity: The service principal that has read access to cloud-init and config blob.
        :param pulumi.Input[str] location: Resource location.
        :param pulumi.Input[str] network_virtual_appliance_name: The name of Network Virtual Appliance.
        :param pulumi.Input[pulumi.InputType['VirtualApplianceSkuPropertiesArgs']] nva_sku: Network Virtual Appliance SKU.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags.
        :param pulumi.Input[float] virtual_appliance_asn: VirtualAppliance ASN.
        :param pulumi.Input[pulumi.InputType['SubResourceArgs']] virtual_hub: The Virtual Hub where Network Virtual Appliance is being deployed.
        """
        pulumi.log.warn("""NetworkVirtualAppliance is deprecated: The 'latest' version is deprecated. Please migrate to the resource in the top-level module: 'azure-native:network:NetworkVirtualAppliance'.""")
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['boot_strap_configuration_blobs'] = boot_strap_configuration_blobs
            __props__['cloud_init_configuration'] = cloud_init_configuration
            __props__['cloud_init_configuration_blobs'] = cloud_init_configuration_blobs
            __props__['id'] = id
            __props__['identity'] = identity
            __props__['location'] = location
            __props__['network_virtual_appliance_name'] = network_virtual_appliance_name
            __props__['nva_sku'] = nva_sku
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['virtual_appliance_asn'] = virtual_appliance_asn
            __props__['virtual_hub'] = virtual_hub
            __props__['address_prefix'] = None
            __props__['etag'] = None
            __props__['inbound_security_rules'] = None
            __props__['name'] = None
            __props__['provisioning_state'] = None
            __props__['type'] = None
            __props__['virtual_appliance_nics'] = None
            __props__['virtual_appliance_sites'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:network/latest:NetworkVirtualAppliance"), pulumi.Alias(type_="azure-native:network:NetworkVirtualAppliance"), pulumi.Alias(type_="azure-nextgen:network:NetworkVirtualAppliance"), pulumi.Alias(type_="azure-native:network/v20191201:NetworkVirtualAppliance"), pulumi.Alias(type_="azure-nextgen:network/v20191201:NetworkVirtualAppliance"), pulumi.Alias(type_="azure-native:network/v20200301:NetworkVirtualAppliance"), pulumi.Alias(type_="azure-nextgen:network/v20200301:NetworkVirtualAppliance"), pulumi.Alias(type_="azure-native:network/v20200401:NetworkVirtualAppliance"), pulumi.Alias(type_="azure-nextgen:network/v20200401:NetworkVirtualAppliance"), pulumi.Alias(type_="azure-native:network/v20200501:NetworkVirtualAppliance"), pulumi.Alias(type_="azure-nextgen:network/v20200501:NetworkVirtualAppliance"), pulumi.Alias(type_="azure-native:network/v20200601:NetworkVirtualAppliance"), pulumi.Alias(type_="azure-nextgen:network/v20200601:NetworkVirtualAppliance"), pulumi.Alias(type_="azure-native:network/v20200701:NetworkVirtualAppliance"), pulumi.Alias(type_="azure-nextgen:network/v20200701:NetworkVirtualAppliance"), pulumi.Alias(type_="azure-native:network/v20200801:NetworkVirtualAppliance"), pulumi.Alias(type_="azure-nextgen:network/v20200801:NetworkVirtualAppliance")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(NetworkVirtualAppliance, __self__).__init__(
            'azure-native:network/latest:NetworkVirtualAppliance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'NetworkVirtualAppliance':
        """
        Get an existing NetworkVirtualAppliance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["address_prefix"] = None
        __props__["boot_strap_configuration_blobs"] = None
        __props__["cloud_init_configuration"] = None
        __props__["cloud_init_configuration_blobs"] = None
        __props__["etag"] = None
        __props__["identity"] = None
        __props__["inbound_security_rules"] = None
        __props__["location"] = None
        __props__["name"] = None
        __props__["nva_sku"] = None
        __props__["provisioning_state"] = None
        __props__["tags"] = None
        __props__["type"] = None
        __props__["virtual_appliance_asn"] = None
        __props__["virtual_appliance_nics"] = None
        __props__["virtual_appliance_sites"] = None
        __props__["virtual_hub"] = None
        return NetworkVirtualAppliance(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="addressPrefix")
    def address_prefix(self) -> pulumi.Output[str]:
        """
        Address Prefix.
        """
        return pulumi.get(self, "address_prefix")

    @property
    @pulumi.getter(name="bootStrapConfigurationBlobs")
    def boot_strap_configuration_blobs(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        BootStrapConfigurationBlobs storage URLs.
        """
        return pulumi.get(self, "boot_strap_configuration_blobs")

    @property
    @pulumi.getter(name="cloudInitConfiguration")
    def cloud_init_configuration(self) -> pulumi.Output[Optional[str]]:
        """
        CloudInitConfiguration string in plain text.
        """
        return pulumi.get(self, "cloud_init_configuration")

    @property
    @pulumi.getter(name="cloudInitConfigurationBlobs")
    def cloud_init_configuration_blobs(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        CloudInitConfigurationBlob storage URLs.
        """
        return pulumi.get(self, "cloud_init_configuration_blobs")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        A unique read-only string that changes whenever the resource is updated.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def identity(self) -> pulumi.Output[Optional['outputs.ManagedServiceIdentityResponse']]:
        """
        The service principal that has read access to cloud-init and config blob.
        """
        return pulumi.get(self, "identity")

    @property
    @pulumi.getter(name="inboundSecurityRules")
    def inbound_security_rules(self) -> pulumi.Output[Sequence['outputs.SubResourceResponse']]:
        """
        List of references to InboundSecurityRules.
        """
        return pulumi.get(self, "inbound_security_rules")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[Optional[str]]:
        """
        Resource location.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nvaSku")
    def nva_sku(self) -> pulumi.Output[Optional['outputs.VirtualApplianceSkuPropertiesResponse']]:
        """
        Network Virtual Appliance SKU.
        """
        return pulumi.get(self, "nva_sku")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        The provisioning state of the resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Resource type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="virtualApplianceAsn")
    def virtual_appliance_asn(self) -> pulumi.Output[Optional[float]]:
        """
        VirtualAppliance ASN.
        """
        return pulumi.get(self, "virtual_appliance_asn")

    @property
    @pulumi.getter(name="virtualApplianceNics")
    def virtual_appliance_nics(self) -> pulumi.Output[Sequence['outputs.VirtualApplianceNicPropertiesResponse']]:
        """
        List of Virtual Appliance Network Interfaces.
        """
        return pulumi.get(self, "virtual_appliance_nics")

    @property
    @pulumi.getter(name="virtualApplianceSites")
    def virtual_appliance_sites(self) -> pulumi.Output[Sequence['outputs.SubResourceResponse']]:
        """
        List of references to VirtualApplianceSite.
        """
        return pulumi.get(self, "virtual_appliance_sites")

    @property
    @pulumi.getter(name="virtualHub")
    def virtual_hub(self) -> pulumi.Output[Optional['outputs.SubResourceResponse']]:
        """
        The Virtual Hub where Network Virtual Appliance is being deployed.
        """
        return pulumi.get(self, "virtual_hub")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

