package main

import (
    "encoding/json"
    "testing"
)

var JSON = `{"birth": 673675640,
  "city": "Росоштадт",
  "country": "Малатрис",
  "email": "evehnensoty@rambler.ru",
  "fname": "Филипп",
  "id": 1,
  "interests": ["Путешествия", "Аватар", "На открытом воздухе"],
  "joined": 1344556800,
  "likes": [{"id": 840, "ts": 1465179671},
  {"id": 8832, "ts": 1456886999},
  {"id": 6842, "ts": 1473416865},
  {"id": 7518, "ts": 1499915559},
  {"id": 1930, "ts": 1488292582},
  {"id": 9968, "ts": 1470472322},
  {"id": 394, "ts": 1461063706},
  {"id": 8140, "ts": 1514416895},
  {"id": 4620, "ts": 1504026891},
  {"id": 9702, "ts": 1477407606},
  {"id": 3156, "ts": 1466420098},
  {"id": 5922, "ts": 1507605010},
  {"id": 4768, "ts": 1456602673},
  {"id": 8818, "ts": 1512780050},
  {"id": 9556, "ts": 1478295315},
  {"id": 178, "ts": 1516099087},
  {"id": 1842, "ts": 1479334825},
  {"id": 22, "ts": 1497325065},
  {"id": 8350, "ts": 1494590491},
  {"id": 1878, "ts": 1492802929},
  {"id": 7040, "ts": 1532977148},
  {"id": 4712, "ts": 1520372752},
  {"id": 598, "ts": 1469848628},
  {"id": 4756, "ts": 1474778459},
  {"id": 7938, "ts": 1511478908},
  {"id": 3976, "ts": 1529563562},
  {"id": 7542, "ts": 1495878743},
  {"id": 8308, "ts": 1537134657},
  {"id": 2658, "ts": 1495707978},
  {"id": 2878, "ts": 1488896049},
  {"id": 2632, "ts": 1462084804},
  {"id": 9636, "ts": 1464512415},
  {"id": 136, "ts": 1469446164},
  {"id": 4952, "ts": 1539009138},
  {"id": 2082, "ts": 1477353483},
  {"id": 1782, "ts": 1456629485},
  {"id": 4578, "ts": 1518741825}],
  "phone": "8(984)8611629",
  "sex": "m",
  "sname": "Пенетасян",
    "status": "всё сложно"}`

func TestJsonParser(t *testing.T) {
    var a Account

    if err := json.Unmarshal([]byte(JSON), &a); err != nil {
        t.Error("Json parsing error")
    }
}

func TestGetAccountMiddle(t *testing.T) {
    testString := "sdfsfsdfsdfsdf"
    if checkString := getAccount(testString); checkString != testString {
        t.Error("Middle check failed")
    }
}

func TestGetAccountNoRight(t *testing.T) {
    testString := "sdfs{fsdfsdfsdf"
    if checkString := getAccount(testString); checkString != "{fsdfsdfsdf" {
        t.Error("No right } check failed")
    }
}

func TestGetAccountRigthAndLeft(t *testing.T) {
    testString := "sdfs{fsdfsdf}sdf"
    if checkString := getAccount(testString); checkString != "{fsdfsdf}" {
        t.Error("Middle {} check failed: %s", checkString)
    }
}
