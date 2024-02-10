/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// MetricsApplyConfiguration represents an declarative configuration of the Metrics type for use
// with apply.
type MetricsApplyConfiguration struct {
	Name        *string                        `json:"name,omitempty"`
	Help        *string                        `json:"help,omitempty"`
	Type        *string                        `json:"type,omitempty"`
	Field       *FieldApplyConfiguration       `json:"field,omitempty"`
	Labels      []LabelApplyConfiguration      `json:"labels,omitempty"`
	Params      []ParameterApplyConfiguration  `json:"params,omitempty"`
	States      *StateApplyConfiguration       `json:"states,omitempty"`
	MetricValue *MetricValueApplyConfiguration `json:"metricValue,omitempty"`
}

// MetricsApplyConfiguration constructs an declarative configuration of the Metrics type for use with
// apply.
func Metrics() *MetricsApplyConfiguration {
	return &MetricsApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *MetricsApplyConfiguration) WithName(value string) *MetricsApplyConfiguration {
	b.Name = &value
	return b
}

// WithHelp sets the Help field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Help field is set to the value of the last call.
func (b *MetricsApplyConfiguration) WithHelp(value string) *MetricsApplyConfiguration {
	b.Help = &value
	return b
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *MetricsApplyConfiguration) WithType(value string) *MetricsApplyConfiguration {
	b.Type = &value
	return b
}

// WithField sets the Field field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Field field is set to the value of the last call.
func (b *MetricsApplyConfiguration) WithField(value *FieldApplyConfiguration) *MetricsApplyConfiguration {
	b.Field = value
	return b
}

// WithLabels adds the given value to the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Labels field.
func (b *MetricsApplyConfiguration) WithLabels(values ...*LabelApplyConfiguration) *MetricsApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithLabels")
		}
		b.Labels = append(b.Labels, *values[i])
	}
	return b
}

// WithParams adds the given value to the Params field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Params field.
func (b *MetricsApplyConfiguration) WithParams(values ...*ParameterApplyConfiguration) *MetricsApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithParams")
		}
		b.Params = append(b.Params, *values[i])
	}
	return b
}

// WithStates sets the States field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the States field is set to the value of the last call.
func (b *MetricsApplyConfiguration) WithStates(value *StateApplyConfiguration) *MetricsApplyConfiguration {
	b.States = value
	return b
}

// WithMetricValue sets the MetricValue field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MetricValue field is set to the value of the last call.
func (b *MetricsApplyConfiguration) WithMetricValue(value *MetricValueApplyConfiguration) *MetricsApplyConfiguration {
	b.MetricValue = value
	return b
}