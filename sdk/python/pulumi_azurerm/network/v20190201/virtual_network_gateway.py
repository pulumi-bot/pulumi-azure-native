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
    Gets a unique read-only string that changes whenever the resource is updated.
    """
    location: pulumi.Output[str]
    """
    Resource location.
    """
    name: pulumi.Output[str]
    """
    Resource name.
    """
    properties: pulumi.Output[dict]
    """
    Properties of the virtual network gateway.
      * `active_active` (`bool`) - ActiveActive flag
      * `bgp_settings` (`dict`) - Virtual network gateway's BGP speaker settings.
        * `asn` (`float`) - The BGP speaker's ASN.
        * `bgp_peering_address` (`str`) - The BGP peering address and BGP identifier of this BGP speaker.
        * `peer_weight` (`float`) - The weight added to routes learned from this BGP speaker.

      * `custom_routes` (`dict`) - The reference of the address space resource which represents the custom routes address space specified by the customer for virtual network gateway and VpnClient.
        * `address_prefixes` (`list`) - A list of address blocks reserved for this virtual network in CIDR notation.

      * `enable_bgp` (`bool`) - Whether BGP is enabled for this virtual network gateway or not.
      * `gateway_default_site` (`dict`) - The reference of the LocalNetworkGateway resource which represents local network site having default routes. Assign Null value in case of removing existing default site setting.
        * `id` (`str`) - Resource ID.

      * `gateway_type` (`str`) - The type of this virtual network gateway. Possible values are: 'Vpn' and 'ExpressRoute'.
      * `ip_configurations` (`list`) - IP configurations for virtual network gateway.
        * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
        * `id` (`str`) - Resource ID.
        * `name` (`str`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
        * `properties` (`dict`) - Properties of the virtual network gateway ip configuration.
          * `private_ip_allocation_method` (`str`) - The private IP address allocation method.
          * `provisioning_state` (`str`) - The provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
          * `public_ip_address` (`dict`) - The reference of the public IP resource.
          * `subnet` (`dict`) - The reference of the subnet resource.

      * `provisioning_state` (`str`) - The provisioning state of the VirtualNetworkGateway resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
      * `resource_guid` (`str`) - The resource GUID property of the VirtualNetworkGateway resource.
      * `sku` (`dict`) - The reference of the VirtualNetworkGatewaySku resource which represents the SKU selected for Virtual network gateway.
        * `capacity` (`float`) - The capacity.
        * `name` (`str`) - Gateway SKU name.
        * `tier` (`str`) - Gateway SKU tier.

      * `vpn_client_configuration` (`dict`) - The reference of the VpnClientConfiguration resource which represents the P2S VpnClient configurations.
        * `radius_server_address` (`str`) - The radius server address property of the VirtualNetworkGateway resource for vpn client connection.
        * `radius_server_secret` (`str`) - The radius secret property of the VirtualNetworkGateway resource for vpn client connection.
        * `vpn_client_address_pool` (`dict`) - The reference of the address space resource which represents Address space for P2S VpnClient.
        * `vpn_client_ipsec_policies` (`list`) - VpnClientIpsecPolicies for virtual network gateway P2S client.
          * `dh_group` (`str`) - The DH Group used in IKE Phase 1 for initial SA.
          * `ike_encryption` (`str`) - The IKE encryption algorithm (IKE phase 2).
          * `ike_integrity` (`str`) - The IKE integrity algorithm (IKE phase 2).
          * `ipsec_encryption` (`str`) - The IPSec encryption algorithm (IKE phase 1).
          * `ipsec_integrity` (`str`) - The IPSec integrity algorithm (IKE phase 1).
          * `pfs_group` (`str`) - The Pfs Group used in IKE Phase 2 for new child SA.
          * `sa_data_size_kilobytes` (`float`) - The IPSec Security Association (also called Quick Mode or Phase 2 SA) payload size in KB for a site to site VPN tunnel.
          * `sa_life_time_seconds` (`float`) - The IPSec Security Association (also called Quick Mode or Phase 2 SA) lifetime in seconds for a site to site VPN tunnel.

        * `vpn_client_protocols` (`list`) - VpnClientProtocols for Virtual network gateway.
        * `vpn_client_revoked_certificates` (`list`) - VpnClientRevokedCertificate for Virtual network gateway.
          * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
          * `id` (`str`) - Resource ID.
          * `name` (`str`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
          * `properties` (`dict`) - Properties of the vpn client revoked certificate.
            * `provisioning_state` (`str`) - The provisioning state of the VPN client revoked certificate resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
            * `thumbprint` (`str`) - The revoked VPN client certificate thumbprint.

        * `vpn_client_root_certificates` (`list`) - VpnClientRootCertificate for virtual network gateway.
          * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
          * `id` (`str`) - Resource ID.
          * `name` (`str`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
          * `properties` (`dict`) - Properties of the vpn client root certificate.
            * `provisioning_state` (`str`) - The provisioning state of the VPN client root certificate resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
            * `public_cert_data` (`str`) - The certificate public data.

      * `vpn_type` (`str`) - The type of this virtual network gateway. Possible values are: 'PolicyBased' and 'RouteBased'.
    """
    tags: pulumi.Output[dict]
    """
    Resource tags.
    """
    type: pulumi.Output[str]
    """
    Resource type.
    """
    def __init__(__self__, resource_name, opts=None, active_active=None, bgp_settings=None, custom_routes=None, enable_bgp=None, etag=None, gateway_default_site=None, gateway_type=None, id=None, ip_configurations=None, location=None, name=None, resource_group_name=None, resource_guid=None, sku=None, tags=None, vpn_client_configuration=None, vpn_type=None, __props__=None, __name__=None, __opts__=None):
        """
        A common class for general resource information

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] active_active: ActiveActive flag
        :param pulumi.Input[dict] bgp_settings: Virtual network gateway's BGP speaker settings.
        :param pulumi.Input[dict] custom_routes: The reference of the address space resource which represents the custom routes address space specified by the customer for virtual network gateway and VpnClient.
        :param pulumi.Input[bool] enable_bgp: Whether BGP is enabled for this virtual network gateway or not.
        :param pulumi.Input[str] etag: Gets a unique read-only string that changes whenever the resource is updated.
        :param pulumi.Input[dict] gateway_default_site: The reference of the LocalNetworkGateway resource which represents local network site having default routes. Assign Null value in case of removing existing default site setting.
        :param pulumi.Input[str] gateway_type: The type of this virtual network gateway. Possible values are: 'Vpn' and 'ExpressRoute'.
        :param pulumi.Input[str] id: Resource ID.
        :param pulumi.Input[list] ip_configurations: IP configurations for virtual network gateway.
        :param pulumi.Input[str] location: Resource location.
        :param pulumi.Input[str] name: The name of the virtual network gateway.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[str] resource_guid: The resource GUID property of the VirtualNetworkGateway resource.
        :param pulumi.Input[dict] sku: The reference of the VirtualNetworkGatewaySku resource which represents the SKU selected for Virtual network gateway.
        :param pulumi.Input[dict] tags: Resource tags.
        :param pulumi.Input[dict] vpn_client_configuration: The reference of the VpnClientConfiguration resource which represents the P2S VpnClient configurations.
        :param pulumi.Input[str] vpn_type: The type of this virtual network gateway. Possible values are: 'PolicyBased' and 'RouteBased'.

        The **bgp_settings** object supports the following:

          * `asn` (`pulumi.Input[float]`) - The BGP speaker's ASN.
          * `bgp_peering_address` (`pulumi.Input[str]`) - The BGP peering address and BGP identifier of this BGP speaker.
          * `peer_weight` (`pulumi.Input[float]`) - The weight added to routes learned from this BGP speaker.

        The **custom_routes** object supports the following:

          * `address_prefixes` (`pulumi.Input[list]`) - A list of address blocks reserved for this virtual network in CIDR notation.

        The **gateway_default_site** object supports the following:

          * `id` (`pulumi.Input[str]`) - Resource ID.

        The **ip_configurations** object supports the following:

          * `etag` (`pulumi.Input[str]`) - A unique read-only string that changes whenever the resource is updated.
          * `id` (`pulumi.Input[str]`) - Resource ID.
          * `name` (`pulumi.Input[str]`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
          * `private_ip_allocation_method` (`pulumi.Input[str]`) - The private IP address allocation method.
          * `public_ip_address` (`pulumi.Input[dict]`) - The reference of the public IP resource.
          * `subnet` (`pulumi.Input[dict]`) - The reference of the subnet resource.

        The **sku** object supports the following:

          * `capacity` (`pulumi.Input[float]`) - The capacity.
          * `name` (`pulumi.Input[str]`) - Gateway SKU name.
          * `tier` (`pulumi.Input[str]`) - Gateway SKU tier.

        The **vpn_client_configuration** object supports the following:

          * `radius_server_address` (`pulumi.Input[str]`) - The radius server address property of the VirtualNetworkGateway resource for vpn client connection.
          * `radius_server_secret` (`pulumi.Input[str]`) - The radius secret property of the VirtualNetworkGateway resource for vpn client connection.
          * `vpn_client_address_pool` (`pulumi.Input[dict]`) - The reference of the address space resource which represents Address space for P2S VpnClient.
          * `vpn_client_ipsec_policies` (`pulumi.Input[list]`) - VpnClientIpsecPolicies for virtual network gateway P2S client.
            * `dh_group` (`pulumi.Input[str]`) - The DH Group used in IKE Phase 1 for initial SA.
            * `ike_encryption` (`pulumi.Input[str]`) - The IKE encryption algorithm (IKE phase 2).
            * `ike_integrity` (`pulumi.Input[str]`) - The IKE integrity algorithm (IKE phase 2).
            * `ipsec_encryption` (`pulumi.Input[str]`) - The IPSec encryption algorithm (IKE phase 1).
            * `ipsec_integrity` (`pulumi.Input[str]`) - The IPSec integrity algorithm (IKE phase 1).
            * `pfs_group` (`pulumi.Input[str]`) - The Pfs Group used in IKE Phase 2 for new child SA.
            * `sa_data_size_kilobytes` (`pulumi.Input[float]`) - The IPSec Security Association (also called Quick Mode or Phase 2 SA) payload size in KB for a site to site VPN tunnel.
            * `sa_life_time_seconds` (`pulumi.Input[float]`) - The IPSec Security Association (also called Quick Mode or Phase 2 SA) lifetime in seconds for a site to site VPN tunnel.

          * `vpn_client_protocols` (`pulumi.Input[list]`) - VpnClientProtocols for Virtual network gateway.
          * `vpn_client_revoked_certificates` (`pulumi.Input[list]`) - VpnClientRevokedCertificate for Virtual network gateway.
            * `etag` (`pulumi.Input[str]`) - A unique read-only string that changes whenever the resource is updated.
            * `id` (`pulumi.Input[str]`) - Resource ID.
            * `name` (`pulumi.Input[str]`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
            * `thumbprint` (`pulumi.Input[str]`) - The revoked VPN client certificate thumbprint.

          * `vpn_client_root_certificates` (`pulumi.Input[list]`) - VpnClientRootCertificate for virtual network gateway.
            * `etag` (`pulumi.Input[str]`) - A unique read-only string that changes whenever the resource is updated.
            * `id` (`pulumi.Input[str]`) - Resource ID.
            * `name` (`pulumi.Input[str]`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
            * `public_cert_data` (`pulumi.Input[str]`) - The certificate public data.
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

            __props__['active_active'] = active_active
            __props__['bgp_settings'] = bgp_settings
            __props__['custom_routes'] = custom_routes
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
            'azurerm:network/v20190201:VirtualNetworkGateway',
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
