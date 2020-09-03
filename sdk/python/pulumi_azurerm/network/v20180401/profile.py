# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Profile']


class Profile(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dns_config: Optional[pulumi.Input[pulumi.InputType['DnsConfigArgs']]] = None,
                 endpoints: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['EndpointArgs']]]]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 max_return: Optional[pulumi.Input[float]] = None,
                 monitor_config: Optional[pulumi.Input[pulumi.InputType['MonitorConfigArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 profile_name: Optional[pulumi.Input[str]] = None,
                 profile_status: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 traffic_routing_method: Optional[pulumi.Input[str]] = None,
                 traffic_view_enrollment_status: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Class representing a Traffic Manager profile.

        ## Example Usage
        ### Profile-PUT-NoEndpoints

        ```python
        import pulumi
        import pulumi_azurerm as azurerm

        profile = azurerm.network.v20180401.Profile("profile",
            dns_config={
                "relativeName": "azsmnet6386",
                "ttl": 35,
            },
            location="global",
            monitor_config={
                "path": "/testpath.aspx",
                "port": 80,
                "protocol": "HTTP",
            },
            profile_name="azsmnet6386",
            profile_status="Enabled",
            resource_group_name="azuresdkfornetautoresttrafficmanager1421",
            traffic_routing_method="Performance")

        ```
        ### Profile-PUT-WithCustomHeaders

        ```python
        import pulumi
        import pulumi_azurerm as azurerm

        profile = azurerm.network.v20180401.Profile("profile",
            dns_config={
                "relativeName": "azuresdkfornetautoresttrafficmanager6192",
                "ttl": 35,
            },
            endpoints=[{
                "name": "My external endpoint",
                "type": "Microsoft.network/TrafficManagerProfiles/ExternalEndpoints",
            }],
            location="global",
            monitor_config={
                "customHeaders": [
                    {
                        "name": "header-1",
                        "value": "value-1",
                    },
                    {
                        "name": "header-2",
                        "value": "value-2",
                    },
                ],
                "expectedStatusCodeRanges": [
                    {
                        "max": 205,
                        "min": 200,
                    },
                    {
                        "max": 410,
                        "min": 400,
                    },
                ],
                "intervalInSeconds": 10,
                "path": "/testpath.aspx",
                "port": 80,
                "protocol": "HTTP",
                "timeoutInSeconds": 5,
                "toleratedNumberOfFailures": 2,
            },
            profile_name="azuresdkfornetautoresttrafficmanager6192",
            profile_status="Enabled",
            resource_group_name="azuresdkfornetautoresttrafficmanager2583",
            traffic_routing_method="Performance",
            traffic_view_enrollment_status="Disabled")

        ```
        ### Profile-PUT-WithEndpoints

        ```python
        import pulumi
        import pulumi_azurerm as azurerm

        profile = azurerm.network.v20180401.Profile("profile",
            dns_config={
                "relativeName": "azuresdkfornetautoresttrafficmanager6192",
                "ttl": 35,
            },
            endpoints=[{
                "name": "My external endpoint",
                "type": "Microsoft.network/TrafficManagerProfiles/ExternalEndpoints",
            }],
            location="global",
            monitor_config={
                "intervalInSeconds": 10,
                "path": "/testpath.aspx",
                "port": 80,
                "protocol": "HTTP",
                "timeoutInSeconds": 5,
                "toleratedNumberOfFailures": 2,
            },
            profile_name="azuresdkfornetautoresttrafficmanager6192",
            profile_status="Enabled",
            resource_group_name="azuresdkfornetautoresttrafficmanager2583",
            traffic_routing_method="Performance")

        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['DnsConfigArgs']] dns_config: The DNS settings of the Traffic Manager profile.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['EndpointArgs']]]] endpoints: The list of endpoints in the Traffic Manager profile.
        :param pulumi.Input[str] id: Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficManagerProfiles/{resourceName}
        :param pulumi.Input[str] location: The Azure Region where the resource lives
        :param pulumi.Input[float] max_return: Maximum number of endpoints to be returned for MultiValue routing type.
        :param pulumi.Input[pulumi.InputType['MonitorConfigArgs']] monitor_config: The endpoint monitoring settings of the Traffic Manager profile.
        :param pulumi.Input[str] name: The name of the resource
        :param pulumi.Input[str] profile_name: The name of the Traffic Manager profile.
        :param pulumi.Input[str] profile_status: The status of the Traffic Manager profile.
        :param pulumi.Input[str] resource_group_name: The name of the resource group containing the Traffic Manager profile.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags.
        :param pulumi.Input[str] traffic_routing_method: The traffic routing method of the Traffic Manager profile.
        :param pulumi.Input[str] traffic_view_enrollment_status: Indicates whether Traffic View is 'Enabled' or 'Disabled' for the Traffic Manager profile. Null, indicates 'Disabled'. Enabling this feature will increase the cost of the Traffic Manage profile.
        :param pulumi.Input[str] type: The type of the resource. Ex- Microsoft.Network/trafficManagerProfiles.
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

            __props__['dns_config'] = dns_config
            __props__['endpoints'] = endpoints
            __props__['id'] = id
            __props__['location'] = location
            __props__['max_return'] = max_return
            __props__['monitor_config'] = monitor_config
            __props__['name'] = name
            if profile_name is None:
                raise TypeError("Missing required property 'profile_name'")
            __props__['profile_name'] = profile_name
            __props__['profile_status'] = profile_status
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['traffic_routing_method'] = traffic_routing_method
            __props__['traffic_view_enrollment_status'] = traffic_view_enrollment_status
            __props__['type'] = type
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azurerm:network/latest:Profile"), pulumi.Alias(type_="azurerm:network/v20151101:Profile"), pulumi.Alias(type_="azurerm:network/v20170301:Profile"), pulumi.Alias(type_="azurerm:network/v20170501:Profile"), pulumi.Alias(type_="azurerm:network/v20180201:Profile"), pulumi.Alias(type_="azurerm:network/v20180301:Profile")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Profile, __self__).__init__(
            'azurerm:network/v20180401:Profile',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Profile':
        """
        Get an existing Profile resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Profile(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dnsConfig")
    def dns_config(self) -> pulumi.Output[Optional['outputs.DnsConfigResponse']]:
        """
        The DNS settings of the Traffic Manager profile.
        """
        return pulumi.get(self, "dns_config")

    @property
    @pulumi.getter
    def endpoints(self) -> pulumi.Output[Optional[List['outputs.EndpointResponse']]]:
        """
        The list of endpoints in the Traffic Manager profile.
        """
        return pulumi.get(self, "endpoints")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[Optional[str]]:
        """
        The Azure Region where the resource lives
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="maxReturn")
    def max_return(self) -> pulumi.Output[Optional[float]]:
        """
        Maximum number of endpoints to be returned for MultiValue routing type.
        """
        return pulumi.get(self, "max_return")

    @property
    @pulumi.getter(name="monitorConfig")
    def monitor_config(self) -> pulumi.Output[Optional['outputs.MonitorConfigResponse']]:
        """
        The endpoint monitoring settings of the Traffic Manager profile.
        """
        return pulumi.get(self, "monitor_config")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="profileStatus")
    def profile_status(self) -> pulumi.Output[Optional[str]]:
        """
        The status of the Traffic Manager profile.
        """
        return pulumi.get(self, "profile_status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="trafficRoutingMethod")
    def traffic_routing_method(self) -> pulumi.Output[Optional[str]]:
        """
        The traffic routing method of the Traffic Manager profile.
        """
        return pulumi.get(self, "traffic_routing_method")

    @property
    @pulumi.getter(name="trafficViewEnrollmentStatus")
    def traffic_view_enrollment_status(self) -> pulumi.Output[Optional[str]]:
        """
        Indicates whether Traffic View is 'Enabled' or 'Disabled' for the Traffic Manager profile. Null, indicates 'Disabled'. Enabling this feature will increase the cost of the Traffic Manage profile.
        """
        return pulumi.get(self, "traffic_view_enrollment_status")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[Optional[str]]:
        """
        The type of the resource. Ex- Microsoft.Network/trafficManagerProfiles.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

