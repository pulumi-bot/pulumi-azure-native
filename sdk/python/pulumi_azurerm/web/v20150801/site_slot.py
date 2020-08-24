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

__all__ = ['SiteSlot']


class SiteSlot(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 client_affinity_enabled: Optional[pulumi.Input[bool]] = None,
                 client_cert_enabled: Optional[pulumi.Input[bool]] = None,
                 cloning_info: Optional[pulumi.Input[pulumi.InputType['CloningInfoArgs']]] = None,
                 container_size: Optional[pulumi.Input[float]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 force_dns_registration: Optional[pulumi.Input[str]] = None,
                 gateway_site_name: Optional[pulumi.Input[str]] = None,
                 host_name_ssl_states: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['HostNameSslStateArgs']]]]] = None,
                 host_names_disabled: Optional[pulumi.Input[bool]] = None,
                 hosting_environment_profile: Optional[pulumi.Input[pulumi.InputType['HostingEnvironmentProfileArgs']]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 max_number_of_workers: Optional[pulumi.Input[float]] = None,
                 micro_service: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 scm_site_also_stopped: Optional[pulumi.Input[bool]] = None,
                 server_farm_id: Optional[pulumi.Input[str]] = None,
                 site_config: Optional[pulumi.Input[pulumi.InputType['SiteConfigArgs']]] = None,
                 skip_custom_domain_verification: Optional[pulumi.Input[str]] = None,
                 skip_dns_registration: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 ttl_in_seconds: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Represents a web app

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] client_affinity_enabled: Specifies if the client affinity is enabled when load balancing http request for multiple instances of the web app
        :param pulumi.Input[bool] client_cert_enabled: Specifies if the client certificate is enabled for the web app
        :param pulumi.Input[pulumi.InputType['CloningInfoArgs']] cloning_info: This is only valid for web app creation. If specified, web app is cloned from 
                           a source web app
        :param pulumi.Input[float] container_size: Size of a function container
        :param pulumi.Input[bool] enabled: True if the site is enabled; otherwise, false. Setting this  value to false disables the site (takes the site off line).
        :param pulumi.Input[str] force_dns_registration: If true, web app hostname is force registered with DNS
        :param pulumi.Input[str] gateway_site_name: Name of gateway app associated with web app
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['HostNameSslStateArgs']]]] host_name_ssl_states: Hostname SSL states are  used to manage the SSL bindings for site's hostnames.
        :param pulumi.Input[bool] host_names_disabled: Specifies if the public hostnames are disabled the web app.
                           If set to true the app is only accessible via API Management process
        :param pulumi.Input[pulumi.InputType['HostingEnvironmentProfileArgs']] hosting_environment_profile: Specification for the hosting environment (App Service Environment) to use for the web app
        :param pulumi.Input[str] id: Resource Id
        :param pulumi.Input[str] kind: Kind of resource
        :param pulumi.Input[str] location: Resource Location
        :param pulumi.Input[float] max_number_of_workers: Maximum number of workers
                           This only applies to function container
        :param pulumi.Input[str] name: Name of web app slot. If not specified then will default to production slot.
        :param pulumi.Input[str] resource_group_name: Name of the resource group
        :param pulumi.Input[bool] scm_site_also_stopped: If set indicates whether to stop SCM (KUDU) site when the web app is stopped. Default is false.
        :param pulumi.Input[pulumi.InputType['SiteConfigArgs']] site_config: Configuration of web app
        :param pulumi.Input[str] skip_custom_domain_verification: If true, custom (non *.azurewebsites.net) domains associated with web app are not verified.
        :param pulumi.Input[str] skip_dns_registration: If true web app hostname is not registered with DNS on creation. This parameter is
                           only used for app creation
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags
        :param pulumi.Input[str] ttl_in_seconds: Time to live in seconds for web app's default domain name
        :param pulumi.Input[str] type: Resource type
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

            __props__['client_affinity_enabled'] = client_affinity_enabled
            __props__['client_cert_enabled'] = client_cert_enabled
            __props__['cloning_info'] = cloning_info
            __props__['container_size'] = container_size
            __props__['enabled'] = enabled
            __props__['force_dns_registration'] = force_dns_registration
            __props__['gateway_site_name'] = gateway_site_name
            __props__['host_name_ssl_states'] = host_name_ssl_states
            __props__['host_names_disabled'] = host_names_disabled
            __props__['hosting_environment_profile'] = hosting_environment_profile
            __props__['id'] = id
            __props__['kind'] = kind
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            __props__['max_number_of_workers'] = max_number_of_workers
            __props__['micro_service'] = micro_service
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['scm_site_also_stopped'] = scm_site_also_stopped
            __props__['server_farm_id'] = server_farm_id
            __props__['site_config'] = site_config
            __props__['skip_custom_domain_verification'] = skip_custom_domain_verification
            __props__['skip_dns_registration'] = skip_dns_registration
            __props__['tags'] = tags
            __props__['ttl_in_seconds'] = ttl_in_seconds
            __props__['type'] = type
            __props__['availability_state'] = None
            __props__['default_host_name'] = None
            __props__['enabled_host_names'] = None
            __props__['host_names'] = None
            __props__['is_default_container'] = None
            __props__['last_modified_time_utc'] = None
            __props__['outbound_ip_addresses'] = None
            __props__['premium_app_deployed'] = None
            __props__['repository_site_name'] = None
            __props__['resource_group'] = None
            __props__['state'] = None
            __props__['target_swap_slot'] = None
            __props__['traffic_manager_host_names'] = None
            __props__['usage_state'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:web/v20160801:SiteSlot"), pulumi.Alias(type_="azurerm:web/v20180201:SiteSlot"), pulumi.Alias(type_="azurerm:web/v20181101:SiteSlot"), pulumi.Alias(type_="azurerm:web/v20190801:SiteSlot"), pulumi.Alias(type_="azurerm:web/v20200601:SiteSlot")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(SiteSlot, __self__).__init__(
            'azurerm:web/v20150801:SiteSlot',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'SiteSlot':
        """
        Get an existing SiteSlot resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return SiteSlot(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="availabilityState")
    def availability_state(self) -> str:
        """
        Management information availability state for the web app. Possible values are Normal or Limited. 
                    Normal means that the site is running correctly and that management information for the site is available. 
                    Limited means that only partial management information for the site is available and that detailed site information is unavailable.
        """
        return pulumi.get(self, "availability_state")

    @property
    @pulumi.getter(name="clientAffinityEnabled")
    def client_affinity_enabled(self) -> Optional[bool]:
        """
        Specifies if the client affinity is enabled when load balancing http request for multiple instances of the web app
        """
        return pulumi.get(self, "client_affinity_enabled")

    @property
    @pulumi.getter(name="clientCertEnabled")
    def client_cert_enabled(self) -> Optional[bool]:
        """
        Specifies if the client certificate is enabled for the web app
        """
        return pulumi.get(self, "client_cert_enabled")

    @property
    @pulumi.getter(name="cloningInfo")
    def cloning_info(self) -> Optional['outputs.CloningInfoResponse']:
        """
        This is only valid for web app creation. If specified, web app is cloned from 
                    a source web app
        """
        return pulumi.get(self, "cloning_info")

    @property
    @pulumi.getter(name="containerSize")
    def container_size(self) -> Optional[float]:
        """
        Size of a function container
        """
        return pulumi.get(self, "container_size")

    @property
    @pulumi.getter(name="defaultHostName")
    def default_host_name(self) -> str:
        """
        Default hostname of the web app
        """
        return pulumi.get(self, "default_host_name")

    @property
    @pulumi.getter
    def enabled(self) -> Optional[bool]:
        """
        True if the site is enabled; otherwise, false. Setting this  value to false disables the site (takes the site off line).
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="enabledHostNames")
    def enabled_host_names(self) -> List[str]:
        """
        Hostnames for the web app that are enabled. Hostnames need to be assigned and enabled. If some hostnames are assigned but not enabled
                    the app is not served on those hostnames
        """
        return pulumi.get(self, "enabled_host_names")

    @property
    @pulumi.getter(name="gatewaySiteName")
    def gateway_site_name(self) -> Optional[str]:
        """
        Name of gateway app associated with web app
        """
        return pulumi.get(self, "gateway_site_name")

    @property
    @pulumi.getter(name="hostNameSslStates")
    def host_name_ssl_states(self) -> Optional[List['outputs.HostNameSslStateResponse']]:
        """
        Hostname SSL states are  used to manage the SSL bindings for site's hostnames.
        """
        return pulumi.get(self, "host_name_ssl_states")

    @property
    @pulumi.getter(name="hostNames")
    def host_names(self) -> List[str]:
        """
        Hostnames associated with web app
        """
        return pulumi.get(self, "host_names")

    @property
    @pulumi.getter(name="hostNamesDisabled")
    def host_names_disabled(self) -> Optional[bool]:
        """
        Specifies if the public hostnames are disabled the web app.
                    If set to true the app is only accessible via API Management process
        """
        return pulumi.get(self, "host_names_disabled")

    @property
    @pulumi.getter(name="hostingEnvironmentProfile")
    def hosting_environment_profile(self) -> Optional['outputs.HostingEnvironmentProfileResponse']:
        """
        Specification for the hosting environment (App Service Environment) to use for the web app
        """
        return pulumi.get(self, "hosting_environment_profile")

    @property
    @pulumi.getter(name="isDefaultContainer")
    def is_default_container(self) -> bool:
        """
        Site is a default container
        """
        return pulumi.get(self, "is_default_container")

    @property
    @pulumi.getter
    def kind(self) -> Optional[str]:
        """
        Kind of resource
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="lastModifiedTimeUtc")
    def last_modified_time_utc(self) -> str:
        """
        Last time web app was modified in UTC
        """
        return pulumi.get(self, "last_modified_time_utc")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        Resource Location
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="maxNumberOfWorkers")
    def max_number_of_workers(self) -> Optional[float]:
        """
        Maximum number of workers
                    This only applies to function container
        """
        return pulumi.get(self, "max_number_of_workers")

    @property
    @pulumi.getter(name="microService")
    def micro_service(self) -> Optional[str]:
        return pulumi.get(self, "micro_service")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Resource Name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="outboundIpAddresses")
    def outbound_ip_addresses(self) -> str:
        """
        List of comma separated IP addresses that this web app uses for outbound connections. Those can be used when configuring firewall rules for databases accessed by this web app.
        """
        return pulumi.get(self, "outbound_ip_addresses")

    @property
    @pulumi.getter(name="premiumAppDeployed")
    def premium_app_deployed(self) -> bool:
        """
        If set indicates whether web app is deployed as a premium app
        """
        return pulumi.get(self, "premium_app_deployed")

    @property
    @pulumi.getter(name="repositorySiteName")
    def repository_site_name(self) -> str:
        """
        Name of repository site
        """
        return pulumi.get(self, "repository_site_name")

    @property
    @pulumi.getter(name="resourceGroup")
    def resource_group(self) -> str:
        """
        Resource group web app belongs to
        """
        return pulumi.get(self, "resource_group")

    @property
    @pulumi.getter(name="scmSiteAlsoStopped")
    def scm_site_also_stopped(self) -> Optional[bool]:
        """
        If set indicates whether to stop SCM (KUDU) site when the web app is stopped. Default is false.
        """
        return pulumi.get(self, "scm_site_also_stopped")

    @property
    @pulumi.getter(name="serverFarmId")
    def server_farm_id(self) -> Optional[str]:
        return pulumi.get(self, "server_farm_id")

    @property
    @pulumi.getter(name="siteConfig")
    def site_config(self) -> Optional['outputs.SiteConfigResponse']:
        """
        Configuration of web app
        """
        return pulumi.get(self, "site_config")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        State of the web app
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        Resource tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="targetSwapSlot")
    def target_swap_slot(self) -> str:
        """
        Read-only property that specifies which slot this app will swap into
        """
        return pulumi.get(self, "target_swap_slot")

    @property
    @pulumi.getter(name="trafficManagerHostNames")
    def traffic_manager_host_names(self) -> List[str]:
        """
        Read-only list of Azure Traffic manager hostnames associated with web app
        """
        return pulumi.get(self, "traffic_manager_host_names")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        Resource type
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="usageState")
    def usage_state(self) -> str:
        """
        State indicating whether web app has exceeded its quota usage
        """
        return pulumi.get(self, "usage_state")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

