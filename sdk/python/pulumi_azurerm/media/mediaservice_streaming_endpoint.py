# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class MediaserviceStreamingEndpoint(pulumi.CustomResource):
    location: pulumi.Output[str]
    """
    The geo-location where the resource lives
    """
    name: pulumi.Output[str]
    """
    The name of the resource
    """
    properties: pulumi.Output[dict]
    """
    The StreamingEndpoint properties.
      * `access_control` (`dict`) - The access control definition of the StreamingEndpoint.
        * `akamai` (`dict`) - The access control of Akamai
          * `akamai_signature_header_authentication_key_list` (`list`) - authentication key list
            * `base64_key` (`str`) - authentication key
            * `expiration` (`str`) - The expiration time of the authentication key.
            * `identifier` (`str`) - identifier of the key

        * `ip` (`dict`) - The IP access control of the StreamingEndpoint.
          * `allow` (`list`) - The IP allow list.
            * `address` (`str`) - The IP address.
            * `name` (`str`) - The friendly name for the IP address range.
            * `subnet_prefix_length` (`float`) - The subnet mask prefix length (see CIDR notation).

      * `availability_set_name` (`str`) - The name of the AvailabilitySet used with this StreamingEndpoint for high availability streaming.  This value can only be set at creation time.
      * `cdn_enabled` (`bool`) - The CDN enabled flag.
      * `cdn_profile` (`str`) - The CDN profile name.
      * `cdn_provider` (`str`) - The CDN provider name.
      * `created` (`str`) - The exact time the StreamingEndpoint was created.
      * `cross_site_access_policies` (`dict`) - The StreamingEndpoint access policies.
        * `client_access_policy` (`str`) - The content of clientaccesspolicy.xml used by Silverlight.
        * `cross_domain_policy` (`str`) - The content of crossdomain.xml used by Silverlight.

      * `custom_host_names` (`list`) - The custom host names of the StreamingEndpoint
      * `description` (`str`) - The StreamingEndpoint description.
      * `free_trial_end_time` (`str`) - The free trial expiration time.
      * `host_name` (`str`) - The StreamingEndpoint host name.
      * `last_modified` (`str`) - The exact time the StreamingEndpoint was last modified.
      * `max_cache_age` (`float`) - Max cache age
      * `provisioning_state` (`str`) - The provisioning state of the StreamingEndpoint.
      * `resource_state` (`str`) - The resource state of the StreamingEndpoint.
      * `scale_units` (`float`) - The number of scale units.  Use the Scale operation to adjust this value.
    """
    tags: pulumi.Output[dict]
    """
    Resource tags.
    """
    type: pulumi.Output[str]
    """
    The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
    """
    def __init__(__self__, resource_name, opts=None, account_name=None, auto_start=None, location=None, name=None, properties=None, resource_group_name=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        The StreamingEndpoint.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: The Media Services account name.
        :param pulumi.Input[bool] auto_start: The flag indicates if the resource should be automatically started on creation.
        :param pulumi.Input[str] location: The geo-location where the resource lives
        :param pulumi.Input[str] name: The name of the StreamingEndpoint.
        :param pulumi.Input[dict] properties: The StreamingEndpoint properties.
        :param pulumi.Input[str] resource_group_name: The name of the resource group within the Azure subscription.
        :param pulumi.Input[dict] tags: Resource tags.

        The **properties** object supports the following:

          * `access_control` (`pulumi.Input[dict]`) - The access control definition of the StreamingEndpoint.
            * `akamai` (`pulumi.Input[dict]`) - The access control of Akamai
              * `akamai_signature_header_authentication_key_list` (`pulumi.Input[list]`) - authentication key list
                * `base64_key` (`pulumi.Input[str]`) - authentication key
                * `expiration` (`pulumi.Input[str]`) - The expiration time of the authentication key.
                * `identifier` (`pulumi.Input[str]`) - identifier of the key

            * `ip` (`pulumi.Input[dict]`) - The IP access control of the StreamingEndpoint.
              * `allow` (`pulumi.Input[list]`) - The IP allow list.
                * `address` (`pulumi.Input[str]`) - The IP address.
                * `name` (`pulumi.Input[str]`) - The friendly name for the IP address range.
                * `subnet_prefix_length` (`pulumi.Input[float]`) - The subnet mask prefix length (see CIDR notation).

          * `availability_set_name` (`pulumi.Input[str]`) - The name of the AvailabilitySet used with this StreamingEndpoint for high availability streaming.  This value can only be set at creation time.
          * `cdn_enabled` (`pulumi.Input[bool]`) - The CDN enabled flag.
          * `cdn_profile` (`pulumi.Input[str]`) - The CDN profile name.
          * `cdn_provider` (`pulumi.Input[str]`) - The CDN provider name.
          * `cross_site_access_policies` (`pulumi.Input[dict]`) - The StreamingEndpoint access policies.
            * `client_access_policy` (`pulumi.Input[str]`) - The content of clientaccesspolicy.xml used by Silverlight.
            * `cross_domain_policy` (`pulumi.Input[str]`) - The content of crossdomain.xml used by Silverlight.

          * `custom_host_names` (`pulumi.Input[list]`) - The custom host names of the StreamingEndpoint
          * `description` (`pulumi.Input[str]`) - The StreamingEndpoint description.
          * `max_cache_age` (`pulumi.Input[float]`) - Max cache age
          * `scale_units` (`pulumi.Input[float]`) - The number of scale units.  Use the Scale operation to adjust this value.
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if account_name is None:
                raise TypeError("Missing required property 'account_name'")
            __props__['account_name'] = account_name
            __props__['auto_start'] = auto_start
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['type'] = None
        super(MediaserviceStreamingEndpoint, __self__).__init__(
            'azurerm:media:MediaserviceStreamingEndpoint',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing MediaserviceStreamingEndpoint resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return MediaserviceStreamingEndpoint(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
