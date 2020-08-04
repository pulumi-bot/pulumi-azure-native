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
    Gets a unique read-only string that changes whenever the resource is updated
    """
    location: pulumi.Output[str]
    """
    Resource location
    """
    name: pulumi.Output[str]
    """
    Resource name
    """
    properties: pulumi.Output[dict]
    """
    Route Table resource
      * `provisioning_state` (`str`) - Gets or sets Provisioning state of the resource Updating/Deleting/Failed
      * `routes` (`list`) - Gets or sets Routes in a Route Table
        * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated
        * `id` (`str`) - Resource Id
        * `name` (`str`) - Gets name of the resource that is unique within a resource group. This name can be used to access the resource
        * `properties` (`dict`) - Route resource
          * `address_prefix` (`str`) - Gets or sets the destination CIDR to which the route applies.
          * `next_hop_ip_address` (`str`) - Gets or sets the IP address packets should be forwarded to. Next hop values are only allowed in routes where the next hop type is VirtualAppliance.
          * `next_hop_type` (`str`) - Gets or sets the type of Azure hop the packet should be sent to.
          * `provisioning_state` (`str`) - Gets or sets Provisioning state of the resource Updating/Deleting/Failed

      * `subnets` (`list`) - Gets collection of references to subnets
        * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated
        * `id` (`str`) - Resource Id
        * `name` (`str`) - Gets name of the resource that is unique within a resource group. This name can be used to access the resource
        * `properties` (`dict`)
          * `address_prefix` (`str`) - Gets or sets Address prefix for the subnet.
          * `ip_configurations` (`list`) - Gets array of references to the network interface IP configurations using subnet
            * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated
            * `id` (`str`) - Resource Id
            * `name` (`str`) - Gets name of the resource that is unique within a resource group. This name can be used to access the resource
            * `properties` (`dict`) - Properties of IPConfiguration
              * `private_ip_address` (`str`) - Gets or sets the privateIPAddress of the IP Configuration
              * `private_ip_allocation_method` (`str`) - Gets or sets PrivateIP allocation method (Static/Dynamic)
              * `provisioning_state` (`str`) - Gets or sets Provisioning state of the PublicIP resource Updating/Deleting/Failed
              * `public_ip_address` (`dict`) - Gets or sets the reference of the PublicIP resource
                * `etag` (`str`) - Gets a unique read-only string that changes whenever the resource is updated
                * `id` (`str`) - Resource Id
                * `location` (`str`) - Resource location
                * `name` (`str`) - Resource name
                * `properties` (`dict`) - PublicIpAddress properties
                  * `dns_settings` (`dict`) - Gets or sets FQDN of the DNS record associated with the public IP address
                    * `domain_name_label` (`str`) - Gets or sets the Domain name label.The concatenation of the domain name label and the regionalized DNS zone make up the fully qualified domain name associated with the public IP address. If a domain name label is specified, an A DNS record is created for the public IP in the Microsoft Azure DNS system.
                    * `fqdn` (`str`) - Gets the FQDN, Fully qualified domain name of the A DNS record associated with the public IP. This is the concatenation of the domainNameLabel and the regionalized DNS zone.
                    * `reverse_fqdn` (`str`) - Gets or Sets the Reverse FQDN. A user-visible, fully qualified domain name that resolves to this public IP address. If the reverseFqdn is specified, then a PTR DNS record is created pointing from the IP address in the in-addr.arpa domain to the reverse FQDN. 

                  * `idle_timeout_in_minutes` (`float`) - Gets or sets the idle timeout of the public IP address
                  * `ip_address` (`str`)
                  * `ip_configuration` (`dict`) - IPConfiguration
                  * `provisioning_state` (`str`) - Gets or sets Provisioning state of the PublicIP resource Updating/Deleting/Failed
                  * `public_ip_address_version` (`str`) - Gets or sets PublicIP address version (IPv4/IPv6)
                  * `public_ip_allocation_method` (`str`) - Gets or sets PublicIP allocation method (Static/Dynamic)
                  * `resource_guid` (`str`) - Gets or sets resource GUID property of the PublicIP resource

                * `tags` (`dict`) - Resource tags
                * `type` (`str`) - Resource type

              * `subnet` (`dict`) - Gets or sets the reference of the subnet resource

          * `network_security_group` (`dict`) - Gets or sets the reference of the NetworkSecurityGroup resource
            * `etag` (`str`) - Gets a unique read-only string that changes whenever the resource is updated
            * `id` (`str`) - Resource Id
            * `location` (`str`) - Resource location
            * `name` (`str`) - Resource name
            * `properties` (`dict`) - Network Security Group resource
              * `default_security_rules` (`list`) - Gets or sets Default security rules of network security group
                * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated
                * `id` (`str`) - Resource Id
                * `name` (`str`) - Gets name of the resource that is unique within a resource group. This name can be used to access the resource
                * `properties` (`dict`)
                  * `access` (`str`) - Gets or sets network traffic is allowed or denied. Possible values are 'Allow' and 'Deny'
                  * `description` (`str`) - Gets or sets a description for this rule. Restricted to 140 chars.
                  * `destination_address_prefix` (`str`) - Gets or sets destination address prefix. CIDR or source IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used. 
                  * `destination_port_range` (`str`) - Gets or sets Destination Port or Range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.
                  * `direction` (`str`) - Gets or sets the direction of the rule.InBound or Outbound. The direction specifies if rule will be evaluated on incoming or outgoing traffic.
                  * `priority` (`float`) - Gets or sets the priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
                  * `protocol` (`str`) - Gets or sets Network protocol this rule applies to. Can be Tcp, Udp or All(*).
                  * `provisioning_state` (`str`) - Gets or sets Provisioning state of the PublicIP resource Updating/Deleting/Failed
                  * `source_address_prefix` (`str`) - Gets or sets source address prefix. CIDR or source IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used. If this is an ingress rule, specifies where network traffic originates from. 
                  * `source_port_range` (`str`) - Gets or sets Source Port or Range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.

              * `network_interfaces` (`list`) - Gets collection of references to Network Interfaces
                * `etag` (`str`) - Gets a unique read-only string that changes whenever the resource is updated
                * `id` (`str`) - Resource Id
                * `location` (`str`) - Resource location
                * `name` (`str`) - Resource name
                * `properties` (`dict`) - NetworkInterface properties. 
                  * `dns_settings` (`dict`) - Gets or sets DNS Settings in  NetworkInterface
                    * `applied_dns_servers` (`list`) - Gets or sets list of Applied DNS servers IP addresses
                    * `dns_servers` (`list`) - Gets or sets list of DNS servers IP addresses
                    * `internal_dns_name_label` (`str`) - Gets or sets the Internal DNS name
                    * `internal_domain_name_suffix` (`str`) - Gets or sets internal domain name suffix of the NIC.
                    * `internal_fqdn` (`str`) - Gets or sets the internal FQDN.

                  * `enable_ip_forwarding` (`bool`) - Gets or sets whether IPForwarding is enabled on the NIC
                  * `ip_configurations` (`list`) - Gets or sets list of IPConfigurations of the NetworkInterface
                    * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated
                    * `id` (`str`) - Resource Id
                    * `name` (`str`) - Gets name of the resource that is unique within a resource group. This name can be used to access the resource
                    * `properties` (`dict`) - Properties of IPConfiguration
                      * `application_gateway_backend_address_pools` (`list`) - Gets or sets the reference of ApplicationGatewayBackendAddressPool resource
                        * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated
                        * `id` (`str`) - Resource Id
                        * `name` (`str`) - Gets name of the resource that is unique within a resource group. This name can be used to access the resource
                        * `properties` (`dict`) - Properties of Backend Address Pool of application gateway
                          * `backend_addresses` (`list`) - Gets or sets the backend addresses
                            * `fqdn` (`str`) - Gets or sets the dns name
                            * `ip_address` (`str`) - Gets or sets the ip address

                          * `backend_ip_configurations` (`list`) - Gets collection of references to IPs defined in NICs
                          * `provisioning_state` (`str`) - Gets or sets Provisioning state of the backend address pool resource Updating/Deleting/Failed

                      * `load_balancer_backend_address_pools` (`list`) - Gets or sets the reference of LoadBalancerBackendAddressPool resource
                        * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated
                        * `id` (`str`) - Resource Id
                        * `name` (`str`) - Gets name of the resource that is unique within a resource group. This name can be used to access the resource
                        * `properties` (`dict`) - Properties of BackendAddressPool
                          * `backend_ip_configurations` (`list`) - Gets collection of references to IPs defined in NICs
                          * `load_balancing_rules` (`list`) - Gets Load Balancing rules that use this Backend Address Pool
                            * `id` (`str`) - Resource Id

                          * `outbound_nat_rule` (`dict`) - Gets outbound rules that use this Backend Address Pool
                          * `provisioning_state` (`str`) - Provisioning state of the PublicIP resource Updating/Deleting/Failed

                      * `load_balancer_inbound_nat_rules` (`list`) - Gets or sets list of references of LoadBalancerInboundNatRules
                        * `etag` (`str`) - A unique read-only string that changes whenever the resource is updated
                        * `id` (`str`) - Resource Id
                        * `name` (`str`) - Gets name of the resource that is unique within a resource group. This name can be used to access the resource
                        * `properties` (`dict`) - Properties of Inbound NAT rule
                          * `backend_ip_configuration` (`dict`) - Gets or sets a reference to a private ip address defined on a NetworkInterface of a VM. Traffic sent to frontendPort of each of the frontendIPConfigurations is forwarded to the backed IP
                          * `backend_port` (`float`) - Gets or sets a port used for internal connections on the endpoint. The localPort attribute maps the eternal port of the endpoint to an internal port on a role. This is useful in scenarios where a role must communicate to an internal component on a port that is different from the one that is exposed externally. If not specified, the value of localPort is the same as the port attribute. Set the value of localPort to '*' to automatically assign an unallocated port that is discoverable using the runtime API
                          * `enable_floating_ip` (`bool`) - Configures a virtual machine's endpoint for the floating IP capability required to configure a SQL AlwaysOn availability Group. This setting is required when using the SQL Always ON availability Groups in SQL server. This setting can't be changed after you create the endpoint
                          * `frontend_ip_configuration` (`dict`) - Gets or sets a reference to frontend IP Addresses
                          * `frontend_port` (`float`) - Gets or sets the port for the external endpoint. You can specify any port number you choose, but the port numbers specified for each role in the service must be unique. Possible values range between 1 and 65535, inclusive
                          * `idle_timeout_in_minutes` (`float`) - Gets or sets the timeout for the Tcp idle connection. The value can be set between 4 and 30 minutes. The default value is 4 minutes. This element is only used when the protocol is set to Tcp
                          * `protocol` (`str`) - Gets or sets the transport protocol for the external endpoint. Possible values are Udp or Tcp
                          * `provisioning_state` (`str`) - Gets or sets Provisioning state of the PublicIP resource Updating/Deleting/Failed

                      * `primary` (`bool`) - Gets whether this is a primary customer address on the NIC
                      * `private_ip_address` (`str`)
                      * `private_ip_address_version` (`str`) - Gets or sets PrivateIP address version (IPv4/IPv6)
                      * `private_ip_allocation_method` (`str`) - Gets or sets PrivateIP allocation method (Static/Dynamic)
                      * `provisioning_state` (`str`)
                      * `public_ip_address` (`dict`) - PublicIPAddress resource
                      * `subnet` (`dict`) - Subnet in a VirtualNetwork resource

                  * `mac_address` (`str`) - Gets the MAC Address of the network interface
                  * `network_security_group` (`dict`) - Gets or sets the reference of the NetworkSecurityGroup resource
                  * `primary` (`bool`) - Gets whether this is a primary NIC on a virtual machine
                  * `provisioning_state` (`str`) - Gets or sets Provisioning state of the PublicIP resource Updating/Deleting/Failed
                  * `resource_guid` (`str`) - Gets or sets resource GUID property of the network interface resource
                  * `virtual_machine` (`dict`) - Gets or sets the reference of a VirtualMachine

                * `tags` (`dict`) - Resource tags
                * `type` (`str`) - Resource type

              * `provisioning_state` (`str`) - Gets or sets Provisioning state of the PublicIP resource Updating/Deleting/Failed
              * `resource_guid` (`str`) - Gets or sets resource GUID property of the network security group resource
              * `security_rules` (`list`) - Gets or sets Security rules of network security group
              * `subnets` (`list`) - Gets collection of references to subnets

            * `tags` (`dict`) - Resource tags
            * `type` (`str`) - Resource type

          * `provisioning_state` (`str`) - Gets or sets Provisioning state of the PublicIP resource Updating/Deleting/Failed
          * `route_table` (`dict`) - Gets or sets the reference of the RouteTable resource
            * `etag` (`str`) - Gets a unique read-only string that changes whenever the resource is updated
            * `id` (`str`) - Resource Id
            * `location` (`str`) - Resource location
            * `name` (`str`) - Resource name
            * `properties` (`dict`) - Route Table resource
            * `tags` (`dict`) - Resource tags
            * `type` (`str`) - Resource type
    """
    tags: pulumi.Output[dict]
    """
    Resource tags
    """
    type: pulumi.Output[str]
    """
    Resource type
    """
    def __init__(__self__, resource_name, opts=None, etag=None, id=None, location=None, name=None, properties=None, resource_group_name=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        RouteTable resource

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] etag: Gets a unique read-only string that changes whenever the resource is updated
        :param pulumi.Input[str] id: Resource Id
        :param pulumi.Input[str] location: Resource location
        :param pulumi.Input[str] name: The name of the route table.
        :param pulumi.Input[dict] properties: Route Table resource
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[dict] tags: Resource tags

        The **properties** object supports the following:

          * `provisioning_state` (`pulumi.Input[str]`) - Gets or sets Provisioning state of the resource Updating/Deleting/Failed
          * `routes` (`pulumi.Input[list]`) - Gets or sets Routes in a Route Table
            * `etag` (`pulumi.Input[str]`) - A unique read-only string that changes whenever the resource is updated
            * `id` (`pulumi.Input[str]`) - Resource Id
            * `name` (`pulumi.Input[str]`) - Gets name of the resource that is unique within a resource group. This name can be used to access the resource
            * `properties` (`pulumi.Input[dict]`) - Route resource
              * `address_prefix` (`pulumi.Input[str]`) - Gets or sets the destination CIDR to which the route applies.
              * `next_hop_ip_address` (`pulumi.Input[str]`) - Gets or sets the IP address packets should be forwarded to. Next hop values are only allowed in routes where the next hop type is VirtualAppliance.
              * `next_hop_type` (`pulumi.Input[str]`) - Gets or sets the type of Azure hop the packet should be sent to.
              * `provisioning_state` (`pulumi.Input[str]`) - Gets or sets Provisioning state of the resource Updating/Deleting/Failed

          * `subnets` (`pulumi.Input[list]`) - Gets collection of references to subnets
            * `etag` (`pulumi.Input[str]`) - A unique read-only string that changes whenever the resource is updated
            * `id` (`pulumi.Input[str]`) - Resource Id
            * `name` (`pulumi.Input[str]`) - Gets name of the resource that is unique within a resource group. This name can be used to access the resource
            * `properties` (`pulumi.Input[dict]`)
              * `address_prefix` (`pulumi.Input[str]`) - Gets or sets Address prefix for the subnet.
              * `ip_configurations` (`pulumi.Input[list]`) - Gets array of references to the network interface IP configurations using subnet
                * `etag` (`pulumi.Input[str]`) - A unique read-only string that changes whenever the resource is updated
                * `id` (`pulumi.Input[str]`) - Resource Id
                * `name` (`pulumi.Input[str]`) - Gets name of the resource that is unique within a resource group. This name can be used to access the resource
                * `properties` (`pulumi.Input[dict]`) - Properties of IPConfiguration
                  * `private_ip_address` (`pulumi.Input[str]`) - Gets or sets the privateIPAddress of the IP Configuration
                  * `private_ip_allocation_method` (`pulumi.Input[str]`) - Gets or sets PrivateIP allocation method (Static/Dynamic)
                  * `provisioning_state` (`pulumi.Input[str]`) - Gets or sets Provisioning state of the PublicIP resource Updating/Deleting/Failed
                  * `public_ip_address` (`pulumi.Input[dict]`) - Gets or sets the reference of the PublicIP resource
                    * `etag` (`pulumi.Input[str]`) - Gets a unique read-only string that changes whenever the resource is updated
                    * `id` (`pulumi.Input[str]`) - Resource Id
                    * `location` (`pulumi.Input[str]`) - Resource location
                    * `properties` (`pulumi.Input[dict]`) - PublicIpAddress properties
                      * `dns_settings` (`pulumi.Input[dict]`) - Gets or sets FQDN of the DNS record associated with the public IP address
                        * `domain_name_label` (`pulumi.Input[str]`) - Gets or sets the Domain name label.The concatenation of the domain name label and the regionalized DNS zone make up the fully qualified domain name associated with the public IP address. If a domain name label is specified, an A DNS record is created for the public IP in the Microsoft Azure DNS system.
                        * `fqdn` (`pulumi.Input[str]`) - Gets the FQDN, Fully qualified domain name of the A DNS record associated with the public IP. This is the concatenation of the domainNameLabel and the regionalized DNS zone.
                        * `reverse_fqdn` (`pulumi.Input[str]`) - Gets or Sets the Reverse FQDN. A user-visible, fully qualified domain name that resolves to this public IP address. If the reverseFqdn is specified, then a PTR DNS record is created pointing from the IP address in the in-addr.arpa domain to the reverse FQDN. 

                      * `idle_timeout_in_minutes` (`pulumi.Input[float]`) - Gets or sets the idle timeout of the public IP address
                      * `ip_address` (`pulumi.Input[str]`)
                      * `ip_configuration` (`pulumi.Input[dict]`) - IPConfiguration
                      * `provisioning_state` (`pulumi.Input[str]`) - Gets or sets Provisioning state of the PublicIP resource Updating/Deleting/Failed
                      * `public_ip_address_version` (`pulumi.Input[str]`) - Gets or sets PublicIP address version (IPv4/IPv6)
                      * `public_ip_allocation_method` (`pulumi.Input[str]`) - Gets or sets PublicIP allocation method (Static/Dynamic)
                      * `resource_guid` (`pulumi.Input[str]`) - Gets or sets resource GUID property of the PublicIP resource

                    * `tags` (`pulumi.Input[dict]`) - Resource tags

                  * `subnet` (`pulumi.Input[dict]`) - Gets or sets the reference of the subnet resource

              * `network_security_group` (`pulumi.Input[dict]`) - Gets or sets the reference of the NetworkSecurityGroup resource
                * `etag` (`pulumi.Input[str]`) - Gets a unique read-only string that changes whenever the resource is updated
                * `id` (`pulumi.Input[str]`) - Resource Id
                * `location` (`pulumi.Input[str]`) - Resource location
                * `properties` (`pulumi.Input[dict]`) - Network Security Group resource
                  * `default_security_rules` (`pulumi.Input[list]`) - Gets or sets Default security rules of network security group
                    * `etag` (`pulumi.Input[str]`) - A unique read-only string that changes whenever the resource is updated
                    * `id` (`pulumi.Input[str]`) - Resource Id
                    * `name` (`pulumi.Input[str]`) - Gets name of the resource that is unique within a resource group. This name can be used to access the resource
                    * `properties` (`pulumi.Input[dict]`)
                      * `access` (`pulumi.Input[str]`) - Gets or sets network traffic is allowed or denied. Possible values are 'Allow' and 'Deny'
                      * `description` (`pulumi.Input[str]`) - Gets or sets a description for this rule. Restricted to 140 chars.
                      * `destination_address_prefix` (`pulumi.Input[str]`) - Gets or sets destination address prefix. CIDR or source IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used. 
                      * `destination_port_range` (`pulumi.Input[str]`) - Gets or sets Destination Port or Range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.
                      * `direction` (`pulumi.Input[str]`) - Gets or sets the direction of the rule.InBound or Outbound. The direction specifies if rule will be evaluated on incoming or outgoing traffic.
                      * `priority` (`pulumi.Input[float]`) - Gets or sets the priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each rule in the collection. The lower the priority number, the higher the priority of the rule.
                      * `protocol` (`pulumi.Input[str]`) - Gets or sets Network protocol this rule applies to. Can be Tcp, Udp or All(*).
                      * `provisioning_state` (`pulumi.Input[str]`) - Gets or sets Provisioning state of the PublicIP resource Updating/Deleting/Failed
                      * `source_address_prefix` (`pulumi.Input[str]`) - Gets or sets source address prefix. CIDR or source IP range. Asterisk '*' can also be used to match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used. If this is an ingress rule, specifies where network traffic originates from. 
                      * `source_port_range` (`pulumi.Input[str]`) - Gets or sets Source Port or Range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match all ports.

                  * `network_interfaces` (`pulumi.Input[list]`) - Gets collection of references to Network Interfaces
                    * `etag` (`pulumi.Input[str]`) - Gets a unique read-only string that changes whenever the resource is updated
                    * `id` (`pulumi.Input[str]`) - Resource Id
                    * `location` (`pulumi.Input[str]`) - Resource location
                    * `properties` (`pulumi.Input[dict]`) - NetworkInterface properties. 
                      * `dns_settings` (`pulumi.Input[dict]`) - Gets or sets DNS Settings in  NetworkInterface
                        * `applied_dns_servers` (`pulumi.Input[list]`) - Gets or sets list of Applied DNS servers IP addresses
                        * `dns_servers` (`pulumi.Input[list]`) - Gets or sets list of DNS servers IP addresses
                        * `internal_dns_name_label` (`pulumi.Input[str]`) - Gets or sets the Internal DNS name
                        * `internal_domain_name_suffix` (`pulumi.Input[str]`) - Gets or sets internal domain name suffix of the NIC.
                        * `internal_fqdn` (`pulumi.Input[str]`) - Gets or sets the internal FQDN.

                      * `enable_ip_forwarding` (`pulumi.Input[bool]`) - Gets or sets whether IPForwarding is enabled on the NIC
                      * `ip_configurations` (`pulumi.Input[list]`) - Gets or sets list of IPConfigurations of the NetworkInterface
                        * `etag` (`pulumi.Input[str]`) - A unique read-only string that changes whenever the resource is updated
                        * `id` (`pulumi.Input[str]`) - Resource Id
                        * `name` (`pulumi.Input[str]`) - Gets name of the resource that is unique within a resource group. This name can be used to access the resource
                        * `properties` (`pulumi.Input[dict]`) - Properties of IPConfiguration
                          * `application_gateway_backend_address_pools` (`pulumi.Input[list]`) - Gets or sets the reference of ApplicationGatewayBackendAddressPool resource
                            * `etag` (`pulumi.Input[str]`) - A unique read-only string that changes whenever the resource is updated
                            * `id` (`pulumi.Input[str]`) - Resource Id
                            * `name` (`pulumi.Input[str]`) - Gets name of the resource that is unique within a resource group. This name can be used to access the resource
                            * `properties` (`pulumi.Input[dict]`) - Properties of Backend Address Pool of application gateway
                              * `backend_addresses` (`pulumi.Input[list]`) - Gets or sets the backend addresses
                                * `fqdn` (`pulumi.Input[str]`) - Gets or sets the dns name
                                * `ip_address` (`pulumi.Input[str]`) - Gets or sets the ip address

                              * `backend_ip_configurations` (`pulumi.Input[list]`) - Gets collection of references to IPs defined in NICs
                              * `provisioning_state` (`pulumi.Input[str]`) - Gets or sets Provisioning state of the backend address pool resource Updating/Deleting/Failed

                          * `load_balancer_backend_address_pools` (`pulumi.Input[list]`) - Gets or sets the reference of LoadBalancerBackendAddressPool resource
                            * `etag` (`pulumi.Input[str]`) - A unique read-only string that changes whenever the resource is updated
                            * `id` (`pulumi.Input[str]`) - Resource Id
                            * `name` (`pulumi.Input[str]`) - Gets name of the resource that is unique within a resource group. This name can be used to access the resource
                            * `properties` (`pulumi.Input[dict]`) - Properties of BackendAddressPool
                              * `backend_ip_configurations` (`pulumi.Input[list]`) - Gets collection of references to IPs defined in NICs
                              * `load_balancing_rules` (`pulumi.Input[list]`) - Gets Load Balancing rules that use this Backend Address Pool
                                * `id` (`pulumi.Input[str]`) - Resource Id

                              * `outbound_nat_rule` (`pulumi.Input[dict]`) - Gets outbound rules that use this Backend Address Pool
                              * `provisioning_state` (`pulumi.Input[str]`) - Provisioning state of the PublicIP resource Updating/Deleting/Failed

                          * `load_balancer_inbound_nat_rules` (`pulumi.Input[list]`) - Gets or sets list of references of LoadBalancerInboundNatRules
                            * `etag` (`pulumi.Input[str]`) - A unique read-only string that changes whenever the resource is updated
                            * `id` (`pulumi.Input[str]`) - Resource Id
                            * `name` (`pulumi.Input[str]`) - Gets name of the resource that is unique within a resource group. This name can be used to access the resource
                            * `properties` (`pulumi.Input[dict]`) - Properties of Inbound NAT rule
                              * `backend_ip_configuration` (`pulumi.Input[dict]`) - Gets or sets a reference to a private ip address defined on a NetworkInterface of a VM. Traffic sent to frontendPort of each of the frontendIPConfigurations is forwarded to the backed IP
                              * `backend_port` (`pulumi.Input[float]`) - Gets or sets a port used for internal connections on the endpoint. The localPort attribute maps the eternal port of the endpoint to an internal port on a role. This is useful in scenarios where a role must communicate to an internal component on a port that is different from the one that is exposed externally. If not specified, the value of localPort is the same as the port attribute. Set the value of localPort to '*' to automatically assign an unallocated port that is discoverable using the runtime API
                              * `enable_floating_ip` (`pulumi.Input[bool]`) - Configures a virtual machine's endpoint for the floating IP capability required to configure a SQL AlwaysOn availability Group. This setting is required when using the SQL Always ON availability Groups in SQL server. This setting can't be changed after you create the endpoint
                              * `frontend_ip_configuration` (`pulumi.Input[dict]`) - Gets or sets a reference to frontend IP Addresses
                              * `frontend_port` (`pulumi.Input[float]`) - Gets or sets the port for the external endpoint. You can specify any port number you choose, but the port numbers specified for each role in the service must be unique. Possible values range between 1 and 65535, inclusive
                              * `idle_timeout_in_minutes` (`pulumi.Input[float]`) - Gets or sets the timeout for the Tcp idle connection. The value can be set between 4 and 30 minutes. The default value is 4 minutes. This element is only used when the protocol is set to Tcp
                              * `protocol` (`pulumi.Input[str]`) - Gets or sets the transport protocol for the external endpoint. Possible values are Udp or Tcp
                              * `provisioning_state` (`pulumi.Input[str]`) - Gets or sets Provisioning state of the PublicIP resource Updating/Deleting/Failed

                          * `primary` (`pulumi.Input[bool]`) - Gets whether this is a primary customer address on the NIC
                          * `private_ip_address` (`pulumi.Input[str]`)
                          * `private_ip_address_version` (`pulumi.Input[str]`) - Gets or sets PrivateIP address version (IPv4/IPv6)
                          * `private_ip_allocation_method` (`pulumi.Input[str]`) - Gets or sets PrivateIP allocation method (Static/Dynamic)
                          * `provisioning_state` (`pulumi.Input[str]`)
                          * `public_ip_address` (`pulumi.Input[dict]`) - PublicIPAddress resource
                          * `subnet` (`pulumi.Input[dict]`) - Subnet in a VirtualNetwork resource

                      * `mac_address` (`pulumi.Input[str]`) - Gets the MAC Address of the network interface
                      * `network_security_group` (`pulumi.Input[dict]`) - Gets or sets the reference of the NetworkSecurityGroup resource
                      * `primary` (`pulumi.Input[bool]`) - Gets whether this is a primary NIC on a virtual machine
                      * `provisioning_state` (`pulumi.Input[str]`) - Gets or sets Provisioning state of the PublicIP resource Updating/Deleting/Failed
                      * `resource_guid` (`pulumi.Input[str]`) - Gets or sets resource GUID property of the network interface resource
                      * `virtual_machine` (`pulumi.Input[dict]`) - Gets or sets the reference of a VirtualMachine

                    * `tags` (`pulumi.Input[dict]`) - Resource tags

                  * `provisioning_state` (`pulumi.Input[str]`) - Gets or sets Provisioning state of the PublicIP resource Updating/Deleting/Failed
                  * `resource_guid` (`pulumi.Input[str]`) - Gets or sets resource GUID property of the network security group resource
                  * `security_rules` (`pulumi.Input[list]`) - Gets or sets Security rules of network security group
                  * `subnets` (`pulumi.Input[list]`) - Gets collection of references to subnets

                * `tags` (`pulumi.Input[dict]`) - Resource tags

              * `provisioning_state` (`pulumi.Input[str]`) - Gets or sets Provisioning state of the PublicIP resource Updating/Deleting/Failed
              * `route_table` (`pulumi.Input[dict]`) - Gets or sets the reference of the RouteTable resource
                * `etag` (`pulumi.Input[str]`) - Gets a unique read-only string that changes whenever the resource is updated
                * `id` (`pulumi.Input[str]`) - Resource Id
                * `location` (`pulumi.Input[str]`) - Resource location
                * `properties` (`pulumi.Input[dict]`) - Route Table resource
                * `tags` (`pulumi.Input[dict]`) - Resource tags
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

            __props__['etag'] = etag
            __props__['id'] = id
            __props__['location'] = location
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            __props__['properties'] = properties
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['type'] = None
        super(RouteTable, __self__).__init__(
            'azurerm:network/v20160330:RouteTable',
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
