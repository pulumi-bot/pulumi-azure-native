// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Compute.V20160430Preview.Inputs
{

    /// <summary>
    /// The API entity reference.
    /// </summary>
    public sealed class ApiEntityReferenceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARM resource id in the form of /subscriptions/{SubscriptionId}/resourceGroups/{ResourceGroupName}/...
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        public ApiEntityReferenceArgs()
        {
        }
    }
}
