# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables
from . import outputs

__all__ = [
    'GetApplicationDefinitionResult',
    'AwaitableGetApplicationDefinitionResult',
    'get_application_definition',
]

@pulumi.output_type
class GetApplicationDefinitionResult:
    """
    Information about managed application definition.
    """
    def __init__(__self__, artifacts=None, authorizations=None, create_ui_definition=None, deployment_policy=None, description=None, display_name=None, is_enabled=None, location=None, lock_level=None, locking_policy=None, main_template=None, managed_by=None, management_policy=None, name=None, notification_policy=None, package_file_uri=None, policies=None, sku=None, tags=None, type=None):
        if artifacts and not isinstance(artifacts, list):
            raise TypeError("Expected argument 'artifacts' to be a list")
        pulumi.set(__self__, "artifacts", artifacts)
        if authorizations and not isinstance(authorizations, list):
            raise TypeError("Expected argument 'authorizations' to be a list")
        pulumi.set(__self__, "authorizations", authorizations)
        if create_ui_definition and not isinstance(create_ui_definition, dict):
            raise TypeError("Expected argument 'create_ui_definition' to be a dict")
        pulumi.set(__self__, "create_ui_definition", create_ui_definition)
        if deployment_policy and not isinstance(deployment_policy, dict):
            raise TypeError("Expected argument 'deployment_policy' to be a dict")
        pulumi.set(__self__, "deployment_policy", deployment_policy)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if is_enabled and not isinstance(is_enabled, bool):
            raise TypeError("Expected argument 'is_enabled' to be a bool")
        pulumi.set(__self__, "is_enabled", is_enabled)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if lock_level and not isinstance(lock_level, str):
            raise TypeError("Expected argument 'lock_level' to be a str")
        pulumi.set(__self__, "lock_level", lock_level)
        if locking_policy and not isinstance(locking_policy, dict):
            raise TypeError("Expected argument 'locking_policy' to be a dict")
        pulumi.set(__self__, "locking_policy", locking_policy)
        if main_template and not isinstance(main_template, dict):
            raise TypeError("Expected argument 'main_template' to be a dict")
        pulumi.set(__self__, "main_template", main_template)
        if managed_by and not isinstance(managed_by, str):
            raise TypeError("Expected argument 'managed_by' to be a str")
        pulumi.set(__self__, "managed_by", managed_by)
        if management_policy and not isinstance(management_policy, dict):
            raise TypeError("Expected argument 'management_policy' to be a dict")
        pulumi.set(__self__, "management_policy", management_policy)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if notification_policy and not isinstance(notification_policy, dict):
            raise TypeError("Expected argument 'notification_policy' to be a dict")
        pulumi.set(__self__, "notification_policy", notification_policy)
        if package_file_uri and not isinstance(package_file_uri, str):
            raise TypeError("Expected argument 'package_file_uri' to be a str")
        pulumi.set(__self__, "package_file_uri", package_file_uri)
        if policies and not isinstance(policies, list):
            raise TypeError("Expected argument 'policies' to be a list")
        pulumi.set(__self__, "policies", policies)
        if sku and not isinstance(sku, dict):
            raise TypeError("Expected argument 'sku' to be a dict")
        pulumi.set(__self__, "sku", sku)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def artifacts(self) -> Optional[List['outputs.ApplicationDefinitionArtifactResponse']]:
        """
        The collection of managed application artifacts. The portal will use the files specified as artifacts to construct the user experience of creating a managed application from a managed application definition.
        """
        return pulumi.get(self, "artifacts")

    @property
    @pulumi.getter
    def authorizations(self) -> Optional[List['outputs.ApplicationAuthorizationResponse']]:
        """
        The managed application provider authorizations.
        """
        return pulumi.get(self, "authorizations")

    @property
    @pulumi.getter(name="createUiDefinition")
    def create_ui_definition(self) -> Optional[Mapping[str, Any]]:
        """
        The createUiDefinition json for the backing template with Microsoft.Solutions/applications resource. It can be a JObject or well-formed JSON string.
        """
        return pulumi.get(self, "create_ui_definition")

    @property
    @pulumi.getter(name="deploymentPolicy")
    def deployment_policy(self) -> Optional['outputs.ApplicationDeploymentPolicyResponse']:
        """
        The managed application deployment policy.
        """
        return pulumi.get(self, "deployment_policy")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The managed application definition description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[str]:
        """
        The managed application definition display name.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="isEnabled")
    def is_enabled(self) -> Optional[bool]:
        """
        A value indicating whether the package is enabled or not.
        """
        return pulumi.get(self, "is_enabled")

    @property
    @pulumi.getter
    def location(self) -> Optional[str]:
        """
        Resource location
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="lockLevel")
    def lock_level(self) -> str:
        """
        The managed application lock level.
        """
        return pulumi.get(self, "lock_level")

    @property
    @pulumi.getter(name="lockingPolicy")
    def locking_policy(self) -> Optional['outputs.ApplicationPackageLockingPolicyDefinitionResponse']:
        """
        The managed application locking policy.
        """
        return pulumi.get(self, "locking_policy")

    @property
    @pulumi.getter(name="mainTemplate")
    def main_template(self) -> Optional[Mapping[str, Any]]:
        """
        The inline main template json which has resources to be provisioned. It can be a JObject or well-formed JSON string.
        """
        return pulumi.get(self, "main_template")

    @property
    @pulumi.getter(name="managedBy")
    def managed_by(self) -> Optional[str]:
        """
        ID of the resource that manages this resource.
        """
        return pulumi.get(self, "managed_by")

    @property
    @pulumi.getter(name="managementPolicy")
    def management_policy(self) -> Optional['outputs.ApplicationManagementPolicyResponse']:
        """
        The managed application management policy that determines publisher's access to the managed resource group.
        """
        return pulumi.get(self, "management_policy")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="notificationPolicy")
    def notification_policy(self) -> Optional['outputs.ApplicationNotificationPolicyResponse']:
        """
        The managed application notification policy.
        """
        return pulumi.get(self, "notification_policy")

    @property
    @pulumi.getter(name="packageFileUri")
    def package_file_uri(self) -> Optional[str]:
        """
        The managed application definition package file Uri. Use this element
        """
        return pulumi.get(self, "package_file_uri")

    @property
    @pulumi.getter
    def policies(self) -> Optional[List['outputs.ApplicationPolicyResponse']]:
        """
        The managed application provider policies.
        """
        return pulumi.get(self, "policies")

    @property
    @pulumi.getter
    def sku(self) -> Optional['outputs.SkuResponse']:
        """
        The SKU of the resource.
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        Resource tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Resource type
        """
        return pulumi.get(self, "type")


