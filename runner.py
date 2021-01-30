import urllib.request
import json
import iso4217parse

link = "http://localhost:3000/"
f = urllib.request.urlopen(link)
myfile = f.read()

rates = json.loads(myfile)

rawData = iso4217parse.parse(str(rates['base']))

print(rawData[0][3][0])

coin = rawData[0][2]
symbol = rawData[0][3][0]

print("Moeda escolhida: " + coin)
print(symbol + " 1,00 se equivale a:")
print("€ " + str(rates['rates']['EUR']) + " Euros")
print("$ " + str(rates['rates']['USD']) + " Dolares Americanos")
print("R$ " + str(rates['rates']['BRL']) + " Reais brasileiros")

