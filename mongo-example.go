package main

import (        
        "gopkg.in/mgo.v2"
        "gopkg.in/mgo.v2/bson"
)

func main() {
        session, err := mgo.Dial("server1.example.com,server2.example.com")

        defer session.Close()		
        c := session.DB("test").C("people")
        err = c.Insert(&Person{"Ale", "+55 53 8116 9639"}, &Person{"Cla", "+55 53 8402 8510"})
                	
        result := Person{}
        err = c.Find(bson.M{"name": "Ale"}).One(&result)
        	
        fmt.Println("Phone:", result.Phone)
}