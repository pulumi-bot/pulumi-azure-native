// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.MachineLearningCompute.V20170801Preview.Inputs
{

    /// <summary>
    /// Global service auth configuration properties. These are the data-plane authorization keys and are used if a service doesn't define it's own.
    /// </summary>
    public sealed class ServiceAuthConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The primary auth key hash. This is not returned in response of GET/PUT on the resource.. To see this please call listKeys API.
        /// </summary>
        [Input("primaryAuthKeyHash", required: true)]
        public Input<string> PrimaryAuthKeyHash { get; set; } = null!;

        /// <summary>
        /// The secondary auth key hash. This is not returned in response of GET/PUT on the resource.. To see this please call listKeys API.
        /// </summary>
        [Input("secondaryAuthKeyHash", required: true)]
        public Input<string> SecondaryAuthKeyHash { get; set; } = null!;

        public ServiceAuthConfigurationArgs()
        {
        }
    }
}
