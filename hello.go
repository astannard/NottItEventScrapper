package main

import (
  "log"
  "strings"
"fmt"

  "github.com/PuerkitoBio/goquery"
)



func main() {
  var groups = getListOfGroups()
  
  for _, each := range groups {
    s := Scrapper(each)
      s.Scrape()
  }  
}

func getListOfGroups()([]Scrapper) {
  groups := []Scrapper{}

  groups = append(groups,MeetupGroup{url: "http://www.meetup.com/Mobile-Notts/", name: "Mobile Notts"})
  groups = append(groups,MeetupGroup{url: "http://www.meetup.com/dotnetnotts/", name: "Dot Net Notts"})
  groups = append(groups,MeetupGroup{url: "http://www.meetup.com/Nott-Tuesday-Getting-Nottinghams-tech-scene-together/", name: "Nott Tuesday"})
  groups = append(groups,MeetupGroup{url: "http://www.meetup.com/Tech-Nottingham/", name: "Tech Nottingham"})
  groups = append(groups,MeetupGroup{url: "http://www.meetup.com/NottinghamProgrammers/", name: "Nottingham Programmers"})
  groups = append(groups,MeetupGroup{url: "http://www.meetup.com/Nottingham-AWS-Meetup/", name: "Nottingham AWS Meetup"})

  return groups
}


type Scrapper interface {

    Scrape() []Event

}



type MeetupGroup struct {
    name, description, url string
}

type WebAPIGroup struct {
    name, description, url string
}


type Event struct {
    groupname string
    eventname string
   //date string
    description string
   // location string
}

//this function Area works on the type Rectangle and has the same function signature defined in the interface Shaper.  Therefore, Rectangle now implements the interface Shaper.
func (g MeetupGroup) Scrape() []Event {
   
    events := []Event{}
    doc, err := goquery.NewDocument(g.url) 
    if err != nil {
      log.Fatal(err)
    }

    doc.Find("#events-list-module .event-list li.event-item").Each(func(i int, s *goquery.Selection) {
      title := strings.TrimSpace(s.Find("h3").Text())
      description := strings.TrimSpace(s.Find(".event-desc").Text())
      var evnt = Event{eventname: title, description: description, groupname:g.name}

      events = append(events,evnt)
      fmt.Println(description)
    })


     return events
}



func (g WebAPIGroup) Scrape() []Event {
   
    events := []Event{}
    

    //Do some cool stuff here!


     return events
}
