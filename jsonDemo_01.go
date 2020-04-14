package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)
//example demeonstrate the serialization and de-serialization
type OrderInfo struct {
	OrderName  string `json:"orderName"`
}

type ForwardingOrder struct {
	OrderName string `json:"orderName"`
	OrderId int `json:"orderId,omitempty"` //omitempty field will help to avoid forming the xml data if value not there
	OrderBy *string `json:"orderBy"` //if we give pointer field or normal field,it is same in json unlike xml
	Label      string `json:",omitempty"`
	OrderDetails []OrderInfo

}

type ForwardingOrderReceiver struct {
	ORDERName string // if it is exported, it is not case sensitive
	OrderId int

}

func main() {
	orderDet1:=OrderInfo{OrderName:"A"}
	orderDet2:=OrderInfo{OrderName:"A"}
	orderDet3:=OrderInfo{OrderName:"A"}
	order:=ForwardingOrder{
		OrderName: "test",
		OrderDetails: []OrderInfo{orderDet1,orderDet2,orderDet3},
		Label:"TESTING",

	}
	// serialization
	out,_:=json.MarshalIndent(order,""," ")
	fmt.Println(string(out))
    // de-serialization
	fwOrder:=ForwardingOrderReceiver{} //An exported field with a tag of "Foo" (see the Go spec for more on struct tags), An exported field named "Foo", or An exported field named "FOO" or "FoO" or some other case-insensitive match of "Foo".
	json.Unmarshal(out,&fwOrder)
	fmt.Println(fwOrder) //{test 0}

    exampleOfUsingEncoding(order)

}

/**
It really depends on what your input is. If you look at the implementation of the Decode method of json.Decoder, it buffers the entire JSON value in memory before unmarshalling it into a Go value. So in most cases it won't be any more memory efficient (although this could easily change in a future version of the language).

So a better rule of thumb is this:

Use json.Decoder if your data is coming from an io.Reader stream, or you need to decode multiple values from a stream of data.
Use json.Unmarshal if you already have the JSON data in memory.
For the case of reading from an HTTP request, I'd pick json.Decoder since you're obviously reading from a stream.
 */

func exampleOfUsingEncoding(order ForwardingOrder) {
	//whenever we want to convert the data to json and print, we can use the encoding
	writer := os.Stdout
	json.NewEncoder(writer).Encode(order)

	//decoding
	fwOrder:=ForwardingOrderReceiver{}
	//jsonInput
	//if we have data already then
	jsonInput:="{\"orderName\":\"test\",\"orderBy\":null,\"Label\":\"TESTING\",\"OrderDetails\":[{\"orderName\":\"A\"},{\"orderName\":\"A\"},{\"orderName\":\"A\"}]}"
	reader := strings.NewReader(jsonInput)
	dec := json.NewDecoder(reader)
	dec.Decode(&fwOrder)
	fmt.Println(fwOrder)


}

