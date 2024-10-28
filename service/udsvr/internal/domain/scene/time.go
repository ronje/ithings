package scene

import (
	"context"
	"gitee.com/unitedrhino/share/crons"
	"gitee.com/unitedrhino/share/errors"
	"gitee.com/unitedrhino/share/tools"
	"gitee.com/unitedrhino/share/utils"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

var secondParser = crons.NewParser(crons.Second | crons.Minute | crons.Hour | crons.Dom | crons.Month | crons.DowOptional | crons.Descriptor)

type TimeRangeType = string

const (
	TimeRangeTypeAllDay      TimeRangeType = "allDay"
	TimeRangeTypeLight       TimeRangeType = "light" //todo
	TimeRangeTypeNight       TimeRangeType = "night" //todo
	TimeRangeTypeCustomRange TimeRangeType = "customRange"
)

type DateRangeType = string

const (
	DateRangeTypeAllDay      DateRangeType = "allDay"
	DateRangeTypeWorkDay     DateRangeType = "workday"
	DateRangeTypeWeekend     DateRangeType = "weekend"
	DateRangeTypeHoliday     DateRangeType = "holiday"
	DateRangeTypeCustomRange DateRangeType = "customRange"
	DateRangeTypeCustomWeek  DateRangeType = "customWeek"
)

// TimeRange 时间范围 只支持后面几种特殊字符:*  - ,
type TimeRange struct {
	Type      TimeRangeType `json:"type"`                //时间类型  allDay:全天 light:白天(从日出到日落) night:夜间(从日落到日出) customRange:自定义范围
	StartTime int64         `json:"startTime,omitempty"` //自定义开始时间 从0点加起来的秒数
	EndTime   int64         `json:"endTime,omitempty"`   //自定义结束时间 从0点加起来的秒数
}

type DateRange struct {
	Type      DateRangeType `json:"type"`                //日期类型 workday: 工作日 weekend: 周末 holiday: 节假日  customRange:自定义日期范围  customWeek:自定义周末
	StartDate string        `json:"startDate,omitempty"` //开始日期 2006-01-02
	EndDate   string        `json:"endDate,omitempty"`   //结束日期 2006-01-02
	Repeat    string        `json:"repeat,omitempty"`    //重复 二进制周一到周日 11111111 这个参数只有定时触发才有
}

type RepeatType = string

const (
	RepeatTypeWeek  RepeatType = "week"  //按周重复(默认)
	RepeatTypeMount RepeatType = "mount" //按月重复
)

var RepeatTypeLen = map[string]int{
	RepeatTypeWeek:  7,
	RepeatTypeMount: 31,
}

type ExecType = string

const (
	ExecTypeAt       ExecType = "at"       //在时间点执行(默认)
	ExecTypeSunRises ExecType = "sunRises" //太阳升起
	ExecTypeSunSet   ExecType = "sunSet"   //太阳落下
	ExecTypeLoop     ExecType = "loop"     //间隔时间执行
)

//type TimeUnit string
//
//const (
//	TimeUnitSeconds TimeUnit = "seconds"
//	TimeUnitMinutes TimeUnit = "minutes"
//	TimeUnitHours   TimeUnit = "hours"
//)
//
//type UnitTime struct {
//	Time int64    `json:"time"` //延迟时间
//	Unit TimeUnit `json:"unit"` //时间单位 second:秒  minute:分钟  hour:小时  week:星期 month:月
//}

func (t *TimeRange) Validate() error {
	if t == nil {
		return errors.Parameter.AddMsg("时间范围需要填写时间内容")
	}
	if !utils.SliceIn(t.Type, TimeRangeTypeAllDay, TimeRangeTypeCustomRange, TimeRangeTypeLight, TimeRangeTypeNight) {
		return errors.Parameter.AddMsg("时间范围类型不正确")
	}
	if t.Type == TimeRangeTypeCustomRange {
		if t.StartTime < 0 || t.StartTime > 24*60*60 || t.EndTime < 0 || t.EndTime > t.StartTime+24*60*60 || t.StartTime > t.EndTime {
			return errors.Parameter.AddMsg("自定义时间范围只能在0到24小时之间")
		}
	}
	return nil
}

func (t *DateRange) Validate() error {
	if t == nil {
		return errors.Parameter.AddMsg("日期范围需要填写日期内容")
	}
	if !utils.SliceIn(t.Type, DateRangeTypeAllDay, DateRangeTypeWorkDay, DateRangeTypeWeekend, DateRangeTypeHoliday, DateRangeTypeCustomRange, DateRangeTypeCustomWeek) {
		return errors.Parameter.AddMsg("日期范围类型不正确")
	}
	if t.Type == DateRangeTypeCustomRange {
		start := utils.FmtNilDateStr(t.StartDate)
		if start == nil {
			return errors.Parameter.AddMsg("日期范围开始时间的格式为:2006-01-02")
		}
		end := utils.FmtNilDateStr(t.EndDate)
		if end == nil {
			return errors.Parameter.AddMsg("日期范围结束时间的格式为:2006-01-02")
		}
	}
	repeat := cast.ToInt64(t.Repeat)
	if t.Type == DateRangeTypeCustomWeek && repeat == 0 || repeat > 1111111 {
		return errors.Parameter.AddMsg("日期范围重复只能在0 7个二进制为高位")
	}
	return nil
}

func (d *DateRange) IsHit(ctx context.Context, t time.Time, repo CheckRepo) bool {
	if d.Type == DateRangeTypeAllDay {
		return true
	}
	if d.Type == DateRangeTypeCustomRange {
		start := utils.FmtDateStr(d.StartDate)
		end := utils.FmtDateStr(d.EndDate)
		if t.Before(end) && t.After(start) {
			return true
		}
		return false
	}
	switch d.Type {
	case DateRangeTypeAllDay:
		return true
	case DateRangeTypeWorkDay, DateRangeTypeHoliday:
		h, err := tools.GetHoliday(ctx, t)
		if err != nil {
			logx.WithContext(ctx).Error(err)
			return false
		}
		if d.Type == DateRangeTypeWorkDay {
			return h.Holiday == tools.HolidayWorkDay
		} else {
			return h.Holiday == tools.HolidayFestival
		}
	case DateRangeTypeWeekend:
		weekday := time.Now().Weekday()
		if utils.SliceIn(weekday, time.Sunday, time.Saturday) {
			return true
		}
	case DateRangeTypeCustomRange:
		start := utils.FmtDateStr(d.StartDate)
		end := utils.FmtDateStr(d.EndDate)
		if t.Before(end) && t.After(start) {
			return true
		}
		return false
	case DateRangeTypeCustomWeek:
		weekday := time.Now().Weekday()
		if weekday == 0 {
			weekday = 6
		} else {
			weekday--
		}
		repeat := utils.BStrToInt64(d.Repeat)
		if repeat&(1<<weekday) > 0 {
			return true
		}
	}
	return false
}

func (d *TimeRange) IsHit(ctx context.Context, t time.Time, repo CheckRepo) bool {
	switch d.Type {
	case TimeRangeTypeAllDay:
		return true
	case TimeRangeTypeCustomRange:
		now := utils.TimeToDaySec(t)
		if now >= d.StartTime && now <= d.EndTime {
			return true
		}
		return false
	}
	return false
}
