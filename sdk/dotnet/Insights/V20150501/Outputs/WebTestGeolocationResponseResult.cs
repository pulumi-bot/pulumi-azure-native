// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Insights.V20150501.Outputs
{

    [OutputType]
    public sealed class WebTestGeolocationResponseResult
    {
        /// <summary>
        /// Location ID for the webtest to run from.
        /// </summary>
        public readonly string? Id;

        [OutputConstructor]
        private WebTestGeolocationResponseResult(string? Id)
        {
            this.Id = Id;
        }
    }
}