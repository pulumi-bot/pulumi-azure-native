// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.MachineLearningServices.V20200901Preview.Inputs
{

    public sealed class RCranPackageArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The package name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The repository name.
        /// </summary>
        [Input("repository")]
        public Input<string>? Repository { get; set; }

        public RCranPackageArgs()
        {
        }
    }
}
