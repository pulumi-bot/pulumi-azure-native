// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ContainerService.V20151101Preview.Outputs
{

    [OutputType]
    public sealed class ContainerServiceAgentPoolProfileResponseResult
    {
        /// <summary>
        /// No. of agents (VMs) that will host docker containers
        /// </summary>
        public readonly int? Count;
        /// <summary>
        /// DNS prefix to be used to create FQDN for this agent pool
        /// </summary>
        public readonly string DnsPrefix;
        /// <summary>
        /// FQDN for the agent pool
        /// </summary>
        public readonly string Fqdn;
        /// <summary>
        /// Unique name of the agent pool profile within the context of the subscription and resource group
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Size of agent VMs
        /// </summary>
        public readonly string? VmSize;

        [OutputConstructor]
        private ContainerServiceAgentPoolProfileResponseResult(
            int? count,

            string dnsPrefix,

            string fqdn,

            string name,

            string? vmSize)
        {
            Count = count;
            DnsPrefix = dnsPrefix;
            Fqdn = fqdn;
            Name = name;
            VmSize = vmSize;
        }
    }
}
