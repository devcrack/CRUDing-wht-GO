package calculation

import (
	"database/sql"
	"math/big"
)

type Calculation struct {
	EmployeeID         sql.NullInt64  `json:"employee_id"`
	CountryID          sql.NullInt64  `json:"country_id"`
	GroupID            sql.NullInt64  `json:"group_id"`
	DivisionID         sql.NullInt64  `json:"division_id"`
	DistrictID         sql.NullInt64  `json:"district_id"`
	RegionID           sql.NullInt64  `json:"region_id"`
	AreaID             sql.NullInt64  `json:"area_id"`
	CompanyID          sql.NullInt64  `json:"company_id"`
	EmployerRegistryID sql.NullInt64  `json:"employer_registry_id"`
	BranchID           sql.NullInt64  `json:"branch_id"`
	DepartmentID       sql.NullInt64  `json:"department_id"`
	PositionId         sql.NullInt64  `json:"position_id"`
	PayrollType        sql.NullInt64  `json:"payroll_type"`
	Year               sql.NullInt32  `json:"year"`
	Month              sql.NullInt32  `json:"month"`
	CalendarID         sql.NullInt64  `json:"calendar_id"`
	Period             sql.NullInt32  `json:"period"`
	ConceptType        sql.NullInt32  `json:"concept_type"`
	ConceptKey         sql.NullString `json:"concept_key"`
	SatKey             sql.NullString `json:"sat_key"`
	ConceptName        sql.NullString `json:"concept_name"`
	Taxed              *big.Float     `json:"taxed"`
	Exempt             *big.Float     `json:"exempt"`
	Amount             *big.Float     `json:"amount"`
	Days               sql.NullInt32  `json:"days"`
	IDUser             sql.NullInt64  `json:"id_user"`
	Formula            sql.NullString `json:"formula"`
	Interpreted        sql.NullString `json:"interpreted"`
	PayrollID          sql.NullInt64  `json:"payroll_id"`
}
