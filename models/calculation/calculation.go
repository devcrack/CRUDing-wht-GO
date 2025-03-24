package calculation

import "math/big"

type Calculation struct {
	EmployeeID         int64     `json:"employee_id"`
	CountryID          int64     `json:"country_id"`
	GroupID            int64     `json:"group_id"`
	DivisionID         int64     `json:"division_id"`
	DistrictID         int64     `json:"district_id"`
	RegionID           int64     `json:"region_id"`
	AreaID             int64     `json:"area_id"`
	CompanyID          int64     `json:"company_id"`
	EmployerRegistryID int64     `json:"employer_registry_id"`
	BranchID           int64     `json:"branch_id"`
	DepartmentID       int64     `json:"department_id"`
	PositionId         int64     `json:"position_id"`
	PayrollType        int64     `json:"payroll_type"`
	Year               int       `json:"year"`
	Month              int       `json:"month"`
	CalendarID         int64     `json:"calendar_id"`
	Period             int       `json:"period"`
	ConceptType        int       `json:"concept_type"`
	ConceptKey         string    `json:"concept_key"`
	SatKey             string    `json:"sat_key"`
	ConceptName        string    `json:"concept_name"`
	Taxed              big.Float `json:"taxed"`
	Exempt             big.Float `json:"exempt"`
	Amount             big.Float `json:"amount"`
	Days               int       `json:"days"`
	IDUser             int64     `json:"id_user"`
	Formula            string    `json:"formula"`
	Interpreted        string    `json:"interpreted"`
	PayrollID          int64     `json:"payroll_id"`
}
