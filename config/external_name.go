package config

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Alertmanager configuration - uses org_id as the identifier
	"mimir_alertmanager_config": config.IdentifierFromProvider,

	// Rule groups - use a composite identifier: namespace/name
	"mimir_rule_group_alerting":  config.TemplatedStringAsIdentifier("name", "{{ .parameters.namespace }}/{{ .external_name }}"),
	"mimir_rule_group_recording": config.TemplatedStringAsIdentifier("name", "{{ .parameters.namespace }}/{{ .external_name }}"),

	// Rules resource - manages multiple rule groups from a file
	"mimir_rules": config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
