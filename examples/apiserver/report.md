![Alt text](./Group_4.png)  

# **Отчет по тестированию**


## Назначение документа
 
**Основная  цель  данного  документа**  –  предоставить  обработанные  и 
систематизированные  результаты  нагрузочного  тестирования  **Веб-сервера SERVER_NAME_CONFIG**,  описать 
отклонения  при  проведении  тестирования  от  методики  и  ограничения 
тестирования. 

## Основные положения

**Основной  целью  проведения  нагрузочного  тестирования**  **Веб-сервераSERVER_NAME_CONFIG**  является 
оценка  возможностей  системы. 

Для  определения  производительности  **Веб-сервера SERVER_NAME_CONFIG**
планируется: 

1. Определение максимальной производительности (количество запросов/секунду) 
Веб-сервера SERVER_NAME_CONFIG ,  на  существующей  конфигурации  (в  соответствии  с  настроенными 
правилами). 
2. Проверка  надежности  **Веб-сервера SERVER_NAME_CONFIG**  во  время  тестирования  в  течение  длительного 
времени. 
3. Определить  отказоустойчивость  **Веб-сервера SERVER_NAME_CONFIG**  при  отказе  одной  из  площадок 
серверов приложения (останов 2-х из 4-х серверов приложения) 
4. Выявление потенциально «узких» мест **Веб-сервера SERVER_NAME_CONFIG**. 

### Методика тестирования
Для тестирования использовался шаговой алгоритм тестирования в течении 20 минут.

### Конфигурация тестирования
```yaml
CONFIG_DATA
```
### Ограничение тестирования

Ограничений тестирования нет.

## Выводы 

#### Показатели тестирования
**Requests Per Second** --  _RPS_CONFIG (rps)_

**Query Per Second** --     _QPS_CONFIG (qps)_

**Average ResponseTime** -- _ART_CONFIG (ms)_

**Max ResponseTime** --     _MRT_CONFIG (ms)_

**Min ResponseTime** --     _MinRT_CONFIG (ms)_

#### Системные метрики

**CpuLoad** --                 _CPU_CONFIG(%)_

**MemoryFree** --              _MEMORY_FREE_CONFIG(Mb)_

**MemoryCached** --            _MEMORY_CACHED_CONFIG(Mb)_

**DiskLoad (IO interface)** -- _DISK_CONFIG(%)_

### График тестирования
![Alt text](./goog_ltm.png)  
#### Зависимость числа запросов от используемой  памяти
![Alt text](two_axis.png)  
