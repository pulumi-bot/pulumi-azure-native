# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Workflow']


class Workflow(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 definition: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 integration_account: Optional[pulumi.Input[pulumi.InputType['ResourceReferenceArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[pulumi.InputType['WorkflowParameterArgs']]]]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 sku: Optional[pulumi.Input[pulumi.InputType['SkuArgs']]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 workflow_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        The workflow type.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, Any]] definition: The definition. See [Schema reference for Workflow Definition Language in Azure Logic Apps](https://aka.ms/logic-apps-workflow-definition-language).
        :param pulumi.Input[pulumi.InputType['ResourceReferenceArgs']] integration_account: The integration account.
        :param pulumi.Input[str] location: The resource location.
        :param pulumi.Input[Mapping[str, pulumi.Input[pulumi.InputType['WorkflowParameterArgs']]]] parameters: The parameters.
        :param pulumi.Input[str] resource_group_name: The resource group name.
        :param pulumi.Input[pulumi.InputType['SkuArgs']] sku: The sku.
        :param pulumi.Input[str] state: The state.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: The resource tags.
        :param pulumi.Input[str] workflow_name: The workflow name.
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

            __props__['definition'] = definition
            __props__['integration_account'] = integration_account
            __props__['location'] = location
            __props__['parameters'] = parameters
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['sku'] = sku
            __props__['state'] = state
            __props__['tags'] = tags
            if workflow_name is None:
                raise TypeError("Missing required property 'workflow_name'")
            __props__['workflow_name'] = workflow_name
            __props__['access_endpoint'] = None
            __props__['changed_time'] = None
            __props__['created_time'] = None
            __props__['name'] = None
            __props__['provisioning_state'] = None
            __props__['type'] = None
            __props__['version'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:logic/latest:Workflow"), pulumi.Alias(type_="azurerm:logic/v20150201preview:Workflow"), pulumi.Alias(type_="azurerm:logic/v20160601:Workflow"), pulumi.Alias(type_="azurerm:logic/v20190501:Workflow")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Workflow, __self__).__init__(
            'azurerm:logic/v20180701preview:Workflow',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Workflow':
        """
        Get an existing Workflow resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Workflow(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessEndpoint")
    def access_endpoint(self) -> pulumi.Output[str]:
        """
        Gets the access endpoint.
        """
        return pulumi.get(self, "access_endpoint")

    @property
    @pulumi.getter(name="changedTime")
    def changed_time(self) -> pulumi.Output[str]:
        """
        Gets the changed time.
        """
        return pulumi.get(self, "changed_time")

    @property
    @pulumi.getter(name="createdTime")
    def created_time(self) -> pulumi.Output[str]:
        """
        Gets the created time.
        """
        return pulumi.get(self, "created_time")

    @property
    @pulumi.getter
    def definition(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        The definition. See [Schema reference for Workflow Definition Language in Azure Logic Apps](https://aka.ms/logic-apps-workflow-definition-language).
        """
        return pulumi.get(self, "definition")

    @property
    @pulumi.getter(name="integrationAccount")
    def integration_account(self) -> pulumi.Output[Optional['outputs.ResourceReferenceResponse']]:
        """
        The integration account.
        """
        return pulumi.get(self, "integration_account")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[Optional[str]]:
        """
        The resource location.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Gets the resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def parameters(self) -> pulumi.Output[Optional[Mapping[str, 'outputs.WorkflowParameterResponse']]]:
        """
        The parameters.
        """
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        Gets the provisioning state.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def sku(self) -> pulumi.Output[Optional['outputs.SkuResponse']]:
        """
        The sku.
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[Optional[str]]:
        """
        The state.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        The resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Gets the resource type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[str]:
        """
        Gets the version.
        """
        return pulumi.get(self, "version")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

