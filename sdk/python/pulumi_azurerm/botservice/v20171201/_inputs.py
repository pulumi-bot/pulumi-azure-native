# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'BotPropertiesArgs',
    'ChannelArgs',
    'ConnectionSettingParameterArgs',
    'ConnectionSettingPropertiesArgs',
    'SkuArgs',
]

@pulumi.input_type
class BotPropertiesArgs:
    def __init__(__self__, *,
                 display_name: pulumi.Input[str],
                 endpoint: pulumi.Input[str],
                 msa_app_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 developer_app_insight_key: Optional[pulumi.Input[str]] = None,
                 developer_app_insights_api_key: Optional[pulumi.Input[str]] = None,
                 developer_app_insights_application_id: Optional[pulumi.Input[str]] = None,
                 icon_url: Optional[pulumi.Input[str]] = None,
                 luis_app_ids: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 luis_key: Optional[pulumi.Input[str]] = None):
        """
        The parameters to provide for the Bot.
        :param pulumi.Input[str] display_name: The Name of the bot
        :param pulumi.Input[str] endpoint: The bot's endpoint
        :param pulumi.Input[str] msa_app_id: Microsoft App Id for the bot
        :param pulumi.Input[str] description: The description of the bot
        :param pulumi.Input[str] developer_app_insight_key: The Application Insights key
        :param pulumi.Input[str] developer_app_insights_api_key: The Application Insights Api Key
        :param pulumi.Input[str] developer_app_insights_application_id: The Application Insights App Id
        :param pulumi.Input[str] icon_url: The Icon Url of the bot
        :param pulumi.Input[List[pulumi.Input[str]]] luis_app_ids: Collection of LUIS App Ids
        :param pulumi.Input[str] luis_key: The LUIS Key
        """
        pulumi.set(__self__, "display_name", display_name)
        pulumi.set(__self__, "endpoint", endpoint)
        pulumi.set(__self__, "msa_app_id", msa_app_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if developer_app_insight_key is not None:
            pulumi.set(__self__, "developer_app_insight_key", developer_app_insight_key)
        if developer_app_insights_api_key is not None:
            pulumi.set(__self__, "developer_app_insights_api_key", developer_app_insights_api_key)
        if developer_app_insights_application_id is not None:
            pulumi.set(__self__, "developer_app_insights_application_id", developer_app_insights_application_id)
        if icon_url is not None:
            pulumi.set(__self__, "icon_url", icon_url)
        if luis_app_ids is not None:
            pulumi.set(__self__, "luis_app_ids", luis_app_ids)
        if luis_key is not None:
            pulumi.set(__self__, "luis_key", luis_key)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Input[str]:
        """
        The Name of the bot
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def endpoint(self) -> pulumi.Input[str]:
        """
        The bot's endpoint
        """
        return pulumi.get(self, "endpoint")

    @endpoint.setter
    def endpoint(self, value: pulumi.Input[str]):
        pulumi.set(self, "endpoint", value)

    @property
    @pulumi.getter(name="msaAppId")
    def msa_app_id(self) -> pulumi.Input[str]:
        """
        Microsoft App Id for the bot
        """
        return pulumi.get(self, "msa_app_id")

    @msa_app_id.setter
    def msa_app_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "msa_app_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the bot
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="developerAppInsightKey")
    def developer_app_insight_key(self) -> Optional[pulumi.Input[str]]:
        """
        The Application Insights key
        """
        return pulumi.get(self, "developer_app_insight_key")

    @developer_app_insight_key.setter
    def developer_app_insight_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "developer_app_insight_key", value)

    @property
    @pulumi.getter(name="developerAppInsightsApiKey")
    def developer_app_insights_api_key(self) -> Optional[pulumi.Input[str]]:
        """
        The Application Insights Api Key
        """
        return pulumi.get(self, "developer_app_insights_api_key")

    @developer_app_insights_api_key.setter
    def developer_app_insights_api_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "developer_app_insights_api_key", value)

    @property
    @pulumi.getter(name="developerAppInsightsApplicationId")
    def developer_app_insights_application_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Application Insights App Id
        """
        return pulumi.get(self, "developer_app_insights_application_id")

    @developer_app_insights_application_id.setter
    def developer_app_insights_application_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "developer_app_insights_application_id", value)

    @property
    @pulumi.getter(name="iconUrl")
    def icon_url(self) -> Optional[pulumi.Input[str]]:
        """
        The Icon Url of the bot
        """
        return pulumi.get(self, "icon_url")

    @icon_url.setter
    def icon_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "icon_url", value)

    @property
    @pulumi.getter(name="luisAppIds")
    def luis_app_ids(self) -> Optional[pulumi.Input[List[pulumi.Input[str]]]]:
        """
        Collection of LUIS App Ids
        """
        return pulumi.get(self, "luis_app_ids")

    @luis_app_ids.setter
    def luis_app_ids(self, value: Optional[pulumi.Input[List[pulumi.Input[str]]]]):
        pulumi.set(self, "luis_app_ids", value)

    @property
    @pulumi.getter(name="luisKey")
    def luis_key(self) -> Optional[pulumi.Input[str]]:
        """
        The LUIS Key
        """
        return pulumi.get(self, "luis_key")

    @luis_key.setter
    def luis_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "luis_key", value)


