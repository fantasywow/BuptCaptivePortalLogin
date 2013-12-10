package main

import(
	"net/http"
	"net/url"
)

func main(){
	http.PostForm("http://gw.bupt.edu.cn",url.Values{"DDDDD": {"2012110669"}, "upass": {"2dbc74044f3b09602559cda5137abb3f123456781"}, "R1": {"0"}, "R2": {"1"}, "para": {"00"}, "0MKKey": {"123456"}})
}


/*
DDDDD:2012110669
upass:2dbc74044f3b09602559cda5137abb3f123456781
R1:0
R2:1
para:00
0MKKey:123456

*/
//Content-type:text/html