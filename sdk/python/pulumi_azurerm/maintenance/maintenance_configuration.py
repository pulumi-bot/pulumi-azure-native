# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class MaintenanceConfiguration(pulumi.CustomResource):
    location: pulumi.Output[str]
    """
    Gets or sets location of the resource
    """
    name: pulumi.Output[str]
    """
    Name of the resource
    """
    properties: pulumi.Output[dict]
    """
    Gets or sets properties of the resource
      * `extension_properties` (`dict`) - Gets or sets extensionProperties of the maintenanceConfiguration. This is for future use only and would be a set of key value pairs for additional information e.g. whether to follow SDP etc.
      * `maintenance_scope` (`str`) - Gets or sets maintenanceScope of the configuration. It represent the impact area of the maintenance
      * `namespace` (`str`) - Gets or sets namespace of the resource e.g. Microsoft.Maintenance or Microsoft.Sql
    """
    tags: pulumi.Output[dict]
    """
    Gets or sets tags of the resource
    """
    type: pulumi.Output[str]
    """
    Type of the resource
    """
    def __init__(__self__, resource_name, opts=None, location=None, name=None, properties=None, resource_group_name=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Maintenance configuration record type

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] location: Gets or sets location of the resource
        :param pulumi.Input[str] name: Resource Identifier
        :param pulumi.Input[dict] properties: Gets or sets properties of the resource
        :param pulumi.Input[str] resource_group_name: Resource Group Name
        :param pulumi.Input[dict] tags: Gets or sets tags of the resource

        The **properties** object supports the following:

          * `extension_properties` (`pulumi.Input[dict]`) - Gets or sets extensionProperties of the maintenanceConfiguration. This is for future use only and would be a set of key value pairs for additional information e.g. whether to follow SDP etc.
          * `maintenance_scope` (`pulumi.Input[str]`) - Gets or sets maintenanceScope of the configuration. It represent the impact area of the maintenance
          * `namespace` (`pulumi.Input[str]`) - Gets or sets namespace of the resource e.g. Microsoft.Maintenance or Microsoft.Sql
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
        super(MaintenanceConfiguration, __self__).__init__(
            'azurerm:maintenance:MaintenanceConfiguration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing MaintenanceConfiguration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return MaintenanceConfiguration(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
