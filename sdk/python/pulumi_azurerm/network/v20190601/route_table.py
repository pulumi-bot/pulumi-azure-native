# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from ... import _utilities, _tables


class RouteTable(pulumi.CustomResource):
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
    Properties of the route table.
      * `disable_bgp_route_propagation` (`bool`) - Gets or sets whether to disable the routes learned by BGP on that route table. True means disable.
      * `provisioning_state` (`str`) - The provisioning state of the resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
      * `routes` (`list`) - Collection of routes contained within a route table.
        * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
        * `id` (`str`) - Resource ID.
        * `name` (`str`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
        * `properties` (`dict`) - Properties of the route.
          * `address_prefix` (`str`) - The destination CIDR to which the route applies.
          * `next_hop_ip_address` (`str`) - The IP address packets should be forwarded to. Next hop values are only allowed in routes where the next hop type is VirtualAppliance.
          * `next_hop_type` (`str`) - The type of Azure hop the packet should be sent to.
          * `provisioning_state` (`str`) - The provisioning state of the resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.

      * `subnets` (`list`) - A collection of references to subnets.
        * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
        * `id` (`str`) - Resource ID.
        * `name` (`str`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
        * `properties` (`dict`) - Properties of the subnet.
          * `address_prefix` (`str`) - The address prefix for the subnet.
          * `address_prefixes` (`list`) - List of address prefixes for the subnet.
          * `delegations` (`list`) - Gets an array of references to the delegations on the subnet.
            * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
            * `id` (`str`) - Resource ID.
            * `name` (`str`) - The name of the resource that is unique within a subnet. This name can be used to access the resource.
            * `properties` (`dict`) - Properties of the subnet.
              * `actions` (`list`) - Describes the actions permitted to the service upon delegation.
              * `provisioning_state` (`str`) - The provisioning state of the resource.
              * `service_name` (`str`) - The name of the service to whom the subnet should be delegated (e.g. Microsoft.Sql/servers).

          * `ip_configuration_profiles` (`list`) - Array of IP configuration profiles which reference this subnet.
            * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
            * `id` (`str`) - Resource ID.
            * `name` (`str`) - The name of the resource. This name can be used to access the resource.
            * `properties` (`dict`) - Properties of the IP configuration profile.
              * `provisioning_state` (`str`) - The provisioning state of the resource.
              * `subnet` (`dict`) - The reference of the subnet resource to create a container network interface ip configuration.

            * `type` (`str`) - Sub Resource type.

          * `ip_configurations` (`list`) - Gets an array of references to the network interface IP configurations using subnet.
            * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
            * `id` (`str`) - Resource ID.
            * `name` (`str`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
            * `properties` (`dict`) - Properties of the IP configuration.
              * `private_ip_address` (`str`) - The private IP address of the IP configuration.
              * `private_ip_allocation_method` (`str`) - The private IP address allocation method.
              * `provisioning_state` (`str`) - Gets the provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
              * `public_ip_address` (`dict`) - The reference of the public IP resource.
                * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
                * `id` (`str`) - Resource ID.
                * `location` (`str`) - Resource location.
                * `name` (`str`) - Resource name.
                * `properties` (`dict`) - Public IP address properties.
                  * `ddos_settings` (`dict`) - The DDoS protection custom policy associated with the public IP address.
                    * `ddos_custom_policy` (`dict`) - The DDoS custom policy associated with the public IP.
                      * `id` (`str`) - Resource ID.

                    * `protection_coverage` (`str`) - The DDoS protection policy customizability of the public IP. Only standard coverage will have the ability to be customized.

                  * `dns_settings` (`dict`) - The FQDN of the DNS record associated with the public IP address.
                    * `domain_name_label` (`str`) - Gets or sets the Domain name label.The concatenation of the domain name label and the regionalized DNS zone make up the fully qualified domain name associated with the public IP address. If a domain name label is specified, an A DNS record is created for the public IP in the Microsoft Azure DNS system.
                    * `fqdn` (`str`) - Gets the FQDN, Fully qualified domain name of the A DNS record associated with the public IP. This is the concatenation of the domainNameLabel and the regionalized DNS zone.
                    * `reverse_fqdn` (`str`) - Gets or Sets the Reverse FQDN. A user-visible, fully qualified domain name that resolves to this public IP address. If the reverseFqdn is specified, then a PTR DNS record is created pointing from the IP address in the in-addr.arpa domain to the reverse FQDN.

                  * `idle_timeout_in_minutes` (`float`) - The idle timeout of the public IP address.
                  * `ip_address` (`str`) - The IP address associated with the public IP address resource.
                  * `ip_configuration` (`dict`) - The IP configuration associated with the public IP address.
                  * `ip_tags` (`list`) - The list of tags associated with the public IP address.
                    * `ip_tag_type` (`str`) - Gets or sets the ipTag type: Example FirstPartyUsage.
                    * `tag` (`str`) - Gets or sets value of the IpTag associated with the public IP. Example SQL, Storage etc.

                  * `provisioning_state` (`str`) - The provisioning state of the PublicIP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
                  * `public_ip_address_version` (`str`) - The public IP address version.
                  * `public_ip_allocation_method` (`str`) - The public IP address allocation method.
                  * `public_ip_prefix` (`dict`) - The Public IP Prefix this Public IP Address should be allocated from.
                  * `resource_guid` (`str`) - The resource GUID property of the public IP resource.

                * `sku` (`dict`) - The public IP address SKU.
                  * `name` (`str`) - Name of a public IP address SKU.

                * `tags` (`dict`) - Resource tags.
                * `type` (`str`) - Resource type.
                * `zones` (`list`) - A list of availability zones denoting the IP allocated for the resource needs to come from.

              * `subnet` (`dict`) - The reference of the subnet resource.

          * `nat_gateway` (`dict`) - Nat gateway associated with this subnet.
          * `network_security_group` (`dict`) - The reference of the NetworkSecurityGroup resource.
            * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
            * `id` (`str`) - Resource ID.
            * `location` (`str`) - Resource location.
            * `name` (`str`) - Resource name.
            * `properties` (`dict`) - Properties of the network security group.
              * `default_security_rules` (`list`) - The default security rules of network security group.
                * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
                * `id` (`str`) - Resource ID.
                * `name` (`str`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
                * `properties` (`dict`) - Properties of the security rule.
                  * `access` (`str`) - The network traffic is allowed or denied.
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
                  * `direction` (`str`) - The direction of the rule. The direction specifies if rule will be evaluated on incoming or outgoing traffic.
                  * `priority` (`float`) - The priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
                  * `protocol` (`str`) - Network protocol this rule applies to.
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
                  * `hosted_workloads` (`list`) - A list of references to linked BareMetal resources.
                  * `ip_configurations` (`list`) - A list of IPConfigurations of the network interface.
                    * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
                    * `id` (`str`) - Resource ID.
                    * `name` (`str`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
                    * `properties` (`dict`) - Network interface IP configuration properties.
                      * `application_gateway_backend_address_pools` (`list`) - The reference of ApplicationGatewayBackendAddressPool resource.
                        * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
                        * `id` (`str`) - Resource ID.
                        * `name` (`str`) - Name of the backend address pool that is unique within an Application Gateway.
                        * `properties` (`dict`) - Properties of the application gateway backend address pool.
                          * `backend_addresses` (`list`) - Backend addresses.
                            * `fqdn` (`str`) - Fully qualified domain name (FQDN).
                            * `ip_address` (`str`) - IP address.

                          * `backend_ip_configurations` (`list`) - Collection of references to IPs defined in network interfaces.
                          * `provisioning_state` (`str`) - Provisioning state of the backend address pool resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.

                        * `type` (`str`) - Type of the resource.

                      * `application_security_groups` (`list`) - Application security groups in which the IP configuration is included.
                      * `load_balancer_backend_address_pools` (`list`) - The reference of LoadBalancerBackendAddressPool resource.
                        * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
                        * `id` (`str`) - Resource ID.
                        * `name` (`str`) - Gets name of the resource that is unique within the set of backend address pools used by the load balancer. This name can be used to access the resource.
                        * `properties` (`dict`) - Properties of load balancer backend address pool.
                          * `backend_ip_configurations` (`list`) - Gets collection of references to IP addresses defined in network interfaces.
                          * `load_balancing_rules` (`list`) - Gets load balancing rules that use this backend address pool.
                          * `outbound_rule` (`dict`) - Gets outbound rules that use this backend address pool.
                          * `outbound_rules` (`list`) - Gets outbound rules that use this backend address pool.
                          * `provisioning_state` (`str`) - Get provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.

                        * `type` (`str`) - Type of the resource.

                      * `load_balancer_inbound_nat_rules` (`list`) - A list of references of LoadBalancerInboundNatRules.
                        * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
                        * `id` (`str`) - Resource ID.
                        * `name` (`str`) - Gets name of the resource that is unique within the set of inbound NAT rules used by the load balancer. This name can be used to access the resource.
                        * `properties` (`dict`) - Properties of load balancer inbound nat rule.
                          * `backend_ip_configuration` (`dict`) - A reference to a private IP address defined on a network interface of a VM. Traffic sent to the frontend port of each of the frontend IP configurations is forwarded to the backend IP.
                          * `backend_port` (`float`) - The port used for the internal endpoint. Acceptable values range from 1 to 65535.
                          * `enable_floating_ip` (`bool`) - Configures a virtual machine's endpoint for the floating IP capability required to configure a SQL AlwaysOn Availability Group. This setting is required when using the SQL AlwaysOn Availability Groups in SQL server. This setting can't be changed after you create the endpoint.
                          * `enable_tcp_reset` (`bool`) - Receive bidirectional TCP Reset on TCP flow idle timeout or unexpected connection termination. This element is only used when the protocol is set to TCP.
                          * `frontend_ip_configuration` (`dict`) - A reference to frontend IP addresses.
                          * `frontend_port` (`float`) - The port for the external endpoint. Port numbers for each rule must be unique within the Load Balancer. Acceptable values range from 1 to 65534.
                          * `idle_timeout_in_minutes` (`float`) - The timeout for the TCP idle connection. The value can be set between 4 and 30 minutes. The default value is 4 minutes. This element is only used when the protocol is set to TCP.
                          * `protocol` (`str`) - The reference to the transport protocol used by the load balancing rule.
                          * `provisioning_state` (`str`) - Gets the provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.

                        * `type` (`str`) - Type of the resource.

                      * `primary` (`bool`) - Gets whether this is a primary customer address on the network interface.
                      * `private_ip_address` (`str`) - Private IP address of the IP configuration.
                      * `private_ip_address_version` (`str`) - Available from Api-Version 2016-03-30 onwards, it represents whether the specific ipconfiguration is IPv4 or IPv6. Default is taken as IPv4.
                      * `private_ip_allocation_method` (`str`) - The private IP address allocation method.
                      * `provisioning_state` (`str`) - The provisioning state of the network interface IP configuration. Possible values are: 'Updating', 'Deleting', and 'Failed'.
                      * `public_ip_address` (`dict`) - Public IP address bound to the IP configuration.
                      * `subnet` (`dict`) - Subnet bound to the IP configuration.
                      * `virtual_network_taps` (`list`) - The reference to Virtual Network Taps.
                        * `etag` (`str`) - Gets a unique read-only string that changes whenever the resource is updated.
                        * `id` (`str`) - Resource ID.
                        * `location` (`str`) - Resource location.
                        * `name` (`str`) - Resource name.
                        * `properties` (`dict`) - Virtual Network Tap Properties.
                          * `destination_load_balancer_front_end_ip_configuration` (`dict`) - The reference to the private IP address on the internal Load Balancer that will receive the tap.
                            * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
                            * `id` (`str`) - Resource ID.
                            * `name` (`str`) - The name of the resource that is unique within the set of frontend IP configurations used by the load balancer. This name can be used to access the resource.
                            * `properties` (`dict`) - Properties of the load balancer probe.
                              * `inbound_nat_pools` (`list`) - Read only. Inbound pools URIs that use this frontend IP.
                              * `inbound_nat_rules` (`list`) - Read only. Inbound rules URIs that use this frontend IP.
                              * `load_balancing_rules` (`list`) - Gets load balancing rules URIs that use this frontend IP.
                              * `outbound_rules` (`list`) - Read only. Outbound rules URIs that use this frontend IP.
                              * `private_ip_address` (`str`) - The private IP address of the IP configuration.
                              * `private_ip_address_version` (`str`) - It represents whether the specific ipconfiguration is IPv4 or IPv6. Default is taken as IPv4.
                              * `private_ip_allocation_method` (`str`) - The Private IP allocation method.
                              * `provisioning_state` (`str`) - Gets the provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
                              * `public_ip_address` (`dict`) - The reference of the Public IP resource.
                              * `public_ip_prefix` (`dict`) - The reference of the Public IP Prefix resource.
                              * `subnet` (`dict`) - The reference of the subnet resource.

                            * `type` (`str`) - Type of the resource.
                            * `zones` (`list`) - A list of availability zones denoting the IP allocated for the resource needs to come from.

                          * `destination_network_interface_ip_configuration` (`dict`) - The reference to the private IP Address of the collector nic that will receive the tap.
                          * `destination_port` (`float`) - The VXLAN destination port that will receive the tapped traffic.
                          * `network_interface_tap_configurations` (`list`) - Specifies the list of resource IDs for the network interface IP configuration that needs to be tapped.
                            * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
                            * `id` (`str`) - Resource ID.
                            * `name` (`str`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
                            * `properties` (`dict`) - Properties of the Virtual Network Tap configuration.
                              * `provisioning_state` (`str`) - The provisioning state of the network interface tap configuration. Possible values are: 'Updating', 'Deleting', and 'Failed'.
                              * `virtual_network_tap` (`dict`) - The reference of the Virtual Network Tap resource.

                            * `type` (`str`) - Sub Resource type.

                          * `provisioning_state` (`str`) - The provisioning state of the virtual network tap. Possible values are: 'Updating', 'Deleting', and 'Failed'.
                          * `resource_guid` (`str`) - The resourceGuid property of the virtual network tap.

                        * `tags` (`dict`) - Resource tags.
                        * `type` (`str`) - Resource type.

                  * `mac_address` (`str`) - The MAC address of the network interface.
                  * `network_security_group` (`dict`) - The reference of the NetworkSecurityGroup resource.
                  * `primary` (`bool`) - Gets whether this is a primary network interface on a virtual machine.
                  * `private_endpoint` (`dict`) - A reference to the private endpoint to which the network interface is linked.
                    * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
                    * `id` (`str`) - Resource ID.
                    * `location` (`str`) - Resource location.
                    * `name` (`str`) - Resource name.
                    * `properties` (`dict`) - Properties of the private endpoint.
                      * `manual_private_link_service_connections` (`list`) - A grouping of information about the connection to the remote resource. Used when the network admin does not have access to approve connections to the remote resource.
                        * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
                        * `id` (`str`) - Resource ID.
                        * `name` (`str`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
                        * `properties` (`dict`) - Properties of the private link service connection.
                          * `group_ids` (`list`) - The ID(s) of the group(s) obtained from the remote resource that this private endpoint should connect to.
                          * `private_link_service_connection_state` (`dict`) - A collection of read-only information about the state of the connection to the remote resource.
                            * `actions_required` (`str`) - A message indicating if changes on the service provider require any updates on the consumer.
                            * `description` (`str`) - The reason for approval/rejection of the connection.
                            * `status` (`str`) - Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.

                          * `private_link_service_id` (`str`) - The resource id of private link service.
                          * `provisioning_state` (`str`) - The provisioning state of the private link service connection.
                          * `request_message` (`str`) - A message passed to the owner of the remote resource with this connection request. Restricted to 140 chars.

                        * `type` (`str`) - The resource type.

                      * `network_interfaces` (`list`) - Gets an array of references to the network interfaces created for this private endpoint.
                      * `private_link_service_connections` (`list`) - A grouping of information about the connection to the remote resource.
                      * `provisioning_state` (`str`) - The provisioning state of the private endpoint.
                      * `subnet` (`dict`) - The ID of the subnet from which the private IP will be allocated.

                    * `tags` (`dict`) - Resource tags.
                    * `type` (`str`) - Resource type.

                  * `provisioning_state` (`str`) - The provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
                  * `resource_guid` (`str`) - The resource GUID property of the network interface resource.
                  * `tap_configurations` (`list`) - A list of TapConfigurations of the network interface.
                  * `virtual_machine` (`dict`) - The reference of a virtual machine.

                * `tags` (`dict`) - Resource tags.
                * `type` (`str`) - Resource type.

              * `provisioning_state` (`str`) - The provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
              * `resource_guid` (`str`) - The resource GUID property of the network security group resource.
              * `security_rules` (`list`) - A collection of security rules of the network security group.
              * `subnets` (`list`) - A collection of references to subnets.

            * `tags` (`dict`) - Resource tags.
            * `type` (`str`) - Resource type.

          * `private_endpoint_network_policies` (`str`) - Enable or Disable apply network policies on private end point in the subnet.
          * `private_endpoints` (`list`) - An array of references to private endpoints.
          * `private_link_service_network_policies` (`str`) - Enable or Disable apply network policies on private link service in the subnet.
          * `provisioning_state` (`str`) - The provisioning state of the resource.
          * `purpose` (`str`) - A read-only string identifying the intention of use for this subnet based on delegations and other user-defined properties.
          * `resource_navigation_links` (`list`) - Gets an array of references to the external resources using subnet.
            * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
            * `id` (`str`) - Resource ID.
            * `name` (`str`) - Name of the resource that is unique within a resource group. This name can be used to access the resource.
            * `properties` (`dict`) - Resource navigation link properties format.
              * `link` (`str`) - Link to the external resource.
              * `linked_resource_type` (`str`) - Resource type of the linked resource.
              * `provisioning_state` (`str`) - Provisioning state of the ResourceNavigationLink resource.

            * `type` (`str`) - Resource type.

          * `route_table` (`dict`) - The reference of the RouteTable resource.
            * `etag` (`str`) - Gets a unique read-only string that changes whenever the resource is updated.
            * `id` (`str`) - Resource ID.
            * `location` (`str`) - Resource location.
            * `name` (`str`) - Resource name.
            * `properties` (`dict`) - Properties of the route table.
            * `tags` (`dict`) - Resource tags.
            * `type` (`str`) - Resource type.

          * `service_association_links` (`list`) - Gets an array of references to services injecting into this subnet.
            * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
            * `id` (`str`) - Resource ID.
            * `name` (`str`) - Name of the resource that is unique within a resource group. This name can be used to access the resource.
            * `properties` (`dict`) - Resource navigation link properties format.
              * `allow_delete` (`bool`) - If true, the resource can be deleted.
              * `link` (`str`) - Link to the external resource.
              * `linked_resource_type` (`str`) - Resource type of the linked resource.
              * `locations` (`list`) - A list of locations.
              * `provisioning_state` (`str`) - Provisioning state of the ServiceAssociationLink resource.

            * `type` (`str`) - Resource type.

          * `service_endpoint_policies` (`list`) - An array of service endpoint policies.
            * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
            * `id` (`str`) - Resource ID.
            * `location` (`str`) - Resource location.
            * `name` (`str`) - Resource name.
            * `properties` (`dict`) - Properties of the service end point policy.
              * `provisioning_state` (`str`) - The provisioning state of the service endpoint policy. Possible values are: 'Updating', 'Deleting', and 'Failed'.
              * `resource_guid` (`str`) - The resource GUID property of the service endpoint policy resource.
              * `service_endpoint_policy_definitions` (`list`) - A collection of service endpoint policy definitions of the service endpoint policy.
                * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated.
                * `id` (`str`) - Resource ID.
                * `name` (`str`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
                * `properties` (`dict`) - Properties of the service endpoint policy definition.
                  * `description` (`str`) - A description for this rule. Restricted to 140 chars.
                  * `provisioning_state` (`str`) - The provisioning state of the service end point policy definition. Possible values are: 'Updating', 'Deleting', and 'Failed'.
                  * `service` (`str`) - Service endpoint name.
                  * `service_resources` (`list`) - A list of service resources.

              * `subnets` (`list`) - A collection of references to subnets.

            * `tags` (`dict`) - Resource tags.
            * `type` (`str`) - Resource type.

          * `service_endpoints` (`list`) - An array of service endpoints.
            * `locations` (`list`) - A list of locations.
            * `provisioning_state` (`str`) - The provisioning state of the resource.
            * `service` (`str`) - The type of the endpoint service.
    """
    tags: pulumi.Output[dict]
    """
    Resource tags.
    """
    type: pulumi.Output[str]
    """
    Resource type.
    """
    def __init__(__self__, resource_name, opts=None, disable_bgp_route_propagation=None, etag=None, id=None, location=None, name=None, provisioning_state=None, resource_group_name=None, routes=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Route table resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] disable_bgp_route_propagation: Gets or sets whether to disable the routes learned by BGP on that route table. True means disable.
        :param pulumi.Input[str] etag: Gets a unique read-only string that changes whenever the resource is updated.
        :param pulumi.Input[str] id: Resource ID.
        :param pulumi.Input[str] location: Resource location.
        :param pulumi.Input[str] name: The name of the route table.
        :param pulumi.Input[str] provisioning_state: The provisioning state of the resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[list] routes: Collection of routes contained within a route table.
        :param pulumi.Input[dict] tags: Resource tags.

        The **routes** object supports the following:

          * `address_prefix` (`pulumi.Input[str]`) - The destination CIDR to which the route applies.
          * `etag` (`pulumi.Input[str]`) - A unique read-only string that changes whenever the resource is updated.
          * `id` (`pulumi.Input[str]`) - Resource ID.
          * `name` (`pulumi.Input[str]`) - The name of the resource that is unique within a resource group. This name can be used to access the resource.
          * `next_hop_ip_address` (`pulumi.Input[str]`) - The IP address packets should be forwarded to. Next hop values are only allowed in routes where the next hop type is VirtualAppliance.
          * `next_hop_type` (`pulumi.Input[str]`) - The type of Azure hop the packet should be sent to.
          * `provisioning_state` (`pulumi.Input[str]`) - The provisioning state of the resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
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

            __props__['disable_bgp_route_propagation'] = disable_bgp_route_propagation
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
            __props__['routes'] = routes
            __props__['tags'] = tags
            __props__['properties'] = None
            __props__['type'] = None
        super(RouteTable, __self__).__init__(
            'azurerm:network/v20190601:RouteTable',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get an existing RouteTable resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return RouteTable(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop