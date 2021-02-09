// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Assessment sizing criterion.
type AssessmentSizingCriterion pulumi.String

const (
	AssessmentSizingCriterionPerformanceBased = AssessmentSizingCriterion("PerformanceBased")
	AssessmentSizingCriterionAsOnPremises     = AssessmentSizingCriterion("AsOnPremises")
)

func (AssessmentSizingCriterion) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e AssessmentSizingCriterion) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AssessmentSizingCriterion) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AssessmentSizingCriterion) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AssessmentSizingCriterion) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// User configurable setting that describes the status of the assessment.
type AssessmentStage pulumi.String

const (
	AssessmentStageInProgress  = AssessmentStage("InProgress")
	AssessmentStageUnderReview = AssessmentStage("UnderReview")
	AssessmentStageApproved    = AssessmentStage("Approved")
)

func (AssessmentStage) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e AssessmentStage) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AssessmentStage) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AssessmentStage) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AssessmentStage) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Storage type selected for this disk.
type AzureDiskType pulumi.String

const (
	AzureDiskTypeUnknown           = AzureDiskType("Unknown")
	AzureDiskTypeStandard          = AzureDiskType("Standard")
	AzureDiskTypePremium           = AzureDiskType("Premium")
	AzureDiskTypeStandardSSD       = AzureDiskType("StandardSSD")
	AzureDiskTypeStandardOrPremium = AzureDiskType("StandardOrPremium")
)

func (AzureDiskType) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e AzureDiskType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureDiskType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureDiskType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AzureDiskType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// AHUB discount on windows virtual machines.
type AzureHybridUseBenefit pulumi.String

const (
	AzureHybridUseBenefitUnknown = AzureHybridUseBenefit("Unknown")
	AzureHybridUseBenefitYes     = AzureHybridUseBenefit("Yes")
	AzureHybridUseBenefitNo      = AzureHybridUseBenefit("No")
)

func (AzureHybridUseBenefit) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e AzureHybridUseBenefit) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureHybridUseBenefit) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureHybridUseBenefit) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AzureHybridUseBenefit) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Target Azure location for which the machines should be assessed. These enums are the same as used by Compute API.
type AzureLocation pulumi.String

const (
	AzureLocationUnknown            = AzureLocation("Unknown")
	AzureLocationEastAsia           = AzureLocation("EastAsia")
	AzureLocationSoutheastAsia      = AzureLocation("SoutheastAsia")
	AzureLocationAustraliaEast      = AzureLocation("AustraliaEast")
	AzureLocationAustraliaSoutheast = AzureLocation("AustraliaSoutheast")
	AzureLocationBrazilSouth        = AzureLocation("BrazilSouth")
	AzureLocationCanadaCentral      = AzureLocation("CanadaCentral")
	AzureLocationCanadaEast         = AzureLocation("CanadaEast")
	AzureLocationWestEurope         = AzureLocation("WestEurope")
	AzureLocationNorthEurope        = AzureLocation("NorthEurope")
	AzureLocationCentralIndia       = AzureLocation("CentralIndia")
	AzureLocationSouthIndia         = AzureLocation("SouthIndia")
	AzureLocationWestIndia          = AzureLocation("WestIndia")
	AzureLocationJapanEast          = AzureLocation("JapanEast")
	AzureLocationJapanWest          = AzureLocation("JapanWest")
	AzureLocationKoreaCentral       = AzureLocation("KoreaCentral")
	AzureLocationKoreaSouth         = AzureLocation("KoreaSouth")
	AzureLocationUkWest             = AzureLocation("UkWest")
	AzureLocationUkSouth            = AzureLocation("UkSouth")
	AzureLocationNorthCentralUs     = AzureLocation("NorthCentralUs")
	AzureLocationEastUs             = AzureLocation("EastUs")
	AzureLocationWestUs2            = AzureLocation("WestUs2")
	AzureLocationSouthCentralUs     = AzureLocation("SouthCentralUs")
	AzureLocationCentralUs          = AzureLocation("CentralUs")
	AzureLocationEastUs2            = AzureLocation("EastUs2")
	AzureLocationWestUs             = AzureLocation("WestUs")
	AzureLocationWestCentralUs      = AzureLocation("WestCentralUs")
	AzureLocationGermanyCentral     = AzureLocation("GermanyCentral")
	AzureLocationGermanyNortheast   = AzureLocation("GermanyNortheast")
	AzureLocationChinaNorth         = AzureLocation("ChinaNorth")
	AzureLocationChinaEast          = AzureLocation("ChinaEast")
	AzureLocationUSGovArizona       = AzureLocation("USGovArizona")
	AzureLocationUSGovTexas         = AzureLocation("USGovTexas")
	AzureLocationUSGovIowa          = AzureLocation("USGovIowa")
	AzureLocationUSGovVirginia      = AzureLocation("USGovVirginia")
	AzureLocationUSDoDCentral       = AzureLocation("USDoDCentral")
	AzureLocationUSDoDEast          = AzureLocation("USDoDEast")
)

