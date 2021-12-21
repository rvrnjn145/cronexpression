package cronExpression

import "testing"

func TestClientImpl_ToDescription(t *testing.T) {
	type fields struct {
		logger Logger
		parser Parser
	}
	field := fields{parser: &ParserImpl{}, logger: &loggerImpl{LoggerFlag: true}}
	type args struct {
		expr string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{

		{

			name:    "test0",
			fields:  field,
			args:    args{expr: "*/15 0 1,15 * 1-5 /usr/bin/find -e foo"},
			want:    "1 2",
			wantErr: false,
		},
		{

			name:    "test1",
			fields:  field,
			args:    args{expr: "1,2 * * * * /folder/file"},
			want:    "1 2",
			wantErr: false,
		},
		{
			name:    "test2",
			fields:  field,
			args:    args{expr: "61,2 * * * * /folder/file"},
			want:    "",
			wantErr: true,
		},
		{
			name:    "test3",
			fields:  field,
			args:    args{expr: "1-4 * * * * /folder/file"},
			want:    "1 2 3 4",
			wantErr: false,
		},
		{
			name:    "test4",
			fields:  field,
			args:    args{expr: "9-3 * * * * /folder/file"},
			want:    "",
			wantErr: true,
		},
		{
			name:    "test5",
			fields:  field,
			args:    args{expr: "61/2 * * * * /folder/file"},
			want:    "",
			wantErr: true,
		},
		{
			name:    "test6",
			fields:  field,
			args:    args{expr: "* * * * * /folder/file"},
			want:    "0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59",
			wantErr: false,
		},
		{
			name:    "test7",
			fields:  field,
			args:    args{expr: "*/15 * * * * /folder/file"},
			want:    "0 15 30 45",
			wantErr: false,
		},
		{
			name:    "test8",
			fields:  field,
			args:    args{expr: "60/15 * * * * /folder/file"},
			want:    "",
			wantErr: true,
		},
		{
			name:    "test9",
			fields:  field,
			args:    args{expr: "0 2,3 * * * /folder/file"},
			want:    "0",
			wantErr: false,
		},
		{
			name:    "test10",
			fields:  field,
			args:    args{expr: "0 2 2/7 * * /folder/file"},
			want:    "0",
			wantErr: false,
		},
		{
			name:    "test11",
			fields:  field,
			args:    args{expr: "0 2 2/7 2,3 * /folder/file"},
			want:    "0",
			wantErr: false,
		},
		{
			name:    "test12",
			fields:  field,
			args:    args{expr: "0 2 2/7 2,3 3-4 /folder/file"},
			want:    "0",
			wantErr: false,
		},
		{
			name:    "test13",
			fields:  field,
			args:    args{expr: "0 2 ? 2,3 3-4 /folder/file"},
			want:    "0",
			wantErr: false,
		},
		{
			name:    "test14",
			fields:  field,
			args:    args{expr: "0 2 1 2,3 ? /folder/file"},
			want:    "0",
			wantErr: false,
		},

		{
			name:    "test15",
			fields:  field,
			args:    args{expr: "0 12 1 * ? /folder/file"},
			want:    "0",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &ClientImpl{
				logger: tt.fields.logger,
				parser: tt.fields.parser,
			}
			got, err := e.CronExpressionDescriptor(tt.args.expr)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClientImpl.ToDescription() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ClientImpl.ToDescription() = %v, want %v", got, tt.want)
			}
		})
	}
}
