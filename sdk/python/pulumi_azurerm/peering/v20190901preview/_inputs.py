# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables

__all__ = [
    'BgpSessionArgs',
    'ContactInfoArgs',
    'DirectConnectionArgs',
    'ExchangeConnectionArgs',
    'PeeringPropertiesDirectArgs',
    'PeeringPropertiesExchangeArgs',
    'PeeringSkuArgs',
    'SubResourceArgs',
]

@pulumi.input_type
class BgpSessionArgs:
    def __init__(__self__, *,
                 max_prefixes_advertised_v4: Optional[pulumi.Input[float]] = None,
                 max_prefixes_advertised_v6: Optional[pulumi.Input[float]] = None,
                 md5_authentication_key: Optional[pulumi.Input[str]] = None,
                 peer_session_i_pv4_address: Optional[pulumi.Input[str]] = None,
                 peer_session_i_pv6_address: Optional[pulumi.Input[str]] = None,
                 session_prefix_v4: Optional[pulumi.Input[str]] = None,
                 session_prefix_v6: Optional[pulumi.Input[str]] = None):
        """
        The properties that define a BGP session.
        :param pulumi.Input[float] max_prefixes_advertised_v4: The maximum number of prefixes advertised over the IPv4 session.
        :param pulumi.Input[float] max_prefixes_advertised_v6: The maximum number of prefixes advertised over the IPv6 session.
        :param pulumi.Input[str] md5_authentication_key: The MD5 authentication key of the session.
        :param pulumi.Input[str] peer_session_i_pv4_address: The IPv4 session address on peer's end.
        :param pulumi.Input[str] peer_session_i_pv6_address: The IPv6 session address on peer's end.
        :param pulumi.Input[str] session_prefix_v4: The IPv4 prefix that contains both ends' IPv4 addresses.
        :param pulumi.Input[str] session_prefix_v6: The IPv6 prefix that contains both ends' IPv6 addresses.
        """
        if max_prefixes_advertised_v4 is not None:
            pulumi.set(__self__, "max_prefixes_advertised_v4", max_prefixes_advertised_v4)
        if max_prefixes_advertised_v6 is not None:
            pulumi.set(__self__, "max_prefixes_advertised_v6", max_prefixes_advertised_v6)
        if md5_authentication_key is not None:
            pulumi.set(__self__, "md5_authentication_key", md5_authentication_key)
        if peer_session_i_pv4_address is not None:
            pulumi.set(__self__, "peer_session_i_pv4_address", peer_session_i_pv4_address)
        if peer_session_i_pv6_address is not None:
            pulumi.set(__self__, "peer_session_i_pv6_address", peer_session_i_pv6_address)
        if session_prefix_v4 is not None:
            pulumi.set(__self__, "session_prefix_v4", session_prefix_v4)
        if session_prefix_v6 is not None:
            pulumi.set(__self__, "session_prefix_v6", session_prefix_v6)

    @property
    @pulumi.getter(name="maxPrefixesAdvertisedV4")
    def max_prefixes_advertised_v4(self) -> Optional[pulumi.Input[float]]:
        """
        The maximum number of prefixes advertised over the IPv4 session.
        """
        return pulumi.get(self, "max_prefixes_advertised_v4")

    @max_prefixes_advertised_v4.setter
    def max_prefixes_advertised_v4(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "max_prefixes_advertised_v4", value)

    @property
    @pulumi.getter(name="maxPrefixesAdvertisedV6")
    def max_prefixes_advertised_v6(self) -> Optional[pulumi.Input[float]]:
        """
        The maximum number of prefixes advertised over the IPv6 session.
        """
        return pulumi.get(self, "max_prefixes_advertised_v6")

    @max_prefixes_advertised_v6.setter
    def max_prefixes_advertised_v6(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "max_prefixes_advertised_v6", value)

    @property
    @pulumi.getter(name="md5AuthenticationKey")
    def md5_authentication_key(self) -> Optional[pulumi.Input[str]]:
        """
        The MD5 authentication key of the session.
        """
        return pulumi.get(self, "md5_authentication_key")

    @md5_authentication_key.setter
    def md5_authentication_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "md5_authentication_key", value)

    @property
    @pulumi.getter(name="peerSessionIPv4Address")
    def peer_session_i_pv4_address(self) -> Optional[pulumi.Input[str]]:
        """
        The IPv4 session address on peer's end.
        """
        return pulumi.get(self, "peer_session_i_pv4_address")

    @peer_session_i_pv4_address.setter
    def peer_session_i_pv4_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peer_session_i_pv4_address", value)

    @property
    @pulumi.getter(name="peerSessionIPv6Address")
    def peer_session_i_pv6_address(self) -> Optional[pulumi.Input[str]]:
        """
        The IPv6 session address on peer's end.
        """
        return pulumi.get(self, "peer_session_i_pv6_address")

    @peer_session_i_pv6_address.setter
    def peer_session_i_pv6_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peer_session_i_pv6_address", value)

    @property
    @pulumi.getter(name="sessionPrefixV4")
    def session_prefix_v4(self) -> Optional[pulumi.Input[str]]:
        """
        The IPv4 prefix that contains both ends' IPv4 addresses.
        """
        return pulumi.get(self, "session_prefix_v4")

    @session_prefix_v4.setter
    def session_prefix_v4(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "session_prefix_v4", value)

    @property
    @pulumi.getter(name="sessionPrefixV6")
    def session_prefix_v6(self) -> Optional[pulumi.Input[str]]:
        """
        The IPv6 prefix that contains both ends' IPv6 addresses.
        """
        return pulumi.get(self, "session_prefix_v6")

    @session_prefix_v6.setter
    def session_prefix_v6(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "session_prefix_v6", value)


