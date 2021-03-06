# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from ._enums import *

__all__ = ['WorkspaceConnection']


class WorkspaceConnection(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_type: Optional[pulumi.Input[str]] = None,
                 category: Optional[pulumi.Input[str]] = None,
                 connection_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 target: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None,
                 value_format: Optional[pulumi.Input[Union[str, 'ValueFormat']]] = None,
                 workspace_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Workspace connection.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auth_type: Authorization type of the workspace connection.
        :param pulumi.Input[str] category: Category of the workspace connection.
        :param pulumi.Input[str] connection_name: Friendly name of the workspace connection
        :param pulumi.Input[str] name: Friendly name of the workspace connection
        :param pulumi.Input[str] resource_group_name: Name of the resource group in which workspace is located.
        :param pulumi.Input[str] target: Target of the workspace connection.
        :param pulumi.Input[str] value: Value details of the workspace connection.
        :param pulumi.Input[Union[str, 'ValueFormat']] value_format: format for the workspace connection value
        :param pulumi.Input[str] workspace_name: Name of Azure Machine Learning workspace.
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

            __props__['auth_type'] = auth_type
            __props__['category'] = category
            __props__['connection_name'] = connection_name
            __props__['name'] = name
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['target'] = target
            __props__['value'] = value
            __props__['value_format'] = value_format
            if workspace_name is None and not opts.urn:
                raise TypeError("Missing required property 'workspace_name'")
            __props__['workspace_name'] = workspace_name
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:machinelearningservices/v20210101:WorkspaceConnection"), pulumi.Alias(type_="azure-native:machinelearningservices:WorkspaceConnection"), pulumi.Alias(type_="azure-nextgen:machinelearningservices:WorkspaceConnection"), pulumi.Alias(type_="azure-native:machinelearningservices/latest:WorkspaceConnection"), pulumi.Alias(type_="azure-nextgen:machinelearningservices/latest:WorkspaceConnection"), pulumi.Alias(type_="azure-native:machinelearningservices/v20200601:WorkspaceConnection"), pulumi.Alias(type_="azure-nextgen:machinelearningservices/v20200601:WorkspaceConnection"), pulumi.Alias(type_="azure-native:machinelearningservices/v20200801:WorkspaceConnection"), pulumi.Alias(type_="azure-nextgen:machinelearningservices/v20200801:WorkspaceConnection"), pulumi.Alias(type_="azure-native:machinelearningservices/v20200901preview:WorkspaceConnection"), pulumi.Alias(type_="azure-nextgen:machinelearningservices/v20200901preview:WorkspaceConnection")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(WorkspaceConnection, __self__).__init__(
            'azure-native:machinelearningservices/v20210101:WorkspaceConnection',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'WorkspaceConnection':
        """
        Get an existing WorkspaceConnection resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["auth_type"] = None
        __props__["category"] = None
        __props__["name"] = None
        __props__["target"] = None
        __props__["type"] = None
        __props__["value"] = None
        __props__["value_format"] = None
        return WorkspaceConnection(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="authType")
    def auth_type(self) -> pulumi.Output[Optional[str]]:
        """
        Authorization type of the workspace connection.
        """
        return pulumi.get(self, "auth_type")

    @property
    @pulumi.getter
    def category(self) -> pulumi.Output[Optional[str]]:
        """
        Category of the workspace connection.
        """
        return pulumi.get(self, "category")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Friendly name of the workspace connection.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def target(self) -> pulumi.Output[Optional[str]]:
        """
        Target of the workspace connection.
        """
        return pulumi.get(self, "target")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Resource type of workspace connection.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def value(self) -> pulumi.Output[Optional[str]]:
        """
        Value details of the workspace connection.
        """
        return pulumi.get(self, "value")

    @property
    @pulumi.getter(name="valueFormat")
    def value_format(self) -> pulumi.Output[Optional[str]]:
        """
        format for the workspace connection value
        """
        return pulumi.get(self, "value_format")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

