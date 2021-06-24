package file

import (
	"reflect"
	"testing"
)

var testData = "Regardless of whether the process provides buffering, the kernel provides buffering, and the system" +
	" provides a buffer (kernel cache) for reading and writing to the disk, and writes the data to the block buffer" +
	" for queueing.When the block buffer reaches a certain amount, the data is written to the disk."

func TestCheckAndCreate(t *testing.T) {
	type args struct {
		dir string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "#1",
			args:    args{"testdata/"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CheckAndCreate(tt.args.dir); (err != nil) != tt.wantErr {
				t.Errorf("CheckAndCreate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIsExist(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Exist",
			args: args{"file.go"},
			want: true,
		},
		{
			name: "NotExist",
			args: args{"files.go"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsExist(tt.args.fileName); got != tt.want {
				t.Errorf("IsExist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsFile(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "IsFile",
			args: args{"file.go"},
			want: true,
		},
		{
			name: "FileNotExist",
			args: args{"files.go"},
		},
		{
			name: "IsDir",
			args: args{"file"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsFile(tt.args.fileName); got != tt.want {
				t.Errorf("IsFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateFile(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Normal",
			args:    args{"testdata/data.txt"},
			wantErr: false,
		},
		{
			name:    "Dir",
			args:    args{"testdata1/"},
			wantErr: true,
		},
		{
			name: "CurDir",
			args: args{"./test.dat"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := CreateFile(tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestReadAll(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name:    "#1",
			args:    args{"testdata/data_all.txt"},
			want:    []byte(testData),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !IsExist(tt.args.fileName) {
				err := WriteFile4OS(tt.args.fileName, testData)
				if err != nil {
					return
				}
			}

			got, err := ReadAll(tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadByBytes(t *testing.T) {
	type args struct {
		fileName string
		bits     int
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name:    "#1",
			args:    args{"testdata/data_bytes.txt", 1024},
			want:    []byte(testData),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !IsExist(tt.args.fileName) {
				err := WriteFile4OS(tt.args.fileName, testData)
				if err != nil {
					return
				}
			}

			got, err := ReadByBytes(tt.args.fileName, tt.args.bits)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadByBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadByBytes() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadFile(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name:    "#1",
			args:    args{"testdata/data_file.txt"},
			want:    []byte(testData),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !IsExist(tt.args.fileName) {
				err := WriteFile4OS(tt.args.fileName, testData)
				if err != nil {
					return
				}
			}

			got, err := ReadFile(tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWriteFile4IOUtil(t *testing.T) {
	type args struct {
		fileName string
		data     []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "#1",
			args: args{
				fileName: "testdata/data.txt",
				data:     []byte(testData),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WriteFile4IOUtil(tt.args.fileName, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("WriteFile4IOUtil() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWriteFile4BufIO(t *testing.T) {
	type args struct {
		fileName string
		data     string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "#1",
			args: args{
				fileName: "testdata/data.txt",
				data:     testData,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WriteFile4BufIO(tt.args.fileName, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("WriteFile4BufIO() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWriteFile(t *testing.T) {
	type args struct {
		fileName string
		data     string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "#1",
			args: args{
				fileName: "testdata/data.txt",
				data:     testData,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WriteFile(tt.args.fileName, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("WriteFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWriteFile4OS(t *testing.T) {
	type args struct {
		fileName string
		data     string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "#1",
			args: args{
				fileName: "testdata/data.txt",
				data:     testData,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WriteFile4OS(tt.args.fileName, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("WriteFile4OS() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func BenchmarkWriteFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = WriteFile("testdata/data_file.txt", testData)
	}
}

func BenchmarkWriteFile4BufIO(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = WriteFile4BufIO("testdata/data_buf.txt", testData)
	}
}

func BenchmarkWriteFile4OS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = WriteFile4OS("testdata/data_os.txt", testData)
	}
}

func BenchmarkWriteFile4IOUtil(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = WriteFile4IOUtil("testdata/data_io.txt", []byte(testData))
	}
}

func BenchmarkReadAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ReadAll("testdata/data.txt")
	}
}

func BenchmarkReadByBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = ReadByBytes("testdata/data_bytes.txt", 1024)
	}
}

func BenchmarkReadFile(b *testing.B) {
	_, _ = ReadFile("testdata/data_file.txt")
}