@pulumi.input_type
class ContactInfoArgs:
    def __init__(__self__, *,
                 emails: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 phone: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None):
        """
        The contact information of the peer.
        :param pulumi.Input[List[pulumi.Input[str]]] emails: The list of email addresses.
        :param pulumi.Input[List[pulumi.Input[str]]] phone: The list of contact numbers.
        """
        if emails is not None:
            pulumi.set(__self__, "emails", emails)
        if phone is not None:
            pulumi.set(__self__, "phone", phone)

    @property
    @pulumi.getter
    def emails(self) -> Optional[pulumi.Input[List[pulumi.Input[str]]]]:
        """
        The list of email addresses.
        """
        return pulumi.get(self, "emails")

    @emails.setter
    def emails(self, value: Optional[pulumi.Input[List[pulumi.Input[str]]]]):
        pulumi.set(self, "emails", value)

    @property
    @pulumi.getter
    def phone(self) -> Optional[pulumi.Input[List[pulumi.Input[str]]]]:
        """
        The list of contact numbers.
        """
        return pulumi.get(self, "phone")

    @phone.setter
    def phone(self, value: Optional[pulumi.Input[List[pulumi.Input[str]]]]):
        pulumi.set(self, "phone", value)


@pulumi.input_type
class DirectConnectionArgs:
    def __init__(__self__, *,
                 bandwidth_in_mbps: Optional[pulumi.Input[float]] = None,
                 bgp_session: Optional[pulumi.Input['BgpSessionArgs']] = None,
                 connection_identifier: Optional[pulumi.Input[str]] = None,
                 peering_db_facility_id: Optional[pulumi.Input[float]] = None,
                 session_address_provider: Optional[pulumi.Input[str]] = None,
                 use_for_peering_service: Optional[pulumi.Input[bool]] = None):
        """
        The properties that define a direct connection.
        :param pulumi.Input[float] bandwidth_in_mbps: The bandwidth of the connection.
        :param pulumi.Input['BgpSessionArgs'] bgp_session: The BGP session associated with the connection.
        :param pulumi.Input[str] connection_identifier: The unique identifier (GUID) for the connection.
        :param pulumi.Input[float] peering_db_facility_id: The PeeringDB.com ID of the facility at which the connection has to be set up.
        :param pulumi.Input[str] session_address_provider: The field indicating if Microsoft provides session ip addresses.
        :param pulumi.Input[bool] use_for_peering_service: The flag that indicates whether or not the connection is used for peering service.
        """
        if bandwidth_in_mbps is not None:
            pulumi.set(__self__, "bandwidth_in_mbps", bandwidth_in_mbps)
        if bgp_session is not None:
            pulumi.set(__self__, "bgp_session", bgp_session)
        if connection_identifier is not None:
            pulumi.set(__self__, "connection_identifier", connection_identifier)
        if peering_db_facility_id is not None:
            pulumi.set(__self__, "peering_db_facility_id", peering_db_facility_id)
        if session_address_provider is not None:
            pulumi.set(__self__, "session_address_provider", session_address_provider)
        if use_for_peering_service is not None:
            pulumi.set(__self__, "use_for_peering_service", use_for_peering_service)

    @property
    @pulumi.getter(name="bandwidthInMbps")
    def bandwidth_in_mbps(self) -> Optional[pulumi.Input[float]]:
        """
        The bandwidth of the connection.
        """
        return pulumi.get(self, "bandwidth_in_mbps")

    @bandwidth_in_mbps.setter
    def bandwidth_in_mbps(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "bandwidth_in_mbps", value)

    @property
    @pulumi.getter(name="bgpSession")
    def bgp_session(self) -> Optional[pulumi.Input['BgpSessionArgs']]:
        """
        The BGP session associated with the connection.
        """
        return pulumi.get(self, "bgp_session")

    @bgp_session.setter
    def bgp_session(self, value: Optional[pulumi.Input['BgpSessionArgs']]):
        pulumi.set(self, "bgp_session", value)

    @property
    @pulumi.getter(name="connectionIdentifier")
    def connection_identifier(self) -> Optional[pulumi.Input[str]]:
        """
        The unique identifier (GUID) for the connection.
        """
        return pulumi.get(self, "connection_identifier")

    @connection_identifier.setter
    def connection_identifier(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "connection_identifier", value)

    @property
    @pulumi.getter(name="peeringDBFacilityId")
    def peering_db_facility_id(self) -> Optional[pulumi.Input[float]]:
        """
        The PeeringDB.com ID of the facility at which the connection has to be set up.
        """
        return pulumi.get(self, "peering_db_facility_id")

    @peering_db_facility_id.setter
    def peering_db_facility_id(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "peering_db_facility_id", value)

    @property
    @pulumi.getter(name="sessionAddressProvider")
    def session_address_provider(self) -> Optional[pulumi.Input[str]]:
        """
        The field indicating if Microsoft provides session ip addresses.
        """
        return pulumi.get(self, "session_address_provider")

    @session_address_provider.setter
    def session_address_provider(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "session_address_provider", value)

    @property
    @pulumi.getter(name="useForPeeringService")
    def use_for_peering_service(self) -> Optional[pulumi.Input[bool]]:
        """
        The flag that indicates whether or not the connection is used for peering service.
        """
        return pulumi.get(self, "use_for_peering_service")

    @use_for_peering_service.setter
    def use_for_peering_service(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "use_for_peering_service", value)


