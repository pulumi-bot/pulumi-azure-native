# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class VpnGateway(pulumi.CustomResource):
    etag: pulumi.Output[str]
    """
    A unique read-only string that changes whenever the resource is updated.
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
    Properties of the VPN gateway.
      * `bgp_settings` (`dict`) - Local network gateway's BGP speaker settings.
        * `asn` (`float`) - The BGP speaker's ASN.
        * `bgp_peering_address` (`str`) - The BGP peering address and BGP identifier of this BGP speaker.
        * `peer_weight` (`float`) - The weight added to routes learned from this BGP speaker.

      * `connections` (`list`) - List of all vpn connections to the gateway.
        * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
        * `id` (`str`) - Resource ID.
        * `name` (`str`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
        * `properties` (`dict`) - Properties of the VPN connection.
          * `connection_bandwidth` (`float`) - Expected bandwidth in MBPS.
          * `connection_status` (`str`) - The connection status.
          * `egress_bytes_transferred` (`float`) - Egress bytes transferred.
          * `enable_bgp` (`bool`) - EnableBgp flag.
          * `enable_internet_security` (`bool`) - Enable internet security.
          * `enable_rate_limiting` (`bool`) - EnableBgp flag.
          * `ingress_bytes_transferred` (`float`) - Ingress bytes transferred.
          * `ipsec_policies` (`list`) - The IPSec Policies to be considered by this connection.
            * `dh_group` (`str`) - The DH Group used in IKE Phase 1 for initial SA.
            * `ike_encryption` (`str`) - The IKE encryption algorithm (IKE phase 2).
            * `ike_integrity` (`str`) - The IKE integrity algorithm (IKE phase 2).
            * `ipsec_encryption` (`str`) - The IPSec encryption algorithm (IKE phase 1).
            * `ipsec_integrity` (`str`) - The IPSec integrity algorithm (IKE phase 1).
            * `pfs_group` (`str`) - The Pfs Group used in IKE Phase 2 for new child SA.
            * `sa_data_size_kilobytes` (`float`) - The IPSec Security Association (also called Quick Mode or Phase 2 SA) payload size in KB for a site to site VPN tunnel.
            * `sa_life_time_seconds` (`float`) - The IPSec Security Association (also called Quick Mode or Phase 2 SA) lifetime in seconds for a site to site VPN tunnel.

          * `provisioning_state` (`str`) - The provisioning state of the VPN connection resource.
          * `remote_vpn_site` (`dict`) - Id of the connected vpn site.
            * `id` (`str`) - Resource ID.

          * `routing_weight` (`float`) - Routing weight for vpn connection.
          * `shared_key` (`str`) - SharedKey for the vpn connection.
          * `use_local_azure_ip_address` (`bool`) - Use local azure ip to initiate connection.
          * `use_policy_based_traffic_selectors` (`bool`) - Enable policy-based traffic selectors.
          * `vpn_connection_protocol_type` (`str`) - Connection protocol used for this connection.
          * `vpn_link_connections` (`list`) - List of all vpn site link connections to the gateway.
            * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
            * `id` (`str`) - Resource ID.
            * `name` (`str`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
            * `properties` (`dict`) - Properties of the VPN site link connection.
              * `connection_bandwidth` (`float`) - Expected bandwidth in MBPS.
              * `connection_status` (`str`) - The connection status.
              * `egress_bytes_transferred` (`float`) - Egress bytes transferred.
              * `enable_bgp` (`bool`) - EnableBgp flag.
              * `enable_rate_limiting` (`bool`) - EnableBgp flag.
              * `ingress_bytes_transferred` (`float`) - Ingress bytes transferred.
              * `ipsec_policies` (`list`) - The IPSec Policies to be considered by this connection.
              * `provisioning_state` (`str`) - The provisioning state of the VPN site link connection resource.
              * `routing_weight` (`float`) - Routing weight for vpn connection.
              * `shared_key` (`str`) - SharedKey for the vpn connection.
              * `use_local_azure_ip_address` (`bool`) - Use local azure ip to initiate connection.
              * `use_policy_based_traffic_selectors` (`bool`) - Enable policy-based traffic selectors.
              * `vpn_connection_protocol_type` (`str`) - Connection protocol used for this connection.
              * `vpn_site_link` (`dict`) - Id of the connected vpn site link.

            * `type` (`str`) - Resource type.

      * `provisioning_state` (`str`) - The provisioning state of the VPN gateway resource.
      * `virtual_hub` (`dict`) - The VirtualHub to which the gateway belongs.
      * `vpn_gateway_scale_unit` (`float`) - The scale unit for this vpn gateway.
    """
    tags: pulumi.Output[dict]
    """
    Resource tags.
    """
    type: pulumi.Output[str]
    """
    Resource type.
    """
    def __init__(__self__, resource_name, opts=None, bgp_settings=None, connections=None, id=None, location=None, name=None, resource_group_name=None, tags=None, virtual_hub=None, vpn_gateway_scale_unit=None, __props__=None, __name__=None, __opts__=None):
        """
        VpnGateway Resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] bgp_settings: Local network gateway's BGP speaker settings.
        :param pulumi.Input[list] connections: List of all vpn connections to the gateway.
        :param pulumi.Input[str] id: Resource ID.
        :param pulumi.Input[str] location: Resource location.
        :param pulumi.Input[str] name: The name of the gateway.
        :param pulumi.Input[str] resource_group_name: The resource group name of the VpnGateway.
        :param pulumi.Input[dict] tags: Resource tags.
        :param pulumi.Input[dict] virtual_hub: The VirtualHub to which the gateway belongs.
        :param pulumi.Input[float] vpn_gateway_scale_unit: The scale unit for this vpn gateway.

        The **bgp_settings** object supports the following:

          * `asn` (`pulumi.Input[float]`) - The BGP speaker's ASN.
          * `bgp_peering_address` (`pulumi.Input[str]`) - The BGP peering address and BGP identifier of this BGP speaker.
          * `peer_weight` (`pulumi.Input[float]`) - The weight added to routes learned from this BGP speaker.

        The **connections** object supports the following:

          * `etag` (`pulumi.Input[str]`) - A unique read-only string that changes whenever the resource is updated.
          * `name` (`pulumi.Input[str]`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
          * `properties` (`pulumi.Input[dict]`) - Properties of the VPN connection.
            * `connection_bandwidth` (`pulumi.Input[float]`) - Expected bandwidth in MBPS.
            * `connection_status` (`pulumi.Input[str]`) - The connection status.
            * `egress_bytes_transferred` (`pulumi.Input[float]`) - Egress bytes transferred.
            * `enable_bgp` (`pulumi.Input[bool]`) - EnableBgp flag.
            * `enable_internet_security` (`pulumi.Input[bool]`) - Enable internet security.
            * `enable_rate_limiting` (`pulumi.Input[bool]`) - EnableBgp flag.
            * `ingress_bytes_transferred` (`pulumi.Input[float]`) - Ingress bytes transferred.
            * `ipsec_policies` (`pulumi.Input[list]`) - The IPSec Policies to be considered by this connection.
              * `dh_group` (`pulumi.Input[str]`) - The DH Group used in IKE Phase 1 for initial SA.
              * `ike_encryption` (`pulumi.Input[str]`) - The IKE encryption algorithm (IKE phase 2).
              * `ike_integrity` (`pulumi.Input[str]`) - The IKE integrity algorithm (IKE phase 2).
              * `ipsec_encryption` (`pulumi.Input[str]`) - The IPSec encryption algorithm (IKE phase 1).
              * `ipsec_integrity` (`pulumi.Input[str]`) - The IPSec integrity algorithm (IKE phase 1).
              * `pfs_group` (`pulumi.Input[str]`) - The Pfs Group used in IKE Phase 2 for new child SA.
              * `sa_data_size_kilobytes` (`pulumi.Input[float]`) - The IPSec Security Association (also called Quick Mode or Phase 2 SA) payload size in KB for a site to site VPN tunnel.
              * `sa_life_time_seconds` (`pulumi.Input[float]`) - The IPSec Security Association (also called Quick Mode or Phase 2 SA) lifetime in seconds for a site to site VPN tunnel.

            * `provisioning_state` (`pulumi.Input[str]`) - The provisioning state of the VPN connection resource.
            * `remote_vpn_site` (`pulumi.Input[dict]`) - Id of the connected vpn site.
              * `id` (`pulumi.Input[str]`) - Resource ID.

            * `routing_weight` (`pulumi.Input[float]`) - Routing weight for vpn connection.
            * `shared_key` (`pulumi.Input[str]`) - SharedKey for the vpn connection.
            * `use_local_azure_ip_address` (`pulumi.Input[bool]`) - Use local azure ip to initiate connection.
            * `use_policy_based_traffic_selectors` (`pulumi.Input[bool]`) - Enable policy-based traffic selectors.
            * `vpn_connection_protocol_type` (`pulumi.Input[str]`) - Connection protocol used for this connection.
            * `vpn_link_connections` (`pulumi.Input[list]`) - List of all vpn site link connections to the gateway.
              * `etag` (`pulumi.Input[str]`) - A unique read-only string that changes whenever the resource is updated.
              * `id` (`pulumi.Input[str]`) - Resource ID.
              * `name` (`pulumi.Input[str]`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
              * `properties` (`pulumi.Input[dict]`) - Properties of the VPN site link connection.
                * `connection_bandwidth` (`pulumi.Input[float]`) - Expected bandwidth in MBPS.
                * `connection_status` (`pulumi.Input[str]`) - The connection status.
                * `egress_bytes_transferred` (`pulumi.Input[float]`) - Egress bytes transferred.
                * `enable_bgp` (`pulumi.Input[bool]`) - EnableBgp flag.
                * `enable_rate_limiting` (`pulumi.Input[bool]`) - EnableBgp flag.
                * `ingress_bytes_transferred` (`pulumi.Input[float]`) - Ingress bytes transferred.
                * `ipsec_policies` (`pulumi.Input[list]`) - The IPSec Policies to be considered by this connection.
                * `provisioning_state` (`pulumi.Input[str]`) - The provisioning state of the VPN site link connection resource.
                * `routing_weight` (`pulumi.Input[float]`) - Routing weight for vpn connection.
                * `shared_key` (`pulumi.Input[str]`) - SharedKey for the vpn connection.
                * `use_local_azure_ip_address` (`pulumi.Input[bool]`) - Use local azure ip to initiate connection.
                * `use_policy_based_traffic_selectors` (`pulumi.Input[bool]`) - Enable policy-based traffic selectors.
                * `vpn_connection_protocol_type` (`pulumi.Input[str]`) - Connection protocol used for this connection.
                * `vpn_site_link` (`pulumi.Input[dict]`) - Id of the connected vpn site link.

              * `type` (`pulumi.Input[str]`) - Resource type.

        The **virtual_hub** object supports the following:

          * `id` (`pulumi.Input[str]`) - Resource ID.
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
            __props__['connections'] = connections
            __props__['id'] = id
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['virtual_hub'] = virtual_hub
            __props__['vpn_gateway_scale_unit'] = vpn_gateway_scale_unit
            __props__['etag'] = None
            __props__['properties'] = None
            __props__['type'] = None
        super(VpnGateway, __self__).__init__(
            'azurerm:network/v20191101:VpnGateway',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing VpnGateway resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return VpnGateway(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
