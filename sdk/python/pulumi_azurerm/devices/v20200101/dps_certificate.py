# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables
from . import outputs

__all__ = ['DpsCertificate']


class DpsCertificate(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 certificate: Optional[pulumi.Input[str]] = None,
                 certificate_name: Optional[pulumi.Input[str]] = None,
                 provisioning_service_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        The X509 Certificate.

        ## Example Usage
        ### DPSCreateOrUpdateCertificate

        ```python
        import pulumi
        import pulumi_azurerm as azurerm

        dps_certificate = azurerm.devices.v20200101.DpsCertificate("dpsCertificate",
            certificate="############################################",
            certificate_name="cert",
            provisioning_service_name="myFirstProvisioningService",
            resource_group_name="myResourceGroup")

        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] certificate: Base-64 representation of the X509 leaf certificate .cer file or just .pem file content.
        :param pulumi.Input[str] certificate_name: The name of the certificate create or update.
        :param pulumi.Input[str] provisioning_service_name: The name of the provisioning service.
        :param pulumi.Input[str] resource_group_name: Resource group identifier.
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
            if certificate_name is None:
                raise TypeError("Missing required property 'certificate_name'")
            __props__['certificate_name'] = certificate_name
            if provisioning_service_name is None:
                raise TypeError("Missing required property 'provisioning_service_name'")
            __props__['provisioning_service_name'] = provisioning_service_name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['etag'] = None
            __props__['name'] = None
            __props__['properties'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:devices/latest:DpsCertificate"), pulumi.Alias(type_="azurerm:devices/v20170821preview:DpsCertificate"), pulumi.Alias(type_="azurerm:devices/v20171115:DpsCertificate"), pulumi.Alias(type_="azurerm:devices/v20180122:DpsCertificate")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(DpsCertificate, __self__).__init__(
            'azurerm:devices/v20200101:DpsCertificate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'DpsCertificate':
        """
        Get an existing DpsCertificate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return DpsCertificate(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        The entity tag.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the certificate.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def properties(self) -> pulumi.Output['outputs.CertificatePropertiesResponse']:
        """
        properties of a certificate
        """
        return pulumi.get(self, "properties")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The resource type.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

