# pj
`pj` is a handy utility for building json strings on the command line. It's meant to complement `jq`, which is excellent for querying json but rather cumbersome when it comes to creating json strings from scratch. The main idea behind `pj` is to use POSIX-compliant argument syntax of the form `--key=value` to build up relatively complex json strings with ease.

## Usage
Json objects are made up of properties of key-value pairs. pj lets you create these properties by first specifying the type with a flag (whether `-s`, `--string`; `-n`, `--number`; or `-b`, `--boolean`) and then listing the keys with their values using the following form:
```
pj --string key1=value key2=value key3=value ... --number --key=value
```
Here's a concrete example:
```
pj -s make=Toyota model=Camry -n year=2021 price=23450 -b isCertifiedPreOwned=true
```
pj produces the following json output, albeit in compact form:
```
{
  "make": "Toyota",
  "model": "Camry",
  "year": 2021,
  "price": 23450,
  "isCertifiedPreOwned": true
}
```

### Stdin
pj accepts json from standard input. Simply pipe your json into pj and then add properties to it. Keys with the same name will overwrite others based on the following precedence: `stdin`, `--json`, `--string`, `--number`, `--boolean` with those coming after overwriting those that came before.

Example file: `weather.json`:
```
{
  "city": "San Diego",
  "temperature": 75.2,
  "unit": "Fahrenheit",
  "isRaining": false,
  "humidity": 55
}
```
Pipe this into pj and overwrite the properties "unit" and "temperature".
```
cat weather.json | pj -s unit=Celsius -n temperature=24
```
json output:
```
{
  "city": "San Diego",
  "temperature": 24,
  "unit": "Celsius",
  "isRaining": false,
  "humidity": 55
}
```

### What about arrays?
For now, pj optimizes for the creation of json objects rather than arrays. However, any json string can be passed in to pj by wrapping it in an object.
```
pj --json '{"numbers": [1, 2, 3, 4, 5, 6, 7]}'
```
