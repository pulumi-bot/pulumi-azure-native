// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DocumentDB.V20200601Preview.Outputs
{

    [OutputType]
    public sealed class ApiPropertiesResponseResult
    {
        /// <summary>
        /// Describes the ServerVersion of an a MongoDB account.
        /// </summary>
        public readonly string? ServerVersion;

        [OutputConstructor]
        private ApiPropertiesResponseResult(string? serverVersion)
        {
            ServerVersion = serverVersion;
        }
    }
}
