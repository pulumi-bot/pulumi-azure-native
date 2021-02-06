// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A Data Lake Analytics account object, containing all information associated with the named Data Lake Analytics account.
// Latest API Version: 2016-11-01.
type Account struct {
	pulumi.CustomResourceState

	// The unique identifier associated with this Data Lake Analytics account.
	AccountId pulumi.StringOutput `pulumi:"accountId"`
	// The list of compute policies associated with this account.
	ComputePolicies ComputePolicyResponseArrayOutput `pulumi:"computePolicies"`
	// The account creation time.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// The commitment tier in use for the current month.
	CurrentTier pulumi.StringOutput `pulumi:"currentTier"`
	// The list of Data Lake Store accounts associated with this account.
	DataLakeStoreAccounts DataLakeStoreAccountInformationResponseArrayOutput `pulumi:"dataLakeStoreAccounts"`
	// The current state of the DebugDataAccessLevel for this account.
	DebugDataAccessLevel pulumi.StringOutput `pulumi:"debugDataAccessLevel"`
	// The default Data Lake Store account associated with this account.
	DefaultDataLakeStoreAccount pulumi.StringOutput `pulumi:"defaultDataLakeStoreAccount"`
	// The full CName endpoint for this account.
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// The current state of allowing or disallowing IPs originating within Azure through the firewall. If the firewall is disabled, this is not enforced.
	FirewallAllowAzureIps pulumi.StringOutput `pulumi:"firewallAllowAzureIps"`
	// The list of firewall rules associated with this account.
	FirewallRules FirewallRuleResponseArrayOutput `pulumi:"firewallRules"`
	// The current state of the IP address firewall for this account.
	FirewallState pulumi.StringOutput `pulumi:"firewallState"`
	// The list of hiveMetastores associated with this account.
	HiveMetastores HiveMetastoreResponseArrayOutput `pulumi:"hiveMetastores"`
	// The account last modified time.
	LastModifiedTime pulumi.StringOutput `pulumi:"lastModifiedTime"`
	// The resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// The maximum supported degree of parallelism for this account.
	MaxDegreeOfParallelism pulumi.IntPtrOutput `pulumi:"maxDegreeOfParallelism"`
	// The maximum supported degree of parallelism per job for this account.
	MaxDegreeOfParallelismPerJob pulumi.IntOutput `pulumi:"maxDegreeOfParallelismPerJob"`
	// The maximum supported jobs running under the account at the same time.
	MaxJobCount pulumi.IntPtrOutput `pulumi:"maxJobCount"`
	// The minimum supported priority per job for this account.
	MinPriorityPerJob pulumi.IntOutput `pulumi:"minPriorityPerJob"`
	// The resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The commitment tier for the next month.
	NewTier pulumi.StringOutput `pulumi:"newTier"`
	// The provisioning status of the Data Lake Analytics account.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The number of days that job metadata is retained.
	QueryStoreRetention pulumi.IntPtrOutput `pulumi:"queryStoreRetention"`
	// The state of the Data Lake Analytics account.
	State pulumi.StringOutput `pulumi:"state"`
	// The list of Azure Blob Storage accounts associated with this account.
	StorageAccounts StorageAccountInformationResponseArrayOutput `pulumi:"storageAccounts"`
	// The system defined maximum supported degree of parallelism for this account, which restricts the maximum value of parallelism the user can set for the account.
	SystemMaxDegreeOfParallelism pulumi.IntOutput `pulumi:"systemMaxDegreeOfParallelism"`
	// The system defined maximum supported jobs running under the account at the same time, which restricts the maximum number of running jobs the user can set for the account.
	SystemMaxJobCount pulumi.IntOutput `pulumi:"systemMaxJobCount"`
	// The resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// The list of virtualNetwork rules associated with this account.
	VirtualNetworkRules VirtualNetworkRuleResponseArrayOutput `pulumi:"virtualNetworkRules"`
}

