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
    Kind of resource.
    """
    location: pulumi.Output[str]
    """
    Resource Location.
    """
    name: pulumi.Output[str]
    """
    Resource Name.
    """
    properties: pulumi.Output[dict]
    """
    Certificate resource specific properties
      * `canonical_name` (`str`) - CNAME of the certificate to be issued via free certificate
      * `cer_blob` (`str`) - Raw bytes of .cer file
      * `expiration_date` (`str`) - Certificate expiration date.
      * `friendly_name` (`str`) - Friendly name of the certificate.
      * `host_names` (`list`) - Host names the certificate applies to.
      * `hosting_environment_profile` (`dict`) - Specification for the App Service Environment to use for the certificate.
        * `id` (`str`) - Resource ID of the App Service Environment.
        * `name` (`str`) - Name of the App Service Environment.
        * `type` (`str`) - Resource type of the App Service Environment.

      * `issue_date` (`str`) - Certificate issue Date.
      * `issuer` (`str`) - Certificate issuer.
      * `key_vault_id` (`str`) - Key Vault Csm resource Id.
      * `key_vault_secret_name` (`str`) - Key Vault secret name.
      * `key_vault_secret_status` (`str`) - Status of the Key Vault secret.
      * `password` (`str`) - Certificate password.
      * `pfx_blob` (`str`) - Pfx blob.
      * `public_key_hash` (`str`) - Public key hash.
      * `self_link` (`str`) - Self link.
      * `server_farm_id` (`str`) - Resource ID of the associated App Service plan, formatted as: "/subscriptions/{subscriptionID}/resourceGroups/{groupName}/providers/Microsoft.Web/serverfarms/{appServicePlanName}".
      * `site_name` (`str`) - App name.
      * `subject_name` (`str`) - Subject name of the certificate.
      * `thumbprint` (`str`) - Certificate thumbprint.
      * `valid` (`bool`) - Is the certificate valid?.
    """
    tags: pulumi.Output[dict]
    """
    Resource tags.
    """
    type: pulumi.Output[str]
    """
    Resource type.
    """
    def __init__(__self__, resource_name, opts=None, kind=None, location=None, name=None, properties=None, resource_group_name=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        SSL certificate for an app.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] kind: Kind of resource.
        :param pulumi.Input[str] location: Resource Location.
        :param pulumi.Input[str] name: Name of the certificate.
        :param pulumi.Input[dict] properties: Certificate resource specific properties
        :param pulumi.Input[str] resource_group_name: Name of the resource group to which the resource belongs.
        :param pulumi.Input[dict] tags: Resource tags.

        The **properties** object supports the following:

          * `canonical_name` (`pulumi.Input[str]`) - CNAME of the certificate to be issued via free certificate
          * `host_names` (`pulumi.Input[list]`) - Host names the certificate applies to.
          * `key_vault_id` (`pulumi.Input[str]`) - Key Vault Csm resource Id.
          * `key_vault_secret_name` (`pulumi.Input[str]`) - Key Vault secret name.
          * `password` (`pulumi.Input[str]`) - Certificate password.
          * `pfx_blob` (`pulumi.Input[str]`) - Pfx blob.
          * `server_farm_id` (`pulumi.Input[str]`) - Resource ID of the associated App Service plan, formatted as: "/subscriptions/{subscriptionID}/resourceGroups/{groupName}/providers/Microsoft.Web/serverfarms/{appServicePlanName}".
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
            __props__['type'] = None
        super(Certificate, __self__).__init__(
            'azurerm:web/v20190801:Certificate',
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
