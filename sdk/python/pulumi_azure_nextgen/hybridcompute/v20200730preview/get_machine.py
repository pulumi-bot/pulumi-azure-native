# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from . import outputs

__all__ = [
    'GetMachineResult',
    'AwaitableGetMachineResult',
    'get_machine',
]

@pulumi.output_type
class GetMachineResult:
    """
    Describes a hybrid machine.
    """
    def __init__(__self__, ad_fqdn=None, agent_version=None, client_public_key=None, display_name=None, dns_fqdn=None, domain_name=None, error_details=None, extensions=None, identity=None, last_status_change=None, location=None, location_data=None, machine_fqdn=None, name=None, os_name=None, os_profile=None, os_sku=None, os_version=None, provisioning_state=None, status=None, tags=None, type=None, vm_id=None, vm_uuid=None):
        if ad_fqdn and not isinstance(ad_fqdn, str):
            raise TypeError("Expected argument 'ad_fqdn' to be a str")
        pulumi.set(__self__, "ad_fqdn", ad_fqdn)
        if agent_version and not isinstance(agent_version, str):
            raise TypeError("Expected argument 'agent_version' to be a str")
        pulumi.set(__self__, "agent_version", agent_version)
        if client_public_key and not isinstance(client_public_key, str):
            raise TypeError("Expected argument 'client_public_key' to be a str")
        pulumi.set(__self__, "client_public_key", client_public_key)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if dns_fqdn and not isinstance(dns_fqdn, str):
            raise TypeError("Expected argument 'dns_fqdn' to be a str")
        pulumi.set(__self__, "dns_fqdn", dns_fqdn)
        if domain_name and not isinstance(domain_name, str):
            raise TypeError("Expected argument 'domain_name' to be a str")
        pulumi.set(__self__, "domain_name", domain_name)
        if error_details and not isinstance(error_details, list):
            raise TypeError("Expected argument 'error_details' to be a list")
        pulumi.set(__self__, "error_details", error_details)
        if extensions and not isinstance(extensions, list):
            raise TypeError("Expected argument 'extensions' to be a list")
        pulumi.set(__self__, "extensions", extensions)
        if identity and not isinstance(identity, dict):
            raise TypeError("Expected argument 'identity' to be a dict")
        pulumi.set(__self__, "identity", identity)
        if last_status_change and not isinstance(last_status_change, str):
            raise TypeError("Expected argument 'last_status_change' to be a str")
        pulumi.set(__self__, "last_status_change", last_status_change)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if location_data and not isinstance(location_data, dict):
            raise TypeError("Expected argument 'location_data' to be a dict")
        pulumi.set(__self__, "location_data", location_data)
        if machine_fqdn and not isinstance(machine_fqdn, str):
            raise TypeError("Expected argument 'machine_fqdn' to be a str")
        pulumi.set(__self__, "machine_fqdn", machine_fqdn)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if os_name and not isinstance(os_name, str):
            raise TypeError("Expected argument 'os_name' to be a str")
        pulumi.set(__self__, "os_name", os_name)
        if os_profile and not isinstance(os_profile, dict):
            raise TypeError("Expected argument 'os_profile' to be a dict")
        pulumi.set(__self__, "os_profile", os_profile)
        if os_sku and not isinstance(os_sku, str):
            raise TypeError("Expected argument 'os_sku' to be a str")
        pulumi.set(__self__, "os_sku", os_sku)
        if os_version and not isinstance(os_version, str):
            raise TypeError("Expected argument 'os_version' to be a str")
        pulumi.set(__self__, "os_version", os_version)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if vm_id and not isinstance(vm_id, str):
            raise TypeError("Expected argument 'vm_id' to be a str")
        pulumi.set(__self__, "vm_id", vm_id)
        if vm_uuid and not isinstance(vm_uuid, str):
            raise TypeError("Expected argument 'vm_uuid' to be a str")
        pulumi.set(__self__, "vm_uuid", vm_uuid)

    @property
    @pulumi.getter(name="adFqdn")
    def ad_fqdn(self) -> str:
        """
        Specifies the AD fully qualified display name.
        """
        return pulumi.get(self, "ad_fqdn")

    @property
    @pulumi.getter(name="agentVersion")
    def agent_version(self) -> str:
        """
        The hybrid machine agent full version.
        """
        return pulumi.get(self, "agent_version")

    @property
    @pulumi.getter(name="clientPublicKey")
    def client_public_key(self) -> Optional[str]:
        """
        Public Key that the client provides to be used during initial resource onboarding
        """
        return pulumi.get(self, "client_public_key")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        Specifies the hybrid machine display name.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="dnsFqdn")
    def dns_fqdn(self) -> str:
        """
        Specifies the DNS fully qualified display name.
        """
        return pulumi.get(self, "dns_fqdn")

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> str:
        """
        Specifies the Windows domain name.
        """
        return pulumi.get(self, "domain_name")

    @property
    @pulumi.getter(name="errorDetails")
    def error_details(self) -> Sequence['outputs.ErrorDetailResponse']:
        """
        Details about the error state.
        """
        return pulumi.get(self, "error_details")

    @property
    @pulumi.getter
    def extensions(self) -> Optional[Sequence['outputs.MachineExtensionInstanceViewResponse']]:
        """
        Machine Extensions information
        """
        return pulumi.get(self, "extensions")

    @property
    @pulumi.getter
    def identity(self) -> Optional['outputs.MachineResponseIdentity']:
        return pulumi.get(self, "identity")

    @property
    @pulumi.getter(name="lastStatusChange")
    def last_status_change(self) -> str:
        """
        The time of the last status change.
        """
        return pulumi.get(self, "last_status_change")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="locationData")
    def location_data(self) -> Optional['outputs.LocationDataResponse']:
        """
        Metadata pertaining to the geographic location of the resource.
        """
        return pulumi.get(self, "location_data")

    @property
    @pulumi.getter(name="machineFqdn")
    def machine_fqdn(self) -> str:
        """
        Specifies the hybrid machine FQDN.
        """
        return pulumi.get(self, "machine_fqdn")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="osName")
    def os_name(self) -> str:
        """
        The Operating System running on the hybrid machine.
        """
        return pulumi.get(self, "os_name")

    @property
    @pulumi.getter(name="osProfile")
    def os_profile(self) -> Optional['outputs.MachinePropertiesResponseOsProfile']:
        """
        Specifies the operating system settings for the hybrid machine.
        """
        return pulumi.get(self, "os_profile")

    @property
    @pulumi.getter(name="osSku")
    def os_sku(self) -> str:
        """
        Specifies the Operating System product SKU.
        """
        return pulumi.get(self, "os_sku")

    @property
    @pulumi.getter(name="osVersion")
    def os_version(self) -> str:
        """
        The version of Operating System running on the hybrid machine.
        """
        return pulumi.get(self, "os_version")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        The provisioning state, which only appears in the response.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        The status of the hybrid machine agent.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="vmId")
    def vm_id(self) -> Optional[str]:
        """
        Specifies the hybrid machine unique ID.
        """
        return pulumi.get(self, "vm_id")

    @property
    @pulumi.getter(name="vmUuid")
    def vm_uuid(self) -> str:
        """
        Specifies the Arc Machine's unique SMBIOS ID
        """
        return pulumi.get(self, "vm_uuid")


