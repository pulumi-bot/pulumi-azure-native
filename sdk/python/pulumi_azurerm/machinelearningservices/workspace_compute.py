# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class WorkspaceCompute(pulumi.CustomResource):
    identity: pulumi.Output[dict]
    """
    The identity of the resource.
      * `principal_id` (`str`) - The principal ID of resource identity.
      * `tenant_id` (`str`) - The tenant ID of resource.
      * `type` (`str`) - The identity type.
      * `user_assigned_identities` (`dict`) - The list of user identities associated with resource. The user identity dictionary key references will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'.
    """
    location: pulumi.Output[str]
    """
    Specifies the location of the resource.
    """
    name: pulumi.Output[str]
    """
    Specifies the name of the resource.
    """
    properties: pulumi.Output[dict]
    """
    Compute properties
      * `compute_location` (`str`) - Location for the underlying compute
      * `compute_type` (`str`) - The type of compute
      * `created_on` (`str`) - The date and time when the compute was created.
      * `description` (`str`) - The description of the Machine Learning compute.
      * `is_attached_compute` (`bool`) - Indicating whether the compute was provisioned by user and brought from outside if true, or machine learning service provisioned it if false.
      * `modified_on` (`str`) - The date and time when the compute was last modified.
      * `provisioning_errors` (`list`) - Errors during provisioning
        * `error` (`dict`) - The error response.
          * `code` (`str`) - Error code.
          * `details` (`list`) - An array of error detail objects.
            * `code` (`str`) - Error code.
            * `message` (`str`) - Error message.

          * `message` (`str`) - Error message.

      * `provisioning_state` (`str`) - The provision state of the cluster. Valid values are Unknown, Updating, Provisioning, Succeeded, and Failed.
      * `resource_id` (`str`) - ARM resource id of the underlying compute
    """
    sku: pulumi.Output[dict]
    """
    The sku of the workspace.
      * `name` (`str`) - Name of the sku
      * `tier` (`str`) - Tier of the sku like Basic or Enterprise
    """
    tags: pulumi.Output[dict]
    """
    Contains resource tags defined as key/value pairs.
    """
    type: pulumi.Output[str]
    """
    Specifies the type of the resource.
    """
    def __init__(__self__, resource_name, opts=None, identity=None, location=None, name=None, properties=None, resource_group_name=None, sku=None, tags=None, workspace_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Machine Learning compute object wrapped into ARM resource envelope.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] identity: The identity of the resource.
        :param pulumi.Input[str] location: Specifies the location of the resource.
        :param pulumi.Input[str] name: Name of the Azure Machine Learning compute.
        :param pulumi.Input[dict] properties: Compute properties
        :param pulumi.Input[str] resource_group_name: Name of the resource group in which workspace is located.
        :param pulumi.Input[dict] sku: The sku of the workspace.
        :param pulumi.Input[dict] tags: Contains resource tags defined as key/value pairs.
        :param pulumi.Input[str] workspace_name: Name of Azure Machine Learning workspace.

        The **identity** object supports the following:

          * `type` (`pulumi.Input[str]`) - The identity type.
          * `user_assigned_identities` (`pulumi.Input[dict]`) - The list of user identities associated with resource. The user identity dictionary key references will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'.

        The **properties** object supports the following:

          * `compute_location` (`pulumi.Input[str]`) - Location for the underlying compute
          * `compute_type` (`pulumi.Input[str]`) - The type of compute
          * `description` (`pulumi.Input[str]`) - The description of the Machine Learning compute.
          * `resource_id` (`pulumi.Input[str]`) - ARM resource id of the underlying compute

        The **sku** object supports the following:

          * `name` (`pulumi.Input[str]`) - Name of the sku
          * `tier` (`pulumi.Input[str]`) - Tier of the sku like Basic or Enterprise
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['identity'] = identity
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['sku'] = sku
            __props__['tags'] = tags
            if workspace_name is None:
                raise TypeError("Missing required property 'workspace_name'")
            __props__['workspace_name'] = workspace_name
            __props__['type'] = None
        super(WorkspaceCompute, __self__).__init__(
            'azurerm:machinelearningservices:WorkspaceCompute',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing WorkspaceCompute resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return WorkspaceCompute(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
