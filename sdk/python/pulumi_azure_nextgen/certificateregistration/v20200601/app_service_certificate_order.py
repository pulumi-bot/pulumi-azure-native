# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from . import outputs
from ._enums import *
from ._inputs import *

__all__ = ['AppServiceCertificateOrder']


class AppServiceCertificateOrder(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_renew: Optional[pulumi.Input[bool]] = None,
                 certificate_order_name: Optional[pulumi.Input[str]] = None,
                 certificates: Optional[pulumi.Input[Mapping[str, pulumi.Input[pulumi.InputType['AppServiceCertificateArgs']]]]] = None,
                 csr: Optional[pulumi.Input[str]] = None,
                 distinguished_name: Optional[pulumi.Input[str]] = None,
                 key_size: Optional[pulumi.Input[int]] = None,
                 kind: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 product_type: Optional[pulumi.Input['CertificateProductType']] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 validity_in_years: Optional[pulumi.Input[int]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        SSL certificate purchase order.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_renew: <code>true</code> if the certificate should be automatically renewed when it expires; otherwise, <code>false</code>.
        :param pulumi.Input[str] certificate_order_name: Name of the certificate order.
        :param pulumi.Input[Mapping[str, pulumi.Input[pulumi.InputType['AppServiceCertificateArgs']]]] certificates: State of the Key Vault secret.
        :param pulumi.Input[str] csr: Last CSR that was created for this order.
        :param pulumi.Input[str] distinguished_name: Certificate distinguished name.
        :param pulumi.Input[int] key_size: Certificate key size.
        :param pulumi.Input[str] kind: Kind of resource.
        :param pulumi.Input[str] location: Resource Location.
        :param pulumi.Input['CertificateProductType'] product_type: Certificate product type.
        :param pulumi.Input[str] resource_group_name: Name of the resource group to which the resource belongs.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags.
        :param pulumi.Input[int] validity_in_years: Duration in years (must be between 1 and 3).
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

            if auto_renew is None:
                auto_renew = True
            __props__['auto_renew'] = auto_renew
            if certificate_order_name is None and not opts.urn:
                raise TypeError("Missing required property 'certificate_order_name'")
            __props__['certificate_order_name'] = certificate_order_name
            __props__['certificates'] = certificates
            __props__['csr'] = csr
            __props__['distinguished_name'] = distinguished_name
            if key_size is None:
                key_size = 2048
            __props__['key_size'] = key_size
            __props__['kind'] = kind
            __props__['location'] = location
            if product_type is None and not opts.urn:
                raise TypeError("Missing required property 'product_type'")
            __props__['product_type'] = product_type
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            if validity_in_years is None:
                validity_in_years = 1
            __props__['validity_in_years'] = validity_in_years
            __props__['app_service_certificate_not_renewable_reasons'] = None
            __props__['domain_verification_token'] = None
            __props__['expiration_time'] = None
            __props__['intermediate'] = None
            __props__['is_private_key_external'] = None
            __props__['last_certificate_issuance_time'] = None
            __props__['name'] = None
            __props__['next_auto_renewal_time_stamp'] = None
            __props__['provisioning_state'] = None
            __props__['root'] = None
            __props__['serial_number'] = None
            __props__['signed_certificate'] = None
            __props__['status'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:certificateregistration/latest:AppServiceCertificateOrder"), pulumi.Alias(type_="azure-nextgen:certificateregistration/v20150801:AppServiceCertificateOrder"), pulumi.Alias(type_="azure-nextgen:certificateregistration/v20180201:AppServiceCertificateOrder"), pulumi.Alias(type_="azure-nextgen:certificateregistration/v20190801:AppServiceCertificateOrder"), pulumi.Alias(type_="azure-nextgen:certificateregistration/v20200901:AppServiceCertificateOrder")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(AppServiceCertificateOrder, __self__).__init__(
            'azure-nextgen:certificateregistration/v20200601:AppServiceCertificateOrder',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'AppServiceCertificateOrder':
        """
        Get an existing AppServiceCertificateOrder resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return AppServiceCertificateOrder(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="appServiceCertificateNotRenewableReasons")
    def app_service_certificate_not_renewable_reasons(self) -> pulumi.Output[Sequence[str]]:
        """
        Reasons why App Service Certificate is not renewable at the current moment.
        """
        return pulumi.get(self, "app_service_certificate_not_renewable_reasons")

    @property
    @pulumi.getter(name="autoRenew")
    def auto_renew(self) -> pulumi.Output[Optional[bool]]:
        """
        <code>true</code> if the certificate should be automatically renewed when it expires; otherwise, <code>false</code>.
        """
        return pulumi.get(self, "auto_renew")

    @property
    @pulumi.getter
    def certificates(self) -> pulumi.Output[Optional[Mapping[str, 'outputs.AppServiceCertificateResponse']]]:
        """
        State of the Key Vault secret.
        """
        return pulumi.get(self, "certificates")

    @property
    @pulumi.getter
    def csr(self) -> pulumi.Output[Optional[str]]:
        """
        Last CSR that was created for this order.
        """
        return pulumi.get(self, "csr")

    @property
    @pulumi.getter(name="distinguishedName")
    def distinguished_name(self) -> pulumi.Output[Optional[str]]:
        """
        Certificate distinguished name.
        """
        return pulumi.get(self, "distinguished_name")

    @property
    @pulumi.getter(name="domainVerificationToken")
    def domain_verification_token(self) -> pulumi.Output[str]:
        """
        Domain verification token.
        """
        return pulumi.get(self, "domain_verification_token")

    @property
    @pulumi.getter(name="expirationTime")
    def expiration_time(self) -> pulumi.Output[str]:
        """
        Certificate expiration time.
        """
        return pulumi.get(self, "expiration_time")

    @property
    @pulumi.getter
    def intermediate(self) -> pulumi.Output['outputs.CertificateDetailsResponse']:
        """
        Intermediate certificate.
        """
        return pulumi.get(self, "intermediate")

    @property
    @pulumi.getter(name="isPrivateKeyExternal")
    def is_private_key_external(self) -> pulumi.Output[bool]:
        """
        <code>true</code> if private key is external; otherwise, <code>false</code>.
        """
        return pulumi.get(self, "is_private_key_external")

    @property
    @pulumi.getter(name="keySize")
    def key_size(self) -> pulumi.Output[Optional[int]]:
        """
        Certificate key size.
        """
        return pulumi.get(self, "key_size")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[Optional[str]]:
        """
        Kind of resource.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="lastCertificateIssuanceTime")
    def last_certificate_issuance_time(self) -> pulumi.Output[str]:
        """
        Certificate last issuance time.
        """
        return pulumi.get(self, "last_certificate_issuance_time")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Resource Location.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource Name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nextAutoRenewalTimeStamp")
    def next_auto_renewal_time_stamp(self) -> pulumi.Output[str]:
        """
        Time stamp when the certificate would be auto renewed next
        """
        return pulumi.get(self, "next_auto_renewal_time_stamp")

    @property
    @pulumi.getter(name="productType")
    def product_type(self) -> pulumi.Output[str]:
        """
        Certificate product type.
        """
        return pulumi.get(self, "product_type")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        Status of certificate order.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def root(self) -> pulumi.Output['outputs.CertificateDetailsResponse']:
        """
        Root certificate.
        """
        return pulumi.get(self, "root")

    @property
    @pulumi.getter(name="serialNumber")
    def serial_number(self) -> pulumi.Output[str]:
        """
        Current serial number of the certificate.
        """
        return pulumi.get(self, "serial_number")

    @property
    @pulumi.getter(name="signedCertificate")
    def signed_certificate(self) -> pulumi.Output['outputs.CertificateDetailsResponse']:
        """
        Signed certificate.
        """
        return pulumi.get(self, "signed_certificate")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Current order status.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Resource type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="validityInYears")
    def validity_in_years(self) -> pulumi.Output[Optional[int]]:
        """
        Duration in years (must be between 1 and 3).
        """
        return pulumi.get(self, "validity_in_years")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

