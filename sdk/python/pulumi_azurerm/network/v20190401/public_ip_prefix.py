# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class PublicIPPrefix(pulumi.CustomResource):
    etag: pulumi.Output[str]
    """
    A unique read-only string that changes whenever the resource is updated.
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
    Public IP prefix properties.
      * `ip_prefix` (`str`) - The allocated Prefix.
      * `ip_tags` (`list`) - The list of tags associated with the public IP prefix.
        * `ip_tag_type` (`str`) - Gets or sets the ipTag type: Example FirstPartyUsage.
        * `tag` (`str`) - Gets or sets value of the IpTag associated with the public IP. Example SQL, Storage etc.

      * `load_balancer_frontend_ip_configuration` (`dict`) - The reference to load balancer frontend IP configuration associated with the public IP prefix.
        * `id` (`str`) - Resource ID.

      * `prefix_length` (`float`) - The Length of the Public IP Prefix.
      * `provisioning_state` (`str`) - The provisioning state of the Public IP prefix resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
      * `public_ip_address_version` (`str`) - The public IP address version.
      * `public_ip_addresses` (`list`) - The list of all referenced PublicIPAddresses.
        * `id` (`str`) - The PublicIPAddress Reference.

      * `resource_guid` (`str`) - The resource GUID property of the public IP prefix resource.
    """
    sku: pulumi.Output[dict]
    """
    The public IP prefix SKU.
      * `name` (`str`) - Name of a public IP prefix SKU.
    """
    tags: pulumi.Output[dict]
    """
    Resource tags.
    """
    type: pulumi.Output[str]
    """
    Resource type.
    """
    zones: pulumi.Output[list]
    """
    A list of availability zones denoting the IP allocated for the resource needs to come from.
    """
    def __init__(__self__, resource_name, opts=None, etag=None, id=None, ip_prefix=None, ip_tags=None, location=None, name=None, prefix_length=None, provisioning_state=None, public_ip_address_version=None, public_ip_addresses=None, resource_group_name=None, resource_guid=None, sku=None, tags=None, zones=None, __props__=None, __name__=None, __opts__=None):
        """
        Public IP prefix resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] etag: A unique read-only string that changes whenever the resource is updated.
        :param pulumi.Input[str] id: Resource ID.
        :param pulumi.Input[str] ip_prefix: The allocated Prefix.
        :param pulumi.Input[list] ip_tags: The list of tags associated with the public IP prefix.
        :param pulumi.Input[str] location: Resource location.
        :param pulumi.Input[str] name: The name of the public IP prefix.
        :param pulumi.Input[float] prefix_length: The Length of the Public IP Prefix.
        :param pulumi.Input[str] provisioning_state: The provisioning state of the Public IP prefix resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        :param pulumi.Input[str] public_ip_address_version: The public IP address version.
        :param pulumi.Input[list] public_ip_addresses: The list of all referenced PublicIPAddresses.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[str] resource_guid: The resource GUID property of the public IP prefix resource.
        :param pulumi.Input[dict] sku: The public IP prefix SKU.
        :param pulumi.Input[dict] tags: Resource tags.
        :param pulumi.Input[list] zones: A list of availability zones denoting the IP allocated for the resource needs to come from.

        The **ip_tags** object supports the following:

          * `ip_tag_type` (`pulumi.Input[str]`) - Gets or sets the ipTag type: Example FirstPartyUsage.
          * `tag` (`pulumi.Input[str]`) - Gets or sets value of the IpTag associated with the public IP. Example SQL, Storage etc.

        The **public_ip_addresses** object supports the following:

          * `id` (`pulumi.Input[str]`) - The PublicIPAddress Reference.

        The **sku** object supports the following:

          * `name` (`pulumi.Input[str]`) - Name of a public IP prefix SKU.
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

            __props__['etag'] = etag
            __props__['id'] = id
            __props__['ip_prefix'] = ip_prefix
            __props__['ip_tags'] = ip_tags
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['prefix_length'] = prefix_length
            __props__['provisioning_state'] = provisioning_state
            __props__['public_ip_address_version'] = public_ip_address_version
            __props__['public_ip_addresses'] = public_ip_addresses
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['resource_guid'] = resource_guid
            __props__['sku'] = sku
            __props__['tags'] = tags
            __props__['zones'] = zones
            __props__['properties'] = None
            __props__['type'] = None
        super(PublicIPPrefix, __self__).__init__(
            'azurerm:network/v20190401:PublicIPPrefix',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing PublicIPPrefix resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return PublicIPPrefix(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
