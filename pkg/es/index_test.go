package es

import (
	"testing"
)

func TestEsClient(t *testing.T) {
	type args struct {
		indexName string
		address   []string
		username  string
		password  string
	}
	arg := args{
		indexName: "test",
		address:   []string{"http://10.89.64.76:9200"},
		username:  "elastic",
		password:  "P@ssw0rd",
	}

	t.Run("createIndex", func(t *testing.T) {
		es, err := NewEsClient(arg.address, arg.username, arg.password)
		if err != nil {
			t.Fatal(err)
		}
		err = es.CreateIndex(arg.indexName)
		if err != nil {
			t.Errorf("CreateIndex() error = %v", err)
		}
	})

	t.Run("getIndex", func(t *testing.T) {
		es, err := NewEsClient(arg.address, arg.username, arg.password)
		if err != nil {
			t.Fatal(err)
		}
		err = es.GetIndex(arg.indexName)
		if err != nil {
			t.Errorf("GetIndex() error = %v", err)
		}
	})

	t.Run("deleteIndex", func(t *testing.T) {
		es, err := NewEsClient(arg.address, arg.username, arg.password)
		if err != nil {
			t.Fatal(err)
		}
		err = es.DeleteIndex(arg.indexName)
		if err != nil {
			t.Errorf("DeleteIndex() error = %v", err)
		}
	})

}
