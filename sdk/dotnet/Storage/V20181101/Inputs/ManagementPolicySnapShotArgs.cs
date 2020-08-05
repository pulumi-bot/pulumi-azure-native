// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Storage.V20181101.Inputs
{

    /// <summary>
    /// Management policy action for snapshot.
    /// </summary>
    public sealed class ManagementPolicySnapShotArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The function to delete the blob snapshot
        /// </summary>
        [Input("delete")]
        public Input<Inputs.DateAfterCreationArgs>? Delete { get; set; }

        public ManagementPolicySnapShotArgs()
        {
        }
    }
}
