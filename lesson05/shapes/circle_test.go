package shapes

import (
	"errors"
	"fmt"
	"math"
	"testing"
)

func TestCircle_SetRadius(t *testing.T) {
	type fields struct {
		radius float64
	}
	type args struct {
		r float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   error
	}{
		{"positive radius", fields{radius: 1}, args{r: 1}, nil},
		{"zero radius", fields{radius: 0}, args{r: -1}, fmt.Errorf("radius needs to be positive and non-zero, got: %.3f", 0.0)},
		{"negative radius", fields{radius: -1}, args{r: -1}, fmt.Errorf("radius needs to be positive and non-zero, got: %.3f", -1.0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Circle{
				radius: tt.fields.radius,
			}
			if got := c.SetRadius(tt.args.r); got == nil && tt.want != nil || got != nil && tt.want == nil {
				t.Errorf("Circle.SetRadius() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCircle_GetRadius(t *testing.T) {
	type fields struct {
		radius float64
	}
	tests := []struct {
		name    string
		fields  fields
		want    float64
		wantErr error
	}{
		{"positive radius", fields{radius: 1}, 1, nil},
		{"zero radius", fields{radius: 0}, -1, fmt.Errorf("radius needs to be positive and non-zero, got: %.3f", 0.0)},
		{"negative radius", fields{radius: -1}, -1, fmt.Errorf("radius needs to be positive and non-zero, got: %.3f", -1.0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Circle{
				radius: tt.fields.radius,
			}
			got, got1 := c.GetRadius()
			if got != tt.want {
				t.Errorf("Circle.GetRadius() got = %v, want %v", got, tt.want)
			}
			if !errors.Is(got1, tt.wantErr) {
				t.Errorf("Circle.GetRadius() gotErr = %v, want %v", got1, tt.wantErr)
			}
		})
	}
}

func TestCircle_String(t *testing.T) {
	type fields struct {
		radius float64
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"positive radius", fields{radius: 1}, "\nCircle: radius 1.00"},
		{"zero radius", fields{radius: 0}, "radius needs to be positive and non-zero, got: 0.000"},
		{"negative radius", fields{radius: -1}, "radius needs to be positive and non-zero, got: -1.000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Circle{
				radius: tt.fields.radius,
			}
			if got := c.String(); got != tt.want {
				t.Errorf("Circle.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCircle_Perimeter(t *testing.T) {
	type fields struct {
		radius float64
	}
	tests := []struct {
		name    string
		fields  fields
		want    float64
		wantErr error
	}{
		{"positive radius", fields{radius: 1}, 2 * math.Pi, nil},
		{"zero radius", fields{radius: 0}, -1, fmt.Errorf("radius needs to be positive and non-zero, got: %.3f", 0.0)},
		{"negative radius", fields{radius: -1}, -1, fmt.Errorf("radius needs to be positive and non-zero, got: %.3f", -1.0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Circle{
				radius: tt.fields.radius,
			}
			got, got1 := c.Perimeter()
			if got != tt.want {
				t.Errorf("Circle.Perimeter() got = %v, want %v", got, tt.want)
			}
			if got1 == nil && tt.wantErr != nil || got1 != nil && tt.wantErr == nil {
				t.Errorf("Circle.Perimeter() got1 = %v, want %v", got1, tt.wantErr)
			}
		})
	}
}

func TestCircle_Area(t *testing.T) {
	type fields struct {
		radius float64
	}
	tests := []struct {
		name    string
		fields  fields
		want    float64
		wantErr error
	}{
		{"positive radius", fields{radius: 1}, math.Pi, nil},
		{"zero radius", fields{radius: 0}, -1, fmt.Errorf("radius needs to be positive and non-zero, got: %.3f", 0.0)},
		{"negative radius", fields{radius: -1}, -1, fmt.Errorf("radius needs to be positive and non-zero, got: %.3f", -1.0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Circle{
				radius: tt.fields.radius,
			}
			got, got1 := c.Area()
			if got != tt.want {
				t.Errorf("Circle.Area() got = %v, want %v", got, tt.want)
			}
			if got1 == nil && tt.wantErr != nil || got1 != nil && tt.wantErr == nil {
				t.Errorf("Circle.Area() got1 = %v, want %v", got1, tt.wantErr)
			}
		})
	}
}
