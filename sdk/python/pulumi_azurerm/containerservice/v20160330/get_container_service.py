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
    'GetContainerServiceResult',
    'AwaitableGetContainerServiceResult',
    'get_container_service',
]

@pulumi.output_type
class GetContainerServiceResult:
    """
    Container service.
    """
    def __init__(__self__, agent_pool_profiles=None, diagnostics_profile=None, linux_profile=None, location=None, master_profile=None, name=None, orchestrator_profile=None, provisioning_state=None, tags=None, type=None, windows_profile=None):
        if agent_pool_profiles and not isinstance(agent_pool_profiles, list):
            raise TypeError("Expected argument 'agent_pool_profiles' to be a list")
        pulumi.set(__self__, "agent_pool_profiles", agent_pool_profiles)
        if diagnostics_profile and not isinstance(diagnostics_profile, dict):
            raise TypeError("Expected argument 'diagnostics_profile' to be a dict")
        pulumi.set(__self__, "diagnostics_profile", diagnostics_profile)
        if linux_profile and not isinstance(linux_profile, dict):
            raise TypeError("Expected argument 'linux_profile' to be a dict")
        pulumi.set(__self__, "linux_profile", linux_profile)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if master_profile and not isinstance(master_profile, dict):
            raise TypeError("Expected argument 'master_profile' to be a dict")
        pulumi.set(__self__, "master_profile", master_profile)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if orchestrator_profile and not isinstance(orchestrator_profile, dict):
            raise TypeError("Expected argument 'orchestrator_profile' to be a dict")
        pulumi.set(__self__, "orchestrator_profile", orchestrator_profile)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if windows_profile and not isinstance(windows_profile, dict):
            raise TypeError("Expected argument 'windows_profile' to be a dict")
        pulumi.set(__self__, "windows_profile", windows_profile)

    @property
    @pulumi.getter(name="agentPoolProfiles")
    def agent_pool_profiles(self) -> List['outputs.ContainerServiceAgentPoolProfileResponse']:
        """
        Properties of the agent pool.
        """
        return pulumi.get(self, "agent_pool_profiles")

    @property
    @pulumi.getter(name="diagnosticsProfile")
    def diagnostics_profile(self) -> Optional['outputs.ContainerServiceDiagnosticsProfileResponse']:
        """
        Properties of the diagnostic agent.
        """
        return pulumi.get(self, "diagnostics_profile")

    @property
    @pulumi.getter(name="linuxProfile")
    def linux_profile(self) -> 'outputs.ContainerServiceLinuxProfileResponse':
        """
        Properties of Linux VMs.
        """
        return pulumi.get(self, "linux_profile")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        Resource location
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="masterProfile")
    def master_profile(self) -> 'outputs.ContainerServiceMasterProfileResponse':
        """
        Properties of master agents.
        """
        return pulumi.get(self, "master_profile")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="orchestratorProfile")
    def orchestrator_profile(self) -> Optional['outputs.ContainerServiceOrchestratorProfileResponse']:
        """
        Properties of the orchestrator.
        """
        return pulumi.get(self, "orchestrator_profile")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        the current deployment or provisioning state, which only appears in the response.
        """
        return pulumi.get(self, "provisioning_state")

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

    @property
    @pulumi.getter(name="windowsProfile")
    def windows_profile(self) -> Optional['outputs.ContainerServiceWindowsProfileResponse']:
        """
        Properties of Windows VMs.
        """
        return pulumi.get(self, "windows_profile")


class AwaitableGetContainerServiceResult(GetContainerServiceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetContainerServiceResult(
            agent_pool_profiles=self.agent_pool_profiles,
            diagnostics_profile=self.diagnostics_profile,
            linux_profile=self.linux_profile,
            location=self.location,
            master_profile=self.master_profile,
            name=self.name,
            orchestrator_profile=self.orchestrator_profile,
            provisioning_state=self.provisioning_state,
            tags=self.tags,
            type=self.type,
            windows_profile=self.windows_profile)


def get_container_service(name: Optional[str] = None,
                          resource_group_name: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetContainerServiceResult:
    """
    Use this data source to access information about an existing resource.

    :param str name: The name of the container service in the specified subscription and resource group.
    :param str resource_group_name: The name of the resource group.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:containerservice/v20160330:getContainerService', __args__, opts=opts, typ=GetContainerServiceResult).value

    return AwaitableGetContainerServiceResult(
        agent_pool_profiles=__ret__.agent_pool_profiles,
        diagnostics_profile=__ret__.diagnostics_profile,
        linux_profile=__ret__.linux_profile,
        location=__ret__.location,
        master_profile=__ret__.master_profile,
        name=__ret__.name,
        orchestrator_profile=__ret__.orchestrator_profile,
        provisioning_state=__ret__.provisioning_state,
        tags=__ret__.tags,
        type=__ret__.type,
        windows_profile=__ret__.windows_profile)
