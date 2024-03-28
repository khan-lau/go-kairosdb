// Copyright 2016 Ajit Yagaty
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package builder

import (
	"fmt"
)

var (
	// Metric Errors.
	ErrorMetricNameInvalid = fmt.Errorf("%s", "Metric name empty")
	ErrorTagNameInvalid    = fmt.Errorf("%s", "Tag name empty")
	ErrorTagValueInvalid   = fmt.Errorf("%s", "Tag value empty")
	ErrorTTLInvalid        = fmt.Errorf("%s", "TTL value invalid")

	// Data Point Errors.
	ErrorDataPointInt64   = fmt.Errorf("%s", "Not an int64 data value")
	ErrorDataPointFloat32 = fmt.Errorf("%s", "Not a float32 data value")
	ErrorDataPointFloat64 = fmt.Errorf("%s", "Not a float64 data value")

	// Query Metric Errors.
	ErrorQMetricNameInvalid     = fmt.Errorf("%s", "Query Metric name empty")
	ErrorQMetricTagNameInvalid  = fmt.Errorf("%s", "Query Metric Tag name empty")
	ErrorQMetricTagValueInvalid = fmt.Errorf("%s", "Query Metric Tag value empty")
	ErrorQMetricLimitInvalid    = fmt.Errorf("%s", "Query Metric Limit must be >= 0")

	// Query Builder Errors.
	ErrorAbsRelativeStartSet      = fmt.Errorf("%s", "Both absolute and relative start times cannot be set")
	ErrorRelativeStartTimeInvalid = fmt.Errorf("%s", "Relative start time duration must be > 0")
	ErrorAbsRelativeEndSet        = fmt.Errorf("%s", "Both absolute and relative end times cannot be set")
	ErrorRelativeEndTimeInvalid   = fmt.Errorf("%s", "Relative end time duration must be > 0")
	ErrorStartTimeNotSpecified    = fmt.Errorf("%s", "Start time not specified")
)
