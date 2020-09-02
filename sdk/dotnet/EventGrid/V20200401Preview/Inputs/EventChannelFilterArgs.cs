// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.EventGrid.V20200401Preview.Inputs
{

    /// <summary>
    /// Filter for the Event Channel.
    /// </summary>
    public sealed class EventChannelFilterArgs : Pulumi.ResourceArgs
    {
        [Input("advancedFilters")]
        private InputList<Inputs.AdvancedFilterArgs>? _advancedFilters;

        /// <summary>
        /// An array of advanced filters that are used for filtering event channels.
        /// </summary>
        public InputList<Inputs.AdvancedFilterArgs> AdvancedFilters
        {
            get => _advancedFilters ?? (_advancedFilters = new InputList<Inputs.AdvancedFilterArgs>());
            set => _advancedFilters = value;
        }

        public EventChannelFilterArgs()
        {
        }
    }
}
