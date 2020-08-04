# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class P2sVpnServerConfiguration(pulumi.CustomResource):
    etag: pulumi.Output[str]
    """
    Gets a unique read-only string that changes whenever the resource is updated.
    """
    name: pulumi.Output[str]
    """
    The name of the resource that is unique within a resource group. This name can be used to access the resource.
    """
    properties: pulumi.Output[dict]
    """
    Parameters for P2SVpnServerConfiguration
      * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
      * `name` (`str`) - The name of the P2SVpnServerConfiguration that is unique within a VirtualWan in a resource group. This name can be used to access the resource along with Parent VirtualWan resource name.
      * `p2_s_vpn_gateways` (`list`)
        * `id` (`str`) - Resource ID.

      * `p2_s_vpn_server_config_radius_client_root_certificates` (`list`) - Radius client root certificate of P2SVpnServerConfiguration.
        * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
        * `id` (`str`) - Resource ID.
        * `name` (`str`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
        * `properties` (`dict`) - Properties of the Radius client root certificate.
          * `provisioning_state` (`str`) - The provisioning state of the Radius client root certificate resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
          * `thumbprint` (`str`) - The Radius client root certificate thumbprint.

      * `p2_s_vpn_server_config_radius_server_root_certificates` (`list`) - Radius Server root certificate of P2SVpnServerConfiguration.
        * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
        * `id` (`str`) - Resource ID.
        * `name` (`str`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
        * `properties` (`dict`) - Properties of the P2SVpnServerConfiguration Radius Server root certificate.
          * `provisioning_state` (`str`) - The provisioning state of the P2SVpnServerConfiguration Radius Server root certificate resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
          * `public_cert_data` (`str`) - The certificate public data.

      * `p2_s_vpn_server_config_vpn_client_revoked_certificates` (`list`) - VPN client revoked certificate of P2SVpnServerConfiguration.
        * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
        * `id` (`str`) - Resource ID.
        * `name` (`str`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
        * `properties` (`dict`) - Properties of the vpn client revoked certificate.
          * `provisioning_state` (`str`) - The provisioning state of the VPN client revoked certificate resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
          * `thumbprint` (`str`) - The revoked VPN client certificate thumbprint.

      * `p2_s_vpn_server_config_vpn_client_root_certificates` (`list`) - VPN client root certificate of P2SVpnServerConfiguration.
        * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
        * `id` (`str`) - Resource ID.
        * `name` (`str`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
        * `properties` (`dict`) - Properties of the P2SVpnServerConfiguration VPN client root certificate.
          * `provisioning_state` (`str`) - The provisioning state of the P2SVpnServerConfiguration VPN client root certificate resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
          * `public_cert_data` (`str`) - The certificate public data.

      * `provisioning_state` (`str`) - The provisioning state of the P2SVpnServerConfiguration resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
      * `radius_server_address` (`str`) - The radius server address property of the P2SVpnServerConfiguration resource for point to site client connection.
      * `radius_server_secret` (`str`) - The radius secret property of the P2SVpnServerConfiguration resource for point to site client connection.
      * `vpn_client_ipsec_policies` (`list`) - VpnClientIpsecPolicies for P2SVpnServerConfiguration.
        * `dh_group` (`str`) - The DH Groups used in IKE Phase 1 for initial SA.
        * `ike_encryption` (`str`) - The IKE encryption algorithm (IKE phase 2).
        * `ike_integrity` (`str`) - The IKE integrity algorithm (IKE phase 2).
        * `ipsec_encryption` (`str`) - The IPSec encryption algorithm (IKE phase 1).
        * `ipsec_integrity` (`str`) - The IPSec integrity algorithm (IKE phase 1).
        * `pfs_group` (`str`) - The Pfs Groups used in IKE Phase 2 for new child SA.
        * `sa_data_size_kilobytes` (`float`) - The IPSec Security Association (also called Quick Mode or Phase 2 SA) payload size in KB for a site to site VPN tunnel.
        * `sa_life_time_seconds` (`float`) - The IPSec Security Association (also called Quick Mode or Phase 2 SA) lifetime in seconds for a site to site VPN tunnel.

      * `vpn_protocols` (`list`) - vpnProtocols for the P2SVpnServerConfiguration.
    """
    def __init__(__self__, resource_name, opts=None, id=None, name=None, properties=None, resource_group_name=None, virtual_wan_name=None, __props__=None, __name__=None, __opts__=None):
        """
        P2SVpnServerConfiguration Resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] id: Resource ID.
        :param pulumi.Input[str] name: The name of the P2SVpnServerConfiguration.
        :param pulumi.Input[dict] properties: Parameters for P2SVpnServerConfiguration
        :param pulumi.Input[str] resource_group_name: The resource group name of the VirtualWan.
        :param pulumi.Input[str] virtual_wan_name: The name of the VirtualWan.

        The **properties** object supports the following:

          * `etag` (`pulumi.Input[str]`) - A unique read-only string that changes whenever the resource is updated.
          * `name` (`pulumi.Input[str]`) - The name of the P2SVpnServerConfiguration that is unique within a VirtualWan in a resource group. This name can be used to access the resource along with Parent VirtualWan resource name.
          * `p2_s_vpn_server_config_radius_client_root_certificates` (`pulumi.Input[list]`) - Radius client root certificate of P2SVpnServerConfiguration.
            * `etag` (`pulumi.Input[str]`) - A unique read-only string that changes whenever the resource is updated.
            * `id` (`pulumi.Input[str]`) - Resource ID.
            * `name` (`pulumi.Input[str]`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
            * `properties` (`pulumi.Input[dict]`) - Properties of the Radius client root certificate.
              * `thumbprint` (`pulumi.Input[str]`) - The Radius client root certificate thumbprint.

          * `p2_s_vpn_server_config_radius_server_root_certificates` (`pulumi.Input[list]`) - Radius Server root certificate of P2SVpnServerConfiguration.
            * `etag` (`pulumi.Input[str]`) - A unique read-only string that changes whenever the resource is updated.
            * `id` (`pulumi.Input[str]`) - Resource ID.
            * `name` (`pulumi.Input[str]`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
            * `properties` (`pulumi.Input[dict]`) - Properties of the P2SVpnServerConfiguration Radius Server root certificate.
              * `public_cert_data` (`pulumi.Input[str]`) - The certificate public data.

          * `p2_s_vpn_server_config_vpn_client_revoked_certificates` (`pulumi.Input[list]`) - VPN client revoked certificate of P2SVpnServerConfiguration.
            * `etag` (`pulumi.Input[str]`) - A unique read-only string that changes whenever the resource is updated.
            * `id` (`pulumi.Input[str]`) - Resource ID.
            * `name` (`pulumi.Input[str]`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
            * `properties` (`pulumi.Input[dict]`) - Properties of the vpn client revoked certificate.
              * `thumbprint` (`pulumi.Input[str]`) - The revoked VPN client certificate thumbprint.

          * `p2_s_vpn_server_config_vpn_client_root_certificates` (`pulumi.Input[list]`) - VPN client root certificate of P2SVpnServerConfiguration.
            * `etag` (`pulumi.Input[str]`) - A unique read-only string that changes whenever the resource is updated.
            * `id` (`pulumi.Input[str]`) - Resource ID.
            * `name` (`pulumi.Input[str]`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
            * `properties` (`pulumi.Input[dict]`) - Properties of the P2SVpnServerConfiguration VPN client root certificate.
              * `public_cert_data` (`pulumi.Input[str]`) - The certificate public data.

          * `radius_server_address` (`pulumi.Input[str]`) - The radius server address property of the P2SVpnServerConfiguration resource for point to site client connection.
          * `radius_server_secret` (`pulumi.Input[str]`) - The radius secret property of the P2SVpnServerConfiguration resource for point to site client connection.
          * `vpn_client_ipsec_policies` (`pulumi.Input[list]`) - VpnClientIpsecPolicies for P2SVpnServerConfiguration.
            * `dh_group` (`pulumi.Input[str]`) - The DH Groups used in IKE Phase 1 for initial SA.
            * `ike_encryption` (`pulumi.Input[str]`) - The IKE encryption algorithm (IKE phase 2).
            * `ike_integrity` (`pulumi.Input[str]`) - The IKE integrity algorithm (IKE phase 2).
            * `ipsec_encryption` (`pulumi.Input[str]`) - The IPSec encryption algorithm (IKE phase 1).
            * `ipsec_integrity` (`pulumi.Input[str]`) - The IPSec integrity algorithm (IKE phase 1).
            * `pfs_group` (`pulumi.Input[str]`) - The Pfs Groups used in IKE Phase 2 for new child SA.
            * `sa_data_size_kilobytes` (`pulumi.Input[float]`) - The IPSec Security Association (also called Quick Mode or Phase 2 SA) payload size in KB for a site to site VPN tunnel.
            * `sa_life_time_seconds` (`pulumi.Input[float]`) - The IPSec Security Association (also called Quick Mode or Phase 2 SA) lifetime in seconds for a site to site VPN tunnel.

          * `vpn_protocols` (`pulumi.Input[list]`) - vpnProtocols for the P2SVpnServerConfiguration.
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

            __props__['id'] = id
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if virtual_wan_name is None:
                raise TypeError("Missing required property 'virtual_wan_name'")
            __props__['virtual_wan_name'] = virtual_wan_name
            __props__['etag'] = None
        super(P2sVpnServerConfiguration, __self__).__init__(
            'azurerm:network/v20180801:P2sVpnServerConfiguration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing P2sVpnServerConfiguration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return P2sVpnServerConfiguration(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
