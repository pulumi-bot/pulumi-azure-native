// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ServiceFabric.V20191101Preview.Outputs
{

    [OutputType]
    public sealed class ServicePlacementPolicyDescriptionResponseResult
    {
        /// <summary>
        /// The type of placement policy for a service fabric service. Following are the possible values.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private ServicePlacementPolicyDescriptionResponseResult(string type)
        {
            Type = type;
        }
    }
}
