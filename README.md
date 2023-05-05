# JSON EVAL 
JSON EVAL is simple json evaluator that may be simplify business rule. 

# Operator 
Operator supported 
| No.      | Operator  |
| ---      | ---       |
| 1        | <         |
| 2        | <=        |
| 3        | >         |
| 4        | >=        |
| 5        | !=        |
| 6        | \|\|      |
| 7        | &&        |

# Evaluate
json-eval can evaluate json value with logical expression.

We have below json:

```json 
{
    "glossary": {
        "title": "example glossary",
        "GlossDiv": {
            "total": 1000
            "title": "S",
            "GlossList": {
                "GlossEntry": {
                    "ID": "SGML",
                    "SortAs": "SGML",
                    "GlossTerm": "Standard Generalized Markup Language",
                    "Acronym": "SGML",
                    "Abbrev": "ISO 8879:1986",
                    "GlossDef": {
                        "para": "A meta-markup language, used to create markup languages such as DocBook.",
                        "GlossSeeAlso": [
                            "GML",
                            "XML"
                        ]
                    },
                    "GlossSee": "markup"
                }
            }
        }
    }
}
```
expression: 

~~~
glossary.title == "example glossary" && glossary.total > 100
~~~
array of index, json eval also has capability to extract array value from json. for example, 
~~~
glossary.GlossDiv.GlossList.GlossEntry.GlossDef.GlossSeeAlso[0]
~~~

Result:
```go 
"GML"
```

# Example
Code snippest
```go
jsnStr := `
    {
        "glossary": {
            "title": "example glossary",
            "GlossDiv": {
                "total": 1000
            }
        }
    }    
`
expression := `glossary.title == "example glossary" && glossary.total > 100`
mapEval := jsoneval.NewJsonEvaluator()
result, err := mapEval.EvaluateJson(expression, jsonStr)
if err != nil {
  fmt.Println(err)
} else {
  fmt.Println(result)
}
```

Result:
```go 
true 
```