// NewAccount registers a new resource with the given unique name, arguments, and options.
func NewAccount(ctx *pulumi.Context,
	name string, args *AccountArgs, opts ...pulumi.ResourceOption) (*Account, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.DataLakeStoreAccounts == nil {
		return nil, errors.New("invalid value for required argument 'DataLakeStoreAccounts'")
	}
	if args.DefaultDataLakeStoreAccount == nil {
		return nil, errors.New("invalid value for required argument 'DefaultDataLakeStoreAccount'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.MaxDegreeOfParallelism == nil {
		args.MaxDegreeOfParallelism = pulumi.IntPtr(30)
	}
	if args.MaxJobCount == nil {
		args.MaxJobCount = pulumi.IntPtr(3)
	}
	if args.QueryStoreRetention == nil {
		args.QueryStoreRetention = pulumi.IntPtr(30)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:datalakeanalytics/v20151001preview:Account"),
		},
		{
			Type: pulumi.String("azure-nextgen:datalakeanalytics/v20161101:Account"),
		},
	})
	opts = append(opts, aliases)
	var resource Account
	err := ctx.RegisterResource("azure-nextgen:datalakeanalytics/latest:Account", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccount gets an existing Account resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccountState, opts ...pulumi.ResourceOption) (*Account, error) {
	var resource Account
	err := ctx.ReadResource("azure-nextgen:datalakeanalytics/latest:Account", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Account resources.
type accountState struct {
	// The unique identifier associated with this Data Lake Analytics account.
	AccountId *string `pulumi:"accountId"`
	// The list of compute policies associated with this account.
	ComputePolicies []ComputePolicyResponse `pulumi:"computePolicies"`
	// The account creation time.
	CreationTime *string `pulumi:"creationTime"`
	// The commitment tier in use for the current month.
	CurrentTier *string `pulumi:"currentTier"`
	// The list of Data Lake Store accounts associated with this account.
	DataLakeStoreAccounts []DataLakeStoreAccountInformationResponse `pulumi:"dataLakeStoreAccounts"`
	// The current state of the DebugDataAccessLevel for this account.
	DebugDataAccessLevel *string `pulumi:"debugDataAccessLevel"`
	// The default Data Lake Store account associated with this account.
	DefaultDataLakeStoreAccount *string `pulumi:"defaultDataLakeStoreAccount"`
	// The full CName endpoint for this account.
	Endpoint *string `pulumi:"endpoint"`
	// The current state of allowing or disallowing IPs originating within Azure through the firewall. If the firewall is disabled, this is not enforced.
	FirewallAllowAzureIps *string `pulumi:"firewallAllowAzureIps"`
	// The list of firewall rules associated with this account.
	FirewallRules []FirewallRuleResponse `pulumi:"firewallRules"`
	// The current state of the IP address firewall for this account.
	FirewallState *string `pulumi:"firewallState"`
	// The list of hiveMetastores associated with this account.
	HiveMetastores []HiveMetastoreResponse `pulumi:"hiveMetastores"`
	// The account last modified time.
	LastModifiedTime *string `pulumi:"lastModifiedTime"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The maximum supported degree of parallelism for this account.
	MaxDegreeOfParallelism *int `pulumi:"maxDegreeOfParallelism"`
	// The maximum supported degree of parallelism per job for this account.
	MaxDegreeOfParallelismPerJob *int `pulumi:"maxDegreeOfParallelismPerJob"`
	// The maximum supported jobs running under the account at the same time.
	MaxJobCount *int `pulumi:"maxJobCount"`
	// The minimum supported priority per job for this account.
	MinPriorityPerJob *int `pulumi:"minPriorityPerJob"`
	// The resource name.
	Name *string `pulumi:"name"`
	// The commitment tier for the next month.
	NewTier *string `pulumi:"newTier"`
	// The provisioning status of the Data Lake Analytics account.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The number of days that job metadata is retained.
	QueryStoreRetention *int `pulumi:"queryStoreRetention"`
	// The state of the Data Lake Analytics account.
	State *string `pulumi:"state"`
	// The list of Azure Blob Storage accounts associated with this account.
	StorageAccounts []StorageAccountInformationResponse `pulumi:"storageAccounts"`
	// The system defined maximum supported degree of parallelism for this account, which restricts the maximum value of parallelism the user can set for the account.
	SystemMaxDegreeOfParallelism *int `pulumi:"systemMaxDegreeOfParallelism"`
	// The system defined maximum supported jobs running under the account at the same time, which restricts the maximum number of running jobs the user can set for the account.
	SystemMaxJobCount *int `pulumi:"systemMaxJobCount"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The resource type.
	Type *string `pulumi:"type"`
	// The list of virtualNetwork rules associated with this account.
	VirtualNetworkRules []VirtualNetworkRuleResponse `pulumi:"virtualNetworkRules"`
}

type AccountState struct {
	// The unique identifier associated with this Data Lake Analytics account.
	AccountId pulumi.StringPtrInput
	// The list of compute policies associated with this account.
	ComputePolicies ComputePolicyResponseArrayInput
	// The account creation time.
	CreationTime pulumi.StringPtrInput
	// The commitment tier in use for the current month.
	CurrentTier pulumi.StringPtrInput
	// The list of Data Lake Store accounts associated with this account.
	DataLakeStoreAccounts DataLakeStoreAccountInformationResponseArrayInput
	// The current state of the DebugDataAccessLevel for this account.
	DebugDataAccessLevel pulumi.StringPtrInput
	// The default Data Lake Store account associated with this account.
	DefaultDataLakeStoreAccount pulumi.StringPtrInput
	// The full CName endpoint for this account.
	Endpoint pulumi.StringPtrInput
	// The current state of allowing or disallowing IPs originating within Azure through the firewall. If the firewall is disabled, this is not enforced.
	FirewallAllowAzureIps pulumi.StringPtrInput
	// The list of firewall rules associated with this account.
	FirewallRules FirewallRuleResponseArrayInput
	// The current state of the IP address firewall for this account.
	FirewallState pulumi.StringPtrInput
	// The list of hiveMetastores associated with this account.
	HiveMetastores HiveMetastoreResponseArrayInput
	// The account last modified time.
	LastModifiedTime pulumi.StringPtrInput
	// The resource location.
	Location pulumi.StringPtrInput
	// The maximum supported degree of parallelism for this account.
	MaxDegreeOfParallelism pulumi.IntPtrInput
	// The maximum supported degree of parallelism per job for this account.
	MaxDegreeOfParallelismPerJob pulumi.IntPtrInput
	// The maximum supported jobs running under the account at the same time.
	MaxJobCount pulumi.IntPtrInput
	// The minimum supported priority per job for this account.
	MinPriorityPerJob pulumi.IntPtrInput
	// The resource name.
	Name pulumi.StringPtrInput
	// The commitment tier for the next month.
	NewTier pulumi.StringPtrInput
	// The provisioning status of the Data Lake Analytics account.
	ProvisioningState pulumi.StringPtrInput
	// The number of days that job metadata is retained.
	QueryStoreRetention pulumi.IntPtrInput
	// The state of the Data Lake Analytics account.
	State pulumi.StringPtrInput
	// The list of Azure Blob Storage accounts associated with this account.
	StorageAccounts StorageAccountInformationResponseArrayInput
	// The system defined maximum supported degree of parallelism for this account, which restricts the maximum value of parallelism the user can set for the account.
	SystemMaxDegreeOfParallelism pulumi.IntPtrInput
	// The system defined maximum supported jobs running under the account at the same time, which restricts the maximum number of running jobs the user can set for the account.
	SystemMaxJobCount pulumi.IntPtrInput
	// The resource tags.
	Tags pulumi.StringMapInput
	// The resource type.
	Type pulumi.StringPtrInput
	// The list of virtualNetwork rules associated with this account.
	VirtualNetworkRules VirtualNetworkRuleResponseArrayInput
}

func (AccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountState)(nil)).Elem()
}

type accountArgs struct {
	// The name of the Data Lake Analytics account.
	AccountName string `pulumi:"accountName"`
	// The list of compute policies associated with this account.
	ComputePolicies []CreateComputePolicyWithAccountParameters `pulumi:"computePolicies"`
	// The list of Data Lake Store accounts associated with this account.
	DataLakeStoreAccounts []AddDataLakeStoreWithAccountParameters `pulumi:"dataLakeStoreAccounts"`
	// The default Data Lake Store account associated with this account.
	DefaultDataLakeStoreAccount string `pulumi:"defaultDataLakeStoreAccount"`
	// The current state of allowing or disallowing IPs originating within Azure through the firewall. If the firewall is disabled, this is not enforced.
	FirewallAllowAzureIps *string `pulumi:"firewallAllowAzureIps"`
	// The list of firewall rules associated with this account.
	FirewallRules []CreateFirewallRuleWithAccountParameters `pulumi:"firewallRules"`
	// The current state of the IP address firewall for this account.
	FirewallState *string `pulumi:"firewallState"`
	// The resource location.
	Location *string `pulumi:"location"`
	// The maximum supported degree of parallelism for this account.
	MaxDegreeOfParallelism *int `pulumi:"maxDegreeOfParallelism"`
	// The maximum supported degree of parallelism per job for this account.
	MaxDegreeOfParallelismPerJob *int `pulumi:"maxDegreeOfParallelismPerJob"`
	// The maximum supported jobs running under the account at the same time.
	MaxJobCount *int `pulumi:"maxJobCount"`
	// The minimum supported priority per job for this account.
	MinPriorityPerJob *int `pulumi:"minPriorityPerJob"`
	// The commitment tier for the next month.
	NewTier *string `pulumi:"newTier"`
	// The number of days that job metadata is retained.
	QueryStoreRetention *int `pulumi:"queryStoreRetention"`
	// The name of the Azure resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The list of Azure Blob Storage accounts associated with this account.
	StorageAccounts []AddStorageAccountWithAccountParameters `pulumi:"storageAccounts"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Account resource.
type AccountArgs struct {
	// The name of the Data Lake Analytics account.
	AccountName pulumi.StringInput
	// The list of compute policies associated with this account.
	ComputePolicies CreateComputePolicyWithAccountParametersArrayInput
	// The list of Data Lake Store accounts associated with this account.
	DataLakeStoreAccounts AddDataLakeStoreWithAccountParametersArrayInput
	// The default Data Lake Store account associated with this account.
	DefaultDataLakeStoreAccount pulumi.StringInput
	// The current state of allowing or disallowing IPs originating within Azure through the firewall. If the firewall is disabled, this is not enforced.
	FirewallAllowAzureIps *FirewallAllowAzureIpsState
	// The list of firewall rules associated with this account.
	FirewallRules CreateFirewallRuleWithAccountParametersArrayInput
	// The current state of the IP address firewall for this account.
	FirewallState *FirewallState
	// The resource location.
	Location pulumi.StringPtrInput
	// The maximum supported degree of parallelism for this account.
	MaxDegreeOfParallelism pulumi.IntPtrInput
	// The maximum supported degree of parallelism per job for this account.
	MaxDegreeOfParallelismPerJob pulumi.IntPtrInput
	// The maximum supported jobs running under the account at the same time.
	MaxJobCount pulumi.IntPtrInput
	// The minimum supported priority per job for this account.
	MinPriorityPerJob pulumi.IntPtrInput
	// The commitment tier for the next month.
	NewTier *TierType
	// The number of days that job metadata is retained.
	QueryStoreRetention pulumi.IntPtrInput
	// The name of the Azure resource group.
	ResourceGroupName pulumi.StringInput
	// The list of Azure Blob Storage accounts associated with this account.
	StorageAccounts AddStorageAccountWithAccountParametersArrayInput
	// The resource tags.
	Tags pulumi.StringMapInput
}

func (AccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accountArgs)(nil)).Elem()
}

type AccountInput interface {
	pulumi.Input

	ToAccountOutput() AccountOutput
	ToAccountOutputWithContext(ctx context.Context) AccountOutput
}

func (*Account) ElementType() reflect.Type {
	return reflect.TypeOf((*Account)(nil))
}

func (i *Account) ToAccountOutput() AccountOutput {
	return i.ToAccountOutputWithContext(context.Background())
}

func (i *Account) ToAccountOutputWithContext(ctx context.Context) AccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccountOutput)
}

type AccountOutput struct {
	*pulumi.OutputState
}

func (AccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Account)(nil))
}

func (o AccountOutput) ToAccountOutput() AccountOutput {
	return o
}

func (o AccountOutput) ToAccountOutputWithContext(ctx context.Context) AccountOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AccountOutput{})
}
