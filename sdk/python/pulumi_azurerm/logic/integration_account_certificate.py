# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class IntegrationAccountCertificate(pulumi.CustomResource):
    location: pulumi.Output[str]
    """
    The resource location.
    """
    name: pulumi.Output[str]
    """
    Gets the resource name.
    """
    properties: pulumi.Output[dict]
    """
    The integration account certificate properties.
      * `changed_time` (`str`) - The changed time.
      * `created_time` (`str`) - The created time.
      * `key` (`dict`) - The key details in the key vault.
        * `key_name` (`str`) - The private key name in key vault.
        * `key_vault` (`dict`) - The key vault reference.
          * `id` (`str`) - The resource id.
          * `name` (`str`) - The resource name.
          * `type` (`str`) - The resource type.

        * `key_version` (`str`) - The private key version in key vault.

      * `metadata` (`dict`) - The metadata.
      * `public_certificate` (`str`) - The public certificate.
    """
    tags: pulumi.Output[dict]
    """
    The resource tags.
    """
    type: pulumi.Output[str]
    """
    Gets the resource type.
    """
    def __init__(__self__, resource_name, opts=None, integration_account_name=None, location=None, name=None, properties=None, resource_group_name=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        The integration account certificate.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] integration_account_name: The integration account name.
        :param pulumi.Input[str] location: The resource location.
        :param pulumi.Input[str] name: The integration account certificate name.
        :param pulumi.Input[dict] properties: The integration account certificate properties.
        :param pulumi.Input[str] resource_group_name: The resource group name.
        :param pulumi.Input[dict] tags: The resource tags.

        The **properties** object supports the following:

          * `key` (`pulumi.Input[dict]`) - The key details in the key vault.
            * `key_name` (`pulumi.Input[str]`) - The private key name in key vault.
            * `key_vault` (`pulumi.Input[dict]`) - The key vault reference.
              * `id` (`pulumi.Input[str]`) - The resource id.

            * `key_version` (`pulumi.Input[str]`) - The private key version in key vault.

          * `metadata` (`pulumi.Input[dict]`) - The metadata.
          * `public_certificate` (`pulumi.Input[str]`) - The public certificate.
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

            if integration_account_name is None:
                raise TypeError("Missing required property 'integration_account_name'")
            __props__['integration_account_name'] = integration_account_name
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if properties is None:
                raise TypeError("Missing required property 'properties'")
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['type'] = None
        super(IntegrationAccountCertificate, __self__).__init__(
            'azurerm:logic:IntegrationAccountCertificate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing IntegrationAccountCertificate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return IntegrationAccountCertificate(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
