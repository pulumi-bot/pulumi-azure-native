# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = ['Certificate']


class Certificate(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 automation_account_name: Optional[pulumi.Input[str]] = None,
                 base64_value: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 is_exportable: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 thumbprint: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Definition of the certificate.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] automation_account_name: The name of the automation account.
        :param pulumi.Input[str] base64_value: Gets or sets the base64 encoded value of the certificate.
        :param pulumi.Input[str] description: Gets or sets the description of the certificate.
        :param pulumi.Input[bool] is_exportable: Gets or sets the is exportable flag of the certificate.
        :param pulumi.Input[str] name: The parameters supplied to the create or update certificate operation.
        :param pulumi.Input[str] resource_group_name: Name of an Azure Resource group.
        :param pulumi.Input[str] thumbprint: Gets or sets the thumbprint of the certificate.
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

            if automation_account_name is None:
                raise TypeError("Missing required property 'automation_account_name'")
            __props__['automation_account_name'] = automation_account_name
            if base64_value is None:
                raise TypeError("Missing required property 'base64_value'")
            __props__['base64_value'] = base64_value
            __props__['description'] = description
            __props__['is_exportable'] = is_exportable
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['thumbprint'] = thumbprint
            __props__['creation_time'] = None
            __props__['expiry_time'] = None
            __props__['last_modified_time'] = None
            __props__['type'] = None
        super(Certificate, __self__).__init__(
            'azurerm:automation/v20151031:Certificate',
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

        return Certificate(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> str:
        """
        Gets the creation time.
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        Gets or sets the description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="expiryTime")
    def expiry_time(self) -> str:
        """
        Gets the expiry time of the certificate.
        """
        return pulumi.get(self, "expiry_time")

    @property
    @pulumi.getter(name="isExportable")
    def is_exportable(self) -> bool:
        """
        Gets the is exportable flag of the certificate.
        """
        return pulumi.get(self, "is_exportable")

    @property
    @pulumi.getter(name="lastModifiedTime")
    def last_modified_time(self) -> str:
        """
        Gets the last modified time.
        """
        return pulumi.get(self, "last_modified_time")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def thumbprint(self) -> str:
        """
        Gets the thumbprint of the certificate.
        """
        return pulumi.get(self, "thumbprint")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the resource.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

