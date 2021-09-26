package domains

type User struct {
	Given_Name         string `"json": "given_name"`
	Family_Name        string `"json": "family_name"`
	Email              string `json:"email"`
	Company_Id         string `json: "company_id"`
	Product_Codes      string `json: "product_codes"`
	Is_Check_Point_Use bool   `json : "is_check_point_use"`
	Pass               string `json:"pass"`
}

var Guser = User{
	Given_Name:         "Hasan",
	Family_Name:        "Gmail",
	Email:              "hsmasud@gmail.com",
	Company_Id:         "orbitax",
	Product_Codes:      "WWITPQ|WWEATQ",
	Is_Check_Point_Use: false,
	Pass:               "asdf",
}

var UserList []User
