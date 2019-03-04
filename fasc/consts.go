// {{{ Copyright (c) Paul R. Tagliamonte <paultag@gmail.com>, 2019
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE. }}}

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

// vim: foldmethod=marker