@pulumi.input_type
class ChannelArgs:
    def __init__(__self__, *,
                 channel_name: pulumi.Input[str]):
        """
        Channel definition
        :param pulumi.Input[str] channel_name: The channel name
        """
        pulumi.set(__self__, "channel_name", channel_name)

    @property
    @pulumi.getter(name="channelName")
    def channel_name(self) -> pulumi.Input[str]:
        """
        The channel name
        """
        return pulumi.get(self, "channel_name")

    @channel_name.setter
    def channel_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "channel_name", value)


@pulumi.input_type
class ConnectionSettingParameterArgs:
    def __init__(__self__, *,
                 key: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None):
        """
        Extra Parameter in a Connection Setting Properties to indicate service provider specific properties
        :param pulumi.Input[str] key: Key for the Connection Setting Parameter.
        :param pulumi.Input[str] value: Value associated with the Connection Setting Parameter.
        """
        if key is not None:
            pulumi.set(__self__, "key", key)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        """
        Key for the Connection Setting Parameter.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        """
        Value associated with the Connection Setting Parameter.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class ConnectionSettingPropertiesArgs:
    def __init__(__self__, *,
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_secret: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[List[pulumi.Input['ConnectionSettingParameterArgs']]]] = None,
                 scopes: Optional[pulumi.Input[str]] = None,
                 service_provider_display_name: Optional[pulumi.Input[str]] = None,
                 service_provider_id: Optional[pulumi.Input[str]] = None):
        """
        Properties for a Connection Setting Item
        :param pulumi.Input[str] client_id: Client Id associated with the Connection Setting.
        :param pulumi.Input[str] client_secret: Client Secret associated with the Connection Setting
        :param pulumi.Input[List[pulumi.Input['ConnectionSettingParameterArgs']]] parameters: Service Provider Parameters associated with the Connection Setting
        :param pulumi.Input[str] scopes: Scopes associated with the Connection Setting
        :param pulumi.Input[str] service_provider_display_name: Service Provider Display Name associated with the Connection Setting
        :param pulumi.Input[str] service_provider_id: Service Provider Id associated with the Connection Setting
        """
        if client_id is not None:
            pulumi.set(__self__, "client_id", client_id)
        if client_secret is not None:
            pulumi.set(__self__, "client_secret", client_secret)
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)
        if scopes is not None:
            pulumi.set(__self__, "scopes", scopes)
        if service_provider_display_name is not None:
            pulumi.set(__self__, "service_provider_display_name", service_provider_display_name)
        if service_provider_id is not None:
            pulumi.set(__self__, "service_provider_id", service_provider_id)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> Optional[pulumi.Input[str]]:
        """
        Client Id associated with the Connection Setting.
        """
        return pulumi.get(self, "client_id")

    @client_id.setter
    def client_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_id", value)

    @property
    @pulumi.getter(name="clientSecret")
    def client_secret(self) -> Optional[pulumi.Input[str]]:
        """
        Client Secret associated with the Connection Setting
        """
        return pulumi.get(self, "client_secret")

    @client_secret.setter
    def client_secret(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_secret", value)

    @property
    @pulumi.getter
    def parameters(self) -> Optional[pulumi.Input[List[pulumi.Input['ConnectionSettingParameterArgs']]]]:
        """
        Service Provider Parameters associated with the Connection Setting
        """
        return pulumi.get(self, "parameters")

    @parameters.setter
    def parameters(self, value: Optional[pulumi.Input[List[pulumi.Input['ConnectionSettingParameterArgs']]]]):
        pulumi.set(self, "parameters", value)

    @property
    @pulumi.getter
    def scopes(self) -> Optional[pulumi.Input[str]]:
        """
        Scopes associated with the Connection Setting
        """
        return pulumi.get(self, "scopes")

    @scopes.setter
    def scopes(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scopes", value)

    @property
    @pulumi.getter(name="serviceProviderDisplayName")
    def service_provider_display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Service Provider Display Name associated with the Connection Setting
        """
        return pulumi.get(self, "service_provider_display_name")

    @service_provider_display_name.setter
    def service_provider_display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_provider_display_name", value)

    @property
    @pulumi.getter(name="serviceProviderId")
    def service_provider_id(self) -> Optional[pulumi.Input[str]]:
        """
        Service Provider Id associated with the Connection Setting
        """
        return pulumi.get(self, "service_provider_id")

    @service_provider_id.setter
    def service_provider_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_provider_id", value)


@pulumi.input_type
class SkuArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str]):
        """
        The SKU of the cognitive services account.
        :param pulumi.Input[str] name: The sku name
        """
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        The sku name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)


