# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from . import outputs

__all__ = ['Certificate']


class Certificate(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 certificate: Optional[pulumi.Input[str]] = None,
                 certificate_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 resource_name_: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        The X509 Certificate.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] certificate: base-64 representation of the X509 leaf certificate .cer file or just .pem file content.
        :param pulumi.Input[str] certificate_name: The name of the certificate
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
            __props__['certificate_name'] = certificate_name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if resource_name_ is None and not opts.urn:
                raise TypeError("Missing required property 'resource_name_'")
            __props__['resource_name'] = resource_name_
            __props__['etag'] = None
            __props__['name'] = None
            __props__['properties'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:devices/v20181201preview:Certificate"), pulumi.Alias(type_="azure-native:devices:Certificate"), pulumi.Alias(type_="azure-nextgen:devices:Certificate"), pulumi.Alias(type_="azure-native:devices/latest:Certificate"), pulumi.Alias(type_="azure-nextgen:devices/latest:Certificate"), pulumi.Alias(type_="azure-native:devices/v20170701:Certificate"), pulumi.Alias(type_="azure-nextgen:devices/v20170701:Certificate"), pulumi.Alias(type_="azure-native:devices/v20180122:Certificate"), pulumi.Alias(type_="azure-nextgen:devices/v20180122:Certificate"), pulumi.Alias(type_="azure-native:devices/v20180401:Certificate"), pulumi.Alias(type_="azure-nextgen:devices/v20180401:Certificate"), pulumi.Alias(type_="azure-native:devices/v20190322:Certificate"), pulumi.Alias(type_="azure-nextgen:devices/v20190322:Certificate"), pulumi.Alias(type_="azure-native:devices/v20190322preview:Certificate"), pulumi.Alias(type_="azure-nextgen:devices/v20190322preview:Certificate"), pulumi.Alias(type_="azure-native:devices/v20190701preview:Certificate"), pulumi.Alias(type_="azure-nextgen:devices/v20190701preview:Certificate"), pulumi.Alias(type_="azure-native:devices/v20191104:Certificate"), pulumi.Alias(type_="azure-nextgen:devices/v20191104:Certificate"), pulumi.Alias(type_="azure-native:devices/v20200301:Certificate"), pulumi.Alias(type_="azure-nextgen:devices/v20200301:Certificate"), pulumi.Alias(type_="azure-native:devices/v20200401:Certificate"), pulumi.Alias(type_="azure-nextgen:devices/v20200401:Certificate"), pulumi.Alias(type_="azure-native:devices/v20200615:Certificate"), pulumi.Alias(type_="azure-nextgen:devices/v20200615:Certificate"), pulumi.Alias(type_="azure-native:devices/v20200710preview:Certificate"), pulumi.Alias(type_="azure-nextgen:devices/v20200710preview:Certificate"), pulumi.Alias(type_="azure-native:devices/v20200801:Certificate"), pulumi.Alias(type_="azure-nextgen:devices/v20200801:Certificate"), pulumi.Alias(type_="azure-native:devices/v20200831:Certificate"), pulumi.Alias(type_="azure-nextgen:devices/v20200831:Certificate"), pulumi.Alias(type_="azure-native:devices/v20200831preview:Certificate"), pulumi.Alias(type_="azure-nextgen:devices/v20200831preview:Certificate")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Certificate, __self__).__init__(
            'azure-native:devices/v20181201preview:Certificate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Certificate':
        """
        Get an existing Certificate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["etag"] = None
        __props__["name"] = None
        __props__["properties"] = None
        __props__["type"] = None
        return Certificate(resource_name, opts=opts, __props__=__props__)

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
        The description of an X509 CA Certificate.
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

