# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from . import outputs
from ._enums import *

__all__ = [
    'AcsClusterPropertiesResponse',
    'AppInsightsCredentialsResponseResult',
    'AppInsightsPropertiesResponse',
    'AutoScaleConfigurationResponse',
    'ContainerRegistryCredentialsResponseResult',
    'ContainerRegistryPropertiesResponse',
    'ContainerServiceCredentialsResponseResult',
    'ErrorDetailResponse',
    'ErrorResponseResponse',
    'ErrorResponseWrapperResponse',
    'GlobalServiceConfigurationResponse',
    'KubernetesClusterPropertiesResponse',
    'ServiceAuthConfigurationResponse',
    'ServicePrincipalPropertiesResponse',
    'SslConfigurationResponse',
    'StorageAccountCredentialsResponseResult',
    'StorageAccountPropertiesResponse',
    'SystemServiceResponse',
]

@pulumi.output_type
class AcsClusterPropertiesResponse(dict):
    """
    Information about the container service backing the cluster
    """
    def __init__(__self__, *,
                 cluster_fqdn: str,
                 orchestrator_type: str,
                 agent_count: Optional[int] = None,
                 agent_vm_size: Optional[str] = None,
                 master_count: Optional[int] = None,
                 orchestrator_properties: Optional['outputs.KubernetesClusterPropertiesResponse'] = None,
                 system_services: Optional[Sequence['outputs.SystemServiceResponse']] = None):
        """
        Information about the container service backing the cluster
        :param str cluster_fqdn: The FQDN of the cluster. 
        :param str orchestrator_type: Type of orchestrator. It cannot be changed once the cluster is created.
        :param int agent_count: The number of agent nodes in the Container Service. This can be changed to scale the cluster.
        :param str agent_vm_size: The Azure VM size of the agent VM nodes. This cannot be changed once the cluster is created. This list is non exhaustive; refer to https://docs.microsoft.com/en-us/azure/virtual-machines/windows/sizes for the possible VM sizes.
        :param int master_count: The number of master nodes in the container service.
        :param 'KubernetesClusterPropertiesResponseArgs' orchestrator_properties: Orchestrator specific properties
        :param Sequence['SystemServiceResponseArgs'] system_services: The system services deployed to the cluster
        """
        pulumi.set(__self__, "cluster_fqdn", cluster_fqdn)
        pulumi.set(__self__, "orchestrator_type", orchestrator_type)
        if agent_count is None:
            agent_count = 2
        if agent_count is not None:
            pulumi.set(__self__, "agent_count", agent_count)
        if agent_vm_size is None:
            agent_vm_size = 'Standard_D3_v2'
        if agent_vm_size is not None:
            pulumi.set(__self__, "agent_vm_size", agent_vm_size)
        if master_count is None:
            master_count = 1
        if master_count is not None:
            pulumi.set(__self__, "master_count", master_count)
        if orchestrator_properties is not None:
            pulumi.set(__self__, "orchestrator_properties", orchestrator_properties)
        if system_services is not None:
            pulumi.set(__self__, "system_services", system_services)

    @property
    @pulumi.getter(name="clusterFqdn")
    def cluster_fqdn(self) -> str:
        """
        The FQDN of the cluster. 
        """
        return pulumi.get(self, "cluster_fqdn")

    @property
    @pulumi.getter(name="orchestratorType")
    def orchestrator_type(self) -> str:
        """
        Type of orchestrator. It cannot be changed once the cluster is created.
        """
        return pulumi.get(self, "orchestrator_type")

    @property
    @pulumi.getter(name="agentCount")
    def agent_count(self) -> Optional[int]:
        """
        The number of agent nodes in the Container Service. This can be changed to scale the cluster.
        """
        return pulumi.get(self, "agent_count")

    @property
    @pulumi.getter(name="agentVmSize")
    def agent_vm_size(self) -> Optional[str]:
        """
        The Azure VM size of the agent VM nodes. This cannot be changed once the cluster is created. This list is non exhaustive; refer to https://docs.microsoft.com/en-us/azure/virtual-machines/windows/sizes for the possible VM sizes.
        """
        return pulumi.get(self, "agent_vm_size")

    @property
    @pulumi.getter(name="masterCount")
    def master_count(self) -> Optional[int]:
        """
        The number of master nodes in the container service.
        """
        return pulumi.get(self, "master_count")

    @property
    @pulumi.getter(name="orchestratorProperties")
    def orchestrator_properties(self) -> Optional['outputs.KubernetesClusterPropertiesResponse']:
        """
        Orchestrator specific properties
        """
        return pulumi.get(self, "orchestrator_properties")

    @property
    @pulumi.getter(name="systemServices")
    def system_services(self) -> Optional[Sequence['outputs.SystemServiceResponse']]:
        """
        The system services deployed to the cluster
        """
        return pulumi.get(self, "system_services")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class AppInsightsCredentialsResponseResult(dict):
    """
    AppInsights credentials.
    """
    def __init__(__self__, *,
                 app_id: Optional[str] = None,
                 instrumentation_key: Optional[str] = None):
        """
        AppInsights credentials.
        :param str app_id: The AppInsights application ID.
        :param str instrumentation_key: The AppInsights instrumentation key. This is not returned in response of GET/PUT on the resource. To see this please call listKeys API.
        """
        if app_id is not None:
            pulumi.set(__self__, "app_id", app_id)
        if instrumentation_key is not None:
            pulumi.set(__self__, "instrumentation_key", instrumentation_key)

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> Optional[str]:
        """
        The AppInsights application ID.
        """
        return pulumi.get(self, "app_id")

    @property
    @pulumi.getter(name="instrumentationKey")
    def instrumentation_key(self) -> Optional[str]:
        """
        The AppInsights instrumentation key. This is not returned in response of GET/PUT on the resource. To see this please call listKeys API.
        """
        return pulumi.get(self, "instrumentation_key")


