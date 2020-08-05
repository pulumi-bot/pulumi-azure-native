# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class VpnSite(pulumi.CustomResource):
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
    Properties of the VPN site.
      * `address_space` (`dict`) - The AddressSpace that contains an array of IP address ranges.
        * `address_prefixes` (`list`) - A list of address blocks reserved for this virtual network in CIDR notation.

      * `bgp_properties` (`dict`) - The set of bgp properties.
        * `asn` (`float`) - The BGP speaker's ASN.
        * `bgp_peering_address` (`str`) - The BGP peering address and BGP identifier of this BGP speaker.
        * `peer_weight` (`float`) - The weight added to routes learned from this BGP speaker.

      * `device_properties` (`dict`) - The device properties
        * `device_model` (`str`) - Model of the device.
        * `device_vendor` (`str`) - Name of the device Vendor.
        * `link_speed_in_mbps` (`float`) - Link speed.

      * `ip_address` (`str`) - The ip-address for the vpn-site.
      * `is_security_site` (`bool`) - IsSecuritySite flag
      * `provisioning_state` (`str`) - The provisioning state of the resource.
      * `site_key` (`str`) - The key for vpn-site that can be used for connections.
      * `virtual_wan` (`dict`) - The VirtualWAN to which the vpnSite belongs
        * `id` (`str`) - Resource ID.
    """
    tags: pulumi.Output[dict]
    """
    Resource tags.
    """
    type: pulumi.Output[str]
    """
    Resource type.
    """
    def __init__(__self__, resource_name, opts=None, address_space=None, bgp_properties=None, device_properties=None, id=None, ip_address=None, is_security_site=None, location=None, name=None, provisioning_state=None, resource_group_name=None, site_key=None, tags=None, virtual_wan=None, __props__=None, __name__=None, __opts__=None):
        """
        VpnSite Resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] address_space: The AddressSpace that contains an array of IP address ranges.
        :param pulumi.Input[dict] bgp_properties: The set of bgp properties.
        :param pulumi.Input[dict] device_properties: The device properties
        :param pulumi.Input[str] id: Resource ID.
        :param pulumi.Input[str] ip_address: The ip-address for the vpn-site.
        :param pulumi.Input[bool] is_security_site: IsSecuritySite flag
        :param pulumi.Input[str] location: Resource location.
        :param pulumi.Input[str] name: The name of the VpnSite being created or updated.
        :param pulumi.Input[str] provisioning_state: The provisioning state of the resource.
        :param pulumi.Input[str] resource_group_name: The resource group name of the VpnSite.
        :param pulumi.Input[str] site_key: The key for vpn-site that can be used for connections.
        :param pulumi.Input[dict] tags: Resource tags.
        :param pulumi.Input[dict] virtual_wan: The VirtualWAN to which the vpnSite belongs

        The **address_space** object supports the following:

          * `address_prefixes` (`pulumi.Input[list]`) - A list of address blocks reserved for this virtual network in CIDR notation.

        The **bgp_properties** object supports the following:

          * `asn` (`pulumi.Input[float]`) - The BGP speaker's ASN.
          * `bgp_peering_address` (`pulumi.Input[str]`) - The BGP peering address and BGP identifier of this BGP speaker.
          * `peer_weight` (`pulumi.Input[float]`) - The weight added to routes learned from this BGP speaker.

        The **device_properties** object supports the following:

          * `device_model` (`pulumi.Input[str]`) - Model of the device.
          * `device_vendor` (`pulumi.Input[str]`) - Name of the device Vendor.
          * `link_speed_in_mbps` (`pulumi.Input[float]`) - Link speed.

        The **virtual_wan** object supports the following:

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

            __props__['address_space'] = address_space
            __props__['bgp_properties'] = bgp_properties
            __props__['device_properties'] = device_properties
            __props__['id'] = id
            __props__['ip_address'] = ip_address
            __props__['is_security_site'] = is_security_site
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['provisioning_state'] = provisioning_state
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['site_key'] = site_key
            __props__['tags'] = tags
            __props__['virtual_wan'] = virtual_wan
            __props__['etag'] = None
            __props__['properties'] = None
            __props__['type'] = None
        super(VpnSite, __self__).__init__(
            'azurerm:network/v20190201:VpnSite',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing VpnSite resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return VpnSite(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
