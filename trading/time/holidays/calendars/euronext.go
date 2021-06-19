package calendars

import (
	"time"
)

// EuroNext implements a EuroNext holiday (non-trading day) calendar.
//
// Since 2020 Dublin and Oslo Børs are part of EuroNext, but they are
// not included because they have extra holidays which will make this
// calendar MIC-dependent.
//
// EuroNext calendar of business days 2021.
//  - Jan 1,  New Year Day
//  - Apr 1,  Maundy Thursday (Oslo Børs)
//  - Apr 2,  Good Friday
//  - Apr 5,  Easter Monday
//  - May 3,  May Day (Euronext Dublin)
//  - May 13, (Oslo Børs)
//  - May 17, Constitution Day (Oslo Børs)
//  - May 24, Whit Monday (Oslo Børs)
//  - Dec 24, Christmas Eve (Oslo Børs)
//  - Dec 27, Substitute for Christmas Day (Euronext Dublin)
//  - Dec 28, Substitute for St Stephens/Boxing Day (Euronext Dublin)
//  - Dec 31, New Year Eve (Oslo Børs)
//
// EuroNext calendar of business days 2020. Oslo Børs has independent calendar.
//  - Jan 1,  New Year Day
//  - Apr 10, Good Friday
//  - Apr 13, Easter Monday
//  - May 1,  Labour Day
//  - Dec 25, Christmas Day
//
// EuroNext calendar of business days 2019.
//  - Jan 1,  New Year Day
//  - Apr 19, Good Friday
//  - Apr 22, Easter Monday
//  - May 1,  Labour Day
//  - Dec 25, Christmas Day
//  - Dec 26, Boxing Day
//
// EuroNext calendar of business days 2018.
//  - Jan 1,  New Year Day
//  - Mar 30, Good Friday
//  - Apr 2,  Easter Monday
//  - May 1,  Labour Day
//  - Dec 25, Christmas Day
//  - Dec 26, Boxing Day
//
// EuroNext calendar of business days 2017.
//  - Apr 14, Good Friday
//  - Apr 17, Easter Monday
//  - May 1,  Labour Day
//  - Dec 25, Christmas Day
//  - Dec 26, Boxing Day
//
// EuroNext calendar of business days 2016.
//  - Jan 1,  New Year Day
//  - Mar 25, Good Friday
//  - Mar 28, Easter Monday
//  - Dec 26, Boxing Day
//
// EuroNext calendar of business days 2015.
//  - Jan 1,  New Year Day
//  - Apr 3,  Good Friday
//  - Apr 6,  Easter Monday
//  - May 1,  Labour Day
//  - Dec 25, Christmas Day
//
// EuroNext calendar of business days 2014.
//  - Jan 1,  New Year Day
//  - Apr 18, Good Friday
//  - Apr 21, Easter Monday
//  - May 1,  Labour Day
//  - Dec 25, Christmas Day
//  - Dec 26, Boxing Day
//
// EuroNext calendar of business days 2013.
//  - Jan 1,  New Year Day
//  - Mar 29, Good Friday
//  - Apr 1,  Easter Monday
//  - May 1,  Labour Day
//  - Dec 25, Christmas Day
//  - Dec 26, Boxing Day
//
// EuroNext calendar of business days 2012.
//  - Jan 1,  New Year Day
//  - Apr 6,  Good Friday
//  - Apr 9,  Easter Monday
//  - May 1,  Labour Day
//  - Dec 25, Christmas Day
//  - Dec 26, Boxing Day
//
// EuroNext calendar of business days 2011.
//  - Apr 22, Good Friday
//  - Apr 25, Easter Monday
//  - Dec 26, Boxing Day
//
// EuroNext calendar of business days 2010.
//  - Jan 1,  New Year Day
//  - Apr 2,  Good Friday
//  - Apr 5,  Easter Monday
//
// EuroNext calendar of business days 2009.
//  - Jan 1,  New Year Day
//  - Apr 10, Good Friday
//  - Apr 13, Easter Monday
//  - May 1,  Labour Day
//  - Dec 25, Christmas Day
//
// EuroNext calendar of business days 2008.
//  - Jan 1,  New Year Day
//  - Mar 21, Good Friday
//  - Mar 24, Easter Monday
//  - May 1,  Labour Day
//  - Dec 25, Christmas Day
//  - Dec 26, Boxing Day
//
// EuroNext calendar of business days 2007.
//  - Jan 1,  New Year Day
//  - Apr 6,  Good Friday
//  - Apr 9,  Easter Monday
//  - May 1,  Labour Day
//  - Dec 25, Christmas Day
//  - Dec 26, Boxing Day
//
// EuroNext calendar of business days 2006.
//  - Apr 14, Good Friday
//  - Apr 17, Easter Monday
//  - May 1,  Labour Day
//  - Dec 25, Christmas Day
//  - Dec 26, Boxing Day
//
// EuroNext calendar of business days 2005.
//  - Mar 25, Good Friday
//  - Mar 28, Easter Monday
//  - Dec 26, Boxing Day
//
// EuroNext calendar of business days 2004.
//  - Jan 1,  New Year Day
//  - Apr 9,  Good Friday
//  - Apr 12, Easter Monday
//
// EuroNext calendar of business days 2003.
//  - Jan 1,  New Year Day
//  - Apr 18, Good Friday
//  - Apr 21, Easter Monday
//  - May 1,  Labour Day
//  - Dec 25, Christmas Day
//  - Dec 26, Boxing Day
//
// EuroNext calendar of business days 2002.
//  - Jan 1,  New Year Day
//  - Mar 29, Good Friday
//  - Apr 1,  Easter Monday
//  - May 1,  Labour Day
//  - Dec 25, Christmas Day
//  - Dec 26, Boxing Day
//
// EuroNext calendar of business days 2001.
//  - Jan 1,  New Year Day
//  - Apr 13, Good Friday
//  - Apr 16, Easter Monday
//  - May 1,  Labour Day
//  - Jun 4,  Whit Monday
//  - Dec 25, Christmas Day
//  - Dec 26, Boxing Day
//  - Dec 31, Last trading day - change over to euro
//
// EuroNext calendar of business days 2000.
//  - Apr 21, Good Friday
//  - Apr 24, Easter Monday
//  - May 1,  Labour Day
//  - Jun 12, Whit Monday
//  - Dec 25, Christmas Day
//  - Dec 26, Boxing Day
//
// EuroNext calendar of business days 1999.
//  - Jan 1,  New Year Day
//  - Apr 2,  Good Friday
//  - Apr 5,  Easter Monday
//  - May 24, Whit Monday
//
type EuroNext struct{}

