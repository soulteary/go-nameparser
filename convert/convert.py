from nameparser import HumanName
import json

def Convert(s):
    result = HumanName(s)
    data = {
        "text": str(result),
        "detail": result.as_dict()
    }
    return json.dumps(data)