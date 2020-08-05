# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class AppServiceCertificateOrder(pulumi.CustomResource):
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
    AppServiceCertificateOrder resource specific properties
      * `app_service_certificate_not_renewable_reasons` (`list`) - Reasons why App Service Certificate is not renewable at the current moment.
      * `auto_renew` (`bool`) - <code>true</code> if the certificate should be automatically renewed when it expires; otherwise, <code>false</code>.
      * `certificates` (`dict`) - State of the Key Vault secret.
      * `csr` (`str`) - Last CSR that was created for this order.
      * `distinguished_name` (`str`) - Certificate distinguished name.
      * `domain_verification_token` (`str`) - Domain verification token.
      * `expiration_time` (`str`) - Certificate expiration time.
      * `intermediate` (`dict`) - Intermediate certificate.
        * `id` (`str`) - Resource Id
        * `kind` (`str`) - Kind of resource
        * `location` (`str`) - Resource Location
        * `name` (`str`) - Resource Name
        * `properties` (`dict`)
          * `issuer` (`str`) - Issuer
          * `not_after` (`str`) - Valid to
          * `not_before` (`str`) - Valid from
          * `raw_data` (`str`) - Raw certificate data
          * `serial_number` (`str`) - Serial Number
          * `signature_algorithm` (`str`) - Signature Algorithm
          * `subject` (`str`) - Subject
          * `thumbprint` (`str`) - Thumbprint
          * `version` (`float`) - Version

        * `tags` (`dict`) - Resource tags
        * `type` (`str`) - Resource type

      * `is_private_key_external` (`bool`) - <code>true</code> if private key is external; otherwise, <code>false</code>.
      * `key_size` (`float`) - Certificate key size.
      * `last_certificate_issuance_time` (`str`) - Certificate last issuance time.
      * `next_auto_renewal_time_stamp` (`str`) - Time stamp when the certificate would be auto renewed next
      * `product_type` (`str`) - Certificate product type.
      * `provisioning_state` (`str`) - Status of certificate order.
      * `root` (`dict`) - Root certificate.
      * `serial_number` (`str`) - Current serial number of the certificate.
      * `signed_certificate` (`dict`) - Signed certificate.
      * `status` (`str`) - Current order status.
      * `validity_in_years` (`float`) - Duration in years (must be between 1 and 3).
    """
    tags: pulumi.Output[dict]
    """
    Resource tags.
    """
    type: pulumi.Output[str]
    """
    Resource type.
    """
    def __init__(__self__, resource_name, opts=None, auto_renew=None, certificates=None, csr=None, distinguished_name=None, key_size=None, kind=None, location=None, name=None, product_type=None, resource_group_name=None, tags=None, validity_in_years=None, __props__=None, __name__=None, __opts__=None):
        """
        SSL certificate purchase order.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_renew: <code>true</code> if the certificate should be automatically renewed when it expires; otherwise, <code>false</code>.
        :param pulumi.Input[dict] certificates: State of the Key Vault secret.
        :param pulumi.Input[str] csr: Last CSR that was created for this order.
        :param pulumi.Input[str] distinguished_name: Certificate distinguished name.
        :param pulumi.Input[float] key_size: Certificate key size.
        :param pulumi.Input[str] kind: Kind of resource.
        :param pulumi.Input[str] location: Resource Location.
        :param pulumi.Input[str] name: Name of the certificate order.
        :param pulumi.Input[str] product_type: Certificate product type.
        :param pulumi.Input[str] resource_group_name: Name of the resource group to which the resource belongs.
        :param pulumi.Input[dict] tags: Resource tags.
        :param pulumi.Input[float] validity_in_years: Duration in years (must be between 1 and 3).
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

            __props__['auto_renew'] = auto_renew
            __props__['certificates'] = certificates
            __props__['csr'] = csr
            __props__['distinguished_name'] = distinguished_name
            __props__['key_size'] = key_size
            __props__['kind'] = kind
            if location is None:
                raise TypeError("Missing required property 'location'")
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if product_type is None:
                raise TypeError("Missing required property 'product_type'")
            __props__['product_type'] = product_type
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['validity_in_years'] = validity_in_years
            __props__['properties'] = None
            __props__['type'] = None
        super(AppServiceCertificateOrder, __self__).__init__(
            'azurerm:certificateregistration/v20150801:AppServiceCertificateOrder',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing AppServiceCertificateOrder resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return AppServiceCertificateOrder(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
