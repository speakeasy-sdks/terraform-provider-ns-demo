// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Security struct {
	APIKey string `security:"scheme,type=apiKey,subtype=header,name=Netskope-Api-Token"`
}
