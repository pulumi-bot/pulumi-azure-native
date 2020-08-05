# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class ServerDetails(pulumi.CustomResource):
    location: pulumi.Output[str]
    """
    Location of the Analysis Services resource.
    """
    name: pulumi.Output[str]
    """
    The name of the Analysis Services resource.
    """
    properties: pulumi.Output[dict]
    """
    Properties of the provision operation request.
      * `as_administrators` (`dict`) - A collection of AS server administrators
        * `members` (`list`) - An array of administrator user identities.

      * `backup_blob_container_uri` (`str`) - The container URI of backup blob.
      * `provisioning_state` (`str`) - The current deployment state of Analysis Services resource. The provisioningState is to indicate states for resource provisioning.
      * `server_full_name` (`str`) - The full name of the Analysis Services resource.
      * `state` (`str`) - The current state of Analysis Services resource. The state is to indicate more states outside of resource provisioning.
    """
    sku: pulumi.Output[dict]
    """
    The SKU of the Analysis Services resource.
      * `name` (`str`) - Name of the SKU level.
      * `tier` (`str`) - The name of the Azure pricing tier to which the SKU applies.
    """
    tags: pulumi.Output[dict]
    """
    Key-value pairs of additional resource provisioning properties.
    """
    type: pulumi.Output[str]
    """
    The type of the Analysis Services resource.
    """
    def __init__(__self__, resource_name, opts=None, as_administrators=None, backup_blob_container_uri=None, location=None, name=None, resource_group_name=None, sku=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Represents an instance of an Analysis Services resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] as_administrators: A collection of AS server administrators
        :param pulumi.Input[str] backup_blob_container_uri: The container URI of backup blob.
        :param pulumi.Input[str] location: Location of the Analysis Services resource.
        :param pulumi.Input[str] name: The name of the Analysis Services server. It must be a minimum of 3 characters, and a maximum of 63.
        :param pulumi.Input[str] resource_group_name: The name of the Azure Resource group of which a given Analysis Services server is part. This name must be at least 1 character in length, and no more than 90.
        :param pulumi.Input[dict] sku: The SKU of the Analysis Services resource.
        :param pulumi.Input[dict] tags: Key-value pairs of additional resource provisioning properties.

        The **as_administrators** object supports the following:

          * `members` (`pulumi.Input[list]`) - An array of administrator user identities.

        The **sku** object supports the following:

          * `name` (`pulumi.Input[str]`) - Name of the SKU level.
          * `tier` (`pulumi.Input[str]`) - The name of the Azure pricing tier to which the SKU applies.
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

            __props__['as_administrators'] = as_administrators
            __props__['backup_blob_container_uri'] = backup_blob_container_uri
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if sku is None:
                raise TypeError("Missing required property 'sku'")
            __props__['sku'] = sku
            __props__['tags'] = tags
            __props__['properties'] = None
            __props__['type'] = None
        super(ServerDetails, __self__).__init__(
            'azurerm:analysisservices/v20160516:ServerDetails',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing ServerDetails resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ServerDetails(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
