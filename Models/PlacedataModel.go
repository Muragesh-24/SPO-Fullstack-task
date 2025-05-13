package model





type Experience struct {
    Name                 string `json:"name" bson:"name"`
    Company              string `json:"company" bson:"company"`
    Role                 string `json:"role" bson:"role"`
    SelectionProcess     string `json:"selection_process" bson:"selection_process"`
    PreparationStrategy  string `json:"preparation_strategy" bson:"preparation_strategy"`
    Tips                 string `json:"tips" bson:"tips"`
}
