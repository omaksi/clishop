package rdg

import (
	"database/sql"
	"time"

	"ondrejmaksi.com/db2project/db"
)

/*
Používateľ zadá na vstupe počiatočný rok. Pre každý mesiac v danom a nasledovných rokoch (januar 2018, februar 2018, …) vypočítajte, aký je pomer medzi tými vyhľadávacími dopytmi, ktoré obsahujú názov produktu a tými, ktorý názov produktu neobsahujú. Vo výstupe musia byť mesiace, pre ktoré nie sú v databáze dáta. Ak daný pomer nie je pre daný mesiac definovaný, vo výstupe má byť NULL.
*/
type SearchStatistic struct {
	Month time.Time
	Share sql.NullFloat64
}

func rowsToSearchStatistics(rows *sql.Rows) []SearchStatistic {
	res := []SearchStatistic{}
	for rows.Next() {
		ss := SearchStatistic{}
		err := rows.Scan(&ss.Month, &ss.Share)
		if err != nil {
			panic(err)
		}
		res = append(res, ss)
	}

	return res
}

func GetSearchStatistics(year string) []SearchStatistic {
	sql := `
	with months as (
		select * from generate_series($1, now()::date, '1 month'::interval) months
	),
	total_searches as (
		select months, count(ts.searched_at) total from months
		join search_log_timestamps ts on ts.searched_at between months::timestamp and months::timestamp + INTERVAL '1 month'
		group by months
	),
	searches_with_name as (
		select months, count(p.name) total from months
		join search_log_timestamps ts on ts.searched_at between months::timestamp and months::timestamp + INTERVAL '1 month'
		join search_log l on ts.search_log_id = l.id
		join products p on l.query like '%' || p.name || '%'
		group by months
	)
	select m.months, /*ts.total total, ns.total with_name,*/ ns.total::decimal / ts.total::decimal pomer from months m
	left join total_searches ts on ts.months = m.months
	left join searches_with_name ns on ns.months = m.months
	;
	`
	rows, err := db.GetDatabase().Query(sql, year+"-01-01")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	return rowsToSearchStatistics(rows)
}
