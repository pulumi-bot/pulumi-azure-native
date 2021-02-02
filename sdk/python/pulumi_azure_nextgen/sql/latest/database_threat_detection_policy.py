# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from ._enums import *

__all__ = ['DatabaseThreatDetectionPolicy']


class DatabaseThreatDetectionPolicy(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 database_name: Optional[pulumi.Input[str]] = None,
                 disabled_alerts: Optional[pulumi.Input[str]] = None,
                 email_account_admins: Optional[pulumi.Input[Union[str, 'SecurityAlertPolicyEmailAccountAdmins']]] = None,
                 email_addresses: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 retention_days: Optional[pulumi.Input[int]] = None,
                 security_alert_policy_name: Optional[pulumi.Input[str]] = None,
                 server_name: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[Union[str, 'SecurityAlertPolicyState']]] = None,
                 storage_account_access_key: Optional[pulumi.Input[str]] = None,
                 storage_endpoint: Optional[pulumi.Input[str]] = None,
                 use_server_default: Optional[pulumi.Input[Union[str, 'SecurityAlertPolicyUseServerDefault']]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Contains information about a database Threat Detection policy.
        Latest API Version: 2014-04-01.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] database_name: The name of the database for which database Threat Detection policy is defined.
        :param pulumi.Input[str] disabled_alerts: Specifies the semicolon-separated list of alerts that are disabled, or empty string to disable no alerts. Possible values: Sql_Injection; Sql_Injection_Vulnerability; Access_Anomaly; Data_Exfiltration; Unsafe_Action.
        :param pulumi.Input[Union[str, 'SecurityAlertPolicyEmailAccountAdmins']] email_account_admins: Specifies that the alert is sent to the account administrators.
        :param pulumi.Input[str] email_addresses: Specifies the semicolon-separated list of e-mail addresses to which the alert is sent.
        :param pulumi.Input[str] location: The geo-location where the resource lives
        :param pulumi.Input[str] resource_group_name: The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        :param pulumi.Input[int] retention_days: Specifies the number of days to keep in the Threat Detection audit logs.
        :param pulumi.Input[str] security_alert_policy_name: The name of the security alert policy.
        :param pulumi.Input[str] server_name: The name of the server.
        :param pulumi.Input[Union[str, 'SecurityAlertPolicyState']] state: Specifies the state of the policy. If state is Enabled, storageEndpoint and storageAccountAccessKey are required.
        :param pulumi.Input[str] storage_account_access_key: Specifies the identifier key of the Threat Detection audit storage account. If state is Enabled, storageAccountAccessKey is required.
        :param pulumi.Input[str] storage_endpoint: Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs. If state is Enabled, storageEndpoint is required.
        :param pulumi.Input[Union[str, 'SecurityAlertPolicyUseServerDefault']] use_server_default: Specifies whether to use the default server policy.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if database_name is None and not opts.urn:
                raise TypeError("Missing required property 'database_name'")
            __props__['database_name'] = database_name
            __props__['disabled_alerts'] = disabled_alerts
            __props__['email_account_admins'] = email_account_admins
            __props__['email_addresses'] = email_addresses
            __props__['location'] = location
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['retention_days'] = retention_days
            if security_alert_policy_name is None and not opts.urn:
                raise TypeError("Missing required property 'security_alert_policy_name'")
            __props__['security_alert_policy_name'] = security_alert_policy_name
            if server_name is None and not opts.urn:
                raise TypeError("Missing required property 'server_name'")
            __props__['server_name'] = server_name
            if state is None and not opts.urn:
                raise TypeError("Missing required property 'state'")
            __props__['state'] = state
            __props__['storage_account_access_key'] = storage_account_access_key
            __props__['storage_endpoint'] = storage_endpoint
            __props__['use_server_default'] = use_server_default
            __props__['kind'] = None
            __props__['name'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:sql/v20140401:DatabaseThreatDetectionPolicy"), pulumi.Alias(type_="azure-nextgen:sql/v20180601preview:DatabaseThreatDetectionPolicy"), pulumi.Alias(type_="azure-nextgen:sql/v20200202preview:DatabaseThreatDetectionPolicy"), pulumi.Alias(type_="azure-nextgen:sql/v20200801preview:DatabaseThreatDetectionPolicy")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(DatabaseThreatDetectionPolicy, __self__).__init__(
            'azure-nextgen:sql/latest:DatabaseThreatDetectionPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'DatabaseThreatDetectionPolicy':
        """
        Get an existing DatabaseThreatDetectionPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return DatabaseThreatDetectionPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="disabledAlerts")
    def disabled_alerts(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the semicolon-separated list of alerts that are disabled, or empty string to disable no alerts. Possible values: Sql_Injection; Sql_Injection_Vulnerability; Access_Anomaly; Data_Exfiltration; Unsafe_Action.
        """
        return pulumi.get(self, "disabled_alerts")

    @property
    @pulumi.getter(name="emailAccountAdmins")
    def email_account_admins(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies that the alert is sent to the account administrators.
        """
        return pulumi.get(self, "email_account_admins")

    @property
    @pulumi.getter(name="emailAddresses")
    def email_addresses(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the semicolon-separated list of e-mail addresses to which the alert is sent.
        """
        return pulumi.get(self, "email_addresses")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[str]:
        """
        Resource kind.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[Optional[str]]:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="retentionDays")
    def retention_days(self) -> pulumi.Output[Optional[int]]:
        """
        Specifies the number of days to keep in the Threat Detection audit logs.
        """
        return pulumi.get(self, "retention_days")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        Specifies the state of the policy. If state is Enabled, storageEndpoint and storageAccountAccessKey are required.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="storageAccountAccessKey")
    def storage_account_access_key(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the identifier key of the Threat Detection audit storage account. If state is Enabled, storageAccountAccessKey is required.
        """
        return pulumi.get(self, "storage_account_access_key")

    @property
    @pulumi.getter(name="storageEndpoint")
    def storage_endpoint(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs. If state is Enabled, storageEndpoint is required.
        """
        return pulumi.get(self, "storage_endpoint")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Resource type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="useServerDefault")
    def use_server_default(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies whether to use the default server policy.
        """
        return pulumi.get(self, "use_server_default")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

