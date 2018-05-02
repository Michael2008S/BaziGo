package BaziGo

import (
	. "github.com/warrially/BaziGo/Common"
	"github.com/warrially/BaziGo/DaYun"
	"github.com/warrially/BaziGo/JieQi"
	"github.com/warrially/BaziGo/LiChun"
	"github.com/warrially/BaziGo/SiZhu"
	"log"
)

// 八字
type TBazi struct {
	SolarDate TDate   // 新历日期
	BaziDate  TDate   // 八字日期
	SiZhu     TSiZhu  // 四柱
	XiYong    TXiYong // 喜用神
	DaYun     TDaYun  // 大运
}

// 从新历获取八字(年, 月, 日, 时, 分, 秒)
func GetBazi(nYear, nMonth, nDay, nHour, nMinute, nSecond, nSex int) TBazi {
	var bazi TBazi

	// 新历年
	bazi.SolarDate.Year = nYear
	bazi.SolarDate.Month = nMonth
	bazi.SolarDate.Day = nDay
	bazi.SolarDate.Hour = nHour
	bazi.SolarDate.Minute = nMinute
	bazi.SolarDate.Second = nSecond

	// 通过立春获取当年的年份
	bazi.BaziDate.Year = LiChun.GetLiChun(bazi.SolarDate)
	// 通过节气获取当年的月份
	bazi.BaziDate.JieQi = JieQi.GetJieQi(bazi.SolarDate)
	// 节气0 是立春 是1月
	bazi.BaziDate.Month = bazi.BaziDate.JieQi/2 + 1

	// 通过八字年来获取年柱
	bazi.SiZhu.YearZhu = SiZhu.GetZhuFromYear(bazi.BaziDate.Year)
	// 通过年干支和八字月
	bazi.SiZhu.MonthZhu = SiZhu.GetZhuFromMonth(bazi.BaziDate.Month, bazi.SiZhu.YearZhu.Gan.Value)
	// 通过公历 年月日计算日柱
	bazi.SiZhu.DayZhu = SiZhu.GetZhuFromDay(nYear, nMonth, nDay)
	//
	bazi.SiZhu.HourZhu = SiZhu.GetZhuFromHour(nHour, bazi.SiZhu.DayZhu.Gan.Value)

	// 计算十神
	SiZhu.CalcShiShen(&bazi.SiZhu)
	// 计算纳音
	SiZhu.CalcNaYin(&bazi.SiZhu)

	// 计算大运
	bazi.DaYun = DaYun.CalcDaYun(&bazi.SiZhu, nSex)

	// 计算喜用神
	bazi.XiYong = SiZhu.CalcXiYong(&bazi.SiZhu)

	return bazi
}

// 从农历获取八字
func GetBaziFromLunar() {

}

// 按照特殊格式化输出(未完成)
func PrintBazi(bazi TBazi) {
	log.Println("======================================================================")
	log.Println("出生日期新历： ", bazi.SolarDate.Year, "年",
		bazi.SolarDate.Month, "月",
		bazi.SolarDate.Day, "日  ",
		bazi.SolarDate.Hour, ":",
		bazi.SolarDate.Minute, ":",
		bazi.SolarDate.Second,
	)
	log.Println("基本八字：",
		bazi.SiZhu.YearZhu.GanZhi.Str,
		bazi.SiZhu.MonthZhu.GanZhi.Str,
		bazi.SiZhu.DayZhu.GanZhi.Str,
		bazi.SiZhu.HourZhu.GanZhi.Str)

	log.Println("命盘解析：")
	log.Println(
		bazi.SiZhu.YearZhu.Gan.Str, "(",
		bazi.SiZhu.YearZhu.Gan.WuXing.Str, ")[",
		bazi.SiZhu.YearZhu.Gan.ShiShen.Str, "]\t",
		bazi.SiZhu.MonthZhu.Gan.Str, "(",
		bazi.SiZhu.MonthZhu.Gan.WuXing.Str, ")[",
		bazi.SiZhu.MonthZhu.Gan.ShiShen.Str, "]\t",
		bazi.SiZhu.DayZhu.Gan.Str, "(",
		bazi.SiZhu.DayZhu.Gan.WuXing.Str, ")[",
		bazi.SiZhu.DayZhu.Gan.ShiShen.Str, "]\t",
		bazi.SiZhu.HourZhu.Gan.Str, "(",
		bazi.SiZhu.HourZhu.Gan.WuXing.Str, ")[",
		bazi.SiZhu.HourZhu.Gan.ShiShen.Str, "]")
	log.Println(
		bazi.SiZhu.YearZhu.Zhi.Str, "(",
		bazi.SiZhu.YearZhu.Zhi.WuXing.Str, ")",
		bazi.SiZhu.YearZhu.Zhi.CangGan[0].ShiShen.Str,
		bazi.SiZhu.YearZhu.Zhi.CangGan[1].ShiShen.Str,
		bazi.SiZhu.YearZhu.Zhi.CangGan[2].ShiShen.Str,
		"\t",
		bazi.SiZhu.MonthZhu.Zhi.Str, "(",
		bazi.SiZhu.MonthZhu.Zhi.WuXing.Str, ")",
		bazi.SiZhu.MonthZhu.Zhi.CangGan[0].ShiShen.Str,
		bazi.SiZhu.MonthZhu.Zhi.CangGan[1].ShiShen.Str,
		bazi.SiZhu.MonthZhu.Zhi.CangGan[2].ShiShen.Str,
		"\t",
		bazi.SiZhu.DayZhu.Zhi.Str, "(",
		bazi.SiZhu.DayZhu.Zhi.WuXing.Str, ")",
		bazi.SiZhu.DayZhu.Zhi.CangGan[0].ShiShen.Str,
		bazi.SiZhu.DayZhu.Zhi.CangGan[1].ShiShen.Str,
		bazi.SiZhu.DayZhu.Zhi.CangGan[2].ShiShen.Str,
		"\t",
		bazi.SiZhu.HourZhu.Zhi.Str, "(",
		bazi.SiZhu.HourZhu.Zhi.WuXing.Str, ")",
		bazi.SiZhu.HourZhu.Zhi.CangGan[0].ShiShen.Str,
		bazi.SiZhu.HourZhu.Zhi.CangGan[1].ShiShen.Str,
		bazi.SiZhu.HourZhu.Zhi.CangGan[2].ShiShen.Str)

	log.Println(
		bazi.SiZhu.YearZhu.GanZhi.NaYin.Str, "\t\t",
		bazi.SiZhu.MonthZhu.GanZhi.NaYin.Str, "\t\t",
		bazi.SiZhu.DayZhu.GanZhi.NaYin.Str, "\t\t",
		bazi.SiZhu.HourZhu.GanZhi.NaYin.Str, "\t\t",
	)

	log.Println("======================================================================")

}
