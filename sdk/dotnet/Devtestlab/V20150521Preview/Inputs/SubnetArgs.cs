// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DevTestLab.V20150521Preview.Inputs
{

    public sealed class SubnetArgs : Pulumi.ResourceArgs
    {
        [Input("allowPublicIp")]
        public Input<string>? AllowPublicIp { get; set; }

        [Input("labSubnetName")]
        public Input<string>? LabSubnetName { get; set; }

        [Input("resourceId")]
        public Input<string>? ResourceId { get; set; }

        public SubnetArgs()
        {
        }
    }
}
