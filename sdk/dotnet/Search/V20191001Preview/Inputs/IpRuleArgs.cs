// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Search.V20191001Preview.Inputs
{

    /// <summary>
    /// The IP restriction rule of the Azure Cognitive Search service.
    /// </summary>
    public sealed class IpRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Value corresponding to a single IPv4 address (eg., 123.1.2.3) or an IP range in CIDR format (eg., 123.1.2.3/24) to be allowed.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public IpRuleArgs()
        {
        }
    }
}
