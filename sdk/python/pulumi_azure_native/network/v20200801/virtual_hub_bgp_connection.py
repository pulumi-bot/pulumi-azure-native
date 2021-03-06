# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables

__all__ = ['VirtualHubBgpConnection']


class VirtualHubBgpConnection(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 connection_name: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 peer_asn: Optional[pulumi.Input[float]] = None,
                 peer_ip: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 virtual_hub_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Virtual Appliance Site resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] connection_name: The name of the connection.
        :param pulumi.Input[str] id: Resource ID.
        :param pulumi.Input[str] name: Name of the connection.
        :param pulumi.Input[float] peer_asn: Peer ASN.
        :param pulumi.Input[str] peer_ip: Peer IP.
        :param pulumi.Input[str] resource_group_name: The resource group name of the VirtualHub.
        :param pulumi.Input[str] virtual_hub_name: The name of the VirtualHub.
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

            __props__['connection_name'] = connection_name
            __props__['id'] = id
            __props__['name'] = name
            __props__['peer_asn'] = peer_asn
            __props__['peer_ip'] = peer_ip
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if virtual_hub_name is None and not opts.urn:
                raise TypeError("Missing required property 'virtual_hub_name'")
            __props__['virtual_hub_name'] = virtual_hub_name
            __props__['connection_state'] = None
            __props__['etag'] = None
            __props__['provisioning_state'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:network/v20200801:VirtualHubBgpConnection"), pulumi.Alias(type_="azure-native:network:VirtualHubBgpConnection"), pulumi.Alias(type_="azure-nextgen:network:VirtualHubBgpConnection"), pulumi.Alias(type_="azure-native:network/latest:VirtualHubBgpConnection"), pulumi.Alias(type_="azure-nextgen:network/latest:VirtualHubBgpConnection"), pulumi.Alias(type_="azure-native:network/v20200501:VirtualHubBgpConnection"), pulumi.Alias(type_="azure-nextgen:network/v20200501:VirtualHubBgpConnection"), pulumi.Alias(type_="azure-native:network/v20200601:VirtualHubBgpConnection"), pulumi.Alias(type_="azure-nextgen:network/v20200601:VirtualHubBgpConnection"), pulumi.Alias(type_="azure-native:network/v20200701:VirtualHubBgpConnection"), pulumi.Alias(type_="azure-nextgen:network/v20200701:VirtualHubBgpConnection")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(VirtualHubBgpConnection, __self__).__init__(
            'azure-native:network/v20200801:VirtualHubBgpConnection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'VirtualHubBgpConnection':
        """
        Get an existing VirtualHubBgpConnection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["connection_state"] = None
        __props__["etag"] = None
        __props__["name"] = None
        __props__["peer_asn"] = None
        __props__["peer_ip"] = None
        __props__["provisioning_state"] = None
        __props__["type"] = None
        return VirtualHubBgpConnection(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="connectionState")
    def connection_state(self) -> pulumi.Output[str]:
        """
        The current state of the VirtualHub to Peer.
        """
        return pulumi.get(self, "connection_state")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        A unique read-only string that changes whenever the resource is updated.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[str]]:
        """
        Name of the connection.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="peerAsn")
    def peer_asn(self) -> pulumi.Output[Optional[float]]:
        """
        Peer ASN.
        """
        return pulumi.get(self, "peer_asn")

    @property
    @pulumi.getter(name="peerIp")
    def peer_ip(self) -> pulumi.Output[Optional[str]]:
        """
        Peer IP.
        """
        return pulumi.get(self, "peer_ip")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        The provisioning state of the resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Connection type.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