//nolint:cyclop,gomnd,funlen
// IsHoliday implements Calendarer interface.
func (EuroNext) IsHoliday(t time.Time) bool {
	if checkWeekend(t) {
		return true
	}

	y, m, d := t.Date()

	switch y {
	case 2021:
		return euronext2021(m, d)
	case 2020:
		return euronext2020(m, d)
	case 2019:
		return euronext2019(m, d)
	case 2018:
		return euronext2018(m, d)
	case 2017:
		return euronext2017(m, d)
	case 2016:
		return euronext2016(m, d)
	case 2015:
		return euronext2015(m, d)
	case 2014:
		return euronext2014(m, d)
	case 2013:
		return euronext2013(m, d)
	case 2012:
		return euronext2012(m, d)
	case 2011:
		return euronext2011(m, d)
	case 2010:
		return euronext2010(m, d)
	case 2009:
		return euronext2009(m, d)
	case 2008:
		return euronext2008(m, d)
	case 2007:
		return euronext2007(m, d)
	case 2006:
		return euronext2006(m, d)
	case 2005:
		return euronext2005(m, d)
	case 2004:
		return euronext2004(m, d)
	case 2003:
		return euronext2003(m, d)
	case 2002:
		return euronext2002(m, d)
	case 2001:
		return euronext2001(m, d)
	case 2000:
		return euronext2000(m, d)
	case 1999:
		return euronext1999(m, d)
	}

	// Out of predefined range of years.
	// Check for New Year Day, Christmas Day and Boxing Day.
	switch {
	case m == time.January && d == 1,
		m == time.May && d == 1,
		m == time.December && (d == 25 || d == 26):
		return true
	}

	// Check for Good Friday and Easter Monday.
	return computusFridayMonday(y, t.YearDay())
}

