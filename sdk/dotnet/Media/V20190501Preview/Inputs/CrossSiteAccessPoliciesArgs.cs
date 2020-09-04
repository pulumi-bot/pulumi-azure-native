// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Media.V20190501Preview.Inputs
{

    /// <summary>
    /// The client access policy.
    /// </summary>
    public sealed class CrossSiteAccessPoliciesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The content of clientaccesspolicy.xml used by Silverlight.
        /// </summary>
        [Input("clientAccessPolicy")]
        public Input<string>? ClientAccessPolicy { get; set; }

        /// <summary>
        /// The content of crossdomain.xml used by Silverlight.
        /// </summary>
        [Input("crossDomainPolicy")]
        public Input<string>? CrossDomainPolicy { get; set; }

        public CrossSiteAccessPoliciesArgs()
        {
        }
    }
}
