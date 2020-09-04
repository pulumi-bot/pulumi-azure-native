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

__all__ = ['ApiManagementService']


class ApiManagementService(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 additional_locations: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['AdditionalLocationArgs']]]]] = None,
                 certificates: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['CertificateConfigurationArgs']]]]] = None,
                 custom_properties: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 enable_client_certificate: Optional[pulumi.Input[bool]] = None,
                 hostname_configurations: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['HostnameConfigurationArgs']]]]] = None,
                 identity: Optional[pulumi.Input[pulumi.InputType['ApiManagementServiceIdentityArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 notification_sender_email: Optional[pulumi.Input[str]] = None,
                 publisher_email: Optional[pulumi.Input[str]] = None,
                 publisher_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 sku: Optional[pulumi.Input[pulumi.InputType['ApiManagementServiceSkuPropertiesArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 virtual_network_configuration: Optional[pulumi.Input[pulumi.InputType['VirtualNetworkConfigurationArgs']]] = None,
                 virtual_network_type: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        A single API Management service resource in List or Get response.

        ## Example Usage
        ### ApiManagementCreateMultiRegionServiceWithCustomHostname

        ```python
        import pulumi
        import pulumi_azurerm as azurerm

        api_management_service = azurerm.apimanagement.v20190101.ApiManagementService("apiManagementService",
            additional_locations=[{
                "location": "North Europe",
                "sku": {
                    "capacity": 1,
                    "name": "Premium",
                },
                "virtualNetworkConfiguration": {
                    "subnetResourceId": "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/westUsVirtualNetwork/subnets/apimSubnet",
                },
            }],
            hostname_configurations=[
                {
                    "certificatePassword": "**************Password of the Certificate************************************************",
                    "encodedCertificate": "************Base 64 Encoded Pfx Certificate************************",
                    "hostName": "proxyhostname1.contoso.com",
                    "type": "Proxy",
                },
                {
                    "certificatePassword": "**************Password of the Certificate************************************************",
                    "encodedCertificate": "************Base 64 Encoded Pfx Certificate************************",
                    "hostName": "proxyhostname2.contoso.com",
                    "negotiateClientCertificate": True,
                    "type": "Proxy",
                },
                {
                    "certificatePassword": "**************Password of the Certificate************************************************",
                    "encodedCertificate": "************Base 64 Encoded Pfx Certificate************************",
                    "hostName": "portalhostname1.contoso.com",
                    "type": "Portal",
                },
            ],
            location="Central US",
            publisher_email="admin@live.com",
            publisher_name="contoso",
            resource_group_name="rg1",
            service_name="apimService1",
            sku={
                "capacity": 1,
                "name": "Premium",
            },
            virtual_network_configuration={
                "subnetResourceId": "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/centralUsVirtualNetwork/subnets/apimSubnet",
            },
            virtual_network_type="External")

        ```
        ### ApiManagementCreateService

        ```python
        import pulumi
        import pulumi_azurerm as azurerm

        api_management_service = azurerm.apimanagement.v20190101.ApiManagementService("apiManagementService",
            location="Central US",
            publisher_email="apim@autorestsdk.com",
            publisher_name="autorestsdk",
            resource_group_name="rg1",
            service_name="apimService1",
            sku={
                "capacity": 1,
                "name": "Developer",
            },
            tags={
                "tag1": "value1",
                "tag2": "value2",
                "tag3": "value3",
            })

        ```
        ### ApiManagementCreateServiceHavingMsi

        ```python
        import pulumi
        import pulumi_azurerm as azurerm

        api_management_service = azurerm.apimanagement.v20190101.ApiManagementService("apiManagementService",
            identity={
                "type": "SystemAssigned",
            },
            location="West US",
            publisher_email="apim@autorestsdk.com",
            publisher_name="autorestsdk",
            resource_group_name="rg1",
            service_name="apimService1",
            sku={
                "name": "Consumption",
            },
            tags={
                "tag1": "value1",
                "tag2": "value2",
                "tag3": "value3",
            })

        ```
        ### ApiManagementCreateServiceWithSystemCertificates

        ```python
        import pulumi
        import pulumi_azurerm as azurerm

        api_management_service = azurerm.apimanagement.v20190101.ApiManagementService("apiManagementService",
            certificates=[{
                "certificatePassword": "Password",
                "encodedCertificate": "*******Base64 encoded Certificate******************",
                "storeName": "CertificateAuthority",
            }],
            location="Central US",
            publisher_email="apim@autorestsdk.com",
            publisher_name="autorestsdk",
            resource_group_name="rg1",
            service_name="apimService1",
            sku={
                "capacity": 1,
                "name": "Basic",
            },
            tags={
                "tag1": "value1",
                "tag2": "value2",
                "tag3": "value3",
            })

        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['AdditionalLocationArgs']]]] additional_locations: Additional datacenter locations of the API Management service.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['CertificateConfigurationArgs']]]] certificates: List of Certificates that need to be installed in the API Management service. Max supported certificates that can be installed is 10.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] custom_properties: Custom properties of the API Management service.</br>Setting `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Ciphers.TripleDes168` will disable the cipher TLS_RSA_WITH_3DES_EDE_CBC_SHA for all TLS(1.0, 1.1 and 1.2).</br>Setting `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls11` can be used to disable just TLS 1.1.</br>Setting `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls10` can be used to disable TLS 1.0 on an API Management service.</br>Setting `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Backend.Protocols.Tls11` can be used to disable just TLS 1.1 for communications with backends.</br>Setting `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Backend.Protocols.Tls10` can be used to disable TLS 1.0 for communications with backends.</br>Setting `Microsoft.WindowsAzure.ApiManagement.Gateway.Protocols.Server.Http2` can be used to enable HTTP2 protocol on an API Management service.</br>Not specifying any of these properties on PATCH operation will reset omitted properties' values to their defaults. For all the settings except Http2 the default value is `True` if the service was created on or before April 1st 2018 and `False` otherwise. Http2 setting's default value is `False`.</br></br>You can disable any of next ciphers by using settings `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Ciphers.[cipher_name]`: TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA, TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA, TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA, TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA, TLS_RSA_WITH_AES_128_GCM_SHA256, TLS_RSA_WITH_AES_256_CBC_SHA256, TLS_RSA_WITH_AES_128_CBC_SHA256, TLS_RSA_WITH_AES_256_CBC_SHA, TLS_RSA_WITH_AES_128_CBC_SHA. For example, `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Ciphers.TLS_RSA_WITH_AES_128_CBC_SHA256`:`false`. The default value is `true` for them.
        :param pulumi.Input[bool] enable_client_certificate: Property only meant to be used for Consumption SKU Service. This enforces a client certificate to be presented on each request to the gateway. This also enables the ability to authenticate the certificate in the policy on the gateway.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['HostnameConfigurationArgs']]]] hostname_configurations: Custom hostname configuration of the API Management service.
        :param pulumi.Input[pulumi.InputType['ApiManagementServiceIdentityArgs']] identity: Managed service identity of the Api Management service.
        :param pulumi.Input[str] location: Resource location.
        :param pulumi.Input[str] notification_sender_email: Email address from which the notification will be sent.
        :param pulumi.Input[str] publisher_email: Publisher email.
        :param pulumi.Input[str] publisher_name: Publisher name.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[str] service_name: The name of the API Management service.
        :param pulumi.Input[pulumi.InputType['ApiManagementServiceSkuPropertiesArgs']] sku: SKU properties of the API Management service.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags.
        :param pulumi.Input[pulumi.InputType['VirtualNetworkConfigurationArgs']] virtual_network_configuration: Virtual network configuration of the API Management service.
        :param pulumi.Input[str] virtual_network_type: The type of VPN in which API Management service needs to be configured in. None (Default Value) means the API Management service is not part of any Virtual Network, External means the API Management deployment is set up inside a Virtual Network having an Internet Facing Endpoint, and Internal means that API Management deployment is setup inside a Virtual Network having an Intranet Facing Endpoint only.
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
            __props__['certificates'] = certificates
            __props__['custom_properties'] = custom_properties
            __props__['enable_client_certificate'] = enable_client_certificate
            __props__['hostname_configurations'] = hostname_configurations
            __props__['identity'] = identity
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            __props__['notification_sender_email'] = notification_sender_email
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
            __props__['virtual_network_configuration'] = virtual_network_configuration
            __props__['virtual_network_type'] = virtual_network_type
            __props__['created_at_utc'] = None
            __props__['developer_portal_url'] = None
            __props__['etag'] = None
            __props__['gateway_regional_url'] = None
            __props__['gateway_url'] = None
            __props__['management_api_url'] = None
            __props__['name'] = None
            __props__['portal_url'] = None
            __props__['private_ip_addresses'] = None
            __props__['provisioning_state'] = None
            __props__['public_ip_addresses'] = None
            __props__['scm_url'] = None
            __props__['target_provisioning_state'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:apimanagement/latest:ApiManagementService"), pulumi.Alias(type_="azurerm:apimanagement/v20160707:ApiManagementService"), pulumi.Alias(type_="azurerm:apimanagement/v20161010:ApiManagementService"), pulumi.Alias(type_="azurerm:apimanagement/v20170301:ApiManagementService"), pulumi.Alias(type_="azurerm:apimanagement/v20180101:ApiManagementService"), pulumi.Alias(type_="azurerm:apimanagement/v20180601preview:ApiManagementService"), pulumi.Alias(type_="azurerm:apimanagement/v20191201:ApiManagementService"), pulumi.Alias(type_="azurerm:apimanagement/v20191201preview:ApiManagementService")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(ApiManagementService, __self__).__init__(
            'azurerm:apimanagement/v20190101:ApiManagementService',
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
    def additional_locations(self) -> pulumi.Output[Optional[List['outputs.AdditionalLocationResponse']]]:
        """
        Additional datacenter locations of the API Management service.
        """
        return pulumi.get(self, "additional_locations")

    @property
    @pulumi.getter
    def certificates(self) -> pulumi.Output[Optional[List['outputs.CertificateConfigurationResponse']]]:
        """
        List of Certificates that need to be installed in the API Management service. Max supported certificates that can be installed is 10.
        """
        return pulumi.get(self, "certificates")

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
        Custom properties of the API Management service.</br>Setting `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Ciphers.TripleDes168` will disable the cipher TLS_RSA_WITH_3DES_EDE_CBC_SHA for all TLS(1.0, 1.1 and 1.2).</br>Setting `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls11` can be used to disable just TLS 1.1.</br>Setting `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Protocols.Tls10` can be used to disable TLS 1.0 on an API Management service.</br>Setting `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Backend.Protocols.Tls11` can be used to disable just TLS 1.1 for communications with backends.</br>Setting `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Backend.Protocols.Tls10` can be used to disable TLS 1.0 for communications with backends.</br>Setting `Microsoft.WindowsAzure.ApiManagement.Gateway.Protocols.Server.Http2` can be used to enable HTTP2 protocol on an API Management service.</br>Not specifying any of these properties on PATCH operation will reset omitted properties' values to their defaults. For all the settings except Http2 the default value is `True` if the service was created on or before April 1st 2018 and `False` otherwise. Http2 setting's default value is `False`.</br></br>You can disable any of next ciphers by using settings `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Ciphers.[cipher_name]`: TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA, TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA, TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA, TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA, TLS_RSA_WITH_AES_128_GCM_SHA256, TLS_RSA_WITH_AES_256_CBC_SHA256, TLS_RSA_WITH_AES_128_CBC_SHA256, TLS_RSA_WITH_AES_256_CBC_SHA, TLS_RSA_WITH_AES_128_CBC_SHA. For example, `Microsoft.WindowsAzure.ApiManagement.Gateway.Security.Ciphers.TLS_RSA_WITH_AES_128_CBC_SHA256`:`false`. The default value is `true` for them.
        """
        return pulumi.get(self, "custom_properties")

    @property
    @pulumi.getter(name="developerPortalUrl")
    def developer_portal_url(self) -> pulumi.Output[str]:
        """
        Developer Portal endpoint URL of the API Management service.
        """
        return pulumi.get(self, "developer_portal_url")

    @property
    @pulumi.getter(name="enableClientCertificate")
    def enable_client_certificate(self) -> pulumi.Output[Optional[bool]]:
        """
        Property only meant to be used for Consumption SKU Service. This enforces a client certificate to be presented on each request to the gateway. This also enables the ability to authenticate the certificate in the policy on the gateway.
        """
        return pulumi.get(self, "enable_client_certificate")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        ETag of the resource.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="gatewayRegionalUrl")
    def gateway_regional_url(self) -> pulumi.Output[str]:
        """
        Gateway URL of the API Management service in the Default Region.
        """
        return pulumi.get(self, "gateway_regional_url")

    @property
    @pulumi.getter(name="gatewayUrl")
    def gateway_url(self) -> pulumi.Output[str]:
        """
        Gateway URL of the API Management service.
        """
        return pulumi.get(self, "gateway_url")

    @property
    @pulumi.getter(name="hostnameConfigurations")
    def hostname_configurations(self) -> pulumi.Output[Optional[List['outputs.HostnameConfigurationResponse']]]:
        """
        Custom hostname configuration of the API Management service.
        """
        return pulumi.get(self, "hostname_configurations")

    @property
    @pulumi.getter
    def identity(self) -> pulumi.Output[Optional['outputs.ApiManagementServiceIdentityResponse']]:
        """
        Managed service identity of the Api Management service.
        """
        return pulumi.get(self, "identity")

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
    def name(self) -> pulumi.Output[str]:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="notificationSenderEmail")
    def notification_sender_email(self) -> pulumi.Output[Optional[str]]:
        """
        Email address from which the notification will be sent.
        """
        return pulumi.get(self, "notification_sender_email")

    @property
    @pulumi.getter(name="portalUrl")
    def portal_url(self) -> pulumi.Output[str]:
        """
        Publisher portal endpoint Url of the API Management service.
        """
        return pulumi.get(self, "portal_url")

    @property
    @pulumi.getter(name="privateIPAddresses")
    def private_ip_addresses(self) -> pulumi.Output[List[str]]:
        """
        Private Static Load Balanced IP addresses of the API Management service in Primary region which is deployed in an Internal Virtual Network. Available only for Basic, Standard and Premium SKU.
        """
        return pulumi.get(self, "private_ip_addresses")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        The current provisioning state of the API Management service which can be one of the following: Created/Activating/Succeeded/Updating/Failed/Stopped/Terminating/TerminationFailed/Deleted.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="publicIPAddresses")
    def public_ip_addresses(self) -> pulumi.Output[List[str]]:
        """
        Public Static Load Balanced IP addresses of the API Management service in Primary region. Available only for Basic, Standard and Premium SKU.
        """
        return pulumi.get(self, "public_ip_addresses")

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
    @pulumi.getter(name="virtualNetworkConfiguration")
    def virtual_network_configuration(self) -> pulumi.Output[Optional['outputs.VirtualNetworkConfigurationResponse']]:
        """
        Virtual network configuration of the API Management service.
        """
        return pulumi.get(self, "virtual_network_configuration")

    @property
    @pulumi.getter(name="virtualNetworkType")
    def virtual_network_type(self) -> pulumi.Output[Optional[str]]:
        """
        The type of VPN in which API Management service needs to be configured in. None (Default Value) means the API Management service is not part of any Virtual Network, External means the API Management deployment is set up inside a Virtual Network having an Internet Facing Endpoint, and Internal means that API Management deployment is setup inside a Virtual Network having an Intranet Facing Endpoint only.
        """
        return pulumi.get(self, "virtual_network_type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

