package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

type doc struct {
	Meeting meeting `xml:"meeting"`
}

type meeting struct {
	Race race `xml:"race"`
}

type race struct {
	ID       int          `xml:"id,attr"`
	Name     string       `xml:"name,attr"`
	Distance int          `xml:"distance,attr"`
	Noms     []nomination `xml:"nomination"`
}

type nomination struct {
	Number int    `xml:"number,attr"`
	ID     int    `xml:"id,attr"`
	Horse  string `xml:"horse,attr"`
	Weight int    `xml:"weight,attr"`
	Rating int    `xml:"rating,attr"`
}

func main() {
	d := doc{}
	err := xml.Unmarshal(sample, &d)
	if err != nil {
		log.Fatalf("Unable to unmarshal XML: %s\n", err)
	}
	fmt.Printf("%#v\n", d)
}

var sample = []byte(`<xml>
<meeting>
  <race id="211075" number="1" nomnumber="7" division="0"
  name="TAB MORE THAN JUST WINNING HANDICAP" mediumname="2YF MDN"
  shortname="2YF MDN" stage="Results" distance="1000"
  minweight="57" raisedweight="0" class="MDN" age="2" grade="0"
  weightcondition="HCP" trophy="0" owner="0" trainer="0" jockey="0"
  strapper="0" totalprize="40000" first="23025" second="7925"
  third="3960" fourth="1885" fifth="955" time="2016-01-07T13:04:00"
  bonustype="BOB6" nomsfee="0" acceptfee="0"
  trackcondition="Heavy 10" timingmethod="Electronic"
  fastesttime="0-59.78" sectionaltime="600/35.41" formavailable="0"
  racebookprize="Of $40000. First $23025, second $7925, third $3960, fourth $1885, fifth $955, sixth $450, seventh $450, eighth $450, ninth $450, tenth $450">

    <nomination number="6" saddlecloth="6" horse="Kentucky Miss"
    id="201261" idnumber="" regnumber="" blinkers="0"
    trainernumber="1942" trainersurname="Pride"
    trainerfirstname="Joseph" trainertrack="Warwick Farm"
    rsbtrainername="Joseph Pride" jockeynumber="40133"
    jockeysurname="Shinn" jockeyfirstname="Blake" barrier="5"
    weight="58" rating="0"
    description="B F 2 Foxwedge x Twelve Pack Shelly(USA) (Deputy Storm(USA))"
    colours="Pink, Black Spots, Black Sleeves With Pink Spots, Pink Cap With Black Spots"
    owners="Aston Bloodstock (Mgr: W J Mula), Mrs R J Mula, Luskin Park Stud Racing (Mgr: P F Whelan), P B Duncan, M A Mcguinness, P Beaumont, D C Carter, S P Whelan, J M Whelan &amp; C G Staff"
    dob="2013-08-13T00:00:00" age="3" sex="F" career="0-0-0-0"
    thistrack="0-0-0-0" thisdistance="0-0-0-0" goodtrack="0-0-0-0"
    heavytrack="0-0-0-0" slowtrack="" deadtrack=""
    fasttrack="0-0-0-0" firstup="0-0-0-0" secondup="0-0-0-0"
    mindistancewin="0" maxdistancewin="0" finished="1"
    weightvariation="0" variedweight="58" decimalmargin="0.00"
    penalty="0" pricestarting="$4.00F" sectional200="0"
    sectional400="0" sectional600="0" sectional800="0"
    sectional1200="0" bonusindicator="E" />
    <nomination number="8" saddlecloth="8" horse="Mega Mall"
    id="201102" idnumber="" regnumber="" blinkers="0"
    trainernumber="81483" trainersurname="Cummings"
    trainerfirstname="James" trainertrack="Randwick"
    rsbtrainername="James Cummings" jockeynumber="57688"
    jockeysurname="Schofield" jockeyfirstname="Glyn" barrier="1"
    weight="58" rating="0"
    description="B F 2 Pendragon(NZ) x So Sydney (Dehere(USA))"
    colours="Black And White Check, Yellow Sleeves, Black And White Checked Cap"
    owners="Think Big Stud (Mgr: Dato Tan Chin Nam &amp; D G Ramage)"
    dob="2013-08-08T00:00:00" age="3" sex="F" career="0-0-0-0"
    thistrack="0-0-0-0" thisdistance="0-0-0-0" goodtrack="0-0-0-0"
    heavytrack="0-0-0-0" slowtrack="" deadtrack=""
    fasttrack="0-0-0-0" firstup="0-0-0-0" secondup="0-0-0-0"
    mindistancewin="0" maxdistancewin="0" finished="2"
    weightvariation="0" variedweight="58" decimalmargin="3.30"
    penalty="0" pricestarting="$12.00" sectional200="0"
    sectional400="0" sectional600="0" sectional800="0"
    sectional1200="0" bonusindicator="E" ></nomination>
    <nomination number="4" saddlecloth="4" horse="Hussterical"
    id="201100" idnumber="" regnumber="" blinkers="1"
    trainernumber="77973" trainersurname="Waterhouse"
    trainerfirstname="Gai" trainertrack="Randwick"
    rsbtrainername="Gai Waterhouse" jockeynumber="51661"
    jockeysurname="Berry" jockeyfirstname="Tommy" barrier="9"
    weight="58" rating="0"
    description="B F 2 Beneteau x Hussterics (Hussonet(USA))"
    colours="White, Darby Racing Logo, Navy Blue Armbands, Navy Blue Cap With White Pom Pom"
    owners="S G Darby, P J Allsop, R J Brown, T Buttenshaw, P C Beyerman, R J Corry, M A Davis, M De Nicolis, M Fradgley, P J Hall, C R Jessop, J C Lehner, R G Loveridge, T A Mesite, D M Preiss, M A Slack-Smith, D Wimbridge &amp; Chanceu Racing (Mgr: D Nicholas)"
    dob="2013-09-06T00:00:00" age="3" sex="F" career="0-0-0-0"
    thistrack="0-0-0-0" thisdistance="0-0-0-0" goodtrack="0-0-0-0"
    heavytrack="0-0-0-0" slowtrack="" deadtrack=""
    fasttrack="0-0-0-0" firstup="0-0-0-0" secondup="0-0-0-0"
    mindistancewin="0" maxdistancewin="0" finished="3"
    weightvariation="0" variedweight="58" decimalmargin="3.50"
    penalty="0" pricestarting="$5.00" sectional200="0"
    sectional400="0" sectional600="0" sectional800="0"
    sectional1200="0" bonusindicator=""></nomination>
  </race>
</meeting></xml>`)
