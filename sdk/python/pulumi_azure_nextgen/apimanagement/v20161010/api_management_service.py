# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['ApiManagementService']


class ApiManagementService(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 additional_locations: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AdditionalRegionArgs']]]]] = None,
                 addresser_email: Optional[pulumi.Input[str]] = None,
                 custom_properties: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 hostname_configurations: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['HostnameConfigurationArgs']]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 publisher_email: Optional[pulumi.Input[str]] = None,
                 publisher_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 sku: Optional[pulumi.Input[pulumi.InputType['ApiManagementServiceSkuPropertiesArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vpn_type: Optional[pulumi.Input[str]] = None,
                 vpnconfiguration: Optional[pulumi.Input[pulumi.InputType['VirtualNetworkConfigurationArgs']]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        A single API Management service resource in List or Get response.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AdditionalRegionArgs']]]] additional_locations: Additional datacenter locations of the API Management service.
        :param pulumi.Input[str] addresser_email: Addresser email.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] custom_properties: Custom properties of the API Management service, like disabling TLS 1.0.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['HostnameConfigurationArgs']]]] hostname_configurations: Custom hostname configuration of the API Management service.
        :param pulumi.Input[str] location: Resource location.
        :param pulumi.Input[str] name: Resource name.
        :param pulumi.Input[str] publisher_email: Publisher email.
        :param pulumi.Input[str] publisher_name: Publisher name.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[str] service_name: The name of the API Management service.
        :param pulumi.Input[pulumi.InputType['ApiManagementServiceSkuPropertiesArgs']] sku: SKU properties of the API Management service.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags.
        :param pulumi.Input[str] vpn_type: The type of VPN in which API Management service needs to be configured in. None (Default Value) means the API Management service is not part of any Virtual Network, External means the API Management deployment is set up inside a Virtual Network having an Internet Facing Endpoint, and Internal means that API Management deployment is setup inside a Virtual Network having an Intranet Facing Endpoint only.
        :param pulumi.Input[pulumi.InputType['VirtualNetworkConfigurationArgs']] vpnconfiguration: Virtual network configuration of the API Management service.
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

            __props__['additional_locations'] = additional_locations
            __props__['addresser_email'] = addresser_email
            __props__['custom_properties'] = custom_properties
            __props__['hostname_configurations'] = hostname_configurations
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            __props__['name'] = name
            if publisher_email is None:
                raise TypeError("Missing required property 'publisher_email'")
            __props__['publisher_email'] = publisher_email
            if publisher_name is None:
                raise TypeError("Missing required property 'publisher_name'")
            __props__['publisher_name'] = publisher_name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if service_name is None:
                raise TypeError("Missing required property 'service_name'")
            __props__['service_name'] = service_name
            if sku is None:
                raise TypeError("Missing required property 'sku'")
            __props__['sku'] = sku
            __props__['tags'] = tags
            __props__['vpn_type'] = vpn_type
            __props__['vpnconfiguration'] = vpnconfiguration
            __props__['created_at_utc'] = None
            __props__['etag'] = None
            __props__['management_api_url'] = None
            __props__['portal_url'] = None
            __props__['provisioning_state'] = None
            __props__['runtime_url'] = None
            __props__['scm_url'] = None
            __props__['static_ips'] = None
            __props__['target_provisioning_state'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:apimanagement/latest:ApiManagementService"), pulumi.Alias(type_="azure-nextgen:apimanagement/v20160707:ApiManagementService"), pulumi.Alias(type_="azure-nextgen:apimanagement/v20170301:ApiManagementService"), pulumi.Alias(type_="azure-nextgen:apimanagement/v20180101:ApiManagementService"), pulumi.Alias(type_="azure-nextgen:apimanagement/v20180601preview:ApiManagementService"), pulumi.Alias(type_="azure-nextgen:apimanagement/v20190101:ApiManagementService"), pulumi.Alias(type_="azure-nextgen:apimanagement/v20191201:ApiManagementService"), pulumi.Alias(type_="azure-nextgen:apimanagement/v20191201preview:ApiManagementService"), pulumi.Alias(type_="azure-nextgen:apimanagement/v20200601preview:ApiManagementService")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(ApiManagementService, __self__).__init__(
            'azure-nextgen:apimanagement/v20161010:ApiManagementService',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ApiManagementService':
        """
        Get an existing ApiManagementService resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ApiManagementService(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="additionalLocations")
    def additional_locations(self) -> pulumi.Output[Optional[Sequence['outputs.AdditionalRegionResponse']]]:
        """
        Additional datacenter locations of the API Management service.
        """
        return pulumi.get(self, "additional_locations")

    @property
    @pulumi.getter(name="addresserEmail")
    def addresser_email(self) -> pulumi.Output[Optional[str]]:
        """
        Addresser email.
        """
        return pulumi.get(self, "addresser_email")

    @property
    @pulumi.getter(name="createdAtUtc")
    def created_at_utc(self) -> pulumi.Output[str]:
        """
        Creation UTC date of the API Management service.The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
        """
        return pulumi.get(self, "created_at_utc")

    @property
    @pulumi.getter(name="customProperties")
    def custom_properties(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Custom properties of the API Management service, like disabling TLS 1.0.
        """
        return pulumi.get(self, "custom_properties")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        ETag of the resource.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="hostnameConfigurations")
    def hostname_configurations(self) -> pulumi.Output[Optional[Sequence['outputs.HostnameConfigurationResponse']]]:
        """
        Custom hostname configuration of the API Management service.
        """
        return pulumi.get(self, "hostname_configurations")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Resource location.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="managementApiUrl")
    def management_api_url(self) -> pulumi.Output[str]:
        """
        Management API endpoint URL of the API Management service.
        """
        return pulumi.get(self, "management_api_url")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[str]]:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="portalUrl")
    def portal_url(self) -> pulumi.Output[str]:
        """
        Publisher portal endpoint Url of the API Management service.
        """
        return pulumi.get(self, "portal_url")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        The current provisioning state of the API Management service which can be one of the following: Created/Activating/Succeeded/Updating/Failed/Stopped/Terminating/TerminationFailed/Deleted.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="publisherEmail")
    def publisher_email(self) -> pulumi.Output[str]:
        """
        Publisher email.
        """
        return pulumi.get(self, "publisher_email")

    @property
    @pulumi.getter(name="publisherName")
    def publisher_name(self) -> pulumi.Output[str]:
        """
        Publisher name.
        """
        return pulumi.get(self, "publisher_name")

    @property
    @pulumi.getter(name="runtimeUrl")
    def runtime_url(self) -> pulumi.Output[str]:
        """
        Proxy endpoint URL of the API Management service.
        """
        return pulumi.get(self, "runtime_url")

    @property
    @pulumi.getter(name="scmUrl")
    def scm_url(self) -> pulumi.Output[str]:
        """
        SCM endpoint URL of the API Management service.
        """
        return pulumi.get(self, "scm_url")

    @property
    @pulumi.getter
    def sku(self) -> pulumi.Output['outputs.ApiManagementServiceSkuPropertiesResponse']:
        """
        SKU properties of the API Management service.
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter(name="staticIPs")
    def static_ips(self) -> pulumi.Output[Sequence[str]]:
        """
        Static IP addresses of the API Management service virtual machines. Available only for Standard and Premium SKU.
        """
        return pulumi.get(self, "static_ips")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="targetProvisioningState")
    def target_provisioning_state(self) -> pulumi.Output[str]:
        """
        The provisioning state of the API Management service, which is targeted by the long running operation started on the service.
        """
        return pulumi.get(self, "target_provisioning_state")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Resource type for API Management resource is set to Microsoft.ApiManagement.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="vpnType")
    def vpn_type(self) -> pulumi.Output[Optional[str]]:
        """
        The type of VPN in which API Management service needs to be configured in. None (Default Value) means the API Management service is not part of any Virtual Network, External means the API Management deployment is set up inside a Virtual Network having an Internet Facing Endpoint, and Internal means that API Management deployment is setup inside a Virtual Network having an Intranet Facing Endpoint only.
        """
        return pulumi.get(self, "vpn_type")

    @property
    @pulumi.getter
    def vpnconfiguration(self) -> pulumi.Output[Optional['outputs.VirtualNetworkConfigurationResponse']]:
        """
        Virtual network configuration of the API Management service.
        """
        return pulumi.get(self, "vpnconfiguration")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

