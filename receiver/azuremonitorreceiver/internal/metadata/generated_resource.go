// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
)

// ResourceBuilder is a helper struct to build resources predefined in metadata.yaml.
type ResourceBuilder struct {
	config ResourceAttributesConfig
	res    pcommon.Resource
}

func NewResourceBuilder(rac ResourceAttributesConfig) *ResourceBuilder {
	return &ResourceBuilder{
		config: rac,
		res:    pcommon.NewResource(),
	}
}

// SetAzuremonitorSubscriptionID sets provided value as "azuremonitor.subscription_id" attribute.
func (rb *ResourceBuilder) SetAzuremonitorSubscriptionID(val string) {
	if rb.config.AzuremonitorSubscriptionID.Enabled {
		rb.res.Attributes().PutStr("azuremonitor.subscription_id", val)
	}
}

// SetAzuremonitorTenantID sets provided value as "azuremonitor.tenant_id" attribute.
func (rb *ResourceBuilder) SetAzuremonitorTenantID(val string) {
	if rb.config.AzuremonitorTenantID.Enabled {
		rb.res.Attributes().PutStr("azuremonitor.tenant_id", val)
	}
}

// Emit returns the built resource and resets the internal builder state.
func (rb *ResourceBuilder) Emit() pcommon.Resource {
	r := rb.res
	rb.res = pcommon.NewResource()
	return r
}
