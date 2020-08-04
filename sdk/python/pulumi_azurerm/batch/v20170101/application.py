# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class Application(pulumi.CustomResource):
    allow_updates: pulumi.Output[bool]
    """
    A value indicating whether packages within the application may be overwritten using the same version string.
    """
    default_version: pulumi.Output[str]
    """
    The package to use if a client requests the application but does not specify a version.
    """
    display_name: pulumi.Output[str]
    """
    The display name for the application.
    """
    packages: pulumi.Output[list]
    """
    The list of packages under this application.
      * `format` (`str`) - The format of the application package, if the package is active.
      * `id` (`str`) - The ID of the application.
      * `last_activation_time` (`str`) - The time at which the package was last activated, if the package is active.
      * `state` (`str`) - The current state of the application package.
      * `storage_url` (`str`) - The storage URL at which the application package is stored.
      * `storage_url_expiry` (`str`) - The UTC time at which the storage URL will expire.
      * `version` (`str`) - The version of the application package. 
    """
    def __init__(__self__, resource_name, opts=None, account_name=None, allow_updates=None, application_id=None, display_name=None, resource_group_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Contains information about an application in a Batch account.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: The name of the Batch account.
        :param pulumi.Input[bool] allow_updates: A value indicating whether packages within the application may be overwritten using the same version string.
        :param pulumi.Input[str] application_id: The ID of the application.
        :param pulumi.Input[str] display_name: The display name for the application.
        :param pulumi.Input[str] resource_group_name: The name of the resource group that contains the Batch account.
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

            if account_name is None:
                raise TypeError("Missing required property 'account_name'")
            __props__['account_name'] = account_name
            __props__['allow_updates'] = allow_updates
            if application_id is None:
                raise TypeError("Missing required property 'application_id'")
            __props__['application_id'] = application_id
            __props__['display_name'] = display_name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['default_version'] = None
            __props__['packages'] = None
        super(Application, __self__).__init__(
            'azurerm:batch/v20170101:Application',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing Application resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Application(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
