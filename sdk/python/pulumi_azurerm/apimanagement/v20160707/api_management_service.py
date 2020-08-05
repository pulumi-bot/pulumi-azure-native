# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class ApiManagementService(pulumi.CustomResource):
    etag: pulumi.Output[str]
    """
    ETag of the resource.
    """
    location: pulumi.Output[str]
    """
    Datacenter location of the API Management service.
    """
    name: pulumi.Output[str]
    """
    Name of the API Management service.
    """
    properties: pulumi.Output[dict]
    """
    Properties of the API Management service.
      * `additional_locations` (`list`) - Additional datacenter locations of the API Management service.
        * `location` (`str`) - The location name of the additional region among Azure Data center regions.
        * `sku_type` (`str`) - The SKU type in the location.
        * `sku_unit_count` (`float`) - The SKU Unit count at the location. The maximum SKU Unit count depends on the SkuType. Maximum allowed for Developer SKU is 1, for Standard SKU is 4, and for Premium SKU is 10, at a location.
        * `static_i_ps` (`list`) - Static IP addresses of the location's virtual machines.
        * `vpnconfiguration` (`dict`) - Virtual network configuration for the location.
          * `location` (`str`) - The location of the virtual network.
          * `subnet_resource_id` (`str`) - The name of the subnet Resource ID. This has format /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/Microsoft.{Network|ClassicNetwork}/VirtualNetworks/{virtual network name}/subnets/{subnet name}.
          * `subnetname` (`str`) - The name of the subnet.
          * `vnetid` (`str`) - The virtual network ID. This is typically a GUID. Expect a null GUID by default.

      * `addresser_email` (`str`) - Addresser email.
      * `created_at_utc` (`str`) - Creation UTC date of the API Management service.The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
      * `custom_properties` (`dict`) - Custom properties of the API Management service, like disabling TLS 1.0.
      * `hostname_configurations` (`list`) - Custom hostname configuration of the API Management service.
        * `certificate` (`dict`) - Certificate information.
          * `expiry` (`str`) - Expiration date of the certificate. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
          * `subject` (`str`) - Subject of the certificate.
          * `thumbprint` (`str`) - Thumbprint of the certificate.

        * `hostname` (`str`) - Hostname.
        * `type` (`str`) - Hostname type.

      * `management_api_url` (`str`) - Management API endpoint URL of the API Management service.
      * `portal_url` (`str`) - Publisher portal endpoint Url of the API Management service.
      * `provisioning_state` (`str`) - The current provisioning state of the API Management service which can be one of the following: Created/Activating/Succeeded/Updating/Failed/Stopped/Terminating/TerminationFailed/Deleted.
      * `publisher_email` (`str`) - Publisher email.
      * `publisher_name` (`str`) - Publisher name.
      * `runtime_url` (`str`) - Proxy endpoint URL of the API Management service.
      * `scm_url` (`str`) - SCM endpoint URL of the API Management service.
      * `static_i_ps` (`list`) - Static IP addresses of the API Management service virtual machines. Available only for Standard and Premium SKU.
      * `target_provisioning_state` (`str`) - The provisioning state of the API Management service, which is targeted by the long running operation started on the service.
      * `vpn_type` (`str`) - The type of VPN in which API Management service needs to be configured in. None (Default Value) means the API Management service is not part of any Virtual Network, External means the API Management deployment is set up inside a Virtual Network having an Internet Facing Endpoint, and Internal means that API Management deployment is setup inside a Virtual Network having an Intranet Facing Endpoint only.
      * `vpnconfiguration` (`dict`) - Virtual network configuration of the API Management service.
    """
    sku: pulumi.Output[dict]
    """
    SKU properties of the API Management service.
      * `capacity` (`float`) - Capacity of the SKU (number of deployed units of the SKU). The default value is 1.
      * `name` (`str`) - Name of the Sku.
    """
    tags: pulumi.Output[dict]
    """
    API Management service tags. A maximum of 10 tags can be provided for a resource, and each tag must have a key no greater than 128 characters (and a value no greater than 256 characters).
    """
    type: pulumi.Output[str]
    """
    Resource type of the API Management service.
    """
    def __init__(__self__, resource_name, opts=None, additional_locations=None, addresser_email=None, custom_properties=None, etag=None, hostname_configurations=None, location=None, name=None, publisher_email=None, publisher_name=None, resource_group_name=None, sku=None, tags=None, vpn_type=None, vpnconfiguration=None, __props__=None, __name__=None, __opts__=None):
        """
        Description of an API Management service resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] additional_locations: Additional datacenter locations of the API Management service.
        :param pulumi.Input[str] addresser_email: Addresser email.
        :param pulumi.Input[dict] custom_properties: Custom properties of the API Management service, like disabling TLS 1.0.
        :param pulumi.Input[str] etag: ETag of the resource.
        :param pulumi.Input[list] hostname_configurations: Custom hostname configuration of the API Management service.
        :param pulumi.Input[str] location: Datacenter location of the API Management service.
        :param pulumi.Input[str] name: The name of the API Management service.
        :param pulumi.Input[str] publisher_email: Publisher email.
        :param pulumi.Input[str] publisher_name: Publisher name.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[dict] sku: SKU properties of the API Management service.
        :param pulumi.Input[dict] tags: API Management service tags. A maximum of 10 tags can be provided for a resource, and each tag must have a key no greater than 128 characters (and a value no greater than 256 characters).
        :param pulumi.Input[str] vpn_type: The type of VPN in which API Management service needs to be configured in. None (Default Value) means the API Management service is not part of any Virtual Network, External means the API Management deployment is set up inside a Virtual Network having an Internet Facing Endpoint, and Internal means that API Management deployment is setup inside a Virtual Network having an Intranet Facing Endpoint only.
        :param pulumi.Input[dict] vpnconfiguration: Virtual network configuration of the API Management service.

        The **additional_locations** object supports the following:

          * `location` (`pulumi.Input[str]`) - The location name of the additional region among Azure Data center regions.
          * `sku_type` (`pulumi.Input[str]`) - The SKU type in the location.
          * `sku_unit_count` (`pulumi.Input[float]`) - The SKU Unit count at the location. The maximum SKU Unit count depends on the SkuType. Maximum allowed for Developer SKU is 1, for Standard SKU is 4, and for Premium SKU is 10, at a location.
          * `vpnconfiguration` (`pulumi.Input[dict]`) - Virtual network configuration for the location.
            * `location` (`pulumi.Input[str]`) - The location of the virtual network.
            * `subnet_resource_id` (`pulumi.Input[str]`) - The name of the subnet Resource ID. This has format /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/Microsoft.{Network|ClassicNetwork}/VirtualNetworks/{virtual network name}/subnets/{subnet name}.

        The **hostname_configurations** object supports the following:

          * `certificate` (`pulumi.Input[dict]`) - Certificate information.
            * `expiry` (`pulumi.Input[str]`) - Expiration date of the certificate. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
            * `subject` (`pulumi.Input[str]`) - Subject of the certificate.
            * `thumbprint` (`pulumi.Input[str]`) - Thumbprint of the certificate.

          * `hostname` (`pulumi.Input[str]`) - Hostname.
          * `type` (`pulumi.Input[str]`) - Hostname type.

        The **sku** object supports the following:

          * `capacity` (`pulumi.Input[float]`) - Capacity of the SKU (number of deployed units of the SKU). The default value is 1.
          * `name` (`pulumi.Input[str]`) - Name of the Sku.
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
            __props__['etag'] = etag
            __props__['hostname_configurations'] = hostname_configurations
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
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
            if sku is None:
                raise TypeError("Missing required property 'sku'")
            __props__['sku'] = sku
            __props__['tags'] = tags
            __props__['vpn_type'] = vpn_type
            __props__['vpnconfiguration'] = vpnconfiguration
            __props__['properties'] = None
            __props__['type'] = None
        super(ApiManagementService, __self__).__init__(
            'azurerm:apimanagement/v20160707:ApiManagementService',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing ApiManagementService resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ApiManagementService(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
