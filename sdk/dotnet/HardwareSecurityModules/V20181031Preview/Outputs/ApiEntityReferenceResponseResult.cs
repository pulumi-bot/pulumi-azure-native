// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.HardwareSecurityModules.V20181031Preview.Outputs
{

    [OutputType]
    public sealed class ApiEntityReferenceResponseResult
    {
        /// <summary>
        /// The ARM resource id in the form of /subscriptions/{SubscriptionId}/resourceGroups/{ResourceGroupName}/...
        /// </summary>
        public readonly string? Id;

        [OutputConstructor]
        private ApiEntityReferenceResponseResult(string? id)
        {
            Id = id;
        }
    }
}
