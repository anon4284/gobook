import "package:angular2/core.dart";
import "../services/entry.service.dart";
import 'dart:async';
import '../services/flashMsg.service.dart';
import 'dart:convert';
import 'dart:html';

@Component(
  selector: 'gb-entry',
  templateUrl: 'entry.tpl.html'
)
class EntryComponent implements OnInit{
  String name;
  String email;
  String content;
  Map entries;

  void send() {
    Map resp = EntryService.addEntry(email, name, content);

    FlashMsgService.msg(resp);

    if(resp["Valid"]) {
      this.addEntryToDomAfterAdd(resp["Problem"]);
    }
  }

  Future getEntries() async {
    String resp = await EntryService.getEntries();
    Map m = JSON.decode(resp);

    for(var i=0; i < m["Count"]; i++) {
      this.addEntryToDom(m["Items"][i]);
    }
  }


  void addEntryToDom(Map entry) {
    var tr = new TableRowElement();
    var tdName = new TableCellElement();
    var tdYear = new TableCellElement();
    var tdContent = new TableCellElement();

    tdName.text = entry["Name"]["S"];
    tdYear.text = entry["Year"]["N"];
    tdContent.text = entry["Content"]["S"];

    tr.children.add(tdName);
    tr.children.add(tdYear);
    tr.children.add(tdContent);



   querySelector("#entryTable").children.add(tr);



  }


  void addEntryToDomAfterAdd(String year) {
    var tr = new TableRowElement();
    var tdName = new TableCellElement();
    var tdYear = new TableCellElement();
    var tdContent = new TableCellElement();

    tdName.text = this.name;
    tdYear.text = year;
    tdContent.text = this.content;

    tr.children.add(tdName);
    tr.children.add(tdYear);
    tr.children.add(tdContent);



   querySelector("#entryTable").children.add(tr);



  }




  @override
  ngOnInit() {
    getEntries();
  }
}
