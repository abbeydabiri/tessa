package database

var sqlTypes = map[string]string{
	"bool":    "bool",
	"time":    "timestamp",
	"string":  "text",
	"int":     "int",
	"uint":    "int",
	"int64":   "int8",
	"uint32":  "int8",
	"uint64":  "int8",
	"float32": "float8",
	"float64": "float8",
}

//Tables ...
type Tables interface {
	ToMap() (mapInterface map[string]interface{})
	FillStruct(tableMap map[string]interface{}) error
}

//AllTables ...
var AllTables = map[string]Tables{
	"Accounts":      &Accounts{},
	"AccountTokens": &AccountTokens{},
	"Activations":   &Activations{},
	"Campaigns":     &Campaigns{},
	"Currencys":     &Currencys{},
	"Contacts":      &Contacts{},
	"Followers":     &Followers{},
	"Hits":          &Hits{},
	"Networks":      &Networks{},
	"Newsletters":   &Newsletters{},
	"Permissions":   &Permissions{},
	"Profiles":      &Profiles{},
	"Roles":         &Roles{},
	"Seocontents":   &Seocontents{},
	"Settings":      &Settings{},
	"Smtps":         &Smtps{},
	"Tasks":         &Tasks{},
	"Tokens":        &Tokens{},
	"Transactions":  &Transactions{},
	"Users":         &Users{},
	"Wallets":       &Wallets{},
}

func createTable(tableName string) (Message []string) {
	switch tableName {
	default:
		tableName = tableName[1:]
		if AllTables[tableName] != nil {
			Message = append(Message, Fields{}.sqlCreate(AllTables[tableName]))
			switch tableName {
			case "Networks":
				networks := new(Networks)
				networks.Setup()
			}
		} else {
			Message = append(Message, "Please Specify Table")
		}

	case "/all":
		for _, table := range AllTables {
			Message = append(Message, Fields{}.sqlCreate(table))
			switch tableName {
			case "Networks":
				networks := new(Networks)
				networks.Setup()
			}
		}
	}
	return Message
}

// SearchParams serves as default parameters used in generating sql prepared statements
type SearchParams struct {
	Field, Text, 
	Workflow string

	Skip, Limit,
	DocumentID, CID uint64

	ID uint64

	OrderBy  string
	OrderAsc bool

	CampaignID, OwnerID,
	UserID uint64

	Filter map[string]string
}
