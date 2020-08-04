# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class VirtualNetwork(pulumi.CustomResource):
    etag: pulumi.Output[str]
    """
    Gets a unique read-only string that changes whenever the resource is updated.
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
    Properties of the virtual network.
      * `address_space` (`dict`) - The AddressSpace that contains an array of IP address ranges that can be used by subnets.
        * `address_prefixes` (`list`) - A list of address blocks reserved for this virtual network in CIDR notation.

      * `dhcp_options` (`dict`) - The dhcpOptions that contains an array of DNS servers available to VMs deployed in the virtual network.
        * `dns_servers` (`list`) - The list of DNS servers IP addresses.

      * `enable_ddos_protection` (`bool`) - Indicates if DDoS protection is enabled for all the protected resources in a Virtual Network.
      * `enable_vm_protection` (`bool`) - Indicates if Vm protection is enabled for all the subnets in a Virtual Network.
      * `provisioning_state` (`str`) - The provisioning state of the PublicIP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
      * `resource_guid` (`str`) - The resourceGuid property of the Virtual Network resource.
      * `subnets` (`list`) - A list of subnets in a Virtual Network.
        * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
        * `id` (`str`) - Resource ID.
        * `name` (`str`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
        * `properties` (`dict`) - Properties of the subnet.
          * `address_prefix` (`str`) - The address prefix for the subnet.
          * `ip_configurations` (`list`) - Gets an array of references to the network interface IP configurations using subnet.
            * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
            * `id` (`str`) - Resource ID.
            * `name` (`str`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
            * `properties` (`dict`) - Properties of the IP configuration
              * `private_ip_address` (`str`) - The private IP address of the IP configuration.
              * `private_ip_allocation_method` (`str`) - The private IP allocation method. Possible values are 'Static' and 'Dynamic'.
              * `provisioning_state` (`str`) - Gets the provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
              * `public_ip_address` (`dict`) - The reference of the public IP resource.
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
                  * `ip_address` (`str`) - The IP address associated with the public IP address resource.
                  * `ip_configuration` (`dict`) - The IP configuration associated with the public IP address.
                  * `provisioning_state` (`str`) - The provisioning state of the PublicIP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
                  * `public_ip_address_version` (`str`) - The public IP address version. Possible values are: 'IPv4' and 'IPv6'.
                  * `public_ip_allocation_method` (`str`) - The public IP allocation method. Possible values are: 'Static' and 'Dynamic'.
                  * `resource_guid` (`str`) - The resource GUID property of the public IP resource.

                * `sku` (`dict`) - The public IP address SKU.
                  * `name` (`str`) - Name of a public IP address SKU.

                * `tags` (`dict`) - Resource tags.
                * `type` (`str`) - Resource type.
                * `zones` (`list`) - A list of availability zones denoting the IP allocated for the resource needs to come from.

              * `subnet` (`dict`) - The reference of the subnet resource.

          * `network_security_group` (`dict`) - The reference of the NetworkSecurityGroup resource.
            * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
            * `id` (`str`) - Resource ID.
            * `location` (`str`) - Resource location.
            * `name` (`str`) - Resource name.
            * `properties` (`dict`) - Properties of the network security group
              * `default_security_rules` (`list`) - The default security rules of network security group.
                * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
                * `id` (`str`) - Resource ID.
                * `name` (`str`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
                * `properties` (`dict`) - Properties of the security rule
                  * `access` (`str`) - The network traffic is allowed or denied. Possible values are: 'Allow' and 'Deny'.
                  * `description` (`str`) - A description for this rule. Restricted to 140 chars.
                  * `destination_address_prefix` (`str`) - The destination address prefix. CIDR or destination IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used.
                  * `destination_address_prefixes` (`list`) - The destination address prefixes. CIDR or destination IP ranges.
                  * `destination_application_security_groups` (`list`) - The application security group specified as destination.
                    * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
                    * `id` (`str`) - Resource ID.
                    * `location` (`str`) - Resource location.
                    * `name` (`str`) - Resource name.
                    * `properties` (`dict`) - Properties of the application security group.
                      * `provisioning_state` (`str`) - The provisioning state of the application security group resource. Possible values are: 'Succeeded', 'Updating', 'Deleting', and 'Failed'.
                      * `resource_guid` (`str`) - The resource GUID property of the application security group resource. It uniquely identifies a resource, even if the user changes its name or migrate the resource across subscriptions or resource groups.

                    * `tags` (`dict`) - Resource tags.
                    * `type` (`str`) - Resource type.

                  * `destination_port_range` (`str`) - The destination port or range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.
                  * `destination_port_ranges` (`list`) - The destination port ranges.
                  * `direction` (`str`) - The direction of the rule. The direction specifies if rule will be evaluated on incoming or outgoing traffic. Possible values are: 'Inbound' and 'Outbound'.
                  * `priority` (`float`) - The priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
                  * `protocol` (`str`) - Network protocol this rule applies to. Possible values are 'Tcp', 'Udp', and '*'.
                  * `provisioning_state` (`str`) - The provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
                  * `source_address_prefix` (`str`) - The CIDR or source IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used. If this is an ingress rule, specifies where network traffic originates from. 
                  * `source_address_prefixes` (`list`) - The CIDR or source IP ranges.
                  * `source_application_security_groups` (`list`) - The application security group specified as source.
                  * `source_port_range` (`str`) - The source port or range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.
                  * `source_port_ranges` (`list`) - The source port ranges.

              * `network_interfaces` (`list`) - A collection of references to network interfaces.
                * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
                * `id` (`str`) - Resource ID.
                * `location` (`str`) - Resource location.
                * `name` (`str`) - Resource name.
                * `properties` (`dict`) - Properties of the network interface.
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
                    * `properties` (`dict`) - Network interface IP configuration properties.
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

                        * `type` (`str`) - Type of the resource.

                      * `application_security_groups` (`list`) - Application security groups in which the IP configuration is included.
                      * `load_balancer_backend_address_pools` (`list`) - The reference of LoadBalancerBackendAddressPool resource.
                        * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
                        * `id` (`str`) - Resource ID.
                        * `name` (`str`) - Gets name of the resource that is unique within a resource group. This name can be used to access the resource.
                        * `properties` (`dict`) - Properties of load balancer backend address pool.
                          * `backend_ip_configurations` (`list`) - Gets collection of references to IP addresses defined in network interfaces.
                          * `load_balancing_rules` (`list`) - Gets load balancing rules that use this backend address pool.
                            * `id` (`str`) - Resource ID.

                          * `outbound_nat_rule` (`dict`) - Gets outbound rules that use this backend address pool.
                          * `provisioning_state` (`str`) - Get provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.

                      * `load_balancer_inbound_nat_rules` (`list`) - A list of references of LoadBalancerInboundNatRules.
                        * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
                        * `id` (`str`) - Resource ID.
                        * `name` (`str`) - Gets name of the resource that is unique within a resource group. This name can be used to access the resource.
                        * `properties` (`dict`) - Properties of load balancer inbound nat rule.
                          * `backend_ip_configuration` (`dict`) - A reference to a private IP address defined on a network interface of a VM. Traffic sent to the frontend port of each of the frontend IP configurations is forwarded to the backend IP.
                          * `backend_port` (`float`) - The port used for the internal endpoint. Acceptable values range from 1 to 65535.
                          * `enable_floating_ip` (`bool`) - Configures a virtual machine's endpoint for the floating IP capability required to configure a SQL AlwaysOn Availability Group. This setting is required when using the SQL AlwaysOn Availability Groups in SQL server. This setting can't be changed after you create the endpoint.
                          * `frontend_ip_configuration` (`dict`) - A reference to frontend IP addresses.
                          * `frontend_port` (`float`) - The port for the external endpoint. Port numbers for each rule must be unique within the Load Balancer. Acceptable values range from 1 to 65534.
                          * `idle_timeout_in_minutes` (`float`) - The timeout for the TCP idle connection. The value can be set between 4 and 30 minutes. The default value is 4 minutes. This element is only used when the protocol is set to TCP.
                          * `protocol` (`str`) - The transport protocol for the endpoint. Possible values are 'Udp' or 'Tcp' or 'All.'
                          * `provisioning_state` (`str`) - Gets the provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.

                      * `primary` (`bool`) - Gets whether this is a primary customer address on the network interface.
                      * `private_ip_address` (`str`) - Private IP address of the IP configuration.
                      * `private_ip_address_version` (`str`) - Available from Api-Version 2016-03-30 onwards, it represents whether the specific ipconfiguration is IPv4 or IPv6. Default is taken as IPv4.  Possible values are: 'IPv4' and 'IPv6'.
                      * `private_ip_allocation_method` (`str`) - Defines how a private IP address is assigned. Possible values are: 'Static' and 'Dynamic'.
                      * `provisioning_state` (`str`) - The provisioning state of the network interface IP configuration. Possible values are: 'Updating', 'Deleting', and 'Failed'.
                      * `public_ip_address` (`dict`) - Public IP address bound to the IP configuration.
                      * `subnet` (`dict`) - Subnet bound to the IP configuration.

                  * `mac_address` (`str`) - The MAC address of the network interface.
                  * `network_security_group` (`dict`) - The reference of the NetworkSecurityGroup resource.
                  * `primary` (`bool`) - Gets whether this is a primary network interface on a virtual machine.
                  * `provisioning_state` (`str`) - The provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
                  * `resource_guid` (`str`) - The resource GUID property of the network interface resource.
                  * `virtual_machine` (`dict`) - The reference of a virtual machine.

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
            * `properties` (`dict`) - Resource navigation link properties format.
              * `link` (`str`) - Link to the external resource
              * `linked_resource_type` (`str`) - Resource type of the linked resource.
              * `provisioning_state` (`str`) - Provisioning state of the ResourceNavigationLink resource.

          * `route_table` (`dict`) - The reference of the RouteTable resource.
            * `etag` (`str`) - Gets a unique read-only string that changes whenever the resource is updated.
            * `id` (`str`) - Resource ID.
            * `location` (`str`) - Resource location.
            * `name` (`str`) - Resource name.
            * `properties` (`dict`) - Properties of the route table.
              * `disable_bgp_route_propagation` (`bool`) - Gets or sets whether to disable the routes learned by BGP on that route table. True means disable.
              * `provisioning_state` (`str`) - The provisioning state of the resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
              * `routes` (`list`) - Collection of routes contained within a route table.
                * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
                * `id` (`str`) - Resource ID.
                * `name` (`str`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
                * `properties` (`dict`) - Properties of the route.
                  * `address_prefix` (`str`) - The destination CIDR to which the route applies.
                  * `next_hop_ip_address` (`str`) - The IP address packets should be forwarded to. Next hop values are only allowed in routes where the next hop type is VirtualAppliance.
                  * `next_hop_type` (`str`) - The type of Azure hop the packet should be sent to. Possible values are: 'VirtualNetworkGateway', 'VnetLocal', 'Internet', 'VirtualAppliance', and 'None'
                  * `provisioning_state` (`str`) - The provisioning state of the resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.

              * `subnets` (`list`) - A collection of references to subnets.

            * `tags` (`dict`) - Resource tags.
            * `type` (`str`) - Resource type.

          * `service_endpoints` (`list`) - An array of service endpoints.
            * `locations` (`list`) - A list of locations.
            * `provisioning_state` (`str`) - The provisioning state of the resource.
            * `service` (`str`) - The type of the endpoint service.

      * `virtual_network_peerings` (`list`) - A list of peerings in a Virtual Network.
        * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
        * `id` (`str`) - Resource ID.
        * `name` (`str`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
        * `properties` (`dict`) - Properties of the virtual network peering.
          * `allow_forwarded_traffic` (`bool`) - Whether the forwarded traffic from the VMs in the remote virtual network will be allowed/disallowed.
          * `allow_gateway_transit` (`bool`) - If gateway links can be used in remote virtual networking to link to this virtual network.
          * `allow_virtual_network_access` (`bool`) - Whether the VMs in the linked virtual network space would be able to access all the VMs in local Virtual network space.
          * `peering_state` (`str`) - The status of the virtual network peering. Possible values are 'Initiated', 'Connected', and 'Disconnected'.
          * `provisioning_state` (`str`) - The provisioning state of the resource.
          * `remote_address_space` (`dict`) - The reference of the remote virtual network address space.
          * `remote_virtual_network` (`dict`) - The reference of the remote virtual network. The remote virtual network can be in the same or different region (preview). See here to register for the preview and learn more (https://docs.microsoft.com/en-us/azure/virtual-network/virtual-network-create-peering).
          * `use_remote_gateways` (`bool`) - If remote gateways can be used on this virtual network. If the flag is set to true, and allowGatewayTransit on remote peering is also true, virtual network will use gateways of remote virtual network for transit. Only one peering can have this flag set to true. This flag cannot be set if virtual network already has a gateway.
    """
    tags: pulumi.Output[dict]
    """
    Resource tags.
    """
    type: pulumi.Output[str]
    """
    Resource type.
    """
    def __init__(__self__, resource_name, opts=None, address_space=None, dhcp_options=None, enable_ddos_protection=None, enable_vm_protection=None, etag=None, id=None, location=None, name=None, provisioning_state=None, resource_group_name=None, resource_guid=None, subnets=None, tags=None, virtual_network_peerings=None, __props__=None, __name__=None, __opts__=None):
        """
        Virtual Network resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] address_space: The AddressSpace that contains an array of IP address ranges that can be used by subnets.
        :param pulumi.Input[dict] dhcp_options: The dhcpOptions that contains an array of DNS servers available to VMs deployed in the virtual network.
        :param pulumi.Input[bool] enable_ddos_protection: Indicates if DDoS protection is enabled for all the protected resources in a Virtual Network.
        :param pulumi.Input[bool] enable_vm_protection: Indicates if Vm protection is enabled for all the subnets in a Virtual Network.
        :param pulumi.Input[str] etag: Gets a unique read-only string that changes whenever the resource is updated.
        :param pulumi.Input[str] id: Resource ID.
        :param pulumi.Input[str] location: Resource location.
        :param pulumi.Input[str] name: The name of the virtual network.
        :param pulumi.Input[str] provisioning_state: The provisioning state of the PublicIP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[str] resource_guid: The resourceGuid property of the Virtual Network resource.
        :param pulumi.Input[list] subnets: A list of subnets in a Virtual Network.
        :param pulumi.Input[dict] tags: Resource tags.
        :param pulumi.Input[list] virtual_network_peerings: A list of peerings in a Virtual Network.

        The **address_space** object supports the following:

          * `address_prefixes` (`pulumi.Input[list]`) - A list of address blocks reserved for this virtual network in CIDR notation.

        The **dhcp_options** object supports the following:

          * `dns_servers` (`pulumi.Input[list]`) - The list of DNS servers IP addresses.

        The **subnets** object supports the following:

          * `address_prefix` (`pulumi.Input[str]`) - The address prefix for the subnet.
          * `etag` (`pulumi.Input[str]`) - A unique read-only string that changes whenever the resource is updated.
          * `id` (`pulumi.Input[str]`) - Resource ID.
          * `name` (`pulumi.Input[str]`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
          * `network_security_group` (`pulumi.Input[dict]`) - The reference of the NetworkSecurityGroup resource.
            * `default_security_rules` (`pulumi.Input[list]`) - The default security rules of network security group.
              * `access` (`pulumi.Input[str]`) - The network traffic is allowed or denied. Possible values are: 'Allow' and 'Deny'.
              * `description` (`pulumi.Input[str]`) - A description for this rule. Restricted to 140 chars.
              * `destination_address_prefix` (`pulumi.Input[str]`) - The destination address prefix. CIDR or destination IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used.
              * `destination_address_prefixes` (`pulumi.Input[list]`) - The destination address prefixes. CIDR or destination IP ranges.
              * `destination_application_security_groups` (`pulumi.Input[list]`) - The application security group specified as destination.
                * `id` (`pulumi.Input[str]`) - Resource ID.
                * `location` (`pulumi.Input[str]`) - Resource location.
                * `tags` (`pulumi.Input[dict]`) - Resource tags.

              * `destination_port_range` (`pulumi.Input[str]`) - The destination port or range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.
              * `destination_port_ranges` (`pulumi.Input[list]`) - The destination port ranges.
              * `direction` (`pulumi.Input[str]`) - The direction of the rule. The direction specifies if rule will be evaluated on incoming or outgoing traffic. Possible values are: 'Inbound' and 'Outbound'.
              * `etag` (`pulumi.Input[str]`) - A unique read-only string that changes whenever the resource is updated.
              * `id` (`pulumi.Input[str]`) - Resource ID.
              * `name` (`pulumi.Input[str]`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
              * `priority` (`pulumi.Input[float]`) - The priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
              * `protocol` (`pulumi.Input[str]`) - Network protocol this rule applies to. Possible values are 'Tcp', 'Udp', and '*'.
              * `provisioning_state` (`pulumi.Input[str]`) - The provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
              * `source_address_prefix` (`pulumi.Input[str]`) - The CIDR or source IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used. If this is an ingress rule, specifies where network traffic originates from. 
              * `source_address_prefixes` (`pulumi.Input[list]`) - The CIDR or source IP ranges.
              * `source_application_security_groups` (`pulumi.Input[list]`) - The application security group specified as source.
              * `source_port_range` (`pulumi.Input[str]`) - The source port or range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.
              * `source_port_ranges` (`pulumi.Input[list]`) - The source port ranges.

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
            * `disable_bgp_route_propagation` (`pulumi.Input[bool]`) - Gets or sets whether to disable the routes learned by BGP on that route table. True means disable.
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

          * `service_endpoints` (`pulumi.Input[list]`) - An array of service endpoints.
            * `locations` (`pulumi.Input[list]`) - A list of locations.
            * `provisioning_state` (`pulumi.Input[str]`) - The provisioning state of the resource.
            * `service` (`pulumi.Input[str]`) - The type of the endpoint service.

        The **virtual_network_peerings** object supports the following:

          * `allow_forwarded_traffic` (`pulumi.Input[bool]`) - Whether the forwarded traffic from the VMs in the remote virtual network will be allowed/disallowed.
          * `allow_gateway_transit` (`pulumi.Input[bool]`) - If gateway links can be used in remote virtual networking to link to this virtual network.
          * `allow_virtual_network_access` (`pulumi.Input[bool]`) - Whether the VMs in the linked virtual network space would be able to access all the VMs in local Virtual network space.
          * `etag` (`pulumi.Input[str]`) - A unique read-only string that changes whenever the resource is updated.
          * `id` (`pulumi.Input[str]`) - Resource ID.
          * `name` (`pulumi.Input[str]`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
          * `peering_state` (`pulumi.Input[str]`) - The status of the virtual network peering. Possible values are 'Initiated', 'Connected', and 'Disconnected'.
          * `provisioning_state` (`pulumi.Input[str]`) - The provisioning state of the resource.
          * `remote_address_space` (`pulumi.Input[dict]`) - The reference of the remote virtual network address space.
          * `remote_virtual_network` (`pulumi.Input[dict]`) - The reference of the remote virtual network. The remote virtual network can be in the same or different region (preview). See here to register for the preview and learn more (https://docs.microsoft.com/en-us/azure/virtual-network/virtual-network-create-peering).
            * `id` (`pulumi.Input[str]`) - Resource ID.

          * `use_remote_gateways` (`pulumi.Input[bool]`) - If remote gateways can be used on this virtual network. If the flag is set to true, and allowGatewayTransit on remote peering is also true, virtual network will use gateways of remote virtual network for transit. Only one peering can have this flag set to true. This flag cannot be set if virtual network already has a gateway.
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

            __props__['address_space'] = address_space
            __props__['dhcp_options'] = dhcp_options
            __props__['enable_ddos_protection'] = enable_ddos_protection
            __props__['enable_vm_protection'] = enable_vm_protection
            __props__['etag'] = etag
            __props__['id'] = id
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['provisioning_state'] = provisioning_state
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['resource_guid'] = resource_guid
            __props__['subnets'] = subnets
            __props__['tags'] = tags
            __props__['virtual_network_peerings'] = virtual_network_peerings
            __props__['properties'] = None
            __props__['type'] = None
        super(VirtualNetwork, __self__).__init__(
            'azurerm:network/v20171001:VirtualNetwork',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing VirtualNetwork resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return VirtualNetwork(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop