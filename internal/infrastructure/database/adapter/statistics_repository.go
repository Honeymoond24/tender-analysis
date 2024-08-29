package adapter

import (
	"context"
	"fmt"
	"github.com/Honeymoond24/tender-analysis/internal/application"
)

type StatisticsRepository struct {
	DbPool *DBPool
}

func NewStatisticsRepository(db *DBPool) application.StatisticsRepository {
	return &StatisticsRepository{DbPool: db}
}

// MostActiveCategoryByTenders returns the category with the most active tenders
// and the number of active tenders in this category
func (s *StatisticsRepository) MostActiveCategoryByTenders() (category string, count int) {
	row := s.DbPool.conn.QueryRow(context.Background(),
		`SELECT count(a.id) AS count, l.ktru_name
		FROM announces AS a
		         JOIN lot_announces AS l ON a.id = l.announce_id
		WHERE COALESCE(a.open_start_date, a.offers_start_date) > now()
		GROUP BY l.ktru_name
		ORDER BY count DESC
		LIMIT 1`)
	err := row.Scan(&count, &category)
	if err != nil {
		return
	}
	return
}

func (s *StatisticsRepository) MostActiveCategoryByPriceSum() (category string, sum float64) {
	row := s.DbPool.conn.QueryRow(context.Background(),
		`SELECT SUM(l.planned_amount) AS tender_sum, l.ktru_name
			FROM lot_announces AS l
					 LEFT JOIN announces AS a ON a.id = l.announce_id
			WHERE (a.open_start_date IS NOT NULL AND a.open_start_date > now())
			   OR (a.open_start_date IS NULL AND a.offers_start_date > now())
				AND l.planned_amount IS NOT NULL
			GROUP BY l.ktru_name
			ORDER BY tender_sum DESC
			LIMIT 1;`)
	err := row.Scan(&sum, &category)
	if err != nil {
		return
	}
	return
}

func (s *StatisticsRepository) ActiveTenders() (count int) {
	row := s.DbPool.conn.QueryRow(context.Background(),
		`SELECT count(1) FROM announces AS a
            WHERE COALESCE(a.open_start_date, a.offers_start_date) > now()`)
	err := row.Scan(&count)
	if err != nil {
		return 0
	}
	return count
}

func (s *StatisticsRepository) CategorySumsCounts() (result []application.CategorySumsCount) {
	rows, err := s.DbPool.conn.Query(context.Background(),
		`SELECT l.ktru_name AS category,
				   SUM(a.sum)  AS category_sum,
				   COUNT(a.id) AS tenders_count
		FROM announces AS a
				 JOIN lot_announces AS l ON a.id = l.announce_id
		WHERE a.sum IS NOT NULL
		GROUP BY l.ktru_name
		ORDER BY tenders_count DESC, category_sum DESC
		LIMIT 20;`)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var category string
		var sum float64
		var count int
		err = rows.Scan(&category, &sum, &count)
		if err != nil {
			fmt.Println("CategorySumsCounts error:", err)
			return
		}
		result = append(result, application.CategorySumsCount{
			Category: category,
			Sum:      sum,
			Count:    count,
		})
	}
	return
}

func (s *StatisticsRepository) MonthsWithMoreTendersThanAverage() (result []application.TendersPerMonth) {
	rows, err := s.DbPool.conn.Query(context.Background(),
		`WITH YearMonths AS (SELECT count(a.id)                                                          AS tenders_count,
                           EXTRACT(YEAR FROM COALESCE(a.offers_start_date, a.open_start_date))  AS year,
                           EXTRACT(MONTH FROM COALESCE(a.offers_start_date, a.open_start_date)) AS month
                    FROM announces AS a
                    WHERE COALESCE(a.offers_start_date, a.open_start_date) IS NOT NULL
                    GROUP BY year, month
                    ORDER BY year, month),
     AveragePerMonth AS (SELECT AVG(ym.tenders_count)
                         FROM YearMonths AS ym)
		SELECT ym.tenders_count,
			   ym.year,
			   ym.month
		FROM YearMonths AS ym
		WHERE ym.tenders_count > (SELECT * FROM AveragePerMonth) * 1.1;`)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var tendersCount int
		var year int
		var month int
		err = rows.Scan(&tendersCount, &year, &month)
		if err != nil {
			return
		}
		result = append(result, application.TendersPerMonth{
			TendersCount: tendersCount,
			Year:         year,
			Month:        month,
		})
	}
	return
}

func (s *StatisticsRepository) DiagramByDate(params application.Params) (result []application.DiagramDataPerMonth) {
	var queryArguments []interface{}
	query := `SELECT count(a.id)`
	if params.ShowSum {
		query += `, sum(a.sum)`
	}
	query += `,
			CASE
				WHEN a.offers_start_date IS NOT NULL THEN
					EXTRACT(YEAR FROM a.offers_start_date)
				ELSE
					EXTRACT(YEAR FROM a.open_start_date)
				END AS year,
		   	CASE
				WHEN a.offers_start_date IS NOT NULL THEN
					EXTRACT(MONTH FROM a.offers_start_date)
				ELSE
					EXTRACT(MONTH FROM a.open_start_date)
				END AS month
		FROM announces AS a`
	if len(params.KeyWords) != 0 || params.CategoryCode != "" {
		query += ` LEFT JOIN lot_announces AS l ON a.id = l.announce_id`
	}
	query += ` WHERE (a.offers_start_date IS NOT NULL OR a.open_start_date IS NOT NULL)`
	if params.SumRangeFrom > 0 {
		queryArguments = append(queryArguments, params.SumRangeFrom)
		query += fmt.Sprintf(` AND a.sum > $%v `, len(queryArguments))
	} else {
		query += ` AND a.sum > 0 `
	}
	if params.SumRangeTo > 0 {
		queryArguments = append(queryArguments, params.SumRangeTo)
		query += fmt.Sprintf(` AND a.sum < $%v`, len(queryArguments))
	}
	if params.CategoryCode != "" {
		queryArguments = append(queryArguments, params.CategoryCode)
		query += fmt.Sprintf(` AND l.ktru_code = $%v`, len(queryArguments))
	}
	if len(params.KeyWords) != 0 {
		for _, keyWord := range params.KeyWords {
			queryArguments = append(queryArguments, fmt.Sprintf("%%%v%", keyWord))
			query += fmt.Sprintf(` AND l.ktru_name LIKE $%v`, len(queryArguments))
		}
	}
	query += ` GROUP BY year, month 
			   ORDER BY year, month;`
	fmt.Println("Final query:", query)
	rows, err := s.DbPool.conn.Query(context.Background(), query, queryArguments...)
	if err != nil {
		fmt.Println("DiagramByDate error:", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var tendersCount int
		var tendersSum float64
		var year int
		var month int
		var destination []interface{}
		if params.ShowSum {
			destination = []interface{}{&tendersCount, &tendersSum, &year, &month}
		} else {
			destination = []interface{}{&tendersCount, &year, &month}
		}
		err = rows.Scan(destination...)
		if err != nil {
			fmt.Println("DiagramByDate rows.Scan error:", err) // TODO: log with logger
			return
		}
		result = append(result, application.DiagramDataPerMonth{
			TendersCount: tendersCount,
			TendersSum:   tendersSum,
			Year:         year,
			Month:        month,
		})
	}
	return
}
