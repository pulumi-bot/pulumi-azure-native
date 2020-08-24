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
    'GetDedicatedHostResult',
    'AwaitableGetDedicatedHostResult',
    'get_dedicated_host',
]

@pulumi.output_type
class GetDedicatedHostResult:
    """
    Specifies information about the Dedicated host.
    """
    def __init__(__self__, auto_replace_on_failure=None, host_id=None, instance_view=None, license_type=None, location=None, name=None, platform_fault_domain=None, provisioning_state=None, provisioning_time=None, sku=None, tags=None, type=None, virtual_machines=None):
        if auto_replace_on_failure and not isinstance(auto_replace_on_failure, bool):
            raise TypeError("Expected argument 'auto_replace_on_failure' to be a bool")
        pulumi.set(__self__, "auto_replace_on_failure", auto_replace_on_failure)
        if host_id and not isinstance(host_id, str):
            raise TypeError("Expected argument 'host_id' to be a str")
        pulumi.set(__self__, "host_id", host_id)
        if instance_view and not isinstance(instance_view, dict):
            raise TypeError("Expected argument 'instance_view' to be a dict")
        pulumi.set(__self__, "instance_view", instance_view)
        if license_type and not isinstance(license_type, str):
            raise TypeError("Expected argument 'license_type' to be a str")
        pulumi.set(__self__, "license_type", license_type)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if platform_fault_domain and not isinstance(platform_fault_domain, float):
            raise TypeError("Expected argument 'platform_fault_domain' to be a float")
        pulumi.set(__self__, "platform_fault_domain", platform_fault_domain)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if provisioning_time and not isinstance(provisioning_time, str):
            raise TypeError("Expected argument 'provisioning_time' to be a str")
        pulumi.set(__self__, "provisioning_time", provisioning_time)
        if sku and not isinstance(sku, dict):
            raise TypeError("Expected argument 'sku' to be a dict")
        pulumi.set(__self__, "sku", sku)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if virtual_machines and not isinstance(virtual_machines, list):
            raise TypeError("Expected argument 'virtual_machines' to be a list")
        pulumi.set(__self__, "virtual_machines", virtual_machines)

    @property
    @pulumi.getter(name="autoReplaceOnFailure")
    def auto_replace_on_failure(self) -> Optional[bool]:
        """
        Specifies whether the dedicated host should be replaced automatically in case of a failure. The value is defaulted to 'true' when not provided.
        """
        return pulumi.get(self, "auto_replace_on_failure")

    @property
    @pulumi.getter(name="hostId")
    def host_id(self) -> str:
        """
        A unique id generated and assigned to the dedicated host by the platform. <br><br> Does not change throughout the lifetime of the host.
        """
        return pulumi.get(self, "host_id")

    @property
    @pulumi.getter(name="instanceView")
    def instance_view(self) -> 'outputs.DedicatedHostInstanceViewResponse':
        """
        The dedicated host instance view.
        """
        return pulumi.get(self, "instance_view")

    @property
    @pulumi.getter(name="licenseType")
    def license_type(self) -> Optional[str]:
        """
        Specifies the software license type that will be applied to the VMs deployed on the dedicated host. <br><br> Possible values are: <br><br> **None** <br><br> **Windows_Server_Hybrid** <br><br> **Windows_Server_Perpetual** <br><br> Default: **None**
        """
        return pulumi.get(self, "license_type")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        Resource location
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="platformFaultDomain")
    def platform_fault_domain(self) -> Optional[float]:
        """
        Fault domain of the dedicated host within a dedicated host group.
        """
        return pulumi.get(self, "platform_fault_domain")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> str:
        """
        The provisioning state, which only appears in the response.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="provisioningTime")
    def provisioning_time(self) -> str:
        """
        The date when the host was first provisioned.
        """
        return pulumi.get(self, "provisioning_time")

    @property
    @pulumi.getter
    def sku(self) -> 'outputs.SkuResponse':
        """
        SKU of the dedicated host for Hardware Generation and VM family. Only name is required to be set. List Microsoft.Compute SKUs for a list of possible values.
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        Resource tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Resource type
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="virtualMachines")
    def virtual_machines(self) -> List['outputs.SubResourceReadOnlyResponse']:
        """
        A list of references to all virtual machines in the Dedicated Host.
        """
        return pulumi.get(self, "virtual_machines")


class AwaitableGetDedicatedHostResult(GetDedicatedHostResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDedicatedHostResult(
            auto_replace_on_failure=self.auto_replace_on_failure,
            host_id=self.host_id,
            instance_view=self.instance_view,
            license_type=self.license_type,
            location=self.location,
            name=self.name,
            platform_fault_domain=self.platform_fault_domain,
            provisioning_state=self.provisioning_state,
            provisioning_time=self.provisioning_time,
            sku=self.sku,
            tags=self.tags,
            type=self.type,
            virtual_machines=self.virtual_machines)


def get_dedicated_host(expand: Optional[str] = None,
                       host_group_name: Optional[str] = None,
                       name: Optional[str] = None,
                       resource_group_name: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDedicatedHostResult:
    """
    Use this data source to access information about an existing resource.

    :param str expand: The expand expression to apply on the operation.
    :param str host_group_name: The name of the dedicated host group.
    :param str name: The name of the dedicated host.
    :param str resource_group_name: The name of the resource group.
    """
    __args__ = dict()
    __args__['expand'] = expand
    __args__['hostGroupName'] = host_group_name
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azurerm:compute/v20190701:getDedicatedHost', __args__, opts=opts, typ=GetDedicatedHostResult).value

    return AwaitableGetDedicatedHostResult(
        auto_replace_on_failure=__ret__.auto_replace_on_failure,
        host_id=__ret__.host_id,
        instance_view=__ret__.instance_view,
        license_type=__ret__.license_type,
        location=__ret__.location,
        name=__ret__.name,
        platform_fault_domain=__ret__.platform_fault_domain,
        provisioning_state=__ret__.provisioning_state,
        provisioning_time=__ret__.provisioning_time,
        sku=__ret__.sku,
        tags=__ret__.tags,
        type=__ret__.type,
        virtual_machines=__ret__.virtual_machines)
