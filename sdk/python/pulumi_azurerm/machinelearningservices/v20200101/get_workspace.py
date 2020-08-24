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
    'GetWorkspaceResult',
    'AwaitableGetWorkspaceResult',
    'get_workspace',
]

@pulumi.output_type
class GetWorkspaceResult:
    """
    An object that represents a machine learning workspace.
    """
    def __init__(__self__, application_insights=None, container_registry=None, creation_time=None, description=None, discovery_url=None, encryption=None, friendly_name=None, hbi_workspace=None, identity=None, key_vault=None, location=None, name=None, provisioning_state=None, service_provisioned_resource_group=None, sku=None, storage_account=None, tags=None, type=None, workspace_id=None):
        if application_insights and not isinstance(application_insights, str):
            raise TypeError("Expected argument 'application_insights' to be a str")
        pulumi.set(__self__, "application_insights", application_insights)
        if container_registry and not isinstance(container_registry, str):
            raise TypeError("Expected argument 'container_registry' to be a str")
        pulumi.set(__self__, "container_registry", container_registry)
        if creation_time and not isinstance(creation_time, str):
            raise TypeError("Expected argument 'creation_time' to be a str")
        pulumi.set(__self__, "creation_time", creation_time)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if discovery_url and not isinstance(discovery_url, str):
            raise TypeError("Expected argument 'discovery_url' to be a str")
        pulumi.set(__self__, "discovery_url", discovery_url)
        if encryption and not isinstance(encryption, dict):
            raise TypeError("Expected argument 'encryption' to be a dict")
        pulumi.set(__self__, "encryption", encryption)
        if friendly_name and not isinstance(friendly_name, str):
            raise TypeError("Expected argument 'friendly_name' to be a str")
        pulumi.set(__self__, "friendly_name", friendly_name)
        if hbi_workspace and not isinstance(hbi_workspace, bool):
            raise TypeError("Expected argument 'hbi_workspace' to be a bool")
        pulumi.set(__self__, "hbi_workspace", hbi_workspace)
        if identity and not isinstance(identity, dict):
            raise TypeError("Expected argument 'identity' to be a dict")
        pulumi.set(__self__, "identity", identity)
        if key_vault and not isinstance(key_vault, str):
            raise TypeError("Expected argument 'key_vault' to be a str")
        pulumi.set(__self__, "key_vault", key_vault)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if service_provisioned_resource_group and not isinstance(service_provisioned_resource_group, str):
            raise TypeError("Expected argument 'service_provisioned_resource_group' to be a str")
        pulumi.set(__self__, "service_provisioned_resource_group", service_provisioned_resource_group)
        if sku and not isinstance(sku, dict):
            raise TypeError("Expected argument 'sku' to be a dict")
        pulumi.set(__self__, "sku", sku)
        if storage_account and not isinstance(storage_account, str):
            raise TypeError("Expected argument 'storage_account' to be a str")
        pulumi.set(__self__, "storage_account", storage_account)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if workspace_id and not isinstance(workspace_id, str):
            raise TypeError("Expected argument 'workspace_id' to be a str")
        pulumi.set(__self__, "workspace_id", workspace_id)

    @property
    @pulumi.getter(name="applicationInsights")
    def application_insights(self) -> Optional[str]:
        """
        ARM id of the application insights associated with this workspace. This cannot be changed once the workspace has been created
        """
        return pulumi.get(self, "application_insights")

    @property
    @pulumi.getter(name="containerRegistry")
    def container_registry(self) -> Optional[str]:
        """
        ARM id of the container registry associated with this workspace. This cannot be changed once the workspace has been created
        """
        return pulumi.get(self, "container_registry")

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> str:
        """
        The creation time of the machine learning workspace in ISO8601 format.
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The description of this workspace.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="discoveryUrl")
    def discovery_url(self) -> Optional[str]:
        """
        Url for the discovery service to identify regional endpoints for machine learning experimentation services
        """
        return pulumi.get(self, "discovery_url")

    @property
    @pulumi.getter
    def encryption(self) -> Optional['outputs.EncryptionPropertyResponse']:
        """
        The encryption settings of Azure ML workspace.
        """
        return pulumi.get(self, "encryption")

    @property
    @pulumi.getter(name="friendlyName")
    def friendly_name(self) -> Optional[str]:
        """
        The friendly name for this workspace. This name in mutable
        """
        return pulumi.get(self, "friendly_name")

    @property
    @pulumi.getter(name="hbiWorkspace")
    def hbi_workspace(self) -> Optional[bool]:
        """
        The flag to signal HBI data in the workspace and reduce diagnostic data collected by the service
        """
        return pulumi.get(self, "hbi_workspace")

    @property
    @pulumi.getter
    def identity(self) -> Optional['outputs.IdentityResponse']:
        """
        The identity of the resource.
        """
        return pulumi.get(self, "identity")

    @property
    @pulumi.getter(name="keyVault")
    def key_vault(self) -> Optional[str]:
        """
        ARM id of the key vault associated with this workspace. This cannot be changed once the workspace has been created
        """
        return pulumi.get(self, "key_vault")

    @property
    @pulumi.getter
    def location(self) -> Optional[str]:
        """
        Specifies the location of the resource.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Specifies the name of the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        The current deployment state of workspace resource. The provisioningState is to indicate states for resource provisioning.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="serviceProvisionedResourceGroup")
    def service_provisioned_resource_group(self) -> str:
        """
        The name of the managed resource group created by workspace RP in customer subscription if the workspace is CMK workspace
        """
        return pulumi.get(self, "service_provisioned_resource_group")

    @property
    @pulumi.getter
    def sku(self) -> Optional['outputs.SkuResponse']:
        """
        The sku of the workspace.
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter(name="storageAccount")
    def storage_account(self) -> Optional[str]:
        """
        ARM id of the storage account associated with this workspace. This cannot be changed once the workspace has been created
        """
        return pulumi.get(self, "storage_account")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        Contains resource tags defined as key/value pairs.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Specifies the type of the resource.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="workspaceId")
    def workspace_id(self) -> str:
        """
        The immutable id associated with this workspace.
        """
        return pulumi.get(self, "workspace_id")


class AwaitableGetWorkspaceResult(GetWorkspaceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetWorkspaceResult(
            application_insights=self.application_insights,
            container_registry=self.container_registry,
            creation_time=self.creation_time,
            description=self.description,
            discovery_url=self.discovery_url,
            encryption=self.encryption,
            friendly_name=self.friendly_name,
            hbi_workspace=self.hbi_workspace,
            identity=self.identity,
            key_vault=self.key_vault,
            location=self.location,
            name=self.name,
            provisioning_state=self.provisioning_state,
            service_provisioned_resource_group=self.service_provisioned_resource_group,
            sku=self.sku,
            storage_account=self.storage_account,
            tags=self.tags,
            type=self.type,
            workspace_id=self.workspace_id)


def get_workspace(name: Optional[str] = None,
                  resource_group_name: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetWorkspaceResult:
    """
    Use this data source to access information about an existing resource.

    :param str name: Name of Azure Machine Learning workspace.
    :param str resource_group_name: Name of the resource group in which workspace is located.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:machinelearningservices/v20200101:getWorkspace', __args__, opts=opts, typ=GetWorkspaceResult).value

    return AwaitableGetWorkspaceResult(
        application_insights=__ret__.application_insights,
        container_registry=__ret__.container_registry,
        creation_time=__ret__.creation_time,
        description=__ret__.description,
        discovery_url=__ret__.discovery_url,
        encryption=__ret__.encryption,
        friendly_name=__ret__.friendly_name,
        hbi_workspace=__ret__.hbi_workspace,
        identity=__ret__.identity,
        key_vault=__ret__.key_vault,
        location=__ret__.location,
        name=__ret__.name,
        provisioning_state=__ret__.provisioning_state,
        service_provisioned_resource_group=__ret__.service_provisioned_resource_group,
        sku=__ret__.sku,
        storage_account=__ret__.storage_account,
        tags=__ret__.tags,
        type=__ret__.type,
        workspace_id=__ret__.workspace_id)
