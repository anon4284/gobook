import "package:angular2/core.dart";
import "dart:html";
import "dart:async";

@Injectable()
class EntryService {
  static Future getTestValue() async {
    var url = "http://localhost:5000/api/test";
    return (await HttpRequest.getString(url));
  }
}
