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

__all__ = ['ApplicationDefinition']


class ApplicationDefinition(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_definition_name: Optional[pulumi.Input[str]] = None,
                 artifacts: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationDefinitionArtifactArgs']]]]] = None,
                 authorizations: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationAuthorizationArgs']]]]] = None,
                 create_ui_definition: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 deployment_policy: Optional[pulumi.Input[pulumi.InputType['ApplicationDeploymentPolicyArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 is_enabled: Optional[pulumi.Input[bool]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 lock_level: Optional[pulumi.Input[str]] = None,
                 locking_policy: Optional[pulumi.Input[pulumi.InputType['ApplicationPackageLockingPolicyDefinitionArgs']]] = None,
                 main_template: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 managed_by: Optional[pulumi.Input[str]] = None,
                 management_policy: Optional[pulumi.Input[pulumi.InputType['ApplicationManagementPolicyArgs']]] = None,
                 notification_policy: Optional[pulumi.Input[pulumi.InputType['ApplicationNotificationPolicyArgs']]] = None,
                 package_file_uri: Optional[pulumi.Input[str]] = None,
                 policies: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationPolicyArgs']]]]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 sku: Optional[pulumi.Input[pulumi.InputType['SkuArgs']]] = None,
                 storage_account_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Information about managed application definition.

        ## Example Usage
        ### Create or update managed application definition

        ```python
        import pulumi
        import pulumi_azurerm as azurerm

        application_definition = azurerm.solutions.latest.ApplicationDefinition("applicationDefinition",
            application_definition_name="myManagedApplicationDef",
            authorizations=[{
                "principalId": "validprincipalguid",
                "roleDefinitionId": "validroleguid",
            }],
            description="myManagedApplicationDef description",
            display_name="myManagedApplicationDef",
            lock_level="None",
            package_file_uri="https://path/to/packagezipfile",
            resource_group_name="rg")

        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_definition_name: The name of the managed application definition.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationDefinitionArtifactArgs']]]] artifacts: The collection of managed application artifacts. The portal will use the files specified as artifacts to construct the user experience of creating a managed application from a managed application definition.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationAuthorizationArgs']]]] authorizations: The managed application provider authorizations.
        :param pulumi.Input[Mapping[str, Any]] create_ui_definition: The createUiDefinition json for the backing template with Microsoft.Solutions/applications resource. It can be a JObject or well-formed JSON string.
        :param pulumi.Input[pulumi.InputType['ApplicationDeploymentPolicyArgs']] deployment_policy: The managed application deployment policy.
        :param pulumi.Input[str] description: The managed application definition description.
        :param pulumi.Input[str] display_name: The managed application definition display name.
        :param pulumi.Input[bool] is_enabled: A value indicating whether the package is enabled or not.
        :param pulumi.Input[str] location: Resource location
        :param pulumi.Input[str] lock_level: The managed application lock level.
        :param pulumi.Input[pulumi.InputType['ApplicationPackageLockingPolicyDefinitionArgs']] locking_policy: The managed application locking policy.
        :param pulumi.Input[Mapping[str, Any]] main_template: The inline main template json which has resources to be provisioned. It can be a JObject or well-formed JSON string.
        :param pulumi.Input[str] managed_by: ID of the resource that manages this resource.
        :param pulumi.Input[pulumi.InputType['ApplicationManagementPolicyArgs']] management_policy: The managed application management policy that determines publisher's access to the managed resource group.
        :param pulumi.Input[pulumi.InputType['ApplicationNotificationPolicyArgs']] notification_policy: The managed application notification policy.
        :param pulumi.Input[str] package_file_uri: The managed application definition package file Uri. Use this element
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ApplicationPolicyArgs']]]] policies: The managed application provider policies.
        :param pulumi.Input[str] resource_group_name: The name of the resource group. The name is case insensitive.
        :param pulumi.Input[pulumi.InputType['SkuArgs']] sku: The SKU of the resource.
        :param pulumi.Input[str] storage_account_id: The storage account id for bring your own storage scenario.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags
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

            if application_definition_name is None:
                raise TypeError("Missing required property 'application_definition_name'")
            __props__['application_definition_name'] = application_definition_name
            __props__['artifacts'] = artifacts
            __props__['authorizations'] = authorizations
            __props__['create_ui_definition'] = create_ui_definition
            __props__['deployment_policy'] = deployment_policy
            __props__['description'] = description
            __props__['display_name'] = display_name
            __props__['is_enabled'] = is_enabled
            __props__['location'] = location
            if lock_level is None:
                raise TypeError("Missing required property 'lock_level'")
            __props__['lock_level'] = lock_level
            __props__['locking_policy'] = locking_policy
            __props__['main_template'] = main_template
            __props__['managed_by'] = managed_by
            __props__['management_policy'] = management_policy
            __props__['notification_policy'] = notification_policy
            __props__['package_file_uri'] = package_file_uri
            __props__['policies'] = policies
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['sku'] = sku
            __props__['storage_account_id'] = storage_account_id
            __props__['tags'] = tags
            __props__['name'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:solutions/v20170901:ApplicationDefinition"), pulumi.Alias(type_="azurerm:solutions/v20180601:ApplicationDefinition"), pulumi.Alias(type_="azurerm:solutions/v20190701:ApplicationDefinition"), pulumi.Alias(type_="azurerm:solutions/v20200821preview:ApplicationDefinition")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(ApplicationDefinition, __self__).__init__(
            'azurerm:solutions/latest:ApplicationDefinition',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ApplicationDefinition':
        """
        Get an existing ApplicationDefinition resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ApplicationDefinition(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def artifacts(self) -> pulumi.Output[Optional[List['outputs.ApplicationDefinitionArtifactResponse']]]:
        """
        The collection of managed application artifacts. The portal will use the files specified as artifacts to construct the user experience of creating a managed application from a managed application definition.
        """
        return pulumi.get(self, "artifacts")

    @property
    @pulumi.getter
    def authorizations(self) -> pulumi.Output[Optional[List['outputs.ApplicationAuthorizationResponse']]]:
        """
        The managed application provider authorizations.
        """
        return pulumi.get(self, "authorizations")

    @property
    @pulumi.getter(name="createUiDefinition")
    def create_ui_definition(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        The createUiDefinition json for the backing template with Microsoft.Solutions/applications resource. It can be a JObject or well-formed JSON string.
        """
        return pulumi.get(self, "create_ui_definition")

    @property
    @pulumi.getter(name="deploymentPolicy")
    def deployment_policy(self) -> pulumi.Output[Optional['outputs.ApplicationDeploymentPolicyResponse']]:
        """
        The managed application deployment policy.
        """
        return pulumi.get(self, "deployment_policy")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The managed application definition description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[Optional[str]]:
        """
        The managed application definition display name.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="isEnabled")
    def is_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        A value indicating whether the package is enabled or not.
        """
        return pulumi.get(self, "is_enabled")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[Optional[str]]:
        """
        Resource location
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="lockLevel")
    def lock_level(self) -> pulumi.Output[str]:
        """
        The managed application lock level.
        """
        return pulumi.get(self, "lock_level")

    @property
    @pulumi.getter(name="lockingPolicy")
    def locking_policy(self) -> pulumi.Output[Optional['outputs.ApplicationPackageLockingPolicyDefinitionResponse']]:
        """
        The managed application locking policy.
        """
        return pulumi.get(self, "locking_policy")

    @property
    @pulumi.getter(name="mainTemplate")
    def main_template(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        The inline main template json which has resources to be provisioned. It can be a JObject or well-formed JSON string.
        """
        return pulumi.get(self, "main_template")

    @property
    @pulumi.getter(name="managedBy")
    def managed_by(self) -> pulumi.Output[Optional[str]]:
        """
        ID of the resource that manages this resource.
        """
        return pulumi.get(self, "managed_by")

    @property
    @pulumi.getter(name="managementPolicy")
    def management_policy(self) -> pulumi.Output[Optional['outputs.ApplicationManagementPolicyResponse']]:
        """
        The managed application management policy that determines publisher's access to the managed resource group.
        """
        return pulumi.get(self, "management_policy")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="notificationPolicy")
    def notification_policy(self) -> pulumi.Output[Optional['outputs.ApplicationNotificationPolicyResponse']]:
        """
        The managed application notification policy.
        """
        return pulumi.get(self, "notification_policy")

    @property
    @pulumi.getter(name="packageFileUri")
    def package_file_uri(self) -> pulumi.Output[Optional[str]]:
        """
        The managed application definition package file Uri. Use this element
        """
        return pulumi.get(self, "package_file_uri")

    @property
    @pulumi.getter
    def policies(self) -> pulumi.Output[Optional[List['outputs.ApplicationPolicyResponse']]]:
        """
        The managed application provider policies.
        """
        return pulumi.get(self, "policies")

    @property
    @pulumi.getter
    def sku(self) -> pulumi.Output[Optional['outputs.SkuResponse']]:
        """
        The SKU of the resource.
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter(name="storageAccountId")
    def storage_account_id(self) -> pulumi.Output[Optional[str]]:
        """
        The storage account id for bring your own storage scenario.
        """
        return pulumi.get(self, "storage_account_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Resource tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Resource type
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

