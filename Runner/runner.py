################################################
#                                              #
#          Simple Python Runner                #
#                                              #
################################################

#Initial imports
import urllib.request
import json
import iso4217parse

# Declaring API url
link = "http://host.docker.internal:3000/"

# Fetching data from API through request
f = urllib.request.urlopen(link)
myfile = f.read()

# Parsing data from API's JSON
rates = json.loads(myfile)
rawData = iso4217parse.parse(str(rates['base']))

# Declaring important data for usage in prints
coin = rawData[0][2]
symbol = rawData[0][3][0]

# Prints showcasing parsed data from API
print("Moeda escolhida: " + coin)
print(symbol + " 1,00 se equivale a:")
print("€ " + str(rates['rates']['EUR']) + " Euros")
print("$ " + str(rates['rates']['USD']) + " Dolares Americanos")
print("R$ " + str(rates['rates']['BRL']) + " Reais brasileiros")

