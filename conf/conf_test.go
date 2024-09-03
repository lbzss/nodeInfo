package conf

import (
	"testing"
)

func TestConfig_Load(t *testing.T) {
	type fields struct {
		ElasticsearchAddress []string
		UserName             string
		Password             string
		IndexPrefix          string
		ServerAddress        string
		ServerPort           int64
	}
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Load config", fields: fields{
			ElasticsearchAddress: []string{"http://1.2.3.4:9200"},
			UserName:             "foo",
			Password:             "bar",
			IndexPrefix:          "node_info",
			ServerAddress:        "1.2.3.4",
			ServerPort:           8080,
		}, args: args{path: "./config.yaml"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{}
			if err := c.Load(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantErr)
			}
			t.Logf("%+v", c)
		})
	}
}
