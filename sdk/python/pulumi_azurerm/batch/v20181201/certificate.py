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
    The ETag of the resource, used for concurrency statements.
    """
    name: pulumi.Output[str]
    """
    The name of the resource.
    """
    properties: pulumi.Output[dict]
    """
    The properties associated with the certificate.
      * `delete_certificate_error` (`dict`) - This is only returned when the certificate provisioningState is 'Failed'.
        * `code` (`str`) - An identifier for the error. Codes are invariant and are intended to be consumed programmatically.
        * `details` (`list`) - A list of additional details about the error.
        * `message` (`str`) - A message describing the error, intended to be suitable for display in a user interface.
        * `target` (`str`) - The target of the particular error. For example, the name of the property in error.

      * `format` (`str`) - The format of the certificate - either Pfx or Cer. If omitted, the default is Pfx.
      * `previous_provisioning_state` (`str`) - The previous provisioned state of the resource
      * `previous_provisioning_state_transition_time` (`str`)
      * `provisioning_state` (`str`)
      * `provisioning_state_transition_time` (`str`)
      * `public_data` (`str`) - The public key of the certificate.
      * `thumbprint` (`str`) - This must match the thumbprint from the name.
      * `thumbprint_algorithm` (`str`) - This must match the first portion of the certificate name. Currently required to be 'SHA1'.
    """
    type: pulumi.Output[str]
    """
    The type of the resource.
    """
    def __init__(__self__, resource_name, opts=None, account_name=None, name=None, properties=None, resource_group_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Contains information about a certificate.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: The name of the Batch account.
        :param pulumi.Input[str] name: The identifier for the certificate. This must be made up of algorithm and thumbprint separated by a dash, and must match the certificate data in the request. For example SHA1-a3d1c5.
        :param pulumi.Input[dict] properties: The properties associated with the certificate.
        :param pulumi.Input[str] resource_group_name: The name of the resource group that contains the Batch account.

        The **properties** object supports the following:

          * `data` (`pulumi.Input[str]`) - The maximum size is 10KB.
          * `format` (`pulumi.Input[str]`) - The format of the certificate - either Pfx or Cer. If omitted, the default is Pfx.
          * `password` (`pulumi.Input[str]`) - This is required if the certificate format is pfx and must be omitted if the certificate format is cer.
          * `thumbprint` (`pulumi.Input[str]`) - This must match the thumbprint from the name.
          * `thumbprint_algorithm` (`pulumi.Input[str]`) - This must match the first portion of the certificate name. Currently required to be 'SHA1'.
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
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['etag'] = None
            __props__['type'] = None
        super(Certificate, __self__).__init__(
            'azurerm:batch/v20181201:Certificate',
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
