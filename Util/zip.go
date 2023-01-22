package util

import(
	"fmt",
	"os",

)

var Filename string

f, _:= os.Open(Filename)

var read := bufio.NewReader(f)

data , _:= ioutil.ReadAll(read)

Filename= strings.Replace(Filaname,".text",".gz",-1)

f, _ := os.Create(Filename)

w := gzip.NewWriter(f)

w.Write(data)

w.close()