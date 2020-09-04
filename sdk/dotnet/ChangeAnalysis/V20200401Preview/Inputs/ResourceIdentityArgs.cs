// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ChangeAnalysis.V20200401Preview.Inputs
{

    /// <summary>
    /// The identity block returned by ARM resource that supports managed identity.
    /// </summary>
    public sealed class ResourceIdentityArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of managed identity used. The type 'SystemAssigned, UserAssigned' includes both an implicitly created identity and a set of user-assigned identities. The type 'None' will remove any identities.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public ResourceIdentityArgs()
        {
        }
    }
}
