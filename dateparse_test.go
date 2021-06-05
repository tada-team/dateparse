package dateparse

import (
	"testing"
	"time"
)

type parserStruct struct {
	date    time.Time
	message string
}

func TestDateparse(t *testing.T) {
	loc, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		t.Fatal("load location fail:", err)
	}
	dt := time.Date(2020, 10, 10, 12, 1, 0, 0, loc) // saturday
	for k, want := range map[string]parserStruct{
		"через час тест": {
			dt.Add(1 * time.Hour),
			"тест",
		},
		"в 15 часов": {
			time.Date(dt.Year(), dt.Month(), dt.Day(), 15, 0, 0, 0, dt.Location()),
			"",
		},
		"в 11 часов": {
			time.Date(dt.Year(), dt.Month(), dt.Day()+1, 11, 0, 0, 0, dt.Location()),
			"",
		},
		"at 11 hours": {
			time.Date(dt.Year(), dt.Month(), dt.Day()+1, 11, 0, 0, 0, dt.Location()),
			"",
		},
		"в 11 комментарий": {
			time.Date(dt.Year(), dt.Month(), dt.Day()+1, 11, 0, 0, 0, dt.Location()),
			"комментарий",
		},
		"сегодня/13:00 почитать": {
			time.Date(dt.Year(), dt.Month(), dt.Day(), 13, 0, 0, 0, dt.Location()),
			"почитать",
		},
		"завтра/13:00": {
			time.Date(dt.Year(), dt.Month(), dt.Day()+1, 13, 0, 0, 0, dt.Location()),
			"",
		},
		"сегодня в полночь": {
			time.Date(dt.Year(), dt.Month(), dt.Day()+1, 0, 0, 0, 0, dt.Location()),
			"",
		},
		"послезавтра днем": {
			time.Date(dt.Year(), dt.Month(), dt.Day()+2, 12, 0, 0, 0, dt.Location()),
			"",
		},
		"at midnight": {
			time.Date(dt.Year(), dt.Month(), dt.Day()+1, 0, 0, 0, 0, dt.Location()),
			"",
		},
		"tomorrow at midnight eat": {
			time.Date(dt.Year(), dt.Month(), dt.Day()+1, 0, 0, 0, 0, dt.Location()),
			"eat",
		},
		"завтра в полночь": {
			time.Date(dt.Year(), dt.Month(), dt.Day()+1, 0, 0, 0, 0, dt.Location()),
			"",
		},
		"послезавтра ночью": {
			time.Date(dt.Year(), dt.Month(), dt.Day()+2, 0, 0, 0, 0, dt.Location()),
			"",
		},
		"сегодня в 22:00 тестируем": {
			time.Date(dt.Year(), dt.Month(), dt.Day(), 22, 0, 0, 0, dt.Location()),
			"тестируем",
		},
		"сегодня в 18:00 умыться": {
			time.Date(dt.Year(), dt.Month(), dt.Day(), 18, 0, 0, 0, dt.Location()),
			"умыться",
		},
		"сегодня в 18": {
			time.Date(dt.Year(), dt.Month(), dt.Day(), 18, 0, 0, 0, dt.Location()),
			"",
		},
		"ровно через год": {
			dt.Add(24 * 365 * time.Hour),
			"",
		},
		"пнуть женю в 16:00": {
			time.Date(dt.Year(), dt.Month(), dt.Day(), 16, 0, 0, 0, dt.Location()),
			"пнуть женю",
		},
		"в 11 утра": {
			time.Date(dt.Year(), dt.Month(), dt.Day()+1, 11, 0, 0, 0, dt.Location()),
			"",
		},
		" в 11 вечера покодить": {
			time.Date(dt.Year(), dt.Month(), dt.Day(), 23, 0, 0, 0, dt.Location()),
			"покодить",
		},
		"в 22:00 спать": {
			time.Date(dt.Year(), dt.Month(), dt.Day(), 22, 0, 0, 0, dt.Location()),
			"спать",
		},
		"в 18:00 декомпозировать": {
			time.Date(dt.Year(), dt.Month(), dt.Day(), 18, 0, 0, 0, dt.Location()),
			"декомпозировать",
		},
		"05.09.2019 в 12:00": {
			time.Date(2019, 9, 5, 12, 0, 0, 0, dt.Location()),
			"",
		},
		"05.09.2019/13:00": {
			time.Date(2019, 9, 5, 13, 0, 0, 0, dt.Location()),
			"",
		},
		"05/26/2021": {
			time.Date(2021, 5, 26, 18, 0, 0, 0, dt.Location()),
			"",
		},
		"05/26/22": {
			time.Date(2022, 5, 26, 18, 0, 0, 0, dt.Location()),
			"",
		},
		"05/26": {
			time.Date(dt.Year()+1, 5, 26, 18, 0, 0, 0, dt.Location()),
			"",
		},
		"26/05/22": {
			time.Date(2022, 5, 26, 18, 0, 0, 0, dt.Location()),
			"",
		},
		"09.07 в 12:00": {
			time.Date(dt.Year()+1, 7, 9, 12, 0, 0, 0, dt.Location()),
			"",
		},
		"1 августа": {
			time.Date(dt.Year()+1, 8, 1, 18, 0, 0, 0, dt.Location()),
			"",
		},
		"23 июня в 9 утра": {
			time.Date(dt.Year()+1, 6, 23, 9, 0, 0, 0, dt.Location()),
			"",
		},
		"1 июля": {
			time.Date(dt.Year()+1, 7, 1, 18, 0, 0, 0, dt.Location()),
			"",
		},
		"31 сентября": {
			time.Date(dt.Year()+1, 9, 31, 18, 0, 0, 0, dt.Location()),
			"",
		},
		"15 ноября выпить молока": {
			time.Date(dt.Year(), 11, 15, 18, 0, 0, 0, dt.Location()),
			"выпить молока",
		},
		"15 октября": {
			time.Date(dt.Year(), 10, 15, 18, 0, 0, 0, dt.Location()),
			"",
		},
		"12/13": {
			time.Date(dt.Year(), 12, 13, 18, 0, 0, 0, dt.Location()),
			"",
		},
		"12:13 проверить что-то": {
			time.Date(dt.Year(), dt.Month(), dt.Day(), 12, 13, 0, 0, dt.Location()),
			"проверить что-то",
		},
		"16.16": {
			time.Date(dt.Year(), dt.Month(), dt.Day(), 16, 16, 0, 0, dt.Location()),
			"",
		},
		"30th morning chil": {
			time.Date(dt.Year(), dt.Month(), 30, 10, 0, 0, 0, dt.Location()),
			"chil",
		},
		"30-го чил": {
			time.Date(dt.Year(), dt.Month(), 30, 18, 0, 0, 0, dt.Location()),
			"чил",
		},
		"18 июня в 14:00 на море": {
			time.Date(dt.Year()+1, 6, 18, 14, 0, 0, 0, dt.Location()),
			"на море",
		},
		"18.03.2019 12:00": {
			time.Date(2019, 3, 18, 12, 0, 0, 0, dt.Location()),
			"",
		},
		"20 марта 2020 года": {
			time.Date(2021, 3, 20, 18, 0, 0, 0, dt.Location()),
			"",
		},
		"20 марта 2021 года купить дошик": {
			time.Date(2021, 3, 20, 18, 0, 0, 0, dt.Location()),
			"купить дошик",
		},
		"20 марта 21 года": {
			time.Date(2021, 3, 20, 18, 0, 0, 0, dt.Location()),
			"",
		},
		"20 мая": {
			time.Date(dt.Year()+1, 5, 20, 18, 0, 0, 0, dt.Location()),
			"",
		},
		"31 мая": {
			time.Date(dt.Year()+1, 5, 31, 18, 0, 0, 0, dt.Location()),
			"",
		},
		"20.11.2019к 9:45 посетить врача": {
			time.Date(2019, 11, 20, 9, 45, 0, 0, dt.Location()),
			"посетить врача",
		},
		"21.09 14:00 чильнуть": {
			time.Date(dt.Year()+1, 9, 21, 14, 0, 0, 0, dt.Location()),
			"чильнуть",
		},
		"23 14:00 поужинать в кафе": {
			time.Date(dt.Year(), dt.Month(), dt.Day(), 14, 0, 0, 0, dt.Location()),
			"23  поужинать в кафе",
		},
		"23-ого в 14:00 срезы": {
			time.Date(dt.Year(), dt.Month(), 23, 14, 0, 0, 0, dt.Location()),
			"срезы",
		},
		"23 числа в 14:00 др начальника": {
			time.Date(dt.Year(), dt.Month(), 23, 14, 0, 0, 0, dt.Location()),
			"др начальника",
		},
		"23/12 в 14 тренировка": {
			time.Date(dt.Year(), 12, 23, 14, 0, 0, 0, dt.Location()),
			"тренировка",
		},
		"23/12 14:00 тренировка 2": {
			time.Date(dt.Year(), 12, 23, 14, 0, 0, 0, dt.Location()),
			"тренировка 2",
		},
		"25.04 в 14 проснуться": {
			time.Date(dt.Year()+1, 4, 25, 14, 0, 0, 0, dt.Location()),
			"проснуться",
		},
		"26.09. 12:00 улыбнуться": {
			time.Date(dt.Year()+1, 9, 26, 12, 0, 0, 0, dt.Location()),
			"улыбнуться",
		},
		"31.05.2019 спектакль": {
			time.Date(2019, 5, 31, 18, 0, 0, 0, dt.Location()),
			"спектакль",
		},
		"7:00 завтрак": {
			time.Date(dt.Year(), dt.Month(), dt.Day()+1, 7, 0, 0, 0, dt.Location()),
			"завтрак",
		},
		"в 10 выпить кофан": {
			time.Date(dt.Year(), dt.Month(), dt.Day()+1, 10, 0, 0, 0, dt.Location()),
			"выпить кофан",
		},
		"в 16:59 позвонить": {
			time.Date(dt.Year(), dt.Month(), dt.Day(), 16, 59, 0, 0, dt.Location()),
			"позвонить",
		},
		"в 1:30": {
			time.Date(dt.Year(), dt.Month(), dt.Day()+1, 1, 30, 0, 0, dt.Location()),
			"",
		},
		"в понедельник убраться": {
			time.Date(dt.Year(), dt.Month(), dt.Day()+2, 18, 0, 0, 0, dt.Location()),
			"убраться",
		},
		"в следующий понедельник": {
			time.Date(dt.Year(), dt.Month(), dt.Day()+9, 18, 0, 0, 0, dt.Location()),
			"",
		},
		"в следующий понедельник утром посмотреть код": {
			time.Date(dt.Year(), dt.Month(), dt.Day()+9, 10, 0, 0, 0, dt.Location()),
			"посмотреть код",
		},
		"в субботу утром": {
			time.Date(dt.Year(), dt.Month(), dt.Day()+7, 10, 0, 0, 0, dt.Location()),
			"",
		},
		"утром": {
			time.Date(dt.Year(), dt.Month(), dt.Day()+1, 10, 0, 0, 0, dt.Location()),
			"",
		},
		"вечером": {
			time.Date(dt.Year(), dt.Month(), dt.Day(), 18, 0, 0, 0, dt.Location()),
			"",
		},
		"понедельник через два дома офис": {
			time.Date(dt.Year(), dt.Month(), dt.Day()+2, 18, 0, 0, 0, dt.Location()),
			"через два дома офис",
		},
		"утром в субботу почистить зубы": {
			time.Date(dt.Year(), dt.Month(), dt.Day()+7, 10, 0, 0, 0, dt.Location()),
			"почистить зубы",
		},
		"в полночь в понедельник помыть голову": {
			time.Date(dt.Year(), dt.Month(), dt.Day()+2, 0, 0, 0, 0, dt.Location()),
			"помыть голову",
		},
		"в субботу": {
			time.Date(dt.Year(), dt.Month(), dt.Day(), 18, 0, 0, 0, dt.Location()),
			"",
		},
		"on monday test": {
			time.Date(dt.Year(), dt.Month(), dt.Day()+2, 18, 0, 0, 0, dt.Location()),
			"test",
		},
		"on monday at 5 p.m": {
			time.Date(dt.Year(), dt.Month(), dt.Day()+2, 17, 0, 0, 0, dt.Location()),
			"",
		},
		"в пятницу в 5 утра": {
			time.Date(dt.Year(), dt.Month(), dt.Day()+6, 5, 0, 0, 0, dt.Location()),
			"",
		},
		"в пятницу в 5 вечера на волгу": {
			time.Date(dt.Year(), dt.Month(), dt.Day()+6, 17, 0, 0, 0, dt.Location()),
			"на волгу",
		},
		"в среду": {
			time.Date(dt.Year(), dt.Month(), dt.Day()+4, 18, 0, 0, 0, dt.Location()),
			"",
		},
		"в пн": {
			time.Date(dt.Year(), dt.Month(), dt.Day()+2, 18, 0, 0, 0, dt.Location()),
			"",
		},
		"пн": {
			time.Date(dt.Year(), dt.Month(), dt.Day()+2, 18, 0, 0, 0, dt.Location()),
			"",
		},
		"вт": {
			time.Date(dt.Year(), dt.Month(), dt.Day()+3, 18, 0, 0, 0, dt.Location()),
			"",
		},
		"ср": {
			time.Date(dt.Year(), dt.Month(), dt.Day()+4, 18, 0, 0, 0, dt.Location()),
			"",
		},
		"чт": {
			time.Date(dt.Year(), dt.Month(), dt.Day()+5, 18, 0, 0, 0, dt.Location()),
			"",
		},
		"пт": {
			time.Date(dt.Year(), dt.Month(), dt.Day()+6, 18, 0, 0, 0, dt.Location()),
			"",
		},
		"сб": {
			time.Date(dt.Year(), dt.Month(), dt.Day(), 18, 0, 0, 0, dt.Location()),
			"",
		},
		"вс": {
			time.Date(dt.Year(), dt.Month(), dt.Day()+1, 18, 0, 0, 0, dt.Location()),
			"",
		},
		"вс тестируем приложение": {
			time.Date(dt.Year(), dt.Month(), dt.Day()+1, 18, 0, 0, 0, dt.Location()),
			"тестируем приложение",
		},
		"в полдень покушать": {
			time.Date(dt.Year(), dt.Month(), dt.Day()+1, 12, 0, 0, 0, dt.Location()),
			"покушать",
		},
		"четверговый митап завтра": {
			time.Date(dt.Year(), dt.Month(), dt.Day()+1, 18, 0, 0, 0, dt.Location()),
			"четверговый митап",
		},
		"что завтра совещание": {
			time.Date(dt.Year(), dt.Month(), dt.Day()+1, 18, 0, 0, 0, dt.Location()),
			"что  совещание",
		},
		"в полночь": {
			time.Date(dt.Year(), dt.Month(), dt.Day()+1, 0, 0, 0, 0, dt.Location()),
			"",
		},
		"среду": {
			time.Date(dt.Year(), dt.Month(), dt.Day()+4, 18, 0, 0, 0, dt.Location()),
			"",
		},
		"в среду в 12:30": {
			time.Date(dt.Year(), dt.Month(), dt.Day()+4, 12, 30, 0, 0, dt.Location()),
			"",
		},
		"в среду утром": {
			time.Date(dt.Year(), dt.Month(), dt.Day()+4, 10, 0, 0, 0, dt.Location()),
			"",
		},
		"во вторник": {
			time.Date(dt.Year(), dt.Month(), dt.Day()+3, 18, 0, 0, 0, dt.Location()),
			"",
		},
		"вчера": {
			time.Date(dt.Year(), dt.Month(), dt.Day()+365, 18, 0, 0, 0, dt.Location()),
			"",
		},
		"завтра в 12": {
			time.Date(dt.Year(), dt.Month(), dt.Day()+1, 12, 0, 0, 0, dt.Location()),
			"",
		},
		"завтра воскресный праздник": {
			time.Date(dt.Year(), dt.Month(), dt.Day()+1, 18, 0, 0, 0, dt.Location()),
			"воскресный праздник",
		},
		"09.12": {
			time.Date(dt.Year(), 12, 9, 18, 0, 0, 0, dt.Location()),
			"",
		},
		"09/12": {
			time.Date(dt.Year(), 12, 9, 18, 0, 0, 0, dt.Location()),
			"",
		},
		"09/12/2050 продать": {
			time.Date(2050, 12, 9, 18, 0, 0, 0, dt.Location()),
			"продать",
		},
		"2020-04-25": {
			time.Date(2021, 4, 25, 18, 0, 0, 0, dt.Location()),
			"",
		},
		"2020-0809": {
			time.Date(2021, 8, 9, 18, 0, 0, 0, dt.Location()),
			"",
		},
		"20-0809": {
			time.Date(2021, 8, 9, 18, 0, 0, 0, dt.Location()),
			"",
		},
		"2020-0809 потеряли дефис": {
			time.Date(2021, 8, 9, 18, 0, 0, 0, dt.Location()),
			"потеряли дефис",
		},
		"2020-04-25 это Iso формат": {
			time.Date(2021, 4, 25, 18, 0, 0, 0, dt.Location()),
			"это iso формат",
		},
		"20-04-25": {
			time.Date(2021, 4, 25, 18, 0, 0, 0, dt.Location()),
			"",
		},
		"200425": {
			time.Date(2021, 4, 25, 18, 0, 0, 0, dt.Location()),
			"",
		},
		"20200425": {
			time.Date(2021, 4, 25, 18, 0, 0, 0, dt.Location()),
			"",
		},
		"20200425 потеряли все": {
			time.Date(2021, 4, 25, 18, 0, 0, 0, dt.Location()),
			"потеряли все",
		},
		"20-04-25 покормить собаку": {
			time.Date(2021, 4, 25, 18, 0, 0, 0, dt.Location()),
			"покормить собаку",
		},
		"10/02/50-го покушать": {
			time.Date(2050, 2, 10, 18, 0, 0, 0, dt.Location()),
			"покушать",
		},
		"сегодня": {
			time.Date(dt.Year(), dt.Month(), dt.Day(), 18, 0, 0, 0, dt.Location()),
			"",
		},
		"сегодня сохранить фото": {
			time.Date(dt.Year(), dt.Month(), dt.Day(), 18, 0, 0, 0, dt.Location()),
			"сохранить фото",
		},
		"30 минут кофе": {
			dt.Add(30 * time.Minute),
			"кофе",
		},
		"завтра": {
			time.Date(dt.Year(), dt.Month(), dt.Day()+1, 18, 0, 0, 0, dt.Location()),
			"",
		},
		"завтра через неделю поход": {
			time.Date(dt.Year(), dt.Month(), dt.Day()+1, 18, 0, 0, 0, dt.Location()),
			"через неделю поход",
		},
		"завтра в среду важное совещание": {
			time.Date(dt.Year(), dt.Month(), dt.Day()+1, 18, 0, 0, 0, dt.Location()),
			"в среду важное совещание",
		},
		"послезавтра помыть голову": {
			time.Date(dt.Year(), dt.Month(), dt.Day()+2, 18, 0, 0, 0, dt.Location()),
			"помыть голову",
		},
		"послепослезавтра": {
			time.Date(dt.Year(), dt.Month(), dt.Day()+3, 18, 0, 0, 0, dt.Location()),
			"",
		},
		"after tomorrow": {
			time.Date(dt.Year(), dt.Month(), dt.Day()+2, 18, 0, 0, 0, dt.Location()),
			"",
		},
		"через 1 минуту моргнуть": {
			dt.Add(1 * time.Minute),
			"моргнуть",
		},
		"in 1 min": {
			dt.Add(1 * time.Minute),
			"",
		},
		"через 3 недели": {
			dt.Add(3 * 7 * 24 * time.Hour),
			"",
		},
		"через 3 неделя": {
			dt.Add(3 * 7 * 24 * time.Hour),
			"",
		},
		"1 неделя": {
			dt.Add(7 * 24 * time.Hour),
			"",
		},
		"1 week wake up": {
			dt.Add(7 * 24 * time.Hour),
			"wake up",
		},
		"через 10 недель": {
			dt.Add(10 * 7 * 24 * time.Hour),
			"",
		},
		"через 30": {
			dt.Add(30 * time.Minute),
			"",
		},
		"через 30 секунд куку": {
			dt.Add(30 * time.Second),
			"куку",
		},
		"через 6 часов": {
			dt.Add(6 * time.Hour),
			"",
		},
		"через 2 часа": {
			dt.Add(2 * time.Hour),
			"",
		},
		"через 7 дней": {
			dt.Add(7 * 24 * time.Hour),
			"",
		},
		"через 2 дня": {
			dt.Add(2 * 24 * time.Hour),
			"",
		},
		"через 100 дней отдохнуть": {
			dt.Add(100 * 24 * time.Hour),
			"отдохнуть",
		},
		"через год": {
			dt.Add(365 * 24 * time.Hour),
			"",
		},
		"через месяц": {
			dt.Add(31 * 24 * time.Hour),
			"",
		},
		"через минуту": {
			dt.Add(1 * time.Minute),
			"",
		},
		"через минуту [username](http://ya.ru)": {
			dt.Add(1 * time.Minute),
			"[username](http://ya.ru)",
		},
		"через 2 минуты": {
			dt.Add(2 * time.Minute),
			"",
		},
		"через неделю выходной": {
			dt.Add(7 * 24 * time.Hour),
			"выходной",
		},
		"через два часа": {
			dt.Add(2 * time.Hour),
			"",
		},
		"через три дня сходить в кино": {
			dt.Add(3 * 24 * time.Hour),
			"сходить в кино",
		},
		"в следующую субботу": {
			time.Date(dt.Year(), dt.Month(), dt.Day()+7, 18, 0, 0, 0, dt.Location()),
			"",
		},
		"в следующую субботу праздник": {
			time.Date(dt.Year(), dt.Month(), dt.Day()+7, 18, 0, 0, 0, dt.Location()),
			"праздник",
		},
		"с утра тест": {
			time.Date(dt.Year(), dt.Month(), dt.Day()+1, 10, 0, 0, 0, dt.Location()),
			"тест",
		},
		"через три минуты тест": {
			dt.Add(time.Minute * 3),
			"тест",
		},
		// FIXME:
		//"в субботу в 11 утра": {
		//	time.Date(dt.Year(), dt.Month(), dt.Day()+7, 11, 0, 0, 0, dt.Location()),
		//	"",
		//},
		//"12 13": {
		//	time.Date(dt.Year(), dt.Month(), dt.Day(), 12, 13, 0, 0, dt.Location()),
		//	"",
		//},
		//"20.11/13:00": {
		//	time.Date(dt.Year(), 11, 20, 13, 0, 0, 0, dt.Location()),
		//	 "",
		//},
		//"завтра про полдень в": {
		//	dt.Add(24 * time.Hour),
		//	"",
		//},
		//"09 00": {
		//	time.Date(dt.Year(), dt.Month(), dt.Day()+1, 9, 0, 0, 0, dt.Location()),
		//	"",
		//},
		//"15": {
		//	time.Date(dt.Year(), dt.Month(), dt.Day(), 15, 0, 0, 0, dt.Location()),
		//	"",
		//},
		//"2": {
		//	time.Date(dt.Year(), dt.Month(), dt.Day()+1, 2, 0, 0, 0, dt.Location()),
		//	"",
		//},
		//"9 14": {
		//	time.Date(dt.Year(), dt.Month(), dt.Day()+1, 9, 14, 0, 0, dt.Location()),
		//	"",
		//},
	} {
		t.Run(k, func(t *testing.T) {
			got, msg := Parse(k, &Opts{Now: dt})
			if got.IsZero() || !got.Equal(want.date) || msg != want.message {
				t.Errorf("dateparse error on '%s': got '%s' (comment: '%s') want '%s' (comment: '%s')", k, got, msg, want.date, want.message)
			}
		})
	}
}

// BenchmarkParse-12    	   15781	     76097 ns/op	     494 B/op	      14 allocs/op
// ==>
// BenchmarkParse-12    	   16088	     75671 ns/op	     439 B/op	       8 allocs/op
func BenchmarkParse(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Parse("сегодня в 18", nil)
	}
}
