package servicebus

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// DefaultAction enumerates the values for default action.
type DefaultAction string

const (
	// Allow ...
	Allow DefaultAction = "Allow"
	// Deny ...
	Deny DefaultAction = "Deny"
)

// PossibleDefaultActionValues returns an array of possible values for the DefaultAction const type.
func PossibleDefaultActionValues() []DefaultAction {
	return []DefaultAction{Allow, Deny}
}

// IdentityType enumerates the values for identity type.
type IdentityType string

const (
	// SystemAssigned ...
	SystemAssigned IdentityType = "SystemAssigned"
)

// PossibleIdentityTypeValues returns an array of possible values for the IdentityType const type.
func PossibleIdentityTypeValues() []IdentityType {
	return []IdentityType{SystemAssigned}
}

// IPAction enumerates the values for ip action.
type IPAction string

const (
	// Accept ...
	Accept IPAction = "Accept"
	// Reject ...
	Reject IPAction = "Reject"
)

// PossibleIPActionValues returns an array of possible values for the IPAction const type.
func PossibleIPActionValues() []IPAction {
	return []IPAction{Accept, Reject}
}

// KeySource enumerates the values for key source.
type KeySource string

const (
	// MicrosoftKeyVault ...
	MicrosoftKeyVault KeySource = "Microsoft.KeyVault"
)

// PossibleKeySourceValues returns an array of possible values for the KeySource const type.
func PossibleKeySourceValues() []KeySource {
	return []KeySource{MicrosoftKeyVault}
}

// NetworkRuleIPAction enumerates the values for network rule ip action.
type NetworkRuleIPAction string

const (
	// NetworkRuleIPActionAllow ...
	NetworkRuleIPActionAllow NetworkRuleIPAction = "Allow"
)

// PossibleNetworkRuleIPActionValues returns an array of possible values for the NetworkRuleIPAction const type.
func PossibleNetworkRuleIPActionValues() []NetworkRuleIPAction {
	return []NetworkRuleIPAction{NetworkRuleIPActionAllow}
}

// SkuName enumerates the values for sku name.
type SkuName string

const (
	// Basic ...
	Basic SkuName = "Basic"
	// Premium ...
	Premium SkuName = "Premium"
	// Standard ...
	Standard SkuName = "Standard"
)

// PossibleSkuNameValues returns an array of possible values for the SkuName const type.
func PossibleSkuNameValues() []SkuName {
	return []SkuName{Basic, Premium, Standard}
}

// SkuTier enumerates the values for sku tier.
type SkuTier string

const (
	// SkuTierBasic ...
	SkuTierBasic SkuTier = "Basic"
	// SkuTierPremium ...
	SkuTierPremium SkuTier = "Premium"
	// SkuTierStandard ...
	SkuTierStandard SkuTier = "Standard"
)

// PossibleSkuTierValues returns an array of possible values for the SkuTier const type.
func PossibleSkuTierValues() []SkuTier {
	return []SkuTier{SkuTierBasic, SkuTierPremium, SkuTierStandard}
}