// EuroNext calendar of business days 2021.
//  - Jan 1,  New Year Day
//  - Apr 1,  Maundy Thursday (Oslo Børs)
//  - Apr 2,  Good Friday
//  - Apr 5,  Easter Monday
//  - May 3,  May Day (Euronext Dublin)
//  - May 13, (Oslo Børs)
//  - May 17, Constitution Day (Oslo Børs)
//  - May 24, Whit Monday (Oslo Børs)
//  - Dec 24, Christmas Eve (Oslo Børs)
//  - Dec 27, Substitute for Christmas Day (Euronext Dublin)
//  - Dec 28, Substitute for St Stephens/Boxing Day (Euronext Dublin)
//  - Dec 31, New Year Eve (Oslo Børs)
func euronext2021(m time.Month, d int) bool {
	switch m {
	case time.January:
		return d == 1
	case time.April:
		return d == 2 || d == 5
	default:
		return false
	}
}

//nolint:gomnd
// EuroNext calendar of business days 2020. Oslo Børs has independent calendar.
//  - Jan 1,  New Year Day
//  - Apr 10, Good Friday
//  - Apr 13, Easter Monday
//  - May 1,  Labour Day
//  - Dec 25, Christmas Day
func euronext2020(m time.Month, d int) bool {
	switch m {
	case time.January:
		return d == 1
	case time.April:
		return d == 10 || d == 13
	case time.May:
		return d == 1
	case time.December:
		return d == 25
	default:
		return false
	}
}

// EuroNext calendar of business days 2019.
//  - Jan 1,  New Year Day
//  - Apr 19, Good Friday
//  - Apr 22, Easter Monday
//  - May 1,  Labour Day
//  - Dec 25, Christmas Day
//  - Dec 26, Boxing Day
func euronext2019(m time.Month, d int) bool {
	switch m {
	case time.January:
		return d == 1
	case time.April:
		return d == 19 || d == 22
	case time.May:
		return d == 1
	case time.December:
		return d == 25 || d == 26
	default:
		return false
	}
}

//nolint:gomnd
// EuroNext calendar of business days 2018.
//  - Jan 1,  New Year Day
//  - Mar 30, Good Friday
//  - Apr 2,  Easter Monday
//  - May 1,  Labour Day
//  - Dec 25, Christmas Day
//  - Dec 26, Boxing Day
func euronext2018(m time.Month, d int) bool {
	switch m {
	case time.January:
		return d == 1
	case time.March:
		return d == 30
	case time.April:
		return d == 2
	case time.May:
		return d == 1
	case time.December:
		return d == 25 || d == 26
	default:
		return false
	}
}

// EuroNext calendar of business days 2017.
//  - Apr 14, Good Friday
//  - Apr 17, Easter Monday
//  - May 1,  Labour Day
//  - Dec 25, Christmas Day
//  - Dec 26, Boxing Day
func euronext2017(m time.Month, d int) bool {
	switch m {
	case time.April:
		return d == 14 || d == 17
	case time.May:
		return d == 1
	case time.December:
		return d == 25 || d == 26
	default:
		return false
	}
}

//nolint:gomnd
// EuroNext calendar of business days 2016.
//  - Jan 1,  New Year Day
//  - Mar 25, Good Friday
//  - Mar 28, Easter Monday
//  - Dec 26, Boxing Day
func euronext2016(m time.Month, d int) bool {
	switch m {
	case time.January:
		return d == 1
	case time.March:
		return d == 25 || d == 28
	case time.December:
		return d == 26
	default:
		return false
	}
}

//nolint:gomnd
// EuroNext calendar of business days 2015.
//  - Jan 1,  New Year Day
//  - Apr 3,  Good Friday
//  - Apr 6,  Easter Monday
//  - May 1,  Labour Day
//  - Dec 25, Christmas Day
func euronext2015(m time.Month, d int) bool {
	switch m {
	case time.January:
		return d == 1
	case time.April:
		return d == 3 || d == 6
	case time.May:
		return d == 1
	case time.December:
		return d == 25
	default:
		return false
	}
}