class AwaitableGetApplicationDefinitionResult(GetApplicationDefinitionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetApplicationDefinitionResult(
            artifacts=self.artifacts,
            authorizations=self.authorizations,
            create_ui_definition=self.create_ui_definition,
            deployment_policy=self.deployment_policy,
            description=self.description,
            display_name=self.display_name,
            is_enabled=self.is_enabled,
            location=self.location,
            lock_level=self.lock_level,
            locking_policy=self.locking_policy,
            main_template=self.main_template,
            managed_by=self.managed_by,
            management_policy=self.management_policy,
            name=self.name,
            notification_policy=self.notification_policy,
            package_file_uri=self.package_file_uri,
            policies=self.policies,
            sku=self.sku,
            tags=self.tags,
            type=self.type)


def get_application_definition(name: Optional[str] = None,
                               resource_group_name: Optional[str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetApplicationDefinitionResult:
    """
    Use this data source to access information about an existing resource.

    :param str name: The name of the managed application definition.
    :param str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:solutions/v20190701:getApplicationDefinition', __args__, opts=opts, typ=GetApplicationDefinitionResult).value

    return AwaitableGetApplicationDefinitionResult(
        artifacts=__ret__.artifacts,
        authorizations=__ret__.authorizations,
        create_ui_definition=__ret__.create_ui_definition,
        deployment_policy=__ret__.deployment_policy,
        description=__ret__.description,
        display_name=__ret__.display_name,
        is_enabled=__ret__.is_enabled,
        location=__ret__.location,
        lock_level=__ret__.lock_level,
        locking_policy=__ret__.locking_policy,
        main_template=__ret__.main_template,
        managed_by=__ret__.managed_by,
        management_policy=__ret__.management_policy,
        name=__ret__.name,
        notification_policy=__ret__.notification_policy,
        package_file_uri=__ret__.package_file_uri,
        policies=__ret__.policies,
        sku=__ret__.sku,
        tags=__ret__.tags,
        type=__ret__.type)
