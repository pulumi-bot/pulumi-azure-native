// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DevTestLab.V20160515.Inputs
{

    /// <summary>
    /// Schedules applicable to a virtual machine. The schedules may have been defined on a VM or on lab level.
    /// </summary>
    public sealed class ApplicableScheduleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The auto-shutdown schedule, if one has been set at the lab or lab resource level.
        /// </summary>
        [Input("labVmsShutdown")]
        public Input<Inputs.ScheduleArgs>? LabVmsShutdown { get; set; }

        /// <summary>
        /// The auto-startup schedule, if one has been set at the lab or lab resource level.
        /// </summary>
        [Input("labVmsStartup")]
        public Input<Inputs.ScheduleArgs>? LabVmsStartup { get; set; }

        /// <summary>
        /// The location of the resource.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// The tags of the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ApplicableScheduleArgs()
        {
        }
    }
}
