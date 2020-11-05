// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Insights.V20201005Preview.Inputs
{

    /// <summary>
    /// An XML configuration specification for a WebTest.
    /// </summary>
    public sealed class WebTestPropertiesConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The XML specification of a WebTest to run against an application.
        /// </summary>
        [Input("webTest")]
        public Input<string>? WebTest { get; set; }

        public WebTestPropertiesConfigurationArgs()
        {
        }
    }
}
