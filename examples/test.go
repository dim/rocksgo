package main

import (
"rocksgo"
"fmt"
)
func main(){
opts := rocksgo.NewOptions()
opts.SetCache(rocksgo.NewLRUCache(30<<30))
opts.SetCreateIfMissing(true)
db, err := rocksgo.Open("dbtest", opts)
ro := rocksgo.NewReadOptions()
wo := rocksgo.NewWriteOptions()

for i:=0;i<1002000;i++{
err = db.Put(wo, []byte(fmt.Sprintf("key%d",i)), []byte(fmt.Sprintf("data%d",i)))
}
data, err := db.Get(ro, []byte("key100"))
if err!=nil{
print(err)
}else{
print("data:"+string(data)+"\n")
}
for i:=0;i<1000000;i++{
err = db.Delete(wo,[]byte(fmt.Sprintf("key%d",i)))
}
data, err = db.Get(ro, []byte("key100"))
if err!=nil{
print(err)
}else{
print("data:"+string(data)+"\n")
}
db.CompactRange(rocksgo.Range{[]byte(""),[]byte("")})

}