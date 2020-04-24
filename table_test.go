package osrm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyTableRequestOptions(t *testing.T) {
	req := TableRequest{}
	assert.Empty(t, req.request().options.encode())
}

func TestNotEmptyTableRequestOptions(t *testing.T) {
	req := TableRequest{
		Sources:      []int{0, 1, 2},
		Destinations: []int{1, 3},
	}
	assert.Equal(t, "destinations=1;3&sources=0;1;2", req.request().options.encode())
}

func TestAnnotations(t *testing.T) {
	req := TableRequest{
		Annotations: AnnotationsDurationDistance,
	}
	assert.Equal(t, "annotations=duration%2Cdistance", req.request().options.encode())
}

func TestFallbackSpeed(t *testing.T) {
	req := TableRequest{
		FallbackSpeed: 5.4,
	}
	assert.Equal(t, "fallback_speed=5.4", req.request().options.encode())
}

func TestFallBackCoordinateInput(t *testing.T) {
	req := TableRequest{
		FallbackSpeed:      11.4,
		FallbackCoordinate: FallbackCoordinateInput,
	}
	assert.Equal(t, "fallback_coordinate=input&fallback_speed=11.4", req.request().options.encode())
}

func TestFallBackCoordinateSnapped(t *testing.T) {
	req := TableRequest{
		FallbackSpeed:      11.4,
		FallbackCoordinate: FallbackCoordinateSnapped,
	}
	assert.Equal(t, "fallback_coordinate=snapped&fallback_speed=11.4", req.request().options.encode())
}

func TestScaleFactor(t *testing.T) {
	req := TableRequest{
		Annotations: AnnotationsDuration,
		ScaleFactor: 2.5,
	}
	assert.Equal(t, "annotations=duration&scale_factor=2.5", req.request().options.encode())
}

func TestSkipWaypoints(t *testing.T) {
	req := TableRequest{
		SkipWaypoints: true,
	}
	assert.Equal(t, "skip_waypoints=true", req.request().options.encode())
}
