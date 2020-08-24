# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'IdentitySourceArgs',
    'ManagementClusterArgs',
    'SkuArgs',
]

@pulumi.input_type
class IdentitySourceArgs:
    def __init__(__self__, *,
                 alias: Optional[pulumi.Input[str]] = None,
                 base_group_dn: Optional[pulumi.Input[str]] = None,
                 base_user_dn: Optional[pulumi.Input[str]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 primary_server: Optional[pulumi.Input[str]] = None,
                 secondary_server: Optional[pulumi.Input[str]] = None,
                 ssl: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None):
        """
        vCenter Single Sign On Identity Source
        :param pulumi.Input[str] alias: The domain's NetBIOS name
        :param pulumi.Input[str] base_group_dn: The base distinguished name for groups
        :param pulumi.Input[str] base_user_dn: The base distinguished name for users
        :param pulumi.Input[str] domain: The domain's dns name
        :param pulumi.Input[str] name: The name of the identity source
        :param pulumi.Input[str] password: The password of the Active Directory user with a minimum of read-only access to Base DN for users and groups.
        :param pulumi.Input[str] primary_server: Primary server URL
        :param pulumi.Input[str] secondary_server: Secondary server URL
        :param pulumi.Input[str] ssl: Protect LDAP communication using SSL certificate (LDAPS)
        :param pulumi.Input[str] username: The ID of an Active Directory user with a minimum of read-only access to Base DN for users and group
        """
        if alias is not None:
            pulumi.set(__self__, "alias", alias)
        if base_group_dn is not None:
            pulumi.set(__self__, "base_group_dn", base_group_dn)
        if base_user_dn is not None:
            pulumi.set(__self__, "base_user_dn", base_user_dn)
        if domain is not None:
            pulumi.set(__self__, "domain", domain)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if primary_server is not None:
            pulumi.set(__self__, "primary_server", primary_server)
        if secondary_server is not None:
            pulumi.set(__self__, "secondary_server", secondary_server)
        if ssl is not None:
            pulumi.set(__self__, "ssl", ssl)
        if username is not None:
            pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter
    def alias(self) -> Optional[pulumi.Input[str]]:
        """
        The domain's NetBIOS name
        """
        return pulumi.get(self, "alias")

    @alias.setter
    def alias(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "alias", value)

    @property
    @pulumi.getter(name="baseGroupDN")
    def base_group_dn(self) -> Optional[pulumi.Input[str]]:
        """
        The base distinguished name for groups
        """
        return pulumi.get(self, "base_group_dn")

    @base_group_dn.setter
    def base_group_dn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "base_group_dn", value)

    @property
    @pulumi.getter(name="baseUserDN")
    def base_user_dn(self) -> Optional[pulumi.Input[str]]:
        """
        The base distinguished name for users
        """
        return pulumi.get(self, "base_user_dn")

    @base_user_dn.setter
    def base_user_dn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "base_user_dn", value)

    @property
    @pulumi.getter
    def domain(self) -> Optional[pulumi.Input[str]]:
        """
        The domain's dns name
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the identity source
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        The password of the Active Directory user with a minimum of read-only access to Base DN for users and groups.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="primaryServer")
    def primary_server(self) -> Optional[pulumi.Input[str]]:
        """
        Primary server URL
        """
        return pulumi.get(self, "primary_server")

    @primary_server.setter
    def primary_server(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "primary_server", value)

    @property
    @pulumi.getter(name="secondaryServer")
    def secondary_server(self) -> Optional[pulumi.Input[str]]:
        """
        Secondary server URL
        """
        return pulumi.get(self, "secondary_server")

    @secondary_server.setter
    def secondary_server(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secondary_server", value)

    @property
    @pulumi.getter
    def ssl(self) -> Optional[pulumi.Input[str]]:
        """
        Protect LDAP communication using SSL certificate (LDAPS)
        """
        return pulumi.get(self, "ssl")

    @ssl.setter
    def ssl(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssl", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of an Active Directory user with a minimum of read-only access to Base DN for users and group
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)


@pulumi.input_type
class ManagementClusterArgs:
    def __init__(__self__, *,
                 cluster_size: pulumi.Input[float]):
        """
        The properties of a default cluster
        :param pulumi.Input[float] cluster_size: The cluster size
        """
        pulumi.set(__self__, "cluster_size", cluster_size)

    @property
    @pulumi.getter(name="clusterSize")
    def cluster_size(self) -> pulumi.Input[float]:
        """
        The cluster size
        """
        return pulumi.get(self, "cluster_size")

    @cluster_size.setter
    def cluster_size(self, value: pulumi.Input[float]):
        pulumi.set(self, "cluster_size", value)


@pulumi.input_type
class SkuArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str]):
        """
        The resource model definition representing SKU
        :param pulumi.Input[str] name: The name of the SKU.
        """
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        The name of the SKU.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)


