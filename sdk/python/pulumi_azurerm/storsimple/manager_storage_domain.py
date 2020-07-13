# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class ManagerStorageDomain(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    The name.
    """
    properties: pulumi.Output[dict]
    """
    The properties.
      * `encryption_key` (`dict`) - The encryption key used to encrypt the data. This is a user secret.
        * `encryption_algorithm` (`str`) - The algorithm used to encrypt "Value".
        * `encryption_cert_thumbprint` (`str`) - Thumbprint certificate that was used to encrypt "Value". If the value in unencrypted, it will be null.
        * `value` (`str`) - The value of the secret.

      * `encryption_status` (`str`) - The encryption status "Enabled | Disabled".
      * `storage_account_credential_ids` (`list`) - The storage account credentials.
    """
    type: pulumi.Output[str]
    """
    The type.
    """
    def __init__(__self__, resource_name, opts=None, manager_name=None, name=None, properties=None, resource_group_name=None, __props__=None, __name__=None, __opts__=None):
        """
        The storage domain.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] manager_name: The manager name
        :param pulumi.Input[str] name: The storage domain name.
        :param pulumi.Input[dict] properties: The properties.
        :param pulumi.Input[str] resource_group_name: The resource group name

        The **properties** object supports the following:

          * `encryption_key` (`pulumi.Input[dict]`) - The encryption key used to encrypt the data. This is a user secret.
            * `encryption_algorithm` (`pulumi.Input[str]`) - The algorithm used to encrypt "Value".
            * `encryption_cert_thumbprint` (`pulumi.Input[str]`) - Thumbprint certificate that was used to encrypt "Value". If the value in unencrypted, it will be null.
            * `value` (`pulumi.Input[str]`) - The value of the secret.

          * `encryption_status` (`pulumi.Input[str]`) - The encryption status "Enabled | Disabled".
          * `storage_account_credential_ids` (`pulumi.Input[list]`) - The storage account credentials.
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

            if manager_name is None:
                raise TypeError("Missing required property 'manager_name'")
            __props__['manager_name'] = manager_name
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if properties is None:
                raise TypeError("Missing required property 'properties'")
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['type'] = None
        super(ManagerStorageDomain, __self__).__init__(
            'azurerm:storsimple:ManagerStorageDomain',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing ManagerStorageDomain resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ManagerStorageDomain(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
