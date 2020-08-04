# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class ListDisasterRecoveryConfigKeysResult:
    """
    Namespace/ServiceBus Connection String
    """
    def __init__(__self__, alias_primary_connection_string=None, alias_secondary_connection_string=None, key_name=None, primary_connection_string=None, primary_key=None, secondary_connection_string=None, secondary_key=None):
        if alias_primary_connection_string and not isinstance(alias_primary_connection_string, str):
            raise TypeError("Expected argument 'alias_primary_connection_string' to be a str")
        __self__.alias_primary_connection_string = alias_primary_connection_string
        """
        Primary connection string of the alias if GEO DR is enabled
        """
        if alias_secondary_connection_string and not isinstance(alias_secondary_connection_string, str):
            raise TypeError("Expected argument 'alias_secondary_connection_string' to be a str")
        __self__.alias_secondary_connection_string = alias_secondary_connection_string
        """
        Secondary  connection string of the alias if GEO DR is enabled
        """
        if key_name and not isinstance(key_name, str):
            raise TypeError("Expected argument 'key_name' to be a str")
        __self__.key_name = key_name
        """
        A string that describes the authorization rule.
        """
        if primary_connection_string and not isinstance(primary_connection_string, str):
            raise TypeError("Expected argument 'primary_connection_string' to be a str")
        __self__.primary_connection_string = primary_connection_string
        """
        Primary connection string of the created namespace authorization rule.
        """
        if primary_key and not isinstance(primary_key, str):
            raise TypeError("Expected argument 'primary_key' to be a str")
        __self__.primary_key = primary_key
        """
        A base64-encoded 256-bit primary key for signing and validating the SAS token.
        """
        if secondary_connection_string and not isinstance(secondary_connection_string, str):
            raise TypeError("Expected argument 'secondary_connection_string' to be a str")
        __self__.secondary_connection_string = secondary_connection_string
        """
        Secondary connection string of the created namespace authorization rule.
        """
        if secondary_key and not isinstance(secondary_key, str):
            raise TypeError("Expected argument 'secondary_key' to be a str")
        __self__.secondary_key = secondary_key
        """
        A base64-encoded 256-bit primary key for signing and validating the SAS token.
        """


class AwaitableListDisasterRecoveryConfigKeysResult(ListDisasterRecoveryConfigKeysResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ListDisasterRecoveryConfigKeysResult(
            alias_primary_connection_string=self.alias_primary_connection_string,
            alias_secondary_connection_string=self.alias_secondary_connection_string,
            key_name=self.key_name,
            primary_connection_string=self.primary_connection_string,
            primary_key=self.primary_key,
            secondary_connection_string=self.secondary_connection_string,
            secondary_key=self.secondary_key)


def list_disaster_recovery_config_keys(alias=None, authorization_rule_name=None, namespace_name=None, resource_group_name=None, opts=None):
    """
    Use this data source to access information about an existing resource.

    :param str alias: The Disaster Recovery configuration name
    :param str authorization_rule_name: The authorization rule name.
    :param str namespace_name: The namespace name
    :param str resource_group_name: Name of the Resource group within the Azure subscription.
    """
    __args__ = dict()
    __args__['alias'] = alias
    __args__['authorizationRuleName'] = authorization_rule_name
    __args__['namespaceName'] = namespace_name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:servicebus/v20170401:listDisasterRecoveryConfigKeys', __args__, opts=opts).value

    return AwaitableListDisasterRecoveryConfigKeysResult(
        alias_primary_connection_string=__ret__.get('aliasPrimaryConnectionString'),
        alias_secondary_connection_string=__ret__.get('aliasSecondaryConnectionString'),
        key_name=__ret__.get('keyName'),
        primary_connection_string=__ret__.get('primaryConnectionString'),
        primary_key=__ret__.get('primaryKey'),
        secondary_connection_string=__ret__.get('secondaryConnectionString'),
        secondary_key=__ret__.get('secondaryKey'))
