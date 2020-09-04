// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Media.V20180601Preview.Outputs
{

    [OutputType]
    public sealed class CrossSiteAccessPoliciesResponseResult
    {
        /// <summary>
        /// The content of clientaccesspolicy.xml used by Silverlight.
        /// </summary>
        public readonly string? ClientAccessPolicy;
        /// <summary>
        /// The content of crossdomain.xml used by Silverlight.
        /// </summary>
        public readonly string? CrossDomainPolicy;

        [OutputConstructor]
        private CrossSiteAccessPoliciesResponseResult(
            string? clientAccessPolicy,

            string? crossDomainPolicy)
        {
            ClientAccessPolicy = clientAccessPolicy;
            CrossDomainPolicy = crossDomainPolicy;
        }
    }
}
