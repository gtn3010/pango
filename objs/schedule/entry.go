package schedule

import (
	"encoding/xml"

	"github.com/PaloAltoNetworks/pango/util"
	"github.com/PaloAltoNetworks/pango/version"
)

// Entry is a normalized, version independent representation of an address
// object.
type Entry struct {
	Name  string
	Type  string
	Value []string
}

// Copy copies the information from source Entry `s` to this object.  As the
// Name field relates to the XPATH of this object, this field is not copied.
func (o *Entry) Copy(s Entry) {
	o.Name = s.Name
	o.Type = s.Type
	o.Value = s.Value
}

/** Structs / functions for normalization. **/

func (o Entry) Specify(v version.Number) (string, interface{}) {
	_, fn := versioning(v)
	return o.Name, fn(o)
}

type normalizer interface {
	Normalize() []Entry
	Names() []string
}

type container_v1 struct {
	Answer []entry_v1 `xml:"entry"`
}

func (o *container_v1) Names() []string {
	ans := make([]string, 0, len(o.Answer))
	for i := range o.Answer {
		ans = append(ans, o.Answer[i].Name)
	}

	return ans
}

func (o *container_v1) Normalize() []Entry {
	ans := make([]Entry, 0, len(o.Answer))
	for i := range o.Answer {
		ans = append(ans, o.Answer[i].normalize())
	}

	return ans
}

type entry_v1 struct {
	XMLName         xml.Name         `xml:"entry"`
	Name            string           `xml:"name,attr"`
	NonRecurring    *util.MemberType `xml:"schedule-type>non-recurring"`
	WeeklySunday    *util.MemberType `xml:"schedule-type>recurring>weekly>sunday"`
	WeeklyMonday    *util.MemberType `xml:"schedule-type>recurring>weekly>monday"`
	WeeklyTuesday   *util.MemberType `xml:"schedule-type>recurring>weekly>tuesday"`
	WeeklyWednesday *util.MemberType `xml:"schedule-type>recurring>weekly>wednesday"`
	WeeklyThursday  *util.MemberType `xml:"schedule-type>recurring>weekly>thursday"`
	WeeklyFriday    *util.MemberType `xml:"schedule-type>recurring>weekly>friday"`
	WeeklySaturday  *util.MemberType `xml:"schedule-type>recurring>weekly>saturday"`
	Daily           *util.MemberType `xml:"schedule-type>recurring>daily"`
}

func specify_v1(e Entry) interface{} {
	ans := entry_v1{
		Name: e.Name,
	}

	switch e.Type {
	case Once:
		ans.NonRecurring = util.StrToMem(e.Value)
	case EveryDay:
		ans.Daily = util.StrToMem(e.Value)
	case EverySunday:
		ans.WeeklySunday = util.StrToMem(e.Value)
	case EveryMonday:
		ans.WeeklyMonday = util.StrToMem(e.Value)
	case EveryTuesday:
		ans.WeeklyTuesday = util.StrToMem(e.Value)
	case EveryWednesday:
		ans.WeeklyWednesday = util.StrToMem(e.Value)
	case EveryThursday:
		ans.WeeklyThursday = util.StrToMem(e.Value)
	case EveryFriday:
		ans.WeeklyFriday = util.StrToMem(e.Value)
	case EverySaturday:
		ans.WeeklySaturday = util.StrToMem(e.Value)
	}

	return ans
}

func (e *entry_v1) normalize() Entry {
	ans := Entry{
		Name: e.Name,
	}

	if e.NonRecurring != nil {
		ans.Type = Once
		ans.Value = util.MemToStr(e.NonRecurring)
	}
	if e.Daily != nil {
		ans.Type = EveryDay
		ans.Value = util.MemToStr(e.Daily)
	}
	if e.WeeklyMonday != nil {
		ans.Type = EveryMonday
		ans.Value = util.MemToStr(e.WeeklyMonday)
	}
	if e.WeeklyTuesday != nil {
		ans.Type = EveryTuesday
		ans.Value = util.MemToStr(e.WeeklyTuesday)
	}
	if e.WeeklyWednesday != nil {
		ans.Type = EveryWednesday
		ans.Value = util.MemToStr(e.WeeklyWednesday)
	}
	if e.WeeklyThursday != nil {
		ans.Type = EveryThursday
		ans.Value = util.MemToStr(e.WeeklyThursday)
	}
	if e.WeeklyFriday != nil {
		ans.Type = EveryFriday
		ans.Value = util.MemToStr(e.WeeklyFriday)
	}
	if e.WeeklySaturday != nil {
		ans.Type = EverySaturday
		ans.Value = util.MemToStr(e.WeeklySaturday)
	}
	if e.WeeklySunday != nil {
		ans.Type = EverySunday
		ans.Value = util.MemToStr(e.WeeklySunday)
	}

	return ans
}
