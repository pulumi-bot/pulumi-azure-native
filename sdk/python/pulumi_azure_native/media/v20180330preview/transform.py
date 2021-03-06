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

__all__ = ['Transform']


class Transform(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 outputs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TransformOutputArgs']]]]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 transform_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        A Transform encapsulates the rules or instructions for generating desired outputs from input media, such as by transcoding or by extracting insights. After the Transform is created, it can be applied to input media by creating Jobs.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: The Media Services account name.
        :param pulumi.Input[str] description: An optional verbose description of the Transform.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TransformOutputArgs']]]] outputs: An array of one or more TransformOutputs that the Transform should generate.
        :param pulumi.Input[str] resource_group_name: The name of the resource group within the Azure subscription.
        :param pulumi.Input[str] transform_name: The Transform name.
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

            if account_name is None and not opts.urn:
                raise TypeError("Missing required property 'account_name'")
            __props__['account_name'] = account_name
            __props__['description'] = description
            if outputs is None and not opts.urn:
                raise TypeError("Missing required property 'outputs'")
            __props__['outputs'] = outputs
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['transform_name'] = transform_name
            __props__['created'] = None
            __props__['last_modified'] = None
            __props__['name'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:media/v20180330preview:Transform"), pulumi.Alias(type_="azure-native:media:Transform"), pulumi.Alias(type_="azure-nextgen:media:Transform"), pulumi.Alias(type_="azure-native:media/latest:Transform"), pulumi.Alias(type_="azure-nextgen:media/latest:Transform"), pulumi.Alias(type_="azure-native:media/v20180601preview:Transform"), pulumi.Alias(type_="azure-nextgen:media/v20180601preview:Transform"), pulumi.Alias(type_="azure-native:media/v20180701:Transform"), pulumi.Alias(type_="azure-nextgen:media/v20180701:Transform"), pulumi.Alias(type_="azure-native:media/v20200501:Transform"), pulumi.Alias(type_="azure-nextgen:media/v20200501:Transform")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Transform, __self__).__init__(
            'azure-native:media/v20180330preview:Transform',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Transform':
        """
        Get an existing Transform resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["created"] = None
        __props__["description"] = None
        __props__["last_modified"] = None
        __props__["name"] = None
        __props__["outputs"] = None
        __props__["type"] = None
        return Transform(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def created(self) -> pulumi.Output[str]:
        """
        The UTC date and time when the Transform was created, in 'YYYY-MM-DDThh:mm:ssZ' format.
        """
        return pulumi.get(self, "created")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        An optional verbose description of the Transform.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="lastModified")
    def last_modified(self) -> pulumi.Output[str]:
        """
        The UTC date and time when the Transform was last updated, in 'YYYY-MM-DDThh:mm:ssZ' format.
        """
        return pulumi.get(self, "last_modified")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def outputs(self) -> pulumi.Output[Sequence['outputs.TransformOutputResponse']]:
        """
        An array of one or more TransformOutputs that the Transform should generate.
        """
        return pulumi.get(self, "outputs")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of the resource.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

