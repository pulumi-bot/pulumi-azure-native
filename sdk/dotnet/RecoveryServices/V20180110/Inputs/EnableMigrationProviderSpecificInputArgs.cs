// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.RecoveryServices.V20180110.Inputs
{

    /// <summary>
    /// Enable migration provider specific input.
    /// </summary>
    public sealed class EnableMigrationProviderSpecificInputArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The class type.
        /// </summary>
        [Input("instanceType", required: true)]
        public Input<string> InstanceType { get; set; } = null!;

        public EnableMigrationProviderSpecificInputArgs()
        {
        }
    }
}