# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'APIServerProfileResponse',
    'ClusterProfileResponse',
    'ConsoleProfileResponse',
    'IngressProfileResponse',
    'MasterProfileResponse',
    'NetworkProfileResponse',
    'ServicePrincipalProfileResponse',
    'WorkerProfileResponse',
]

@pulumi.output_type
class APIServerProfileResponse(dict):
    """
    APIServerProfile represents an API server profile.
    """
    def __init__(__self__, *,
                 ip: Optional[str] = None,
                 url: Optional[str] = None,
                 visibility: Optional[str] = None):
        """
        APIServerProfile represents an API server profile.
        :param str ip: The IP of the cluster API server (immutable).
        :param str url: The URL to access the cluster API server (immutable).
        :param str visibility: API server visibility (immutable).
        """
        if ip is not None:
            pulumi.set(__self__, "ip", ip)
        if url is not None:
            pulumi.set(__self__, "url", url)
        if visibility is not None:
            pulumi.set(__self__, "visibility", visibility)

    @property
    @pulumi.getter
    def ip(self) -> Optional[str]:
        """
        The IP of the cluster API server (immutable).
        """
        return pulumi.get(self, "ip")

    @property
    @pulumi.getter
    def url(self) -> Optional[str]:
        """
        The URL to access the cluster API server (immutable).
        """
        return pulumi.get(self, "url")

    @property
    @pulumi.getter
    def visibility(self) -> Optional[str]:
        """
        API server visibility (immutable).
        """
        return pulumi.get(self, "visibility")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ClusterProfileResponse(dict):
    """
    ClusterProfile represents a cluster profile.
    """
    def __init__(__self__, *,
                 domain: Optional[str] = None,
                 pull_secret: Optional[str] = None,
                 resource_group_id: Optional[str] = None,
                 version: Optional[str] = None):
        """
        ClusterProfile represents a cluster profile.
        :param str domain: The domain for the cluster (immutable).
        :param str pull_secret: The pull secret for the cluster (immutable).
        :param str resource_group_id: The ID of the cluster resource group (immutable).
        :param str version: The version of the cluster (immutable).
        """
        if domain is not None:
            pulumi.set(__self__, "domain", domain)
        if pull_secret is not None:
            pulumi.set(__self__, "pull_secret", pull_secret)
        if resource_group_id is not None:
            pulumi.set(__self__, "resource_group_id", resource_group_id)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def domain(self) -> Optional[str]:
        """
        The domain for the cluster (immutable).
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter(name="pullSecret")
    def pull_secret(self) -> Optional[str]:
        """
        The pull secret for the cluster (immutable).
        """
        return pulumi.get(self, "pull_secret")

    @property
    @pulumi.getter(name="resourceGroupId")
    def resource_group_id(self) -> Optional[str]:
        """
        The ID of the cluster resource group (immutable).
        """
        return pulumi.get(self, "resource_group_id")

    @property
    @pulumi.getter
    def version(self) -> Optional[str]:
        """
        The version of the cluster (immutable).
        """
        return pulumi.get(self, "version")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ConsoleProfileResponse(dict):
    """
    ConsoleProfile represents a console profile.
    """
    def __init__(__self__, *,
                 url: Optional[str] = None):
        """
        ConsoleProfile represents a console profile.
        :param str url: The URL to access the cluster console (immutable).
        """
        if url is not None:
            pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter
    def url(self) -> Optional[str]:
        """
        The URL to access the cluster console (immutable).
        """
        return pulumi.get(self, "url")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class IngressProfileResponse(dict):
    """
    IngressProfile represents an ingress profile.
    """
    def __init__(__self__, *,
                 ip: Optional[str] = None,
                 name: Optional[str] = None,
                 visibility: Optional[str] = None):
        """
        IngressProfile represents an ingress profile.
        :param str ip: The IP of the ingress (immutable).
        :param str name: The ingress profile name.  Must be "default" (immutable).
        :param str visibility: Ingress visibility (immutable).
        """
        if ip is not None:
            pulumi.set(__self__, "ip", ip)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if visibility is not None:
            pulumi.set(__self__, "visibility", visibility)

    @property
    @pulumi.getter
    def ip(self) -> Optional[str]:
        """
        The IP of the ingress (immutable).
        """
        return pulumi.get(self, "ip")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The ingress profile name.  Must be "default" (immutable).
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def visibility(self) -> Optional[str]:
        """
        Ingress visibility (immutable).
        """
        return pulumi.get(self, "visibility")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class MasterProfileResponse(dict):
    """
    MasterProfile represents a master profile.
    """
    def __init__(__self__, *,
                 subnet_id: Optional[str] = None,
                 vm_size: Optional[str] = None):
        """
        MasterProfile represents a master profile.
        :param str subnet_id: The Azure resource ID of the master subnet (immutable).
        :param str vm_size: The size of the master VMs (immutable).
        """
        if subnet_id is not None:
            pulumi.set(__self__, "subnet_id", subnet_id)
        if vm_size is not None:
            pulumi.set(__self__, "vm_size", vm_size)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[str]:
        """
        The Azure resource ID of the master subnet (immutable).
        """
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter(name="vmSize")
    def vm_size(self) -> Optional[str]:
        """
        The size of the master VMs (immutable).
        """
        return pulumi.get(self, "vm_size")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class NetworkProfileResponse(dict):
    """
    NetworkProfile represents a network profile.
    """
    def __init__(__self__, *,
                 pod_cidr: Optional[str] = None,
                 service_cidr: Optional[str] = None):
        """
        NetworkProfile represents a network profile.
        :param str pod_cidr: The CIDR used for OpenShift/Kubernetes Pods (immutable).
        :param str service_cidr: The CIDR used for OpenShift/Kubernetes Services (immutable).
        """
        if pod_cidr is not None:
            pulumi.set(__self__, "pod_cidr", pod_cidr)
        if service_cidr is not None:
            pulumi.set(__self__, "service_cidr", service_cidr)

    @property
    @pulumi.getter(name="podCidr")
    def pod_cidr(self) -> Optional[str]:
        """
        The CIDR used for OpenShift/Kubernetes Pods (immutable).
        """
        return pulumi.get(self, "pod_cidr")

    @property
    @pulumi.getter(name="serviceCidr")
    def service_cidr(self) -> Optional[str]:
        """
        The CIDR used for OpenShift/Kubernetes Services (immutable).
        """
        return pulumi.get(self, "service_cidr")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ServicePrincipalProfileResponse(dict):
    """
    ServicePrincipalProfile represents a service principal profile.
    """
    def __init__(__self__, *,
                 client_id: Optional[str] = None,
                 client_secret: Optional[str] = None):
        """
        ServicePrincipalProfile represents a service principal profile.
        :param str client_id: The client ID used for the cluster (immutable).
        :param str client_secret: The client secret used for the cluster (immutable).
        """
        if client_id is not None:
            pulumi.set(__self__, "client_id", client_id)
        if client_secret is not None:
            pulumi.set(__self__, "client_secret", client_secret)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> Optional[str]:
        """
        The client ID used for the cluster (immutable).
        """
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter(name="clientSecret")
    def client_secret(self) -> Optional[str]:
        """
        The client secret used for the cluster (immutable).
        """
        return pulumi.get(self, "client_secret")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class WorkerProfileResponse(dict):
    """
    WorkerProfile represents a worker profile.
    """
    def __init__(__self__, *,
                 count: Optional[float] = None,
                 disk_size_gb: Optional[float] = None,
                 name: Optional[str] = None,
                 subnet_id: Optional[str] = None,
                 vm_size: Optional[str] = None):
        """
        WorkerProfile represents a worker profile.
        :param float count: The number of worker VMs.  Must be between 3 and 20 (immutable).
        :param float disk_size_gb: The disk size of the worker VMs.  Must be 128 or greater (immutable).
        :param str name: The worker profile name.  Must be "worker" (immutable).
        :param str subnet_id: The Azure resource ID of the worker subnet (immutable).
        :param str vm_size: The size of the worker VMs (immutable).
        """
        if count is not None:
            pulumi.set(__self__, "count", count)
        if disk_size_gb is not None:
            pulumi.set(__self__, "disk_size_gb", disk_size_gb)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if subnet_id is not None:
            pulumi.set(__self__, "subnet_id", subnet_id)
        if vm_size is not None:
            pulumi.set(__self__, "vm_size", vm_size)

    @property
    @pulumi.getter
    def count(self) -> Optional[float]:
        """
        The number of worker VMs.  Must be between 3 and 20 (immutable).
        """
        return pulumi.get(self, "count")

    @property
    @pulumi.getter(name="diskSizeGB")
    def disk_size_gb(self) -> Optional[float]:
        """
        The disk size of the worker VMs.  Must be 128 or greater (immutable).
        """
        return pulumi.get(self, "disk_size_gb")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The worker profile name.  Must be "worker" (immutable).
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[str]:
        """
        The Azure resource ID of the worker subnet (immutable).
        """
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter(name="vmSize")
    def vm_size(self) -> Optional[str]:
        """
        The size of the worker VMs (immutable).
        """
        return pulumi.get(self, "vm_size")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


