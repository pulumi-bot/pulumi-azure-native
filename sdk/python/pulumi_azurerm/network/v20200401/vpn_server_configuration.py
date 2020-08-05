# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class VpnServerConfiguration(pulumi.CustomResource):
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
    Properties of the P2SVpnServer configuration.
      * `aad_authentication_parameters` (`dict`) - The set of aad vpn authentication parameters.
        * `aad_audience` (`str`) - AAD Vpn authentication parameter AAD audience.
        * `aad_issuer` (`str`) - AAD Vpn authentication parameter AAD issuer.
        * `aad_tenant` (`str`) - AAD Vpn authentication parameter AAD tenant.

      * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
      * `name` (`str`) - The name of the VpnServerConfiguration that is unique within a resource group.
      * `p2_s_vpn_gateways` (`list`) - List of references to P2SVpnGateways.
        * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
        * `id` (`str`) - Resource ID.
        * `location` (`str`) - Resource location.
        * `name` (`str`) - Resource name.
        * `properties` (`dict`) - Properties of the P2SVpnGateway.
          * `p2_s_connection_configurations` (`list`) - List of all p2s connection configurations of the gateway.
            * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
            * `id` (`str`) - Resource ID.
            * `name` (`str`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
            * `properties` (`dict`) - Properties of the P2S connection configuration.
              * `provisioning_state` (`str`) - The provisioning state of the P2SConnectionConfiguration resource.
              * `routing_configuration` (`dict`) - The Routing Configuration indicating the associated and propagated route tables on this connection.
                * `associated_route_table` (`dict`) - The resource id RouteTable associated with this RoutingConfiguration.
                  * `id` (`str`) - Resource ID.

                * `propagated_route_tables` (`dict`) - The list of RouteTables to advertise the routes to.
                  * `ids` (`list`) - The list of resource ids of all the RouteTables.
                  * `labels` (`list`) - The list of labels.

                * `vnet_routes` (`dict`) - List of routes that control routing from VirtualHub into a virtual network connection.
                  * `static_routes` (`list`) - List of all Static Routes.
                    * `address_prefixes` (`list`) - List of all address prefixes.
                    * `name` (`str`) - The name of the StaticRoute that is unique within a VnetRoute.
                    * `next_hop_ip_address` (`str`) - The ip address of the next hop.

              * `vpn_client_address_pool` (`dict`) - The reference to the address space resource which represents Address space for P2S VpnClient.
                * `address_prefixes` (`list`) - A list of address blocks reserved for this virtual network in CIDR notation.

          * `provisioning_state` (`str`) - The provisioning state of the P2S VPN gateway resource.
          * `virtual_hub` (`dict`) - The VirtualHub to which the gateway belongs.
          * `vpn_client_connection_health` (`dict`) - All P2S VPN clients' connection health status.
            * `allocated_ip_addresses` (`list`) - List of allocated ip addresses to the connected p2s vpn clients.
            * `total_egress_bytes_transferred` (`float`) - Total of the Egress Bytes Transferred in this connection.
            * `total_ingress_bytes_transferred` (`float`) - Total of the Ingress Bytes Transferred in this P2S Vpn connection.
            * `vpn_client_connections_count` (`float`) - The total of p2s vpn clients connected at this time to this P2SVpnGateway.

          * `vpn_gateway_scale_unit` (`float`) - The scale unit for this p2s vpn gateway.
          * `vpn_server_configuration` (`dict`) - The VpnServerConfiguration to which the p2sVpnGateway is attached to.

        * `tags` (`dict`) - Resource tags.
        * `type` (`str`) - Resource type.

      * `provisioning_state` (`str`) - The provisioning state of the VpnServerConfiguration resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
      * `radius_client_root_certificates` (`list`) - Radius client root certificate of VpnServerConfiguration.
        * `name` (`str`) - The certificate name.
        * `thumbprint` (`str`) - The Radius client root certificate thumbprint.

      * `radius_server_address` (`str`) - The radius server address property of the VpnServerConfiguration resource for point to site client connection.
      * `radius_server_root_certificates` (`list`) - Radius Server root certificate of VpnServerConfiguration.
        * `name` (`str`) - The certificate name.
        * `public_cert_data` (`str`) - The certificate public data.

      * `radius_server_secret` (`str`) - The radius secret property of the VpnServerConfiguration resource for point to site client connection.
      * `radius_servers` (`list`) - Multiple Radius Server configuration for VpnServerConfiguration.
        * `radius_server_address` (`str`) - The address of this radius server.
        * `radius_server_score` (`float`) - The initial score assigned to this radius server.
        * `radius_server_secret` (`str`) - The secret used for this radius server.

      * `vpn_authentication_types` (`list`) - VPN authentication types for the VpnServerConfiguration.
      * `vpn_client_ipsec_policies` (`list`) - VpnClientIpsecPolicies for VpnServerConfiguration.
        * `dh_group` (`str`) - The DH Group used in IKE Phase 1 for initial SA.
        * `ike_encryption` (`str`) - The IKE encryption algorithm (IKE phase 2).
        * `ike_integrity` (`str`) - The IKE integrity algorithm (IKE phase 2).
        * `ipsec_encryption` (`str`) - The IPSec encryption algorithm (IKE phase 1).
        * `ipsec_integrity` (`str`) - The IPSec integrity algorithm (IKE phase 1).
        * `pfs_group` (`str`) - The Pfs Group used in IKE Phase 2 for new child SA.
        * `sa_data_size_kilobytes` (`float`) - The IPSec Security Association (also called Quick Mode or Phase 2 SA) payload size in KB for a site to site VPN tunnel.
        * `sa_life_time_seconds` (`float`) - The IPSec Security Association (also called Quick Mode or Phase 2 SA) lifetime in seconds for a site to site VPN tunnel.

      * `vpn_client_revoked_certificates` (`list`) - VPN client revoked certificate of VpnServerConfiguration.
        * `name` (`str`) - The certificate name.
        * `thumbprint` (`str`) - The revoked VPN client certificate thumbprint.

      * `vpn_client_root_certificates` (`list`) - VPN client root certificate of VpnServerConfiguration.
        * `name` (`str`) - The certificate name.
        * `public_cert_data` (`str`) - The certificate public data.

      * `vpn_protocols` (`list`) - VPN protocols for the VpnServerConfiguration.
    """
    tags: pulumi.Output[dict]
    """
    Resource tags.
    """
    type: pulumi.Output[str]
    """
    Resource type.
    """
    def __init__(__self__, resource_name, opts=None, aad_authentication_parameters=None, id=None, location=None, name=None, radius_client_root_certificates=None, radius_server_address=None, radius_server_root_certificates=None, radius_server_secret=None, radius_servers=None, resource_group_name=None, tags=None, vpn_authentication_types=None, vpn_client_ipsec_policies=None, vpn_client_revoked_certificates=None, vpn_client_root_certificates=None, vpn_protocols=None, __props__=None, __name__=None, __opts__=None):
        """
        VpnServerConfiguration Resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] aad_authentication_parameters: The set of aad vpn authentication parameters.
        :param pulumi.Input[str] id: Resource ID.
        :param pulumi.Input[str] location: Resource location.
        :param pulumi.Input[str] name: The name of the VpnServerConfiguration being created or updated.
        :param pulumi.Input[list] radius_client_root_certificates: Radius client root certificate of VpnServerConfiguration.
        :param pulumi.Input[str] radius_server_address: The radius server address property of the VpnServerConfiguration resource for point to site client connection.
        :param pulumi.Input[list] radius_server_root_certificates: Radius Server root certificate of VpnServerConfiguration.
        :param pulumi.Input[str] radius_server_secret: The radius secret property of the VpnServerConfiguration resource for point to site client connection.
        :param pulumi.Input[list] radius_servers: Multiple Radius Server configuration for VpnServerConfiguration.
        :param pulumi.Input[str] resource_group_name: The resource group name of the VpnServerConfiguration.
        :param pulumi.Input[dict] tags: Resource tags.
        :param pulumi.Input[list] vpn_authentication_types: VPN authentication types for the VpnServerConfiguration.
        :param pulumi.Input[list] vpn_client_ipsec_policies: VpnClientIpsecPolicies for VpnServerConfiguration.
        :param pulumi.Input[list] vpn_client_revoked_certificates: VPN client revoked certificate of VpnServerConfiguration.
        :param pulumi.Input[list] vpn_client_root_certificates: VPN client root certificate of VpnServerConfiguration.
        :param pulumi.Input[list] vpn_protocols: VPN protocols for the VpnServerConfiguration.

        The **aad_authentication_parameters** object supports the following:

          * `aad_audience` (`pulumi.Input[str]`) - AAD Vpn authentication parameter AAD audience.
          * `aad_issuer` (`pulumi.Input[str]`) - AAD Vpn authentication parameter AAD issuer.
          * `aad_tenant` (`pulumi.Input[str]`) - AAD Vpn authentication parameter AAD tenant.

        The **radius_client_root_certificates** object supports the following:

          * `name` (`pulumi.Input[str]`) - The certificate name.
          * `thumbprint` (`pulumi.Input[str]`) - The Radius client root certificate thumbprint.

        The **radius_server_root_certificates** object supports the following:

          * `name` (`pulumi.Input[str]`) - The certificate name.
          * `public_cert_data` (`pulumi.Input[str]`) - The certificate public data.

        The **radius_servers** object supports the following:

          * `radius_server_address` (`pulumi.Input[str]`) - The address of this radius server.
          * `radius_server_score` (`pulumi.Input[float]`) - The initial score assigned to this radius server.
          * `radius_server_secret` (`pulumi.Input[str]`) - The secret used for this radius server.

        The **vpn_client_ipsec_policies** object supports the following:

          * `dh_group` (`pulumi.Input[str]`) - The DH Group used in IKE Phase 1 for initial SA.
          * `ike_encryption` (`pulumi.Input[str]`) - The IKE encryption algorithm (IKE phase 2).
          * `ike_integrity` (`pulumi.Input[str]`) - The IKE integrity algorithm (IKE phase 2).
          * `ipsec_encryption` (`pulumi.Input[str]`) - The IPSec encryption algorithm (IKE phase 1).
          * `ipsec_integrity` (`pulumi.Input[str]`) - The IPSec integrity algorithm (IKE phase 1).
          * `pfs_group` (`pulumi.Input[str]`) - The Pfs Group used in IKE Phase 2 for new child SA.
          * `sa_data_size_kilobytes` (`pulumi.Input[float]`) - The IPSec Security Association (also called Quick Mode or Phase 2 SA) payload size in KB for a site to site VPN tunnel.
          * `sa_life_time_seconds` (`pulumi.Input[float]`) - The IPSec Security Association (also called Quick Mode or Phase 2 SA) lifetime in seconds for a site to site VPN tunnel.

        The **vpn_client_revoked_certificates** object supports the following:

          * `name` (`pulumi.Input[str]`) - The certificate name.
          * `thumbprint` (`pulumi.Input[str]`) - The revoked VPN client certificate thumbprint.

        The **vpn_client_root_certificates** object supports the following:

          * `name` (`pulumi.Input[str]`) - The certificate name.
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

            __props__['aad_authentication_parameters'] = aad_authentication_parameters
            __props__['id'] = id
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['radius_client_root_certificates'] = radius_client_root_certificates
            __props__['radius_server_address'] = radius_server_address
            __props__['radius_server_root_certificates'] = radius_server_root_certificates
            __props__['radius_server_secret'] = radius_server_secret
            __props__['radius_servers'] = radius_servers
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['vpn_authentication_types'] = vpn_authentication_types
            __props__['vpn_client_ipsec_policies'] = vpn_client_ipsec_policies
            __props__['vpn_client_revoked_certificates'] = vpn_client_revoked_certificates
            __props__['vpn_client_root_certificates'] = vpn_client_root_certificates
            __props__['vpn_protocols'] = vpn_protocols
            __props__['etag'] = None
            __props__['properties'] = None
            __props__['type'] = None
        super(VpnServerConfiguration, __self__).__init__(
            'azurerm:network/v20200401:VpnServerConfiguration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing VpnServerConfiguration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return VpnServerConfiguration(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
