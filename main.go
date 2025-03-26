package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

// WeatherResponse представляет весь JSON-ответ
type WeatherResponse struct {
	Request  Request  `json:"request"`
	Location Location `json:"location"`
	Current  Current  `json:"current"`
}

// Request содержит информацию о запросе
type Request struct {
	Type     string `json:"type"`
	Query    string `json:"query"`
	Language string `json:"language"`
	Unit     string `json:"unit"`
}

// Location содержит информацию о местоположении
type Location struct {
	Name           string `json:"name"`
	Country        string `json:"country"`
	Region         string `json:"region"`
	Lat            string `json:"lat"`
	Lon            string `json:"lon"`
	TimezoneID     string `json:"timezone_id"`
	Localtime      string `json:"localtime"`
	LocaltimeEpoch int64  `json:"localtime_epoch"`
	UTCOffset      string `json:"utc_offset"`
}

// Current содержит текущие погодные данные
type Current struct {
	ObservationTime     string   `json:"observation_time"`
	Temperature         int      `json:"temperature"`
	WeatherCode         int      `json:"weather_code"`
	WeatherIcons        []string `json:"weather_icons"`
	WeatherDescriptions []string `json:"weather_descriptions"`
	Astro               Astro    `json:"astro"`
	WindSpeed           int      `json:"wind_speed"`
	WindDegree          int      `json:"wind_degree"`
	WindDir             string   `json:"wind_dir"`
	Pressure            int      `json:"pressure"`
	Precip              float64  `json:"precip"`
	Humidity            int      `json:"humidity"`
	Cloudcover          int      `json:"cloudcover"`
	Feelslike           int      `json:"feelslike"`
	UVIndex             int      `json:"uv_index"`
	Visibility          int      `json:"visibility"`
	IsDay               string   `json:"is_day"`
}

// Astro содержит астрономические данные (восход, закат и т.д.)
type Astro struct {
	Sunrise          string `json:"sunrise"`
	Sunset           string `json:"sunset"`
	Moonrise         string `json:"moonrise"`
	Moonset          string `json:"moonset"`
	MoonPhase        string `json:"moon_phase"`
	MoonIllumination int    `json:"moon_illumination"`
}

func ExampleRun() {
	data, err := GetInfoWeather("ВАШ_КЛЮЧ", "Russia", "Samara")
	if err != nil {
		log.Fatal(err)
	}
	var weather WeatherResponse
	err = json.Unmarshal(data, &weather)
	if err != nil {
		log.Fatal(err)
	}

	PrintWeatherReport(weather)
}

func GetInfoWeather(key, country, region string) ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.weatherstack.com/current?access_key=%v&query=%v,%v", key, country, region))
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("error get pages")
	}
	defer resp.Body.Close()
	bs := make([]byte, 1014)
	n, err := resp.Body.Read(bs)

	if n == 0 || err != nil {
		return nil, errors.New("error return get pages")
	}
	return bs[:n], nil
}

func PrintWeatherReport(weather WeatherResponse) {
	fmt.Println("=== Отчёт о погоде ===")
	fmt.Println()

	// Общая информация
	fmt.Println("1. Общая информация:")
	fmt.Printf("- Тип запроса: %s\n", weather.Request.Type)
	fmt.Printf("- Местоположение: %s, %s\n", weather.Location.Name, weather.Location.Country)
	fmt.Printf("- Координаты: Широта — %s, Долгота — %s\n", weather.Location.Lat, weather.Location.Lon)
	fmt.Printf("- Временная зона: %s (UTC%s)\n", weather.Location.TimezoneID, weather.Location.UTCOffset)
	fmt.Printf("- Локальное время: %s\n", weather.Location.Localtime)
	fmt.Println()

	// Текущая погода
	fmt.Println("2. Текущая погода:")
	fmt.Printf("- Температура: %d°C\n", weather.Current.Temperature)
	fmt.Printf("- Ощущается как: %d°C\n", weather.Current.Feelslike)
	fmt.Printf("- Погодные условия: %s\n", weather.Current.WeatherDescriptions[0])
	fmt.Printf("- Иконка погоды: %s\n", weather.Current.WeatherIcons[0])
	fmt.Println()

	// Ветер
	fmt.Println("3. Ветер:")
	fmt.Printf("- Скорость ветра: %d км/ч\n", weather.Current.WindSpeed)
	fmt.Printf("- Направление ветра: %s (%d°)\n", weather.Current.WindDir, weather.Current.WindDegree)
	fmt.Println()

	// Атмосферные условия
	fmt.Println("4. Атмосферные условия:")
	fmt.Printf("- Давление: %d мбар\n", weather.Current.Pressure)
	fmt.Printf("- Влажность: %d%%\n", weather.Current.Humidity)
	fmt.Printf("- Облачность: %d%%\n", weather.Current.Cloudcover)
	fmt.Printf("- Видимость: %d км\n", weather.Current.Visibility)
	fmt.Printf("- Осадки: %.1f мм\n", weather.Current.Precip)
	fmt.Println()

	// Астрономические данные
	fmt.Println("5. Астрономические данные:")
	fmt.Printf("- Восход солнца: %s\n", weather.Current.Astro.Sunrise)
	fmt.Printf("- Закат солнца: %s\n", weather.Current.Astro.Sunset)
	fmt.Printf("- Восход луны: %s\n", weather.Current.Astro.Moonrise)
	fmt.Printf("- Заход луны: %s\n", weather.Current.Astro.Moonset)
	fmt.Printf("- Фаза луны: %s\n", weather.Current.Astro.MoonPhase)
	fmt.Printf("- Освещённость луны: %d%%\n", weather.Current.Astro.MoonIllumination)
	fmt.Println()

	// Дополнительные наблюдения
	fmt.Println("7. Дополнительные наблюдения:")
	fmt.Printf("- УФ-индекс: %d\n", weather.Current.UVIndex)
	fmt.Printf("- День или ночь: %s\n", weather.Current.IsDay)
	fmt.Println()

	// Выводы и рекомендации
	fmt.Println("=== Выводы и рекомендации ===")
	fmt.Println("- Погода прохладная, одевайтесь теплее.")
	fmt.Println("- Воздух чистый, можно планировать прогулки.")
	fmt.Println("- День длится примерно 12 часов, наслаждайтесь световым днём.")
}
