# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['PrivateCloud']


class PrivateCloud(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 identity_sources: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['IdentitySourceArgs']]]]] = None,
                 internet: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 management_cluster: Optional[pulumi.Input[pulumi.InputType['ManagementClusterArgs']]] = None,
                 network_block: Optional[pulumi.Input[str]] = None,
                 nsxt_password: Optional[pulumi.Input[str]] = None,
                 private_cloud_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 sku: Optional[pulumi.Input[pulumi.InputType['SkuArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vcenter_password: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        A private cloud resource

        ## Example Usage
        ### PrivateClouds_CreateOrUpdate

        ```python
        import pulumi
        import pulumi_azurerm as azurerm

        private_cloud = azurerm.avs.v20200320.PrivateCloud("privateCloud",
            location="eastus2",
            management_cluster={
                "clusterSize": 4,
            },
            network_block="192.168.48.0/22",
            private_cloud_name="cloud1",
            resource_group_name="group1",
            sku={
                "name": "AV36",
            },
            tags={})

        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['IdentitySourceArgs']]]] identity_sources: vCenter Single Sign On Identity Sources
        :param pulumi.Input[str] internet: Connectivity to internet is enabled or disabled
        :param pulumi.Input[str] location: Resource location
        :param pulumi.Input[pulumi.InputType['ManagementClusterArgs']] management_cluster: The default cluster used for management
        :param pulumi.Input[str] network_block: The block of addresses should be unique across VNet in your subscription as well as on-premise. Make sure the CIDR format is conformed to (A.B.C.D/X) where A,B,C,D are between 0 and 255, and X is between 0 and 22
        :param pulumi.Input[str] nsxt_password: Optionally, set the NSX-T Manager password when the private cloud is created
        :param pulumi.Input[str] private_cloud_name: Name of the private cloud
        :param pulumi.Input[str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[pulumi.InputType['SkuArgs']] sku: The private cloud SKU
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags
        :param pulumi.Input[str] vcenter_password: Optionally, set the vCenter admin password when the private cloud is created
        """
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

            __props__['identity_sources'] = identity_sources
            __props__['internet'] = internet
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            if management_cluster is None:
                raise TypeError("Missing required property 'management_cluster'")
            __props__['management_cluster'] = management_cluster
            if network_block is None:
                raise TypeError("Missing required property 'network_block'")
            __props__['network_block'] = network_block
            __props__['nsxt_password'] = nsxt_password
            if private_cloud_name is None:
                raise TypeError("Missing required property 'private_cloud_name'")
            __props__['private_cloud_name'] = private_cloud_name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if sku is None:
                raise TypeError("Missing required property 'sku'")
            __props__['sku'] = sku
            __props__['tags'] = tags
            __props__['vcenter_password'] = vcenter_password
            __props__['circuit'] = None
            __props__['endpoints'] = None
            __props__['management_network'] = None
            __props__['name'] = None
            __props__['nsxt_certificate_thumbprint'] = None
            __props__['provisioning_network'] = None
            __props__['provisioning_state'] = None
            __props__['type'] = None
            __props__['vcenter_certificate_thumbprint'] = None
            __props__['vmotion_network'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:avs/latest:PrivateCloud"), pulumi.Alias(type_="azurerm:avs/v20190809preview:PrivateCloud")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(PrivateCloud, __self__).__init__(
            'azurerm:avs/v20200320:PrivateCloud',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'PrivateCloud':
        """
        Get an existing PrivateCloud resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return PrivateCloud(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def circuit(self) -> pulumi.Output[Optional['outputs.CircuitResponse']]:
        """
        An ExpressRoute Circuit
        """
        return pulumi.get(self, "circuit")

    @property
    @pulumi.getter
    def endpoints(self) -> pulumi.Output['outputs.EndpointsResponse']:
        """
        The endpoints
        """
        return pulumi.get(self, "endpoints")

    @property
    @pulumi.getter(name="identitySources")
    def identity_sources(self) -> pulumi.Output[Optional[List['outputs.IdentitySourceResponse']]]:
        """
        vCenter Single Sign On Identity Sources
        """
        return pulumi.get(self, "identity_sources")

    @property
    @pulumi.getter
    def internet(self) -> pulumi.Output[Optional[str]]:
        """
        Connectivity to internet is enabled or disabled
        """
        return pulumi.get(self, "internet")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Resource location
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="managementCluster")
    def management_cluster(self) -> pulumi.Output['outputs.ManagementClusterResponse']:
        """
        The default cluster used for management
        """
        return pulumi.get(self, "management_cluster")

    @property
    @pulumi.getter(name="managementNetwork")
    def management_network(self) -> pulumi.Output[str]:
        """
        Network used to access vCenter Server and NSX-T Manager
        """
        return pulumi.get(self, "management_network")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkBlock")
    def network_block(self) -> pulumi.Output[str]:
        """
        The block of addresses should be unique across VNet in your subscription as well as on-premise. Make sure the CIDR format is conformed to (A.B.C.D/X) where A,B,C,D are between 0 and 255, and X is between 0 and 22
        """
        return pulumi.get(self, "network_block")

    @property
    @pulumi.getter(name="nsxtCertificateThumbprint")
    def nsxt_certificate_thumbprint(self) -> pulumi.Output[str]:
        """
        Thumbprint of the NSX-T Manager SSL certificate
        """
        return pulumi.get(self, "nsxt_certificate_thumbprint")

    @property
    @pulumi.getter(name="nsxtPassword")
    def nsxt_password(self) -> pulumi.Output[Optional[str]]:
        """
        Optionally, set the NSX-T Manager password when the private cloud is created
        """
        return pulumi.get(self, "nsxt_password")

    @property
    @pulumi.getter(name="provisioningNetwork")
    def provisioning_network(self) -> pulumi.Output[str]:
        """
        Used for virtual machine cold migration, cloning, and snapshot migration
        """
        return pulumi.get(self, "provisioning_network")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        The provisioning state
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def sku(self) -> pulumi.Output['outputs.SkuResponse']:
        """
        The private cloud SKU
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Resource tags
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
    @pulumi.getter(name="vcenterCertificateThumbprint")
    def vcenter_certificate_thumbprint(self) -> pulumi.Output[str]:
        """
        Thumbprint of the vCenter Server SSL certificate
        """
        return pulumi.get(self, "vcenter_certificate_thumbprint")

    @property
    @pulumi.getter(name="vcenterPassword")
    def vcenter_password(self) -> pulumi.Output[Optional[str]]:
        """
        Optionally, set the vCenter admin password when the private cloud is created
        """
        return pulumi.get(self, "vcenter_password")

    @property
    @pulumi.getter(name="vmotionNetwork")
    def vmotion_network(self) -> pulumi.Output[str]:
        """
        Used for live migration of virtual machines
        """
        return pulumi.get(self, "vmotion_network")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

