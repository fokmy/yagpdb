// Code generated by "esc -o templates.go -pkg docs templates/"; DO NOT EDIT.

package docs

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// FS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func FS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// Dir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func Dir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// FSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func FSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		_ = f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// FSMustByte is the same as FSByte, but panics if name is not present.
func FSMustByte(useLocal bool, name string) []byte {
	b, err := FSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// FSString is the string version of FSByte.
func FSString(useLocal bool, name string) (string, error) {
	b, err := FSByte(useLocal, name)
	return string(b), err
}

// FSMustString is the string version of FSMustByte.
func FSMustString(useLocal bool, name string) string {
	return string(FSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/templates/ads.md": {
		local:   "templates/ads.md",
		size:    691,
		modtime: 1518879672,
		compressed: `
H4sIAAAAAAAC/zxSQW7bMBC86xUDtEBsVHYcI4GBnnrooe65H1hTK4k1xRW4Kyvs6wvSiXngYTkznBny
jMTRIIvBRkaQlRMC9wYnKXKCCUYOM2bK6CVVlBM1SI+0xOjjgEzD3F32OPfIsmClaIVXlW0kg85ibaFG
9MwBfWIuCCfRyBkmbsHeRk6QCJ7IB2z+SiRNp9fTj6EM9k6mLaQifnp1kjpsfhfM6fX05XA4vGz3TfOL
Ez8pVCZGx0Y+aCE8XHxvsMOfmiFakoCZIpd9YMXApqAkS+zwcrjWKW6eV8XMCSvz9UFfUo03J+8YXtH7
d+5Ahpevzx3lFmd0giiGkW4Mihma1XiqLVIHNTJtoYsbQQoXvLvCxiTLMGoLHzEHcozM1kIF3p4UVHyr
lYKLTJR1/+lI/b9q5Ph2eD8e3xqU9a0eJSaVeFfQiUIouAs7WpQ/dMNKWXHz6i+BsblDjcw7CiFjFvXm
JXK3rReeITFkUAiygj6Q8FMprMQrwXc3TnlHMUoun+TmO5bNtBh32+fB958GHz3dxfyuTzSxVpmYbSzc
4K9c33DfNM3/AAAA///DWIAXswIAAA==
`,
	},

	"/templates/docs.ghtml": {
		local:   "templates/docs.ghtml",
		size:    994,
		modtime: 1518359152,
		compressed: `
H4sIAAAAAAAC/5xSTW+cPBA+v/sr5nUviVQvSs+ES3KpVFU5VOqxmrUHsGJmqG1YrSz+ewULCVn11AtY
9swz83zkbKl2TKCsmKgZRzVNh9K76gAAUCK0gepH9UlVpQPjMcZHVSPUqB3XMv/rs6rKwlXwLGboiBMm
Jwxl7JG3Do8n8rB89RkDO25U9fPrS1nMVdWH2hoBQ5AF9fpa4LrN4LcixhEYRx3JCFvtaSSvqsN/c1nO
AbkhOD6LecGG4jQt7QvExuztAneQ2jt+VSvlYlakyHkI/vdA4QJ3Xs4U4PgdO7qfJlXlvJyn6W3DBbGY
Z6ybENt1evm/1lAcb5cGrVdyxeCrw7V56zvc2NNjQ7M/OSfqeo+JQJn+V0toFRxn46wbNz6LhFfo3a0R
r32jH76o3ca791PUBr2XIcH78d2zD9r9aF0EFyG1BHZv/uflCk0a0IMRrl0zhGssXASWBFaYoKVAR/gm
jWNAthDJk0lwkSFApDBSAMcL1ND3FCC4pk1gJDAFuKvFWwoQL91JPAhDJyfn6X5nhHXjJu5y3H7/plL7
sD3PPuhZdQpLCp6GEIjTmre3ULQPf9e4w/Bq5cz6JPaym5DzHNkn4USc9qF9J3LD5DYJtUiicM3CGqI/
AQAA//+stVcF4gMAAA==
`,
	},

	"/templates/helping-out.md": {
		local:   "templates/helping-out.md",
		size:    3415,
		modtime: 1518879672,
		compressed: `
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
`,
	},

	"/templates/index.md": {
		local:   "templates/index.md",
		size:    1674,
		modtime: 1518879672,
		compressed: `
H4sIAAAAAAAC/4xUTW/bOBC9+1cM0kMTwFViu3Hc3JK4+ejmkKbd7bZBgVDSSGIsceghaVkt+t8XQzlu
mu0CexPEmcf3Zt7jx0o70A58hZBTFho0XnlNBgpi+HxycTM/TQaDKy9FWWBG4+sOVsgdaJNRY2v0CMrk
/b+WeAHagGUqGZ0bgiPQBXQUoFXGgyeosLZAwUPL2mtTysVuCLZG5YSFeemhQqe98ij1OYEjIdHD9AWF
NnlkrU1B3PScOwovGaEmWghuQTyUGgNXwJhR06DJI8YDaRO7XbCW2INDXiEPoa10VkGmDKQIBQWTg0pp
hclg8OIFzMkoYbzl8igplwMEYrhLQxeRVf7KWZXh19190bevcrcXewQ8SoLghOWN8oxkhnCqfUb650em
XDWEt75CxtDEGV9rj3J0PBjAq8fOY6i8t+54f79t28T2P5OMmv1OlTZPpXQDeQx396PLkB9dfhrfTu2H
i2BP518W79bLxfm35ccP5YPh/PT+6+737058kMHOktO+N7Gm3PnxY+8J3JlyVYScLz+P31OKBY8ulrfV
YvKlu5mer/6cVKfpt7Yop7+HFIVPYR+1CubB+mA2Ph+pPDsbqdl0mh0dKvV2djYZ43hyMBu9nqmDbHw4
mxw+w8YNyFPg7dzg7v76ejq9Pvljnb95v/xU8uX66O/FyB++cYu381vq8N1fzwDrTfMWcHBGxqvMQ4O/
ejuGyBADUy3OhKaDuXYZcb5xWO+jWxTTyfKvnAvoHv0k5q3USg50PIBW+yripuSHMZO4VhI6UCCGFlMw
+sBGmhw16Cv5qvUCYedc6RpzSJJkp8+BUG10WXm5BgUBQWVMzoGCNJQxLlvKJzdX0KhOorBhZZnSGhs3
FKtL6YRzsIp9F/XpDMFXym+pLQOyRge7ugBluj15Rf4FpY2cQSbpF4KuolDn/5VR2K0pUx5z2Bx7spAq
3osBqTFuZWGoleAGWYp2yWBwUntkSe8K6264DaKqHQFZFAr9zIF62LsL7S9DClaVGAe/2cLX3ce0ldpX
IY1BeyCj3NHro03i9iCu+RyVDywCyhKdvFDbTffzV0+O4h0KDLZQbPqIQTeWaYXyLLu+wgCutYvmIYM/
laAW34OTJ05Bg84J8d+PkPiZZv2rZkZL/0/nk8Ek8LFCqJX3coNB2bVlLJAZ8yE4bTKEK/BC0JN0lhjX
Y0qXwOCfAAAA//9GoFGYigYAAA==
`,
	},

	"/templates/quickstart.md": {
		local:   "templates/quickstart.md",
		size:    1290,
		modtime: 1518879672,
		compressed: `
H4sIAAAAAAAC/3xUTW/cNhC961e8ugfHgCy3TQsYvqUxYAQt2gBOD0HRA0WOJHYpjjocriL/+oJcbbYB
itwIzcd782aemg+TT1jMSFh9CJgoLNg445/s7SFsSKTIC3Qi9Kxd0zyT5uWhab7v8MRQxqS6pIe7u82M
i+u7T9vLHZofOvwea1VeFpJb8eOkt5YlkrSwwdsDOOLqVx59xB/JxxGPPlkWd9W87vDs5yVsCDWca3jj
LOccWCFHUb0JqUWfFY7jtWJlka0tqddCBWDPv05YqU9eCZHXFivBMSIrRlKYuIGHE8AXjUufENAThJwX
skoOHC3VyJFgsk4s/oUcPr55ev/4cxHk6GmtoyeSI0m60OmaHzu8OdfsJUWxUmY5RrJanpWJsZZz1K75
6ata5uhIcJUolGKzg161SCcJ98CFT13vaqKel1qI4TdWesBHzrAmgmPYYJz7nOKrPJjMkTCbWO5l77aQ
zD4lzzF1TfPnh4mE4BNMSAyDo3fE0Kws3gTMxhH6Db/4OL7lXsy39/ffvcY6eTvBcpVr5qRlHwV7IKNZ
KLUofb/569Xna+Osuevp7uXp/j0dlsfbw+Gmac4DRF6LoIMfsxDWySj9d/QOb+sJ6iScx6lCHY14zicz
JJjoQJ+WwEIwCF41UFlNolO7i3jDQFJGf3eSSHKEj8r1qnxKuZD/m/1pgykvC4uexTM9H6mFCTpVHksg
k+ptqmwFbh+AsyJpHob6KueRKAwXYr3XFiZh3qB+Lh2j+wJPyczXaQ/6hOBnr+RaRAbHurCRVIvPFuMd
BhZo+TW88kOZRKcSenc9I3B148yRtuKwiVMtK4gry6G8Odbimw7PvDuz+Gw27nxIl48Gg0kKobRwTASW
4k6lEGqiMoTMaRjHNmHwkrTS2m3VU4GU7Aj/K3ILO5loy06FzkUjx2gqfm9iJHfT/RsAAP//TTPcCwoF
AAA=
`,
	},

	"/templates/templates.md": {
		local:   "templates/templates.md",
		size:    4489,
		modtime: 1518359152,
		compressed: `
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
mUz+FwAA//99wtxUiREAAA==
`,
	},

	"/": {
		isDir: true,
		local: "",
	},

	"/templates": {
		isDir: true,
		local: "templates",
	},
}
