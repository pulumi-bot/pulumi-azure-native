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

__all__ = ['SignalR']


class SignalR(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cors: Optional[pulumi.Input[pulumi.InputType['SignalRCorsSettingsArgs']]] = None,
                 features: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['SignalRFeatureArgs']]]]] = None,
                 host_name_prefix: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_acls: Optional[pulumi.Input[pulumi.InputType['SignalRNetworkACLsArgs']]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 sku: Optional[pulumi.Input[pulumi.InputType['ResourceSkuArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 upstream: Optional[pulumi.Input[pulumi.InputType['ServerlessUpstreamSettingsArgs']]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        A class represent a SignalR service resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['SignalRCorsSettingsArgs']] cors: Cross-Origin Resource Sharing (CORS) settings.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['SignalRFeatureArgs']]]] features: List of SignalR featureFlags. e.g. ServiceMode.
               
               FeatureFlags that are not included in the parameters for the update operation will not be modified.
               And the response will only include featureFlags that are explicitly set. 
               When a featureFlag is not explicitly set, SignalR service will use its globally default value. 
               But keep in mind, the default value doesn't mean "false". It varies in terms of different FeatureFlags.
        :param pulumi.Input[str] host_name_prefix: Prefix for the hostName of the SignalR service. Retained for future use.
               The hostname will be of format: &lt;hostNamePrefix&gt;.service.signalr.net.
        :param pulumi.Input[str] kind: The kind of the service - e.g. "SignalR", or "RawWebSockets" for "Microsoft.SignalRService/SignalR"
        :param pulumi.Input[str] location: The GEO location of the SignalR service. e.g. West US | East US | North Central US | South Central US.
        :param pulumi.Input[str] name: The name of the SignalR resource.
        :param pulumi.Input[pulumi.InputType['SignalRNetworkACLsArgs']] network_acls: Network ACLs
        :param pulumi.Input[str] resource_group_name: The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        :param pulumi.Input[pulumi.InputType['ResourceSkuArgs']] sku: The billing information of the resource.(e.g. Free, Standard)
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Tags of the service which is a list of key value pairs that describe the resource.
        :param pulumi.Input[pulumi.InputType['ServerlessUpstreamSettingsArgs']] upstream: Upstream settings when the Azure SignalR is in server-less mode.
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

            __props__['cors'] = cors
            __props__['features'] = features
            __props__['host_name_prefix'] = host_name_prefix
            __props__['kind'] = kind
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['network_acls'] = network_acls
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['sku'] = sku
            __props__['tags'] = tags
            __props__['upstream'] = upstream
            __props__['external_ip'] = None
            __props__['host_name'] = None
            __props__['private_endpoint_connections'] = None
            __props__['provisioning_state'] = None
            __props__['public_port'] = None
            __props__['server_port'] = None
            __props__['type'] = None
            __props__['version'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:signalrservice/v20181001:SignalR")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(SignalR, __self__).__init__(
            'azurerm:signalrservice/v20200501:SignalR',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'SignalR':
        """
        Get an existing SignalR resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return SignalR(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def cors(self) -> Optional['outputs.SignalRCorsSettingsResponse']:
        """
        Cross-Origin Resource Sharing (CORS) settings.
        """
        return pulumi.get(self, "cors")

    @property
    @pulumi.getter(name="externalIP")
    def external_ip(self) -> str:
        """
        The publicly accessible IP of the SignalR service.
        """
        return pulumi.get(self, "external_ip")

    @property
    @pulumi.getter
    def features(self) -> Optional[List['outputs.SignalRFeatureResponse']]:
        """
        List of SignalR featureFlags. e.g. ServiceMode.
        
        FeatureFlags that are not included in the parameters for the update operation will not be modified.
        And the response will only include featureFlags that are explicitly set. 
        When a featureFlag is not explicitly set, SignalR service will use its globally default value. 
        But keep in mind, the default value doesn't mean "false". It varies in terms of different FeatureFlags.
        """
        return pulumi.get(self, "features")

    @property
    @pulumi.getter(name="hostName")
    def host_name(self) -> str:
        """
        FQDN of the SignalR service instance. Format: xxx.service.signalr.net
        """
        return pulumi.get(self, "host_name")

    @property
    @pulumi.getter(name="hostNamePrefix")
    def host_name_prefix(self) -> Optional[str]:
        """
        Prefix for the hostName of the SignalR service. Retained for future use.
        The hostname will be of format: &lt;hostNamePrefix&gt;.service.signalr.net.
        """
        return pulumi.get(self, "host_name_prefix")

    @property
    @pulumi.getter
    def kind(self) -> Optional[str]:
        """
        The kind of the service - e.g. "SignalR", or "RawWebSockets" for "Microsoft.SignalRService/SignalR"
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def location(self) -> Optional[str]:
        """
        The GEO location of the SignalR service. e.g. West US | East US | North Central US | South Central US.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkACLs")
    def network_acls(self) -> Optional['outputs.SignalRNetworkACLsResponse']:
        """
        Network ACLs
        """
        return pulumi.get(self, "network_acls")

    @property
    @pulumi.getter(name="privateEndpointConnections")
    def private_endpoint_connections(self) -> List['outputs.PrivateEndpointConnectionResponse']:
        """
        Private endpoint connections to the SignalR resource.
        """
        return pulumi.get(self, "private_endpoint_connections")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        Provisioning state of the resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="publicPort")
    def public_port(self) -> float:
        """
        The publicly accessible port of the SignalR service which is designed for browser/client side usage.
        """
        return pulumi.get(self, "public_port")

    @property
    @pulumi.getter(name="serverPort")
    def server_port(self) -> float:
        """
        The publicly accessible port of the SignalR service which is designed for customer server side usage.
        """
        return pulumi.get(self, "server_port")

    @property
    @pulumi.getter
    def sku(self) -> Optional['outputs.ResourceSkuResponse']:
        """
        The billing information of the resource.(e.g. Free, Standard)
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        Tags of the service which is a list of key value pairs that describe the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the resource - e.g. "Microsoft.SignalRService/SignalR"
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def upstream(self) -> Optional['outputs.ServerlessUpstreamSettingsResponse']:
        """
        Upstream settings when the Azure SignalR is in server-less mode.
        """
        return pulumi.get(self, "upstream")

    @property
    @pulumi.getter
    def version(self) -> str:
        """
        Version of the SignalR resource. Probably you need the same or higher version of client SDKs.
        """
        return pulumi.get(self, "version")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

