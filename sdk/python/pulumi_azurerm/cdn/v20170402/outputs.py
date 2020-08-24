# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'DeepCreatedOriginResponse',
    'GeoFilterResponse',
    'SkuResponse',
]

@pulumi.output_type
class DeepCreatedOriginResponse(dict):
    """
    The main origin of CDN content which is added when creating a CDN endpoint.
    """
    def __init__(__self__, *,
                 host_name: str,
                 name: str,
                 http_port: Optional[float] = None,
                 https_port: Optional[float] = None):
        """
        The main origin of CDN content which is added when creating a CDN endpoint.
        :param str host_name: The address of the origin. It can be a domain name, IPv4 address, or IPv6 address.
        :param str name: Origin name
        :param float http_port: The value of the HTTP port. Must be between 1 and 65535
        :param float https_port: The value of the HTTPS port. Must be between 1 and 65535
        """
        pulumi.set(__self__, "host_name", host_name)
        pulumi.set(__self__, "name", name)
        if http_port is not None:
            pulumi.set(__self__, "http_port", http_port)
        if https_port is not None:
            pulumi.set(__self__, "https_port", https_port)

    @property
    @pulumi.getter(name="hostName")
    def host_name(self) -> str:
        """
        The address of the origin. It can be a domain name, IPv4 address, or IPv6 address.
        """
        return pulumi.get(self, "host_name")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Origin name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="httpPort")
    def http_port(self) -> Optional[float]:
        """
        The value of the HTTP port. Must be between 1 and 65535
        """
        return pulumi.get(self, "http_port")

    @property
    @pulumi.getter(name="httpsPort")
    def https_port(self) -> Optional[float]:
        """
        The value of the HTTPS port. Must be between 1 and 65535
        """
        return pulumi.get(self, "https_port")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GeoFilterResponse(dict):
    """
    Rules defining user's geo access within a CDN endpoint.
    """
    def __init__(__self__, *,
                 action: str,
                 country_codes: List[str],
                 relative_path: str):
        """
        Rules defining user's geo access within a CDN endpoint.
        :param str action: Action of the geo filter, i.e. allow or block access.
        :param List[str] country_codes: Two letter country codes defining user country access in a geo filter, e.g. AU, MX, US.
        :param str relative_path: Relative path applicable to geo filter. (e.g. '/mypictures', '/mypicture/kitty.jpg', and etc.)
        """
        pulumi.set(__self__, "action", action)
        pulumi.set(__self__, "country_codes", country_codes)
        pulumi.set(__self__, "relative_path", relative_path)

    @property
    @pulumi.getter
    def action(self) -> str:
        """
        Action of the geo filter, i.e. allow or block access.
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter(name="countryCodes")
    def country_codes(self) -> List[str]:
        """
        Two letter country codes defining user country access in a geo filter, e.g. AU, MX, US.
        """
        return pulumi.get(self, "country_codes")

    @property
    @pulumi.getter(name="relativePath")
    def relative_path(self) -> str:
        """
        Relative path applicable to geo filter. (e.g. '/mypictures', '/mypicture/kitty.jpg', and etc.)
        """
        return pulumi.get(self, "relative_path")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class SkuResponse(dict):
    """
    The pricing tier (defines a CDN provider, feature list and rate) of the CDN profile.
    """
    def __init__(__self__, *,
                 name: Optional[str] = None):
        """
        The pricing tier (defines a CDN provider, feature list and rate) of the CDN profile.
        :param str name: Name of the pricing tier.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Name of the pricing tier.
        """
        return pulumi.get(self, "name")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


