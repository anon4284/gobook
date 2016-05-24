import "package:angular2/core.dart";
import "services/entry.service.dart";
import 'dart:async';
import 'partials/navbar.component.dart';
import 'components/entry.component.dart';

@Component(
  selector: 'goprofile',
  templateUrl:'app.tpl.html',
  directives: const [Navbar, EntryComponent]
)
class AppComponent implements OnInit {

  Future printTestValue() async {
    String test = await EntryService.getTestValue();
    print(test);
  }


  @override
  ngOnInit() {
    printTestValue();
  }
}
