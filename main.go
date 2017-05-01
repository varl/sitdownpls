package main

import (
    "encoding/json"
    "fmt"
    "os"
    "math"
)

type UnitFormat struct {
    Unit []GuestFormat
}

type GuestFormat struct {
    Name string `json:"name"`
    Gender string `json:"gender"`
}

type Table struct {
    Seats int
    Guests []GuestFormat
}

func (a *UnitFormat) UnmarshalJSON(b []byte) error {
    var s []GuestFormat
    if err := json.Unmarshal(b, &s); err != nil {
        return err
    }
    a.Unit = s
    return nil
}

func round(float float64) float64 {
    return math.Floor(float + 0.5)
}

func grab(amount int, groups []UnitFormat) []GuestFormat {
    var totBoys, totGirls, totAll int

    tableGuests := make([]GuestFormat, amount)

    for i := 0; totAll < amount; i++ {
        unit := groups[i].Unit
        totAll += len(unit)
    }

    fmt.Println(totBoys, totGirls, totAll, len(tableGuests))

    return tableGuests
}

func main() {
    file, err := os.Open("guest-list.json")

    if err != nil {
        panic(err)
    }
    defer file.Close()

    var units = make([]UnitFormat, 0)
    err = json.NewDecoder(file).Decode(&units)

    if err != nil {
        panic(err)
    }

    var guestCount float64
    var maleCount int
    var femaleCount int

    for i := 0; i < len(units); i++ {
        unit := units[i].Unit
        for k := 0; k < len(unit); k++ {
            guestCount += 1

            switch unit[k].Gender {
                case "male": maleCount += 1
                case "female": femaleCount += 1
            }

            fmt.Println(guestCount, unit[k].Name)
        }
    }

    guestsPerTable := 6.0
    tableCount := int(round(guestCount / guestsPerTable))
    lastTableGuestCount := int(math.Mod(guestCount, guestsPerTable))

    fmt.Println("\nStatistics")
    fmt.Println("----------")

    fmt.Println("Boys\t\t\t", maleCount)
    fmt.Println("Girls\t\t\t", femaleCount)
    fmt.Println("")
    fmt.Println("Guest count:\t\t", guestCount)
    fmt.Println("Guests per table:\t", guestsPerTable)
    fmt.Println("Table count:\t\t", tableCount)
    fmt.Println("...one table with:\t", lastTableGuestCount)
    fmt.Println("")

    tables := make([]Table, tableCount)

    for i := 0; i < tableCount; i++ {
        var count int
        if i != tableCount-1 {
            count = int(guestsPerTable)
        } else {
            count = int(lastTableGuestCount)
        }
        guests := grab(count, units)
        tables[i] = Table{Seats: count, Guests: guests}
    }

    fmt.Println(tables)
}
