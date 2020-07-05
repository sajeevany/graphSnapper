package credentials

import "testing"

func TestAddGrafanaReadUserV1_IsValid(t *testing.T) {
	tests := []struct {
		name string
		user AddGrafanaReadUserV1
		want bool
	}{
		{
			name: "test0 empty API key",
			user: AddGrafanaReadUserV1{
				APIKey:      "",
				Host:        "10.2.3.4",
				Port:        9000,
				Description: "blah",
			},
			want: false,
		},
		{
			name: "test1 empty hostkey",
			user: AddGrafanaReadUserV1{
				APIKey:      "abcdefg",
				Host:        "",
				Port:        9000,
				Description: "blah",
			},
			want: false,
		},
		{
			name: "test2 0 val port",
			user: AddGrafanaReadUserV1{
				APIKey:      "abcdefg",
				Host:        "10.2.3.4",
				Port:        0,
				Description: "blah",
			},
			want: false,
		},
		{
			name: "test3 negative port",
			user: AddGrafanaReadUserV1{
				APIKey:      "abcdefg",
				Host:        "10.2.3.4",
				Port:        -1,
				Description: "blah",
			},
			want: false,
		},
		{
			name: "test4 over max port value",
			user: AddGrafanaReadUserV1{
				APIKey:      "abcdefg",
				Host:        "10.2.3.4",
				Port:        999999999999,
				Description: "blah",
			},
			want: false,
		},
		{
			name: "test5 valid entry",
			user: AddGrafanaReadUserV1{
				APIKey:      "abcdefg",
				Host:        "10.2.3.4",
				Port:        8090,
				Description: "blah",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.user.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddConfluenceServerUserV1_IsValid(t *testing.T) {
	tests := []struct {
		name string
		user AddConfluenceServerUserV1
		want bool
	}{
		{
			name: "test0 empty username key",
			user: AddConfluenceServerUserV1{
				Username: "",
				Password: "qwerty",
				Host:     "10.2.3.4",
				Port:     9000,
			},
			want: false,
		},
		{
			name: "test1 empty password",
			user: AddConfluenceServerUserV1{
				Username: "user",
				Password: "",
				Host:     "10.2.3.4",
				Port:     9000,
			},
			want: false,
		},
		{
			name: "test2 0 val port",
			user: AddConfluenceServerUserV1{
				Username: "asc",
				Password: "qwerty",
				Host:     "10.2.3.4",
				Port:     0,
			},
			want: false,
		},
		{
			name: "test3 negative port",
			user: AddConfluenceServerUserV1{
				Username: "asc",
				Password: "qwerty",
				Host:     "10.2.3.4",
				Port:     -1,
			},
			want: false,
		},
		{
			name: "test4 over max port value",
			user: AddConfluenceServerUserV1{
				Username: "asc",
				Password: "qwerty",
				Host:     "10.2.3.4",
				Port:     999999999999,
			},
			want: false,
		},
		{
			name: "test5 valid entry",
			user: AddConfluenceServerUserV1{
				Username: "asc",
				Password: "qwerty",
				Host:     "10.2.3.4",
				Port:     8090,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.user.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}