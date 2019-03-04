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

import (
	"fmt"
)

type AgencyCode int

func (o AgencyCode) String() string {
	v, ok := FIPS95_2AgencyCodes[int(o)]
	if !ok {
		return fmt.Sprintf("%d", o)
	}
	return v
}

var (
	FIPS95_2AgencyCodes = map[int]string{
		0300: "LIBRARY OF CONGRESS",
		0500: "GOVERNMENT ACCOUNTABILITY OFFICE",
		1027: "ADMINISTRATIVE OFFICE OF THE U.S. COURTS",
		1100: "EXECUTIVE OFFICE OF THE PRESIDENT",
		1145: "PEACE CORPS",
		1153: "UNITED STATES TRADE AND DEVELOPMENT AGENCY",
		1200: "AGRICULTURE, DEPARTMENT OF",
		1300: "COMMERCE, DEPARTMENT OF",
		1400: "INTERIOR, DEPARTMENT OF THE",
		1500: "JUSTICE, DEPARTMENT OF",
		1600: "LABOR, DEPARTMENT OF",
		1665: "PENSION BENEFIT GUARANTY CORPORATION",
		1800: "POSTAL SERVICE",
		1900: "STATE, DEPARTMENT OF",
		2000: "TREASURY, DEPARTMENT OF THE",
		2300: "U.S. TAX COURT",
		2400: "OFFICE OF PERSONNEL MANAGEMENT",
		2700: "FEDERAL COMMUNICATIONS COMMISSION",
		2800: "SOCIAL SECURITY ADMINISTRATION",
		2900: "FEDERAL TRADE COMMISSION",
		3100: "NUCLEAR REGULATORY COMMISSION",
		3300: "SMITHSONIAN INSTITUTION",
		3352: "J. F. KENNEDY CENTER FOR THE PERFORMING ARTS",
		3355: "NATIONAL GALLERY OF ART",
		3400: "INTERNATIONAL TRADE COMMISSION",
		3600: "VETERANS AFFAIRS, DEPARTMENT OF",
		4100: "MERIT SYSTEMS PROTECTION BOARD",
		4500: "EQUAL EMPLOYMENT OPPORTUNITY COMMISSION",
		4602: "APPALACHIAN REGIONAL COMMISSION",
		4700: "GENERAL SERVICES ADMINISTRATION",
		4900: "NATIONAL SCIENCE FOUNDATION",
		5000: "SECURITIES AND EXCHANGE COMMISSION",
		5300: "THE INSTITUE OF MUSEUM AND LIBRARY SERVICES",
		5400: "FEDERAL LABOR RELATIONS AUTHORITY",
		5920: "NATIONAL ENDOWMENT FOR THE ARTS",
		5940: "NATIONAL ENDOWMENT FOR THE HUMANITIES",
		6000: "RAILROAD RETIREMENT BOARD",
		6100: "CONSUMER PRODUCT SAFETY COMMISSION",
		6201: "OFFICE OF SPECIAL COUNSEL",
		6300: "NATIONAL LABOR RELATIONS BOARD",
		6500: "FEDERAL MARITIME COMMISSION",
		6800: "ENVIRONMENTAL PROTECTION AGENCY",
		6900: "TRANSPORTATION, DEPARTMENT OF",
		7000: "HOMELAND SECURITY, DEPARTMENT OF",
		7100: "OVERSEAS PRIVATE INVESTMENT CORPORATION",
		7200: "AGENCY FOR INTERNATIONAL DEVELOPMENT",
		7300: "SMALL BUSINESS ADMINISTRATION",
		7400: "AMERICAN BATTLE MONUMENTS COMMISSION",
		7500: "HEALTH AND HUMAN SERVICES, DEPARTMENT OF",
		8000: "NATIONAL AERONAUTICS AND SPACE ADMINISTRATION",
		8300: "EXPORT-IMPORT BANK OF THE U.S.",
		8600: "HOUSING AND URBAN DEVELOPMENT, DEPARTMENT OF",
		8800: "NATIONAL ARCHIVES AND RECORDS ADMINISTRATION",
		8900: "ENERGY, DEPARTMENT OF",
		9000: "SELECTIVE SERVICE SYSTEM",
		9100: "EDUCATION, DEPARTMENT OF",
		9300: "FEDERAL MEDIATION AND CONCILIATION SERVICE",
		9502: "NATIONAL CAPITAL PLANNING COMMISSION",
		9504: "FEDERAL MINE SAFETY AND HEALTH REVIEW COMMISSION",
		9505: "SURFACE TRANSPORTATION BOARD",
		9506: "FEDERAL ELECTION COMMISSION",
		9507: "COMMODITY FUTURES TRADING COMMISSION",
		9508: "NATIONAL TRANSPORTATION SAFETY BOARD",
		9511: "THE COUNCIL OF THE INSPECTORS GENERAL ON INTEGRITY AND EFFICIENCY ",
		9513: "MARINE MAMMAL COMMISSION",
		9514: "OCCUPATIONAL SAFETY AND HEALTH REVIEW COMMISSION",
		9515: "ADMINISTRATIVE CONFERENCE OF THE U. S.",
		9516: "DEFENSE NUCLEAR FACILITIES SAFETY BOARD",
		9518: "COMMITTEE FOR PURCHASE FROM PEOPLE WHO ARE BLIND OR SEVERELY DISABLED",
		9523: "ELECTION ASSISTANCE COMMISSION",
		9524: "NATIONAL MEDIATION BOARD",
		9526: "PRIVACY AND CIVIL LIBERTIES OVERSIGHT BOARD",
		9532: "ARCHITECTURAL AND TRANSPORTATION BARRIERS COMPLIANCE BOARD",
		9533: "GULF COAST ECOSYSTEM RESTORATION COUNCIL",
		9536: "VIETNAM EDUCATION FOUNDATION",
		9537: "COMMISSION OF FINE ARTS",
		9542: "FEDERAL HOUSING FINANCE AGENCY",
		9543: "MILLENIUM CHALLENGE CORPORATION",
		9545: "MORRIS K. UDALL SCHOLARSHIP AND EXCELLENCE IN NATIONAL ENVIRONMENTAL POLICY FOUNDATION",
		9549: "OFFICE OF GOVERNMENT ETHICS",
		9557: "NORTHERN BORDER REGIONAL COMMISSION",
		9565: "CHEMICAL SAFETY AND HAZARD INVESTIGATION BOARD",
		9568: "UNITED STATES AGENCY FOR GLOBAL MEDIA, BBG",
		9572: "DENALI COMMISSION",
		9577: "CORPORATION FOR NATIONAL AND COMMUNITY SERVICE",
		9594: "COURT SERVICES AND OFFENDER SUPERVISION AGENCY",
		9700: "DEPT OF DEFENSE",
		5700: "DEPT OF AIR FORCE",
		1700: "DEPT OF NAVY",
	}
)

// vim: foldmethod=marker
