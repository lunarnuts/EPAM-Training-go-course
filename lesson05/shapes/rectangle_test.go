package shapes

import (
	"testing"
)

func TestRectangle_GetHeight(t *testing.T) {
	type fields struct {
		height float64
		width  float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
		want1  bool
	}{
		{"positive height", fields{height: 1, width: 2}, 1, true},
		{"zero height", fields{height: 0, width: 2}, -1, false},
		{"negative height", fields{height: -1, width: 2}, -1, false},
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
			if got1 != tt.want1 {
				t.Errorf("Rectangle.GetHeight() got1 = %v, want %v", got1, tt.want1)
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
		name   string
		fields fields
		want   float64
		want1  bool
	}{
		{"positive width", fields{height: 1, width: 2}, 2, true},
		{"zero width", fields{height: 0, width: 0}, -1, false},
		{"negative width", fields{height: -1, width: -2}, -1, false},
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
			if got1 != tt.want1 {
				t.Errorf("Rectangle.GetWidth() got1 = %v, want %v", got1, tt.want1)
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
		want   bool
	}{
		{"positive width", fields{height: 1, width: 2}, args{w: 1}, true},
		{"zero width", fields{height: 0, width: 2}, args{w: 0}, false},
		{"negative width", fields{height: -1, width: 2}, args{w: -2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Rectangle{
				height: tt.fields.height,
				width:  tt.fields.width,
			}
			if got := r.SetWidth(tt.args.w); got != tt.want {
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
		want   bool
	}{
		{"positive height", fields{height: 1, width: 2}, args{h: 1}, true},
		{"zero height", fields{height: 0, width: 2}, args{h: 0}, false},
		{"negative height", fields{height: -1, width: 2}, args{h: -2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Rectangle{
				height: tt.fields.height,
				width:  tt.fields.width,
			}
			if got := r.SetHeight(tt.args.h); got != tt.want {
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
			if got := r.String(); got != tt.want {
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
		name   string
		fields fields
		want   float64
		want1  bool
	}{
		{"positive height and width", fields{height: 1, width: 2}, 6, true},
		{"zero height", fields{height: 0, width: 2}, -1, false},
		{"negative height", fields{height: -1, width: 2}, -1, false},
		{"zero width", fields{height: 1, width: 0}, -1, false},
		{"negative width", fields{height: 1, width: -2}, -1, false},
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
			if got1 != tt.want1 {
				t.Errorf("Rectangle.Perimeter() got1 = %v, want %v", got1, tt.want1)
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
		name   string
		fields fields
		want   float64
		want1  bool
	}{
		{"positive height and width", fields{height: 1, width: 2}, 2, true},
		{"zero height", fields{height: 0, width: 2}, -1, false},
		{"negative height", fields{height: -1, width: 2}, -1, false},
		{"zero width", fields{height: 1, width: 0}, -1, false},
		{"negative width", fields{height: 1, width: -2}, -1, false},
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
			if got1 != tt.want1 {
				t.Errorf("Rectangle.Area() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
