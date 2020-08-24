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

__all__ = ['WebApp']


class WebApp(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 client_affinity_enabled: Optional[pulumi.Input[bool]] = None,
                 client_cert_enabled: Optional[pulumi.Input[bool]] = None,
                 client_cert_exclusion_paths: Optional[pulumi.Input[str]] = None,
                 client_cert_mode: Optional[pulumi.Input[str]] = None,
                 cloning_info: Optional[pulumi.Input[pulumi.InputType['CloningInfoArgs']]] = None,
                 container_size: Optional[pulumi.Input[float]] = None,
                 custom_domain_verification_id: Optional[pulumi.Input[str]] = None,
                 daily_memory_time_quota: Optional[pulumi.Input[float]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 host_name_ssl_states: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['HostNameSslStateArgs']]]]] = None,
                 host_names_disabled: Optional[pulumi.Input[bool]] = None,
                 hosting_environment_profile: Optional[pulumi.Input[pulumi.InputType['HostingEnvironmentProfileArgs']]] = None,
                 https_only: Optional[pulumi.Input[bool]] = None,
                 hyper_v: Optional[pulumi.Input[bool]] = None,
                 identity: Optional[pulumi.Input[pulumi.InputType['ManagedServiceIdentityArgs']]] = None,
                 is_xenon: Optional[pulumi.Input[bool]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 redundancy_mode: Optional[pulumi.Input[str]] = None,
                 reserved: Optional[pulumi.Input[bool]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 scm_site_also_stopped: Optional[pulumi.Input[bool]] = None,
                 server_farm_id: Optional[pulumi.Input[str]] = None,
                 site_config: Optional[pulumi.Input[pulumi.InputType['SiteConfigArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        A web app, a mobile app backend, or an API app.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] client_affinity_enabled: <code>true</code> to enable client affinity; <code>false</code> to stop sending session affinity cookies, which route client requests in the same session to the same instance. Default is <code>true</code>.
        :param pulumi.Input[bool] client_cert_enabled: <code>true</code> to enable client certificate authentication (TLS mutual authentication); otherwise, <code>false</code>. Default is <code>false</code>.
        :param pulumi.Input[str] client_cert_exclusion_paths: client certificate authentication comma-separated exclusion paths
        :param pulumi.Input[str] client_cert_mode: This composes with ClientCertEnabled setting.
               - ClientCertEnabled: false means ClientCert is ignored.
               - ClientCertEnabled: true and ClientCertMode: Required means ClientCert is required.
               - ClientCertEnabled: true and ClientCertMode: Optional means ClientCert is optional or accepted.
        :param pulumi.Input[pulumi.InputType['CloningInfoArgs']] cloning_info: If specified during app creation, the app is cloned from a source app.
        :param pulumi.Input[float] container_size: Size of the function container.
        :param pulumi.Input[str] custom_domain_verification_id: Unique identifier that verifies the custom domains assigned to the app. Customer will add this id to a txt record for verification.
        :param pulumi.Input[float] daily_memory_time_quota: Maximum allowed daily memory-time quota (applicable on dynamic apps only).
        :param pulumi.Input[bool] enabled: <code>true</code> if the app is enabled; otherwise, <code>false</code>. Setting this value to false disables the app (takes the app offline).
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['HostNameSslStateArgs']]]] host_name_ssl_states: Hostname SSL states are used to manage the SSL bindings for app's hostnames.
        :param pulumi.Input[bool] host_names_disabled: <code>true</code> to disable the public hostnames of the app; otherwise, <code>false</code>.
                If <code>true</code>, the app is only accessible via API management process.
        :param pulumi.Input[pulumi.InputType['HostingEnvironmentProfileArgs']] hosting_environment_profile: App Service Environment to use for the app.
        :param pulumi.Input[bool] https_only: HttpsOnly: configures a web site to accept only https requests. Issues redirect for
               http requests
        :param pulumi.Input[bool] hyper_v: Hyper-V sandbox.
        :param pulumi.Input[pulumi.InputType['ManagedServiceIdentityArgs']] identity: Managed service identity.
        :param pulumi.Input[bool] is_xenon: Obsolete: Hyper-V sandbox.
        :param pulumi.Input[str] kind: Kind of resource.
        :param pulumi.Input[str] location: Resource Location.
        :param pulumi.Input[str] name: Unique name of the app to create or update. To create or update a deployment slot, use the {slot} parameter.
        :param pulumi.Input[str] redundancy_mode: Site redundancy mode
        :param pulumi.Input[bool] reserved: <code>true</code> if reserved; otherwise, <code>false</code>.
        :param pulumi.Input[str] resource_group_name: Name of the resource group to which the resource belongs.
        :param pulumi.Input[bool] scm_site_also_stopped: <code>true</code> to stop SCM (KUDU) site when the app is stopped; otherwise, <code>false</code>. The default is <code>false</code>.
        :param pulumi.Input[str] server_farm_id: Resource ID of the associated App Service plan, formatted as: "/subscriptions/{subscriptionID}/resourceGroups/{groupName}/providers/Microsoft.Web/serverfarms/{appServicePlanName}".
        :param pulumi.Input[pulumi.InputType['SiteConfigArgs']] site_config: Configuration of the app.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags.
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
            __props__['client_cert_exclusion_paths'] = client_cert_exclusion_paths
            __props__['client_cert_mode'] = client_cert_mode
            __props__['cloning_info'] = cloning_info
            __props__['container_size'] = container_size
            __props__['custom_domain_verification_id'] = custom_domain_verification_id
            __props__['daily_memory_time_quota'] = daily_memory_time_quota
            __props__['enabled'] = enabled
            __props__['host_name_ssl_states'] = host_name_ssl_states
            __props__['host_names_disabled'] = host_names_disabled
            __props__['hosting_environment_profile'] = hosting_environment_profile
            __props__['https_only'] = https_only
            __props__['hyper_v'] = hyper_v
            __props__['identity'] = identity
            __props__['is_xenon'] = is_xenon
            __props__['kind'] = kind
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['redundancy_mode'] = redundancy_mode
            __props__['reserved'] = reserved
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['scm_site_also_stopped'] = scm_site_also_stopped
            __props__['server_farm_id'] = server_farm_id
            __props__['site_config'] = site_config
            __props__['tags'] = tags
            __props__['availability_state'] = None
            __props__['default_host_name'] = None
            __props__['enabled_host_names'] = None
            __props__['host_names'] = None
            __props__['in_progress_operation_id'] = None
            __props__['is_default_container'] = None
            __props__['last_modified_time_utc'] = None
            __props__['max_number_of_workers'] = None
            __props__['outbound_ip_addresses'] = None
            __props__['possible_outbound_ip_addresses'] = None
            __props__['repository_site_name'] = None
            __props__['resource_group'] = None
            __props__['slot_swap_status'] = None
            __props__['state'] = None
            __props__['suspended_till'] = None
            __props__['target_swap_slot'] = None
            __props__['traffic_manager_host_names'] = None
            __props__['type'] = None
            __props__['usage_state'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:web/v20150801:WebApp"), pulumi.Alias(type_="azurerm:web/v20160801:WebApp"), pulumi.Alias(type_="azurerm:web/v20180201:WebApp"), pulumi.Alias(type_="azurerm:web/v20181101:WebApp"), pulumi.Alias(type_="azurerm:web/v20190801:WebApp")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(WebApp, __self__).__init__(
            'azurerm:web/v20200601:WebApp',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'WebApp':
        """
        Get an existing WebApp resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return WebApp(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="availabilityState")
    def availability_state(self) -> str:
        """
        Management information availability state for the app.
        """
        return pulumi.get(self, "availability_state")

    @property
    @pulumi.getter(name="clientAffinityEnabled")
    def client_affinity_enabled(self) -> Optional[bool]:
        """
        <code>true</code> to enable client affinity; <code>false</code> to stop sending session affinity cookies, which route client requests in the same session to the same instance. Default is <code>true</code>.
        """
        return pulumi.get(self, "client_affinity_enabled")

    @property
    @pulumi.getter(name="clientCertEnabled")
    def client_cert_enabled(self) -> Optional[bool]:
        """
        <code>true</code> to enable client certificate authentication (TLS mutual authentication); otherwise, <code>false</code>. Default is <code>false</code>.
        """
        return pulumi.get(self, "client_cert_enabled")

    @property
    @pulumi.getter(name="clientCertExclusionPaths")
    def client_cert_exclusion_paths(self) -> Optional[str]:
        """
        client certificate authentication comma-separated exclusion paths
        """
        return pulumi.get(self, "client_cert_exclusion_paths")

    @property
    @pulumi.getter(name="clientCertMode")
    def client_cert_mode(self) -> Optional[str]:
        """
        This composes with ClientCertEnabled setting.
        - ClientCertEnabled: false means ClientCert is ignored.
        - ClientCertEnabled: true and ClientCertMode: Required means ClientCert is required.
        - ClientCertEnabled: true and ClientCertMode: Optional means ClientCert is optional or accepted.
        """
        return pulumi.get(self, "client_cert_mode")

    @property
    @pulumi.getter(name="cloningInfo")
    def cloning_info(self) -> Optional['outputs.CloningInfoResponse']:
        """
        If specified during app creation, the app is cloned from a source app.
        """
        return pulumi.get(self, "cloning_info")

    @property
    @pulumi.getter(name="containerSize")
    def container_size(self) -> Optional[float]:
        """
        Size of the function container.
        """
        return pulumi.get(self, "container_size")

    @property
    @pulumi.getter(name="customDomainVerificationId")
    def custom_domain_verification_id(self) -> Optional[str]:
        """
        Unique identifier that verifies the custom domains assigned to the app. Customer will add this id to a txt record for verification.
        """
        return pulumi.get(self, "custom_domain_verification_id")

    @property
    @pulumi.getter(name="dailyMemoryTimeQuota")
    def daily_memory_time_quota(self) -> Optional[float]:
        """
        Maximum allowed daily memory-time quota (applicable on dynamic apps only).
        """
        return pulumi.get(self, "daily_memory_time_quota")

    @property
    @pulumi.getter(name="defaultHostName")
    def default_host_name(self) -> str:
        """
        Default hostname of the app. Read-only.
        """
        return pulumi.get(self, "default_host_name")

    @property
    @pulumi.getter
    def enabled(self) -> Optional[bool]:
        """
        <code>true</code> if the app is enabled; otherwise, <code>false</code>. Setting this value to false disables the app (takes the app offline).
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="enabledHostNames")
    def enabled_host_names(self) -> List[str]:
        """
        Enabled hostnames for the app.Hostnames need to be assigned (see HostNames) AND enabled. Otherwise,
        the app is not served on those hostnames.
        """
        return pulumi.get(self, "enabled_host_names")

    @property
    @pulumi.getter(name="hostNameSslStates")
    def host_name_ssl_states(self) -> Optional[List['outputs.HostNameSslStateResponse']]:
        """
        Hostname SSL states are used to manage the SSL bindings for app's hostnames.
        """
        return pulumi.get(self, "host_name_ssl_states")

    @property
    @pulumi.getter(name="hostNames")
    def host_names(self) -> List[str]:
        """
        Hostnames associated with the app.
        """
        return pulumi.get(self, "host_names")

    @property
    @pulumi.getter(name="hostNamesDisabled")
    def host_names_disabled(self) -> Optional[bool]:
        """
        <code>true</code> to disable the public hostnames of the app; otherwise, <code>false</code>.
         If <code>true</code>, the app is only accessible via API management process.
        """
        return pulumi.get(self, "host_names_disabled")

    @property
    @pulumi.getter(name="hostingEnvironmentProfile")
    def hosting_environment_profile(self) -> Optional['outputs.HostingEnvironmentProfileResponse']:
        """
        App Service Environment to use for the app.
        """
        return pulumi.get(self, "hosting_environment_profile")

    @property
    @pulumi.getter(name="httpsOnly")
    def https_only(self) -> Optional[bool]:
        """
        HttpsOnly: configures a web site to accept only https requests. Issues redirect for
        http requests
        """
        return pulumi.get(self, "https_only")

    @property
    @pulumi.getter(name="hyperV")
    def hyper_v(self) -> Optional[bool]:
        """
        Hyper-V sandbox.
        """
        return pulumi.get(self, "hyper_v")

    @property
    @pulumi.getter
    def identity(self) -> Optional['outputs.ManagedServiceIdentityResponse']:
        """
        Managed service identity.
        """
        return pulumi.get(self, "identity")

    @property
    @pulumi.getter(name="inProgressOperationId")
    def in_progress_operation_id(self) -> str:
        """
        Specifies an operation id if this site has a pending operation.
        """
        return pulumi.get(self, "in_progress_operation_id")

    @property
    @pulumi.getter(name="isDefaultContainer")
    def is_default_container(self) -> bool:
        """
        <code>true</code> if the app is a default container; otherwise, <code>false</code>.
        """
        return pulumi.get(self, "is_default_container")

    @property
    @pulumi.getter(name="isXenon")
    def is_xenon(self) -> Optional[bool]:
        """
        Obsolete: Hyper-V sandbox.
        """
        return pulumi.get(self, "is_xenon")

    @property
    @pulumi.getter
    def kind(self) -> Optional[str]:
        """
        Kind of resource.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="lastModifiedTimeUtc")
    def last_modified_time_utc(self) -> str:
        """
        Last time the app was modified, in UTC. Read-only.
        """
        return pulumi.get(self, "last_modified_time_utc")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        Resource Location.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="maxNumberOfWorkers")
    def max_number_of_workers(self) -> float:
        """
        Maximum number of workers.
        This only applies to Functions container.
        """
        return pulumi.get(self, "max_number_of_workers")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource Name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="outboundIpAddresses")
    def outbound_ip_addresses(self) -> str:
        """
        List of IP addresses that the app uses for outbound connections (e.g. database access). Includes VIPs from tenants that site can be hosted with current settings. Read-only.
        """
        return pulumi.get(self, "outbound_ip_addresses")

    @property
    @pulumi.getter(name="possibleOutboundIpAddresses")
    def possible_outbound_ip_addresses(self) -> str:
        """
        List of IP addresses that the app uses for outbound connections (e.g. database access). Includes VIPs from all tenants except dataComponent. Read-only.
        """
        return pulumi.get(self, "possible_outbound_ip_addresses")

    @property
    @pulumi.getter(name="redundancyMode")
    def redundancy_mode(self) -> Optional[str]:
        """
        Site redundancy mode
        """
        return pulumi.get(self, "redundancy_mode")

    @property
    @pulumi.getter(name="repositorySiteName")
    def repository_site_name(self) -> str:
        """
        Name of the repository site.
        """
        return pulumi.get(self, "repository_site_name")

    @property
    @pulumi.getter
    def reserved(self) -> Optional[bool]:
        """
        <code>true</code> if reserved; otherwise, <code>false</code>.
        """
        return pulumi.get(self, "reserved")

    @property
    @pulumi.getter(name="resourceGroup")
    def resource_group(self) -> str:
        """
        Name of the resource group the app belongs to. Read-only.
        """
        return pulumi.get(self, "resource_group")

    @property
    @pulumi.getter(name="scmSiteAlsoStopped")
    def scm_site_also_stopped(self) -> Optional[bool]:
        """
        <code>true</code> to stop SCM (KUDU) site when the app is stopped; otherwise, <code>false</code>. The default is <code>false</code>.
        """
        return pulumi.get(self, "scm_site_also_stopped")

    @property
    @pulumi.getter(name="serverFarmId")
    def server_farm_id(self) -> Optional[str]:
        """
        Resource ID of the associated App Service plan, formatted as: "/subscriptions/{subscriptionID}/resourceGroups/{groupName}/providers/Microsoft.Web/serverfarms/{appServicePlanName}".
        """
        return pulumi.get(self, "server_farm_id")

    @property
    @pulumi.getter(name="siteConfig")
    def site_config(self) -> Optional['outputs.SiteConfigResponse']:
        """
        Configuration of the app.
        """
        return pulumi.get(self, "site_config")

    @property
    @pulumi.getter(name="slotSwapStatus")
    def slot_swap_status(self) -> 'outputs.SlotSwapStatusResponse':
        """
        Status of the last deployment slot swap operation.
        """
        return pulumi.get(self, "slot_swap_status")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        Current state of the app.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="suspendedTill")
    def suspended_till(self) -> str:
        """
        App suspended till in case memory-time quota is exceeded.
        """
        return pulumi.get(self, "suspended_till")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="targetSwapSlot")
    def target_swap_slot(self) -> str:
        """
        Specifies which deployment slot this app will swap into. Read-only.
        """
        return pulumi.get(self, "target_swap_slot")

    @property
    @pulumi.getter(name="trafficManagerHostNames")
    def traffic_manager_host_names(self) -> List[str]:
        """
        Azure Traffic Manager hostnames associated with the app. Read-only.
        """
        return pulumi.get(self, "traffic_manager_host_names")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Resource type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="usageState")
    def usage_state(self) -> str:
        """
        State indicating whether the app has exceeded its quota usage. Read-only.
        """
        return pulumi.get(self, "usage_state")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

