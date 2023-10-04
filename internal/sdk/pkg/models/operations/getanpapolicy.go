// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetANpaPolicyRequest struct {
	// Return values only from specified fields
	Fields *string `queryParam:"style=form,explode=true,name=fields"`
	ID     string  `pathParam:"style=simple,explode=false,name=id"`
}

// GetANpaPolicy400ApplicationJSON - Invalid request
type GetANpaPolicy400ApplicationJSON struct {
	Result *string `json:"result,omitempty"`
	Status *string `json:"status,omitempty"`
}

type GetANpaPolicy200ApplicationJSONDataRuleDataMatchCriteriaAction struct {
	ActionName *string `json:"action_name,omitempty"`
}

type GetANpaPolicy200ApplicationJSONDataRuleDataPrivateAppsWithActivitiesActivities struct {
	Activity          *string  `json:"activity,omitempty"`
	ListOfConstraints []string `json:"list_of_constraints,omitempty"`
}

type GetANpaPolicy200ApplicationJSONDataRuleDataPrivateAppsWithActivities struct {
	Activities []GetANpaPolicy200ApplicationJSONDataRuleDataPrivateAppsWithActivitiesActivities `json:"activities,omitempty"`
	AppName    *string                                                                          `json:"appName,omitempty"`
}

type GetANpaPolicy200ApplicationJSONDataRuleData struct {
	AccessMethod              *string                                                                `json:"access_method,omitempty"`
	BNegateNetLocation        *string                                                                `json:"b_negateNetLocation,omitempty"`
	BNegateSrcCountries       *string                                                                `json:"b_negateSrcCountries,omitempty"`
	Classification            *string                                                                `json:"classification,omitempty"`
	ExcludedUsers             []string                                                               `json:"excludedUsers,omitempty"`
	ExternalDlp               *string                                                                `json:"external_dlp,omitempty"`
	JSONVersion               *string                                                                `json:"json_version,omitempty"`
	MatchCriteriaAction       *GetANpaPolicy200ApplicationJSONDataRuleDataMatchCriteriaAction        `json:"match_criteria_action,omitempty"`
	NetLocationObj            []string                                                               `json:"net_location_obj,omitempty"`
	PolicyType                *string                                                                `json:"policy_type,omitempty"`
	PrivateAppIds             []string                                                               `json:"privateAppIds,omitempty"`
	PrivateAppTagIds          []string                                                               `json:"privateAppTagIds,omitempty"`
	PrivateAppTags            []string                                                               `json:"privateAppTags,omitempty"`
	PrivateApps               []string                                                               `json:"privateApps,omitempty"`
	PrivateAppsWithActivities []GetANpaPolicy200ApplicationJSONDataRuleDataPrivateAppsWithActivities `json:"privateAppsWithActivities,omitempty"`
	ShowDlpProfileActionTable *string                                                                `json:"show_dlp_profile_action_table,omitempty"`
	SrcCountries              []string                                                               `json:"srcCountries,omitempty"`
	UserType                  *string                                                                `json:"userType,omitempty"`
	Users                     []string                                                               `json:"users,omitempty"`
	Version                   *string                                                                `json:"version,omitempty"`
}

type GetANpaPolicy200ApplicationJSONData struct {
	GroupID  *string                                      `json:"group_id,omitempty"`
	RuleData *GetANpaPolicy200ApplicationJSONDataRuleData `json:"rule_data,omitempty"`
	RuleID   *string                                      `json:"rule_id,omitempty"`
	RuleName *string                                      `json:"rule_name,omitempty"`
}

// GetANpaPolicy200ApplicationJSON - successful operation
type GetANpaPolicy200ApplicationJSON struct {
	Data   *GetANpaPolicy200ApplicationJSONData `json:"data,omitempty"`
	Status *string                              `json:"status,omitempty"`
}

type GetANpaPolicyResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// successful operation
	GetANpaPolicy200ApplicationJSONObject *GetANpaPolicy200ApplicationJSON
	// Invalid request
	GetANpaPolicy400ApplicationJSONObject *GetANpaPolicy400ApplicationJSON
}
