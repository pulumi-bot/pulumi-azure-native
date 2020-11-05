# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables

__all__ = [
    'ConnectedClusterAADProfileResponse',
    'ConnectedClusterIdentityResponse',
    'CredentialResultResponseResult',
]

@pulumi.output_type
class ConnectedClusterAADProfileResponse(dict):
    """
    AAD profile of the connected cluster
    """
    def __init__(__self__, *,
                 client_app_id: str,
                 server_app_id: str,
                 tenant_id: str):
        """
        AAD profile of the connected cluster
        :param str client_app_id: The client app id configured on target K8 cluster 
        :param str server_app_id: The server app id to access AD server
        :param str tenant_id: The aad tenant id which is configured on target K8s cluster
        """
        pulumi.set(__self__, "client_app_id", client_app_id)
        pulumi.set(__self__, "server_app_id", server_app_id)
        pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter(name="clientAppId")
    def client_app_id(self) -> str:
        """
        The client app id configured on target K8 cluster 
        """
        return pulumi.get(self, "client_app_id")

    @property
    @pulumi.getter(name="serverAppId")
    def server_app_id(self) -> str:
        """
        The server app id to access AD server
        """
        return pulumi.get(self, "server_app_id")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> str:
        """
        The aad tenant id which is configured on target K8s cluster
        """
        return pulumi.get(self, "tenant_id")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ConnectedClusterIdentityResponse(dict):
    """
    Identity for the connected cluster.
    """
    def __init__(__self__, *,
                 principal_id: str,
                 tenant_id: str,
                 type: str):
        """
        Identity for the connected cluster.
        :param str principal_id: The principal id of connected cluster identity. This property will only be provided for a system assigned identity.
        :param str tenant_id: The tenant id associated with the connected cluster. This property will only be provided for a system assigned identity.
        :param str type: The type of identity used for the connected cluster. The type 'SystemAssigned, includes a system created identity. The type 'None' means no identity is assigned to the connected cluster.
        """
        pulumi.set(__self__, "principal_id", principal_id)
        pulumi.set(__self__, "tenant_id", tenant_id)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> str:
        """
        The principal id of connected cluster identity. This property will only be provided for a system assigned identity.
        """
        return pulumi.get(self, "principal_id")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> str:
        """
        The tenant id associated with the connected cluster. This property will only be provided for a system assigned identity.
        """
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of identity used for the connected cluster. The type 'SystemAssigned, includes a system created identity. The type 'None' means no identity is assigned to the connected cluster.
        """
        return pulumi.get(self, "type")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class CredentialResultResponseResult(dict):
    """
    The credential result response.
    """
    def __init__(__self__, *,
                 name: str,
                 value: str):
        """
        The credential result response.
        :param str name: The name of the credential.
        :param str value: Base64-encoded Kubernetes configuration file.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the credential.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        Base64-encoded Kubernetes configuration file.
        """
        return pulumi.get(self, "value")