@pulumi.input_type
class ExchangeConnectionArgs:
    def __init__(__self__, *,
                 bgp_session: Optional[pulumi.Input['BgpSessionArgs']] = None,
                 connection_identifier: Optional[pulumi.Input[str]] = None,
                 peering_db_facility_id: Optional[pulumi.Input[float]] = None):
        """
        The properties that define an exchange connection.
        :param pulumi.Input['BgpSessionArgs'] bgp_session: The BGP session associated with the connection.
        :param pulumi.Input[str] connection_identifier: The unique identifier (GUID) for the connection.
        :param pulumi.Input[float] peering_db_facility_id: The PeeringDB.com ID of the facility at which the connection has to be set up.
        """
        if bgp_session is not None:
            pulumi.set(__self__, "bgp_session", bgp_session)
        if connection_identifier is not None:
            pulumi.set(__self__, "connection_identifier", connection_identifier)
        if peering_db_facility_id is not None:
            pulumi.set(__self__, "peering_db_facility_id", peering_db_facility_id)

    @property
    @pulumi.getter(name="bgpSession")
    def bgp_session(self) -> Optional[pulumi.Input['BgpSessionArgs']]:
        """
        The BGP session associated with the connection.
        """
        return pulumi.get(self, "bgp_session")

    @bgp_session.setter
    def bgp_session(self, value: Optional[pulumi.Input['BgpSessionArgs']]):
        pulumi.set(self, "bgp_session", value)

    @property
    @pulumi.getter(name="connectionIdentifier")
    def connection_identifier(self) -> Optional[pulumi.Input[str]]:
        """
        The unique identifier (GUID) for the connection.
        """
        return pulumi.get(self, "connection_identifier")

    @connection_identifier.setter
    def connection_identifier(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "connection_identifier", value)

    @property
    @pulumi.getter(name="peeringDBFacilityId")
    def peering_db_facility_id(self) -> Optional[pulumi.Input[float]]:
        """
        The PeeringDB.com ID of the facility at which the connection has to be set up.
        """
        return pulumi.get(self, "peering_db_facility_id")

    @peering_db_facility_id.setter
    def peering_db_facility_id(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "peering_db_facility_id", value)


