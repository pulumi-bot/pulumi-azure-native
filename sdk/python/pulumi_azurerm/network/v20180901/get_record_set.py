# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from ... import _utilities, _tables
from . import outputs

__all__ = [
    'GetRecordSetResult',
    'AwaitableGetRecordSetResult',
    'get_record_set',
]

@pulumi.output_type
class GetRecordSetResult:
    """
    Describes a DNS record set (a collection of DNS records with the same name and type) in a Private DNS zone.
    """
    def __init__(__self__, a_records=None, aaaa_records=None, cname_record=None, etag=None, fqdn=None, is_auto_registered=None, metadata=None, mx_records=None, name=None, ptr_records=None, soa_record=None, srv_records=None, ttl=None, txt_records=None, type=None):
        if a_records and not isinstance(a_records, list):
            raise TypeError("Expected argument 'a_records' to be a list")
        pulumi.set(__self__, "a_records", a_records)
        if aaaa_records and not isinstance(aaaa_records, list):
            raise TypeError("Expected argument 'aaaa_records' to be a list")
        pulumi.set(__self__, "aaaa_records", aaaa_records)
        if cname_record and not isinstance(cname_record, dict):
            raise TypeError("Expected argument 'cname_record' to be a dict")
        pulumi.set(__self__, "cname_record", cname_record)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if fqdn and not isinstance(fqdn, str):
            raise TypeError("Expected argument 'fqdn' to be a str")
        pulumi.set(__self__, "fqdn", fqdn)
        if is_auto_registered and not isinstance(is_auto_registered, bool):
            raise TypeError("Expected argument 'is_auto_registered' to be a bool")
        pulumi.set(__self__, "is_auto_registered", is_auto_registered)
        if metadata and not isinstance(metadata, dict):
            raise TypeError("Expected argument 'metadata' to be a dict")
        pulumi.set(__self__, "metadata", metadata)
        if mx_records and not isinstance(mx_records, list):
            raise TypeError("Expected argument 'mx_records' to be a list")
        pulumi.set(__self__, "mx_records", mx_records)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if ptr_records and not isinstance(ptr_records, list):
            raise TypeError("Expected argument 'ptr_records' to be a list")
        pulumi.set(__self__, "ptr_records", ptr_records)
        if soa_record and not isinstance(soa_record, dict):
            raise TypeError("Expected argument 'soa_record' to be a dict")
        pulumi.set(__self__, "soa_record", soa_record)
        if srv_records and not isinstance(srv_records, list):
            raise TypeError("Expected argument 'srv_records' to be a list")
        pulumi.set(__self__, "srv_records", srv_records)
        if ttl and not isinstance(ttl, float):
            raise TypeError("Expected argument 'ttl' to be a float")
        pulumi.set(__self__, "ttl", ttl)
        if txt_records and not isinstance(txt_records, list):
            raise TypeError("Expected argument 'txt_records' to be a list")
        pulumi.set(__self__, "txt_records", txt_records)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="aRecords")
    def a_records(self) -> Optional[List['outputs.ARecordResponse']]:
        """
        The list of A records in the record set.
        """
        return pulumi.get(self, "a_records")

    @property
    @pulumi.getter(name="aaaaRecords")
    def aaaa_records(self) -> Optional[List['outputs.AaaaRecordResponse']]:
        """
        The list of AAAA records in the record set.
        """
        return pulumi.get(self, "aaaa_records")

    @property
    @pulumi.getter(name="cnameRecord")
    def cname_record(self) -> Optional['outputs.CnameRecordResponse']:
        """
        The CNAME record in the record set.
        """
        return pulumi.get(self, "cname_record")

    @property
    @pulumi.getter
    def etag(self) -> Optional[str]:
        """
        The ETag of the record set.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def fqdn(self) -> str:
        """
        Fully qualified domain name of the record set.
        """
        return pulumi.get(self, "fqdn")

    @property
    @pulumi.getter(name="isAutoRegistered")
    def is_auto_registered(self) -> bool:
        """
        Is the record set auto-registered in the Private DNS zone through a virtual network link?
        """
        return pulumi.get(self, "is_auto_registered")

    @property
    @pulumi.getter
    def metadata(self) -> Optional[Mapping[str, str]]:
        """
        The metadata attached to the record set.
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter(name="mxRecords")
    def mx_records(self) -> Optional[List['outputs.MxRecordResponse']]:
        """
        The list of MX records in the record set.
        """
        return pulumi.get(self, "mx_records")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="ptrRecords")
    def ptr_records(self) -> Optional[List['outputs.PtrRecordResponse']]:
        """
        The list of PTR records in the record set.
        """
        return pulumi.get(self, "ptr_records")

    @property
    @pulumi.getter(name="soaRecord")
    def soa_record(self) -> Optional['outputs.SoaRecordResponse']:
        """
        The SOA record in the record set.
        """
        return pulumi.get(self, "soa_record")

    @property
    @pulumi.getter(name="srvRecords")
    def srv_records(self) -> Optional[List['outputs.SrvRecordResponse']]:
        """
        The list of SRV records in the record set.
        """
        return pulumi.get(self, "srv_records")

    @property
    @pulumi.getter
    def ttl(self) -> Optional[float]:
        """
        The TTL (time-to-live) of the records in the record set.
        """
        return pulumi.get(self, "ttl")

    @property
    @pulumi.getter(name="txtRecords")
    def txt_records(self) -> Optional[List['outputs.TxtRecordResponse']]:
        """
        The list of TXT records in the record set.
        """
        return pulumi.get(self, "txt_records")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the resource. Example - 'Microsoft.Network/privateDnsZones'.
        """
        return pulumi.get(self, "type")


class AwaitableGetRecordSetResult(GetRecordSetResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRecordSetResult(
            a_records=self.a_records,
            aaaa_records=self.aaaa_records,
            cname_record=self.cname_record,
            etag=self.etag,
            fqdn=self.fqdn,
            is_auto_registered=self.is_auto_registered,
            metadata=self.metadata,
            mx_records=self.mx_records,
            name=self.name,
            ptr_records=self.ptr_records,
            soa_record=self.soa_record,
            srv_records=self.srv_records,
            ttl=self.ttl,
            txt_records=self.txt_records,
            type=self.type)


def get_record_set(name: Optional[str] = None,
                   private_zone_name: Optional[str] = None,
                   record_type: Optional[str] = None,
                   resource_group_name: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRecordSetResult:
    """
    Use this data source to access information about an existing resource.

    :param str name: The name of the record set, relative to the name of the zone.
    :param str private_zone_name: The name of the Private DNS zone (without a terminating dot).
    :param str record_type: The type of DNS record in this record set.
    :param str resource_group_name: The name of the resource group.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['privateZoneName'] = private_zone_name
    __args__['recordType'] = record_type
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:network/v20180901:getRecordSet', __args__, opts=opts, typ=GetRecordSetResult).value

    return AwaitableGetRecordSetResult(
        a_records=__ret__.a_records,
        aaaa_records=__ret__.aaaa_records,
        cname_record=__ret__.cname_record,
        etag=__ret__.etag,
        fqdn=__ret__.fqdn,
        is_auto_registered=__ret__.is_auto_registered,
        metadata=__ret__.metadata,
        mx_records=__ret__.mx_records,
        name=__ret__.name,
        ptr_records=__ret__.ptr_records,
        soa_record=__ret__.soa_record,
        srv_records=__ret__.srv_records,
        ttl=__ret__.ttl,
        txt_records=__ret__.txt_records,
        type=__ret__.type)
