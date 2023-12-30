package fstr

import "testing"

func TestInterpolate(t *testing.T) {
	tests := []struct {
		name   string
		format string
		data   map[string]interface{}
		want   string
	}{
		{
			name:   "Simple interpolation",
			format: "My name is {name} and I am {age} years old.",
			data: map[string]interface{}{
				"name": "Ziad Mansour",
				"age":  23,
			},
			want: "My name is Ziad Mansour and I am 23 years old.",
		},
		{
			name:   "Formatted interpolation",
			format: "My name is {name} and I am {age} years old. My GPA is {gpa:.2f}.",
			data: map[string]interface{}{
				"name": "Ziad Mansour",
				"age":  23,
				"gpa":  3.14959265359,
			},
			want: "My name is Ziad Mansour and I am 23 years old. My GPA is 3.15.",
		},
		// Add more test cases as needed here.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Interpolate(tt.format, tt.data)
			if err != nil {
				t.Errorf("Interpolate() error = %v", err)
				return
			}
			if got != tt.want {
				t.Errorf("Interpolate() = %v, want %v", got, tt.want)
			}
		})
	}
}
