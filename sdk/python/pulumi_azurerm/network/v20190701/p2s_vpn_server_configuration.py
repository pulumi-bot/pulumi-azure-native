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

__all__ = ['P2sVpnServerConfiguration']


class P2sVpnServerConfiguration(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 p2_s_vpn_server_config_radius_client_root_certificates: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['P2SVpnServerConfigRadiusClientRootCertificateArgs']]]]] = None,
                 p2_s_vpn_server_config_radius_server_root_certificates: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['P2SVpnServerConfigRadiusServerRootCertificateArgs']]]]] = None,
                 p2_s_vpn_server_config_vpn_client_revoked_certificates: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['P2SVpnServerConfigVpnClientRevokedCertificateArgs']]]]] = None,
                 p2_s_vpn_server_config_vpn_client_root_certificates: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['P2SVpnServerConfigVpnClientRootCertificateArgs']]]]] = None,
                 radius_server_address: Optional[pulumi.Input[str]] = None,
                 radius_server_secret: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 virtual_wan_name: Optional[pulumi.Input[str]] = None,
                 vpn_client_ipsec_policies: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['IpsecPolicyArgs']]]]] = None,
                 vpn_protocols: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        P2SVpnServerConfiguration Resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] etag: A unique read-only string that changes whenever the resource is updated.
        :param pulumi.Input[str] id: Resource ID.
        :param pulumi.Input[str] name: The name of the P2SVpnServerConfiguration.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['P2SVpnServerConfigRadiusClientRootCertificateArgs']]]] p2_s_vpn_server_config_radius_client_root_certificates: Radius client root certificate of P2SVpnServerConfiguration.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['P2SVpnServerConfigRadiusServerRootCertificateArgs']]]] p2_s_vpn_server_config_radius_server_root_certificates: Radius Server root certificate of P2SVpnServerConfiguration.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['P2SVpnServerConfigVpnClientRevokedCertificateArgs']]]] p2_s_vpn_server_config_vpn_client_revoked_certificates: VPN client revoked certificate of P2SVpnServerConfiguration.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['P2SVpnServerConfigVpnClientRootCertificateArgs']]]] p2_s_vpn_server_config_vpn_client_root_certificates: VPN client root certificate of P2SVpnServerConfiguration.
        :param pulumi.Input[str] radius_server_address: The radius server address property of the P2SVpnServerConfiguration resource for point to site client connection.
        :param pulumi.Input[str] radius_server_secret: The radius secret property of the P2SVpnServerConfiguration resource for point to site client connection.
        :param pulumi.Input[str] resource_group_name: The resource group name of the VirtualWan.
        :param pulumi.Input[str] virtual_wan_name: The name of the VirtualWan.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['IpsecPolicyArgs']]]] vpn_client_ipsec_policies: VpnClientIpsecPolicies for P2SVpnServerConfiguration.
        :param pulumi.Input[List[pulumi.Input[str]]] vpn_protocols: VPN protocols for the P2SVpnServerConfiguration.
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
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['p2_s_vpn_server_config_radius_client_root_certificates'] = p2_s_vpn_server_config_radius_client_root_certificates
            __props__['p2_s_vpn_server_config_radius_server_root_certificates'] = p2_s_vpn_server_config_radius_server_root_certificates
            __props__['p2_s_vpn_server_config_vpn_client_revoked_certificates'] = p2_s_vpn_server_config_vpn_client_revoked_certificates
            __props__['p2_s_vpn_server_config_vpn_client_root_certificates'] = p2_s_vpn_server_config_vpn_client_root_certificates
            __props__['radius_server_address'] = radius_server_address
            __props__['radius_server_secret'] = radius_server_secret
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if virtual_wan_name is None:
                raise TypeError("Missing required property 'virtual_wan_name'")
            __props__['virtual_wan_name'] = virtual_wan_name
            __props__['vpn_client_ipsec_policies'] = vpn_client_ipsec_policies
            __props__['vpn_protocols'] = vpn_protocols
            __props__['p2_s_vpn_gateways'] = None
            __props__['provisioning_state'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:network/v20180801:P2sVpnServerConfiguration"), pulumi.Alias(type_="azurerm:network/v20181001:P2sVpnServerConfiguration"), pulumi.Alias(type_="azurerm:network/v20181101:P2sVpnServerConfiguration"), pulumi.Alias(type_="azurerm:network/v20181201:P2sVpnServerConfiguration"), pulumi.Alias(type_="azurerm:network/v20190201:P2sVpnServerConfiguration"), pulumi.Alias(type_="azurerm:network/v20190401:P2sVpnServerConfiguration"), pulumi.Alias(type_="azurerm:network/v20190601:P2sVpnServerConfiguration")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(P2sVpnServerConfiguration, __self__).__init__(
            'azurerm:network/v20190701:P2sVpnServerConfiguration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'P2sVpnServerConfiguration':
        """
        Get an existing P2sVpnServerConfiguration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return P2sVpnServerConfiguration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def etag(self) -> str:
        """
        A unique read-only string that changes whenever the resource is updated.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The name of the P2SVpnServerConfiguration that is unique within a VirtualWan in a resource group. This name can be used to access the resource along with Paren VirtualWan resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="p2SVpnGateways")
    def p2_s_vpn_gateways(self) -> List['outputs.SubResourceResponse']:
        """
        List of references to P2SVpnGateways.
        """
        return pulumi.get(self, "p2_s_vpn_gateways")

    @property
    @pulumi.getter(name="p2SVpnServerConfigRadiusClientRootCertificates")
    def p2_s_vpn_server_config_radius_client_root_certificates(self) -> Optional[List['outputs.P2SVpnServerConfigRadiusClientRootCertificateResponse']]:
        """
        Radius client root certificate of P2SVpnServerConfiguration.
        """
        return pulumi.get(self, "p2_s_vpn_server_config_radius_client_root_certificates")

    @property
    @pulumi.getter(name="p2SVpnServerConfigRadiusServerRootCertificates")
    def p2_s_vpn_server_config_radius_server_root_certificates(self) -> Optional[List['outputs.P2SVpnServerConfigRadiusServerRootCertificateResponse']]:
        """
        Radius Server root certificate of P2SVpnServerConfiguration.
        """
        return pulumi.get(self, "p2_s_vpn_server_config_radius_server_root_certificates")

    @property
    @pulumi.getter(name="p2SVpnServerConfigVpnClientRevokedCertificates")
    def p2_s_vpn_server_config_vpn_client_revoked_certificates(self) -> Optional[List['outputs.P2SVpnServerConfigVpnClientRevokedCertificateResponse']]:
        """
        VPN client revoked certificate of P2SVpnServerConfiguration.
        """
        return pulumi.get(self, "p2_s_vpn_server_config_vpn_client_revoked_certificates")

    @property
    @pulumi.getter(name="p2SVpnServerConfigVpnClientRootCertificates")
    def p2_s_vpn_server_config_vpn_client_root_certificates(self) -> Optional[List['outputs.P2SVpnServerConfigVpnClientRootCertificateResponse']]:
        """
        VPN client root certificate of P2SVpnServerConfiguration.
        """
        return pulumi.get(self, "p2_s_vpn_server_config_vpn_client_root_certificates")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        The provisioning state of the P2S VPN server configuration resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="radiusServerAddress")
    def radius_server_address(self) -> Optional[str]:
        """
        The radius server address property of the P2SVpnServerConfiguration resource for point to site client connection.
        """
        return pulumi.get(self, "radius_server_address")

    @property
    @pulumi.getter(name="radiusServerSecret")
    def radius_server_secret(self) -> Optional[str]:
        """
        The radius secret property of the P2SVpnServerConfiguration resource for point to site client connection.
        """
        return pulumi.get(self, "radius_server_secret")

    @property
    @pulumi.getter(name="vpnClientIpsecPolicies")
    def vpn_client_ipsec_policies(self) -> Optional[List['outputs.IpsecPolicyResponse']]:
        """
        VpnClientIpsecPolicies for P2SVpnServerConfiguration.
        """
        return pulumi.get(self, "vpn_client_ipsec_policies")

    @property
    @pulumi.getter(name="vpnProtocols")
    def vpn_protocols(self) -> Optional[List[str]]:
        """
        VPN protocols for the P2SVpnServerConfiguration.
        """
        return pulumi.get(self, "vpn_protocols")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

