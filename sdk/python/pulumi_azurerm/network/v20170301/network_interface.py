# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class NetworkInterface(pulumi.CustomResource):
    etag: pulumi.Output[str]
    """
    A unique read-only string that changes whenever the resource is updated.
    """
    location: pulumi.Output[str]
    """
    Resource location.
    """
    name: pulumi.Output[str]
    """
    Resource name.
    """
    properties: pulumi.Output[dict]
    """
    NetworkInterface properties. 
      * `dns_settings` (`dict`) - The DNS settings in network interface.
        * `applied_dns_servers` (`list`) - If the VM that uses this NIC is part of an Availability Set, then this list will have the union of all DNS servers from all NICs that are part of the Availability Set. This property is what is configured on each of those VMs.
        * `dns_servers` (`list`) - List of DNS servers IP addresses. Use 'AzureProvidedDNS' to switch to azure provided DNS resolution. 'AzureProvidedDNS' value cannot be combined with other IPs, it must be the only value in dnsServers collection.
        * `internal_dns_name_label` (`str`) - Relative DNS name for this NIC used for internal communications between VMs in the same virtual network.
        * `internal_domain_name_suffix` (`str`) - Even if internalDnsNameLabel is not specified, a DNS entry is created for the primary NIC of the VM. This DNS name can be constructed by concatenating the VM name with the value of internalDomainNameSuffix.
        * `internal_fqdn` (`str`) - Fully qualified DNS name supporting internal communications between VMs in the same virtual network.

      * `enable_accelerated_networking` (`bool`) - If the network interface is accelerated networking enabled.
      * `enable_ip_forwarding` (`bool`) - Indicates whether IP forwarding is enabled on this network interface.
      * `ip_configurations` (`list`) - A list of IPConfigurations of the network interface.
        * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
        * `id` (`str`) - Resource ID.
        * `name` (`str`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
        * `properties` (`dict`) - Properties of IP configuration.
          * `application_gateway_backend_address_pools` (`list`) - The reference of ApplicationGatewayBackendAddressPool resource.
            * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
            * `id` (`str`) - Resource ID.
            * `name` (`str`) - Resource that is unique within a resource group. This name can be used to access the resource.
            * `properties` (`dict`) - Properties of Backend Address Pool of an application gateway.
              * `backend_addresses` (`list`) - Backend addresses
                * `fqdn` (`str`) - Fully qualified domain name (FQDN).
                * `ip_address` (`str`) - IP address

              * `backend_ip_configurations` (`list`) - Collection of references to IPs defined in network interfaces.
              * `provisioning_state` (`str`) - Provisioning state of the backend address pool resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.

          * `load_balancer_backend_address_pools` (`list`) - The reference of LoadBalancerBackendAddressPool resource.
            * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
            * `id` (`str`) - Resource ID.
            * `name` (`str`) - Gets name of the resource that is unique within a resource group. This name can be used to access the resource.
            * `properties` (`dict`) - Properties of the backend address pool.
              * `backend_ip_configurations` (`list`) - Gets collection of references to IP addresses defined in network interfaces.
              * `load_balancing_rules` (`list`) - Gets load balancing rules that use this backend address pool.
                * `id` (`str`) - Resource ID.

              * `outbound_nat_rule` (`dict`) - Gets outbound rules that use this backend address pool.
              * `provisioning_state` (`str`) - Get provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.

          * `load_balancer_inbound_nat_rules` (`list`) - A list of references of LoadBalancerInboundNatRules.
            * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
            * `id` (`str`) - Resource ID.
            * `name` (`str`) - Gets name of the resource that is unique within a resource group. This name can be used to access the resource.
            * `properties` (`dict`) - Properties of the inbound NAT rule.
              * `backend_ip_configuration` (`dict`) - A reference to a private IP address defined on a network interface of a VM. Traffic sent to the frontend port of each of the frontend IP configurations is forwarded to the backed IP.
              * `backend_port` (`float`) - The port used for the internal endpoint. Acceptable values range from 1 to 65535.
              * `enable_floating_ip` (`bool`) - Configures a virtual machine's endpoint for the floating IP capability required to configure a SQL AlwaysOn Availability Group. This setting is required when using the SQL AlwaysOn Availability Groups in SQL server. This setting can't be changed after you create the endpoint.
              * `frontend_ip_configuration` (`dict`) - A reference to frontend IP addresses.
              * `frontend_port` (`float`) - The port for the external endpoint. Port numbers for each rule must be unique within the Load Balancer. Acceptable values range from 1 to 65534.
              * `idle_timeout_in_minutes` (`float`) - The timeout for the TCP idle connection. The value can be set between 4 and 30 minutes. The default value is 4 minutes. This element is only used when the protocol is set to TCP.
              * `protocol` (`str`) - The transport protocol for the endpoint. Possible values are: 'Udp' or 'Tcp'
              * `provisioning_state` (`str`) - Gets the provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.

          * `primary` (`bool`) - Gets whether this is a primary customer address on the network interface.
          * `private_ip_address` (`str`)
          * `private_ip_address_version` (`str`) - Available from Api-Version 2016-03-30 onwards, it represents whether the specific ipconfiguration is IPv4 or IPv6. Default is taken as IPv4.  Possible values are: 'IPv4' and 'IPv6'.
          * `private_ip_allocation_method` (`str`) - Defines how a private IP address is assigned. Possible values are: 'Static' and 'Dynamic'.
          * `provisioning_state` (`str`)
          * `public_ip_address` (`dict`) - Public IP address resource.
            * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
            * `id` (`str`) - Resource ID.
            * `location` (`str`) - Resource location.
            * `name` (`str`) - Resource name.
            * `properties` (`dict`) - Public IP address properties.
              * `dns_settings` (`dict`) - The FQDN of the DNS record associated with the public IP address.
                * `domain_name_label` (`str`) - Gets or sets the Domain name label.The concatenation of the domain name label and the regionalized DNS zone make up the fully qualified domain name associated with the public IP address. If a domain name label is specified, an A DNS record is created for the public IP in the Microsoft Azure DNS system.
                * `fqdn` (`str`) - Gets the FQDN, Fully qualified domain name of the A DNS record associated with the public IP. This is the concatenation of the domainNameLabel and the regionalized DNS zone.
                * `reverse_fqdn` (`str`) - Gets or Sets the Reverse FQDN. A user-visible, fully qualified domain name that resolves to this public IP address. If the reverseFqdn is specified, then a PTR DNS record is created pointing from the IP address in the in-addr.arpa domain to the reverse FQDN. 

              * `idle_timeout_in_minutes` (`float`) - The idle timeout of the public IP address.
              * `ip_address` (`str`)
              * `ip_configuration` (`dict`) - IPConfiguration
                * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
                * `id` (`str`) - Resource ID.
                * `name` (`str`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
                * `properties` (`dict`) - Properties of IP configuration.
                  * `private_ip_address` (`str`) - The private IP address of the IP configuration.
                  * `private_ip_allocation_method` (`str`) - The private IP allocation method. Possible values are 'Static' and 'Dynamic'.
                  * `provisioning_state` (`str`) - Gets the provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
                  * `public_ip_address` (`dict`) - The reference of the public IP resource.
                  * `subnet` (`dict`) - The reference of the subnet resource.
                    * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
                    * `id` (`str`) - Resource ID.
                    * `name` (`str`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
                    * `properties` (`dict`)
                      * `address_prefix` (`str`) - The address prefix for the subnet.
                      * `ip_configurations` (`list`) - Gets an array of references to the network interface IP configurations using subnet.
                      * `network_security_group` (`dict`) - The reference of the NetworkSecurityGroup resource.
                        * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
                        * `id` (`str`) - Resource ID.
                        * `location` (`str`) - Resource location.
                        * `name` (`str`) - Resource name.
                        * `properties` (`dict`) - Network Security Group resource.
                          * `default_security_rules` (`list`) - The default security rules of network security group.
                            * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
                            * `id` (`str`) - Resource ID.
                            * `name` (`str`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
                            * `properties` (`dict`)
                              * `access` (`str`) - The network traffic is allowed or denied. Possible values are: 'Allow' and 'Deny'.
                              * `description` (`str`) - A description for this rule. Restricted to 140 chars.
                              * `destination_address_prefix` (`str`) - The destination address prefix. CIDR or source IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used.
                              * `destination_port_range` (`str`) - The destination port or range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.
                              * `direction` (`str`) - The direction of the rule. The direction specifies if rule will be evaluated on incoming or outgoing traffic. Possible values are: 'Inbound' and 'Outbound'.
                              * `priority` (`float`) - The priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
                              * `protocol` (`str`) - Network protocol this rule applies to. Possible values are 'Tcp', 'Udp', and '*'.
                              * `provisioning_state` (`str`) - The provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
                              * `source_address_prefix` (`str`) - The CIDR or source IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used. If this is an ingress rule, specifies where network traffic originates from. 
                              * `source_port_range` (`str`) - The source port or range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.

                          * `network_interfaces` (`list`) - A collection of references to network interfaces.
                            * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
                            * `id` (`str`) - Resource ID.
                            * `location` (`str`) - Resource location.
                            * `name` (`str`) - Resource name.
                            * `properties` (`dict`) - NetworkInterface properties. 
                            * `tags` (`dict`) - Resource tags.
                            * `type` (`str`) - Resource type.

                          * `provisioning_state` (`str`) - The provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
                          * `resource_guid` (`str`) - The resource GUID property of the network security group resource.
                          * `security_rules` (`list`) - A collection of security rules of the network security group.
                          * `subnets` (`list`) - A collection of references to subnets.

                        * `tags` (`dict`) - Resource tags.
                        * `type` (`str`) - Resource type.

                      * `provisioning_state` (`str`) - The provisioning state of the resource.
                      * `resource_navigation_links` (`list`) - Gets an array of references to the external resources using subnet.
                        * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
                        * `id` (`str`) - Resource ID.
                        * `name` (`str`) - Name of the resource that is unique within a resource group. This name can be used to access the resource.
                        * `properties` (`dict`) - Properties of ResourceNavigationLink.
                          * `link` (`str`) - Link to the external resource
                          * `linked_resource_type` (`str`) - Resource type of the linked resource.
                          * `provisioning_state` (`str`) - Provisioning state of the ResourceNavigationLink resource.

                      * `route_table` (`dict`) - The reference of the RouteTable resource.
                        * `etag` (`str`) - Gets a unique read-only string that changes whenever the resource is updated.
                        * `id` (`str`) - Resource ID.
                        * `location` (`str`) - Resource location.
                        * `name` (`str`) - Resource name.
                        * `properties` (`dict`) - Route Table resource
                          * `provisioning_state` (`str`) - The provisioning state of the resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
                          * `routes` (`list`) - Collection of routes contained within a route table.
                            * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
                            * `id` (`str`) - Resource ID.
                            * `name` (`str`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
                            * `properties` (`dict`) - Route resource
                              * `address_prefix` (`str`) - The destination CIDR to which the route applies.
                              * `next_hop_ip_address` (`str`) - The IP address packets should be forwarded to. Next hop values are only allowed in routes where the next hop type is VirtualAppliance.
                              * `next_hop_type` (`str`) - The type of Azure hop the packet should be sent to. Possible values are: 'VirtualNetworkGateway', 'VnetLocal', 'Internet', 'VirtualAppliance', and 'None'
                              * `provisioning_state` (`str`) - The provisioning state of the resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.

                          * `subnets` (`list`) - A collection of references to subnets.

                        * `tags` (`dict`) - Resource tags.
                        * `type` (`str`) - Resource type.

              * `provisioning_state` (`str`) - The provisioning state of the PublicIP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
              * `public_ip_address_version` (`str`) - The public IP address version. Possible values are: 'IPv4' and 'IPv6'.
              * `public_ip_allocation_method` (`str`) - The public IP allocation method. Possible values are: 'Static' and 'Dynamic'.
              * `resource_guid` (`str`) - The resource GUID property of the public IP resource.

            * `tags` (`dict`) - Resource tags.
            * `type` (`str`) - Resource type.

          * `subnet` (`dict`) - Subnet in a virtual network resource.

      * `mac_address` (`str`) - The MAC address of the network interface.
      * `network_security_group` (`dict`) - The reference of the NetworkSecurityGroup resource.
      * `primary` (`bool`) - Gets whether this is a primary network interface on a virtual machine.
      * `provisioning_state` (`str`) - The provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
      * `resource_guid` (`str`) - The resource GUID property of the network interface resource.
      * `virtual_machine` (`dict`) - The reference of a virtual machine.
    """
    tags: pulumi.Output[dict]
    """
    Resource tags.
    """
    type: pulumi.Output[str]
    """
    Resource type.
    """
    def __init__(__self__, resource_name, opts=None, dns_settings=None, enable_accelerated_networking=None, enable_ip_forwarding=None, etag=None, id=None, ip_configurations=None, location=None, mac_address=None, name=None, network_security_group=None, primary=None, provisioning_state=None, resource_group_name=None, resource_guid=None, tags=None, virtual_machine=None, __props__=None, __name__=None, __opts__=None):
        """
        A network interface in a resource group.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] dns_settings: The DNS settings in network interface.
        :param pulumi.Input[bool] enable_accelerated_networking: If the network interface is accelerated networking enabled.
        :param pulumi.Input[bool] enable_ip_forwarding: Indicates whether IP forwarding is enabled on this network interface.
        :param pulumi.Input[str] etag: A unique read-only string that changes whenever the resource is updated.
        :param pulumi.Input[str] id: Resource ID.
        :param pulumi.Input[list] ip_configurations: A list of IPConfigurations of the network interface.
        :param pulumi.Input[str] location: Resource location.
        :param pulumi.Input[str] mac_address: The MAC address of the network interface.
        :param pulumi.Input[str] name: The name of the network interface.
        :param pulumi.Input[dict] network_security_group: The reference of the NetworkSecurityGroup resource.
        :param pulumi.Input[bool] primary: Gets whether this is a primary network interface on a virtual machine.
        :param pulumi.Input[str] provisioning_state: The provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[str] resource_guid: The resource GUID property of the network interface resource.
        :param pulumi.Input[dict] tags: Resource tags.
        :param pulumi.Input[dict] virtual_machine: The reference of a virtual machine.

        The **dns_settings** object supports the following:

          * `applied_dns_servers` (`pulumi.Input[list]`) - If the VM that uses this NIC is part of an Availability Set, then this list will have the union of all DNS servers from all NICs that are part of the Availability Set. This property is what is configured on each of those VMs.
          * `dns_servers` (`pulumi.Input[list]`) - List of DNS servers IP addresses. Use 'AzureProvidedDNS' to switch to azure provided DNS resolution. 'AzureProvidedDNS' value cannot be combined with other IPs, it must be the only value in dnsServers collection.
          * `internal_dns_name_label` (`pulumi.Input[str]`) - Relative DNS name for this NIC used for internal communications between VMs in the same virtual network.
          * `internal_domain_name_suffix` (`pulumi.Input[str]`) - Even if internalDnsNameLabel is not specified, a DNS entry is created for the primary NIC of the VM. This DNS name can be constructed by concatenating the VM name with the value of internalDomainNameSuffix.
          * `internal_fqdn` (`pulumi.Input[str]`) - Fully qualified DNS name supporting internal communications between VMs in the same virtual network.

        The **ip_configurations** object supports the following:

          * `application_gateway_backend_address_pools` (`pulumi.Input[list]`) - The reference of ApplicationGatewayBackendAddressPool resource.
            * `backend_addresses` (`pulumi.Input[list]`) - Backend addresses
              * `fqdn` (`pulumi.Input[str]`) - Fully qualified domain name (FQDN).
              * `ip_address` (`pulumi.Input[str]`) - IP address

            * `backend_ip_configurations` (`pulumi.Input[list]`) - Collection of references to IPs defined in network interfaces.
            * `etag` (`pulumi.Input[str]`) - A unique read-only string that changes whenever the resource is updated.
            * `id` (`pulumi.Input[str]`) - Resource ID.
            * `name` (`pulumi.Input[str]`) - Resource that is unique within a resource group. This name can be used to access the resource.
            * `provisioning_state` (`pulumi.Input[str]`) - Provisioning state of the backend address pool resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.

          * `etag` (`pulumi.Input[str]`) - A unique read-only string that changes whenever the resource is updated.
          * `id` (`pulumi.Input[str]`) - Resource ID.
          * `load_balancer_backend_address_pools` (`pulumi.Input[list]`) - The reference of LoadBalancerBackendAddressPool resource.
            * `etag` (`pulumi.Input[str]`) - A unique read-only string that changes whenever the resource is updated.
            * `id` (`pulumi.Input[str]`) - Resource ID.
            * `name` (`pulumi.Input[str]`) - Gets name of the resource that is unique within a resource group. This name can be used to access the resource.
            * `provisioning_state` (`pulumi.Input[str]`) - Get provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.

          * `load_balancer_inbound_nat_rules` (`pulumi.Input[list]`) - A list of references of LoadBalancerInboundNatRules.
            * `backend_port` (`pulumi.Input[float]`) - The port used for the internal endpoint. Acceptable values range from 1 to 65535.
            * `enable_floating_ip` (`pulumi.Input[bool]`) - Configures a virtual machine's endpoint for the floating IP capability required to configure a SQL AlwaysOn Availability Group. This setting is required when using the SQL AlwaysOn Availability Groups in SQL server. This setting can't be changed after you create the endpoint.
            * `etag` (`pulumi.Input[str]`) - A unique read-only string that changes whenever the resource is updated.
            * `frontend_ip_configuration` (`pulumi.Input[dict]`) - A reference to frontend IP addresses.
              * `id` (`pulumi.Input[str]`) - Resource ID.

            * `frontend_port` (`pulumi.Input[float]`) - The port for the external endpoint. Port numbers for each rule must be unique within the Load Balancer. Acceptable values range from 1 to 65534.
            * `id` (`pulumi.Input[str]`) - Resource ID.
            * `idle_timeout_in_minutes` (`pulumi.Input[float]`) - The timeout for the TCP idle connection. The value can be set between 4 and 30 minutes. The default value is 4 minutes. This element is only used when the protocol is set to TCP.
            * `name` (`pulumi.Input[str]`) - Gets name of the resource that is unique within a resource group. This name can be used to access the resource.
            * `protocol` (`pulumi.Input[str]`) - The transport protocol for the endpoint. Possible values are: 'Udp' or 'Tcp'
            * `provisioning_state` (`pulumi.Input[str]`) - Gets the provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.

          * `name` (`pulumi.Input[str]`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
          * `primary` (`pulumi.Input[bool]`) - Gets whether this is a primary customer address on the network interface.
          * `private_ip_address` (`pulumi.Input[str]`)
          * `private_ip_address_version` (`pulumi.Input[str]`) - Available from Api-Version 2016-03-30 onwards, it represents whether the specific ipconfiguration is IPv4 or IPv6. Default is taken as IPv4.  Possible values are: 'IPv4' and 'IPv6'.
          * `private_ip_allocation_method` (`pulumi.Input[str]`) - Defines how a private IP address is assigned. Possible values are: 'Static' and 'Dynamic'.
          * `provisioning_state` (`pulumi.Input[str]`)
          * `public_ip_address` (`pulumi.Input[dict]`) - Public IP address resource.
            * `dns_settings` (`pulumi.Input[dict]`) - The FQDN of the DNS record associated with the public IP address.
              * `domain_name_label` (`pulumi.Input[str]`) - Gets or sets the Domain name label.The concatenation of the domain name label and the regionalized DNS zone make up the fully qualified domain name associated with the public IP address. If a domain name label is specified, an A DNS record is created for the public IP in the Microsoft Azure DNS system.
              * `fqdn` (`pulumi.Input[str]`) - Gets the FQDN, Fully qualified domain name of the A DNS record associated with the public IP. This is the concatenation of the domainNameLabel and the regionalized DNS zone.
              * `reverse_fqdn` (`pulumi.Input[str]`) - Gets or Sets the Reverse FQDN. A user-visible, fully qualified domain name that resolves to this public IP address. If the reverseFqdn is specified, then a PTR DNS record is created pointing from the IP address in the in-addr.arpa domain to the reverse FQDN. 

            * `etag` (`pulumi.Input[str]`) - A unique read-only string that changes whenever the resource is updated.
            * `id` (`pulumi.Input[str]`) - Resource ID.
            * `idle_timeout_in_minutes` (`pulumi.Input[float]`) - The idle timeout of the public IP address.
            * `ip_address` (`pulumi.Input[str]`)
            * `location` (`pulumi.Input[str]`) - Resource location.
            * `provisioning_state` (`pulumi.Input[str]`) - The provisioning state of the PublicIP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
            * `public_ip_address_version` (`pulumi.Input[str]`) - The public IP address version. Possible values are: 'IPv4' and 'IPv6'.
            * `public_ip_allocation_method` (`pulumi.Input[str]`) - The public IP allocation method. Possible values are: 'Static' and 'Dynamic'.
            * `resource_guid` (`pulumi.Input[str]`) - The resource GUID property of the public IP resource.
            * `tags` (`pulumi.Input[dict]`) - Resource tags.

          * `subnet` (`pulumi.Input[dict]`) - Subnet in a virtual network resource.
            * `address_prefix` (`pulumi.Input[str]`) - The address prefix for the subnet.
            * `etag` (`pulumi.Input[str]`) - A unique read-only string that changes whenever the resource is updated.
            * `id` (`pulumi.Input[str]`) - Resource ID.
            * `name` (`pulumi.Input[str]`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
            * `network_security_group` (`pulumi.Input[dict]`) - The reference of the NetworkSecurityGroup resource.
              * `default_security_rules` (`pulumi.Input[list]`) - The default security rules of network security group.
                * `access` (`pulumi.Input[str]`) - The network traffic is allowed or denied. Possible values are: 'Allow' and 'Deny'.
                * `description` (`pulumi.Input[str]`) - A description for this rule. Restricted to 140 chars.
                * `destination_address_prefix` (`pulumi.Input[str]`) - The destination address prefix. CIDR or source IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used.
                * `destination_port_range` (`pulumi.Input[str]`) - The destination port or range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.
                * `direction` (`pulumi.Input[str]`) - The direction of the rule. The direction specifies if rule will be evaluated on incoming or outgoing traffic. Possible values are: 'Inbound' and 'Outbound'.
                * `etag` (`pulumi.Input[str]`) - A unique read-only string that changes whenever the resource is updated.
                * `id` (`pulumi.Input[str]`) - Resource ID.
                * `name` (`pulumi.Input[str]`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
                * `priority` (`pulumi.Input[float]`) - The priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
                * `protocol` (`pulumi.Input[str]`) - Network protocol this rule applies to. Possible values are 'Tcp', 'Udp', and '*'.
                * `provisioning_state` (`pulumi.Input[str]`) - The provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
                * `source_address_prefix` (`pulumi.Input[str]`) - The CIDR or source IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used. If this is an ingress rule, specifies where network traffic originates from. 
                * `source_port_range` (`pulumi.Input[str]`) - The source port or range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.

              * `etag` (`pulumi.Input[str]`) - A unique read-only string that changes whenever the resource is updated.
              * `id` (`pulumi.Input[str]`) - Resource ID.
              * `location` (`pulumi.Input[str]`) - Resource location.
              * `provisioning_state` (`pulumi.Input[str]`) - The provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
              * `resource_guid` (`pulumi.Input[str]`) - The resource GUID property of the network security group resource.
              * `security_rules` (`pulumi.Input[list]`) - A collection of security rules of the network security group.
              * `tags` (`pulumi.Input[dict]`) - Resource tags.

            * `provisioning_state` (`pulumi.Input[str]`) - The provisioning state of the resource.
            * `resource_navigation_links` (`pulumi.Input[list]`) - Gets an array of references to the external resources using subnet.
              * `id` (`pulumi.Input[str]`) - Resource ID.
              * `link` (`pulumi.Input[str]`) - Link to the external resource
              * `linked_resource_type` (`pulumi.Input[str]`) - Resource type of the linked resource.
              * `name` (`pulumi.Input[str]`) - Name of the resource that is unique within a resource group. This name can be used to access the resource.

            * `route_table` (`pulumi.Input[dict]`) - The reference of the RouteTable resource.
              * `etag` (`pulumi.Input[str]`) - Gets a unique read-only string that changes whenever the resource is updated.
              * `id` (`pulumi.Input[str]`) - Resource ID.
              * `location` (`pulumi.Input[str]`) - Resource location.
              * `provisioning_state` (`pulumi.Input[str]`) - The provisioning state of the resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
              * `routes` (`pulumi.Input[list]`) - Collection of routes contained within a route table.
                * `address_prefix` (`pulumi.Input[str]`) - The destination CIDR to which the route applies.
                * `etag` (`pulumi.Input[str]`) - A unique read-only string that changes whenever the resource is updated.
                * `id` (`pulumi.Input[str]`) - Resource ID.
                * `name` (`pulumi.Input[str]`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
                * `next_hop_ip_address` (`pulumi.Input[str]`) - The IP address packets should be forwarded to. Next hop values are only allowed in routes where the next hop type is VirtualAppliance.
                * `next_hop_type` (`pulumi.Input[str]`) - The type of Azure hop the packet should be sent to. Possible values are: 'VirtualNetworkGateway', 'VnetLocal', 'Internet', 'VirtualAppliance', and 'None'
                * `provisioning_state` (`pulumi.Input[str]`) - The provisioning state of the resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.

              * `tags` (`pulumi.Input[dict]`) - Resource tags.
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

            __props__['dns_settings'] = dns_settings
            __props__['enable_accelerated_networking'] = enable_accelerated_networking
            __props__['enable_ip_forwarding'] = enable_ip_forwarding
            __props__['etag'] = etag
            __props__['id'] = id
            __props__['ip_configurations'] = ip_configurations
            __props__['location'] = location
            __props__['mac_address'] = mac_address
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['network_security_group'] = network_security_group
            __props__['primary'] = primary
            __props__['provisioning_state'] = provisioning_state
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['resource_guid'] = resource_guid
            __props__['tags'] = tags
            __props__['virtual_machine'] = virtual_machine
            __props__['properties'] = None
            __props__['type'] = None
        super(NetworkInterface, __self__).__init__(
            'azurerm:network/v20170301:NetworkInterface',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing NetworkInterface resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return NetworkInterface(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