@pulumi.input_type
class PeeringPropertiesDirectArgs:
    def __init__(__self__, *,
                 connections: Optional[pulumi.Input[List[pulumi.Input['DirectConnectionArgs']]]] = None,
                 direct_peering_type: Optional[pulumi.Input[str]] = None,
                 peer_asn: Optional[pulumi.Input['SubResourceArgs']] = None):
        """
        The properties that define a direct peering.
        :param pulumi.Input[List[pulumi.Input['DirectConnectionArgs']]] connections: The set of connections that constitute a direct peering.
        :param pulumi.Input[str] direct_peering_type: The type of direct peering.
        :param pulumi.Input['SubResourceArgs'] peer_asn: The reference of the peer ASN.
        """
        if connections is not None:
            pulumi.set(__self__, "connections", connections)
        if direct_peering_type is not None:
            pulumi.set(__self__, "direct_peering_type", direct_peering_type)
        if peer_asn is not None:
            pulumi.set(__self__, "peer_asn", peer_asn)

    @property
    @pulumi.getter
    def connections(self) -> Optional[pulumi.Input[List[pulumi.Input['DirectConnectionArgs']]]]:
        """
        The set of connections that constitute a direct peering.
        """
        return pulumi.get(self, "connections")

    @connections.setter
    def connections(self, value: Optional[pulumi.Input[List[pulumi.Input['DirectConnectionArgs']]]]):
        pulumi.set(self, "connections", value)

    @property
    @pulumi.getter(name="directPeeringType")
    def direct_peering_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of direct peering.
        """
        return pulumi.get(self, "direct_peering_type")

    @direct_peering_type.setter
    def direct_peering_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "direct_peering_type", value)

    @property
    @pulumi.getter(name="peerAsn")
    def peer_asn(self) -> Optional[pulumi.Input['SubResourceArgs']]:
        """
        The reference of the peer ASN.
        """
        return pulumi.get(self, "peer_asn")

    @peer_asn.setter
    def peer_asn(self, value: Optional[pulumi.Input['SubResourceArgs']]):
        pulumi.set(self, "peer_asn", value)


@pulumi.input_type
class PeeringPropertiesExchangeArgs:
    def __init__(__self__, *,
                 connections: Optional[pulumi.Input[List[pulumi.Input['ExchangeConnectionArgs']]]] = None,
                 peer_asn: Optional[pulumi.Input['SubResourceArgs']] = None):
        """
        The properties that define an exchange peering.
        :param pulumi.Input[List[pulumi.Input['ExchangeConnectionArgs']]] connections: The set of connections that constitute an exchange peering.
        :param pulumi.Input['SubResourceArgs'] peer_asn: The reference of the peer ASN.
        """
        if connections is not None:
            pulumi.set(__self__, "connections", connections)
        if peer_asn is not None:
            pulumi.set(__self__, "peer_asn", peer_asn)

    @property
    @pulumi.getter
    def connections(self) -> Optional[pulumi.Input[List[pulumi.Input['ExchangeConnectionArgs']]]]:
        """
        The set of connections that constitute an exchange peering.
        """
        return pulumi.get(self, "connections")

    @connections.setter
    def connections(self, value: Optional[pulumi.Input[List[pulumi.Input['ExchangeConnectionArgs']]]]):
        pulumi.set(self, "connections", value)

    @property
    @pulumi.getter(name="peerAsn")
    def peer_asn(self) -> Optional[pulumi.Input['SubResourceArgs']]:
        """
        The reference of the peer ASN.
        """
        return pulumi.get(self, "peer_asn")

    @peer_asn.setter
    def peer_asn(self, value: Optional[pulumi.Input['SubResourceArgs']]):
        pulumi.set(self, "peer_asn", value)


@pulumi.input_type
class PeeringSkuArgs:
    def __init__(__self__, *,
                 family: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 size: Optional[pulumi.Input[str]] = None,
                 tier: Optional[pulumi.Input[str]] = None):
        """
        The SKU that defines the tier and kind of the peering.
        :param pulumi.Input[str] family: The family of the peering SKU.
        :param pulumi.Input[str] name: The name of the peering SKU.
        :param pulumi.Input[str] size: The size of the peering SKU.
        :param pulumi.Input[str] tier: The tier of the peering SKU.
        """
        if family is not None:
            pulumi.set(__self__, "family", family)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if size is not None:
            pulumi.set(__self__, "size", size)
        if tier is not None:
            pulumi.set(__self__, "tier", tier)

    @property
    @pulumi.getter
    def family(self) -> Optional[pulumi.Input[str]]:
        """
        The family of the peering SKU.
        """
        return pulumi.get(self, "family")

    @family.setter
    def family(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "family", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the peering SKU.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def size(self) -> Optional[pulumi.Input[str]]:
        """
        The size of the peering SKU.
        """
        return pulumi.get(self, "size")

    @size.setter
    def size(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "size", value)

    @property
    @pulumi.getter
    def tier(self) -> Optional[pulumi.Input[str]]:
        """
        The tier of the peering SKU.
        """
        return pulumi.get(self, "tier")

    @tier.setter
    def tier(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tier", value)


@pulumi.input_type
class SubResourceArgs:
    def __init__(__self__, *,
                 id: Optional[pulumi.Input[str]] = None):
        """
        The sub resource.
        :param pulumi.Input[str] id: The identifier of the referenced resource.
        """
        if id is not None:
            pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        The identifier of the referenced resource.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)


