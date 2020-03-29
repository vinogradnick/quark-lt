# Задачи по курсовой



### Действия для системы сборки 
- [ ] Запуск тестирования
- [ ] Выбор агента
- [ ] Выбор сценария тестирования
- [ ] Загрузка конфигурационного файла

### Действия для Менеджера
- [ ] Просмотр отчетов
- [ ] Моделирование тестирования


#### Действия тестировщика
- [ ] Добавление сценариев
- [ ] Выбор алгоритма
- [ ] Выбор объекта тестирования
- [ ] Загрузка файла с сайта




#### Действия Администратора
- [ ] Просмотр истории тестировани
- [ ] Добавление квоты
- [ ] Изменение квоты
- [ ] Добавление рабочих нод
- [ ] Изменение рабочих нод



```go

	go func() {
		for {
			select {
			case <-worker.StatusChan:
				fmt.Println("done")
				worker.Wg.Done()
				worker.WorkerCtx()
				return
			case data := <-worker.MetricChan:

				r, err := json.Marshal(data)
				fmt.Println(r)
				res, err := http.Post(worker.ExportUrl, "application/json", bytes.NewBuffer(r))
				if res.StatusCode != 200 {
					fmt.Println(res.StatusCode)
					close(worker.StatusChan)
				}
				log.Println("send")
				worker.RpcClient.Send(*data)


			}
		}
	}()
```
