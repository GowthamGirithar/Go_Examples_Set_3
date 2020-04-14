package main

import (
	"encoding/xml"
	"fmt"
)

type OrderDetail struct {
	XMLName xml.Name `xml:"orderDetail"` //If a field is called XMLName and its type is xml.Name
	OrderName  string `xml:"orderName"`
}

type Order struct {
	XMLName xml.Name `xml:"orderinfo"`  // field name have to XMLName , type has to be xml.Name to get this is xml parent ,otherwise it will take as normal field
	OrderName string `xml:"orderName"`
	OrderId int `xml:"orderId,omitempty"` //omitempty field will help to avoid forming the xml data if value not there
	OrderBy *string `xml:"orderBy"` //if we give pointer field, it will not consider that value because it willl be nil , otherwise default value is sent .this will not print unless it has value assigned
	Label      string `xml:",omitempty,attr"`  // attr field will be added to the tag like <orderinfo Label="TESTING">
	OrderDetails []OrderDetail `xml:"abc>bangalore>orderDetail"`  //it will form a hierarchy as <abc> <bangalore> <orderDetail>
	Comments string `xml:",comment"` // this will added to first tag of xml like  <!--sample example for xml-->
}

func main() {
	orderDet1:=OrderDetail{OrderName:"A"}
	orderDet2:=OrderDetail{OrderName:"A"}
	orderDet3:=OrderDetail{OrderName:"A"}
	order:=Order{
		OrderName: "test",
		OrderDetails: []OrderDetail{orderDet1,orderDet2,orderDet3},
		Label:"TESTING",
		Comments:"sample example for xml",
	}

	out,_:=xml.MarshalIndent(order,""," ")
	fmt.Println(string(out))
}

/**
Output is :

<orderinfo Label="TESTING">
 <orderName>test</orderName>
 <abc>
  <bangalore>
   <orderDetail>
    <orderName>A</orderName>
   </orderDetail>
   <orderDetail>
    <orderName>A</orderName>
   </orderDetail>
   <orderDetail>
    <orderName>A</orderName>
   </orderDetail>
  </bangalore>
 </abc>
 <!--sample example for xml-->
</orderinfo>

 */
