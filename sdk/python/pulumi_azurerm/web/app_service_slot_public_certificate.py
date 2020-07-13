# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class AppServiceSlotPublicCertificate(pulumi.CustomResource):
    kind: pulumi.Output[str]
    """
    Kind of resource.
    """
    name: pulumi.Output[str]
    """
    Resource Name.
    """
    properties: pulumi.Output[dict]
    """
    PublicCertificate resource specific properties
      * `blob` (`str`) - Public Certificate byte array
      * `public_certificate_location` (`str`) - Public Certificate Location
      * `thumbprint` (`str`) - Certificate Thumbprint
    """
    type: pulumi.Output[str]
    """
    Resource type.
    """
    def __init__(__self__, resource_name, opts=None, kind=None, name=None, properties=None, public_certificate_name=None, resource_group_name=None, slot=None, __props__=None, __name__=None, __opts__=None):
        """
        Public certificate object

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] kind: Kind of resource.
        :param pulumi.Input[str] name: Name of the app.
        :param pulumi.Input[dict] properties: PublicCertificate resource specific properties
        :param pulumi.Input[str] public_certificate_name: Public certificate name.
        :param pulumi.Input[str] resource_group_name: Name of the resource group to which the resource belongs.
        :param pulumi.Input[str] slot: Name of the deployment slot. If a slot is not specified, the API will create a binding for the production slot.

        The **properties** object supports the following:

          * `blob` (`pulumi.Input[str]`) - Public Certificate byte array
          * `public_certificate_location` (`pulumi.Input[str]`) - Public Certificate Location
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

            __props__['kind'] = kind
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['properties'] = properties
            if public_certificate_name is None:
                raise TypeError("Missing required property 'public_certificate_name'")
            __props__['public_certificate_name'] = public_certificate_name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if slot is None:
                raise TypeError("Missing required property 'slot'")
            __props__['slot'] = slot
            __props__['type'] = None
        super(AppServiceSlotPublicCertificate, __self__).__init__(
            'azurerm:web:AppServiceSlotPublicCertificate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing AppServiceSlotPublicCertificate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return AppServiceSlotPublicCertificate(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
