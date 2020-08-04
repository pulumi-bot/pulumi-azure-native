# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class Certificate(pulumi.CustomResource):
    etag: pulumi.Output[str]
    """
    The entity tag.
    """
    name: pulumi.Output[str]
    """
    The name of the certificate.
    """
    properties: pulumi.Output[dict]
    """
    The description of an X509 CA Certificate.
      * `certificate` (`str`) - The certificate content
      * `created` (`str`) - The certificate's create date and time.
      * `expiry` (`str`) - The certificate's expiration date and time.
      * `is_verified` (`bool`) - Determines whether certificate has been verified.
      * `subject` (`str`) - The certificate's subject name.
      * `thumbprint` (`str`) - The certificate's thumbprint.
      * `updated` (`str`) - The certificate's last update date and time.
    """
    type: pulumi.Output[str]
    """
    The resource type.
    """
    def __init__(__self__, resource_name, opts=None, certificate=None, name=None, resource_group_name=None, resource_name_=None, __props__=None, __name__=None, __opts__=None):
        """
        The X509 Certificate.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] certificate: base-64 representation of the X509 leaf certificate .cer file or just .pem file content.
        :param pulumi.Input[str] name: The name of the certificate
        :param pulumi.Input[str] resource_group_name: The name of the resource group that contains the IoT hub.
        :param pulumi.Input[str] resource_name_: The name of the IoT hub.
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

            __props__['certificate'] = certificate
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if resource_name_ is None:
                raise TypeError("Missing required property 'resource_name_'")
            __props__['resource_name'] = resource_name_
            __props__['etag'] = None
            __props__['properties'] = None
            __props__['type'] = None
        super(Certificate, __self__).__init__(
            'azurerm:devices/v20200301:Certificate',
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