func (AzureLocation) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e AzureLocation) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureLocation) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureLocation) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AzureLocation) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Offer code according to which cost estimation is done.
type AzureOfferCode pulumi.String

const (
	AzureOfferCodeUnknown         = AzureOfferCode("Unknown")
	AzureOfferCodeMSAZR0003P      = AzureOfferCode("MSAZR0003P")
	AzureOfferCodeMSAZR0044P      = AzureOfferCode("MSAZR0044P")
	AzureOfferCodeMSAZR0059P      = AzureOfferCode("MSAZR0059P")
	AzureOfferCodeMSAZR0060P      = AzureOfferCode("MSAZR0060P")
	AzureOfferCodeMSAZR0062P      = AzureOfferCode("MSAZR0062P")
	AzureOfferCodeMSAZR0063P      = AzureOfferCode("MSAZR0063P")
	AzureOfferCodeMSAZR0064P      = AzureOfferCode("MSAZR0064P")
	AzureOfferCodeMSAZR0029P      = AzureOfferCode("MSAZR0029P")
	AzureOfferCodeMSAZR0022P      = AzureOfferCode("MSAZR0022P")
	AzureOfferCodeMSAZR0023P      = AzureOfferCode("MSAZR0023P")
	AzureOfferCodeMSAZR0148P      = AzureOfferCode("MSAZR0148P")
	AzureOfferCodeMSAZR0025P      = AzureOfferCode("MSAZR0025P")
	AzureOfferCodeMSAZR0036P      = AzureOfferCode("MSAZR0036P")
	AzureOfferCodeMSAZR0120P      = AzureOfferCode("MSAZR0120P")
	AzureOfferCodeMSAZR0121P      = AzureOfferCode("MSAZR0121P")
	AzureOfferCodeMSAZR0122P      = AzureOfferCode("MSAZR0122P")
	AzureOfferCodeMSAZR0123P      = AzureOfferCode("MSAZR0123P")
	AzureOfferCodeMSAZR0124P      = AzureOfferCode("MSAZR0124P")
	AzureOfferCodeMSAZR0125P      = AzureOfferCode("MSAZR0125P")
	AzureOfferCodeMSAZR0126P      = AzureOfferCode("MSAZR0126P")
	AzureOfferCodeMSAZR0127P      = AzureOfferCode("MSAZR0127P")
	AzureOfferCodeMSAZR0128P      = AzureOfferCode("MSAZR0128P")
	AzureOfferCodeMSAZR0129P      = AzureOfferCode("MSAZR0129P")
	AzureOfferCodeMSAZR0130P      = AzureOfferCode("MSAZR0130P")
	AzureOfferCodeMSAZR0111P      = AzureOfferCode("MSAZR0111P")
	AzureOfferCodeMSAZR0144P      = AzureOfferCode("MSAZR0144P")
	AzureOfferCodeMSAZR0149P      = AzureOfferCode("MSAZR0149P")
	AzureOfferCodeMSMCAZR0044P    = AzureOfferCode("MSMCAZR0044P")
	AzureOfferCodeMSMCAZR0059P    = AzureOfferCode("MSMCAZR0059P")
	AzureOfferCodeMSMCAZR0060P    = AzureOfferCode("MSMCAZR0060P")
	AzureOfferCodeMSMCAZR0063P    = AzureOfferCode("MSMCAZR0063P")
	AzureOfferCodeMSMCAZR0120P    = AzureOfferCode("MSMCAZR0120P")
	AzureOfferCodeMSMCAZR0121P    = AzureOfferCode("MSMCAZR0121P")
	AzureOfferCodeMSMCAZR0125P    = AzureOfferCode("MSMCAZR0125P")
	AzureOfferCodeMSMCAZR0128P    = AzureOfferCode("MSMCAZR0128P")
	AzureOfferCodeMSAZRDE0003P    = AzureOfferCode("MSAZRDE0003P")
	AzureOfferCodeMSAZRDE0044P    = AzureOfferCode("MSAZRDE0044P")
	AzureOfferCodeMSAZRUSGOV0003P = AzureOfferCode("MSAZRUSGOV0003P")
	AzureOfferCodeEA              = AzureOfferCode("EA")
)

