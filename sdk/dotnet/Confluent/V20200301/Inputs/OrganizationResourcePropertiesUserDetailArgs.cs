// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Confluent.V20200301.Inputs
{

    /// <summary>
    /// Subscriber detail
    /// </summary>
    public sealed class OrganizationResourcePropertiesUserDetailArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Email address
        /// </summary>
        [Input("emailAddress")]
        public Input<string>? EmailAddress { get; set; }

        /// <summary>
        /// First name
        /// </summary>
        [Input("firstName")]
        public Input<string>? FirstName { get; set; }

        /// <summary>
        /// Last name
        /// </summary>
        [Input("lastName")]
        public Input<string>? LastName { get; set; }

        public OrganizationResourcePropertiesUserDetailArgs()
        {
        }
    }
}
