// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Logic.V20150801Preview.Outputs
{

    [OutputType]
    public sealed class B2BPartnerContentResponseResult
    {
        /// <summary>
        /// The list of partner business identities.
        /// </summary>
        public readonly ImmutableArray<Outputs.BusinessIdentityResponseResult> BusinessIdentities;

        [OutputConstructor]
        private B2BPartnerContentResponseResult(ImmutableArray<Outputs.BusinessIdentityResponseResult> businessIdentities)
        {
            BusinessIdentities = businessIdentities;
        }
    }
}
