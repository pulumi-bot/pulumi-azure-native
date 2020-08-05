# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class VirtualNetworkGateway(pulumi.CustomResource):
    etag: pulumi.Output[str]
    """
    Gets a unique read-only string that changes whenever the resource is updated
    """
    location: pulumi.Output[str]
    """
    Resource location
    """
    name: pulumi.Output[str]
    """
    Resource name
    """
    properties: pulumi.Output[dict]
    """
    VirtualNetworkGateway properties
      * `bgp_settings` (`dict`) - Virtual network gateway's BGP speaker settings
        * `asn` (`float`) - Gets or sets this BGP speaker's ASN
        * `bgp_peering_address` (`str`) - Gets or sets the BGP peering address and BGP identifier of this BGP speaker
        * `peer_weight` (`float`) - Gets or sets the weight added to routes learned from this BGP speaker

      * `enable_bgp` (`bool`) - EnableBgp Flag
      * `gateway_default_site` (`dict`) - Gets or sets the reference of the LocalNetworkGateway resource which represents Local network site having default routes. Assign Null value in case of removing existing default site setting.
        * `id` (`str`) - Resource Id

      * `gateway_type` (`str`) - The type of this virtual network gateway.
      * `ip_configurations` (`list`) - IpConfigurations for Virtual network gateway.
        * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated
        * `id` (`str`) - Resource Id
        * `name` (`str`) - Gets name of the resource that is unique within a resource group. This name can be used to access the resource
        * `properties` (`dict`) - Properties of VirtualNetworkGatewayIPConfiguration
          * `private_ip_address` (`str`) - Gets or sets the privateIPAddress of the IP Configuration
          * `private_ip_allocation_method` (`str`) - Gets or sets PrivateIP allocation method (Static/Dynamic)
          * `provisioning_state` (`str`) - Gets or sets Provisioning state of the PublicIP resource Updating/Deleting/Failed
          * `public_ip_address` (`dict`) - Gets or sets the reference of the PublicIP resource
          * `subnet` (`dict`) - Gets or sets the reference of the subnet resource

      * `provisioning_state` (`str`) - Gets or sets Provisioning state of the VirtualNetworkGateway resource Updating/Deleting/Failed
      * `resource_guid` (`str`) - Gets or sets resource GUID property of the VirtualNetworkGateway resource
      * `sku` (`dict`) - Gets or sets the reference of the VirtualNetworkGatewaySku resource which represents the sku selected for Virtual network gateway.
        * `capacity` (`float`) - The capacity
        * `name` (`str`) - Gateway sku name -Basic/HighPerformance/Standard
        * `tier` (`str`) - Gateway sku tier -Basic/HighPerformance/Standard

      * `vpn_client_configuration` (`dict`) - Gets or sets the reference of the VpnClientConfiguration resource which represents the P2S VpnClient configurations.
        * `vpn_client_address_pool` (`dict`) - Gets or sets the reference of the Address space resource which represents Address space for P2S VpnClient.
          * `address_prefixes` (`list`) - Gets or sets List of address blocks reserved for this virtual network in CIDR notation

        * `vpn_client_revoked_certificates` (`list`) - VpnClientRevokedCertificate for Virtual network gateway.
          * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated
          * `id` (`str`) - Resource Id
          * `name` (`str`) - Gets name of the resource that is unique within a resource group. This name can be used to access the resource
          * `properties` (`dict`) - Properties of the revoked VPN client certificate of virtual network gateway
            * `provisioning_state` (`str`) - Gets or sets Provisioning state of the VPN client revoked certificate resource Updating/Deleting/Failed
            * `thumbprint` (`str`) - Gets or sets the revoked Vpn client certificate thumbprint

        * `vpn_client_root_certificates` (`list`) - VpnClientRootCertificate for Virtual network gateway.
          * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated
          * `id` (`str`) - Resource Id
          * `name` (`str`) - Gets name of the resource that is unique within a resource group. This name can be used to access the resource
          * `properties` (`dict`) - Properties of SSL certificates of application gateway
            * `provisioning_state` (`str`) - Gets or sets Provisioning state of the VPN client root certificate resource Updating/Deleting/Failed
            * `public_cert_data` (`str`) - Gets or sets the certificate public data

      * `vpn_type` (`str`) - The type of this virtual network gateway.
    """
    tags: pulumi.Output[dict]
    """
    Resource tags
    """
    type: pulumi.Output[str]
    """
    Resource type
    """
    def __init__(__self__, resource_name, opts=None, bgp_settings=None, enable_bgp=None, etag=None, gateway_default_site=None, gateway_type=None, id=None, ip_configurations=None, location=None, name=None, provisioning_state=None, resource_group_name=None, resource_guid=None, sku=None, tags=None, vpn_client_configuration=None, vpn_type=None, __props__=None, __name__=None, __opts__=None):
        """
        A common class for general resource information

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] bgp_settings: Virtual network gateway's BGP speaker settings
        :param pulumi.Input[bool] enable_bgp: EnableBgp Flag
        :param pulumi.Input[str] etag: Gets a unique read-only string that changes whenever the resource is updated
        :param pulumi.Input[dict] gateway_default_site: Gets or sets the reference of the LocalNetworkGateway resource which represents Local network site having default routes. Assign Null value in case of removing existing default site setting.
        :param pulumi.Input[str] gateway_type: The type of this virtual network gateway.
        :param pulumi.Input[str] id: Resource Id
        :param pulumi.Input[list] ip_configurations: IpConfigurations for Virtual network gateway.
        :param pulumi.Input[str] location: Resource location
        :param pulumi.Input[str] name: The name of the virtual network gateway.
        :param pulumi.Input[str] provisioning_state: Gets or sets Provisioning state of the VirtualNetworkGateway resource Updating/Deleting/Failed
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[str] resource_guid: Gets or sets resource GUID property of the VirtualNetworkGateway resource
        :param pulumi.Input[dict] sku: Gets or sets the reference of the VirtualNetworkGatewaySku resource which represents the sku selected for Virtual network gateway.
        :param pulumi.Input[dict] tags: Resource tags
        :param pulumi.Input[dict] vpn_client_configuration: Gets or sets the reference of the VpnClientConfiguration resource which represents the P2S VpnClient configurations.
        :param pulumi.Input[str] vpn_type: The type of this virtual network gateway.

        The **bgp_settings** object supports the following:

          * `asn` (`pulumi.Input[float]`) - Gets or sets this BGP speaker's ASN
          * `bgp_peering_address` (`pulumi.Input[str]`) - Gets or sets the BGP peering address and BGP identifier of this BGP speaker
          * `peer_weight` (`pulumi.Input[float]`) - Gets or sets the weight added to routes learned from this BGP speaker

        The **gateway_default_site** object supports the following:

          * `id` (`pulumi.Input[str]`) - Resource Id

        The **ip_configurations** object supports the following:

          * `etag` (`pulumi.Input[str]`) - A unique read-only string that changes whenever the resource is updated
          * `id` (`pulumi.Input[str]`) - Resource Id
          * `name` (`pulumi.Input[str]`) - Gets name of the resource that is unique within a resource group. This name can be used to access the resource
          * `private_ip_address` (`pulumi.Input[str]`) - Gets or sets the privateIPAddress of the IP Configuration
          * `private_ip_allocation_method` (`pulumi.Input[str]`) - Gets or sets PrivateIP allocation method (Static/Dynamic)
          * `provisioning_state` (`pulumi.Input[str]`) - Gets or sets Provisioning state of the PublicIP resource Updating/Deleting/Failed
          * `public_ip_address` (`pulumi.Input[dict]`) - Gets or sets the reference of the PublicIP resource
          * `subnet` (`pulumi.Input[dict]`) - Gets or sets the reference of the subnet resource

        The **sku** object supports the following:

          * `capacity` (`pulumi.Input[float]`) - The capacity
          * `name` (`pulumi.Input[str]`) - Gateway sku name -Basic/HighPerformance/Standard
          * `tier` (`pulumi.Input[str]`) - Gateway sku tier -Basic/HighPerformance/Standard

        The **vpn_client_configuration** object supports the following:

          * `vpn_client_address_pool` (`pulumi.Input[dict]`) - Gets or sets the reference of the Address space resource which represents Address space for P2S VpnClient.
            * `address_prefixes` (`pulumi.Input[list]`) - Gets or sets List of address blocks reserved for this virtual network in CIDR notation

          * `vpn_client_revoked_certificates` (`pulumi.Input[list]`) - VpnClientRevokedCertificate for Virtual network gateway.
            * `etag` (`pulumi.Input[str]`) - A unique read-only string that changes whenever the resource is updated
            * `id` (`pulumi.Input[str]`) - Resource Id
            * `name` (`pulumi.Input[str]`) - Gets name of the resource that is unique within a resource group. This name can be used to access the resource
            * `provisioning_state` (`pulumi.Input[str]`) - Gets or sets Provisioning state of the VPN client revoked certificate resource Updating/Deleting/Failed
            * `thumbprint` (`pulumi.Input[str]`) - Gets or sets the revoked Vpn client certificate thumbprint

          * `vpn_client_root_certificates` (`pulumi.Input[list]`) - VpnClientRootCertificate for Virtual network gateway.
            * `etag` (`pulumi.Input[str]`) - A unique read-only string that changes whenever the resource is updated
            * `id` (`pulumi.Input[str]`) - Resource Id
            * `name` (`pulumi.Input[str]`) - Gets name of the resource that is unique within a resource group. This name can be used to access the resource
            * `provisioning_state` (`pulumi.Input[str]`) - Gets or sets Provisioning state of the VPN client root certificate resource Updating/Deleting/Failed
            * `public_cert_data` (`pulumi.Input[str]`) - Gets or sets the certificate public data
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

            __props__['bgp_settings'] = bgp_settings
            __props__['enable_bgp'] = enable_bgp
            __props__['etag'] = etag
            __props__['gateway_default_site'] = gateway_default_site
            __props__['gateway_type'] = gateway_type
            __props__['id'] = id
            __props__['ip_configurations'] = ip_configurations
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['provisioning_state'] = provisioning_state
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['resource_guid'] = resource_guid
            __props__['sku'] = sku
            __props__['tags'] = tags
            __props__['vpn_client_configuration'] = vpn_client_configuration
            __props__['vpn_type'] = vpn_type
            __props__['properties'] = None
            __props__['type'] = None
        super(VirtualNetworkGateway, __self__).__init__(
            'azurerm:network/v20160330:VirtualNetworkGateway',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing VirtualNetworkGateway resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return VirtualNetworkGateway(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
