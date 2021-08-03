package shapes

import (
	"errors"
	"testing"
)

func TestRectangle_GetHeight(t *testing.T) {
	type fields struct {
		height float64
		width  float64
	}
	tests := []struct {
		name    string
		fields  fields
		want    float64
		wantErr error
	}{
		{"positive height", fields{height: 1, width: 2}, 1, nil},
		{"zero height", fields{height: 0, width: 2}, -1, errors.New("height needs to be positive and non-zero, got: 0.000")},
		{"negative height", fields{height: -1, width: 2}, -1, errors.New("height needs to be positive and non-zero, got: -1.000")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rectangle{
				height: tt.fields.height,
				width:  tt.fields.width,
			}
			got, got1 := r.GetHeight()
			if got != tt.want {
				t.Errorf("Rectangle.GetHeight() got = %v, want %v", got, tt.want)
			}
			if got1 == nil && got1 != tt.wantErr {
				t.Errorf("Rectangle.GetHeight() got1 = %v, want %v", got1, tt.wantErr)
			}
			if got1 != nil && got1.Error() != tt.wantErr.Error() {
				t.Errorf("Rectangle.GetHeight() got1 = %v, want %v", got1, tt.wantErr)
			}
		})
	}
}

func TestRectangle_GetWidth(t *testing.T) {
	type fields struct {
		height float64
		width  float64
	}
	tests := []struct {
		name    string
		fields  fields
		want    float64
		wantErr error
	}{
		{"positive width", fields{height: 1, width: 2}, 2, nil},
		{"zero width", fields{height: 0, width: 0}, -1, errors.New("width needs to be positive and non-zero, got: 0.000")},
		{"negative width", fields{height: -1, width: -2}, -1, errors.New("width needs to be positive and non-zero, got: -2.000")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rectangle{
				height: tt.fields.height,
				width:  tt.fields.width,
			}
			got, got1 := r.GetWidth()
			if got != tt.want {
				t.Errorf("Rectangle.GetWidth() got = %v, want %v", got, tt.want)
			}
			if got1 == nil && got1 != tt.wantErr {
				t.Errorf("Rectangle.GetWidth() got1 = %v, want %v", got1, tt.wantErr)
			}
			if got1 != nil && got1.Error() != tt.wantErr.Error() {
				t.Errorf("Rectangle.GetWidth() got1 = %v, want %v", got1, tt.wantErr)
			}
		})
	}
}

