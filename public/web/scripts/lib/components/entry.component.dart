import "package:angular2/core.dart";
import "../services/entry.service.dart";
import 'dart:async';

@Component(
  selector: 'gb-entry',
  templateUrl: '/web/tpl/components/entry.tpl.html'
)
class Entry {
  String name;
  String email;
  String content;

  Future send() async {
    String resp = await EntryService.addEntry(email, name, "wdwad");
    print(resp);
  }
}
