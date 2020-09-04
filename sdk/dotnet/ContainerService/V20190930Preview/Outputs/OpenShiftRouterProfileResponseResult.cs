// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ContainerService.V20190930Preview.Outputs
{

    [OutputType]
    public sealed class OpenShiftRouterProfileResponseResult
    {
        /// <summary>
        /// Auto-allocated FQDN for the OpenShift router.
        /// </summary>
        public readonly string Fqdn;
        /// <summary>
        /// Name of the router profile.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// DNS subdomain for OpenShift router.
        /// </summary>
        public readonly string PublicSubdomain;

        [OutputConstructor]
        private OpenShiftRouterProfileResponseResult(
            string fqdn,

            string? name,

            string publicSubdomain)
        {
            Fqdn = fqdn;
            Name = name;
            PublicSubdomain = publicSubdomain;
        }
    }
}
