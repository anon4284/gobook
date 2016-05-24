import "package:angular2/core.dart";
import "dart:html";
import "dart:async";
import 'dart:convert';

@Injectable()
class EntryService {
  static Future getTestValue() async {
    var url = "/api/test";
    return (await HttpRequest.getString(url));
  }
  static addEntry(String email, String name, String content){
    HttpRequest request = new HttpRequest();

    var url = '/api/entry/add';
    var map = new Map<String,String>();
    map["Email"] = email;
    map["Name"] = name;
    map["Content"] = content;

    request.open("POST", url, async: false);

    var jsonData = JSON.encode(map);

    request.send(jsonData);

    Map resp = JSON.decode(request.response.toString());

    return resp;
  }

  static Future getEntries() async {
    var url = "/api/entry/get";
    return (await HttpRequest.getString(url));
  }

}