// EuroNext calendar of business days 2014.
//  - Jan 1,  New Year Day
//  - Apr 18, Good Friday
//  - Apr 21, Easter Monday
//  - May 1,  Labour Day
//  - Dec 25, Christmas Day
//  - Dec 26, Boxing Day
func euronext2014(m time.Month, d int) bool {
	switch m {
	case time.January:
		return d == 1
	case time.April:
		return d == 18 || d == 21
	case time.May:
		return d == 1
	case time.December:
		return d == 25 || d == 26
	default:
		return false
	}
}

//nolint:gomnd
// EuroNext calendar of business days 2013.
//  - Jan 1,  New Year Day
//  - Mar 29, Good Friday
//  - Apr 1,  Easter Monday
//  - May 1,  Labour Day
//  - Dec 25, Christmas Day
//  - Dec 26, Boxing Day
func euronext2013(m time.Month, d int) bool {
	switch m {
	case time.January:
		return d == 1
	case time.March:
		return d == 29
	case time.April:
		return d == 1
	case time.May:
		return d == 1
	case time.December:
		return d == 25 || d == 26
	default:
		return false
	}
}

// EuroNext calendar of business days 2012.
//  - Jan 1,  New Year Day (Sunday)
//  - Apr 6,  Good Friday
//  - Apr 9,  Easter Monday
//  - May 1,  Labour Day
//  - Dec 25, Christmas Day
//  - Dec 26, Boxing Day
func euronext2012(m time.Month, d int) bool {
	switch m {
	case time.April:
		return d == 6 || d == 9
	case time.May:
		return d == 1
	case time.December:
		return d == 25 || d == 26
	default:
		return false
	}
}

//nolint:gomnd
// EuroNext calendar of business days 2011.
//  - Apr 22, Good Friday
//  - Apr 25, Easter Monday
//  - Dec 26, Boxing Day
func euronext2011(m time.Month, d int) bool {
	switch m {
	case time.April:
		return d == 22 || d == 25
	case time.December:
		return d == 26
	default:
		return false
	}
}

// EuroNext calendar of business days 2010.
//  - Jan 1,  New Year Day
//  - Apr 2,  Good Friday
//  - Apr 5,  Easter Monday
func euronext2010(m time.Month, d int) bool {
	switch m {
	case time.January:
		return d == 1
	case time.April:
		return d == 2 || d == 5
	default:
		return false
	}
}

//nolint:gomnd
// EuroNext calendar of business days 2009.
//  - Jan 1,  New Year Day
//  - Apr 10, Good Friday
//  - Apr 13, Easter Monday
//  - May 1,  Labour Day
//  - Dec 25, Christmas Day
func euronext2009(m time.Month, d int) bool {
	switch m {
	case time.January:
		return d == 1
	case time.April:
		return d == 10 || d == 13
	case time.May:
		return d == 1
	case time.December:
		return d == 25
	default:
		return false
	}
}

// EuroNext calendar of business days 2008.
//  - Jan 1,  New Year Day
//  - Mar 21, Good Friday
//  - Mar 24, Easter Monday
//  - May 1,  Labour Day
//  - Dec 25, Christmas Day
//  - Dec 26, Boxing Day
func euronext2008(m time.Month, d int) bool {
	switch m {
	case time.January:
		return d == 1
	case time.March:
		return d == 21 || d == 24
	case time.May:
		return d == 1
	case time.December:
		return d == 25 || d == 26
	default:
		return false
	}
}

// EuroNext calendar of business days 2007.
//  - Jan 1,  New Year Day
//  - Apr 6,  Good Friday
//  - Apr 9,  Easter Monday
//  - May 1,  Labour Day
//  - Dec 25, Christmas Day
//  - Dec 26, Boxing Day
func euronext2007(m time.Month, d int) bool {
	switch m {
	case time.January:
		return d == 1
	case time.April:
		return d == 6 || d == 9
	case time.May:
		return d == 1
	case time.December:
		return d == 25 || d == 26
	default:
		return false
	}
}