class AwaitableGetMachineResult(GetMachineResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMachineResult(
            ad_fqdn=self.ad_fqdn,
            agent_version=self.agent_version,
            client_public_key=self.client_public_key,
            display_name=self.display_name,
            dns_fqdn=self.dns_fqdn,
            domain_name=self.domain_name,
            error_details=self.error_details,
            extensions=self.extensions,
            identity=self.identity,
            last_status_change=self.last_status_change,
            location=self.location,
            location_data=self.location_data,
            machine_fqdn=self.machine_fqdn,
            name=self.name,
            os_name=self.os_name,
            os_profile=self.os_profile,
            os_sku=self.os_sku,
            os_version=self.os_version,
            provisioning_state=self.provisioning_state,
            status=self.status,
            tags=self.tags,
            type=self.type,
            vm_id=self.vm_id,
            vm_uuid=self.vm_uuid)


def get_machine(expand: Optional[str] = None,
                name: Optional[str] = None,
                resource_group_name: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMachineResult:
    """
    Use this data source to access information about an existing resource.

    :param str expand: The expand expression to apply on the operation.
    :param str name: The name of the hybrid machine.
    :param str resource_group_name: The name of the resource group.
    """
    __args__ = dict()
    __args__['expand'] = expand
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure-nextgen:hybridcompute/v20200730preview:getMachine', __args__, opts=opts, typ=GetMachineResult).value

    return AwaitableGetMachineResult(
        ad_fqdn=__ret__.ad_fqdn,
        agent_version=__ret__.agent_version,
        client_public_key=__ret__.client_public_key,
        display_name=__ret__.display_name,
        dns_fqdn=__ret__.dns_fqdn,
        domain_name=__ret__.domain_name,
        error_details=__ret__.error_details,
        extensions=__ret__.extensions,
        identity=__ret__.identity,
        last_status_change=__ret__.last_status_change,
        location=__ret__.location,
        location_data=__ret__.location_data,
        machine_fqdn=__ret__.machine_fqdn,
        name=__ret__.name,
        os_name=__ret__.os_name,
        os_profile=__ret__.os_profile,
        os_sku=__ret__.os_sku,
        os_version=__ret__.os_version,
        provisioning_state=__ret__.provisioning_state,
        status=__ret__.status,
        tags=__ret__.tags,
        type=__ret__.type,
        vm_id=__ret__.vm_id,
        vm_uuid=__ret__.vm_uuid)