func (AzureOfferCode) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e AzureOfferCode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureOfferCode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureOfferCode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AzureOfferCode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Pricing tier for Size evaluation.
type AzurePricingTier pulumi.String

const (
	AzurePricingTierStandard = AzurePricingTier("Standard")
	AzurePricingTierBasic    = AzurePricingTier("Basic")
)

func (AzurePricingTier) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e AzurePricingTier) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzurePricingTier) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzurePricingTier) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AzurePricingTier) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Storage Redundancy type offered by Azure.
type AzureStorageRedundancy pulumi.String

const (
	AzureStorageRedundancyUnknown                = AzureStorageRedundancy("Unknown")
	AzureStorageRedundancyLocallyRedundant       = AzureStorageRedundancy("LocallyRedundant")
	AzureStorageRedundancyZoneRedundant          = AzureStorageRedundancy("ZoneRedundant")
	AzureStorageRedundancyGeoRedundant           = AzureStorageRedundancy("GeoRedundant")
	AzureStorageRedundancyReadAccessGeoRedundant = AzureStorageRedundancy("ReadAccessGeoRedundant")
)

func (AzureStorageRedundancy) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e AzureStorageRedundancy) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureStorageRedundancy) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureStorageRedundancy) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AzureStorageRedundancy) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Azure VM family.
type AzureVmFamily pulumi.String

const (
	AzureVmFamilyUnknown          = AzureVmFamily("Unknown")
	AzureVmFamily_Basic_A0_A4     = AzureVmFamily("Basic_A0_A4")
	AzureVmFamily_Standard_A0_A7  = AzureVmFamily("Standard_A0_A7")
	AzureVmFamily_Standard_A8_A11 = AzureVmFamily("Standard_A8_A11")
	AzureVmFamily_Av2_series      = AzureVmFamily("Av2_series")
	AzureVmFamily_D_series        = AzureVmFamily("D_series")
	AzureVmFamily_Dv2_series      = AzureVmFamily("Dv2_series")
	AzureVmFamily_DS_series       = AzureVmFamily("DS_series")
	AzureVmFamily_DSv2_series     = AzureVmFamily("DSv2_series")
	AzureVmFamily_F_series        = AzureVmFamily("F_series")
	AzureVmFamily_Fs_series       = AzureVmFamily("Fs_series")
	AzureVmFamily_G_series        = AzureVmFamily("G_series")
	AzureVmFamily_GS_series       = AzureVmFamily("GS_series")
	AzureVmFamily_H_series        = AzureVmFamily("H_series")
	AzureVmFamily_Ls_series       = AzureVmFamily("Ls_series")
	AzureVmFamily_Dsv3_series     = AzureVmFamily("Dsv3_series")
	AzureVmFamily_Dv3_series      = AzureVmFamily("Dv3_series")
	AzureVmFamily_Fsv2_series     = AzureVmFamily("Fsv2_series")
	AzureVmFamily_Ev3_series      = AzureVmFamily("Ev3_series")
	AzureVmFamily_Esv3_series     = AzureVmFamily("Esv3_series")
	AzureVmFamily_M_series        = AzureVmFamily("M_series")
	AzureVmFamily_DC_Series       = AzureVmFamily("DC_Series")
)

func (AzureVmFamily) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e AzureVmFamily) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureVmFamily) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureVmFamily) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AzureVmFamily) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Currency to report prices in.
type Currency pulumi.String

