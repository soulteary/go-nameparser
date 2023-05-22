# Go Name Parser

Nameparser API services, python-nameparser implemented by golang.

## Usage

```bash
# curl --request POST 'http://127.0.0.1:8080/api/convert' --header 'Content-Type: application/json' --data-raw '{"name": "Dr. Juan Q. Xavier de la Vega III (Doc Vega)"}'

{"text":"Dr. Juan Q. Xavier de la Vega III (Doc Vega)","detail":{"title":"Dr.","first":"Juan","middle":"Q. Xavier","last":"de la Vega","suffix":"III","nickname":"Doc Vega"}}
```