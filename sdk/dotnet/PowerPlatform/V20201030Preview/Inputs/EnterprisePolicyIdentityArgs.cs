// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.PowerPlatform.V20201030Preview.Inputs
{

    /// <summary>
    /// The identity of the EnterprisePolicy.
    /// </summary>
    public sealed class EnterprisePolicyIdentityArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of identity used for the EnterprisePolicy. Currently, the only supported type is 'SystemAssigned', which implicitly creates an identity.
        /// </summary>
        [Input("type")]
        public Input<Pulumi.AzureNextGen.PowerPlatform.V20201030Preview.ResourceIdentityType>? Type { get; set; }

        public EnterprisePolicyIdentityArgs()
        {
        }
    }
}
