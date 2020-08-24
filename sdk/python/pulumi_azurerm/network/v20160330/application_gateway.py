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

__all__ = ['ApplicationGateway']


class ApplicationGateway(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend_address_pools: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayBackendAddressPoolArgs']]]]] = None,
                 backend_http_settings_collection: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayBackendHttpSettingsArgs']]]]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 frontend_ip_configurations: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayFrontendIPConfigurationArgs']]]]] = None,
                 frontend_ports: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayFrontendPortArgs']]]]] = None,
                 gateway_ip_configurations: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayIPConfigurationArgs']]]]] = None,
                 http_listeners: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayHttpListenerArgs']]]]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 probes: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayProbeArgs']]]]] = None,
                 provisioning_state: Optional[pulumi.Input[str]] = None,
                 request_routing_rules: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayRequestRoutingRuleArgs']]]]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 resource_guid: Optional[pulumi.Input[str]] = None,
                 sku: Optional[pulumi.Input[pulumi.InputType['ApplicationGatewaySkuArgs']]] = None,
                 ssl_certificates: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewaySslCertificateArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 url_path_maps: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayUrlPathMapArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        ApplicationGateways resource

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayBackendAddressPoolArgs']]]] backend_address_pools: Gets or sets backend address pool of application gateway resource
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayBackendHttpSettingsArgs']]]] backend_http_settings_collection: Gets or sets backend http settings of application gateway resource
        :param pulumi.Input[str] etag: Gets a unique read-only string that changes whenever the resource is updated
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayFrontendIPConfigurationArgs']]]] frontend_ip_configurations: Gets or sets frontend IP addresses of application gateway resource
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayFrontendPortArgs']]]] frontend_ports: Gets or sets frontend ports of application gateway resource
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayIPConfigurationArgs']]]] gateway_ip_configurations: Gets or sets subnets of application gateway resource
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayHttpListenerArgs']]]] http_listeners: Gets or sets HTTP listeners of application gateway resource
        :param pulumi.Input[str] id: Resource Id
        :param pulumi.Input[str] location: Resource location
        :param pulumi.Input[str] name: The name of the ApplicationGateway.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayProbeArgs']]]] probes: Gets or sets probes of application gateway resource
        :param pulumi.Input[str] provisioning_state: Gets or sets Provisioning state of the ApplicationGateway resource Updating/Deleting/Failed
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayRequestRoutingRuleArgs']]]] request_routing_rules: Gets or sets request routing rules of application gateway resource
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[str] resource_guid: Gets or sets resource GUID property of the ApplicationGateway resource
        :param pulumi.Input[pulumi.InputType['ApplicationGatewaySkuArgs']] sku: Gets or sets sku of application gateway resource
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewaySslCertificateArgs']]]] ssl_certificates: Gets or sets ssl certificates of application gateway resource
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationGatewayUrlPathMapArgs']]]] url_path_maps: Gets or sets URL path map of application gateway resource
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

            __props__['backend_address_pools'] = backend_address_pools
            __props__['backend_http_settings_collection'] = backend_http_settings_collection
            __props__['etag'] = etag
            __props__['frontend_ip_configurations'] = frontend_ip_configurations
            __props__['frontend_ports'] = frontend_ports
            __props__['gateway_ip_configurations'] = gateway_ip_configurations
            __props__['http_listeners'] = http_listeners
            __props__['id'] = id
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['probes'] = probes
            __props__['provisioning_state'] = provisioning_state
            __props__['request_routing_rules'] = request_routing_rules
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['resource_guid'] = resource_guid
            __props__['sku'] = sku
            __props__['ssl_certificates'] = ssl_certificates
            __props__['tags'] = tags
            __props__['url_path_maps'] = url_path_maps
            __props__['operational_state'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:network/v20150615:ApplicationGateway"), pulumi.Alias(type_="azurerm:network/v20160601:ApplicationGateway"), pulumi.Alias(type_="azurerm:network/v20160901:ApplicationGateway"), pulumi.Alias(type_="azurerm:network/v20161201:ApplicationGateway"), pulumi.Alias(type_="azurerm:network/v20170301:ApplicationGateway"), pulumi.Alias(type_="azurerm:network/v20170601:ApplicationGateway"), pulumi.Alias(type_="azurerm:network/v20170801:ApplicationGateway"), pulumi.Alias(type_="azurerm:network/v20170901:ApplicationGateway"), pulumi.Alias(type_="azurerm:network/v20171001:ApplicationGateway"), pulumi.Alias(type_="azurerm:network/v20171101:ApplicationGateway"), pulumi.Alias(type_="azurerm:network/v20180101:ApplicationGateway"), pulumi.Alias(type_="azurerm:network/v20180201:ApplicationGateway"), pulumi.Alias(type_="azurerm:network/v20180401:ApplicationGateway"), pulumi.Alias(type_="azurerm:network/v20180601:ApplicationGateway"), pulumi.Alias(type_="azurerm:network/v20180701:ApplicationGateway"), pulumi.Alias(type_="azurerm:network/v20180801:ApplicationGateway"), pulumi.Alias(type_="azurerm:network/v20181001:ApplicationGateway"), pulumi.Alias(type_="azurerm:network/v20181101:ApplicationGateway"), pulumi.Alias(type_="azurerm:network/v20181201:ApplicationGateway"), pulumi.Alias(type_="azurerm:network/v20190201:ApplicationGateway"), pulumi.Alias(type_="azurerm:network/v20190401:ApplicationGateway"), pulumi.Alias(type_="azurerm:network/v20190601:ApplicationGateway"), pulumi.Alias(type_="azurerm:network/v20190701:ApplicationGateway"), pulumi.Alias(type_="azurerm:network/v20190801:ApplicationGateway"), pulumi.Alias(type_="azurerm:network/v20190901:ApplicationGateway"), pulumi.Alias(type_="azurerm:network/v20191101:ApplicationGateway"), pulumi.Alias(type_="azurerm:network/v20191201:ApplicationGateway"), pulumi.Alias(type_="azurerm:network/v20200301:ApplicationGateway"), pulumi.Alias(type_="azurerm:network/v20200401:ApplicationGateway"), pulumi.Alias(type_="azurerm:network/v20200501:ApplicationGateway"), pulumi.Alias(type_="azurerm:network/v20200601:ApplicationGateway")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(ApplicationGateway, __self__).__init__(
            'azurerm:network/v20160330:ApplicationGateway',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ApplicationGateway':
        """
        Get an existing ApplicationGateway resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ApplicationGateway(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="backendAddressPools")
    def backend_address_pools(self) -> Optional[List['outputs.ApplicationGatewayBackendAddressPoolResponse']]:
        """
        Gets or sets backend address pool of application gateway resource
        """
        return pulumi.get(self, "backend_address_pools")

    @property
    @pulumi.getter(name="backendHttpSettingsCollection")
    def backend_http_settings_collection(self) -> Optional[List['outputs.ApplicationGatewayBackendHttpSettingsResponse']]:
        """
        Gets or sets backend http settings of application gateway resource
        """
        return pulumi.get(self, "backend_http_settings_collection")

    @property
    @pulumi.getter
    def etag(self) -> Optional[str]:
        """
        Gets a unique read-only string that changes whenever the resource is updated
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="frontendIPConfigurations")
    def frontend_ip_configurations(self) -> Optional[List['outputs.ApplicationGatewayFrontendIPConfigurationResponse']]:
        """
        Gets or sets frontend IP addresses of application gateway resource
        """
        return pulumi.get(self, "frontend_ip_configurations")

    @property
    @pulumi.getter(name="frontendPorts")
    def frontend_ports(self) -> Optional[List['outputs.ApplicationGatewayFrontendPortResponse']]:
        """
        Gets or sets frontend ports of application gateway resource
        """
        return pulumi.get(self, "frontend_ports")

    @property
    @pulumi.getter(name="gatewayIPConfigurations")
    def gateway_ip_configurations(self) -> Optional[List['outputs.ApplicationGatewayIPConfigurationResponse']]:
        """
        Gets or sets subnets of application gateway resource
        """
        return pulumi.get(self, "gateway_ip_configurations")

    @property
    @pulumi.getter(name="httpListeners")
    def http_listeners(self) -> Optional[List['outputs.ApplicationGatewayHttpListenerResponse']]:
        """
        Gets or sets HTTP listeners of application gateway resource
        """
        return pulumi.get(self, "http_listeners")

    @property
    @pulumi.getter
    def location(self) -> Optional[str]:
        """
        Resource location
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="operationalState")
    def operational_state(self) -> str:
        """
        Gets operational state of application gateway resource
        """
        return pulumi.get(self, "operational_state")

    @property
    @pulumi.getter
    def probes(self) -> Optional[List['outputs.ApplicationGatewayProbeResponse']]:
        """
        Gets or sets probes of application gateway resource
        """
        return pulumi.get(self, "probes")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> Optional[str]:
        """
        Gets or sets Provisioning state of the ApplicationGateway resource Updating/Deleting/Failed
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="requestRoutingRules")
    def request_routing_rules(self) -> Optional[List['outputs.ApplicationGatewayRequestRoutingRuleResponse']]:
        """
        Gets or sets request routing rules of application gateway resource
        """
        return pulumi.get(self, "request_routing_rules")

    @property
    @pulumi.getter(name="resourceGuid")
    def resource_guid(self) -> Optional[str]:
        """
        Gets or sets resource GUID property of the ApplicationGateway resource
        """
        return pulumi.get(self, "resource_guid")

    @property
    @pulumi.getter
    def sku(self) -> Optional['outputs.ApplicationGatewaySkuResponse']:
        """
        Gets or sets sku of application gateway resource
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter(name="sslCertificates")
    def ssl_certificates(self) -> Optional[List['outputs.ApplicationGatewaySslCertificateResponse']]:
        """
        Gets or sets ssl certificates of application gateway resource
        """
        return pulumi.get(self, "ssl_certificates")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        Resource tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Resource type
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="urlPathMaps")
    def url_path_maps(self) -> Optional[List['outputs.ApplicationGatewayUrlPathMapResponse']]:
        """
        Gets or sets URL path map of application gateway resource
        """
        return pulumi.get(self, "url_path_maps")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