@pulumi.output_type
class AppInsightsPropertiesResponse(dict):
    """
    Properties of App Insights.
    """
    def __init__(__self__, *,
                 resource_id: Optional[str] = None):
        """
        Properties of App Insights.
        :param str resource_id: ARM resource ID of the App Insights.
        """
        if resource_id is not None:
            pulumi.set(__self__, "resource_id", resource_id)

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> Optional[str]:
        """
        ARM resource ID of the App Insights.
        """
        return pulumi.get(self, "resource_id")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class AutoScaleConfigurationResponse(dict):
    """
    AutoScale configuration properties.
    """
    def __init__(__self__, *,
                 max_replicas: Optional[int] = None,
                 min_replicas: Optional[int] = None,
                 refresh_period_in_seconds: Optional[int] = None,
                 status: Optional[str] = None,
                 target_utilization: Optional[float] = None):
        """
        AutoScale configuration properties.
        :param int max_replicas: The maximum number of replicas for each service.
        :param int min_replicas: The minimum number of replicas for each service.
        :param int refresh_period_in_seconds: Refresh period in seconds.
        :param str status: If auto-scale is enabled for all services. Each service can turn it off individually.
        :param float target_utilization: The target utilization.
        """
        if max_replicas is None:
            max_replicas = 100
        if max_replicas is not None:
            pulumi.set(__self__, "max_replicas", max_replicas)
        if min_replicas is None:
            min_replicas = 1
        if min_replicas is not None:
            pulumi.set(__self__, "min_replicas", min_replicas)
        if refresh_period_in_seconds is not None:
            pulumi.set(__self__, "refresh_period_in_seconds", refresh_period_in_seconds)
        if status is None:
            status = 'Disabled'
        if status is not None:
            pulumi.set(__self__, "status", status)
        if target_utilization is not None:
            pulumi.set(__self__, "target_utilization", target_utilization)

    @property
    @pulumi.getter(name="maxReplicas")
    def max_replicas(self) -> Optional[int]:
        """
        The maximum number of replicas for each service.
        """
        return pulumi.get(self, "max_replicas")

    @property
    @pulumi.getter(name="minReplicas")
    def min_replicas(self) -> Optional[int]:
        """
        The minimum number of replicas for each service.
        """
        return pulumi.get(self, "min_replicas")

    @property
    @pulumi.getter(name="refreshPeriodInSeconds")
    def refresh_period_in_seconds(self) -> Optional[int]:
        """
        Refresh period in seconds.
        """
        return pulumi.get(self, "refresh_period_in_seconds")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        If auto-scale is enabled for all services. Each service can turn it off individually.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="targetUtilization")
    def target_utilization(self) -> Optional[float]:
        """
        The target utilization.
        """
        return pulumi.get(self, "target_utilization")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ContainerRegistryCredentialsResponseResult(dict):
    """
    Information about the Azure Container Registry which contains the images deployed to the cluster.
    """
    def __init__(__self__, *,
                 login_server: str,
                 password: str,
                 password2: str,
                 username: str):
        """
        Information about the Azure Container Registry which contains the images deployed to the cluster.
        :param str login_server: The ACR login server name. User name is the first part of the FQDN.
        :param str password: The ACR primary password.
        :param str password2: The ACR secondary password.
        :param str username: The ACR login username.
        """
        pulumi.set(__self__, "login_server", login_server)
        pulumi.set(__self__, "password", password)
        pulumi.set(__self__, "password2", password2)
        pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter(name="loginServer")
    def login_server(self) -> str:
        """
        The ACR login server name. User name is the first part of the FQDN.
        """
        return pulumi.get(self, "login_server")

    @property
    @pulumi.getter
    def password(self) -> str:
        """
        The ACR primary password.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def password2(self) -> str:
        """
        The ACR secondary password.
        """
        return pulumi.get(self, "password2")

    @property
    @pulumi.getter
    def username(self) -> str:
        """
        The ACR login username.
        """
        return pulumi.get(self, "username")


@pulumi.output_type
class ContainerRegistryPropertiesResponse(dict):
    """
    Properties of Azure Container Registry.
    """
    def __init__(__self__, *,
                 resource_id: Optional[str] = None):
        """
        Properties of Azure Container Registry.
        :param str resource_id: ARM resource ID of the Azure Container Registry used to store Docker images for web services in the cluster. If not provided one will be created. This cannot be changed once the cluster is created.
        """
        if resource_id is not None:
            pulumi.set(__self__, "resource_id", resource_id)

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> Optional[str]:
        """
        ARM resource ID of the Azure Container Registry used to store Docker images for web services in the cluster. If not provided one will be created. This cannot be changed once the cluster is created.
        """
        return pulumi.get(self, "resource_id")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ContainerServiceCredentialsResponseResult(dict):
    """
    Information about the Azure Container Registry which contains the images deployed to the cluster.
    """
    def __init__(__self__, *,
                 acs_kube_config: str,
                 image_pull_secret_name: str,
                 service_principal_configuration: 'outputs.ServicePrincipalPropertiesResponse'):
        """
        Information about the Azure Container Registry which contains the images deployed to the cluster.
        :param str acs_kube_config: The ACS kube config file.
        :param str image_pull_secret_name: The ACR image pull secret name which was created in Kubernetes.
        :param 'ServicePrincipalPropertiesResponseArgs' service_principal_configuration: Service principal configuration used by Kubernetes.
        """
        pulumi.set(__self__, "acs_kube_config", acs_kube_config)
        pulumi.set(__self__, "image_pull_secret_name", image_pull_secret_name)
        pulumi.set(__self__, "service_principal_configuration", service_principal_configuration)

    @property
    @pulumi.getter(name="acsKubeConfig")
    def acs_kube_config(self) -> str:
        """
        The ACS kube config file.
        """
        return pulumi.get(self, "acs_kube_config")

    @property
    @pulumi.getter(name="imagePullSecretName")
    def image_pull_secret_name(self) -> str:
        """
        The ACR image pull secret name which was created in Kubernetes.
        """
        return pulumi.get(self, "image_pull_secret_name")

    @property
    @pulumi.getter(name="servicePrincipalConfiguration")
    def service_principal_configuration(self) -> 'outputs.ServicePrincipalPropertiesResponse':
        """
        Service principal configuration used by Kubernetes.
        """
        return pulumi.get(self, "service_principal_configuration")


@pulumi.output_type
class ErrorDetailResponse(dict):
    """
    Error detail information.
    """
    def __init__(__self__, *,
                 code: str,
                 message: str):
        """
        Error detail information.
        :param str code: Error code.
        :param str message: Error message.
        """
        pulumi.set(__self__, "code", code)
        pulumi.set(__self__, "message", message)

    @property
    @pulumi.getter
    def code(self) -> str:
        """
        Error code.
        """
        return pulumi.get(self, "code")

    @property
    @pulumi.getter
    def message(self) -> str:
        """
        Error message.
        """
        return pulumi.get(self, "message")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ErrorResponseResponse(dict):
    """
    Error response information.
    """
    def __init__(__self__, *,
                 code: str,
                 message: str,
                 details: Optional[Sequence['outputs.ErrorDetailResponse']] = None):
        """
        Error response information.
        :param str code: Error code.
        :param str message: Error message.
        :param Sequence['ErrorDetailResponseArgs'] details: An array of error detail objects.
        """
        pulumi.set(__self__, "code", code)
        pulumi.set(__self__, "message", message)
        if details is not None:
            pulumi.set(__self__, "details", details)

    @property
    @pulumi.getter
    def code(self) -> str:
        """
        Error code.
        """
        return pulumi.get(self, "code")

    @property
    @pulumi.getter
    def message(self) -> str:
        """
        Error message.
        """
        return pulumi.get(self, "message")

    @property
    @pulumi.getter
    def details(self) -> Optional[Sequence['outputs.ErrorDetailResponse']]:
        """
        An array of error detail objects.
        """
        return pulumi.get(self, "details")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ErrorResponseWrapperResponse(dict):
    """
    Wrapper for error response to follow ARM guidelines.
    """
    def __init__(__self__, *,
                 error: Optional['outputs.ErrorResponseResponse'] = None):
        """
        Wrapper for error response to follow ARM guidelines.
        :param 'ErrorResponseResponseArgs' error: The error response.
        """
        if error is not None:
            pulumi.set(__self__, "error", error)

    @property
    @pulumi.getter
    def error(self) -> Optional['outputs.ErrorResponseResponse']:
        """
        The error response.
        """
        return pulumi.get(self, "error")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GlobalServiceConfigurationResponse(dict):
    """
    Global configuration for services in the cluster.
    """
    def __init__(__self__, *,
                 auto_scale: Optional['outputs.AutoScaleConfigurationResponse'] = None,
                 etag: Optional[str] = None,
                 service_auth: Optional['outputs.ServiceAuthConfigurationResponse'] = None,
                 ssl: Optional['outputs.SslConfigurationResponse'] = None):
        """
        Global configuration for services in the cluster.
        :param 'AutoScaleConfigurationResponseArgs' auto_scale: The auto-scale configuration
        :param str etag: The configuration ETag for updates.
        :param 'ServiceAuthConfigurationResponseArgs' service_auth: Optional global authorization keys for all user services deployed in cluster. These are used if the service does not have auth keys.
        :param 'SslConfigurationResponseArgs' ssl: The SSL configuration properties
        """
        if auto_scale is not None:
            pulumi.set(__self__, "auto_scale", auto_scale)
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if service_auth is not None:
            pulumi.set(__self__, "service_auth", service_auth)
        if ssl is not None:
            pulumi.set(__self__, "ssl", ssl)

    @property
    @pulumi.getter(name="autoScale")
    def auto_scale(self) -> Optional['outputs.AutoScaleConfigurationResponse']:
        """
        The auto-scale configuration
        """
        return pulumi.get(self, "auto_scale")

    @property
    @pulumi.getter
    def etag(self) -> Optional[str]:
        """
        The configuration ETag for updates.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="serviceAuth")
    def service_auth(self) -> Optional['outputs.ServiceAuthConfigurationResponse']:
        """
        Optional global authorization keys for all user services deployed in cluster. These are used if the service does not have auth keys.
        """
        return pulumi.get(self, "service_auth")

    @property
    @pulumi.getter
    def ssl(self) -> Optional['outputs.SslConfigurationResponse']:
        """
        The SSL configuration properties
        """
        return pulumi.get(self, "ssl")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class KubernetesClusterPropertiesResponse(dict):
    """
    Kubernetes cluster specific properties
    """
    def __init__(__self__, *,
                 service_principal: Optional['outputs.ServicePrincipalPropertiesResponse'] = None):
        """
        Kubernetes cluster specific properties
        :param 'ServicePrincipalPropertiesResponseArgs' service_principal: The Azure Service Principal used by Kubernetes
        """
        if service_principal is not None:
            pulumi.set(__self__, "service_principal", service_principal)

    @property
    @pulumi.getter(name="servicePrincipal")
    def service_principal(self) -> Optional['outputs.ServicePrincipalPropertiesResponse']:
        """
        The Azure Service Principal used by Kubernetes
        """
        return pulumi.get(self, "service_principal")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ServiceAuthConfigurationResponse(dict):
    """
    Global service auth configuration properties. These are the data-plane authorization keys and are used if a service doesn't define it's own.
    """
    def __init__(__self__, *,
                 primary_auth_key_hash: str,
                 secondary_auth_key_hash: str):
        """
        Global service auth configuration properties. These are the data-plane authorization keys and are used if a service doesn't define it's own.
        :param str primary_auth_key_hash: The primary auth key hash. This is not returned in response of GET/PUT on the resource.. To see this please call listKeys API.
        :param str secondary_auth_key_hash: The secondary auth key hash. This is not returned in response of GET/PUT on the resource.. To see this please call listKeys API.
        """
        pulumi.set(__self__, "primary_auth_key_hash", primary_auth_key_hash)
        pulumi.set(__self__, "secondary_auth_key_hash", secondary_auth_key_hash)

    @property
    @pulumi.getter(name="primaryAuthKeyHash")
    def primary_auth_key_hash(self) -> str:
        """
        The primary auth key hash. This is not returned in response of GET/PUT on the resource.. To see this please call listKeys API.
        """
        return pulumi.get(self, "primary_auth_key_hash")

    @property
    @pulumi.getter(name="secondaryAuthKeyHash")
    def secondary_auth_key_hash(self) -> str:
        """
        The secondary auth key hash. This is not returned in response of GET/PUT on the resource.. To see this please call listKeys API.
        """
        return pulumi.get(self, "secondary_auth_key_hash")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ServicePrincipalPropertiesResponse(dict):
    """
    The Azure service principal used by Kubernetes for configuring load balancers
    """
    def __init__(__self__, *,
                 client_id: str,
                 secret: str):
        """
        The Azure service principal used by Kubernetes for configuring load balancers
        :param str client_id: The service principal client ID
        :param str secret: The service principal secret. This is not returned in response of GET/PUT on the resource. To see this please call listKeys.
        """
        pulumi.set(__self__, "client_id", client_id)
        pulumi.set(__self__, "secret", secret)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> str:
        """
        The service principal client ID
        """
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter
    def secret(self) -> str:
        """
        The service principal secret. This is not returned in response of GET/PUT on the resource. To see this please call listKeys.
        """
        return pulumi.get(self, "secret")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class SslConfigurationResponse(dict):
    """
    SSL configuration. If configured data-plane calls to user services will be exposed over SSL only.
    """
    def __init__(__self__, *,
                 cert: Optional[str] = None,
                 cname: Optional[str] = None,
                 key: Optional[str] = None,
                 status: Optional[str] = None):
        """
        SSL configuration. If configured data-plane calls to user services will be exposed over SSL only.
        :param str cert: The SSL cert data in PEM format.
        :param str cname: The CName of the certificate.
        :param str key: The SSL key data in PEM format. This is not returned in response of GET/PUT on the resource. To see this please call listKeys API.
        :param str status: SSL status. Allowed values are Enabled and Disabled.
        """
        if cert is not None:
            pulumi.set(__self__, "cert", cert)
        if cname is not None:
            pulumi.set(__self__, "cname", cname)
        if key is not None:
            pulumi.set(__self__, "key", key)
        if status is None:
            status = 'Enabled'
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def cert(self) -> Optional[str]:
        """
        The SSL cert data in PEM format.
        """
        return pulumi.get(self, "cert")

    @property
    @pulumi.getter
    def cname(self) -> Optional[str]:
        """
        The CName of the certificate.
        """
        return pulumi.get(self, "cname")

    @property
    @pulumi.getter
    def key(self) -> Optional[str]:
        """
        The SSL key data in PEM format. This is not returned in response of GET/PUT on the resource. To see this please call listKeys API.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        SSL status. Allowed values are Enabled and Disabled.
        """
        return pulumi.get(self, "status")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class StorageAccountCredentialsResponseResult(dict):
    """
    Access information for the storage account.
    """
    def __init__(__self__, *,
                 primary_key: str,
                 resource_id: str,
                 secondary_key: str):
        """
        Access information for the storage account.
        :param str primary_key: The primary key of the storage account.
        :param str resource_id: The ARM resource ID of the storage account.
        :param str secondary_key: The secondary key of the storage account.
        """
        pulumi.set(__self__, "primary_key", primary_key)
        pulumi.set(__self__, "resource_id", resource_id)
        pulumi.set(__self__, "secondary_key", secondary_key)

    @property
    @pulumi.getter(name="primaryKey")
    def primary_key(self) -> str:
        """
        The primary key of the storage account.
        """
        return pulumi.get(self, "primary_key")

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> str:
        """
        The ARM resource ID of the storage account.
        """
        return pulumi.get(self, "resource_id")

    @property
    @pulumi.getter(name="secondaryKey")
    def secondary_key(self) -> str:
        """
        The secondary key of the storage account.
        """
        return pulumi.get(self, "secondary_key")


