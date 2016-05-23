import "package:angular2/core.dart";
import "dart:html";
import "dart:async";

@Injectable()
class EntryService {
  static Future getTestValue() async {
    var url = "http://localhost:5000/api/test";
    return (await HttpRequest.getString(url));
  }
  static Future addEntry(String email, String name, String content) async {
    var url = '/api/entry/add';
    var map = new Map<String,String>();
    map["email"] = email;
    map["name"] = name;
    map["content"] = content;

    return (await HttpRequest.postFormData(url, map));
  }
}
