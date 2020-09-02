// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DocumentDB.V20200601Preview.Inputs
{

    public sealed class ApiPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Describes the ServerVersion of an a MongoDB account.
        /// </summary>
        [Input("serverVersion")]
        public Input<string>? ServerVersion { get; set; }

        public ApiPropertiesArgs()
        {
        }
    }
}
