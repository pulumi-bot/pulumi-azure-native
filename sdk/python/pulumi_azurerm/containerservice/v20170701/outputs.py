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
    'ContainerServiceAgentPoolProfileResponse',
    'ContainerServiceCustomProfileResponse',
    'ContainerServiceDiagnosticsProfileResponse',
    'ContainerServiceLinuxProfileResponse',
    'ContainerServiceMasterProfileResponse',
    'ContainerServiceOrchestratorProfileResponse',
    'ContainerServiceServicePrincipalProfileResponse',
    'ContainerServiceSshConfigurationResponse',
    'ContainerServiceSshPublicKeyResponse',
    'ContainerServiceVMDiagnosticsResponse',
    'ContainerServiceWindowsProfileResponse',
    'KeyVaultSecretRefResponse',
]

@pulumi.output_type
class ContainerServiceAgentPoolProfileResponse(dict):
    """
    Profile for the container service agent pool.
    """
    def __init__(__self__, *,
                 fqdn: str,
                 name: str,
                 vm_size: str,
                 count: Optional[float] = None,
                 dns_prefix: Optional[str] = None,
                 os_disk_size_gb: Optional[float] = None,
                 os_type: Optional[str] = None,
                 ports: Optional[List[float]] = None,
                 storage_profile: Optional[str] = None,
                 vnet_subnet_id: Optional[str] = None):
        """
        Profile for the container service agent pool.
        :param str fqdn: FQDN for the agent pool.
        :param str name: Unique name of the agent pool profile in the context of the subscription and resource group.
        :param str vm_size: Size of agent VMs.
        :param float count: Number of agents (VMs) to host docker containers. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 1. 
        :param str dns_prefix: DNS prefix to be used to create the FQDN for the agent pool.
        :param float os_disk_size_gb: OS Disk Size in GB to be used to specify the disk size for every machine in this master/agent pool. If you specify 0, it will apply the default osDisk size according to the vmSize specified.
        :param str os_type: OsType to be used to specify os type. Choose from Linux and Windows. Default to Linux.
        :param List[float] ports: Ports number array used to expose on this agent pool. The default opened ports are different based on your choice of orchestrator.
        :param str storage_profile: Storage profile specifies what kind of storage used. Choose from StorageAccount and ManagedDisks. Leave it empty, we will choose for you based on the orchestrator choice.
        :param str vnet_subnet_id: VNet SubnetID specifies the VNet's subnet identifier.
        """
        pulumi.set(__self__, "fqdn", fqdn)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "vm_size", vm_size)
        if count is not None:
            pulumi.set(__self__, "count", count)
        if dns_prefix is not None:
            pulumi.set(__self__, "dns_prefix", dns_prefix)
        if os_disk_size_gb is not None:
            pulumi.set(__self__, "os_disk_size_gb", os_disk_size_gb)
        if os_type is not None:
            pulumi.set(__self__, "os_type", os_type)
        if ports is not None:
            pulumi.set(__self__, "ports", ports)
        if storage_profile is not None:
            pulumi.set(__self__, "storage_profile", storage_profile)
        if vnet_subnet_id is not None:
            pulumi.set(__self__, "vnet_subnet_id", vnet_subnet_id)

    @property
    @pulumi.getter
    def fqdn(self) -> str:
        """
        FQDN for the agent pool.
        """
        return pulumi.get(self, "fqdn")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Unique name of the agent pool profile in the context of the subscription and resource group.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="vmSize")
    def vm_size(self) -> str:
        """
        Size of agent VMs.
        """
        return pulumi.get(self, "vm_size")

    @property
    @pulumi.getter
    def count(self) -> Optional[float]:
        """
        Number of agents (VMs) to host docker containers. Allowed values must be in the range of 1 to 100 (inclusive). The default value is 1. 
        """
        return pulumi.get(self, "count")

    @property
    @pulumi.getter(name="dnsPrefix")
    def dns_prefix(self) -> Optional[str]:
        """
        DNS prefix to be used to create the FQDN for the agent pool.
        """
        return pulumi.get(self, "dns_prefix")

    @property
    @pulumi.getter(name="osDiskSizeGB")
    def os_disk_size_gb(self) -> Optional[float]:
        """
        OS Disk Size in GB to be used to specify the disk size for every machine in this master/agent pool. If you specify 0, it will apply the default osDisk size according to the vmSize specified.
        """
        return pulumi.get(self, "os_disk_size_gb")

    @property
    @pulumi.getter(name="osType")
    def os_type(self) -> Optional[str]:
        """
        OsType to be used to specify os type. Choose from Linux and Windows. Default to Linux.
        """
        return pulumi.get(self, "os_type")

    @property
    @pulumi.getter
    def ports(self) -> Optional[List[float]]:
        """
        Ports number array used to expose on this agent pool. The default opened ports are different based on your choice of orchestrator.
        """
        return pulumi.get(self, "ports")

    @property
    @pulumi.getter(name="storageProfile")
    def storage_profile(self) -> Optional[str]:
        """
        Storage profile specifies what kind of storage used. Choose from StorageAccount and ManagedDisks. Leave it empty, we will choose for you based on the orchestrator choice.
        """
        return pulumi.get(self, "storage_profile")

    @property
    @pulumi.getter(name="vnetSubnetID")
    def vnet_subnet_id(self) -> Optional[str]:
        """
        VNet SubnetID specifies the VNet's subnet identifier.
        """
        return pulumi.get(self, "vnet_subnet_id")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ContainerServiceCustomProfileResponse(dict):
    """
    Properties to configure a custom container service cluster.
    """
    def __init__(__self__, *,
                 orchestrator: str):
        """
        Properties to configure a custom container service cluster.
        :param str orchestrator: The name of the custom orchestrator to use.
        """
        pulumi.set(__self__, "orchestrator", orchestrator)

    @property
    @pulumi.getter
    def orchestrator(self) -> str:
        """
        The name of the custom orchestrator to use.
        """
        return pulumi.get(self, "orchestrator")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ContainerServiceDiagnosticsProfileResponse(dict):
    """
    Profile for diagnostics on the container service cluster.
    """
    def __init__(__self__, *,
                 vm_diagnostics: 'outputs.ContainerServiceVMDiagnosticsResponse'):
        """
        Profile for diagnostics on the container service cluster.
        :param 'ContainerServiceVMDiagnosticsResponseArgs' vm_diagnostics: Profile for diagnostics on the container service VMs.
        """
        pulumi.set(__self__, "vm_diagnostics", vm_diagnostics)

    @property
    @pulumi.getter(name="vmDiagnostics")
    def vm_diagnostics(self) -> 'outputs.ContainerServiceVMDiagnosticsResponse':
        """
        Profile for diagnostics on the container service VMs.
        """
        return pulumi.get(self, "vm_diagnostics")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ContainerServiceLinuxProfileResponse(dict):
    """
    Profile for Linux VMs in the container service cluster.
    """
    def __init__(__self__, *,
                 admin_username: str,
                 ssh: 'outputs.ContainerServiceSshConfigurationResponse'):
        """
        Profile for Linux VMs in the container service cluster.
        :param str admin_username: The administrator username to use for Linux VMs.
        :param 'ContainerServiceSshConfigurationResponseArgs' ssh: SSH configuration for Linux-based VMs running on Azure.
        """
        pulumi.set(__self__, "admin_username", admin_username)
        pulumi.set(__self__, "ssh", ssh)

    @property
    @pulumi.getter(name="adminUsername")
    def admin_username(self) -> str:
        """
        The administrator username to use for Linux VMs.
        """
        return pulumi.get(self, "admin_username")

    @property
    @pulumi.getter
    def ssh(self) -> 'outputs.ContainerServiceSshConfigurationResponse':
        """
        SSH configuration for Linux-based VMs running on Azure.
        """
        return pulumi.get(self, "ssh")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ContainerServiceMasterProfileResponse(dict):
    """
    Profile for the container service master.
    """
    def __init__(__self__, *,
                 dns_prefix: str,
                 fqdn: str,
                 vm_size: str,
                 count: Optional[float] = None,
                 first_consecutive_static_ip: Optional[str] = None,
                 os_disk_size_gb: Optional[float] = None,
                 storage_profile: Optional[str] = None,
                 vnet_subnet_id: Optional[str] = None):
        """
        Profile for the container service master.
        :param str dns_prefix: DNS prefix to be used to create the FQDN for the master pool.
        :param str fqdn: FQDN for the master pool.
        :param str vm_size: Size of agent VMs.
        :param float count: Number of masters (VMs) in the container service cluster. Allowed values are 1, 3, and 5. The default value is 1.
        :param str first_consecutive_static_ip: FirstConsecutiveStaticIP used to specify the first static ip of masters.
        :param float os_disk_size_gb: OS Disk Size in GB to be used to specify the disk size for every machine in this master/agent pool. If you specify 0, it will apply the default osDisk size according to the vmSize specified.
        :param str storage_profile: Storage profile specifies what kind of storage used. Choose from StorageAccount and ManagedDisks. Leave it empty, we will choose for you based on the orchestrator choice.
        :param str vnet_subnet_id: VNet SubnetID specifies the VNet's subnet identifier.
        """
        pulumi.set(__self__, "dns_prefix", dns_prefix)
        pulumi.set(__self__, "fqdn", fqdn)
        pulumi.set(__self__, "vm_size", vm_size)
        if count is not None:
            pulumi.set(__self__, "count", count)
        if first_consecutive_static_ip is not None:
            pulumi.set(__self__, "first_consecutive_static_ip", first_consecutive_static_ip)
        if os_disk_size_gb is not None:
            pulumi.set(__self__, "os_disk_size_gb", os_disk_size_gb)
        if storage_profile is not None:
            pulumi.set(__self__, "storage_profile", storage_profile)
        if vnet_subnet_id is not None:
            pulumi.set(__self__, "vnet_subnet_id", vnet_subnet_id)

    @property
    @pulumi.getter(name="dnsPrefix")
    def dns_prefix(self) -> str:
        """
        DNS prefix to be used to create the FQDN for the master pool.
        """
        return pulumi.get(self, "dns_prefix")

    @property
    @pulumi.getter
    def fqdn(self) -> str:
        """
        FQDN for the master pool.
        """
        return pulumi.get(self, "fqdn")

    @property
    @pulumi.getter(name="vmSize")
    def vm_size(self) -> str:
        """
        Size of agent VMs.
        """
        return pulumi.get(self, "vm_size")

    @property
    @pulumi.getter
    def count(self) -> Optional[float]:
        """
        Number of masters (VMs) in the container service cluster. Allowed values are 1, 3, and 5. The default value is 1.
        """
        return pulumi.get(self, "count")

    @property
    @pulumi.getter(name="firstConsecutiveStaticIP")
    def first_consecutive_static_ip(self) -> Optional[str]:
        """
        FirstConsecutiveStaticIP used to specify the first static ip of masters.
        """
        return pulumi.get(self, "first_consecutive_static_ip")

    @property
    @pulumi.getter(name="osDiskSizeGB")
    def os_disk_size_gb(self) -> Optional[float]:
        """
        OS Disk Size in GB to be used to specify the disk size for every machine in this master/agent pool. If you specify 0, it will apply the default osDisk size according to the vmSize specified.
        """
        return pulumi.get(self, "os_disk_size_gb")

    @property
    @pulumi.getter(name="storageProfile")
    def storage_profile(self) -> Optional[str]:
        """
        Storage profile specifies what kind of storage used. Choose from StorageAccount and ManagedDisks. Leave it empty, we will choose for you based on the orchestrator choice.
        """
        return pulumi.get(self, "storage_profile")

    @property
    @pulumi.getter(name="vnetSubnetID")
    def vnet_subnet_id(self) -> Optional[str]:
        """
        VNet SubnetID specifies the VNet's subnet identifier.
        """
        return pulumi.get(self, "vnet_subnet_id")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ContainerServiceOrchestratorProfileResponse(dict):
    """
    Profile for the container service orchestrator.
    """
    def __init__(__self__, *,
                 orchestrator_type: str,
                 orchestrator_version: Optional[str] = None):
        """
        Profile for the container service orchestrator.
        :param str orchestrator_type: The orchestrator to use to manage container service cluster resources. Valid values are Kubernetes, Swarm, DCOS, DockerCE and Custom.
        :param str orchestrator_version: The version of the orchestrator to use. You can specify the major.minor.patch part of the actual version.For example, you can specify version as "1.6.11".
        """
        pulumi.set(__self__, "orchestrator_type", orchestrator_type)
        if orchestrator_version is not None:
            pulumi.set(__self__, "orchestrator_version", orchestrator_version)

    @property
    @pulumi.getter(name="orchestratorType")
    def orchestrator_type(self) -> str:
        """
        The orchestrator to use to manage container service cluster resources. Valid values are Kubernetes, Swarm, DCOS, DockerCE and Custom.
        """
        return pulumi.get(self, "orchestrator_type")

    @property
    @pulumi.getter(name="orchestratorVersion")
    def orchestrator_version(self) -> Optional[str]:
        """
        The version of the orchestrator to use. You can specify the major.minor.patch part of the actual version.For example, you can specify version as "1.6.11".
        """
        return pulumi.get(self, "orchestrator_version")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ContainerServiceServicePrincipalProfileResponse(dict):
    """
    Information about a service principal identity for the cluster to use for manipulating Azure APIs. Either secret or keyVaultSecretRef must be specified.
    """
    def __init__(__self__, *,
                 client_id: str,
                 key_vault_secret_ref: Optional['outputs.KeyVaultSecretRefResponse'] = None,
                 secret: Optional[str] = None):
        """
        Information about a service principal identity for the cluster to use for manipulating Azure APIs. Either secret or keyVaultSecretRef must be specified.
        :param str client_id: The ID for the service principal.
        :param 'KeyVaultSecretRefResponseArgs' key_vault_secret_ref: Reference to a secret stored in Azure Key Vault.
        :param str secret: The secret password associated with the service principal in plain text.
        """
        pulumi.set(__self__, "client_id", client_id)
        if key_vault_secret_ref is not None:
            pulumi.set(__self__, "key_vault_secret_ref", key_vault_secret_ref)
        if secret is not None:
            pulumi.set(__self__, "secret", secret)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> str:
        """
        The ID for the service principal.
        """
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter(name="keyVaultSecretRef")
    def key_vault_secret_ref(self) -> Optional['outputs.KeyVaultSecretRefResponse']:
        """
        Reference to a secret stored in Azure Key Vault.
        """
        return pulumi.get(self, "key_vault_secret_ref")

    @property
    @pulumi.getter
    def secret(self) -> Optional[str]:
        """
        The secret password associated with the service principal in plain text.
        """
        return pulumi.get(self, "secret")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ContainerServiceSshConfigurationResponse(dict):
    """
    SSH configuration for Linux-based VMs running on Azure.
    """
    def __init__(__self__, *,
                 public_keys: List['outputs.ContainerServiceSshPublicKeyResponse']):
        """
        SSH configuration for Linux-based VMs running on Azure.
        :param List['ContainerServiceSshPublicKeyResponseArgs'] public_keys: The list of SSH public keys used to authenticate with Linux-based VMs. Only expect one key specified.
        """
        pulumi.set(__self__, "public_keys", public_keys)

    @property
    @pulumi.getter(name="publicKeys")
    def public_keys(self) -> List['outputs.ContainerServiceSshPublicKeyResponse']:
        """
        The list of SSH public keys used to authenticate with Linux-based VMs. Only expect one key specified.
        """
        return pulumi.get(self, "public_keys")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ContainerServiceSshPublicKeyResponse(dict):
    """
    Contains information about SSH certificate public key data.
    """
    def __init__(__self__, *,
                 key_data: str):
        """
        Contains information about SSH certificate public key data.
        :param str key_data: Certificate public key used to authenticate with VMs through SSH. The certificate must be in PEM format with or without headers.
        """
        pulumi.set(__self__, "key_data", key_data)

    @property
    @pulumi.getter(name="keyData")
    def key_data(self) -> str:
        """
        Certificate public key used to authenticate with VMs through SSH. The certificate must be in PEM format with or without headers.
        """
        return pulumi.get(self, "key_data")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ContainerServiceVMDiagnosticsResponse(dict):
    """
    Profile for diagnostics on the container service VMs.
    """
    def __init__(__self__, *,
                 enabled: bool,
                 storage_uri: str):
        """
        Profile for diagnostics on the container service VMs.
        :param bool enabled: Whether the VM diagnostic agent is provisioned on the VM.
        :param str storage_uri: The URI of the storage account where diagnostics are stored.
        """
        pulumi.set(__self__, "enabled", enabled)
        pulumi.set(__self__, "storage_uri", storage_uri)

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        """
        Whether the VM diagnostic agent is provisioned on the VM.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="storageUri")
    def storage_uri(self) -> str:
        """
        The URI of the storage account where diagnostics are stored.
        """
        return pulumi.get(self, "storage_uri")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ContainerServiceWindowsProfileResponse(dict):
    """
    Profile for Windows VMs in the container service cluster.
    """
    def __init__(__self__, *,
                 admin_password: str,
                 admin_username: str):
        """
        Profile for Windows VMs in the container service cluster.
        :param str admin_password: The administrator password to use for Windows VMs.
        :param str admin_username: The administrator username to use for Windows VMs.
        """
        pulumi.set(__self__, "admin_password", admin_password)
        pulumi.set(__self__, "admin_username", admin_username)

    @property
    @pulumi.getter(name="adminPassword")
    def admin_password(self) -> str:
        """
        The administrator password to use for Windows VMs.
        """
        return pulumi.get(self, "admin_password")

    @property
    @pulumi.getter(name="adminUsername")
    def admin_username(self) -> str:
        """
        The administrator username to use for Windows VMs.
        """
        return pulumi.get(self, "admin_username")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class KeyVaultSecretRefResponse(dict):
    """
    Reference to a secret stored in Azure Key Vault.
    """
    def __init__(__self__, *,
                 secret_name: str,
                 vault_id: str,
                 version: Optional[str] = None):
        """
        Reference to a secret stored in Azure Key Vault.
        :param str secret_name: The secret name.
        :param str vault_id: Key vault identifier.
        :param str version: The secret version.
        """
        pulumi.set(__self__, "secret_name", secret_name)
        pulumi.set(__self__, "vault_id", vault_id)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="secretName")
    def secret_name(self) -> str:
        """
        The secret name.
        """
        return pulumi.get(self, "secret_name")

    @property
    @pulumi.getter(name="vaultID")
    def vault_id(self) -> str:
        """
        Key vault identifier.
        """
        return pulumi.get(self, "vault_id")

    @property
    @pulumi.getter
    def version(self) -> Optional[str]:
        """
        The secret version.
        """
        return pulumi.get(self, "version")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

