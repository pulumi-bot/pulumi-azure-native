# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class VirtualNetworkGatewayConnection(pulumi.CustomResource):
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
    Properties of the virtual network gateway connection.
      * `authorization_key` (`str`) - The authorizationKey.
      * `connection_protocol` (`str`) - Connection protocol used for this connection
      * `connection_status` (`str`) - Virtual network Gateway connection status. Possible values are 'Unknown', 'Connecting', 'Connected' and 'NotConnected'.
      * `connection_type` (`str`) - Gateway connection type. Possible values are: 'Ipsec','Vnet2Vnet','ExpressRoute', and 'VPNClient.
      * `egress_bytes_transferred` (`float`) - The egress bytes transferred in this connection.
      * `enable_bgp` (`bool`) - EnableBgp flag
      * `express_route_gateway_bypass` (`bool`) - Bypass ExpressRoute Gateway for data forwarding
      * `ingress_bytes_transferred` (`float`) - The ingress bytes transferred in this connection.
      * `ipsec_policies` (`list`) - The IPSec Policies to be considered by this connection.
        * `dh_group` (`str`) - The DH Groups used in IKE Phase 1 for initial SA.
        * `ike_encryption` (`str`) - The IKE encryption algorithm (IKE phase 2).
        * `ike_integrity` (`str`) - The IKE integrity algorithm (IKE phase 2).
        * `ipsec_encryption` (`str`) - The IPSec encryption algorithm (IKE phase 1).
        * `ipsec_integrity` (`str`) - The IPSec integrity algorithm (IKE phase 1).
        * `pfs_group` (`str`) - The Pfs Groups used in IKE Phase 2 for new child SA.
        * `sa_data_size_kilobytes` (`float`) - The IPSec Security Association (also called Quick Mode or Phase 2 SA) payload size in KB for a site to site VPN tunnel.
        * `sa_life_time_seconds` (`float`) - The IPSec Security Association (also called Quick Mode or Phase 2 SA) lifetime in seconds for a site to site VPN tunnel.

      * `local_network_gateway2` (`dict`) - The reference to local network gateway resource.
        * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
        * `id` (`str`) - Resource ID.
        * `location` (`str`) - Resource location.
        * `name` (`str`) - Resource name.
        * `properties` (`dict`) - Properties of the local network gateway.
          * `bgp_settings` (`dict`) - Local network gateway's BGP speaker settings.
            * `asn` (`float`) - The BGP speaker's ASN.
            * `bgp_peering_address` (`str`) - The BGP peering address and BGP identifier of this BGP speaker.
            * `peer_weight` (`float`) - The weight added to routes learned from this BGP speaker.

          * `gateway_ip_address` (`str`) - IP address of local network gateway.
          * `local_network_address_space` (`dict`) - Local network site address space.
            * `address_prefixes` (`list`) - A list of address blocks reserved for this virtual network in CIDR notation.

          * `provisioning_state` (`str`) - The provisioning state of the LocalNetworkGateway resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
          * `resource_guid` (`str`) - The resource GUID property of the LocalNetworkGateway resource.

        * `tags` (`dict`) - Resource tags.
        * `type` (`str`) - Resource type.

      * `peer` (`dict`) - The reference to peerings resource.
        * `id` (`str`) - Resource ID.

      * `provisioning_state` (`str`) - The provisioning state of the VirtualNetworkGatewayConnection resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
      * `resource_guid` (`str`) - The resource GUID property of the VirtualNetworkGatewayConnection resource.
      * `routing_weight` (`float`) - The routing weight.
      * `shared_key` (`str`) - The IPSec shared key.
      * `tunnel_connection_status` (`list`) - Collection of all tunnels' connection health status.
        * `connection_status` (`str`) - Virtual network Gateway connection status
        * `egress_bytes_transferred` (`float`) - The Egress Bytes Transferred in this connection
        * `ingress_bytes_transferred` (`float`) - The Ingress Bytes Transferred in this connection
        * `last_connection_established_utc_time` (`str`) - The time at which connection was established in Utc format.
        * `tunnel` (`str`) - Tunnel name.

      * `use_policy_based_traffic_selectors` (`bool`) - Enable policy-based traffic selectors.
      * `virtual_network_gateway1` (`dict`) - The reference to virtual network gateway resource.
        * `etag` (`str`) - Gets a unique read-only string that changes whenever the resource is updated.
        * `id` (`str`) - Resource ID.
        * `location` (`str`) - Resource location.
        * `name` (`str`) - Resource name.
        * `properties` (`dict`) - Properties of the virtual network gateway.
          * `active_active` (`bool`) - ActiveActive flag
          * `bgp_settings` (`dict`) - Virtual network gateway's BGP speaker settings.
          * `enable_bgp` (`bool`) - Whether BGP is enabled for this virtual network gateway or not.
          * `gateway_default_site` (`dict`) - The reference of the LocalNetworkGateway resource which represents local network site having default routes. Assign Null value in case of removing existing default site setting.
          * `gateway_type` (`str`) - The type of this virtual network gateway. Possible values are: 'Vpn' and 'ExpressRoute'.
          * `ip_configurations` (`list`) - IP configurations for virtual network gateway.
            * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
            * `id` (`str`) - Resource ID.
            * `name` (`str`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
            * `properties` (`dict`) - Properties of the virtual network gateway ip configuration.
              * `private_ip_allocation_method` (`str`) - The private IP allocation method. Possible values are: 'Static' and 'Dynamic'.
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

        * `tags` (`dict`) - Resource tags.
        * `type` (`str`) - Resource type.

      * `virtual_network_gateway2` (`dict`) - The reference to virtual network gateway resource.
    """
    tags: pulumi.Output[dict]
    """
    Resource tags.
    """
    type: pulumi.Output[str]
    """
    Resource type.
    """
    def __init__(__self__, resource_name, opts=None, etag=None, id=None, location=None, name=None, properties=None, resource_group_name=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        A common class for general resource information

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] etag: Gets a unique read-only string that changes whenever the resource is updated.
        :param pulumi.Input[str] id: Resource ID.
        :param pulumi.Input[str] location: Resource location.
        :param pulumi.Input[str] name: The name of the virtual network gateway connection.
        :param pulumi.Input[dict] properties: Properties of the virtual network gateway connection.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[dict] tags: Resource tags.

        The **properties** object supports the following:

          * `authorization_key` (`pulumi.Input[str]`) - The authorizationKey.
          * `connection_protocol` (`pulumi.Input[str]`) - Connection protocol used for this connection
          * `connection_type` (`pulumi.Input[str]`) - Gateway connection type. Possible values are: 'Ipsec','Vnet2Vnet','ExpressRoute', and 'VPNClient.
          * `enable_bgp` (`pulumi.Input[bool]`) - EnableBgp flag
          * `express_route_gateway_bypass` (`pulumi.Input[bool]`) - Bypass ExpressRoute Gateway for data forwarding
          * `ipsec_policies` (`pulumi.Input[list]`) - The IPSec Policies to be considered by this connection.
            * `dh_group` (`pulumi.Input[str]`) - The DH Groups used in IKE Phase 1 for initial SA.
            * `ike_encryption` (`pulumi.Input[str]`) - The IKE encryption algorithm (IKE phase 2).
            * `ike_integrity` (`pulumi.Input[str]`) - The IKE integrity algorithm (IKE phase 2).
            * `ipsec_encryption` (`pulumi.Input[str]`) - The IPSec encryption algorithm (IKE phase 1).
            * `ipsec_integrity` (`pulumi.Input[str]`) - The IPSec integrity algorithm (IKE phase 1).
            * `pfs_group` (`pulumi.Input[str]`) - The Pfs Groups used in IKE Phase 2 for new child SA.
            * `sa_data_size_kilobytes` (`pulumi.Input[float]`) - The IPSec Security Association (also called Quick Mode or Phase 2 SA) payload size in KB for a site to site VPN tunnel.
            * `sa_life_time_seconds` (`pulumi.Input[float]`) - The IPSec Security Association (also called Quick Mode or Phase 2 SA) lifetime in seconds for a site to site VPN tunnel.

          * `local_network_gateway2` (`pulumi.Input[dict]`) - The reference to local network gateway resource.
            * `etag` (`pulumi.Input[str]`) - A unique read-only string that changes whenever the resource is updated.
            * `location` (`pulumi.Input[str]`) - Resource location.
            * `name` (`pulumi.Input[str]`) - Resource name.
            * `properties` (`pulumi.Input[dict]`) - Properties of the local network gateway.
              * `bgp_settings` (`pulumi.Input[dict]`) - Local network gateway's BGP speaker settings.
                * `asn` (`pulumi.Input[float]`) - The BGP speaker's ASN.
                * `bgp_peering_address` (`pulumi.Input[str]`) - The BGP peering address and BGP identifier of this BGP speaker.
                * `peer_weight` (`pulumi.Input[float]`) - The weight added to routes learned from this BGP speaker.

              * `gateway_ip_address` (`pulumi.Input[str]`) - IP address of local network gateway.
              * `local_network_address_space` (`pulumi.Input[dict]`) - Local network site address space.
                * `address_prefixes` (`pulumi.Input[list]`) - A list of address blocks reserved for this virtual network in CIDR notation.

              * `provisioning_state` (`pulumi.Input[str]`) - The provisioning state of the LocalNetworkGateway resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
              * `resource_guid` (`pulumi.Input[str]`) - The resource GUID property of the LocalNetworkGateway resource.

            * `tags` (`pulumi.Input[dict]`) - Resource tags.
            * `type` (`pulumi.Input[str]`) - Resource type.

          * `peer` (`pulumi.Input[dict]`) - The reference to peerings resource.
            * `id` (`pulumi.Input[str]`) - Resource ID.

          * `resource_guid` (`pulumi.Input[str]`) - The resource GUID property of the VirtualNetworkGatewayConnection resource.
          * `routing_weight` (`pulumi.Input[float]`) - The routing weight.
          * `shared_key` (`pulumi.Input[str]`) - The IPSec shared key.
          * `use_policy_based_traffic_selectors` (`pulumi.Input[bool]`) - Enable policy-based traffic selectors.
          * `virtual_network_gateway1` (`pulumi.Input[dict]`) - The reference to virtual network gateway resource.
            * `etag` (`pulumi.Input[str]`) - Gets a unique read-only string that changes whenever the resource is updated.
            * `location` (`pulumi.Input[str]`) - Resource location.
            * `name` (`pulumi.Input[str]`) - Resource name.
            * `properties` (`pulumi.Input[dict]`) - Properties of the virtual network gateway.
              * `active_active` (`pulumi.Input[bool]`) - ActiveActive flag
              * `bgp_settings` (`pulumi.Input[dict]`) - Virtual network gateway's BGP speaker settings.
              * `enable_bgp` (`pulumi.Input[bool]`) - Whether BGP is enabled for this virtual network gateway or not.
              * `gateway_default_site` (`pulumi.Input[dict]`) - The reference of the LocalNetworkGateway resource which represents local network site having default routes. Assign Null value in case of removing existing default site setting.
                * `id` (`pulumi.Input[str]`) - Resource ID.

              * `gateway_type` (`pulumi.Input[str]`) - The type of this virtual network gateway. Possible values are: 'Vpn' and 'ExpressRoute'.
              * `ip_configurations` (`pulumi.Input[list]`) - IP configurations for virtual network gateway.
                * `etag` (`pulumi.Input[str]`) - A unique read-only string that changes whenever the resource is updated.
                * `id` (`pulumi.Input[str]`) - Resource ID.
                * `name` (`pulumi.Input[str]`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
                * `properties` (`pulumi.Input[dict]`) - Properties of the virtual network gateway ip configuration.
                  * `private_ip_allocation_method` (`pulumi.Input[str]`) - The private IP allocation method. Possible values are: 'Static' and 'Dynamic'.
                  * `provisioning_state` (`pulumi.Input[str]`) - The provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
                  * `public_ip_address` (`pulumi.Input[dict]`) - The reference of the public IP resource.
                  * `subnet` (`pulumi.Input[dict]`) - The reference of the subnet resource.

              * `provisioning_state` (`pulumi.Input[str]`) - The provisioning state of the VirtualNetworkGateway resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
              * `resource_guid` (`pulumi.Input[str]`) - The resource GUID property of the VirtualNetworkGateway resource.
              * `sku` (`pulumi.Input[dict]`) - The reference of the VirtualNetworkGatewaySku resource which represents the SKU selected for Virtual network gateway.
                * `capacity` (`pulumi.Input[float]`) - The capacity.
                * `name` (`pulumi.Input[str]`) - Gateway SKU name.
                * `tier` (`pulumi.Input[str]`) - Gateway SKU tier.

              * `vpn_client_configuration` (`pulumi.Input[dict]`) - The reference of the VpnClientConfiguration resource which represents the P2S VpnClient configurations.
                * `radius_server_address` (`pulumi.Input[str]`) - The radius server address property of the VirtualNetworkGateway resource for vpn client connection.
                * `radius_server_secret` (`pulumi.Input[str]`) - The radius secret property of the VirtualNetworkGateway resource for vpn client connection.
                * `vpn_client_address_pool` (`pulumi.Input[dict]`) - The reference of the address space resource which represents Address space for P2S VpnClient.
                * `vpn_client_ipsec_policies` (`pulumi.Input[list]`) - VpnClientIpsecPolicies for virtual network gateway P2S client.
                  * `dh_group` (`pulumi.Input[str]`) - The DH Groups used in IKE Phase 1 for initial SA.
                  * `ike_encryption` (`pulumi.Input[str]`) - The IKE encryption algorithm (IKE phase 2).
                  * `ike_integrity` (`pulumi.Input[str]`) - The IKE integrity algorithm (IKE phase 2).
                  * `ipsec_encryption` (`pulumi.Input[str]`) - The IPSec encryption algorithm (IKE phase 1).
                  * `ipsec_integrity` (`pulumi.Input[str]`) - The IPSec integrity algorithm (IKE phase 1).
                  * `pfs_group` (`pulumi.Input[str]`) - The Pfs Groups used in IKE Phase 2 for new child SA.
                  * `sa_data_size_kilobytes` (`pulumi.Input[float]`) - The IPSec Security Association (also called Quick Mode or Phase 2 SA) payload size in KB for a site to site VPN tunnel.
                  * `sa_life_time_seconds` (`pulumi.Input[float]`) - The IPSec Security Association (also called Quick Mode or Phase 2 SA) lifetime in seconds for a site to site VPN tunnel.

                * `vpn_client_protocols` (`pulumi.Input[list]`) - VpnClientProtocols for Virtual network gateway.
                * `vpn_client_revoked_certificates` (`pulumi.Input[list]`) - VpnClientRevokedCertificate for Virtual network gateway.
                  * `etag` (`pulumi.Input[str]`) - A unique read-only string that changes whenever the resource is updated.
                  * `id` (`pulumi.Input[str]`) - Resource ID.
                  * `name` (`pulumi.Input[str]`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
                  * `properties` (`pulumi.Input[dict]`) - Properties of the vpn client revoked certificate.
                    * `provisioning_state` (`pulumi.Input[str]`) - The provisioning state of the VPN client revoked certificate resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
                    * `thumbprint` (`pulumi.Input[str]`) - The revoked VPN client certificate thumbprint.

                * `vpn_client_root_certificates` (`pulumi.Input[list]`) - VpnClientRootCertificate for virtual network gateway.
                  * `etag` (`pulumi.Input[str]`) - A unique read-only string that changes whenever the resource is updated.
                  * `id` (`pulumi.Input[str]`) - Resource ID.
                  * `name` (`pulumi.Input[str]`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
                  * `properties` (`pulumi.Input[dict]`) - Properties of the vpn client root certificate.
                    * `provisioning_state` (`pulumi.Input[str]`) - The provisioning state of the VPN client root certificate resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
                    * `public_cert_data` (`pulumi.Input[str]`) - The certificate public data.

              * `vpn_type` (`pulumi.Input[str]`) - The type of this virtual network gateway. Possible values are: 'PolicyBased' and 'RouteBased'.

            * `tags` (`pulumi.Input[dict]`) - Resource tags.
            * `type` (`pulumi.Input[str]`) - Resource type.

          * `virtual_network_gateway2` (`pulumi.Input[dict]`) - The reference to virtual network gateway resource.
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

            __props__['etag'] = etag
            __props__['id'] = id
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if properties is None:
                raise TypeError("Missing required property 'properties'")
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['type'] = None
        super(VirtualNetworkGatewayConnection, __self__).__init__(
            'azurerm:network/v20181001:VirtualNetworkGatewayConnection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing VirtualNetworkGatewayConnection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return VirtualNetworkGatewayConnection(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