func TestRectangle_SetWidth(t *testing.T) {
	type fields struct {
		height float64
		width  float64
	}
	type args struct {
		w float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   error
	}{
		{"positive width", fields{height: 1, width: 2}, args{w: 1}, nil},
		{"zero width", fields{height: 0, width: 2}, args{w: 0}, errors.New("width needs to be positive and non-zero, got: 2.000")},
		{"negative width", fields{height: -1, width: 2}, args{w: -2}, errors.New("width needs to be positive and non-zero, got: 2.000")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Rectangle{
				height: tt.fields.height,
				width:  tt.fields.width,
			}
			got := r.SetWidth(tt.args.w)
			if got == nil && got != tt.want {
				t.Errorf("Rectangle.SetWidth() = %v, want %v", got, tt.want)
			}
			if got != nil && got.Error() != tt.want.Error() {
				t.Errorf("Rectangle.SetWidth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRectangle_SetHeight(t *testing.T) {
	type fields struct {
		height float64
		width  float64
	}
	type args struct {
		h float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   error
	}{
		{"positive height", fields{height: 1, width: 2}, args{h: 1}, nil},
		{"zero height", fields{height: 0, width: 2}, args{h: 0}, errors.New("height needs to be positive and non-zero, got: 0.000")},
		{"negative height", fields{height: -1, width: 2}, args{h: -2}, errors.New("height needs to be positive and non-zero, got: -1.000")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Rectangle{
				height: tt.fields.height,
				width:  tt.fields.width,
			}
			got := r.SetHeight(tt.args.h)
			if got == nil && got != tt.want {
				t.Errorf("Rectangle.SetHeight() = %v, want %v", got, tt.want)
			}
			if got != nil && got.Error() != tt.want.Error() {
				t.Errorf("Rectangle.SetHeight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRectangle_String(t *testing.T) {
	type fields struct {
		height float64
		width  float64
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"positive height", fields{height: 1, width: 2}, "\nRectangle with height 1.00 and width 2.00"},
		{"zero height", fields{height: 0, width: 2}, "height needs to be positive and non-zero, got: 0.000"},
		{"negative height", fields{height: -1, width: 2}, "height needs to be positive and non-zero, got: -1.000"},
		{"zero width", fields{height: 1, width: 0}, "width needs to be positive and non-zero, got: 0.000"},
		{"negative width", fields{height: 1, width: -2}, "width needs to be positive and non-zero, got: -2.000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rectangle{
				height: tt.fields.height,
				width:  tt.fields.width,
			}
			got := r.String()
			if got != tt.want {
				t.Errorf("Rectangle.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRectangle_Perimeter(t *testing.T) {
	type fields struct {
		height float64
		width  float64
	}
	tests := []struct {
		name    string
		fields  fields
		want    float64
		wantErr error
	}{
		{"positive height and width", fields{height: 1, width: 2}, 6, nil},
		{"zero height", fields{height: 0, width: 2}, -1, errors.New("height needs to be positive and non-zero, got: 0.000")},
		{"negative height", fields{height: -1, width: 2}, -1, errors.New("height needs to be positive and non-zero, got: -1.000")},
		{"zero width", fields{height: 1, width: 0}, -1, errors.New("width needs to be positive and non-zero, got: 0.000")},
		{"negative width", fields{height: 1, width: -2}, -1, errors.New("width needs to be positive and non-zero, got: -2.000")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rectangle{
				height: tt.fields.height,
				width:  tt.fields.width,
			}
			got, got1 := r.Perimeter()
			if got != tt.want {
				t.Errorf("Rectangle.Perimeter() got = %v, want %v", got, tt.want)
			}
			if got1 == nil && got1 != tt.wantErr {
				t.Errorf("Rectangle.Perimeter() got1 = %v, want %v", got1, tt.wantErr)
			}
			if got1 != nil && got1.Error() != tt.wantErr.Error() {
				t.Errorf("Rectangle.Perimeter() got1 = %v, want %v", got1, tt.wantErr)
			}
		})
	}
}

func TestRectangle_Area(t *testing.T) {
	type fields struct {
		height float64
		width  float64
	}
	tests := []struct {
		name    string
		fields  fields
		want    float64
		wantErr error
	}{
		{"positive height and width", fields{height: 1, width: 2}, 2, nil},
		{"zero height", fields{height: 0, width: 2}, -1, errors.New("height needs to be positive and non-zero, got: 0.000")},
		{"negative height", fields{height: -1, width: 2}, -1, errors.New("height needs to be positive and non-zero, got: -1.000")},
		{"zero width", fields{height: 1, width: 0}, -1, errors.New("width needs to be positive and non-zero, got: 0.000")},
		{"negative width", fields{height: 1, width: -2}, -1, errors.New("width needs to be positive and non-zero, got: -2.000")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rectangle{
				height: tt.fields.height,
				width:  tt.fields.width,
			}
			got, got1 := r.Area()
			if got != tt.want {
				t.Errorf("Rectangle.Area() got = %v, want %v", got, tt.want)
			}
			if got1 == nil && got1 != tt.wantErr {
				t.Errorf("Rectangle.Area() got1 = %v, want %v", got1, tt.wantErr)
			}
			if got1 != nil && got1.Error() != tt.wantErr.Error() {
				t.Errorf("Rectangle.Area() got1 = %v, want %v", got1, tt.wantErr)
			}
		})
	}
}
