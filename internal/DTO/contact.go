package DTO

import "time"

type ContactResponse struct {
    ID int
    Username string
    GivenName string
    FamilyName string
    Phone []struct{
        TypeID int
        CountryCode int
        Operator int
        Number int
    }
    Email []string
    Birthdate time.Time
}

type GroupResponse struct {
    ID int
    Title string
    Description string
    Contacts []int
}
