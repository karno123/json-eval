# JSON EVAL 
JSON EVAL is simple json evaluator that may be simplify business rule. 

#Operator 
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

#Evaluate
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
```json
glossary.title == "example glossary" && glossary.total > 100
```


