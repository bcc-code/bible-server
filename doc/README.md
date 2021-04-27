# HTTP API

{% api-method method="get" host="https://bibleapi.bcc.media" path="/v1/:bible/books" %}
{% api-method-summary %}
Get all Books
{% endapi-method-summary %}

{% api-method-description %}
Returns a list of all books in the selected bible
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-path-parameters %}
{% api-method-parameter name="bible" type="string" required=true %}
ID of an available bible
{% endapi-method-parameter %}
{% endapi-method-path-parameters %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}
Cake successfully retrieved.
{% endapi-method-response-example-description %}

```
[{"Number":10,"LongName":"1 Mosebok Genesis","ShortName":"1Mo","ID":""},{"Number":20,"LongName":"2 Mosebok Exodus","ShortName":"2Mo","ID":""},{"Number":30,"LongName":"3 Mosebok Leviticus","ShortName":"3Mo","ID":""},{"Number":40,"LongName":"4 Mosebok Numeri","ShortName":"4Mo","ID":""},{"Number":50,"LongName":"5 Mosebok Deuteronomium","ShortName":"5Mo","ID":""},{"Number":60,"LongName":"Josva","ShortName":"Jos","ID":""},{"Number":70,"LongName":"Dommerne","ShortName":"Dom","ID":""},{"Number":80,"LongName":"Rut","ShortName":"Rut","ID":""},{"Number":90,"LongName":"1 Samuelsbok","ShortName":"1Sa","ID":""},{"Number":100,"LongName":"2 Samuelsbok","ShortName":"2Sa","ID":""},{"Number":110,"LongName":"1 Kongebok","ShortName":"1Ko","ID":""},{"Number":120,"LongName":"2 Kongebok","ShortName":"2Ko","ID":""},{"Number":130,"LongName":"1 Krønikebok","ShortName":"1Kr","ID":""},{"Number":140,"LongName":"2 Krønikebok","ShortName":"2Kr","ID":""},{"Number":150,"LongName":"Esra","ShortName":"Esr","ID":""},{"Number":160,"LongName":"Nehemja","ShortName":"Neh","ID":""},{"Number":190,"LongName":"Ester","ShortName":"Est","ID":""},{"Number":220,"LongName":"Job","ShortName":"Job","ID":""},{"Number":230,"LongName":"Salmenes bok","ShortName":"Sal","ID":""},{"Number":240,"LongName":"Salomos Ordsprog","ShortName":"Ords","ID":""},{"Number":250,"LongName":"Forkynneren","ShortName":"Fork","ID":""},{"Number":260,"LongName":"Salomos Høisang","ShortName":"Høys","ID":""},{"Number":290,"LongName":"Jesaja","ShortName":"Jes","ID":""},{"Number":300,"LongName":"Jeremia","ShortName":"Jer","ID":""},{"Number":310,"LongName":"Klagesangene","ShortName":"Klag","ID":""},{"Number":330,"LongName":"Esekiel","ShortName":"Esek","ID":""},{"Number":340,"LongName":"Daniel","ShortName":"Dan","ID":""},{"Number":350,"LongName":"Hosea","ShortName":"Hos","ID":""},{"Number":360,"LongName":"Joel","ShortName":"Joel","ID":""},{"Number":370,"LongName":"Amos","ShortName":"Am","ID":""},{"Number":380,"LongName":"Obadja","ShortName":"Ob","ID":""},{"Number":390,"LongName":"Jona","ShortName":"Jona","ID":""},{"Number":400,"LongName":"Mika","ShortName":"Mi","ID":""},{"Number":410,"LongName":"Nahum","ShortName":"Nah","ID":""},{"Number":420,"LongName":"Habakkuk","ShortName":"Hab","ID":""},{"Number":430,"LongName":"Sefanja","ShortName":"Sef","ID":""},{"Number":440,"LongName":"Haggai","ShortName":"Hag","ID":""},{"Number":450,"LongName":"Sakarja","ShortName":"Sak","ID":""},{"Number":460,"LongName":"Malaki","ShortName":"Mal","ID":""},{"Number":470,"LongName":"Matteus","ShortName":"Matt","ID":""},{"Number":480,"LongName":"Markus","ShortName":"Mark","ID":""},{"Number":490,"LongName":"Lukas","ShortName":"Luk","ID":""},{"Number":500,"LongName":"Johannes","ShortName":"Joh","ID":""},{"Number":510,"LongName":"Apostlenes gjerninger","ShortName":"Apg","ID":""},{"Number":520,"LongName":"Romerne","ShortName":"Rom","ID":""},{"Number":530,"LongName":"1 korinterne","ShortName":"1Kor","ID":""},{"Number":540,"LongName":"2 korinterne","ShortName":"2Kor","ID":""},{"Number":550,"LongName":"Galaterne","ShortName":"Gal","ID":""},{"Number":560,"LongName":"Efeserne","ShortName":"Ef","ID":""},{"Number":570,"LongName":"Filipperne","ShortName":"Fil","ID":""},{"Number":580,"LongName":"Kolosserne","ShortName":"Kol","ID":""},{"Number":590,"LongName":"1 tessalonikerne","ShortName":"1Ts","ID":""},{"Number":600,"LongName":"2 tessalonikerne","ShortName":"2Ts","ID":""},{"Number":610,"LongName":"1 Timoteus","ShortName":"1Ti","ID":""},{"Number":620,"LongName":"2 Timoteus","ShortName":"2Ti","ID":""},{"Number":630,"LongName":"Titus","ShortName":"Tit","ID":""},{"Number":640,"LongName":"Filemon","ShortName":"Flm","ID":""},{"Number":650,"LongName":"Hebreerne","ShortName":"Hebr","ID":""},{"Number":660,"LongName":"Jakobs","ShortName":"Jak","ID":""},{"Number":670,"LongName":"1 Peters","ShortName":"1Pet","ID":""},{"Number":680,"LongName":"2 Peters","ShortName":"2Pet","ID":""},{"Number":690,"LongName":"1 Johannes","ShortName":"1Joh","ID":""},{"Number":700,"LongName":"2 Johannes","ShortName":"2Joh","ID":""},{"Number":710,"LongName":"3 Johannes","ShortName":"3Joh","ID":""},{"Number":720,"LongName":"Judas","ShortName":"Jud","ID":""},{"Number":730,"LongName":"Johannes’ åpenbaring","ShortName":"Åp","ID":""}]
```
{% endapi-method-response-example %}

{% api-method-response-example httpCode=404 %}
{% api-method-response-example-description %}
Could not find a bible matching this query.
{% endapi-method-response-example-description %}

```
{ "message": "Selected bible not found" }
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

{% api-method method="get" host="https://bibleapi.bcc.media" path="/v1/:bible/:book/:chapter/:verse\_from/:verse\_to" %}
{% api-method-summary %}
Get verse\(s\)
{% endapi-method-summary %}

{% api-method-description %}

{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-path-parameters %}
{% api-method-parameter name="verse\_to" type="integer" required=false %}
Last verse you wish to retrieve. If left empty a single verse will be returned
{% endapi-method-parameter %}

{% api-method-parameter name="verse\_from" type="integer" required=true %}
Number of the starting verse
{% endapi-method-parameter %}

{% api-method-parameter name="chapter" type="integer" required=true %}
Chapter number starting with 1
{% endapi-method-parameter %}

{% api-method-parameter name="book" type="string" required=true %}
Canonical short bible book OD
{% endapi-method-parameter %}

{% api-method-parameter name="bible" type="string" required=true %}
Bible ID
{% endapi-method-parameter %}
{% endapi-method-path-parameters %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}

{% endapi-method-response-example-description %}

```

```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

