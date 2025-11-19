package ruler

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("mimir_rule_group_alerting", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "mimir"
		r.ShortGroup = "ruler"
	})

	p.AddResourceConfigurator("mimir_rule_group_recording", func(r *config.Resource) {
		r.ShortGroup = "ruler"
	})

	p.AddResourceConfigurator("mimir_rules", func(r *config.Resource) {
		r.ShortGroup = "ruler"
	})
}
