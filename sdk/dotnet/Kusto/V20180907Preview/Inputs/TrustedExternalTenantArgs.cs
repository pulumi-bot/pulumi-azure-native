// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Kusto.V20180907Preview.Inputs
{

    public sealed class TrustedExternalTenantArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// GUID representing an external tenant.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public TrustedExternalTenantArgs()
        {
        }
    }
}
