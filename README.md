# hhApi
Easy api for site hh.ru

> Install 
```
go get -u github.com/YuranIgnatenko/hh_api
```

> Example:
```
// Главная функция
func main() {
  
  // Инициализация структуры
	hh := HH{}
  
  // получение вакансий и зарплвты
  // GetData (city string, vacancy string) []string
  
  // [vacancy 1][][] // is empty deposit
  // [vacancy 2][100 000][] // fixed deposit
  // [vacancy 3][100 000][128 000] //range deposit
  res1 := hh.GetData("Ryazan", "it")
  res2 := hh.GetData("Moscow", "cook")
  
  // отображение списка
  // view ( []string )
	view(res1)
	view(res2)
}
```
