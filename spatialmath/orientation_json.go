package spatialmath

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

// OrientationType defines what orientation representations are known.
type OrientationType string

// The set of allowed representations for orientation.
const (
	NoOrientationType            = OrientationType("")
	OrientationVectorDegreesType = OrientationType("ov_degrees")
	OrientationVectorRadiansType = OrientationType("ov_radians")
	EulerAnglesType              = OrientationType("euler_angles")
	AxisAnglesType               = OrientationType("axis_angles")
)

// OrientationMap encodes the orientation interface to something serializable and human readable.
func OrientationMap(o Orientation) (map[string]interface{}, error) {
	switch v := o.(type) {
	case *R4AA:
		return map[string]interface{}{"type": string(AxisAnglesType), "value": v}, nil
	case *OrientationVector:
		return map[string]interface{}{"type": string(OrientationVectorRadiansType), "value": v}, nil
	case *OrientationVectorDegrees:
		return map[string]interface{}{"type": string(OrientationVectorDegreesType), "value": v}, nil
	case *EulerAngles:
		return map[string]interface{}{"type": string(EulerAnglesType), "value": v}, nil
	default:
		return nil, errors.Errorf("do not know how to map Orientation type %T to json fields", v)
	}
}

// OrientationConfig holds the underlying type of orientation, and the value.
type OrientationConfig struct {
	Type  string          `json:"type"`
	Value json.RawMessage `json:"value"`
}

// ParseConfig will use the Type in OrientationConfig and convert into the correct struct that implements Orientation.
func (config OrientationConfig) ParseConfig() (Orientation, error) {
	var err error
	// use the type to unmarshal the value
	switch OrientationType(config.Type) {
	case NoOrientationType:
		return NewZeroOrientation(), nil
	case OrientationVectorDegreesType:
		var o OrientationVectorDegrees
		err = json.Unmarshal(config.Value, &o)
		if err != nil {
			return nil, err
		}
		return &o, nil
	case OrientationVectorRadiansType:
		var o OrientationVector
		err = json.Unmarshal(config.Value, &o)
		if err != nil {
			return nil, err
		}
		return &o, nil
	case AxisAnglesType:
		var o R4AA
		err = json.Unmarshal(config.Value, &o)
		if err != nil {
			return nil, err
		}
		return &o, nil
	case EulerAnglesType:
		var o EulerAngles
		err = json.Unmarshal(config.Value, &o)
		if err != nil {
			return nil, err
		}
		return &o, nil
	default:
		return nil, fmt.Errorf("orientation type %s not recognized", config.Type)
	}
}
