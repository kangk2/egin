package db

import (
	"reflect"
	"testing"
)

func TestAttrToQuery(t *testing.T) {
	type args struct {
		attr Attr
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 []interface{}
	}{
		{
			name: "Attr转换",
			args: args{
				attr: Attr{
					OrderBy: "id desc",
				},
			},
			want:  "order by ?",
			want1: []interface{}{"id desc"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := AttrToQuery(tt.args.attr)
			if got != tt.want {
				t.Errorf("AttrToQuery() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("AttrToQuery() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFilterToQuery(t *testing.T) {
	type args struct {
		filter Filter
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 []interface{}
	}{
		{
			name: "where转换",
			args: args{
				filter: Filter{
					"name": "daodao",
				},
			},
			want:  "where `name` = ?",
			want1: []interface{}{"daodao"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := FilterToQuery(tt.args.filter)
			if got != tt.want {
				t.Errorf("FilterToQuery() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("FilterToQuery() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestInsertRecordToQuery(t *testing.T) {
	type args struct {
		record Record
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 []interface{}
	}{
		{
			name: "Insert转换",
			args: args{
				record: Record{
					"name": "daodao",
				},
			},
			want:  "insert into %s (`name`) values (?)",
			want1: []interface{}{"daodao"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := InsertRecordToQuery(tt.args.record)
			if got != tt.want {
				t.Errorf("InsertRecordToQuery() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("InsertRecordToQuery() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
