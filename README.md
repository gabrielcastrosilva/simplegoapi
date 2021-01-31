# Simple Go API

A simple API that returns the current quotation of the Brazilian Real, United States Dollar and Euro in any choosen coin, developed in Go that communicates with an [external api](exchangeratesapi.io) and is parsed by a Python script inside Docker.


## Instructions

Since the application utilizes docker we must first create the images using the following commands:

Inside API folder:
`docker build -t go-api .`

Inside Runner folder:
`docker build -t py-runner .`

It is of utmost importance for these commands to be ran in their respective folders, otherwise they will **not work**.

Then, the following commands need to be ran in order to:

Start the API docker container: `docker-compose run --rm --name go-api -p 3000:3000 api USD`

Start the Runner docker container: `docker-compose run --rm --name py-runner runner`

Have in mind that they must be ran in that exact order, as the Runner is directly dependant on the API being up. Also it's important to note that every time you reset the API to change the parameter, you will need to run the Runner again aswell to update the message.

On that note, you might have noticed that in the API command it says `USD` in the end, that's the parameter which will change the output in the Runner. The exchangerates API utilizes ISO 4217 codes to identify their currency, so you can feel free to switch the parameter to any of the following codes:

| Code     | Respective Currency   |
| -------- |  :----------------:   | 
| CAD      | Canadian Dollar	     | 
| HKD      | Hong Kong Dollar      | 
| ISK      | Iceland Krona	       |
| PHP      | Philippine Peso       |
| DKK      | Danish Krone          |
| HUF      | Forint                |
| CZK      | Czech Koruna          |
| AUD      | Australian Dollar     |
| RON      | New Romanian Leu      |
| SEK      | Swedish Krona         |
| IDR      | Rupiah                |
| INR      | Indian Rupee          |
| BRL      | Brazilian Real	       |
| RUB      | Russian Ruble         |
| HRK      | Croatian Kuna         |
| JPY      | Yen                   |
| THB      | Baht                  |
| CHF      | Swiss Franc           |
| SGD      | Singapore Dollar      |
| PLN      | Zloty	               |
| BGN      | Bulgarian Lev         |
| TRY      | Turkish Lira          |
| CNY      | Yuan Renminbi         |
| NOK      | Norwegian Krone       |
| NZD      | New Zealand Dollar    |
| ZAR      | Rand                  |
| USD      | US Dollar             |
| MXN      | Mexican Peso          |
| ILS      | New Israeli Sheqel    |
| GBP      | Pound Sterling        |
| KRW      | Won                   |
| MYR      | Malaysian Ringgit     |
