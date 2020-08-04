# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class Certificate(pulumi.CustomResource):
    kind: pulumi.Output[str]
    """
    Kind of resource
    """
    location: pulumi.Output[str]
    """
    Resource Location
    """
    name: pulumi.Output[str]
    """
    Resource Name
    """
    properties: pulumi.Output[dict]
    tags: pulumi.Output[dict]
    """
    Resource tags
    """
    type: pulumi.Output[str]
    """
    Resource type
    """
    def __init__(__self__, resource_name, opts=None, id=None, kind=None, location=None, name=None, properties=None, resource_group_name=None, tags=None, type=None, __props__=None, __name__=None, __opts__=None):
        """
        App certificate

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] id: Resource Id
        :param pulumi.Input[str] kind: Kind of resource
        :param pulumi.Input[str] location: Resource Location
        :param pulumi.Input[str] name: Resource Name
        :param pulumi.Input[str] resource_group_name: Name of the resource group
        :param pulumi.Input[dict] tags: Resource tags
        :param pulumi.Input[str] type: Resource type

        The **properties** object supports the following:

          * `cer_blob` (`pulumi.Input[str]`) - Raw bytes of .cer file
          * `expiration_date` (`pulumi.Input[str]`) - Certificate expiration date
          * `friendly_name` (`pulumi.Input[str]`) - Friendly name of the certificate
          * `host_names` (`pulumi.Input[list]`) - Host names the certificate applies to
          * `hosting_environment_profile` (`pulumi.Input[dict]`) - Specification for the hosting environment (App Service Environment) to use for the certificate
            * `id` (`pulumi.Input[str]`) - Resource id of the hostingEnvironment (App Service Environment)
            * `name` (`pulumi.Input[str]`) - Name of the hostingEnvironment (App Service Environment) (read only)
            * `type` (`pulumi.Input[str]`) - Resource type of the hostingEnvironment (App Service Environment) (read only)

          * `issue_date` (`pulumi.Input[str]`) - Certificate issue Date
          * `issuer` (`pulumi.Input[str]`) - Certificate issuer
          * `password` (`pulumi.Input[str]`) - Certificate password
          * `pfx_blob` (`pulumi.Input[str]`) - Pfx blob
          * `public_key_hash` (`pulumi.Input[str]`) - Public key hash
          * `self_link` (`pulumi.Input[str]`) - Self link
          * `site_name` (`pulumi.Input[str]`) - App name
          * `subject_name` (`pulumi.Input[str]`) - Subject name of the certificate
          * `thumbprint` (`pulumi.Input[str]`) - Certificate thumbprint
          * `valid` (`pulumi.Input[bool]`) - Is the certificate valid?
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

            __props__['id'] = id
            __props__['kind'] = kind
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
            __props__['type'] = type
        super(Certificate, __self__).__init__(
            'azurerm:web/v20150801:Certificate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing Certificate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Certificate(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
