// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.StorageCache.Inputs
{

    /// <summary>
    /// Cache security settings.
    /// </summary>
    public sealed class CacheSecuritySettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// root squash of cache property.
        /// </summary>
        [Input("rootSquash")]
        public Input<bool>? RootSquash { get; set; }

        public CacheSecuritySettingsArgs()
        {
        }
    }
}