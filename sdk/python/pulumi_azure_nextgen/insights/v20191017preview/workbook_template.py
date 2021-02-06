# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['WorkbookTemplate']


class WorkbookTemplate(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 author: Optional[pulumi.Input[str]] = None,
                 galleries: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WorkbookTemplateGalleryArgs']]]]] = None,
                 localized: Optional[pulumi.Input[Mapping[str, pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WorkbookTemplateLocalizedGalleryArgs']]]]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 resource_name_: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 template_data: Optional[Any] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        An Application Insights workbook template definition.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] author: Information about the author of the workbook template.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WorkbookTemplateGalleryArgs']]]] galleries: Workbook galleries supported by the template.
        :param pulumi.Input[Mapping[str, pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WorkbookTemplateLocalizedGalleryArgs']]]]]] localized: Key value pair of localized gallery. Each key is the locale code of languages supported by the Azure portal.
        :param pulumi.Input[str] location: Resource location
        :param pulumi.Input[int] priority: Priority of the template. Determines which template to open when a workbook gallery is opened in viewer mode.
        :param pulumi.Input[str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[str] resource_name_: The name of the Application Insights component resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags
        :param Any template_data: Valid JSON object containing workbook template payload.
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

            __props__['author'] = author
            if galleries is None and not opts.urn:
                raise TypeError("Missing required property 'galleries'")
            __props__['galleries'] = galleries
            __props__['localized'] = localized
            __props__['location'] = location
            __props__['priority'] = priority
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if resource_name_ is None and not opts.urn:
                raise TypeError("Missing required property 'resource_name_'")
            __props__['resource_name'] = resource_name_
            __props__['tags'] = tags
            if template_data is None and not opts.urn:
                raise TypeError("Missing required property 'template_data'")
            __props__['template_data'] = template_data
            __props__['name'] = None
            __props__['type'] = None
        super(WorkbookTemplate, __self__).__init__(
            'azure-nextgen:insights/v20191017preview:WorkbookTemplate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'WorkbookTemplate':
        """
        Get an existing WorkbookTemplate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return WorkbookTemplate(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def author(self) -> pulumi.Output[Optional[str]]:
        """
        Information about the author of the workbook template.
        """
        return pulumi.get(self, "author")

    @property
    @pulumi.getter
    def galleries(self) -> pulumi.Output[Sequence['outputs.WorkbookTemplateGalleryResponse']]:
        """
        Workbook galleries supported by the template.
        """
        return pulumi.get(self, "galleries")

    @property
    @pulumi.getter
    def localized(self) -> pulumi.Output[Optional[Mapping[str, Sequence['outputs.WorkbookTemplateLocalizedGalleryResponse']]]]:
        """
        Key value pair of localized gallery. Each key is the locale code of languages supported by the Azure portal.
        """
        return pulumi.get(self, "localized")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        """
        Resource location
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Azure resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def priority(self) -> pulumi.Output[Optional[int]]:
        """
        Priority of the template. Determines which template to open when a workbook gallery is opened in viewer mode.
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Resource tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="templateData")
    def template_data(self) -> pulumi.Output[Any]:
        """
        Valid JSON object containing workbook template payload.
        """
        return pulumi.get(self, "template_data")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Azure resource type
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

