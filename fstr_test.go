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
		{
			name:   "Comma formatting",
			format: "{name} - {age} - {balance:,} - {gpa:.4f} - {total:,.3f} - {sum:,}",
			data: map[string]interface{}{
				"name":    "Ziad Mansour",
				"age":     23,
				"gpa":     3.1657,
				"balance": 123456789.111,
				"sum":     123456789,
				"total":   123456789.9787968,
			},
			want: "Ziad Mansour - 23 - 123,456,789 - 3.1657 - 123,456,789.979 - 123,456,789",
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

func TestEval(t *testing.T) {
	tests := []struct {
		name   string
		format string
		data   map[string]interface{}
		want   string
	}{
		{
			name:   "All cases",
			format: "{name} - {age} - {balance:,} - {gpa:.4f} - {total:,.3f} - {sum:,}",
			data: map[string]interface{}{
				"name":    "Ziad Mansour",
				"age":     23,
				"gpa":     3.1657,
				"balance": 123456789.111,
				"sum":     123456789,
				"total":   123456789.9787968,
			},
			want: "Ziad Mansour - 23 - 123,456,789 - 3.1657 - 123,456,789.979 - 123,456,789",
		},
		{
			name:   "Advanced cases",
			format: "{name=} - {age=} - {balance=:,} - {gpa=:.4f} - {total=:,.3f} - {sum=:,.2f}",
			data: map[string]interface{}{
				"name":    "Ziad Mansour",
				"age":     23,
				"gpa":     3.1657,
				"balance": 123456789.111,
				"sum":     123456789,
				"total":   123456789.9787968,
			},
			want: "name=Ziad Mansour - age=23 - balance=123,456,789 - gpa=3.1657 - total=123,456,789.979 - sum=123,456,789.00",
		},
		// Add more test cases as needed here.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Eval(tt.format, tt.data)
			if got != tt.want {
				t.Errorf("Eval() = %v, want %v", got, tt.want)
			}
		})
	}
}