@pulumi.output_type
class StorageAccountPropertiesResponse(dict):
    """
    Properties of Storage Account.
    """
    def __init__(__self__, *,
                 resource_id: Optional[str] = None):
        """
        Properties of Storage Account.
        :param str resource_id: ARM resource ID of the Azure Storage Account to store CLI specific files. If not provided one will be created. This cannot be changed once the cluster is created.
        """
        if resource_id is not None:
            pulumi.set(__self__, "resource_id", resource_id)

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> Optional[str]:
        """
        ARM resource ID of the Azure Storage Account to store CLI specific files. If not provided one will be created. This cannot be changed once the cluster is created.
        """
        return pulumi.get(self, "resource_id")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class SystemServiceResponse(dict):
    """
    Information about a system service deployed in the cluster
    """
    def __init__(__self__, *,
                 public_ip_address: str,
                 system_service_type: str,
                 version: str):
        """
        Information about a system service deployed in the cluster
        :param str public_ip_address: The public IP address of the system service
        :param str system_service_type: The system service type
        :param str version: The state of the system service
        """
        pulumi.set(__self__, "public_ip_address", public_ip_address)
        pulumi.set(__self__, "system_service_type", system_service_type)
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="publicIpAddress")
    def public_ip_address(self) -> str:
        """
        The public IP address of the system service
        """
        return pulumi.get(self, "public_ip_address")

    @property
    @pulumi.getter(name="systemServiceType")
    def system_service_type(self) -> str:
        """
        The system service type
        """
        return pulumi.get(self, "system_service_type")

    @property
    @pulumi.getter
    def version(self) -> str:
        """
        The state of the system service
        """
        return pulumi.get(self, "version")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