// EuroNext calendar of business days 2006.
//  - Apr 14, Good Friday
//  - Apr 17, Easter Monday
//  - May 1,  Labour Day
//  - Dec 25, Christmas Day
//  - Dec 26, Boxing Day
func euronext2006(m time.Month, d int) bool {
	switch m {
	case time.April:
		return d == 14 || d == 17
	case time.May:
		return d == 1
	case time.December:
		return d == 25 || d == 26
	default:
		return false
	}
}

//nolint:gomnd
// EuroNext calendar of business days 2005.
//  - Mar 25, Good Friday
//  - Mar 28, Easter Monday
//  - Dec 26, Boxing Day
func euronext2005(m time.Month, d int) bool {
	switch m {
	case time.March:
		return d == 25 || d == 28
	case time.December:
		return d == 26
	default:
		return false
	}
}

// EuroNext calendar of business days 2004.
//  - Jan 1,  New Year Day
//  - Apr 9,  Good Friday
//  - Apr 12, Easter Monday
func euronext2004(m time.Month, d int) bool {
	switch m {
	case time.January:
		return d == 1
	case time.April:
		return d == 9 || d == 12
	default:
		return false
	}
}

// EuroNext calendar of business days 2003.
//  - Jan 1,  New Year Day
//  - Apr 18, Good Friday
//  - Apr 21, Easter Monday
//  - May 1,  Labour Day
//  - Dec 25, Christmas Day
//  - Dec 26, Boxing Day
func euronext2003(m time.Month, d int) bool {
	switch m {
	case time.January:
		return d == 1
	case time.April:
		return d == 18 || d == 21
	case time.May:
		return d == 1
	case time.December:
		return d == 25 || d == 26
	default:
		return false
	}
}

//nolint:gomnd
// EuroNext calendar of business days 2002.
//  - Jan 1,  New Year Day
//  - Mar 29, Good Friday
//  - Apr 1,  Easter Monday
//  - May 1,  Labour Day
//  - Dec 25, Christmas Day
//  - Dec 26, Boxing Day
func euronext2002(m time.Month, d int) bool {
	switch m {
	case time.January:
		return d == 1
	case time.March:
		return d == 29
	case time.April:
		return d == 1
	case time.May:
		return d == 1
	case time.December:
		return d == 25 || d == 26
	default:
		return false
	}
}

//nolint:gomnd
// EuroNext calendar of business days 2001.
//  - Jan 1,  New Year Day
//  - Apr 13, Good Friday
//  - Apr 16, Easter Monday
//  - May 1,  Labour Day
//  - Jun 4,  Whit Monday
//  - Dec 25, Christmas Day
//  - Dec 26, Boxing Day
//  - Dec 31, Last trading day - change over to euro
func euronext2001(m time.Month, d int) bool {
	switch m {
	case time.January:
		return d == 1
	case time.April:
		return d == 13 || d == 16
	case time.May:
		return d == 1
	case time.June:
		return d == 4
	case time.December:
		return d == 25 || d == 26 || d == 31
	default:
		return false
	}
}

//nolint:gomnd
// EuroNext calendar of business days 2000.
//  - Apr 21, Good Friday
//  - Apr 24, Easter Monday
//  - May 1,  Labour Day
//  - Jun 12, Whit Monday
//  - Dec 25, Christmas Day
//  - Dec 26, Boxing Day
func euronext2000(m time.Month, d int) bool {
	switch m {
	case time.April:
		return d == 21 || d == 24
	case time.May:
		return d == 1
	case time.June:
		return d == 12
	case time.December:
		return d == 25 || d == 26
	default:
		return false
	}
}

//nolint:gomnd
// EuroNext calendar of business days 1999.
//  - Jan 1,  New Year Day
//  - Apr 2,  Good Friday
//  - Apr 5,  Easter Monday
//  - May 24, Whit Monday
func euronext1999(m time.Month, d int) bool {
	switch m {
	case time.January:
		return d == 1
	case time.April:
		return d == 2 || d == 5
	case time.May:
		return d == 24
	default:
		return false
	}
}
