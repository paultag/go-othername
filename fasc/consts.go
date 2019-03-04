package fasc

type OrganizationalCategory int
type AssociationCategory int

const (
	OrganizationalCategoryFederalGoverment     OrganizationalCategory = 1
	OrganizationalCategoryStateGovernment      OrganizationalCategory = 2
	OrganizationalCategoryCommercialEnterprise OrganizationalCategory = 3
	OrganizationalCategoryForeignGovernment    OrganizationalCategory = 4

	AssociationCategoryEmployee          AssociationCategory = 1
	AssociationCategoryCivil             AssociationCategory = 2
	AssociationCategoryExecutiveStaff    AssociationCategory = 3
	AssociationCategoryUniformedServivce AssociationCategory = 4
	AssociationCategoryContractor        AssociationCategory = 5
	AssociationCategoryAffiliate         AssociationCategory = 6
	AssociationCategoryBeneficiary       AssociationCategory = 7
)

func (oc OrganizationalCategory) String() string {
	switch oc {
	case OrganizationalCategoryFederalGoverment:
		return "federal government"
	case OrganizationalCategoryStateGovernment:
		return "state government"
	case OrganizationalCategoryCommercialEnterprise:
		return "commercial enterprise"
	case OrganizationalCategoryForeignGovernment:
		return "foreign government"
	}
	return "unknown"
}

func (ac AssociationCategory) String() string {
	switch ac {
	case AssociationCategoryEmployee:
		return "employee"
	case AssociationCategoryCivil:
		return "civil"
	case AssociationCategoryExecutiveStaff:
		return "execstaff"
	case AssociationCategoryUniformedServivce:
		return "uniformed service"
	case AssociationCategoryContractor:
		return "contractor"
	case AssociationCategoryAffiliate:
		return "affiliate"
	case AssociationCategoryBeneficiary:
		return "beneficiary"
	}
	return "unknown"
}
