# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class VpnConnection(pulumi.CustomResource):
    etag: pulumi.Output[str]
    """
    A unique read-only string that changes whenever the resource is updated.
    """
    name: pulumi.Output[str]
    """
    The name of the resource that is unique within a resource group. This name can be used to access the resource.
    """
    properties: pulumi.Output[dict]
    """
    Properties of the VPN connection.
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
    """
    def __init__(__self__, resource_name, opts=None, connection_bandwidth=None, connection_status=None, enable_bgp=None, enable_internet_security=None, enable_rate_limiting=None, gateway_name=None, id=None, ipsec_policies=None, name=None, provisioning_state=None, remote_vpn_site=None, resource_group_name=None, routing_weight=None, shared_key=None, use_local_azure_ip_address=None, use_policy_based_traffic_selectors=None, vpn_connection_protocol_type=None, vpn_link_connections=None, __props__=None, __name__=None, __opts__=None):
        """
        VpnConnection Resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] connection_bandwidth: Expected bandwidth in MBPS.
        :param pulumi.Input[str] connection_status: The connection status.
        :param pulumi.Input[bool] enable_bgp: EnableBgp flag.
        :param pulumi.Input[bool] enable_internet_security: Enable internet security.
        :param pulumi.Input[bool] enable_rate_limiting: EnableBgp flag.
        :param pulumi.Input[str] gateway_name: The name of the gateway.
        :param pulumi.Input[str] id: Resource ID.
        :param pulumi.Input[list] ipsec_policies: The IPSec Policies to be considered by this connection.
        :param pulumi.Input[str] name: The name of the connection.
        :param pulumi.Input[str] provisioning_state: The provisioning state of the VPN connection resource.
        :param pulumi.Input[dict] remote_vpn_site: Id of the connected vpn site.
        :param pulumi.Input[str] resource_group_name: The resource group name of the VpnGateway.
        :param pulumi.Input[float] routing_weight: Routing weight for vpn connection.
        :param pulumi.Input[str] shared_key: SharedKey for the vpn connection.
        :param pulumi.Input[bool] use_local_azure_ip_address: Use local azure ip to initiate connection.
        :param pulumi.Input[bool] use_policy_based_traffic_selectors: Enable policy-based traffic selectors.
        :param pulumi.Input[str] vpn_connection_protocol_type: Connection protocol used for this connection.
        :param pulumi.Input[list] vpn_link_connections: List of all vpn site link connections to the gateway.

        The **ipsec_policies** object supports the following:

          * `dh_group` (`pulumi.Input[str]`) - The DH Group used in IKE Phase 1 for initial SA.
          * `ike_encryption` (`pulumi.Input[str]`) - The IKE encryption algorithm (IKE phase 2).
          * `ike_integrity` (`pulumi.Input[str]`) - The IKE integrity algorithm (IKE phase 2).
          * `ipsec_encryption` (`pulumi.Input[str]`) - The IPSec encryption algorithm (IKE phase 1).
          * `ipsec_integrity` (`pulumi.Input[str]`) - The IPSec integrity algorithm (IKE phase 1).
          * `pfs_group` (`pulumi.Input[str]`) - The Pfs Group used in IKE Phase 2 for new child SA.
          * `sa_data_size_kilobytes` (`pulumi.Input[float]`) - The IPSec Security Association (also called Quick Mode or Phase 2 SA) payload size in KB for a site to site VPN tunnel.
          * `sa_life_time_seconds` (`pulumi.Input[float]`) - The IPSec Security Association (also called Quick Mode or Phase 2 SA) lifetime in seconds for a site to site VPN tunnel.

        The **remote_vpn_site** object supports the following:

          * `id` (`pulumi.Input[str]`) - Resource ID.

        The **vpn_link_connections** object supports the following:

          * `connection_bandwidth` (`pulumi.Input[float]`) - Expected bandwidth in MBPS.
          * `connection_status` (`pulumi.Input[str]`) - The connection status.
          * `enable_bgp` (`pulumi.Input[bool]`) - EnableBgp flag.
          * `enable_rate_limiting` (`pulumi.Input[bool]`) - EnableBgp flag.
          * `id` (`pulumi.Input[str]`) - Resource ID.
          * `ipsec_policies` (`pulumi.Input[list]`) - The IPSec Policies to be considered by this connection.
          * `name` (`pulumi.Input[str]`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
          * `provisioning_state` (`pulumi.Input[str]`) - The provisioning state of the VPN site link connection resource.
          * `routing_weight` (`pulumi.Input[float]`) - Routing weight for vpn connection.
          * `shared_key` (`pulumi.Input[str]`) - SharedKey for the vpn connection.
          * `use_local_azure_ip_address` (`pulumi.Input[bool]`) - Use local azure ip to initiate connection.
          * `use_policy_based_traffic_selectors` (`pulumi.Input[bool]`) - Enable policy-based traffic selectors.
          * `vpn_connection_protocol_type` (`pulumi.Input[str]`) - Connection protocol used for this connection.
          * `vpn_site_link` (`pulumi.Input[dict]`) - Id of the connected vpn site link.
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

            __props__['connection_bandwidth'] = connection_bandwidth
            __props__['connection_status'] = connection_status
            __props__['enable_bgp'] = enable_bgp
            __props__['enable_internet_security'] = enable_internet_security
            __props__['enable_rate_limiting'] = enable_rate_limiting
            if gateway_name is None:
                raise TypeError("Missing required property 'gateway_name'")
            __props__['gateway_name'] = gateway_name
            __props__['id'] = id
            __props__['ipsec_policies'] = ipsec_policies
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['provisioning_state'] = provisioning_state
            __props__['remote_vpn_site'] = remote_vpn_site
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['routing_weight'] = routing_weight
            __props__['shared_key'] = shared_key
            __props__['use_local_azure_ip_address'] = use_local_azure_ip_address
            __props__['use_policy_based_traffic_selectors'] = use_policy_based_traffic_selectors
            __props__['vpn_connection_protocol_type'] = vpn_connection_protocol_type
            __props__['vpn_link_connections'] = vpn_link_connections
            __props__['etag'] = None
            __props__['properties'] = None
        super(VpnConnection, __self__).__init__(
            'azurerm:network/v20190701:VpnConnection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing VpnConnection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return VpnConnection(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