const (
	CurrencyUnknown = Currency("Unknown")
	CurrencyUSD     = Currency("USD")
	CurrencyDKK     = Currency("DKK")
	CurrencyCAD     = Currency("CAD")
	CurrencyIDR     = Currency("IDR")
	CurrencyJPY     = Currency("JPY")
	CurrencyKRW     = Currency("KRW")
	CurrencyNZD     = Currency("NZD")
	CurrencyNOK     = Currency("NOK")
	CurrencyRUB     = Currency("RUB")
	CurrencySAR     = Currency("SAR")
	CurrencyZAR     = Currency("ZAR")
	CurrencySEK     = Currency("SEK")
	CurrencyTRY     = Currency("TRY")
	CurrencyGBP     = Currency("GBP")
	CurrencyMXN     = Currency("MXN")
	CurrencyMYR     = Currency("MYR")
	CurrencyINR     = Currency("INR")
	CurrencyHKD     = Currency("HKD")
	CurrencyBRL     = Currency("BRL")
	CurrencyTWD     = Currency("TWD")
	CurrencyEUR     = Currency("EUR")
	CurrencyCHF     = Currency("CHF")
	CurrencyARS     = Currency("ARS")
	CurrencyAUD     = Currency("AUD")
	CurrencyCNY     = Currency("CNY")
)

func (Currency) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e Currency) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Currency) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Currency) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Currency) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Percentile of performance data used to recommend Azure size.
type Percentile pulumi.String

const (
	PercentilePercentile50 = Percentile("Percentile50")
	PercentilePercentile90 = Percentile("Percentile90")
	PercentilePercentile95 = Percentile("Percentile95")
	PercentilePercentile99 = Percentile("Percentile99")
)

func (Percentile) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e Percentile) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Percentile) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Percentile) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Percentile) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Assessment project status.
type ProjectStatus pulumi.String

const (
	ProjectStatusActive   = ProjectStatus("Active")
	ProjectStatusInactive = ProjectStatus("Inactive")
)

func (ProjectStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e ProjectStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProjectStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProjectStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ProjectStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Azure reserved instance.
type ReservedInstance pulumi.String

const (
	ReservedInstanceNone    = ReservedInstance("None")
	ReservedInstanceRI1Year = ReservedInstance("RI1Year")
	ReservedInstanceRI3Year = ReservedInstance("RI3Year")
)

func (ReservedInstance) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e ReservedInstance) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ReservedInstance) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ReservedInstance) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ReservedInstance) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// The type of identity used for the resource mover service.
type ResourceIdentityType pulumi.String

const (
	ResourceIdentityTypeNone           = ResourceIdentityType("None")
	ResourceIdentityTypeSystemAssigned = ResourceIdentityType("SystemAssigned")
	ResourceIdentityTypeUserAssigned   = ResourceIdentityType("UserAssigned")
)

func (ResourceIdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e ResourceIdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceIdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceIdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ResourceIdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Gets or sets the target availability zone.
type TargetAvailabilityZone pulumi.String

const (
	TargetAvailabilityZoneOne   = TargetAvailabilityZone("1")
	TargetAvailabilityZoneTwo   = TargetAvailabilityZone("2")
	TargetAvailabilityZoneThree = TargetAvailabilityZone("3")
	TargetAvailabilityZoneNA    = TargetAvailabilityZone("NA")
)

func (TargetAvailabilityZone) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e TargetAvailabilityZone) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TargetAvailabilityZone) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TargetAvailabilityZone) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TargetAvailabilityZone) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Time range of performance data used to recommend a size.
type TimeRange pulumi.String

const (
	TimeRangeDay    = TimeRange("Day")
	TimeRangeWeek   = TimeRange("Week")
	TimeRangeMonth  = TimeRange("Month")
	TimeRangeCustom = TimeRange("Custom")
)

func (TimeRange) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e TimeRange) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TimeRange) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TimeRange) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TimeRange) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

// Defines the zone redundant resource setting.
type ZoneRedundant pulumi.String

const (
	ZoneRedundantEnable  = ZoneRedundant("Enable")
	ZoneRedundantDisable = ZoneRedundant("Disable")
)

func (ZoneRedundant) ElementType() reflect.Type {
	return reflect.TypeOf((*pulumi.String)(nil)).Elem()
}

func (e ZoneRedundant) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ZoneRedundant) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ZoneRedundant) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ZoneRedundant) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}
