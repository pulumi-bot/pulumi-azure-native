# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class Workflow(pulumi.CustomResource):
    location: pulumi.Output[str]
    """
    The resource location.
    """
    name: pulumi.Output[str]
    """
    Gets the resource name.
    """
    properties: pulumi.Output[dict]
    """
    The workflow properties.
      * `access_control` (`dict`) - The access control configuration.
        * `actions` (`dict`) - The access control configuration for workflow actions.
          * `allowed_caller_ip_addresses` (`list`) - The allowed caller IP address ranges.
            * `address_range` (`str`) - The IP address range.

          * `open_authentication_policies` (`dict`) - The authentication policies for workflow.
            * `policies` (`dict`) - Open authentication policies.

        * `contents` (`dict`) - The access control configuration for accessing workflow run contents.
        * `triggers` (`dict`) - The access control configuration for invoking workflow triggers.
        * `workflow_management` (`dict`) - The access control configuration for workflow management.

      * `access_endpoint` (`str`) - Gets the access endpoint.
      * `changed_time` (`str`) - Gets the changed time.
      * `created_time` (`str`) - Gets the created time.
      * `definition` (`dict`) - The definition.
      * `endpoints_configuration` (`dict`) - The endpoints configuration.
        * `connector` (`dict`) - The connector endpoints.
          * `access_endpoint_ip_addresses` (`list`) - The access endpoint ip address.
            * `address` (`str`) - The address.

          * `outgoing_ip_addresses` (`list`) - The outgoing ip address.

        * `workflow` (`dict`) - The workflow endpoints.

      * `integration_account` (`dict`) - The integration account.
        * `id` (`str`) - The resource id.
        * `name` (`str`) - Gets the resource name.
        * `type` (`str`) - Gets the resource type.

      * `integration_service_environment` (`dict`) - The integration service environment.
      * `parameters` (`dict`) - The parameters.
      * `provisioning_state` (`str`) - Gets the provisioning state.
      * `sku` (`dict`) - The sku.
        * `name` (`str`) - The name.
        * `plan` (`dict`) - The reference to plan.

      * `state` (`str`) - The state.
      * `version` (`str`) - Gets the version.
    """
    tags: pulumi.Output[dict]
    """
    The resource tags.
    """
    type: pulumi.Output[str]
    """
    Gets the resource type.
    """
    def __init__(__self__, resource_name, opts=None, access_control=None, definition=None, endpoints_configuration=None, integration_account=None, integration_service_environment=None, location=None, name=None, parameters=None, resource_group_name=None, state=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        The workflow type.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] access_control: The access control configuration.
        :param pulumi.Input[dict] definition: The definition.
        :param pulumi.Input[dict] endpoints_configuration: The endpoints configuration.
        :param pulumi.Input[dict] integration_account: The integration account.
        :param pulumi.Input[dict] integration_service_environment: The integration service environment.
        :param pulumi.Input[str] location: The resource location.
        :param pulumi.Input[str] name: The workflow name.
        :param pulumi.Input[dict] parameters: The parameters.
        :param pulumi.Input[str] resource_group_name: The resource group name.
        :param pulumi.Input[str] state: The state.
        :param pulumi.Input[dict] tags: The resource tags.

        The **access_control** object supports the following:

          * `actions` (`pulumi.Input[dict]`) - The access control configuration for workflow actions.
            * `allowed_caller_ip_addresses` (`pulumi.Input[list]`) - The allowed caller IP address ranges.
              * `address_range` (`pulumi.Input[str]`) - The IP address range.

            * `open_authentication_policies` (`pulumi.Input[dict]`) - The authentication policies for workflow.
              * `policies` (`pulumi.Input[dict]`) - Open authentication policies.

          * `contents` (`pulumi.Input[dict]`) - The access control configuration for accessing workflow run contents.
          * `triggers` (`pulumi.Input[dict]`) - The access control configuration for invoking workflow triggers.
          * `workflow_management` (`pulumi.Input[dict]`) - The access control configuration for workflow management.

        The **endpoints_configuration** object supports the following:

          * `connector` (`pulumi.Input[dict]`) - The connector endpoints.
            * `access_endpoint_ip_addresses` (`pulumi.Input[list]`) - The access endpoint ip address.
              * `address` (`pulumi.Input[str]`) - The address.

            * `outgoing_ip_addresses` (`pulumi.Input[list]`) - The outgoing ip address.

          * `workflow` (`pulumi.Input[dict]`) - The workflow endpoints.

        The **integration_account** object supports the following:

          * `id` (`pulumi.Input[str]`) - The resource id.
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

            __props__['access_control'] = access_control
            __props__['definition'] = definition
            __props__['endpoints_configuration'] = endpoints_configuration
            __props__['integration_account'] = integration_account
            __props__['integration_service_environment'] = integration_service_environment
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['parameters'] = parameters
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['state'] = state
            __props__['tags'] = tags
            __props__['properties'] = None
            __props__['type'] = None
        super(Workflow, __self__).__init__(
            'azurerm:logic/v20190501:Workflow',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing Workflow resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Workflow(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
