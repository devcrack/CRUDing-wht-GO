package calculations

import (
	"database/sql"
	"math"
	"math/big"
	"payroll_calculator/models/calculation"
	"payroll_calculator/repository"
)

type CalculationRepository struct {
	*repository.Repository
}

func (r *CalculationRepository) GettingAllCalculations() ([]calculation.Calculation, error) {
	query := `
    SELECT 
        int_employee_calculation,
        int_country_calculation,
        int_group_calculation,
        int_division_calculation,
        int_district_calculation,
        int_region_calculation,
        int_area_calculation,
        int_company_calculation,
        int_employer_registry_calculation,
        int_branch_calculation,
        int_department_calculation,
        int_position_calculation,
        int_payroll_type_calculation,
        int_year_calculation,
        int_month_calculation,
        int_calendars_calculation,
        int_period_calculation,
        int_concept_type_calculation,
        var_concept_key_calculation,
        var_sat_key_calculation,
        var_concept_name_calculation,
        dec_taxed_calculation,
        dec_exempt_calculation,
        dec_amount_calculation,
        dec_datum_calculation,
        fk_int_user,
        var_formula_calculation,
        var_interpreted_formula_calculation,
        int_payroll_calculation
    FROM calculations`
	rows, err := r.RawSql.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var calculations []calculation.Calculation
	for rows.Next() {
		var calculation calculation.Calculation
		var taxedStr, exemptStr, amountStr sql.NullString
		var daysNull sql.NullFloat64

		nwErr := rows.Scan(&calculation.EmployeeID,
			&calculation.CountryID,
			&calculation.GroupID,
			&calculation.DivisionID,
			&calculation.DistrictID,
			&calculation.RegionID,
			&calculation.AreaID,
			&calculation.CompanyID,
			&calculation.EmployerRegistryID,
			&calculation.BranchID,
			&calculation.DepartmentID,
			&calculation.PositionId,
			&calculation.PayrollType,
			&calculation.Year,
			&calculation.Month,
			&calculation.CalendarID,
			&calculation.Period,
			&calculation.ConceptType,
			&calculation.ConceptKey,
			&calculation.SatKey,
			&calculation.ConceptName,
			&taxedStr,
			&exemptStr,
			&amountStr,
			&daysNull,
			&calculation.IDUser,
			&calculation.Formula,
			&calculation.Interpreted,
			&calculation.PayrollID)
		if nwErr != nil {
			return nil, nwErr
		}

		calculation.Taxed = new(big.Float)
		calculation.Exempt = new(big.Float)
		calculation.Amount = new(big.Float)

		if taxedStr.Valid {
			calculation.Taxed.SetString(taxedStr.String)
		}
		if exemptStr.Valid {
			calculation.Exempt.SetString(exemptStr.String)
		}
		if amountStr.Valid {
			calculation.Amount.SetString(amountStr.String)
		}
		if daysNull.Valid {
			calculation.Days = sql.NullInt32{Int32: int32(math.Round(daysNull.Float64)),
				Valid: true}
		}
		calculations = append(calculations, calculation)
	}
	return calculations, nil
}
